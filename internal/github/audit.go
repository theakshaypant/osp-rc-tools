package github

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/v88/github"
)

// ParsePatchVersion parses "1.22.3" into majorMinor="1.22" and previous="1.22.2".
var patchVersionRe = regexp.MustCompile(`^(\d+\.\d+)\.(\d+)$`)

func ParsePatchVersion(version string) (majorMinor string, prevVersion string, err error) {
	matches := patchVersionRe.FindStringSubmatch(version)
	if len(matches) != 3 {
		return "", "", fmt.Errorf("invalid patch version format: %s (expected X.Y.Z)", version)
	}

	majorMinor = matches[1]
	patch, err := strconv.Atoi(matches[2])
	if err != nil {
		return "", "", fmt.Errorf("invalid patch number in version %s: %w", version, err)
	}
	if patch < 1 {
		return "", "", fmt.Errorf("patch version must be >= 1, got %s", version)
	}

	prevVersion = fmt.Sprintf("%s.%d", majorMinor, patch-1)
	return majorMinor, prevVersion, nil
}

type ProgressFunc func(format string, args ...any)

// GetPatchCommits fetches all component commits for a patch release.
// Handles both released patches (bounded by bot commits) and unreleased
// patches (uses latest head SHA as the upper bound).
func GetPatchCommits(ctx context.Context, client *github.Client, version string, toDateOverride *time.Time, progress ProgressFunc) (*AuditResult, error) {
	if progress == nil {
		progress = func(string, ...any) {}
	}

	majorMinor, prevVersion, err := ParsePatchVersion(version)
	if err != nil {
		return nil, err
	}

	progress("Fetching release config for %s...", majorMinor)
	releaseCfg, err := FetchReleaseConfig(ctx, client, majorMinor)
	if err != nil {
		return nil, err
	}

	progress("Fetching patch history for %s...", majorMinor)
	history, err := GetPatchHistory(ctx, client, majorMinor)
	if err != nil {
		return nil, err
	}

	bounds, err := boundsFromHistory(history, version)
	if err != nil {
		return nil, err
	}

	if toDateOverride != nil {
		bounds.To = *toDateOverride
	}

	released := !bounds.To.IsZero()
	if released {
		progress("  From: %s", bounds.From.Format("2006-01-02 15:04:05 MST"))
		progress("  To:   %s", bounds.To.Format("2006-01-02 15:04:05 MST"))
	} else {
		progress("  From: %s", bounds.From.Format("2006-01-02 15:04:05 MST"))
		progress("  To:   (unreleased — using latest HEAD)")
	}

	downstreamBranch := fmt.Sprintf("release-v%s.x", majorMinor)
	var results []ComponentCommits

	for component, branchCfg := range releaseCfg.Branches {
		progress("Processing %s...", component)

		cc, err := getComponentCommits(ctx, client, component, branchCfg, downstreamBranch, bounds, released, toDateOverride)
		if err != nil {
			results = append(results, ComponentCommits{
				Name:  component,
				Error: err.Error(),
			})
			continue
		}

		progress("  %s: %d commits", component, len(cc.Commits))
		if len(cc.UnsyncedCommits) > 0 {
			progress("  ⚠ %s: %d unsynced commits on %s (downstream HEAD %s, upstream HEAD %s)",
				component, len(cc.UnsyncedCommits), cc.UpstreamBranch, cc.ToSHA[:12], cc.UpstreamHead[:12])
		}
		results = append(results, *cc)
	}

	toDate := bounds.To
	if toDate.IsZero() {
		toDate = time.Now()
	}

	return &AuditResult{
		Version:    version,
		Previous:   prevVersion,
		FromDate:   bounds.From.Format("2006-01-02T15:04:05Z"),
		ToDate:     toDate.Format("2006-01-02T15:04:05Z"),
		Released:   released,
		Components: results,
	}, nil
}

func getComponentCommits(ctx context.Context, client *github.Client, component string, branchCfg BranchConfig, downstreamBranch string, bounds *PatchBounds, released bool, toDateOverride *time.Time) (*ComponentCommits, error) {
	repoCfg, err := FetchRepoConfig(ctx, client, component)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch repo config: %w", err)
	}

	upstreamParts := strings.SplitN(repoCfg.Upstream, "/", 2)
	if len(upstreamParts) != 2 {
		return nil, fmt.Errorf("invalid upstream format: %s", repoCfg.Upstream)
	}

	cc := ComponentCommits{
		Name:             component,
		Upstream:         repoCfg.Upstream,
		Downstream:       fmt.Sprintf("openshift-pipelines/%s", repoCfg.Repo),
		UpstreamBranch:   branchCfg.Upstream,
		DownstreamBranch: downstreamBranch,
	}

	fromSHA, err := GetHeadSHAAt(ctx, client, repoCfg.Repo, downstreamBranch, &bounds.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get head SHA at from-date: %w", err)
	}
	cc.FromSHA = fromSHA

	var toSHA string
	if released {
		toSHA, err = GetHeadSHAAt(ctx, client, repoCfg.Repo, downstreamBranch, &bounds.To)
	} else {
		toSHA, err = GetHeadSHAAt(ctx, client, repoCfg.Repo, downstreamBranch, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get head SHA at to-date: %w", err)
	}
	cc.ToSHA = toSHA

	if fromSHA != toSHA {
		commits, err := GetCommitsBetweenRefs(ctx, client, upstreamParts[0], upstreamParts[1], fromSHA, toSHA)
		if err != nil {
			return nil, fmt.Errorf("failed to compare commits: %w", err)
		}
		fromDateStr := bounds.From.Format("2006-01-02")
		var filtered []Commit
		for _, c := range commits {
			if c.Date >= fromDateStr {
				filtered = append(filtered, c)
			}
		}
		cc.Commits = filtered
		if len(filtered) > 0 {
			ResolvePRs(ctx, client, upstreamParts[0], upstreamParts[1], cc.Commits)
		}
	} else {
		cc.Commits = []Commit{}
	}

	var upstreamHead string
	if toDateOverride != nil {
		upstreamHead, err = GetBranchHeadAt(ctx, client, upstreamParts[0], upstreamParts[1], branchCfg.Upstream, *toDateOverride)
	} else {
		upstreamHead, err = GetBranchHead(ctx, client, upstreamParts[0], upstreamParts[1], branchCfg.Upstream)
	}
	if err == nil {
		cc.UpstreamHead = upstreamHead
		if upstreamHead != toSHA {
			unsynced, err := GetCommitsBetweenRefs(ctx, client, upstreamParts[0], upstreamParts[1], toSHA, upstreamHead)
			if err == nil {
				ResolvePRs(ctx, client, upstreamParts[0], upstreamParts[1], unsynced)
				cc.UnsyncedCommits = unsynced
			}
		}
	}

	return &cc, nil
}

func boundsFromHistory(history []PatchVersionEntry, targetVersion string) (*PatchBounds, error) {
	for i, entry := range history {
		if entry.Version == targetVersion {
			bounds := &PatchBounds{From: entry.Date}
			if i+1 < len(history) {
				bounds.To = history[i+1].Date
			}
			return bounds, nil
		}
	}

	return nil, fmt.Errorf("could not find version %s in patch history", targetVersion)
}
