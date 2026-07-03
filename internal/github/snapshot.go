package github

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/google/go-github/v88/github"
	"github.com/theakshaypant/osp-rc-tools/internal/jira"
)

type SnapshotComponent struct {
	Name     string `json:"name"`
	Repo     string `json:"repo"`
	Revision string `json:"revision"`
}

type ChangedRepo struct {
	Repo        string
	OldRevision string
	NewRevision string
	Components  []string
}

var snapshotMMRe = regexp.MustCompile(`-(\d+)-(\d+)-\d{8}-`)

func FetchSnapshot(server, token, namespace, name string) ([]SnapshotComponent, error) {
	cmd := exec.Command("oc", "get", "snapshot", name,
		"-n", namespace,
		"--server="+server,
		"--token="+token,
		"--insecure-skip-tls-verify",
		"-o", "json",
	)
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("oc get snapshot %s failed: %s", name, string(exitErr.Stderr))
		}
		return nil, fmt.Errorf("oc get snapshot %s failed: %w", name, err)
	}

	var snapshot struct {
		Spec struct {
			Components []struct {
				Name   string `json:"name"`
				Source struct {
					Git struct {
						URL      string `json:"url"`
						Revision string `json:"revision"`
					} `json:"git"`
				} `json:"source"`
			} `json:"components"`
		} `json:"spec"`
	}
	if err := json.Unmarshal(out, &snapshot); err != nil {
		return nil, fmt.Errorf("failed to parse snapshot %s: %w", name, err)
	}

	var components []SnapshotComponent
	for _, c := range snapshot.Spec.Components {
		repo := strings.TrimSuffix(c.Source.Git.URL, ".git")
		repo = strings.TrimPrefix(repo, "https://github.com/")
		components = append(components, SnapshotComponent{
			Name:     c.Name,
			Repo:     repo,
			Revision: c.Source.Git.Revision,
		})
	}
	return components, nil
}

func DiffSnapshots(oldComps, newComps []SnapshotComponent) (changed []ChangedRepo, unchanged []string) {
	oldByRepo := map[string]string{}
	for _, c := range oldComps {
		oldByRepo[c.Repo] = c.Revision
	}

	newByRepo := map[string]struct {
		revision   string
		components []string
	}{}
	for _, c := range newComps {
		entry := newByRepo[c.Repo]
		entry.revision = c.Revision
		entry.components = append(entry.components, c.Name)
		newByRepo[c.Repo] = entry
	}

	for repo, entry := range newByRepo {
		oldRev, exists := oldByRepo[repo]
		if !exists || oldRev != entry.revision {
			changed = append(changed, ChangedRepo{
				Repo:        repo,
				OldRevision: oldRev,
				NewRevision: entry.revision,
				Components:  entry.components,
			})
		} else {
			unchanged = append(unchanged, repo)
		}
	}
	return changed, unchanged
}

func ExtractMajorMinor(snapshotName string) (string, error) {
	m := snapshotMMRe.FindStringSubmatch(snapshotName)
	if len(m) != 3 {
		return "", fmt.Errorf("cannot extract version from snapshot name: %s", snapshotName)
	}
	return m[1] + "." + m[2], nil
}

func GetSnapshotDiffCommits(ctx context.Context, client *github.Client, jiraClient *jira.Client, server, token, namespace, oldSnap, newSnap, version string, progress ProgressFunc, onResult ResultFunc) (*AuditResult, error) {
	if progress == nil {
		progress = func(string, ...any) {}
	}

	majorMinor, err := ExtractMajorMinor(newSnap)
	if err != nil {
		return nil, err
	}

	progress("Fetching old snapshot: %s", oldSnap)
	oldComps, err := FetchSnapshot(server, token, namespace, oldSnap)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch old snapshot: %w", err)
	}
	progress("  %d components", len(oldComps))

	progress("Fetching new snapshot: %s", newSnap)
	newComps, err := FetchSnapshot(server, token, namespace, newSnap)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch new snapshot: %w", err)
	}
	progress("  %d components", len(newComps))

	changed, unchanged := DiffSnapshots(oldComps, newComps)
	progress("%d repos changed, %d unchanged", len(changed), len(unchanged))

	progress("Fetching release config for %s...", majorMinor)
	releaseCfg, err := FetchReleaseConfig(ctx, client, majorMinor, "main")
	if err != nil {
		return nil, err
	}

	if version == "" {
		version = releaseCfg.EffectiveVersion()
		progress("Derived version: %s", version)
	}

	downstreamBranch := fmt.Sprintf("release-v%s.x", majorMinor)
	repoMap := buildRepoToComponentMap(ctx, client, releaseCfg)

	var results []ComponentCommits

	for _, cr := range changed {
		mapping, ok := repoMap[cr.Repo]
		if !ok {
			progress("  %s: not in release config, skipping", cr.Repo)
			continue
		}
		comp := mapping.Component

		progress("Processing %s (%s)...", comp, cr.Repo)

		repoCfg, err := FetchRepoConfig(ctx, client, comp)
		if err != nil {
			results = append(results, ComponentCommits{
				Name:  comp,
				Error: fmt.Sprintf("failed to fetch repo config: %v", err),
			})
			continue
		}

		upstreamParts := strings.SplitN(repoCfg.Upstream, "/", 2)
		if len(upstreamParts) != 2 {
			results = append(results, ComponentCommits{
				Name:  comp,
				Error: fmt.Sprintf("invalid upstream format: %s", repoCfg.Upstream),
			})
			continue
		}

		cc := ComponentCommits{
			Name:             comp,
			Upstream:         repoCfg.Upstream,
			Downstream:       cr.Repo,
			UpstreamBranch:   mapping.Branch.Upstream,
			DownstreamBranch: downstreamBranch,
		}

		if cr.OldRevision == "" {
			progress("  %s: new component (not in old snapshot)", comp)
			cc.Error = "new component — not present in old snapshot"
			results = append(results, cc)
			continue
		}

		oldUpstreamSHA, err := readHeadFileAt(ctx, client, repoName(cr.Repo), cr.OldRevision)
		if err != nil {
			progress("  %s: failed to read head at old revision: %v", comp, err)
			cc.Error = fmt.Sprintf("failed to read head file at old revision %s: %v", cr.OldRevision[:min(12, len(cr.OldRevision))], err)
			results = append(results, cc)
			continue
		}
		cc.FromSHA = oldUpstreamSHA

		newUpstreamSHA, err := readHeadFileAt(ctx, client, repoName(cr.Repo), cr.NewRevision)
		if err != nil {
			progress("  %s: failed to read head at new revision: %v", comp, err)
			cc.Error = fmt.Sprintf("failed to read head file at new revision %s: %v", cr.NewRevision[:min(12, len(cr.NewRevision))], err)
			results = append(results, cc)
			continue
		}
		cc.ToSHA = newUpstreamSHA

		if oldUpstreamSHA == newUpstreamSHA {
			progress("  %s: upstream unchanged (downstream-only changes)", comp)
			cc.Commits = []Commit{}
			results = append(results, cc)
			if onResult != nil {
				onResult(buildSnapshotResult(version, oldSnap, results))
			}
			continue
		}

		commits, err := GetCommitsBetweenRefs(ctx, client, upstreamParts[0], upstreamParts[1], oldUpstreamSHA, newUpstreamSHA)
		if err != nil {
			cc.Error = fmt.Sprintf("failed to compare commits: %v", err)
			results = append(results, cc)
			continue
		}
		cc.Commits = commits

		if len(commits) > 0 {
			ResolvePRs(ctx, client, upstreamParts[0], upstreamParts[1], cc.Commits)
			ResolveOriginalPRs(ctx, client, upstreamParts[0], upstreamParts[1], cc.Commits)
			resolveJiras(jiraClient, cc.Commits)
		}

		progress("  %s: %d commits (%s → %s)", comp, len(cc.Commits), oldUpstreamSHA[:min(12, len(oldUpstreamSHA))], newUpstreamSHA[:min(12, len(newUpstreamSHA))])
		results = append(results, cc)

		if onResult != nil {
			onResult(buildSnapshotResult(version, oldSnap, results))
		}
	}

	return buildSnapshotResult(version, oldSnap, results), nil
}

func buildSnapshotResult(version, oldSnap string, components []ComponentCommits) *AuditResult {
	return &AuditResult{
		Version:    version,
		Previous:   oldSnap,
		Components: components,
	}
}

type repoMapping struct {
	Component string
	Branch    BranchConfig
}

func buildRepoToComponentMap(ctx context.Context, client *github.Client, releaseCfg *ReleaseConfig) map[string]repoMapping {
	result := map[string]repoMapping{}
	for component, branchCfg := range releaseCfg.Branches {
		repoCfg, err := FetchRepoConfig(ctx, client, component)
		if err != nil {
			continue
		}
		// The hack repo config uses either `name` or `repo` for the downstream
		// repo name. `name` is the primary field; `repo` is an override used
		// when the downstream repo has a different name (e.g. p12n-* prefixed).
		repoName := repoCfg.Name
		if repoName == "" {
			repoName = repoCfg.Repo
		}
		downstream := fmt.Sprintf("%s/%s", downstreamOrg, repoName)
		result[downstream] = repoMapping{Component: component, Branch: branchCfg}
	}
	return result
}

func repoName(fullRepo string) string {
	parts := strings.SplitN(fullRepo, "/", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return fullRepo
}
