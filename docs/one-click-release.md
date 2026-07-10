# One-Click Release Handbook

A guide for the `/one-click-release` Claude Code skill, which automates the OpenShift Pipelines release workflow from config through image delivery.

This document is for people running releases or onboarding to the release process. It explains what each step does and why, without reproducing the underlying commands (those live in the skill files under `.claude/skills/one-click-release/`).

---

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [How It Works](#how-it-works)
- [Stage 1: Config](#stage-1-config)
- [Stage 2: Build](#stage-2-build)
- [Stage 3: Image Copy](#stage-3-image-copy)
- [Stage 4: Test](#stage-4-test)
- [Stage 5: Stage Release](#stage-5-stage-release)
- [Stage 6: Prod Release](#stage-6-prod-release)
- [Reports](#reports)
- [Troubleshooting](#troubleshooting)

---

## Introduction

The one-click-release skill walks through the full OpenShift Pipelines release lifecycle in stages. It replaces a manual checklist of ~22 steps that previously required an engineer to SSH into clusters, run scripts, monitor GitHub workflows, merge PRs, and coordinate across repos — all while tracking state in a spreadsheet.

The skill automates this by querying external systems (GitHub, GitLab, Konflux cluster) to verify each step's status, then offering to execute fixes when something isn't done. Steps run in order, and the skill stops at the first step that needs attention, so you always know exactly where the release stands.

**What it covers today:**

| Stage | Steps | What it does |
|-------|-------|-------------|
| Config | 1.1–1.13 | All configuration needed before builds can start — hack repo version bump, Konflux cluster config, RPA/Pyxis, OLM bundle version, operator project.yaml, OPC versions, CLI submodules |
| Build | 2.1–2.7 | Process release PRs, wait for Konflux snapshots, merge nudge PRs, render OLM catalog, verify FBC builds, CDN release, code freeze |
| Image Copy | 3.1–3.2 | Extract index image digests from Konflux and copy them to quay.io for QE testing |

Stages 4–6 (Test, Stage Release, Prod Release) are not yet implemented.

---

## Prerequisites

### Credentials

All credentials are stored in `.env` at the repo root. The skill sources this file at startup.

| Credential | What it gates | How to get it |
|-----------|---------------|---------------|
| `gh` CLI auth | All GitHub operations (every step) | `gh auth login` — the skill uses your existing `gh` session |
| `KONFLUX_SERVER` + `KONFLUX_TOKEN` | Cluster queries — checking snapshots, applications, components, releases | Get from your Konflux cluster credentials. Token must have read access to `tekton-ecosystem-tenant` namespace |
| `GITLAB_URL` + `GITLAB_TOKEN` | RPA checks (step 1.5), Pyxis checks (step 1.6), CLI product version (step 1.12), CLI CDN RP/RPA (step 1.13) | GitLab personal access token at `gitlab.cee.redhat.com` — **read-only** is sufficient |
| `JIRA_URL` + `JIRA_EMAIL` + `JIRA_TOKEN` | Jira-related steps (used by the audit skill, not heavily used in the release flow itself) | Jira API token |

If a credential is missing, the skill will note which steps are skipped rather than failing.

### Tooling

- **`gh`** — GitHub CLI, authenticated
- **`oc` / `kubectl`** — Kubernetes CLI, used for Konflux cluster queries (read-only)
- **`skopeo`** — Container image copying tool, needed for Stage 3 (Image Copy). If missing, the skill generates a copy script instead
- **`python3`** — Used for JSON parsing in verify commands

### Permissions

The skill follows a strict permission model:

- **Verify commands** (read-only) run automatically — GitHub API reads, `oc get`, GitLab `GET` requests
- **Execute commands** (write operations) are always shown to you first and run **only after you approve**

This means you can safely run the skill to check status without risk of changing anything. The skill will ask "Step X.Y: [description]. Execute?" before doing anything that modifies state.

---

## Quick Start

```
/one-click-release 1.21.3
```

The version argument must be `X.Y.Z` format. The skill derives everything else:

- `MAJOR_MINOR` = `1.21`
- `RELEASE_BRANCH` = `release-v1.21.x`
- `KONFLUX_NS` = `tekton-ecosystem-tenant` (always)

**What happens next:**

1. **Setup** — the skill parses the version, sources `.env`, fetches the hack repo release config (`config/downstream/releases/{MAJOR_MINOR}.yaml`), and creates the report directory tree under `reports/{MAJOR_MINOR}/{VERSION}/`
2. It runs Stage 1 (Config), checking each step in order
3. At each step, it runs a verify command to check the current state
4. If a step is done, it reports the status and moves on
5. If a step needs action, it shows you the execute command and asks for approval
6. If you approve, it runs the command and re-verifies
7. If you decline, it stops and reports the blocking step
8. After a stage completes (or is blocked), it writes a report file

You can run the skill repeatedly. It's idempotent — verify commands just check state, and execute commands check whether work is already done before acting.

---

## How It Works

### The Verify-Execute Loop

Every step follows the same pattern:

```
Verify → DONE? → yes → report, next step
                → no  → show execute command
                        → user approves → execute → re-verify
                        → user declines → stop, report blocker
```

This is the core interaction model. You never have to remember what to check or what command to run — the skill handles that. Your job is to review the execute commands and approve or decline.

### Sequential Stopping

Steps within each stage are sequential. The skill stops at the **first** step that isn't done and requires your action. This is by design: later steps depend on earlier ones. For example:

- Step 1.3 (Konflux config PR) can't be done until step 1.2 (release-manager PR) merges
- Step 2.2 (core snapshot) depends on step 2.1 (release PRs merged)
- Step 2.5 (FBC builds) depends on step 2.4 (OLM catalog render)

When you unblock a step (either by approving execution or doing it manually), re-run the skill to continue from where it left off.

### Read-Only Safety Net

All cluster operations (Konflux, GitLab) during verify are strictly read-only:

- Konflux: only `oc get` / `kubectl get` — never `apply`, `create`, or `delete`
- GitLab: only `GET` API requests — never `POST`, `PUT`, or `DELETE`
- GitHub: reads are automatic; writes (PR merge, workflow trigger) require approval

### Report Generation

Each stage writes a markdown report to `reports/{MAJOR_MINOR}/{VERSION}/{stage}/report_{timestamp}.md`. Reports include:

- A summary table with status, details, and links for every step
- Detailed per-step sections with timestamps, PR links, SHA comparisons
- A blocking step section if the stage didn't complete

Reports are timestamped so you can track progress across multiple runs.

---

## Stage 1: Config

**Purpose:** Get everything configured before builds can start. This is the longest stage (13 steps) because the release touches many systems: the hack repo, Konflux cluster config, GitLab release data, the operator repo's OLM and version settings, OPC component versions, and the CLI build pipeline.

**When to run:** At the start of a new patch or minor release, before any builds.

**Key dependency chain:**
```
1.1 (version bump) → 1.2 (merge PR) → 1.3 (Konflux config PR) → 1.4 (apply config on cluster)
                                                                        ↓
                                                              2.1 (release PRs need config applied first)

1.7 (OLM bundle version) ──── must complete BEFORE ──── 2.7 (code freeze)
1.8 (operator project.yaml) ─ must complete BEFORE ──── 2.2 (core snapshot / build validation)
1.12 (CLI product version) ── must complete BEFORE ──── 1.13 (CLI CDN RP/RPA)
```

Steps 1.5–1.13 are parallel in theory (independent of 1.1–1.4) but the skill checks them sequentially to give a complete status picture. The cross-stage dependencies above are critical — completing them out of order leads to broken builds or failed releases.

---

### Step 1.1: Create new patch version

**What it does:** Bumps the `release-tag` field in the hack repo's release config (`config/downstream/releases/{MAJOR_MINOR}.yaml`) to the target version. Also resets `code-freeze: false`.

**Why it matters:** The hack repo release config is the source of truth for the release version. Every downstream workflow reads this config to know which version to build. Until it's bumped, nothing else can proceed.

**What "done" looks like:** The `release-tag` field in the hack config equals your target version (e.g. `1.21.3`).

**How execution works:** Triggers the `release-new-patch.yaml` GitHub Actions workflow on `openshift-pipelines/hack`. This workflow creates a PR — it does not directly modify the config. The PR must be merged in the next step.

**What can go wrong:**
- The workflow run fails — check the workflow run URL in the output
- The release config for this major.minor doesn't exist yet — you'll need to create a new minor release config first

---

### Step 1.2: Merge release-manager PR

**What it does:** Merges the PR created by the `release-new-patch` workflow in step 1.1. The PR is titled `[bot: {MAJOR_MINOR}] Release Action: new-patch` and targets the `main` branch of `openshift-pipelines/hack`.

**Why it matters:** Until this PR merges, the hack repo still has the old version. Merging it triggers downstream automation (step 1.3).

**What "done" looks like:** The PR has state `MERGED`. The skill reports the merge timestamp.

**How execution works:** Merges the open PR via `gh pr merge` with rebase strategy.

**What can go wrong:**
- No PR exists — the step 1.1 workflow may still be running, or it failed silently. Check the workflow run.
- PR has merge conflicts — rare, but possible if someone else pushed to main. Resolve manually.

---

### Step 1.3: Konflux config PR merged

**What it does:** After step 1.2 merges, a `generate-konflux` GitHub Actions workflow runs automatically on push to main. It generates Konflux Application and Component YAML files and creates a PR to add them.

**Why it matters:** These YAML files define what Konflux will build — which repos, which branches, which Dockerfiles. Without them, Konflux doesn't know about this release.

**What "done" looks like:** The generated PR (on branch `actions/update/hack-update-konflux-main-{MAJOR_MINOR}`) has state `MERGED`.

**How execution works:** Two options:
1. Wait for the `merge-hack-pull-requests` hourly workflow to auto-merge it (PRs with the `hack` label are auto-merged)
2. Merge immediately via the skill

**What can go wrong:**
- The `generate-konflux` workflow hasn't run yet — it triggers on push to main, so wait a few minutes after step 1.2
- The generated config has unexpected changes — review the PR diff before merging

---

### Step 1.4: Konflux config applied on cluster

**What it does:** Verifies that the Konflux cluster state matches what the hack repo generated. Checks two levels:

- **1.4a: Applications** — compares directories in the hack repo's `.konflux/openshift-pipelines/{MM_DASHED}/` against Applications on the cluster
- **1.4b: Components per Application** — for each Application, compares the component directories against Components on the cluster

**Naming convention:** Dots in OCP version numbers become dashes on the cluster (e.g. hack directory `index-4.17` maps to cluster Application `index-4-17`). The skill handles this conversion automatically.

**Why it matters:** The cluster must have the right Applications and Components before builds can run. If a Component is missing or assigned to the wrong Application, builds will fail or produce images in the wrong snapshot.

**What "done" looks like:** Every hack repo directory has a matching cluster resource. Status is `OK` for all entries.

**How execution works:** Clones the hack repo and runs `kubectl apply` on the generated YAML files. Excludes `role.yaml` and `service-account.yaml` (these are one-time admin resources requiring elevated RBAC).

**What can go wrong:**
- `DRIFT` status — a Component exists but is assigned to the wrong Application. The execute command re-applies the correct config.
- `MISSING` status — a Component doesn't exist on the cluster at all. The execute command creates it.
- RBAC errors — `role.yaml` and `service-account.yaml` are excluded from the execute command because they require elevated RBAC. These are one-time admin resources that only need setup when onboarding a brand-new release stream (not typical for patch releases). If they do need updating, contact a cluster admin.

**Critical dependency:** Do NOT merge component PRs in downstream repos (step 2.1) before this config is applied. Pipelines triggered by those PRs will fail without the correct config on the cluster.

---

### Step 1.5: RPA in konflux-release-data

**What it does:** Checks that ReleasePlanAdmission (RPA) YAML files exist in the `konflux-release-data` GitLab repo for this version. RPAs tell Konflux how to release images — which registry, which signing keys, which advisory.

**Why it matters:** Without RPAs, Konflux builds will complete but can't be released. The release step will fail with "no matching ReleasePlanAdmission."

**What "done" looks like:** RPA files exist for all four categories: `core`, `bundle`, `fbc`, `cdn` (each with `stage` and `prod` variants = 8 files total).

**How execution works:** MANUAL. This cannot be automated — you must copy RPAs from the hack repo's `.konflux/` directory to the `konflux-release-data` GitLab repo via a merge request.

**What can go wrong:**
- Missing RPAs for a specific category — check against a reference MR from a previous release
- GitLab token not configured — the skill will use a GitHub-based fallback search, but verification is less reliable

---

### Step 1.6: Pyxis config

**What it does:** Checks that Pyxis repo configurations exist for `openshift-pipelines` container images. Pyxis is Red Hat's container certification/metadata system.

**Why it matters:** Pyxis config defines how images appear in the Red Hat container catalog. Missing entries mean images won't be listed.

**What "done" looks like:** Config entries exist under `products/openshift-pipelines/`. For patch releases, existing config from a previous patch is usually sufficient — new entries are only needed when shipping new component images.

**How execution works:** MANUAL. Copy Pyxis configuration from the hack repo to the `pyxis-repo-configs` GitLab repo.

**What can go wrong:**
- This step is often a no-op for patch releases — existing config covers the same images
- For minor releases with new components, this must be done before the first build

---

### Step 1.7: OLM bundle version in olm/config.yaml

**What it does:** Checks that the operator repo's `olm/config.yaml` on the release branch lists the target version as the last bundle entry. The `render-olm-catalog` script reads this file to generate catalog JSON.

**Why it matters:** If the version is missing from `olm/config.yaml`, the `render-olm-catalog` script produces **empty `catalog.json` files**. This breaks all index builds silently — they'll build but ship empty catalogs, meaning the operator won't appear in the OCP OperatorHub.

**What "done" looks like:** The last `version:` entry in `olm/config.yaml` equals your target version.

**How execution works:** Creates a PR on the release branch of `openshift-pipelines/operator` that replaces the previous patch version with the target version.

**What can go wrong:**
- The `render-catalog.sh` CI check auto-updates `catalog-template.json` files and may re-add the old version — if so, remove it again
- Initial CI failure on a stale image SHA is expected — the image gets updated by later steps
- **Must be done BEFORE code freeze.** The `update-sources` workflow is disabled during freeze (`if: false`), so if you miss this, you'll need to manually re-enable the workflow

---

### Step 1.8: Operator project.yaml version

**What it does:** Checks that the operator's `project.yaml` on the release branch has `versions.current` set to the target version. The `previous` field is set to what was previously `current`.

**Why it matters:** The operator binary reads its version from `project.yaml` at build time. If this isn't updated before the final image build, the operator will report the wrong version string in production.

**What "done" looks like:** `versions.current` equals your target version.

**How execution works:** Creates a PR on the release branch that updates both `current` and `previous` fields in `project.yaml`.

**What can go wrong:**
- If merged after the operator image is already built, the image ships the wrong version. Requires a rebuild.
- Make sure this completes before build validation (Stage 2)

---

### Step 1.9: OPC version.json

**What it does:** Checks that the OPC (OpenShift Pipelines Client) repo's `pkg/version.json` on the release branch has current upstream component versions. Compares each component (pac, tkn, results, manualapprovalgate, assist) against its latest upstream release tag.

**Why it matters:** OPC bundles multiple CLI tools. If component versions are stale, the CLI ships outdated binaries that may have known bugs or missing features.

**What "done" looks like:** All components show status `CURRENT` and the `opc` field equals the target version.

**How execution works:** MANUAL. The user must update `pkg/version.json`, update `go.mod` dependencies, run `go mod tidy && go mod vendor`, and create a PR targeting the release branch.

**What can go wrong:**
- A component version in `version.json` has a different major/minor series than upstream (e.g. `0.42.x` vs `0.43.x`) — the skill flags this as `REVIEW` instead of `OUTDATED`, because it may be intentional for this release branch
- Dependency conflicts in `go.mod` — sometimes upstream bumps require cascading updates

---

### Step 1.10: p12n-opc sync

**What it does:** Checks that the `p12n-opc` repo's `upstream/pkg/version.json` matches OPC's `pkg/version.json` on the release branch. Also compares the HEAD SHAs of both repos.

**Why it matters:** `p12n-opc` is a productization wrapper around OPC. Its `upstream/` directory must be an exact mirror. If they diverge, the productized CLI binary won't match what was tested upstream.

**What "done" looks like:** The `version.json` files in both repos are identical.

**How execution works:** MANUAL. Sync the `upstream/` directory in `p12n-opc` to match OPC on the release branch and create a PR.

---

### Step 1.11: serve-tkn-cli submodules

**What it does:** Checks that the `serve-tkn-cli` repo's git submodules under `sources/` point to the correct commit SHAs. Each submodule tracks a specific branch in its upstream repo (CLI tracks `tektoncd/cli`, OPC tracks `openshift-pipelines/opc`, PAC tracks `openshift-pipelines/pipelines-as-code`).

**Why it matters:** `serve-tkn-cli` builds the actual CLI binaries that ship on the CDN. If submodule SHAs are stale, the binaries are built from old code.

**What "done" looks like:** Every submodule's recorded SHA matches the HEAD of its tracking branch.

**How execution works:** MANUAL. Run `git submodule update --init --remote --force --checkout` on the release branch. Note: plain `git submodule update` (without `--remote --force --checkout`) only checks out the already-recorded SHA and will report "already up to date" even when the tracking branch has moved.

---

### Step 1.12: CLI product version config

**What it does:** Checks that a product version YAML file exists in the `konflux-release-data` GitLab repo at `data/external/developer-portal/openshift-pipelines/{VERSION}.yaml`.

**Why it matters:** This defines how the release version appears on the Red Hat Developer Portal. The CDN RP/RPA (step 1.13) references this version — you can't create a CDN release plan for a version that doesn't exist.

**What "done" looks like:** The file exists with `versionName` matching the target version.

**How execution works:** MANUAL. Create a GitLab merge request adding the version YAML file. Key fields:
- `ga` — `false` for patches, `true` for minor releases and patches on the latest minor
- `releaseDate` — set to the planned release date
- `hidden` / `invisible` — both `false`

**Critical dependency:** Must be merged BEFORE step 1.13 (CLI CDN RP/RPA).

**Skipped if:** `GITLAB_TOKEN` is not configured.

---

### Step 1.13: CLI CDN RP/RPA

**What it does:** Checks for CDN-specific ReleasePlanAdmission and ReleasePlan resources in the `konflux-release-data` GitLab repo. These are separate from the image-based RPAs in step 1.5 — they handle CLI binary releases to the CDN.

**Why it matters:** Without CDN RP/RPAs, CLI binaries can't be released to the developer download portal, even if they're built.

**What "done" looks like:** CDN RPA files exist for `stage` and `prod`, and corresponding ReleasePlan files exist in both the tenant config and auto-generated directories.

**How execution works:** MANUAL. Copy RP/RPA from a previous version and update the version numbers.

**Critical dependency:** Step 1.12 must be merged first — the ReleasePlanAdmission references the product version, which must already exist.

**Skipped if:** `GITLAB_TOKEN` is not configured.

---

## Stage 2: Build

**Purpose:** Get all component images built, verified, and ready for release. This stage processes release PRs across all downstream repos, waits for Konflux to produce snapshots with the correct commits, handles the nudge/OLM/FBC pipeline, releases CLI binaries to CDN, and sets code freeze.

**When to run:** After Stage 1 (Config) is complete — all configuration must be in place before builds start.

**Key dependency chain:**
```
2.1 (release PRs) → 2.2 (core snapshot) → 2.3 (nudge PRs) → 2.4 (OLM render) → 2.5 (FBC builds) → 2.7 (code freeze)
                     ↓
                     2.6 (CDN release — uses core snapshot)
```

---

### Step 2.1: Process release PRs

**What it does:** Searches for open PRs across all `openshift-pipelines` repos on the release branch that have release-related labels (`hack`, `upstream`, `automated`). For each PR, it checks CI status and mergeability, then offers to merge the ones that are ready.

**Why it matters:** When the hack repo generates new config (Stage 1), it creates PRs in downstream component repos to update Dockerfiles, dependencies, and build config. These PRs must all be merged before Konflux can build the release. This is the "fan-out" moment where config changes propagate to ~20 component repos.

**What "done" looks like:** No open PRs remain on the release branch with release labels.

**How execution works:** For ready PRs: adds `lgtm`, `approved`, and `one-click-release` labels, approves the PR, and sets auto-merge with delete-branch and rebase. For behind PRs: rebases the branch. For CI-failing PRs: reports the failures and skips.

**What can go wrong:**
- **CI FAILING** — the most common blocker. Could be a flaky test, a real build failure introduced by the config change, or a missing dependency. Check the failing check's URL.
- **CONFLICT** — means someone merged something else to the release branch between the PR creation and now. The PR needs manual rebase.
- **CI PENDING** — normal shortly after PR creation. Wait a few minutes and re-run.

---

### Step 2.2: Wait for core snapshot

**What it does:** Checks the latest Konflux "core" snapshot for this release and verifies that each component's revision matches the HEAD of the release branch in its source repo. The core snapshot is a manifest listing every component image and the source commit it was built from.

**Why it matters:** The core snapshot is the input to everything downstream — nudge PRs update the operator with SHAs from this snapshot, and the snapshot is what gets released. If a component is built from a stale commit, the release ships old code.

**What "done" looks like:** All non-operator repos in the snapshot have revision matching their release branch HEAD. The operator is excluded from this check because nudge commits (step 2.3) will push it ahead.

**How execution works:** For stale repos, pushes a placeholder file to `.konflux/patches/.placeholder` to trigger a Konflux rebuild of all components in that repo.

**What can go wrong:**
- **STALE** status — a component was built from an older commit. Usually means a PR merged after the last build. The execute command triggers a rebuild.
- **SPLIT** status — the same repo has components built from different revisions. This happens when a multi-component repo had only some components rebuild. The execute command fixes this by triggering a full rebuild.
- Rebuilds take time (5–15 minutes typically). The skill doesn't wait — re-run it after the pipelines finish.

---

### Step 2.3: Process nudge PRs

**What it does:** After core components build, Konflux creates "nudge" PRs in the operator repo. These PRs update image SHA references in `project.yaml` to point to the newly built images. The skill searches for open PRs with the `konflux-nudge` label and offers to merge them.

**Why it matters:** Nudge PRs are how the operator learns about new component images. Without merging them, the operator's `project.yaml` still references old image SHAs, and the operator would deploy stale components.

**What "done" looks like:** No open nudge PRs remain.

**How execution works:** Same as step 2.1 — label, approve, auto-merge. For failing CI, suggests `/retest` as a PR comment (many nudge PR CI failures are transient).

**What can go wrong:**
- Nudge PRs arrive in waves — as each component finishes building, Konflux creates a separate nudge PR. You may need to run this step multiple times.
- CI failures on nudge PRs are often transient — retesting usually resolves them
- Nudge PRs can conflict with each other if multiple arrive for the same `project.yaml` section. Merge one at a time and rebase others.

---

### Step 2.4: OLM catalog render

**What it does:** Verifies that the `render-olm-catalog` GitHub Actions workflow has run successfully on the operator repo's release branch after the latest nudge PRs merged. This workflow generates `catalog.json` files for each supported OCP version.

**Why it matters:** The catalog JSON files are what OLM reads to discover the operator in OperatorHub. Without a fresh render after nudge PRs merge, the catalog may reference stale image SHAs — the operator would install but run old code.

**What "done" looks like:** A `render-olm-catalog` workflow run completed successfully after the latest nudge PRs merged, and catalog JSON commits are present on the release branch.

**How execution works:** Triggers the `render-olm-catalog` workflow with the release branch as a parameter.

**What can go wrong:**
- The workflow runs daily at 1 AM UTC automatically, but after nudge PRs it may not have run yet. Trigger it manually.
- Workflow failure usually means the OLM config from step 1.7 is wrong or missing

---

### Step 2.5: Wait for FBC build

**What it does:** Verifies that bundle and FBC (File-Based Catalog) index snapshots exist and are current. Checks two levels:

- **2.5a: Bundle snapshot** — the operator bundle image, built from the operator repo
- **2.5b: Index snapshots** — one per supported OCP version (e.g. `index-4-14`, `index-4-16`, etc.)

**Why it matters:** The index images are what OCP clusters pull to populate OperatorHub. They're the final built artifact before release. If they're stale or missing, the release has nothing to ship.

**What "done" looks like:** The bundle snapshot exists — the skill checks whether the bundle revision exactly matches operator HEAD, and if not, examines the commits between them. If all intervening commits are automated (catalog renders, CSV updates, nudge merges), the gap is accepted as normal. All index snapshots with non-empty config directories must be current.

**How execution works:** For stale builds, pushes placeholder files to the appropriate `.konflux/olm-catalog/bundle/` or `.konflux/olm-catalog/index/` directory to trigger Konflux rebuilds.

**What can go wrong:**
- Index apps with empty config dirs (no components for this release) will have no snapshot — this is expected for OCP versions that don't ship this operator version
- Bundle revision may not exactly match operator HEAD — this is OK if the gap is only automated commits (catalog renders, CSV updates). The skill checks this.

---

### Step 2.6: CDN production release

**What it does:** Creates a Konflux Release CR that triggers the CDN release pipeline. This releases CLI binaries (tkn, opc, pac) to the Red Hat developer download portal.

**Why it matters:** CLI binaries are shipped separately from container images. Users download them from the CDN, not from a container registry. This step gets the binaries into the download pipeline.

**What "done" looks like:** A CDN production release exists with status `Succeeded`.

**How execution works:** The skill generates a Release YAML file but does **not** apply it directly (the Konflux cluster is treated as read-only for direct `create` operations). Instead, it saves the YAML to `release-{VERSION}-cdn-prod.yaml` and shows you the `oc create -f` command to run. You apply it from outside the skill.

**Why production directly (not stage first):** Stage CDN releases require manual product version configuration in the stage CDN environment. Going directly to production with `invisible: true` skips this overhead. After the release succeeds, you update the product version YAML to set `invisible: false`.

**What can go wrong:**
- The core snapshot must be verified (step 2.2) before creating a CDN release — the release references the snapshot
- Release pipeline failures are typically RPA/RP configuration issues — check step 1.12 and 1.13

---

### Step 2.7: Code freeze

**What it does:** Sets `code-freeze: true` in the hack repo's release config for this major.minor version.

**Why it matters:** When code freeze is set, the `update-sources` workflow is disabled on the release branch. This prevents upstream changes from landing on the release branch while QE is testing. Without freeze, a new upstream commit could break the already-verified build.

**What "done" looks like:** The `code-freeze` field in the hack config is `true`.

**How execution works:** Creates a PR in `openshift-pipelines/hack` that changes `code-freeze: false` to `code-freeze: true`, then merges it.

**What can go wrong:**
- If freeze is set too early (before all builds complete), you'll need to temporarily un-freeze to let `update-sources` run
- The `create-new-patch` workflow (step 1.1) resets freeze to `false`, so this must be re-set for each new patch

---

## Stage 3: Image Copy

**Purpose:** Copy FBC-built index images from the Konflux build registry to quay.io so QE can access them for testing.

**Note:** This is a **temporary stage** standing in for the full Test stage. Once the test automation is built, this will be replaced by a more comprehensive testing workflow.

**When to run:** After Stage 2 (Build) completes — all snapshots must be current and code freeze set.

---

### Step 3.1: Extract index image digests

**What it does:** For each index application (one per supported OCP version), fetches the latest push snapshot from Konflux and extracts the `containerImage` reference (a `quay.io/...@sha256:...` digest).

**Why it matters:** The Konflux build registry is not directly accessible to QE environments. Index images need to be copied to `quay.io/openshift-pipeline/` with a version-tagged reference for QE to consume.

**What "done" looks like:** All index apps with snapshots have `containerImage` references extracted. The skill reports a mapping table from index app to OCP version to quay target tag.

**How execution works:** This step is read-only — it just extracts the information needed for step 3.2.

**What can go wrong:**
- A snapshot exists but has no `containerImage` — the FBC build pipeline may not have completed. Go back to step 2.5.
- Some index apps have no snapshot at all — this is expected for OCP versions that don't ship this operator version (e.g. too old or too new)

---

### Step 3.2: Copy index images to quay.io

**What it does:** Uses `skopeo` to copy each index image from the Konflux build registry to `quay.io/openshift-pipeline/pipelines-index-{OCP_VERSION}:v{VERSION}-stage`.

**Why it matters:** QE tests against images on quay.io. Until this step runs, the images exist only in the Konflux internal registry.

**What "done" looks like:** All index images exist on quay.io with correct digests matching the source.

**How execution works:** Two paths:
1. If `skopeo` is available, runs `skopeo copy --all` with `--preserve-digests` for each image
2. If `skopeo` is not available, generates a shell script at `scripts/copy-index-images-{VERSION}-stage.sh` that you can run yourself

**What can go wrong:**
- Quay.io authentication — `skopeo` needs credentials to push to the `openshift-pipeline` org. Use `podman login quay.io` or set up `~/.docker/config.json`
- Digest mismatch — if the source image was rebuilt between extract and copy. Re-run from step 3.1.

---

## Stage 4: Test

> **Not yet implemented.** This stage will automate QE handover and test verification.

Based on the pipeline design, this stage will:

- Notify QE that index images are available for testing on quay.io
- Track test status across QE environments
- Gate progression to Stage Release on test results

Currently, QE coordination happens manually via Slack and Jira after Stage 3 completes.

---

## Stage 5: Stage Release

> **Not yet implemented.** This stage will automate staging releases.

Based on the pipeline design, this stage will:

- Create Konflux Release CRs for stage environment (core, bundle, FBC)
- Monitor release pipeline status
- Verify images land in the stage registry
- Validate stage advisory creation

The Release YAML files use `generateName` and must be applied with `oc create -f`, not `oc apply -f`.

---

## Stage 6: Prod Release

> **Not yet implemented.** This stage will automate production releases.

Based on the pipeline design, this stage will:

- Create Konflux Release CRs for production environment
- Monitor release pipeline and advisory status
- Verify images in production registry
- Send Slack notifications on completion
- Update product version visibility (set `invisible: false`)

---

## Reports

### Where reports go

```
reports/{MAJOR_MINOR}/{VERSION}/{stage}/report_{timestamp}.md
```

Example: `reports/1.21/1.21.3/config/report_2026-07-08_14-30-00_IST.md`

Each run generates a new timestamped report, so you can see the history of a release's progression.

### Report structure

Every stage report follows the same format:

1. **Header** — version, branch, generation timestamp
2. **Summary table** — one row per step with status, details, and links
3. **Step details** — expanded information per step (SHAs, PR links, comparison tables)
4. **Blocking step** — which step stopped progress and why (omitted if all steps are done)

### Status values

| Status | Meaning |
|--------|---------|
| `DONE` | Step verified successfully |
| `ACTION NEEDED` | Step requires execution — either automated (user approves) or manual |
| `SKIPPED` | Step was skipped (missing credentials, not applicable for this release type) |

### Link formats

Reports use consistent link formatting:

- GitHub PRs: `repo [#NUM](URL)` — e.g. `hack [#123](https://github.com/openshift-pipelines/hack/pull/123)`
- Workflow runs: `[workflow-name](URL)`
- GitLab MRs: `[MR](URL)`
- Commit SHAs: `[12-char-short](URL)`
- No links available: `—`

---

## Troubleshooting

### Missing credentials

**Symptom:** Steps are reported as `SKIPPED` with a note about missing credentials.

**Fix:** Check your `.env` file has all required variables. The most common missing ones:
- `KONFLUX_TOKEN` — expires periodically, needs refresh from your cluster credentials
- `GITLAB_TOKEN` — steps 1.5, 1.6, 1.12, 1.13 are skipped without it

### Stale Konflux snapshots

**Symptom:** Step 2.2 or 2.5 shows `STALE` status for one or more repos.

**Cause:** A PR merged after the last Konflux build, so the snapshot has an older commit.

**Fix:** Approve the execute command — it pushes a placeholder file to trigger a rebuild. Wait 5–15 minutes for Konflux pipelines to complete, then re-run the skill.

### CI failures on release PRs

**Symptom:** Step 2.1 reports PRs with `CI FAILING` status.

**Common causes:**
- Flaky tests — re-run CI (`/retest` comment) or merge and let the build re-trigger
- Missing config — step 1.4 (Konflux config on cluster) wasn't completed before builds started
- Dependency issues — upstream bumped something that breaks the build. Needs manual investigation.

### Nudge PR conflicts

**Symptom:** Step 2.3 shows nudge PRs with `CONFLICT` status.

**Cause:** Multiple nudge PRs target the same `project.yaml` lines. Konflux creates one nudge PR per component build, and they can overlap.

**Fix:** Merge one nudge PR at a time. After each merge, the remaining PRs need rebase. The skill can rebase `BEHIND` PRs.

### Empty OLM catalog

**Symptom:** Step 2.5 succeeds but index images have empty catalogs.

**Cause:** Step 1.7 (OLM bundle version) wasn't completed, so `render-olm-catalog` produced empty `catalog.json` files.

**Fix:** Go back to step 1.7, fix `olm/config.yaml`, then re-run the OLM render (step 2.4) and wait for FBC rebuilds (step 2.5).

### Code freeze blocks updates

**Symptom:** You need to land a fix on the release branch, but `update-sources` is disabled because code freeze is set.

**Fix:** Temporarily unset code freeze (change `code-freeze: true` to `false` in the hack config), land the fix, wait for the build to propagate, then re-set code freeze (step 2.7).

### "oc create" vs "oc apply"

**Symptom:** `oc apply -f release-*.yaml` fails with an error about `generateName`.

**Fix:** Use `oc create -f`, not `oc apply -f`. Release YAMLs use `generateName` which is incompatible with `apply`.

### Skill stops at step I already fixed manually

**Symptom:** You fixed something outside the skill, but re-running still stops at the same step.

**Cause:** The skill only checks system state via verify commands — it doesn't cache results. If the verify still shows "not done," the manual fix may not have fully propagated.

**Fix:** Check that the manual change is visible in the system the verify command queries (GitHub API, Konflux cluster, GitLab). Merges, workflow triggers, and cluster syncs all have propagation delays.
