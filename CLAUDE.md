# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

osp-release-audit is a toolkit for auditing and managing OpenShift Pipelines releases. It combines Go CLI tools for commit/Jira analysis with Claude Code skills that automate the 16-step release verification workflow.

## Build & Run

```bash
go build ./...                    # build all CLI tools
go build ./cmd/audit              # build the main audit CLI
go run ./cmd/audit                # run audit directly
```

No test suite, linter config, or Makefile — standard Go tooling only.

## Architecture

**Go CLI tools** (`cmd/` + `internal/`):
- `cmd/audit` — collect commits across ~20 component repos, link to Jira tickets
- `cmd/jira-update` — re-fetch Jira metadata for existing audit data
- `cmd/snapshot-diff` — compare two Konflux snapshots
- `cmd/image-diff` — compare images between snapshots
- `internal/github` — GitHub API helpers (uses `go-github`)
- `internal/jira` — Jira REST API client

**Claude Code skills** (`.claude/skills/`) — the primary automation layer:

The *release checklist* is a 16-step sequential workflow split into sub-skills. Each sub-skill owns specific step numbers and report sections (see `REPORT.md` for ownership mapping):

| Sub-skill | Steps | What it checks |
|-----------|-------|----------------|
| `check-hack` | 1-2 | Hack repo release config |
| `check-konflux-config` | 3-5 | Konflux cluster config, RPA, Pyxis |
| `check-components` | 6-8 | Component PRs, OLM bundle version, upstream sync, nudges |
| `check-builds` | 9-10 | Operator version, SHA validation against downstream repos |
| `check-olm` | 11-12 | OLM catalog render, index images, code freeze |
| `check-releases` | 13-16 | Stage/prod releases, QA handover, advisory |

Standalone skills: `audit-release`, `build-diff`, `new-build`, `jira-gaps`, `cli-release`.

**Reports** (`reports/`): generated markdown/JSON output. Naming convention: `{type}-{version}.{md|json}`.

**Release yamls** (root): Konflux Release CRs for applying stage/prod releases via `oc create -f` (not `oc apply` — they use `generateName`).

## External Systems & Credentials

Credentials live in `.env`. The following systems are used:

- **GitHub**: `gh` CLI with existing auth — read/write
- **GitLab** (`gitlab.cee.redhat.com`): `GITLAB_URL` + `GITLAB_TOKEN` — **READ-ONLY only**
- **Konflux cluster**: `KONFLUX_SERVER` + `KONFLUX_TOKEN` — **READ-ONLY only** (`oc get`/`kubectl get`, never `apply`/`create`/`delete`)
- **Jira**: `JIRA_URL` + `JIRA_EMAIL` + `JIRA_TOKEN`
- **Quay**: no credentials configured

All Konflux operations target the `tekton-ecosystem-tenant` namespace.

## Skills

The release checklist can be run in full or by individual step groups:

- **release-checklist** (`.claude/skills/release-checklist/SKILL.md`) - full release checklist. Trigger: `/release-checklist`
- **release-checklist:check-hack** (`.claude/skills/release-checklist/check-hack/SKILL.md`) - hack repo config (steps 1-2). Trigger: `/release-checklist:check-hack`
- **release-checklist:check-konflux-config** (`.claude/skills/release-checklist/check-konflux-config/SKILL.md`) - Konflux config, RPA, Pyxis (steps 3-5). Trigger: `/release-checklist:check-konflux-config`
- **release-checklist:check-components** (`.claude/skills/release-checklist/check-components/SKILL.md`) - component PRs, OLM bundle version, upstream sync, nudges (steps 6-8). Trigger: `/release-checklist:check-components`
- **release-checklist:check-builds** (`.claude/skills/release-checklist/check-builds/SKILL.md`) - operator version and SHA validation (steps 9-10). Trigger: `/release-checklist:check-builds`
- **release-checklist:check-olm** (`.claude/skills/release-checklist/check-olm/SKILL.md`) - OLM render/index status and code freeze (steps 11-12). Trigger: `/release-checklist:check-olm`
- **release-checklist:check-releases** (`.claude/skills/release-checklist/check-releases/SKILL.md`) - stage/prod releases, QA, advisory (steps 13-16). Trigger: `/release-checklist:check-releases`

When the user types any of the above triggers, invoke the Skill tool with the corresponding `skill` name before doing anything else.

Other skills:

- **build-diff** (`.claude/skills/build-diff/SKILL.md`) - compare two Konflux snapshots. Trigger: `/build-diff`
- **audit-release** (`.claude/skills/audit-release/SKILL.md`) - enrich release commits with Jira links. Trigger: `/audit-release`
- **new-build** (`.claude/skills/new-build/SKILL.md`) - verify build health and generate diff for QE. Trigger: `/new-build`
- **jira-gaps** (`.claude/skills/jira-gaps/SKILL.md`) - identify Jiras missing fixVersion or release note fields. Trigger: `/jira-gaps`
- **cli-release** (`.claude/skills/cli-release/SKILL.md`) - track CLI release progress for serve-tkn-cli binaries to CDN. Trigger: `/cli-release`

When the user types any of the above triggers, invoke the Skill tool with the corresponding `skill` name before doing anything else.

## Key Conventions

- **Report formatting**: PR numbers as `[#NUM](URL)`, commit SHAs as `[SHORT](URL)` (12-char short), all timestamps as absolute local time with timezone (e.g. `2026-07-07 10:30 IST`).
- **Release yamls use `generateName`**: apply with `oc create -f`, not `oc apply -f`.
- **Skill step ownership**: each CHECK file owns specific step rows and detail sections in the report. Never modify rows belonging to other CHECK files (see `REPORT.md`).
- **Early stop rule**: checklist steps are sequential. Stop at the first non-DONE step.
