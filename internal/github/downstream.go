package github

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/v88/github"
)

const downstreamOrg = "openshift-pipelines"

func GetHeadSHAAt(ctx context.Context, client *github.Client, repo, branch string, before *time.Time) (string, error) {
	opts := &github.CommitsListOptions{
		SHA:         branch,
		Path:        "head",
		ListOptions: github.ListOptions{PerPage: 1},
	}
	if before != nil {
		opts.Until = *before
	}

	commits, _, err := client.Repositories.ListCommits(ctx, downstreamOrg, repo, opts)
	if err != nil {
		return "", fmt.Errorf("failed to list head file commits for %s/%s: %w", downstreamOrg, repo, err)
	}
	if len(commits) == 0 {
		return "", fmt.Errorf("no head file commits found for %s/%s", downstreamOrg, repo)
	}

	return readHeadFileAt(ctx, client, repo, commits[0].GetSHA())
}

func readHeadFileAt(ctx context.Context, client *github.Client, repo, commitSHA string) (string, error) {
	content, _, _, err := client.Repositories.GetContents(ctx, downstreamOrg, repo, "head", &github.RepositoryContentGetOptions{Ref: commitSHA})
	if err != nil {
		return "", fmt.Errorf("failed to read head file at %s: %w", commitSHA, err)
	}

	raw, err := content.GetContent()
	if err != nil {
		return "", fmt.Errorf("failed to decode head file: %w", err)
	}

	return strings.TrimSpace(raw), nil
}
