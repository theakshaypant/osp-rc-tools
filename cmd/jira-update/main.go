package main

import (
	"encoding/json"
	"fmt"
	"os"

	gh "github.com/theakshaypant/osp-rc-tools/internal/github"
	"github.com/theakshaypant/osp-rc-tools/internal/jira"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run cmd/jira-update/main.go <release_X.Y.Z.json>")
		os.Exit(1)
	}
	inputFile := os.Args[1]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		os.Exit(1)
	}

	var result gh.AuditResult
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
		os.Exit(1)
	}

	progress := func(format string, args ...any) {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}

	jiraClient, err := jira.NewClient()
	if err != nil || jiraClient == nil {
		fmt.Fprintln(os.Stderr, "Jira client not configured. Set JIRA_URL, JIRA_EMAIL, JIRA_TOKEN.")
		os.Exit(1)
	}

	progress("Fetching release note fields...")
	rnCache := map[string]*jira.ReleaseNote{}
	matched := map[string]bool{}
	rnCount := 0
	for ci := range result.Components {
		rnCount += attachReleaseNotes(jiraClient, result.Components[ci].Commits, rnCache, matched)
		rnCount += attachReleaseNotes(jiraClient, result.Components[ci].UnsyncedCommits, rnCache, matched)
	}
	progress("%d commits with release note fields", rnCount)

	progress("Searching Jira for fixVersion tickets...")
	result.UnmatchedJiras = nil
	tickets, jerr := jiraClient.FindTicketsForFixVersion(result.Version)
	if jerr != nil {
		progress("Warning: Jira fixVersion search failed: %v", jerr)
	} else {
		for _, t := range tickets {
			if !matched[t.Key] {
				result.UnmatchedJiras = append(result.UnmatchedJiras, t)
			}
		}
		progress("%d tickets with fixVersion, %d unmatched", len(tickets), len(result.UnmatchedJiras))
	}

	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling result:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(inputFile, out, 0644); err != nil {
		fmt.Fprintln(os.Stderr, "Error writing file:", err)
		os.Exit(1)
	}

	progress("Updated %s", inputFile)
}

func attachReleaseNotes(jiraClient *jira.Client, commits []gh.Commit, cache map[string]*jira.ReleaseNote, matched map[string]bool) int {
	count := 0
	for i := range commits {
		if commits[i].Jira == "" {
			continue
		}
		key := jira.KeyFromURL(commits[i].Jira)
		matched[key] = true

		rn, ok := cache[key]
		if !ok {
			var err error
			rn, err = jiraClient.FetchReleaseNotes(key)
			if err != nil {
				rn = nil
			}
			cache[key] = rn
		}
		if rn != nil {
			commits[i].ReleaseNote = rn
			count++
		}
	}
	return count
}
