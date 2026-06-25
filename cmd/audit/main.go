package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	gh "github.com/theakshaypant/osp-rc-tools/internal/github"
	"github.com/theakshaypant/osp-rc-tools/internal/jira"
)

func main() {
	toDateStr := flag.String("to-date", "", "upper bound date (YYYY-MM-DD); defaults to current date for unreleased patches")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: go run cmd/audit/main.go [flags] <patch-version>")
		fmt.Fprintln(os.Stderr, "  e.g. go run cmd/audit/main.go 1.22.3")
		fmt.Fprintln(os.Stderr, "       go run cmd/audit/main.go --to-date 2025-06-01 1.22.3")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	var toDate *time.Time
	if *toDateStr != "" {
		t, err := time.Parse("2006-01-02", *toDateStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid --to-date format, expected YYYY-MM-DD:", err)
			os.Exit(1)
		}
		t = t.Add(24*time.Hour - time.Second)
		toDate = &t
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		out, err := exec.Command("gh", "auth", "token").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, "No GITHUB_TOKEN set and `gh auth token` failed")
			os.Exit(1)
		}
		token = strings.TrimSpace(string(out))
	}

	client, err := gh.NewClient(token)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating client:", err)
		os.Exit(1)
	}

	jiraClient, err := jira.NewClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Warning: Jira client error:", err)
	}

	ctx := context.Background()
	version := flag.Arg(0)
	if err := os.MkdirAll("reports", 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Error creating reports directory:", err)
		os.Exit(1)
	}
	outputFile := fmt.Sprintf("reports/release_%s.json", version)

	progress := func(format string, args ...any) {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}

	if jiraClient != nil {
		progress("Jira integration enabled")
	}

	writeResult := func(result *gh.AuditResult) {
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to marshal result: %v\n", err)
			return
		}
		if err := os.WriteFile(outputFile, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to write %s: %v\n", outputFile, err)
		}
	}

	var result *gh.AuditResult
	if _, _, merr := gh.ParseMinorVersion(version); merr == nil {
		result, err = gh.GetMinorCommits(ctx, client, jiraClient, version, toDate, progress, writeResult)
	} else {
		result, err = gh.GetPatchCommits(ctx, client, jiraClient, version, toDate, progress, writeResult)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if jiraClient != nil {
		progress("Fetching release note fields...")
		rnCache := map[string]*jira.ReleaseNote{}
		matched := map[string]bool{}
		rnCount := 0
		for ci := range result.Components {
			rnCount += attachReleaseNotes(jiraClient, result.Components[ci].Commits, rnCache, matched)
			rnCount += attachReleaseNotes(jiraClient, result.Components[ci].UnsyncedCommits, rnCache, matched)
		}
		progress("%d commits with release note fields", rnCount)
		writeResult(result)

		progress("Searching Jira for fixVersion tickets...")
		tickets, jerr := jiraClient.FindTicketsForFixVersion(version)
		if jerr != nil {
			progress("Warning: Jira fixVersion search failed: %v", jerr)
		} else {
			for _, t := range tickets {
				if !matched[t.Key] {
					result.UnmatchedJiras = append(result.UnmatchedJiras, t)
				}
			}
			progress("%d tickets with fixVersion, %d unmatched", len(tickets), len(result.UnmatchedJiras))
			writeResult(result)
		}
	}

	fmt.Fprintf(os.Stderr, "Results written to %s\n", outputFile)
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
