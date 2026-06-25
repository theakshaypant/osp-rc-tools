package github

import "time"

type Commit struct {
	SHA     string `json:"sha"`
	Message string `json:"message"`
	Author  string `json:"author"`
	Date    string `json:"date"`
	PR      string `json:"pr,omitempty"`
}

type ComponentCommits struct {
	Name             string   `json:"name"`
	Upstream         string   `json:"upstream"`
	Downstream       string   `json:"downstream"`
	UpstreamBranch   string   `json:"upstream_branch"`
	DownstreamBranch string   `json:"downstream_branch"`
	FromSHA          string   `json:"from_sha,omitempty"`
	ToSHA            string   `json:"to_sha,omitempty"`
	Commits          []Commit `json:"commits,omitempty"`
	UpstreamHead     string   `json:"upstream_head,omitempty"`
	UnsyncedCommits  []Commit `json:"unsynced_commits,omitempty"`
	Error            string   `json:"error,omitempty"`
}

type ReleaseConfig struct {
	Version    string                  `yaml:"version"`
	ReleaseTag string                  `yaml:"release-tag"`
	PatchVer   string                  `yaml:"patch-version"`
	Branches   map[string]BranchConfig `yaml:"branches"`
}

func (c ReleaseConfig) EffectiveVersion() string {
	if c.ReleaseTag != "" {
		return c.ReleaseTag
	}
	return c.PatchVer
}

type BranchConfig struct {
	Upstream string `yaml:"upstream"`
}

type RepoConfig struct {
	Name     string `yaml:"name"`
	Repo     string `yaml:"repo"`
	Upstream string `yaml:"upstream"`
}

type PatchVersionEntry struct {
	Version string    `json:"version"`
	Date    time.Time `json:"date"`
}

type PatchBounds struct {
	From time.Time
	To   time.Time
}

type AuditResult struct {
	Version    string             `json:"version"`
	Previous   string             `json:"previous"`
	FromDate   string             `json:"from_date"`
	ToDate     string             `json:"to_date"`
	Released   bool               `json:"released"`
	Components []ComponentCommits `json:"components"`
}
