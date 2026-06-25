# Release Audit: 1.22.4

**Source:** `release_1.22.4.json`<br>
**Total commits:** 18 across 7 components<br>
**Already linked (by Go tool):** 5 · **Newly linked:** 1 · **Unlinked:** 3 · **Dependabot (skipped):** 9<br>

---

## tektoncd-chains

**Upstream:** tektoncd/chains · **Commits:** 1

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| ac5a570e | chore(deps): bump the all group with 14 updates (#1728) | dependabot[bot] | [#1725](https://github.com/tektoncd/chains/pull/1725) | — | skipped |

---

## tektoncd-triggers

**Upstream:** tektoncd/triggers · **Commits:** 2

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 0d62ad94 | fix: Skip GitHub API-dependent examples in e2e tests | Shiv Verma | [#2047](https://github.com/tektoncd/triggers/pull/2047) | [SRVKP-12354](https://redhat.atlassian.net/browse/SRVKP-12354) | **Type:** CVE<br>**Text:** Fixes critical and important CVEs (CVE-2026-34986, CVE-2026-29181, CVE-2026-39883, CVE-2026-39821, CVE-2026-27136, CVE-2026-25681, CVE-2026-42502, CVE-2026-46595, CVE-2026-42508) by updating vulnerable Go dependencies. *(generated)* |
| e00d998b | bump(deps): bump tektoncd/pipelione and google.golang.org/grpc to fix CVEs | Shiv Verma | [#2051](https://github.com/tektoncd/triggers/pull/2051) | [SRVKP-12354](https://redhat.atlassian.net/browse/SRVKP-12354) | **Type:** CVE<br>**Text:** Bumps tektoncd/pipeline and google.golang.org/grpc dependencies to fix CVEs. *(generated)* |

---

## console-plugin

**Upstream:** openshift-pipelines/console-plugin · **Commits:** 2

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| f9699c58 | feat: add CVE fix skill with dependency analysis script | Ankur Sinha | [#1136](https://github.com/openshift-pipelines/console-plugin/pull/1136) → [#1140](https://github.com/openshift-pipelines/console-plugin/pull/1140) | [SRVKP-12265](https://redhat.atlassian.net/browse/SRVKP-12265) | **Type:** Enhancement<br>**Text:** Adds a CVE remediation skill and dependency analysis script to automate vulnerability fixes across console-plugin release branches, with Jira integration and FedRAMP image triage. *(generated)* |
| 22161b27 | Merge pull request #1140 from openshift-cherrypick-robot/... | openshift-merge-bot[bot] | [#1140](https://github.com/openshift-pipelines/console-plugin/pull/1140) | [SRVKP-12265](https://redhat.atlassian.net/browse/SRVKP-12265) *(PR title)* | *(merge commit — same PR as above)* |

---

## tektoncd-pipeline

**Upstream:** tektoncd/pipeline · **Commits:** 1

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 7c4eb3ca | fix: bump Go and x deps for CVEs | Vibhav Bobade | [#10337](https://github.com/tektoncd/pipeline/pull/10337) | [SRVKP-12360](https://redhat.atlassian.net/browse/SRVKP-12360) | **Type:** CVE<br>**Text:** Bumps Go to 1.25.10, golang.org/x/crypto to v0.52.0, and golang.org/x/net to v0.55.0 for CVE remediation. *(extracted)* |

---

## tektoncd-results

**Upstream:** tektoncd/results · **Commits:** 1

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| ad9dbfdc | Merge pull request #1335 from divyansh42/cherrypick-11380 | Emil Natan | [#1335](https://github.com/tektoncd/results/pull/1335) | — | **Type:** Feature<br>**Text:** Improved performance for `describe` and `logs` commands with optimized filtering; when multiple runs share the same name, the most recent one is returned (use `--uid` to target a specific run). *(extracted)* |

---

## operator

**Upstream:** tektoncd/operator · **Commits:** 10

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 88fce7e9 | chore(deps): update openshift images sha | Shiv Verma | [#3514](https://github.com/tektoncd/operator/pull/3514) | — | **Type:** Release Notes Not Required<br>**Text:** *(none)* |
| 856d60c3 | Merge multikueueclusters permission to scheduler role | Pramod Bindal | [#3518](https://github.com/tektoncd/operator/pull/3518) → [#3522](https://github.com/tektoncd/operator/pull/3522) | [SRVKP-12211](https://redhat.atlassian.net/browse/SRVKP-12211) | **Type:** Bug Fix<br>**Text:** tekton-multicluster-proxy-aae-role and tekton-multicluster-proxy-aae-rolebinding are removed. Required permissions are moved to tekton-scheduler-role role. On OpenShift, tekton-scheduler-role-binding changes the subject to openshift-operator SA so there is no orphaned role or rolebinding present. *(Jira)* |
| 1ce87aad | chore(deps): bump chainguard-dev/actions from 1.6.5 to 1.6.23 | dependabot[bot] | [#3567](https://github.com/tektoncd/operator/pull/3567) | — | skipped |
| 164a8d5b | chore(deps): bump actions/checkout from 6.0.2 to 6.0.3 | dependabot[bot] | [#3570](https://github.com/tektoncd/operator/pull/3570) | — | skipped |
| 3c4c4bfd | chore(deps): bump golangci/golangci-lint-action from 9.2.0 to 9.2.1 | dependabot[bot] | [#3568](https://github.com/tektoncd/operator/pull/3568) | — | skipped |
| 9f88e772 | chore(deps): bump actions/cache from 5.0.3 to 5.0.5 | dependabot[bot] | [#3569](https://github.com/tektoncd/operator/pull/3569) | — | skipped |
| 4d299695 | chore(deps): bump peter-evans/create-pull-request from 8.1.0 to 8.1.1 | dependabot[bot] | [#3566](https://github.com/tektoncd/operator/pull/3566) | — | skipped |
| ead77a65 | chore(deps): bump github.com/sigstore/cosign/v2 from 2.6.2 to 2.6.3 | dependabot[bot] | [#3576](https://github.com/tektoncd/operator/pull/3576) | — | skipped |
| 1dbf76a6 | chore(deps): bump github.com/tektoncd/pipeline from 1.9.3 to 1.9.4 | dependabot[bot] | [#3575](https://github.com/tektoncd/operator/pull/3575) | — | skipped |
| 184d0d67 | chore(deps): bump github.com/cert-manager/cert-manager | dependabot[bot] | [#3577](https://github.com/tektoncd/operator/pull/3577) | — | skipped |

---

## opc

**Upstream:** openshift-pipelines/opc · **Commits:** 1

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 5fabea4c | bump tektoncd/cli to v0.44.2 | Shiv Verma | [#468](https://github.com/openshift-pipelines/opc/pull/468) | — | **Type:** Enhancement<br>**Text:** Updates the embedded Tekton CLI to v0.44.2, bringing in the latest upstream bug fixes and improvements. *(generated)* |

---

## Components with no changes

tekton-caches, tektoncd-cli, multicluster-proxy-aae, tektoncd-hub, tektoncd-pruner, syncer-service, manual-approval-gate, console-plugin-pf5, pipelines-as-code, tekton-kueue.

Components with errors (repo not found): git-init, tekton-assist.

---

## Unmatched Jiras

8 tickets in fixVersion "Pipeline 1.22.4" not linked to any commit. All are GovCloud image CVE remediation tickets — container image rebuilds, not source code changes.

| Key | Summary | Status |
|-----|---------|--------|
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
| tektoncd-triggers | 2 | 2 | 0 | 0 | 0 | 2 | 0 |
| console-plugin | 2 | 1 | 1 | 0 | 0 | 1 | 0 |
| tektoncd-pipeline | 1 | 1 | 0 | 0 | 0 | 1 | 0 |
| tektoncd-results | 1 | 0 | 0 | 1 | 0 | 1 | 0 |
| operator | 10 | 1 | 0 | 1 | 1 | 0 | 8 |
| opc | 1 | 0 | 0 | 1 | 0 | 1 | 0 |
| **Total** | **18** | **5** | **1** | **3** | **1** | **6** | **9** |
