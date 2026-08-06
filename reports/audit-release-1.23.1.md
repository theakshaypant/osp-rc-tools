# Release Audit: 1.23.1

**Source:** `reports/release_1.23.1.json`<br>
**Total commits:** 27 across 10 active components (18 tracked; 3 with errors)<br>
**Already linked (by Go tool):** 0 · **Newly linked:** 8 · **Unlinked:** 9 · **Bots (skipped):** 10<br>

## Index

| Component | Commits | Jira | Unsynced | Unmatched Jiras |
|-----------|---------|------|----------|-----------------|
| [operator](#operator) | 5 | 0/4 | 0 | 1 |
| [opc](#opc) | 3 | 1/1 | 0 | 0 |
| [console-plugin-pf5](#console-plugin-pf5) | 1 | 1/1 | 0 | 5 |
| [tekton-assist](#tekton-assist) | Error | — | — | 0 |
| [tektoncd-hub](#tektoncd-hub) | 4 | 4/4 | 0 | 1 |
| [tektoncd-pruner](#tektoncd-pruner) | 0 | — | — | 0 |
| [tektoncd-triggers](#tektoncd-triggers) | 0 | — | — | 0 |
| [syncer-service](#syncer-service) | 0 | — | — | 0 |
| [manual-approval-gate](#manual-approval-gate) | 2 | 2/2 | 0 | 0 |
| [tektoncd-pipeline](#tektoncd-pipeline) | 4 | — | 0 | 1 |
| [multicluster-proxy-aae](#multicluster-proxy-aae) | 0 | — | — | 0 |
| [git-init](#git-init) | Error | — | — | 0 |
| [tekton-caches](#tekton-caches) | 0 | — | — | 0 |
| [tekton-kueue](#tekton-kueue) | 0 | — | — | 0 |
| [tektoncd-chains](#tektoncd-chains) | 1 | — | 0 | 0 |
| [tektoncd-cli](#tektoncd-cli) | 4 | 0/2 | 0 | 0 |
| [pipelines-as-code](#pipelines-as-code) | 2 | 0/2 | 0 | 0 |
| [tektoncd-results](#tektoncd-results) | 1 | 0/1 | 0 | 0 |

> **Jira column:** linked non-bot commits / total non-bot commits. `—` means all commits were bots.

---

<details>
<summary><h2>operator</h2> — 5 commits, 0 linked, 1 dependabot</summary>

**Upstream:** tektoncd/operator · **Branch:** release-v0.80.x → release-v1.23.x · **Commits:** 5 · **Unsynced:** 0 · **Dependabot:** 1

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| [ca311009cf97](https://github.com/tektoncd/operator/commit/ca311009cf97d2e9eaa7ab70f1a77cb2f90a44d6) | chore: update third-party image digests | Divyanshu Agrawal | [#3778](https://github.com/tektoncd/operator/pull/3778) | — | — |
| [b4cf798f9d9a](https://github.com/tektoncd/operator/commit/b4cf798f9d9ae765190051ff7632573366086588) | chore: update third-party image digests | Divyanshu Agrawal | [#3788](https://github.com/tektoncd/operator/pull/3788) | — | — |
| [3648a8cccf02](https://github.com/tektoncd/operator/commit/3648a8cccf02da6df76ed79f3f604ee8a8b66445) | chore: bump component versions | Vincent Demeester | [#3759](https://github.com/tektoncd/operator/pull/3759) | — | — |
| [793d41b09273](https://github.com/tektoncd/operator/commit/793d41b09273c01dd65820cdc6aa085d648ba7fe) | chore: bump hub from v1.24.1 to v1.24.2 | Vincent Demeester | [#3806](https://github.com/tektoncd/operator/pull/3806) | — | — |

### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-10831](https://redhat.atlassian.net/browse/SRVKP-10831) | [QA] 3 webhooks are missing namespace owner reference | Release Pending |

</details>

<details>
<summary><h2>opc</h2> — 3 commits, 1 linked, 2 bots</summary>

**Upstream:** openshift-pipelines/opc · **Branch:** release-v1.23.x → release-v1.23.x · **Commits:** 3 · **Unsynced:** 0 · **Bots (github-actions):** 2

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| [ad3cba96ddd1](https://github.com/openshift-pipelines/opc/commit/ad3cba96ddd19da3fb76fc90b1590eeae183fa5c) | fix(cve): fix few critical CVEs | Shiv Verma | [#503](https://github.com/openshift-pipelines/opc/pull/503) | [SRVKP-13051](https://redhat.atlassian.net/browse/SRVKP-13051) *(keyword)* | **Type:** CVE<br>**Text:** Updates the OPC component to address critical CVEs by upgrading the Hub dependency to v1.24.2 and pgx/v5 to v5.10.0. *(generated)* |

</details>

<details>
<summary><h2>console-plugin-pf5</h2> — 1 commit, 1 linked, 0 bots</summary>

**Note:** Component encountered an error fetching branch metadata (p12n-console-plugin-pf5 returned 404); commit was manually provided.<br>
**Commits:** 1 · **Bots:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| [71b270edbf9a](https://github.com/openshift-pipelines/console-plugin/commit/71b270edbf9a66d2a750367806cd61d37844f3b2) | manual cherrypick of refresh-yarn-lockfile workflow into main | Arvind | [#1207](https://github.com/openshift-pipelines/console-plugin/pull/1207) | [SRVKP-12826](https://issues.redhat.com/browse/SRVKP-12826) *(PR body)* | **Type:** CVE<br>**Text:** Adds an automated workflow to refresh the yarn lockfile weekly for the console-plugin, keeping JavaScript dependencies current and reducing exposure to known vulnerabilities. *(generated)* |

### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-12879](https://redhat.atlassian.net/browse/SRVKP-12879) | OCP 4.22 \| Incorrect tab is highlighted in Overview page, pipelines list section when switching between 'Per pipeline' & 'Per Repository' | Dev Complete |
| [SRVKP-12178](https://redhat.atlassian.net/browse/SRVKP-12178) | In pipelines widget of Overview page, search text in pipeline widget is not honored | Tasking and Estimation |
| [SRVKP-12111](https://redhat.atlassian.net/browse/SRVKP-12111) | OSP UI redirects to incorrect namespace/page for resolver Tasks causing misleading 404 | Dev Complete |
| [SRVKP-11468](https://redhat.atlassian.net/browse/SRVKP-11468) | Logs misalignment in Pipeline run with logs having mix of tabs & spaces tabular style | Closed |
| [SRVKP-11467](https://redhat.atlassian.net/browse/SRVKP-11467) | Logs misalignment in Pipeline run with logs of ASC\|\| tabular style | Closed |

</details>

<details>
<summary><h2>tekton-assist</h2> — Error</summary>

> **Error:** failed to get head SHA at from-date: failed to list head file commits for openshift-pipelines/p12n-tekton-assist on branch release-v1.23.x: 404 Not Found

</details>

<details>
<summary><h2>tektoncd-hub</h2> — 4 commits, 4 linked, 0 dependabot</summary>

**Upstream:** openshift-pipelines/hub · **Branch:** release-v1.24.2 → release-v1.23.x · **Commits:** 4 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| [9e79a4912553](https://github.com/openshift-pipelines/hub/commit/9e79a4912553e9dd74fd70a01a7d3e0cc5e0b8db) | fix(cve): fix CVE vulnerabilities in the UI and Swagger | Shiv Verma | — | [SRVKP-13051](https://redhat.atlassian.net/browse/SRVKP-13051) *(keyword)* | **Type:** CVE<br>**Text:** Addresses CVE vulnerabilities in the Hub UI and Swagger components. *(generated)* |
| [ab9e68015562](https://github.com/openshift-pipelines/hub/commit/ab9e68015562c22c89c8d90e6a016729cacda853) | fix(cve): fix critical CVEs and update go version to 1.26.4 | Shiv Verma | — | [SRVKP-13051](https://redhat.atlassian.net/browse/SRVKP-13051) *(keyword)* | **Type:** CVE<br>**Text:** Fixes critical CVEs and updates the Go runtime to version 1.26.4 in the Hub component. *(generated)* |
| [8d545e4a9e99](https://github.com/openshift-pipelines/hub/commit/8d545e4a9e994643f12171976094c9d5df1b6177) | update the go version in dockerfile | Shiv Verma | — | [SRVKP-13051](https://redhat.atlassian.net/browse/SRVKP-13051) *(keyword)* | **Type:** CVE<br>**Text:** Updates the Go version in the Hub Dockerfile as part of CVE remediation. *(generated)* |
| [82bd871583da](https://github.com/openshift-pipelines/hub/commit/82bd871583da83bc46a3cb793154e66408acf69b) | fix(cve): fix few critical CVEs in API | Shiv Verma | — | [SRVKP-13051](https://redhat.atlassian.net/browse/SRVKP-13051) *(keyword)* | **Type:** CVE<br>**Text:** Fixes critical CVEs in the Hub API component. *(generated)* |

### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-13051](https://redhat.atlassian.net/browse/SRVKP-13051) | Hub release with CVEs fixes for OSP 1.23.1 | Closed |

</details>

<details>
<summary><h2>tektoncd-pruner</h2> — 0 commits</summary>

**Upstream:** tektoncd/pruner · **Branch:** release-v0.4.x → release-v1.23.x · No changes in this period.

</details>

<details>
<summary><h2>tektoncd-triggers</h2> — 0 commits</summary>

**Upstream:** tektoncd/triggers · **Branch:** release-v0.36.x → release-v1.23.x · No changes in this period.

</details>

<details>
<summary><h2>syncer-service</h2> — 0 commits</summary>

**Upstream:** openshift-pipelines/syncer-service · **Branch:** release-v0.1.x → release-v1.23.x · No changes in this period.

</details>

<details>
<summary><h2>manual-approval-gate</h2> — 2 commits, 2 linked, 0 dependabot</summary>

**Upstream:** openshift-pipelines/manual-approval-gate · **Branch:** release-v0.9.0 → release-v1.23.x · **Commits:** 2 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| [bcd1d4e2b076](https://github.com/openshift-pipelines/manual-approval-gate/commit/bcd1d4e2b076bd8329ffd472cd5971978791d8f6) | fix(cve): update golang.org/x/text to v0.39.0 — CVE-2026-56852 | diagrawa | [#1035](https://github.com/openshift-pipelines/manual-approval-gate/pull/1035) | [SRVKP-12834](https://redhat.atlassian.net/browse/SRVKP-12834) *(PR body)* | **Type:** CVE<br>**Text:** Security fix: upgrade golang.org/x/text v0.35.0 → v0.39.0 to address CVE-2026-56852 (infinite loop on invalid unicode input) on release-v0.9.0 *(extracted)* |
| [7ed9c6cfc181](https://github.com/openshift-pipelines/manual-approval-gate/commit/7ed9c6cfc181094438d7dbc3b655cc41d454c86a) | fix(cve): upgrade Go stdlib go1.25.8 → go1.25.12, x/net v0.52.0 → v0.56.0, otel v1.42.0 → v1.43.0 | diagrawa | [#1030](https://github.com/openshift-pipelines/manual-approval-gate/pull/1030) | [SRVKP-12834](https://redhat.atlassian.net/browse/SRVKP-12834) *(keyword)* | **Type:** CVE<br>**Text:** Security fix: upgrade Go stdlib go1.25.8 → go1.25.12, golang.org/x/net v0.52.0 → v0.56.0, otel v1.42.0 → v1.43.0 to address 14 CVEs on release-v0.9.0 *(extracted)* |

</details>

<details>
<summary><h2>tektoncd-pipeline</h2> — 4 commits, 0 linked, 4 dependabot</summary>

**Upstream:** tektoncd/pipeline · **Branch:** release-v1.12.x → release-v1.23.x · **Commits:** 4 · **Unsynced:** 0 · **Dependabot:** 4

All 4 commits are from `dependabot[bot]` and are not shown in the table.

### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-7717](https://redhat.atlassian.net/browse/SRVKP-7717) | Termination messages are truncated and become unparseable causing missing results | Dev Complete |

</details>

<details>
<summary><h2>multicluster-proxy-aae</h2> — 0 commits</summary>

**Upstream:** openshift-pipelines/multicluster-proxy-aae · **Branch:** release-v0.1.x → release-v1.23.x · No changes in this period.

</details>

<details>
<summary><h2>git-init</h2> — Error</summary>

> **Error:** failed to get head SHA at from-date: failed to list head file commits for openshift-pipelines/git-init on branch release-v1.23.x: 404 Not Found

</details>

<details>
<summary><h2>tekton-caches</h2> — 0 commits</summary>

**Upstream:** openshift-pipelines/tekton-caches · **Branch:** release-v0.3.x → release-v1.23.x · No changes in this period.

</details>

<details>
<summary><h2>tekton-kueue</h2> — 0 commits</summary>

**Upstream:** konflux-ci/tekton-kueue · **Branch:** release-v0.3.x → release-v1.23.x · No changes in this period.

</details>

<details>
<summary><h2>tektoncd-chains</h2> — 1 commit, 0 linked, 1 dependabot</summary>

**Upstream:** tektoncd/chains · **Branch:** release-v0.27.x → release-v1.23.x · **Commits:** 1 · **Unsynced:** 0 · **Dependabot:** 1

All 1 commit is from `dependabot[bot]` and is not shown in the table.

</details>

<details>
<summary><h2>tektoncd-cli</h2> — 4 commits, 0 linked, 2 dependabot</summary>

**Upstream:** tektoncd/cli · **Branch:** release-v0.45.x → release-v1.23.x · **Commits:** 4 · **Unsynced:** 0 · **Dependabot:** 2

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| [492f193a1361](https://github.com/tektoncd/cli/commit/492f193a136170a29c6d2a78dbb7466102bdc035) | fix(cve): fix few critical CVEs | Shiv Verma | [#3063](https://github.com/tektoncd/cli/pull/3063) | — | — |
| [a8f1285514d5](https://github.com/tektoncd/cli/commit/a8f1285514d5cdd639b5c6a73e84f1ecfc06504d) | chore(deps): remove unused tektoncd/chains dependency | Arvind | [#3061](https://github.com/tektoncd/cli/pull/3061) | — | — |

</details>

<details>
<summary><h2>pipelines-as-code</h2> — 2 commits, 0 linked, 0 dependabot</summary>

**Upstream:** tektoncd/pipelines-as-code · **Branch:** release-v0.48.x → release-v1.23.x · **Commits:** 2 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| [97c68eb92803](https://github.com/tektoncd/pipelines-as-code/commit/97c68eb92803b37eebcfd054685baa85f9a727ab) | fix(ci): linting issue after new release | Zaki Shaikh | [#2857](https://github.com/tektoncd/pipelines-as-code/pull/2857) | — | — |
| [fa5b3da72b3c](https://github.com/tektoncd/pipelines-as-code/commit/fa5b3da72b3c4177139ef4044b70ca03151b24eb) | dep(go): bump go version 1.26.4 | Zaki Shaikh | [#2857](https://github.com/tektoncd/pipelines-as-code/pull/2857) | — | — |

</details>

<details>
<summary><h2>tektoncd-results</h2> — 1 commit, 0 linked, 0 dependabot</summary>

**Upstream:** tektoncd/results · **Branch:** release-v0.19.x → release-v1.23.x · **Commits:** 1 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| [0810beb3806d](https://github.com/tektoncd/results/commit/0810beb3806dead78cf43db924510210f8e8ddce) | fix(security): update Go stdlib to 1.25.12 and golang.org/x/text to v0.39.0 | Divyanshu Agrawal | [#1393](https://github.com/tektoncd/results/pull/1393) | — | **Type:** CVE<br>**Text:** Security: Update Go stdlib to 1.25.12 and golang.org/x/text to v0.39.0 to address multiple CVEs *(extracted)* |

</details>

---

## Unmatched Jiras

Tickets with no matching release component (or no component label):

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-12182](https://redhat.atlassian.net/browse/SRVKP-12182) | In pipelines page, sort with task status is not working as expected | Tasking and Estimation |

---

## Summary

| Component | Commits | Linked (tool) | Linked (audit) | Unlinked | Has RN | Generated RN | Dependabot |
|-----------|---------|---------------|----------------|----------|--------|--------------|------------|
| operator | 5 | 0 | 0 | 4 | 0 | 0 | 1 |
| opc | 3 | 0 | 1 | 0 | 1 | 1 | 2 |
| console-plugin-pf5 | 1 | 0 | 1 | 0 | 1 | 1 | 0 |
| tekton-assist | Error | — | — | — | — | — | — |
| tektoncd-hub | 4 | 0 | 4 | 0 | 4 | 4 | 0 |
| tektoncd-pruner | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| tektoncd-triggers | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| syncer-service | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| manual-approval-gate | 2 | 0 | 2 | 0 | 2 | 0 | 0 |
| tektoncd-pipeline | 4 | 0 | 0 | 0 | 0 | 0 | 4 |
| multicluster-proxy-aae | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| git-init | Error | — | — | — | — | — | — |
| tekton-caches | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| tekton-kueue | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| tektoncd-chains | 1 | 0 | 0 | 0 | 0 | 0 | 1 |
| tektoncd-cli | 4 | 0 | 0 | 2 | 0 | 0 | 2 |
| pipelines-as-code | 2 | 0 | 0 | 2 | 0 | 0 | 0 |
| tektoncd-results | 1 | 0 | 0 | 1 | 1 | 0 | 0 |
| **Total** | **27** | **0** | **8** | **9** | **9** | **6** | **10** |
