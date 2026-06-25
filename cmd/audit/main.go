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

	ctx := context.Background()
	version := flag.Arg(0)

	result, err := gh.GetPatchCommits(ctx, client, version, toDate, func(format string, args ...any) {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)
}
