# Proposal: End-to-End Release Automation for OpenShift Pipelines

## Problem

The OpenShift Pipelines release process spans ~20 component repos, 16 sequential steps, and 5 external systems (GitHub, GitLab, Konflux, Jira, container registries). Today, a release engineer manually checks each step, performs each action, and re-checks. A typical patch release takes 2-3 days of intermittent attention — not because the work is hard, but because the human is the scheduler, the state machine, and the executor.

Specific pain points:

- **No single source of truth** for release progress. The engineer tracks state mentally or in ad-hoc notes.
- **Repeated manual verification** — the same checks (is this PR merged? does this SHA match? is the pipeline done?) are performed dozens of times per release.
- **CVE exposure** — there is no systematic pre-release scan of upstream dependencies. Vulnerabilities are discovered reactively, sometimes after images ship to stage.
- **Knowledge concentration** — the release process lives in runbooks and tribal knowledge. When the primary release engineer is unavailable, releases slow down.
- **Error-prone sequencing** — the release has strict ordering constraints (core → CLI → operator → index). Applying a release YAML out of order or with the wrong snapshot wastes hours.

## Proposed Solution

Build a **state-driven release orchestrator** as a set of Claude Code skills backed by Go CLI tools, running in the `openshift-pipelines/one-click-release` repo. Each patch release runs on its own branch (`release-v1.21.3`), and every phase commits its reports and state to that branch — making the release a trackable, auditable git history that any team member can inspect or pick up.

A single command — `/release 1.21.3` — walks the entire release lifecycle: scanning for CVEs, verifying each phase against external systems, executing actions with human approval, and advancing to the next phase. Re-running the command picks up where the last run left off. A different release captain can clone the branch and continue — the state travels with the repo, not the person.

The system is composed of:

1. **A release repo** (`openshift-pipelines/one-click-release`) with a branch per release (`release-vX.Y.Z`) that tracks all state and reports.
2. **An orchestrator skill** (`/release`) that reads state from the branch, determines the current phase, and delegates to the appropriate sub-skill.
3. **Phase skills** that each own a slice of the release lifecycle — both verifying the expected state and executing actions to reach it.
4. **A CVE scanner/fixer** that gates entry to the release — no phase begins until upstream repos are clean at the configured severity threshold.
5. **Go CLI tools** that collect structured data (commits, snapshots, Jira metadata) consumed by the skills.
6. **Utility skills** for ad-hoc tasks: diffing snapshots, generating QE handover reports.

---

## Architecture

### Release Repo and Branching

The orchestrator runs inside `openshift-pipelines/one-click-release`. Each release gets its own branch:

```
main                          ← skills, Go tools, shared config
├── release-v1.21.2           ← completed release (reports committed)
├── release-v1.21.3           ← in-progress release
├── release-v1.22.4           ← in-progress release (different stream)
```

On first invocation for a version, the orchestrator creates the branch from `main` and commits the initial state file. Every phase that completes commits its updated state and reports to the branch. The git log on the branch becomes the release audit trail:

```
$ git log --oneline release-v1.21.3
a3f2c1d phase 8: prod release complete — advisory pending
9b4e7a2 phase 7: audit-release + jira-gaps complete
1c8d3f5 phase 6: stage release complete — QA gate reached
...
e7a1b4c phase 1: check-hack complete
d2f9c8e prep-step: operator-bump — PAC v0.29.3, Results v0.12.1 (tektoncd/operator#2150)
c4b7a1f prep-step: cve-scan clean (19 repos, 2 fixed)
b1a3e5d init: release 1.21.3 started
```

**Release captain handoff:** Any team member clones the repo, checks out `release-v1.21.3`, and runs `/release 1.21.3`. The orchestrator reads the state from the branch and picks up where the previous captain left off.

### Orchestrator

```
/release 1.21.3
```

On every invocation:

1. Checkout or create branch `release-v1.21.3` from `main`
2. Load the state file `reports/1.21/1.21.3/release.json` from the branch
3. **Prep-step gates (parallel):**
   - **CVE scan:** Scan all upstream repos. If any CVE at or above the severity threshold exists, create fix PRs, wait for merges, re-scan.
   - **Operator bump:** Check `tektoncd/operator` `components.yaml` on the release branch. If any component has a newer release tag, update the file and open a PR.
   Both gates must pass before release phases begin.
4. Once both prep-steps pass, for each phase starting from the last incomplete one:
   - Verify the current external state (GitHub, Konflux, Jira)
   - If done → mark complete, advance
   - If action needed → execute, confirm with user, re-verify
   - If blocked (waiting for pipeline, QA, external) → report and stop
5. Commit updated state file and any new reports to the branch
6. Push the branch
7. Print a summary: what completed, what's next, what's blocking

The orchestrator is **re-entrant**. Each invocation reads the state from the branch and the external world, then advances as far as it can. No flags or modes — it always verifies and executes. Individual phase skills can still be run standalone for verification-only use.

### Patch Release Flow

```
              ┌─────────────────────────────────────────────────────┐
  Prep-step:   │                                                     │
              ▼                                                     ▼
  ┌──────────────────────────────┐ ┌──────────────────────────────────┐
  │  cve-scan                    │ │  operator-bump                   │
  │  Scan upstream repos for     │ │  Bump component versions in      │
  │  CVEs ≥ high                 │ │  tektoncd/operator components    │
  │  Execute: create fix PRs,    │ │  .yaml on the release branch     │
  │  re-scan after merge         │ │  Execute: resolve latest tags,   │
  │                              │ │  update YAML, open PR            │
  └──────────┬───────────────────┘ └──────────┬───────────────────────┘
             └─────────────┬──────────────────┘
                           │
  ═════════════════════════╪══════════════════════════════════════════
   GATE: All upstream repos CVE-clean at threshold severity.
         Operator components.yaml PR merged with latest versions.
   Release phases will not start until both gates pass.
  ═════════════════════════╪══════════════════════════════════════════
                           │
                           ▼
   Phase 1: ┌───────────────────────────────────────────────────────────┐
            │  check-hack                                              │
            │  Verify hack repo config, release-manager PR             │
            │  Execute: trigger release-manager workflow, merge PR     │
            └───────────────────────────┬───────────────────────────────┘
                                        │
                          ┌─────────────┴─────────────┐
                          ▼                           ▼
   Phase 2: ┌──────────────────────────┐ ┌──────────────────────────┐
            │  check-konflux-config    │ │  check-components        │
            │  Konflux config PR,      │ │  Component PRs,          │
            │  RPA/RP, Pyxis           │ │  upstream sync, nudges   │
            │  Execute: create/merge   │ │  Execute: merge nudge    │
            │  config PR, oc apply     │ │  PRs, trigger sync       │
            └──────────┬───────────────┘ └──────────┬───────────────┘
                       └─────────────┬──────────────┘
                                     ▼
   Phase 3: ┌───────────────────────────────────────────────────────────┐
            │  check-builds                                            │
            │  Validate operator version, build SHAs                   │
            │  Execute: create operator version bump PR                │
            └───────────────────────────┬───────────────────────────────┘
                                        │
                                        ▼
   Phase 4: ┌───────────────────────────────────────────────────────────┐
            │  check-olm                                               │
            │  Verify OLM catalog, index builds, code freeze           │
            │  Execute: trigger index-render, toggle freeze            │
            └───────────────────────────┬───────────────────────────────┘
                                        │
                                        ▼
   Phase 5: ┌───────────────────────────────────────────────────────────┐
            │  audit-release + jira-gaps                               │
            │  First snapshot ready + code freeze = commit window       │
            │  is closed. Audit commits, fix Jira gaps, generate       │
            │  release notes. Share report with docs team.             │
            │  Execute: enrich JSON, update Jira fixVersion/RN fields, │
            │           notify docs team with release notes report     │
            └───────────────────────────┬───────────────────────────────┘
                                        │
                                        ▼
   Phase 6: ┌───────────────────────────────────────────────────────────┐
            │  check-releases (stage)                                  │
            │  Detect latest snapshot, apply release YAMLs             │
            │  Execute: oc apply core → CLI → operator → index         │
            │  Poll pipeline completion                                │
            └───────────────────────────┬───────────────────────────────┘
                                        │
                                        ▼
  ══════════════════════════════════════════════════════════════════════
   GATE: QA sign-off required
  ══════════════════════════════════════════════════════════════════════
                                        │
                          ┌─────────────┴─────────────┐
                          ▼                           ▼
   Phase 7: ┌──────────────────────────┐ ┌──────────────────────────┐
            │  check-releases (prod)   │ │  cli-release             │
            │  Apply prod release      │ │  OPC sync, p12n-opc,     │
            │  YAMLs                   │ │  serve-tkn-cli, CDN      │
            │  Execute: oc apply       │ │  Execute: trigger OPC    │
            │  core → CLI → op → idx   │ │  sync, submodule update, │
            │  Poll completion         │ │  builds, CDN release     │
            └──────────┬───────────────┘ └──────────┬───────────────┘
                       └─────────────┬──────────────┘
                                     ▼
   Phase 8: ┌───────────────────────────────────────────────────────────┐
            │  check-releases (advisory)                               │
            │  Verify advisory published                               │
            │  Execute: manual (external process)                      │
            └───────────────────────────────────────────────────────────┘

  Helpers (anytime):
    build-diff     — diff two snapshots
    new-build      — build health + QE diff

    cve-scan       — re-scan at any point (standalone)
```

### State File

`reports/X.Y/X.Y.Z/release.json` on the `release-vX.Y.Z` branch is the single source of truth for release progress:

```json
{
  "version": "1.21.3",
  "type": "patch",
  "started": "2026-07-06T10:00:00+05:30",
  "prep_step": {
    "cve-scan": {
      "status": "done",
      "completed_at": "2026-07-06T10:30:00+05:30",
      "threshold": "high",
      "report": "reports/1.21/1.21.3/cve-scan-20260706.md",
      "summary": { "scanned": 19, "clean": 17, "fixed": 2, "blocking": 0 },
      "actions": [
        {
          "type": "gh_pr_create",
          "target": "tektoncd/pipeline#8205",
          "cve": "CVE-2026-1234",
          "package": "golang.org/x/net",
          "result": "merged",
          "at": "2026-07-06T10:15:00+05:30"
        }
      ]
    },
    "operator-bump": {
      "status": "done",
      "completed_at": "2026-07-06T10:25:00+05:30",
      "actions": [
        {
          "type": "gh_pr_create",
          "target": "tektoncd/operator#2150",
          "branch": "release-v0.72.x",
          "bumped": ["pac v0.29.3", "results v0.12.1"],
          "result": "merged",
          "at": "2026-07-06T10:24:00+05:30"
        }
      ]
    }
  },
  "phases": {
    "check-hack": {
      "status": "done",
      "completed_at": "2026-07-06T11:05:00+05:30",
      "actions": [
        {
          "type": "gh_pr_merge",
          "target": "openshift-pipelines/hack#892",
          "result": "merged",
          "at": "2026-07-06T11:04:30+05:30"
        }
      ]
    },
    "check-releases-stage": {
      "status": "in_progress",
      "started_at": "2026-07-06T14:30:00+05:30",
      "actions": [
        {
          "type": "oc_apply",
          "target": "openshift-pipelines-core-1.21.3-release.yaml",
          "snapshot": "openshift-pipelines-core-1-21-20260706-143000-000",
          "result": "applied",
          "at": "2026-07-06T14:31:00+05:30"
        }
      ],
      "waiting_for": "core release pipeline completion"
    }
  }
}
```

When the orchestrator checks out `release-v1.21.3` and loads this file, it sees the prep-step passed (`cve-scan: done`) and `check-releases-stage` is `in_progress`. It checks whether the pipeline has completed, advances if so, commits the updated state, and pushes.

---

## Skills

### Phase Skills

Each skill owns both sides of its domain — verification and execution. When run standalone, it verifies and reports. When invoked by the orchestrator, it also executes actions to advance the phase.

| Skill | Phase | Verify | Execute |
|-------|-------|--------|---------|
| **cve-scan** | Prep-step (parallel) | Scan upstream repos for CVEs above threshold (trivy fs + govulncheck) | Create fix PRs on upstream repos, re-scan after merge |
| **operator-bump** | Prep-step (parallel) | Check `tektoncd/operator` `components.yaml` for outdated component versions on release branch | Resolve latest release tags per component, update `components.yaml`, open PR on operator release branch ([SRVKP-12488](https://redhat.atlassian.net/browse/SRVKP-12488)) |
| **check-hack** | 1 | Check hack repo release config, release-manager PR | Trigger `release-manager` workflow, merge PR |
| **check-konflux-config** | 2 (parallel) | Check Konflux config PR, RPA/RP resources, Pyxis | Create/merge config PR, `oc apply` RPA/RP |
| **check-components** | 2 (parallel) | Check component PRs, upstream sync, nudge PRs | Merge ready nudge PRs, trigger sync workflows |
| **check-builds** | 3 | Validate operator `project.yaml` version, build SHAs | Create operator version bump PR |
| **check-olm** | 4 | Check OLM catalog config, index builds, code freeze | Trigger `index-render-template` workflow, toggle freeze |
| **audit-release** | 5 | Enrich commits with Jira links and release notes | Generate audit JSON and markdown report, share with docs team |
| **jira-gaps** | 5 | Find Jiras missing fixVersion or release notes | Update fixVersion and release note fields in Jira |
| **check-releases** (stage) | 6 | Check stage release pipeline status | Auto-detect snapshot, apply release YAMLs (core → CLI → operator → index), poll completion |
| **check-releases** (prod) | 7 (parallel) | Check prod release pipeline status | Apply prod release YAMLs, poll completion |
| **cli-release** | 7 (parallel) | CLI binary release: OPC → p12n-opc → serve-tkn-cli → CDN | Trigger OPC sync, submodule updates, serve-tkn-cli builds, CDN release |
| **check-releases** (advisory) | 8 | Verify advisory published | Manual (external process) |

### Utility Skills

These run independently, not as orchestrator phases.

| Skill | Trigger | Purpose |
|-------|---------|---------|
| **build-diff** | `/build-diff snap1 snap2` | Compare two Konflux snapshots — show which repos changed, upstream commits, PRs, Jiras |
| **new-build** | `/new-build 1.22.4` | When a fix lands mid-release: check build health, diff baseline vs latest snapshot, generate QE handover report |


### Go CLI Tools

Go tools provide the data infrastructure — structured collection from GitHub, Konflux, and Jira APIs, producing JSON artifacts that the skills consume and enrich.

| Tool | Purpose | Output |
|------|---------|--------|
| `cmd/audit` | Collect commits across all ~20 components for a release version. Resolves PRs (including original upstream PRs for cherry-picks), links to Jira tickets via PR custom field, fetches release note fields. | `reports/X.Y/X.Y.Z/release.json` |
| `cmd/snapshot-diff` | Compare two Konflux snapshots. For each changed component, resolve the upstream commit diff via downstream `head` files. | `reports/X.Y/X.Y.Z/release.json` |
| `cmd/image-diff` | Compare upstream SHAs from operator `project.yaml` images (via `skopeo inspect`) against a Konflux snapshot. Alternative to snapshot-diff when a prior snapshot doesn't exist. | `reports/X.Y/X.Y.Z/release.json` |
| `cmd/jira-update` | Re-enrich an existing audit JSON with fresh Jira data (release notes, fixVersion search). With `--write-back`, apply fixVersion and release note fields back to Jira. | Overwrites input JSON |
| `cmd/cve-scan` | Scan upstream repos for vulnerabilities. Clone each repo at release branch, run `trivy fs` + `govulncheck`, cross-reference with Jira, filter by severity. | `reports/X.Y/X.Y.Z/cve-scan.json` |

Data flow:

```
cmd/cve-scan ──→ X.Y/X.Y.Z/cve-scan.json ──→ cve-scan skill ──→ fix PRs + report
cmd/audit ──→ X.Y/X.Y.Z/release.json ──→ audit-release skill ──→ enriched JSON + report
                                       ──→ jira-gaps skill ──→ Jira updates + report
cmd/snapshot-diff ──→ X.Y/X.Y.Z/release.json (alternative input to audit-release/jira-gaps)
cmd/image-diff ──→ X.Y/X.Y.Z/release.json (alternative input)
```

All outputs are written under `reports/{MAJOR.MINOR}/{X.Y.Z}/`. This groups all artifacts for a release version together and organizes them by minor release stream.

---

## CVE Scanner and Fixer

### Purpose

The CVE scan is a **prep-step gate** — the orchestrator will not start any release phase until all upstream repos are clean at the configured severity threshold. This ensures no release ships with known high-severity vulnerabilities in its dependency tree.

### Invocation

```
/cve-scan 1.21.3                      # scan, default threshold: high
/cve-scan 1.21.3 --severity critical  # only critical
/cve-scan 1.21.3 --fix                # scan + create fix PRs on upstream repos
```

When invoked by the orchestrator, it runs with `--fix` automatically.

### Scan

1. Fetch release config from hack repo — get the list of components, upstream repos, and release branches
2. For each upstream repo on its release branch:
   - Clone or fetch the repo at the release branch HEAD
   - Run `trivy fs --severity HIGH,CRITICAL --scanners vuln --format json .` — scans `go.sum` for Go dependency CVEs
   - Run `govulncheck ./...` — reachability analysis to filter out CVEs in code paths that aren't actually called
3. Cross-reference findings with Jira: search for existing `SecurityTracking` tickets (`project = SRVKP AND labels = SecurityTracking AND summary ~ "CVE-XXXX-YYYY"`)
4. Produce a report: `reports/X.Y/X.Y.Z/cve-scan-{timestamp}.md`

Report includes:
- Per-repo findings table (CVE, severity, package, installed version, fixed version, reachability, linked Jira)
- Cross-repo summary (one CVE affecting multiple repos shown once with all affected repos listed)
- Clean repos list

### Fix

When `--fix` is passed or the orchestrator invokes the skill:

1. **Group by CVE, not by repo** — a single vulnerable package (e.g. `golang.org/x/net`) may affect multiple repos
2. **For each affected repo:**
   - Fork or clone the upstream repo
   - Checkout the release branch
   - `go get <package>@<fixed-version>` for each vulnerable dependency
   - `go mod tidy`
   - `go build ./...` to verify compilation
   - Create a PR on the release branch with CVE details in the description
   - Link to existing Jira ticket if one exists
3. **Wait for CI** — poll `gh pr checks` until pass or fail
4. **Merge** — with user confirmation
5. **Re-scan** — after all fix PRs merge, re-scan affected repos. Gate only passes when clean.

The fixer does **not** touch downstream repos (`openshift-pipelines/*`). Upstream fixes flow downstream via the existing sync/nudge pipeline in phase 3 of the release flow.

### Prep-step gate logic

```
CVEs found above threshold?
  No  → orchestrator enters release phases (phase 1)
  Yes and fixable → fix PRs on upstream repos, re-scan (loop until clean)
  Yes and unfixable (no fix available) → stop, report as blocker
    User can override with explicit approval to enter release phases
```

### Severity Threshold

Default: `HIGH` and above (HIGH + CRITICAL).

- `high` — default, recommended for patch releases
- `critical` — only CRITICAL, for time-sensitive hotfixes where known HIGHs are accepted
- `medium` — MEDIUM and above, for minor releases with more runway

---

## Output Locations

All outputs live on the release branch in `openshift-pipelines/one-click-release`, organized under `reports/{MAJOR.MINOR}/{X.Y.Z}/`:

```
openshift-pipelines/one-click-release  (branch: release-v1.21.3)
└── reports/
    └── 1.21/
        └── 1.21.3/
            ├── release.json            ← state file (orchestrator)
            ├── cve-scan.json           ← CVE scan structured data
            ├── cve-scan-20260706.md    ← CVE scan report
            ├── audit-release.md        ← commit audit with Jira links
            ├── jira-gaps.md            ← Jiras missing fields
            ├── checklist.md            ← phase verification report
            ├── new-build-20260706.md   ← QE handover diff
            └── build-diff-20260702.md  ← snapshot comparison
```

Each release branch contains only its own version's reports. The `main` branch holds the skills, Go tools, and shared configuration — no report data.

This gives each release:
- **A git-native audit trail** — every phase completion is a commit with a descriptive message
- **Release captain handoff** — any team member checks out the branch and runs `/release 1.21.3` to continue
- **Visibility** — `git log release-v1.21.3` shows exactly where the release stands, who advanced each phase, and when
- **History** — completed release branches are preserved as a record of past releases

---

## Prerequisites for Local Use

To run the skills locally, a user needs the following credentials, CLI tools, and access grants.

### Credentials (`.env`)

| Variable | Source | Required For |
|----------|--------|-------------|
| `GITHUB_TOKEN` | GitHub personal access token (or `gh auth login`) | All skills — repo access, PR operations, workflow triggers |
| `GITLAB_URL` | GitLab instance URL (e.g. `https://gitlab.cee.redhat.com`) | check-konflux-config — Pyxis verification |
| `GITLAB_TOKEN` | GitLab personal access token (read-only API scope) | check-konflux-config — Pyxis verification |
| `KONFLUX_SERVER` | Konflux cluster API URL | check-builds, check-olm, check-releases, new-build, build-diff |
| `KONFLUX_TOKEN` | Konflux cluster service account or user token | check-builds, check-olm, check-releases, new-build, build-diff |
| `JIRA_URL` | Jira instance URL (e.g. `https://redhat.atlassian.net`) | audit-release, jira-gaps, cve-scan |
| `JIRA_EMAIL` | Jira account email | audit-release, jira-gaps |
| `JIRA_TOKEN` | Jira API token | audit-release, jira-gaps |

### CLI Tools

| Tool | Version | Required For |
|------|---------|-------------|
| `gh` | 2.x+ | All GitHub operations — authenticated via `gh auth login` |
| `oc` / `kubectl` | 4.x+ | Konflux cluster queries and release YAML application |
| `go` | 1.22+ | Building Go CLI tools (`cmd/audit`, `cmd/snapshot-diff`, etc.) |
| `trivy` | 0.50+ | CVE scanning (`trivy fs` on upstream repos) |
| `govulncheck` | latest | CVE reachability analysis on Go code |
| `skopeo` | 1.x+ | `cmd/image-diff` — inspect operator images in container registries |

### Access Grants

| System | Access Needed | Scope |
|--------|--------------|-------|
| **GitHub (openshift-pipelines/*)** | Read + Write (push, PR create/merge, workflow dispatch) | All downstream repos, hack repo, one-click-release repo |
| **GitHub (tektoncd/*)** | Read + Write (PR create for CVE fixes) | Upstream repos listed in hack release config |
| **GitLab (gitlab.cee.redhat.com)** | Read-only API access | Pyxis project for container metadata verification |
| **Konflux cluster** | `get` on snapshots, releases, pipelineruns in `tekton-ecosystem-tenant`; `apply` for RPA/RP and release YAMLs | Single namespace: `tekton-ecosystem-tenant` |
| **Jira (redhat.atlassian.net)** | Read (search, issue fetch); Write (update fixVersion, release note fields) | Project `SRVKP` and related projects |
| **Container registries** | Pull access to `registry.redhat.io` and `quay.io` | For `skopeo inspect` in image-diff |

### Minimum Viable Setup

Not all credentials are required for every skill. A user can start with a subset:

| Want to run | Minimum credentials |
|-------------|-------------------|
| `release-checklist` (verify only) | `GITHUB_TOKEN` + `KONFLUX_*` + `GITLAB_*` |
| `audit-release` | `GITHUB_TOKEN` + `JIRA_*` |
| `jira-gaps` | `GITHUB_TOKEN` + `JIRA_*` |
| `build-diff` / `new-build` | `GITHUB_TOKEN` + `KONFLUX_*` + `JIRA_*` |
| `cve-scan` | `GITHUB_TOKEN` + `trivy` + `govulncheck` installed |
| Full `/release` orchestrator | All credentials above |

---

## Permission Model

### Tier 0 — Observe

All external systems read-only. Skills verify state and produce markdown reports. The human reads reports and acts manually.

- **GitHub:** `gh api` GET, `gh pr list`, `gh run list`, `gh pr view`
- **GitLab:** GET requests only
- **Konflux:** `oc get` / `kubectl get` only
- **Jira:** GET requests only (REST API search and fetch)
- **Local:** Writes only to `reports/{MAJOR.MINOR}/{X.Y.Z}/` directory

### Tier 1 — Supervised Execute

Every write operation gated by user confirmation — the human sees exactly what will happen and approves.

| System | Write Operations | Used By |
|--------|-----------------|---------|
| **GitHub** | `gh pr merge` | check-hack, check-components, check-builds, check-olm, cve-scan |
| **GitHub** | `gh pr create` (upstream tektoncd/* repos) | cve-scan |
| **GitHub** | `gh pr create` (openshift-pipelines/* repos) | check-builds, check-konflux-config |
| **GitHub** | `gh workflow run` | check-hack, check-olm, check-releases, cli-release |
| **Konflux** | `oc apply -f release.yaml` | check-konflux-config (RPA/RP), check-releases (release YAMLs) |
| **Jira** | `PUT /rest/api/2/issue/{key}` | jira-gaps |
| **GitLab** | No writes needed | — |

### Tier 2 — Autonomous

Same writes, confirmation removed except at mandatory gates:
- **CVE override** — proceeding with unfixable CVEs above threshold
- **QA sign-off** — between stage and prod release
- **Prodsec contact** — before prod release
- **Advisory** — external process, always manual

Requires: dry-run mode, structured audit trail in state file, failure notification.

---

## Risk Mitigation

| Risk | Mitigation |
|------|-----------|
| **Wrong snapshot applied** | Auto-detect via Konflux API; show snapshot details for confirmation |
| **PR merged with failing CI** | Check CI status before `gh pr merge`; refuse if checks failing |
| **Konflux release with wrong version** | Re-verify after execute; mismatch = stop |
| **CVE fix introduces regression** | Fix PRs go through upstream CI; `go build ./...` verified before PR creation; re-scan after merge |
| **CVE fix PR on upstream repo rejected** | Report rejection; user negotiates with upstream or overrides gate |
| **Unfixable CVE blocks release** | User can override with explicit approval; logged in state file with justification |
| **Runaway automation** | Mandatory gates at QA and prod; autonomous requires explicit opt-in |
| **Session ends mid-release** | State file enables resumption; steps are idempotent |
| **Jira writes update wrong tickets** | Preview before apply; dry-run mode |
| **Release YAML applied out of order** | Enforces core → CLI → operator → index ordering; each waits for previous |

---

## What This Enables

**Without automation:** A release engineer manually checks each step, performs each action, and re-checks. The process lives in runbooks and tribal knowledge. A patch release takes 2-3 days. CVEs are caught reactively.

**Phase 1 (verify):** The engineer runs `/release 1.21.3` and gets a live checklist showing exactly where the release stands and what to do next. No more mental state tracking. Any team member can run it.

**Phase 2 (CVE gate):** Before any release work begins, `/cve-scan` ensures upstream repos are clean. Vulnerabilities are caught and fixed proactively, with PRs created automatically on upstream repos.

**Phase 4 (full orchestrator):** The engineer runs `/release 1.21.3` once. It walks through all phases, performs each action with approval, and stops when waiting for something external. Re-running picks up where it left off. The engineer's role shifts from "execute and re-check" to "review and approve."

**Phase 5 (autonomous):** The orchestrator runs on a schedule. It advances through configuration, CVE remediation, build validation, and stage release without intervention. It stops at QA sign-off and prod release, notifying the engineer. Routine patch releases complete with the engineer touching only the mandatory gates.
