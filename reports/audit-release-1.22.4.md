# Release Audit: 1.22.4

**Source:** `reports/release_1.22.4.json`<br>
**Total commits:** 18 across 7 components<br>
**Already linked (by Go tool):** 5 · **Newly linked:** 0 · **Unlinked:** 4 · **Dependabot (skipped):** 9<br>

## Index

| Component | Commits | Jira | Unsynced | Unmatched Jiras |
|-----------|---------|------|----------|-----------------|
| [tektoncd-chains](#tektoncd-chains) | 1 | 0/1 | 0 | 0 |
| [tektoncd-triggers](#tektoncd-triggers) | 2 | 2/2 | 0 | 0 |
| [console-plugin](#console-plugin) | 2 | 1/2 | 0 | 0 |
| [tektoncd-pipeline](#tektoncd-pipeline) | 1 | 1/1 | 0 | 0 |
| [tektoncd-results](#tektoncd-results) | 1 | 0/1 | 0 | 0 |
| [operator](#operator) | 10 | 1/10 | 0 | 0 |
| [opc](#opc) | 1 | 0/1 | 0 | 0 |

---

<details>
<summary><h2>tektoncd-chains</h2> — 1 commits, 0 linked, 1 dependabot</summary>

**Upstream:** tektoncd/chains · **Commits:** 1 · **Unsynced:** 0 · **Dependabot:** 1


</details>

---

<details>
<summary><h2>tektoncd-triggers</h2> — 2 commits, 2 linked</summary>

**Upstream:** tektoncd/triggers · **Commits:** 2 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 0d62ad94 | fix: Skip GitHub API-dependent examples in e2e tests | Shiv Verma | [#2047](https://github.com/tektoncd/triggers/pull/2047) | [SRVKP-12354](https://redhat.atlassian.net/browse/SRVKP-12354) | — |
| e00d998b | bump(deps): bump tektoncd/pipelione and google.golang.org/grpc to fix CVEs | Shiv Verma | [#2051](https://github.com/tektoncd/triggers/pull/2051) | [SRVKP-12354](https://redhat.atlassian.net/browse/SRVKP-12354) | — |

</details>

---

<details>
<summary><h2>console-plugin</h2> — 2 commits, 1 linked</summary>

**Upstream:** openshift-pipelines/console-plugin · **Commits:** 2 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| f9699c58 | feat: add CVE fix skill with dependency analysis script | Ankur Sinha | [#1136](https://github.com/openshift-pipelines/console-plugin/pull/1136) → [#1140](https://github.com/openshift-pipelines/console-plugin/pull/1140) | [SRVKP-12265](https://redhat.atlassian.net/browse/SRVKP-12265) | — |
| 22161b27 | Merge pull request #1140 from openshift-cherrypick-robot/cherry-pick-1136-to-rel | openshift-merge-bot[bot] | [#1140](https://github.com/openshift-pipelines/console-plugin/pull/1140) | — | — |

</details>

---

<details>
<summary><h2>tektoncd-pipeline</h2> — 1 commits, 1 linked</summary>

**Upstream:** tektoncd/pipeline · **Commits:** 1 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 7c4eb3ca | fix: bump Go and x deps for CVEs | Vibhav Bobade | [#10337](https://github.com/tektoncd/pipeline/pull/10337) | [SRVKP-12360](https://redhat.atlassian.net/browse/SRVKP-12360) | — |

</details>

---

<details>
<summary><h2>tektoncd-results</h2> — 1 commits, 0 linked</summary>

**Upstream:** tektoncd/results · **Commits:** 1 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| ad9dbfdc | Merge pull request #1335 from divyansh42/cherrypick-11380 | Emil Natan | [#1335](https://github.com/tektoncd/results/pull/1335) | — | — |

</details>

---

<details>
<summary><h2>operator</h2> — 10 commits, 1 linked, 8 dependabot</summary>

**Upstream:** tektoncd/operator · **Commits:** 10 · **Unsynced:** 0 · **Dependabot:** 8

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 88fce7e9 | chore(deps): update openshift images sha | Shiv Verma | [#3514](https://github.com/tektoncd/operator/pull/3514) | — | — |
| 856d60c3 | Merge multikueueclusters permission to scheduler role | Pramod Bindal | [#3518](https://github.com/tektoncd/operator/pull/3518) → [#3522](https://github.com/tektoncd/operator/pull/3522) | [SRVKP-12211](https://redhat.atlassian.net/browse/SRVKP-12211) | **Bug Fix:** tekton-multicluster-proxy-aae-role and tekton-multicluster-proxy-aae-rolebinding are removed. Required permissions are moved to tekton-scheduler-ro... *(Jira)* |

</details>

---

<details>
<summary><h2>opc</h2> — 1 commits, 0 linked</summary>

**Upstream:** openshift-pipelines/opc · **Commits:** 1 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 5fabea4c | bump tektoncd/cli to v0.44.2 | Shiv Verma | [#468](https://github.com/openshift-pipelines/opc/pull/468) | — | — |

</details>

---

## Components with no changes

tekton-caches, tektoncd-cli, multicluster-proxy-aae, tektoncd-hub, tektoncd-pruner, syncer-service, manual-approval-gate, console-plugin-pf5, pipelines-as-code, tekton-kueue.

Components with errors (repo not found): git-init, tekton-assist.

---

## Unmatched Jiras

Tickets with no Jira component label or no matching release component.

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-12584](https://redhat.atlassian.net/browse/SRVKP-12584) | CLI testing | New |
| [SRVKP-12583](https://redhat.atlassian.net/browse/SRVKP-12583) | UI Regression testing | New |
| [SRVKP-12582](https://redhat.atlassian.net/browse/SRVKP-12582) | Acceptance test | In Progress |
| [SRVKP-12581](https://redhat.atlassian.net/browse/SRVKP-12581) | Upgrade test | New |
| [SRVKP-12580](https://redhat.atlassian.net/browse/SRVKP-12580) | Test Openshift Pipelines 1.22.4 | In Progress |
| [SRVKP-12361](https://redhat.atlassian.net/browse/SRVKP-12361) | [GovCloud] pipelines-serve-tkn-cli-rhel9 - 9 CVEs (CRITICAL) | Closed |
| [SRVKP-12359](https://redhat.atlassian.net/browse/SRVKP-12359) | [GovCloud] pipelines-events-rhel9 - 11 CVEs (CRITICAL) | Closed |
| [SRVKP-12353](https://redhat.atlassian.net/browse/SRVKP-12353) | [GovCloud] pipelines-triggers-controller-rhel9 - 14 CVEs (CRITICAL) | Closed |
| [SRVKP-12352](https://redhat.atlassian.net/browse/SRVKP-12352) | [GovCloud] pipelines-controller-rhel9 - 22 CVEs (CRITICAL) | Closed |
| [SRVKP-12350](https://redhat.atlassian.net/browse/SRVKP-12350) | [GovCloud] pipelines-resolvers-rhel9 - 22 CVEs (CRITICAL) | Closed |
| [SRVKP-12348](https://redhat.atlassian.net/browse/SRVKP-12348) | [GovCloud] pipelines-triggers-core-interceptors-rhel9 - 24 CVEs (CRITICAL) | Closed |
| [SRVKP-12343](https://redhat.atlassian.net/browse/SRVKP-12343) | [GovCloud] pipelines-cli-tkn-rhel9 - 28 CVEs (CRITICAL) | Closed |
| [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) | [GovCloud] June 2026 - FedRAMP CVE Remediation - OpenShift Pipelines | Dev Complete |

---

## Summary

| Component | Commits | Linked (tool) | Linked (audit) | Unlinked | Has RN | Generated RN | Dependabot |
|-----------|---------|---------------|----------------|----------|--------|--------------|------------|
| tektoncd-chains | 1 | 0 | 0 | 0 | 0 | 0 | 1 |
| tektoncd-triggers | 2 | 2 | 0 | 0 | 0 | 0 | 0 |
| console-plugin | 2 | 1 | 0 | 1 | 0 | 0 | 0 |
| tektoncd-pipeline | 1 | 1 | 0 | 0 | 0 | 0 | 0 |
| tektoncd-results | 1 | 0 | 0 | 1 | 0 | 0 | 0 |
| operator | 10 | 1 | 0 | 1 | 1 | 0 | 8 |
| opc | 1 | 0 | 0 | 1 | 0 | 0 | 0 |
| **Total** | **18** | **5** | **0** | **4** | **1** | **0** | **9** |
