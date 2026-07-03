package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	gh "github.com/theakshaypant/osp-rc-tools/internal/github"
	"github.com/theakshaypant/osp-rc-tools/internal/jira"
)

func main() {
	version := flag.String("version", "", "release version (e.g. 1.21.3); auto-derived from hack config if omitted")
	namespace := flag.String("namespace", "tekton-ecosystem-tenant", "Konflux namespace")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: go run cmd/snapshot-diff/main.go [flags] <old-snapshot> <new-snapshot>")
		fmt.Fprintln(os.Stderr, "  e.g. go run cmd/snapshot-diff/main.go openshift-pipelines-core-1-21-20260530-122116-000 openshift-pipelines-core-1-21-20260702-114559-000-lv")
		fmt.Fprintln(os.Stderr)
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}
	oldSnap := flag.Arg(0)
	newSnap := flag.Arg(1)

	server := os.Getenv("KONFLUX_SERVER")
	token := os.Getenv("KONFLUX_TOKEN")
	if server == "" || token == "" {
		fmt.Fprintln(os.Stderr, "KONFLUX_SERVER and KONFLUX_TOKEN must be set")
		os.Exit(1)
	}

	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		out, err := exec.Command("gh", "auth", "token").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, "No GITHUB_TOKEN set and `gh auth token` failed")
			os.Exit(1)
		}
		ghToken = strings.TrimSpace(string(out))
	}

	client, err := gh.NewClient(ghToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating GitHub client:", err)
		os.Exit(1)
	}

	jiraClient, err := jira.NewClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Warning: Jira client error:", err)
	}

	ctx := context.Background()
	if err := os.MkdirAll("reports", 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Error creating reports directory:", err)
		os.Exit(1)
	}

	progress := func(format string, args ...any) {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}

	if jiraClient != nil {
		progress("Jira integration enabled")
	}

	var outputFile string
	writeResult := func(result *gh.AuditResult) {
		if outputFile == "" {
			outputFile = fmt.Sprintf("reports/release_%s.json", result.Version)
		}
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to marshal result: %v\n", err)
			return
		}
		if err := os.WriteFile(outputFile, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to write %s: %v\n", outputFile, err)
		}
	}

	result, err := gh.GetSnapshotDiffCommits(ctx, client, jiraClient, server, token, *namespace, oldSnap, newSnap, *version, progress, writeResult)
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
		}
		progress("%d commits with release note fields", rnCount)
		writeResult(result)

		progress("Searching Jira for fixVersion tickets...")
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
