# Jira Gaps: Release 1.21.3

Jiras linked to commits in 1.21.3 that are missing `fixVersion: Pipelines 1.21.3` and/or release note fields.
Excludes Jiras "CVE" in their name and Jiras with all fields already set.

## Overview

| Jira | Component | Status | fixVersion 1.21.3 | RN Fields | Linked via | Commits | PRs |
|------|-----------|--------|-------------------|-----------|------------|---------|-----|
| [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) | tektoncd-pipeline | No linked SRVKP | **MISSING** | **MISSING** | PR body | 8 | 1 |
| [SRVKP-11927](https://redhat.atlassian.net/browse/SRVKP-11927) | operator | Closed | **MISSING** | YES | Go tool (Jira custom field) | 1 | 1 |
| [SRVKP-11380](https://redhat.atlassian.net/browse/SRVKP-11380) | tektoncd-results | Dev Complete | **MISSING** | YES | Go tool (Jira custom field) | 2 | 1 |

---

## [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727)

**Summary:** (different Jira project — inaccessible)<br>
**Component:** tektoncd-pipeline (tektoncd/pipeline)<br>
**Status:** There is no SRVKP linked to this KONFLUX ticket.<br>
**Current fixVersions:** —<br>
**Linked via:** PR body — KONFLUX-6727 found in the body of upstream PR [#9665](https://github.com/tektoncd/pipeline/pull/9665)<br>
**Missing:** fixVersion `Pipelines 1.21.3`, Release Note fields (type + text)

| SHA | Message | Author | PR |
|-----|---------|--------|----|
| b728378 | Remove deprecated `// +build` directive from most files | Vincent Demeester | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) |
| c1810c5 | test: implement parallel/serial test categorization system | Vincent Demeester | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) |
| f7f9584 | refactor: add clock injection to cache for testing | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) |
| c639f88 | fix: Align cache configstore with framework implementation | Seunghun Shin | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) |
| 48765b2 | fix: resolve cache race with singleflight | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) |
| 192a6c1 | fix: remove corrupted cache entries on type error | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) |
| 6bbcfb1 | fix: use TextParser struct directly for prometheus/common v0.62.0 compatibility | Vibhav Bobade | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) |
| 6e0504b | fix: add missing @test:execution annotation to e2e test | Vibhav Bobade | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) |

## [SRVKP-11927](https://redhat.atlassian.net/browse/SRVKP-11927)

**Summary:** Operator should be able to pick the console-plugin version based on OCP Version<br>
**Component:** operator (tektoncd/operator)<br>
**Status:** Closed<br>
**Current fixVersions:** Pipelines 1.23.0<br>
**RN Type:** Enhancement · **RN Status:** Proposed<br>
**Linked via:** Go tool — Jira key found via PR's custom field (Git Pull Request link)<br>
**Missing:** fixVersion `Pipelines 1.21.3`

| SHA | Message | Author | PR |
|-----|---------|--------|----|
| 9d408fc | Feat: Console Plugin Image should be picked conditionally based OCP Version | Pramod Bindal | [#3386](https://github.com/tektoncd/operator/pull/3386) → [#3467](https://github.com/tektoncd/operator/pull/3467) |

## [SRVKP-11380](https://redhat.atlassian.net/browse/SRVKP-11380)

**Summary:** [Results CLI] Describe/logs operations slow with large record count due to substring matching<br>
**Component:** tektoncd-results (tektoncd/results)<br>
**Status:** Dev Complete<br>
**Current fixVersions:** Pipelines 1.21.2, Pipelines 1.22.2<br>
**RN Type:** Enhancement · **RN Status:** Proposed<br>
**Linked via:** Go tool — Jira key found via PR's custom field (Git Pull Request link)<br>
**Missing:** fixVersion `Pipelines 1.21.3`

| SHA | Message | Author | PR |
|-----|---------|--------|----|
| 8cda1d8 | perf(cli): use exact match for describe/logs command | divyansh42 | [#1283](https://github.com/tektoncd/results/pull/1283) → [#1336](https://github.com/tektoncd/results/pull/1336) |
| ae090bc | Merge pull request #1336 from divyansh42/cherrypick-11380-0.17 | Emil Natan | [#1283](https://github.com/tektoncd/results/pull/1283) → [#1336](https://github.com/tektoncd/results/pull/1336) |
