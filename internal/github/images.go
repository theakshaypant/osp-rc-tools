package github

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/google/go-github/v88/github"
	"github.com/theakshaypant/osp-rc-tools/internal/jira"
	"gopkg.in/yaml.v3"
)

type ImageEntry struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type ImageLabels struct {
	UpstreamRepo string
	UpstreamRef  string
}

type projectYAML struct {
	Images []ImageEntry `yaml:"images"`
}

func FetchProjectYAML(ctx context.Context, client *github.Client, commit string) ([]ImageEntry, error) {
	content, _, _, err := client.Repositories.GetContents(ctx, downstreamOrg, "operator", "project.yaml", &github.RepositoryContentGetOptions{Ref: commit})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch project.yaml at %s: %w", commit, err)
	}
	raw, err := content.GetContent()
	if err != nil {
		return nil, fmt.Errorf("failed to decode project.yaml: %w", err)
	}

	var proj projectYAML
	if err := yaml.Unmarshal([]byte(raw), &proj); err != nil {
		return nil, fmt.Errorf("failed to parse project.yaml: %w", err)
	}
	return proj.Images, nil
}

func DeduplicateImages(entries []ImageEntry) []ImageEntry {
	seen := map[string]bool{}
	var result []ImageEntry
	for _, e := range entries {
		if !seen[e.Value] {
			seen[e.Value] = true
			result = append(result, e)
		}
	}
	return result
}

func InspectImage(imageRef string) (*ImageLabels, error) {
	cmd := exec.Command("skopeo", "inspect", "--no-creds", "docker://"+imageRef)
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("skopeo inspect failed: %s", string(exitErr.Stderr))
		}
		return nil, fmt.Errorf("skopeo inspect failed: %w", err)
	}

	var result struct {
		Labels map[string]string `json:"Labels"`
	}
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, fmt.Errorf("failed to parse skopeo output: %w", err)
	}

	repo := result.Labels["upstream-vcs-location"]
	ref := result.Labels["upstream-vcs-ref"]
	if repo == "" || ref == "" {
		return nil, fmt.Errorf("missing upstream-vcs labels (location=%q, ref=%q)", repo, ref)
	}

	return &ImageLabels{
		UpstreamRepo: normalizeRepoURL(repo),
		UpstreamRef:  ref,
	}, nil
}

func normalizeRepoURL(u string) string {
	u = strings.TrimPrefix(u, "https://github.com/")
	u = strings.TrimPrefix(u, "http://github.com/")
	u = strings.TrimSuffix(u, ".git")
	u = strings.TrimSuffix(u, "/")
	return u
}

func GetImageDiffCommits(ctx context.Context, client *github.Client, jiraClient *jira.Client, server, token, namespace, operatorCommit, snapshot, version string, progress ProgressFunc, onResult ResultFunc) (*AuditResult, error) {
	if progress == nil {
		progress = func(string, ...any) {}
	}

	majorMinor, err := ExtractMajorMinor(snapshot)
	if err != nil {
		return nil, err
	}

	progress("Fetching project.yaml at %s...", operatorCommit)
	images, err := FetchProjectYAML(ctx, client, operatorCommit)
	if err != nil {
		return nil, err
	}
	progress("  %d image entries", len(images))

	unique := DeduplicateImages(images)
	progress("  %d unique images (after dedup)", len(unique))

	progress("Inspecting %d images with skopeo (parallel)...", len(unique))
	type inspectResult struct {
		index  int
		labels *ImageLabels
		err    error
		ref    string
	}
	inspections := make(chan inspectResult, len(unique))
	sem := make(chan struct{}, 8)
	var wg sync.WaitGroup
	for i, img := range unique {
		wg.Add(1)
		go func(idx int, entry ImageEntry) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			labels, err := InspectImage(entry.Value)
			shortRef := entry.Value
			if i := strings.LastIndex(shortRef, "/"); i >= 0 {
				shortRef = shortRef[i+1:]
			}
			if len(shortRef) > 60 {
				shortRef = shortRef[:60] + "..."
			}
			inspections <- inspectResult{index: idx, labels: labels, err: err, ref: shortRef}
		}(i, img)
	}
	go func() { wg.Wait(); close(inspections) }()

	oldUpstream := map[string]string{}
	done := 0
	for r := range inspections {
		done++
		if r.err != nil {
			progress("  [%d/%d] %s ⚠ %v", done, len(unique), r.ref, r.err)
			continue
		}
		oldUpstream[r.labels.UpstreamRepo] = r.labels.UpstreamRef
		progress("  [%d/%d] %s → %s @ %s", done, len(unique), r.ref, r.labels.UpstreamRepo, r.labels.UpstreamRef[:min(12, len(r.labels.UpstreamRef))])
	}
	progress("%d upstream repos resolved from images", len(oldUpstream))

	progress("Fetching snapshot: %s", snapshot)
	snapComps, err := FetchSnapshot(server, token, namespace, snapshot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch snapshot: %w", err)
	}
	progress("  %d components", len(snapComps))

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

	// Group snapshot components by downstream repo
	type repoSnap struct {
		repo     string
		revision string
	}
	snapByRepo := map[string]repoSnap{}
	for _, c := range snapComps {
		if _, exists := snapByRepo[c.Repo]; !exists {
			snapByRepo[c.Repo] = repoSnap{repo: c.Repo, revision: c.Revision}
		}
	}

	var results []ComponentCommits

	for downstreamRepo, snap := range snapByRepo {
		mapping, ok := repoMap[downstreamRepo]
		if !ok {
			continue
		}
		comp := mapping.Component

		repoCfg, err := FetchRepoConfig(ctx, client, comp)
		if err != nil {
			continue
		}

		upstreamParts := strings.SplitN(repoCfg.Upstream, "/", 2)
		if len(upstreamParts) != 2 {
			continue
		}
		upstreamRepo := repoCfg.Upstream

		cc := ComponentCommits{
			Name:             comp,
			Upstream:         upstreamRepo,
			Downstream:       downstreamRepo,
			UpstreamBranch:   mapping.Branch.Upstream,
			DownstreamBranch: downstreamBranch,
		}

		newUpstreamSHA, err := readHeadFileAt(ctx, client, repoName(downstreamRepo), snap.revision)
		if err != nil {
			progress("  %s: failed to read head file: %v", comp, err)
			cc.Error = fmt.Sprintf("failed to read head file: %v", err)
			results = append(results, cc)
			continue
		}
		cc.ToSHA = newUpstreamSHA

		oldUpstreamSHA, ok := oldUpstream[upstreamRepo]
		if !ok {
			progress("  %s: no matching image found in project.yaml for %s", comp, upstreamRepo)
			cc.Error = fmt.Sprintf("no image in project.yaml maps to %s", upstreamRepo)
			results = append(results, cc)
			continue
		}
		cc.FromSHA = oldUpstreamSHA

		if oldUpstreamSHA == newUpstreamSHA {
			progress("  %s: unchanged", comp)
			cc.Commits = []Commit{}
			results = append(results, cc)
			if onResult != nil {
				onResult(buildSnapshotResult(version, operatorCommit, results))
			}
			continue
		}

		progress("Processing %s (%s)...", comp, upstreamRepo)
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
			onResult(buildSnapshotResult(version, operatorCommit, results))
		}
	}

	return buildSnapshotResult(version, operatorCommit, results), nil
}
