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

var (
	cherryPickSHARe    = regexp.MustCompile(`\(cherry picked from commit ([0-9a-f]{40})\)`)
	cherryPickMentionRe = regexp.MustCompile(`(?i)cherry[- ]?pick`)
	prURLRe            = regexp.MustCompile(`https://github\.com/[^/]+/[^/]+/pull/\d+`)
	prNumberRe         = regexp.MustCompile(`#(\d+)`)
	prURLNumberRe      = regexp.MustCompile(`/pull/(\d+)`)
)

func GetCommitsBetweenRefs(ctx context.Context, client *github.Client, owner, repo, base, head string) ([]Commit, error) {
	var allCommits []Commit
	opts := &github.ListOptions{PerPage: 250}

	for {
		comparison, resp, err := client.Repositories.CompareCommits(ctx, owner, repo, base, head, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to compare %s...%s in %s/%s: %w", base[:min(12, len(base))], head[:min(12, len(head))], owner, repo, err)
		}

		for _, c := range comparison.Commits {
			fullMsg := c.Commit.GetMessage()
			subject := strings.SplitN(fullMsg, "\n", 2)[0]
			commit := Commit{
				SHA:         c.GetSHA(),
				Message:     subject,
				Author:      c.Commit.Author.GetName(),
				Date:        c.Commit.Author.GetDate().Time.Format("2006-01-02"),
				fullMessage: fullMsg,
			}
			if m := cherryPickSHARe.FindStringSubmatch(fullMsg); len(m) == 2 {
				commit.cherryPickOf = m[1]
			}
			allCommits = append(allCommits, commit)
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allCommits, nil
}

func ResolvePRs(ctx context.Context, client *github.Client, owner, repo string, commits []Commit) {
	for i, c := range commits {
		prs, _, err := client.PullRequests.ListPullRequestsWithCommit(ctx, owner, repo, c.SHA, &github.ListOptions{PerPage: 1})
		if err != nil || len(prs) == 0 {
			continue
		}
		commits[i].PR = prs[0].GetHTMLURL()
	}
}

func ResolveOriginalPRs(ctx context.Context, client *github.Client, owner, repo string, commits []Commit) {
	for i, c := range commits {
		// Method 1: cherry-picked-from SHA in commit message trailer
		if c.cherryPickOf != "" {
			prs, _, err := client.PullRequests.ListPullRequestsWithCommit(ctx, owner, repo, c.cherryPickOf, &github.ListOptions{PerPage: 1})
			if err == nil && len(prs) > 0 {
				commits[i].OriginalPR = prs[0].GetHTMLURL()
				continue
			}
		}

		// Method 2: cherry-pick mention + PR ref in commit message
		if num, url := extractCherryPickPRRef(c.fullMessage); url != "" {
			commits[i].OriginalPR = url
			continue
		} else if num > 0 {
			pr, _, err := client.PullRequests.Get(ctx, owner, repo, num)
			if err == nil {
				commits[i].OriginalPR = pr.GetHTMLURL()
				continue
			}
		}

		// Method 3: cherry-pick mention + PR ref in the release-branch PR
		if c.PR == "" {
			continue
		}
		prNum := extractPRNumberFromURL(c.PR)
		if prNum == 0 {
			continue
		}
		pr, _, err := client.PullRequests.Get(ctx, owner, repo, prNum)
		if err != nil {
			continue
		}
		text := pr.GetTitle() + "\n" + pr.GetBody()
		if num, url := extractCherryPickPRRef(text); url != "" {
			commits[i].OriginalPR = url
		} else if num > 0 {
			origPR, _, err := client.PullRequests.Get(ctx, owner, repo, num)
			if err == nil {
				commits[i].OriginalPR = origPR.GetHTMLURL()
			}
		}
	}
}

func extractCherryPickPRRef(text string) (num int, url string) {
	for _, line := range strings.Split(text, "\n") {
		if !cherryPickMentionRe.MatchString(line) {
			continue
		}
		if m := prURLRe.FindString(line); m != "" {
			return 0, m
		}
		if m := prNumberRe.FindStringSubmatch(line); len(m) == 2 {
			n, err := strconv.Atoi(m[1])
			if err == nil {
				return n, ""
			}
		}
	}
	return 0, ""
}

func extractPRNumberFromURL(url string) int {
	m := prURLNumberRe.FindStringSubmatch(url)
	if len(m) != 2 {
		return 0
	}
	n, _ := strconv.Atoi(m[1])
	return n
}

func GetBranchHead(ctx context.Context, client *github.Client, owner, repo, branch string) (string, error) {
	ref, _, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+branch)
	if err != nil {
		return "", fmt.Errorf("failed to get branch head for %s/%s@%s: %w", owner, repo, branch, err)
	}
	return ref.Object.GetSHA(), nil
}

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
