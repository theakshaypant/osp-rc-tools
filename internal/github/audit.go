package github

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/v88/github"
	"github.com/theakshaypant/osp-rc-tools/internal/jira"
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

func ParseMinorVersion(version string) (majorMinor string, prevMajorMinor string, err error) {
	matches := patchVersionRe.FindStringSubmatch(version)
	if len(matches) != 3 || matches[2] != "0" {
		return "", "", fmt.Errorf("invalid minor version format: %s (expected X.Y.0)", version)
	}

	majorMinor = matches[1]
	parts := strings.SplitN(majorMinor, ".", 2)
	major, _ := strconv.Atoi(parts[0])
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", "", fmt.Errorf("invalid minor number in version %s: %w", version, err)
	}
	if minor < 1 {
		return "", "", fmt.Errorf("minor version must be >= 1, got %s", version)
	}

	prevMajorMinor = fmt.Sprintf("%d.%d", major, minor-1)
	return majorMinor, prevMajorMinor, nil
}

type ProgressFunc func(format string, args ...any)
type ResultFunc func(*AuditResult)

func buildAuditResult(version, prevVersion string, bounds *PatchBounds, released bool, components []ComponentCommits) *AuditResult {
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
		Components: components,
	}
}

// GetPatchCommits fetches all component commits for a patch release.
// Handles both released patches (bounded by bot commits) and unreleased
// patches (uses latest head SHA as the upper bound).
func GetPatchCommits(ctx context.Context, client *github.Client, jiraClient *jira.Client, version string, toDateOverride *time.Time, progress ProgressFunc, onResult ResultFunc) (*AuditResult, error) {
	if progress == nil {
		progress = func(string, ...any) {}
	}

	majorMinor, prevVersion, err := ParsePatchVersion(version)
	if err != nil {
		return nil, err
	}

	progress("Fetching release config for %s...", majorMinor)
	releaseCfg, err := FetchReleaseConfig(ctx, client, majorMinor, "main")
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

		cc, err := getComponentCommits(ctx, client, jiraClient, component, branchCfg, downstreamBranch, downstreamBranch, bounds, released, true, toDateOverride, version, prevVersion, progress)
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

		if onResult != nil {
			onResult(buildAuditResult(version, prevVersion, bounds, released, results))
		}
	}

	return buildAuditResult(version, prevVersion, bounds, released, results), nil
}

func GetMinorCommits(ctx context.Context, client *github.Client, jiraClient *jira.Client, version string, toDateOverride *time.Time, progress ProgressFunc, onResult ResultFunc) (*AuditResult, error) {
	if progress == nil {
		progress = func(string, ...any) {}
	}

	majorMinor, prevMajorMinor, err := ParseMinorVersion(version)
	if err != nil {
		return nil, err
	}
	prevVersion := prevMajorMinor + ".0"

	progress("Fetching release config for %s...", majorMinor)
	releaseCfg, err := FetchReleaseConfig(ctx, client, majorMinor, "main")
	if err != nil {
		return nil, err
	}

	progress("Fetching release config for %s (previous minor)...", prevMajorMinor)
	prevReleaseCfg, err := FetchReleaseConfig(ctx, client, prevMajorMinor, "main")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch previous minor release config: %w", err)
	}

	progress("Fetching patch history for %s...", prevMajorMinor)
	prevHistory, err := GetPatchHistory(ctx, client, prevMajorMinor)
	if err != nil {
		return nil, err
	}
	prevBounds, err := boundsFromHistory(prevHistory, prevVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to find %s in patch history: %w", prevVersion, err)
	}

	progress("Fetching patch history for %s...", majorMinor)
	currentHistory, err := GetPatchHistory(ctx, client, majorMinor)
	if err != nil {
		return nil, err
	}
	currentBounds, err := boundsFromHistory(currentHistory, version)
	if err != nil {
		return nil, fmt.Errorf("failed to find %s in patch history: %w", version, err)
	}

	fromTime := prevBounds.From
	if !prevBounds.To.IsZero() {
		fromTime = prevBounds.To
	}
	bounds := &PatchBounds{
		From: fromTime,
		To:   currentBounds.To,
	}

	if toDateOverride != nil {
		bounds.To = *toDateOverride
	}

	released := !bounds.To.IsZero()
	if released {
		progress("  From: %s (end of %s)", fromTime.Format("2006-01-02 15:04:05 MST"), prevVersion)
		progress("  To:   %s", bounds.To.Format("2006-01-02 15:04:05 MST"))
	} else {
		progress("  From: %s (end of %s)", fromTime.Format("2006-01-02 15:04:05 MST"), prevVersion)
		progress("  To:   (unreleased — using latest HEAD)")
	}

	fromBranch := fmt.Sprintf("release-v%s.x", prevMajorMinor)
	toBranch := fmt.Sprintf("release-v%s.x", majorMinor)
	var results []ComponentCommits

	for component, branchCfg := range releaseCfg.Branches {
		prevBranchCfg, exists := prevReleaseCfg.Branches[component]
		if !exists {
			progress("  %s: new in %s, skipping (no %s baseline)", component, majorMinor, prevMajorMinor)
			results = append(results, ComponentCommits{
				Name:  component,
				Error: fmt.Sprintf("new component in %s, no %s baseline", majorMinor, prevMajorMinor),
			})
			continue
		}

		progress("Processing %s...", component)

		cc, err := getComponentCommits(ctx, client, jiraClient, component, branchCfg, fromBranch, toBranch, bounds, released, false, toDateOverride, version, prevVersion, progress)
		if err != nil {
			results = append(results, ComponentCommits{
				Name:  component,
				Error: err.Error(),
			})
			continue
		}

		cc.PrevUpstreamBranch = prevBranchCfg.Upstream
		progress("  %s: %d commits (%s → %s)", component, len(cc.Commits), prevBranchCfg.Upstream, branchCfg.Upstream)
		if len(cc.UnsyncedCommits) > 0 {
			progress("  ⚠ %s: %d unsynced commits on %s (downstream HEAD %s, upstream HEAD %s)",
				component, len(cc.UnsyncedCommits), cc.UpstreamBranch, cc.ToSHA[:12], cc.UpstreamHead[:12])
		}
		results = append(results, *cc)

		if onResult != nil {
			onResult(buildAuditResult(version, prevVersion, bounds, released, results))
		}
	}

	return buildAuditResult(version, prevVersion, bounds, released, results), nil
}

func getComponentCommits(ctx context.Context, client *github.Client, jiraClient *jira.Client, component string, branchCfg BranchConfig, fromBranch, toBranch string, bounds *PatchBounds, released bool, filterByDate bool, toDateOverride *time.Time, version, prevVersion string, progress ProgressFunc) (*ComponentCommits, error) {
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
		DownstreamBranch: toBranch,
	}

	fromSHA, err := resolveHeadSHA(ctx, client, component, repoCfg.Repo, fmt.Sprintf("v%s", prevVersion), fromBranch, &bounds.From, progress)
	if err != nil {
		return nil, fmt.Errorf("failed to get head SHA at from-date: %w", err)
	}
	cc.FromSHA = fromSHA

	var toSHA string
	if released {
		toSHA, err = resolveHeadSHA(ctx, client, component, repoCfg.Repo, fmt.Sprintf("v%s", version), toBranch, &bounds.To, progress)
	} else {
		toSHA, err = GetHeadSHAAt(ctx, client, repoCfg.Repo, toBranch, nil)
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
		if filterByDate {
			fromDateStr := bounds.From.Format("2006-01-02")
			var filtered []Commit
			for _, c := range commits {
				if c.Date >= fromDateStr {
					filtered = append(filtered, c)
				}
			}
			commits = filtered
		}
		cc.Commits = commits
		if len(commits) > 0 {
			ResolvePRs(ctx, client, upstreamParts[0], upstreamParts[1], cc.Commits)
			if filterByDate {
				ResolveOriginalPRs(ctx, client, upstreamParts[0], upstreamParts[1], cc.Commits)
			}
			resolveJiras(jiraClient, cc.Commits)
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
				if filterByDate {
					ResolveOriginalPRs(ctx, client, upstreamParts[0], upstreamParts[1], unsynced)
				}
				resolveJiras(jiraClient, unsynced)
				cc.UnsyncedCommits = unsynced
			}
		}
	}

	return &cc, nil
}

func resolveJiras(jiraClient *jira.Client, commits []Commit) {
	if jiraClient == nil {
		return
	}
	for i := range commits {
		if commits[i].PR != "" {
			if ticket := jiraClient.FindTicketForPR(commits[i].PR); ticket != "" {
				commits[i].Jira = ticket
				continue
			}
		}
		if commits[i].OriginalPR != "" {
			if ticket := jiraClient.FindTicketForPR(commits[i].OriginalPR); ticket != "" {
				commits[i].Jira = ticket
			}
		}
	}
}

func resolveHeadSHA(ctx context.Context, client *github.Client, component, repo, tag, branch string, fallbackTime *time.Time, progress ProgressFunc) (string, error) {
	sha, err := readHeadFileAt(ctx, client, repo, tag)
	if err == nil {
		progress("  %s: resolved tag %s", component, tag)
		return sha, nil
	}
	progress("  %s: tag %s not found, using time-based resolution", component, tag)
	return GetHeadSHAAt(ctx, client, repo, branch, fallbackTime)
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
