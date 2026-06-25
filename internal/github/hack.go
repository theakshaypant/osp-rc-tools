package github

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/google/go-github/v88/github"
	"gopkg.in/yaml.v3"
)

const hackRepo = "hack"

func releaseConfigPath(majorMinor string) string {
	return fmt.Sprintf("config/downstream/releases/%s.yaml", majorMinor)
}

func FetchReleaseConfig(ctx context.Context, client *github.Client, majorMinor, ref string) (*ReleaseConfig, error) {
	path := releaseConfigPath(majorMinor)
	content, _, _, err := client.Repositories.GetContents(ctx, downstreamOrg, hackRepo, path, &github.RepositoryContentGetOptions{Ref: ref})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch release config for %s: %w", majorMinor, err)
	}

	raw, err := content.GetContent()
	if err != nil {
		return nil, fmt.Errorf("failed to decode release config: %w", err)
	}

	var cfg ReleaseConfig
	if err := yaml.Unmarshal([]byte(raw), &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse release config: %w", err)
	}

	cfg.ReleaseTag = cfg.EffectiveVersion()

	return &cfg, nil
}

func FetchRepoConfig(ctx context.Context, client *github.Client, component string) (*RepoConfig, error) {
	path := fmt.Sprintf("config/downstream/repos/%s.yaml", component)
	content, _, _, err := client.Repositories.GetContents(ctx, downstreamOrg, hackRepo, path, &github.RepositoryContentGetOptions{Ref: "main"})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch repo config for %s: %w", component, err)
	}

	raw, err := content.GetContent()
	if err != nil {
		return nil, fmt.Errorf("failed to decode repo config: %w", err)
	}

	var cfg RepoConfig
	if err := yaml.Unmarshal([]byte(raw), &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse repo config for %s: %w", component, err)
	}

	if cfg.Repo == "" {
		cfg.Repo = component
	}

	return &cfg, nil
}

// GetPatchHistory returns the full timeline of patch version changes for a
// release by walking all commits that touched the release config file and
// recording when each version was first set. Entries are sorted
// chronologically (oldest first).
func GetPatchHistory(ctx context.Context, client *github.Client, majorMinor string) ([]PatchVersionEntry, error) {
	path := releaseConfigPath(majorMinor)

	// Track dates per version. ListCommits returns newest-first, so we keep
	// overwriting — the final date for each version is its oldest commit
	// (i.e., the commit that first set that version).
	dates := map[string]time.Time{}
	var order []string

	opts := &github.CommitsListOptions{
		SHA:         "main",
		Path:        path,
		ListOptions: github.ListOptions{PerPage: 50},
	}

	for {
		commits, resp, err := client.Repositories.ListCommits(ctx, downstreamOrg, hackRepo, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list hack repo commits: %w", err)
		}

		for _, c := range commits {
			ver, err := getVersionAtCommit(ctx, client, majorMinor, c.GetSHA())
			if err != nil {
				continue
			}

			if _, exists := dates[ver]; !exists {
				order = append(order, ver)
			}
			dates[ver] = c.Commit.Author.GetDate().Time
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	// Build entries in discovery order (newest-first), then reverse
	entries := make([]PatchVersionEntry, 0, len(order))
	for _, ver := range order {
		entries = append(entries, PatchVersionEntry{Version: ver, Date: dates[ver]})
	}

	slices.Reverse(entries)

	return entries, nil
}

func getVersionAtCommit(ctx context.Context, client *github.Client, majorMinor, sha string) (string, error) {
	cfg, err := FetchReleaseConfig(ctx, client, majorMinor, sha)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(cfg.EffectiveVersion()), nil
}
