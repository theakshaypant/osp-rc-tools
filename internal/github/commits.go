package github

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/v88/github"
)

// GetCommitsBetweenRefs returns commits between two SHAs in an upstream repo.
// Handles pagination since the Compare API caps at 250 commits per page.
func GetCommitsBetweenRefs(ctx context.Context, client *github.Client, owner, repo, base, head string) ([]Commit, error) {
	var allCommits []Commit
	opts := &github.ListOptions{PerPage: 250}

	for {
		comparison, resp, err := client.Repositories.CompareCommits(ctx, owner, repo, base, head, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to compare %s...%s in %s/%s: %w", base[:min(12, len(base))], head[:min(12, len(head))], owner, repo, err)
		}

		for _, c := range comparison.Commits {
			sha := c.GetSHA()
			subject := strings.SplitN(c.Commit.GetMessage(), "\n", 2)[0]
			allCommits = append(allCommits, Commit{
				SHA:     sha,
				Message: subject,
				Author:  c.Commit.Author.GetName(),
				Date:    c.Commit.Author.GetDate().Time.Format("2006-01-02"),
			})
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allCommits, nil
}

// ResolvePRs populates the PR field for each commit by looking up the
// pull request that introduced it on the given branch.
func ResolvePRs(ctx context.Context, client *github.Client, owner, repo string, commits []Commit) {
	for i, c := range commits {
		prs, _, err := client.PullRequests.ListPullRequestsWithCommit(ctx, owner, repo, c.SHA, &github.ListOptions{PerPage: 1})
		if err != nil || len(prs) == 0 {
			continue
		}
		commits[i].PR = prs[0].GetHTMLURL()
	}
}

// GetBranchHead returns the HEAD SHA of a branch in a given repo.
func GetBranchHead(ctx context.Context, client *github.Client, owner, repo, branch string) (string, error) {
	ref, _, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch)
	if err != nil {
		return "", fmt.Errorf("failed to get branch head for %s/%s@%s: %w", owner, repo, branch, err)
	}
	return ref.Object.GetSHA(), nil
}

// GetBranchHeadAt returns the HEAD SHA of a branch at a specific point in time.
func GetBranchHeadAt(ctx context.Context, client *github.Client, owner, repo, branch string, before time.Time) (string, error) {
	commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, &github.CommitsListOptions{
		SHA:         branch,
		Until:       before,
		ListOptions: github.ListOptions{PerPage: 1},
	})
	if err != nil {
		return "", fmt.Errorf("failed to list commits for %s/%s@%s: %w", owner, repo, branch, err)
	}
	if len(commits) == 0 {
		return "", fmt.Errorf("no commits found for %s/%s@%s before %s", owner, repo, branch, before.Format(time.RFC3339))
	}
	return commits[0].GetSHA(), nil
}
