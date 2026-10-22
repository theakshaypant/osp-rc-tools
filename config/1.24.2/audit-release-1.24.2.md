# Release Audit: 1.24.2

**Snapshot:** openshift-pipelines-core-1-24-20261019-170342-000
**Previous:** 1.24.1
**Generated:** 2026-10-22 04:24 UTC

---

## Summary

**Total Components:** 17
**Processed:** 17
**Successful:** 17
**Errors:** 0

---

## Components

### operator

**Upstream:** tektoncd/operator
**Commits:** 18 (9 shown, 9 dependabot hidden)
**Range:** [`29ad15322f5e`](https://github.com/tektoncd/operator/commit/29ad15322f5e07246030c63c058c806a3daf9bb3) → [`6fe35f18b138`](https://github.com/tektoncd/operator/commit/6fe35f18b138f844eb5a98f84caa8ff13fe9813c)

| SHA | Message | Author | Date | PR | Original PR |
|-----|---------|--------|------|----|----|
| [f33318e635bc](https://github.com/tektoncd/operator/commit/f33318e635bc2c691eedb968d1bfb555d80e2ff2) | chore: bump chains from v0.29.5 to v0.29.6 | Vincent Demeester | 2026-09-25 | [4160](https://github.com/tektoncd/operator/pull/4160) | — |
| [1c085a78be40](https://github.com/tektoncd/operator/commit/1c085a78be40e6330e13cee0b254c6d738f164e0) | chore: update third-party image digests | Divyanshu Agrawal | 2026-10-01 | [4187](https://github.com/tektoncd/operator/pull/4187) | — |
| [98df9d34174b](https://github.com/tektoncd/operator/commit/98df9d34174b821f47dd860290517e1345015d5f) | chore: bump chains from v0.29.6 to v0.29.7 | Vincent Demeester | 2026-10-02 | [4193](https://github.com/tektoncd/operator/pull/4193) | — |
| [e037472d9adc](https://github.com/tektoncd/operator/commit/e037472d9adc6541fbe4950a010f951265476ff4) | fix(bundle): drop duplicate MulticlusterProxyAAE CRD from OpenShift bundle | Jawed khelil | 2026-10-08 | [4219](https://github.com/tektoncd/operator/pull/4219) | — |
| [7b37043c4515](https://github.com/tektoncd/operator/commit/7b37043c451554074a07ed5383cbb0dd073d9494) | fix(openshift): block SCC escalation via annotations | pratap0007 | 2026-10-09 | [4221](https://github.com/tektoncd/operator/pull/4221) | — |
| [1eb452c99309](https://github.com/tektoncd/operator/commit/1eb452c9930960caf4adc891559e6f611729afc3) | feat(pruner): add default resource limits to controller and webhook | Shubham Bhardwaj | 2026-10-14 | [4256](https://github.com/tektoncd/operator/pull/4256) | — |
| [7f865db42acd](https://github.com/tektoncd/operator/commit/7f865db42acd542aa70ff77ef90c7eab4b15b838) | chore: update third-party image digests | Divyanshu Agrawal | 2026-10-14 | [4271](https://github.com/tektoncd/operator/pull/4271) | — |
| [bd476b3d1909](https://github.com/tektoncd/operator/commit/bd476b3d1909e43a0f2cb0731fa082e6f3658d4d) | fix(chains): allow slsa/v2alpha4 attestation format in validation | Jawed khelil | 2026-10-14 | [4259](https://github.com/tektoncd/operator/pull/4259) | — |
| [927ec44751f2](https://github.com/tektoncd/operator/commit/927ec44751f2fe5a0874f7a0b382e6bd29291b01) | test(chains): drive format validation from allowlist vars | Jawed khelil | 2026-10-14 | [4259](https://github.com/tektoncd/operator/pull/4259) | — |

---

### p12n-console-plugin

**Upstream:** openshift-pipelines/console-plugin
**Commits:** 20 (20 shown, 0 dependabot hidden)
**Range:** [`23678baa99d2`](https://github.com/openshift-pipelines/console-plugin/commit/23678baa99d239f0ac30b4a2f07707c33c3c0a66) → [`1837c380c775`](https://github.com/openshift-pipelines/console-plugin/commit/1837c380c775f53dd765a92efff27bd7905b431d)

| SHA | Message | Author | Date | PR | Original PR | Jira | RN Type | RN Status | RN Text |
|-----|---------|--------|------|----|----|------|---------|-----------|---------|
| [33c9ee133d43](https://github.com/openshift-pipelines/console-plugin/commit/33c9ee133d430442117776fb41419a62701b507a) | fix: prevent Pipelines Overview page duration sorting error | Ankur Sinha | 2026-09-24 | [1380](https://github.com/openshift-pipelines/console-plugin/pull/1380) | — | [SRVKP-13905](https://redhat.atlassian.net/browse/SRVKP-13905) | Release Note Not Required | — | — |
| [17590f314d62](https://github.com/openshift-pipelines/console-plugin/commit/17590f314d627202cc56590caaed2074f2436d4b) | disable links for resolver pipelines in Overview page | Arvind | 2026-09-24 | [1381](https://github.com/openshift-pipelines/console-plugin/pull/1381) | — | [SRVKP-11221](https://redhat.atlassian.net/browse/SRVKP-11221) | Bug Fix | Proposed | The Pipeline Overview page lists resolver based PipelineRuns without a hyperlink as there is no Pipeline definition associated with it. Also a tooltip is added which explains that the pipeline definition does not exist. |
| [ed96101bb77f](https://github.com/openshift-pipelines/console-plugin/commit/ed96101bb77fc735d449df38b5ba387a5d81de02) | feat: added check and filter for ns deletion in pipelines list ( overview page ) assistedBy: cursor | Anwesha Palit | 2026-09-24 | [1383](https://github.com/openshift-pipelines/console-plugin/pull/1383) | — | [SRVKP-12176](https://redhat.atlassian.net/browse/SRVKP-12176) | Bug Fix | Proposed | The Pipeline List on the Overview page now handles deleted namespaces more clearly. Pipelines associated with deleted namespaces remain included in the aggregated metrics to keep counts consistent across the dashboard. To help identify pipelines from deleted namespaces, we introduced two enhancements: Links to the pipeline, PipelineRun count, and namespace are disabled. A tooltip indicates that the namespace has been deleted. |
| [2597bd1c708e](https://github.com/openshift-pipelines/console-plugin/commit/2597bd1c708ed52c5fd79922d1b22b2ec31dec8f) | chore(i18n): update Memsource translations | openshift-pipelines-bot | 2026-09-24 | [1385](https://github.com/openshift-pipelines/console-plugin/pull/1385) | — | — | — | — | — |
| [0f60b2eafff0](https://github.com/openshift-pipelines/console-plugin/commit/0f60b2eafff04085bcb66c220f0d26f8ced50669) | Merge pull request #1383 from openshift-cherrypick-robot/cherry-pick-1377-to-release-v1.24.x | openshift-merge-bot[bot] | 2026-09-24 | [1383](https://github.com/openshift-pipelines/console-plugin/pull/1383) | — | [SRVKP-12176](https://redhat.atlassian.net/browse/SRVKP-12176) | Bug Fix | Proposed | The Pipeline List on the Overview page now handles deleted namespaces more clearly. Pipelines associated with deleted namespaces remain included in the aggregated metrics to keep counts consistent across the dashboard. To help identify pipelines from deleted namespaces, we introduced two enhancements: Links to the pipeline, PipelineRun count, and namespace are disabled. A tooltip indicates that the namespace has been deleted. |
| [4e0a73222661](https://github.com/openshift-pipelines/console-plugin/commit/4e0a73222661087f1081d4247f22e12479d07d94) | Merge pull request #1385 from openshift-cherrypick-robot/cherry-pick-1384-to-release-v1.24.x | openshift-merge-bot[bot] | 2026-09-24 | [1385](https://github.com/openshift-pipelines/console-plugin/pull/1385) | — | — | — | — | — |
| [e190fffb83ab](https://github.com/openshift-pipelines/console-plugin/commit/e190fffb83ab93f6bed649347c3c47c3cc894c84) | chore: regenerate yarn.lock | openshift-pipelines-bot | 2026-09-27 | [1389](https://github.com/openshift-pipelines/console-plugin/pull/1389) | — | — | — | — | — |
| [c8968e4569be](https://github.com/openshift-pipelines/console-plugin/commit/c8968e4569be421a83d512caac179e67a1bac9d2) | Merge pull request #1389 from openshift-pipelines/automation/refresh-yarn-lockfile-release-v1.24.x | openshift-merge-bot[bot] | 2026-09-27 | [1389](https://github.com/openshift-pipelines/console-plugin/pull/1389) | — | — | — | — | — |
| [4845c2b42e2b](https://github.com/openshift-pipelines/console-plugin/commit/4845c2b42e2b936159c9ee4f919f4003a03e5b87) | chore: regenerate yarn.lock | openshift-pipelines-bot | 2026-10-04 | [1405](https://github.com/openshift-pipelines/console-plugin/pull/1405) | — | — | — | — | — |
| [532c343763b3](https://github.com/openshift-pipelines/console-plugin/commit/532c343763b326c8f88f38c89aa0135df8d26898) | Merge pull request #1405 from openshift-pipelines/automation/refresh-yarn-lockfile-release-v1.24.x | openshift-merge-bot[bot] | 2026-10-04 | [1405](https://github.com/openshift-pipelines/console-plugin/pull/1405) | — | — | — | — | — |
| [859a87377469](https://github.com/openshift-pipelines/console-plugin/commit/859a8737746990186ff052a4b9b59dc58b655cfe) | chore: update install script for MAG | Anwesha Palit | 2026-10-05 | [1412](https://github.com/openshift-pipelines/console-plugin/pull/1412) | — | — | — | — | — |
| [eb0d12fa6755](https://github.com/openshift-pipelines/console-plugin/commit/eb0d12fa6755de863c3fd00b6d21f109a3ad4c62) | fix: add null check for etcdRuns in useRuns | Anwesha Palit | 2026-10-05 | [1415](https://github.com/openshift-pipelines/console-plugin/pull/1415) | — | [SRVKP-14864](https://redhat.atlassian.net/browse/SRVKP-14864) | Bug Fix | Proposed | The issue is not resolved and the pipeline run list and details page works as expected on cancelling / stopping a pipelinerun |
| [3889e419d682](https://github.com/openshift-pipelines/console-plugin/commit/3889e419d6824c3396606696570e3712422c4447) | Merge pull request #1415 from openshift-cherrypick-robot/cherry-pick-1413-to-release-v1.24.x | openshift-merge-bot[bot] | 2026-10-05 | [1415](https://github.com/openshift-pipelines/console-plugin/pull/1415) | — | [SRVKP-14864](https://redhat.atlassian.net/browse/SRVKP-14864) | Bug Fix | Proposed | The issue is not resolved and the pipeline run list and details page works as expected on cancelling / stopping a pipelinerun |
| [bde2a7418b32](https://github.com/openshift-pipelines/console-plugin/commit/bde2a7418b32f8f69b16765910471f7d47dd3021) | fix: refactor applyParentUpgrade, correct applyResolutions orderfix and use exact-version resolution entries and discover historical descriptors Co-Authored-By: Claude Opus 4.6 | Anwesha Palit | 2026-10-05 | [1416](https://github.com/openshift-pipelines/console-plugin/pull/1416) | — | — | — | — | — |
| [42de4ae09004](https://github.com/openshift-pipelines/console-plugin/commit/42de4ae09004e50da603581368385658c3b91111) | fix: show restricted access instead of loading on PipelineRuns RBAC error | Ankur Sinha | 2026-10-06 | [1423](https://github.com/openshift-pipelines/console-plugin/pull/1423) | — | [SRVKP-14135](https://redhat.atlassian.net/browse/SRVKP-14135) | Release Note Not Required | — | — |
| [f85d854cbd31](https://github.com/openshift-pipelines/console-plugin/commit/f85d854cbd31f0e8e012e79a2884ce2006f13915) | Merge pull request #1423 from openshift-cherrypick-robot/cherry-pick-1407-to-release-v1.24.x | openshift-merge-bot[bot] | 2026-10-08 | [1423](https://github.com/openshift-pipelines/console-plugin/pull/1423) | — | [SRVKP-14135](https://redhat.atlassian.net/browse/SRVKP-14135) | Release Note Not Required | — | — |
| [c81bc7feacb5](https://github.com/openshift-pipelines/console-plugin/commit/c81bc7feacb562a6668e31a9c715f856c02bfe32) | chore: regenerate yarn.lock | openshift-pipelines-bot | 2026-10-11 | [1425](https://github.com/openshift-pipelines/console-plugin/pull/1425) | — | — | — | — | — |
| [7d86e78d008b](https://github.com/openshift-pipelines/console-plugin/commit/7d86e78d008bdb5c51aa6f5290edfa4db7782016) | Merge pull request #1425 from openshift-pipelines/automation/refresh-yarn-lockfile-release-v1.24.x | openshift-merge-bot[bot] | 2026-10-11 | [1425](https://github.com/openshift-pipelines/console-plugin/pull/1425) | — | — | — | — | — |
| [ab89206641a3](https://github.com/openshift-pipelines/console-plugin/commit/ab89206641a3ec9ba6e195a7449cbe52b5b50db4) | chore: regenerate yarn.lock | openshift-pipelines-bot | 2026-10-18 | [1438](https://github.com/openshift-pipelines/console-plugin/pull/1438) | — | — | — | — | — |
| [1837c380c775](https://github.com/openshift-pipelines/console-plugin/commit/1837c380c775f53dd765a92efff27bd7905b431d) | Merge pull request #1438 from openshift-pipelines/automation/refresh-yarn-lockfile-release-v1.24.x | openshift-merge-bot[bot] | 2026-10-18 | [1438](https://github.com/openshift-pipelines/console-plugin/pull/1438) | — | — | — | — | — |

---

### p12n-console-plugin-pf5

**Upstream:** openshift-pipelines/console-plugin
**Commits:** 13 (13 shown, 0 dependabot hidden)
**Range:** [`87072cc1800b`](https://github.com/openshift-pipelines/console-plugin/commit/87072cc1800b024ac420fbf8f80984c086374804) → [`8095000a88f8`](https://github.com/openshift-pipelines/console-plugin/commit/8095000a88f8b0094b14c35a888f38c19589b96e)

| SHA | Message | Author | Date | PR | Original PR | Jira | RN Type | RN Status | RN Text |
|-----|---------|--------|------|----|----|------|---------|-----------|---------|
| [ba3fc9e7ef79](https://github.com/openshift-pipelines/console-plugin/commit/ba3fc9e7ef79a56e7a7f9b02413634cbad121537) | chore: regenerate yarn.lock | openshift-pipelines-bot | 2026-09-27 | [1394](https://github.com/openshift-pipelines/console-plugin/pull/1394) | — | [SRVKP-14195](https://redhat.atlassian.net/browse/SRVKP-14195)<br>[SRVKP-14200](https://redhat.atlassian.net/browse/SRVKP-14200)<br>[SRVKP-14243](https://redhat.atlassian.net/browse/SRVKP-14243)<br>[SRVKP-14245](https://redhat.atlassian.net/browse/SRVKP-14245)<br>[SRVKP-14247](https://redhat.atlassian.net/browse/SRVKP-14247)<br>[SRVKP-14340](https://redhat.atlassian.net/browse/SRVKP-14340) | —<br>—<br>—<br>—<br>—<br>Release Note Not Required | —<br>—<br>—<br>—<br>—<br>— | —<br>—<br>—<br>—<br>—<br>— |
| [79504eb65ccd](https://github.com/openshift-pipelines/console-plugin/commit/79504eb65ccd38f091e2e6bc8ba4392b946800d7) | Merge pull request #1394 from openshift-pipelines/automation/refresh-yarn-lockfile-release-v1.22.x | openshift-merge-bot[bot] | 2026-09-27 | [1394](https://github.com/openshift-pipelines/console-plugin/pull/1394) | — | [SRVKP-14195](https://redhat.atlassian.net/browse/SRVKP-14195)<br>[SRVKP-14200](https://redhat.atlassian.net/browse/SRVKP-14200)<br>[SRVKP-14243](https://redhat.atlassian.net/browse/SRVKP-14243)<br>[SRVKP-14245](https://redhat.atlassian.net/browse/SRVKP-14245)<br>[SRVKP-14247](https://redhat.atlassian.net/browse/SRVKP-14247)<br>[SRVKP-14340](https://redhat.atlassian.net/browse/SRVKP-14340) | —<br>—<br>—<br>—<br>—<br>Release Note Not Required | —<br>—<br>—<br>—<br>—<br>— | —<br>—<br>—<br>—<br>—<br>— |
| [6b8878ec19c0](https://github.com/openshift-pipelines/console-plugin/commit/6b8878ec19c0942e73f17a90f7792530659978b0) | feat: remove all cypress related files from main and 1.22.x release branch | Anwesha Palit | 2026-10-01 | [1398](https://github.com/openshift-pipelines/console-plugin/pull/1398) | — | [SRVKP-14414](https://redhat.atlassian.net/browse/SRVKP-14414) | Release Note Not Required | — | — |
| [c38144ece7b2](https://github.com/openshift-pipelines/console-plugin/commit/c38144ece7b22d9c89b4c33ed66d7b89aa3d5112) | Merge pull request #1398 from anwesha-palit-redhat/SRVKP-14414/release-v1.22.x | openshift-merge-bot[bot] | 2026-10-01 | [1398](https://github.com/openshift-pipelines/console-plugin/pull/1398) | — | [SRVKP-14414](https://redhat.atlassian.net/browse/SRVKP-14414) | Release Note Not Required | — | — |
| [d8bd688a7bee](https://github.com/openshift-pipelines/console-plugin/commit/d8bd688a7bee321a1875318fdf052a56d865e61c) | chore: regenerate yarn.lock | openshift-pipelines-bot | 2026-10-04 | [1403](https://github.com/openshift-pipelines/console-plugin/pull/1403) | — | — | — | — | — |
| [2581fae8030b](https://github.com/openshift-pipelines/console-plugin/commit/2581fae8030b61abe6cee86ddbef967af27f70d8) | Merge pull request #1403 from openshift-pipelines/automation/refresh-yarn-lockfile-release-v1.22.x | openshift-merge-bot[bot] | 2026-10-04 | [1403](https://github.com/openshift-pipelines/console-plugin/pull/1403) | — | — | — | — | — |
| [fc12aebb80d0](https://github.com/openshift-pipelines/console-plugin/commit/fc12aebb80d07f76383c4d50c2d2cee44552beee) | fix: refactor applyParentUpgrade, correct applyResolutions orderfix and use exact-version resolution entries and discover historical descriptors Co-Authored-By: Claude Opus 4.6 | Anwesha Palit | 2026-10-05 | [1418](https://github.com/openshift-pipelines/console-plugin/pull/1418) | — | — | — | — | — |
| [a8e55d59966d](https://github.com/openshift-pipelines/console-plugin/commit/a8e55d59966d6621f52939161291ced3d66e8aa5) | chore: fix CVEs | openshift-pipelines-bot | 2026-10-05 | [1421](https://github.com/openshift-pipelines/console-plugin/pull/1421) | — | [SRVKP-14131](https://redhat.atlassian.net/browse/SRVKP-14131)<br>[SRVKP-14212](https://redhat.atlassian.net/browse/SRVKP-14212)<br>[SRVKP-14560](https://redhat.atlassian.net/browse/SRVKP-14560)<br>[SRVKP-14568](https://redhat.atlassian.net/browse/SRVKP-14568) | —<br>—<br>Release Note Not Required<br>Release Note Not Required | —<br>—<br>—<br>— | —<br>—<br>—<br>— |
| [1afce627ae8c](https://github.com/openshift-pipelines/console-plugin/commit/1afce627ae8cb9d964ca5dc0e15387112e0912eb) | Merge pull request #1421 from openshift-pipelines/auto/cve/fix-release-v1.22.x | openshift-merge-bot[bot] | 2026-10-05 | [1421](https://github.com/openshift-pipelines/console-plugin/pull/1421) | — | [SRVKP-14131](https://redhat.atlassian.net/browse/SRVKP-14131)<br>[SRVKP-14212](https://redhat.atlassian.net/browse/SRVKP-14212)<br>[SRVKP-14560](https://redhat.atlassian.net/browse/SRVKP-14560)<br>[SRVKP-14568](https://redhat.atlassian.net/browse/SRVKP-14568) | —<br>—<br>Release Note Not Required<br>Release Note Not Required | —<br>—<br>—<br>— | —<br>—<br>—<br>— |
| [101c8a742d89](https://github.com/openshift-pipelines/console-plugin/commit/101c8a742d89c466a3605d1468f78a0e2ee09a42) | chore: regenerate yarn.lock | openshift-pipelines-bot | 2026-10-11 | [1429](https://github.com/openshift-pipelines/console-plugin/pull/1429) | — | — | — | — | — |
| [0c45d6dc5994](https://github.com/openshift-pipelines/console-plugin/commit/0c45d6dc5994ca66400fc6cd48f5360789faa9da) | Merge pull request #1429 from openshift-pipelines/automation/refresh-yarn-lockfile-release-v1.22.x | openshift-merge-bot[bot] | 2026-10-11 | [1429](https://github.com/openshift-pipelines/console-plugin/pull/1429) | — | — | — | — | — |
| [94a10da74304](https://github.com/openshift-pipelines/console-plugin/commit/94a10da74304b1c2d436a3a49d234ccf03e66447) | chore: regenerate yarn.lock | openshift-pipelines-bot | 2026-10-18 | [1441](https://github.com/openshift-pipelines/console-plugin/pull/1441) | — | — | — | — | — |
| [8095000a88f8](https://github.com/openshift-pipelines/console-plugin/commit/8095000a88f8b0094b14c35a888f38c19589b96e) | Merge pull request #1441 from openshift-pipelines/automation/refresh-yarn-lockfile-release-v1.22.x | openshift-merge-bot[bot] | 2026-10-18 | [1441](https://github.com/openshift-pipelines/console-plugin/pull/1441) | — | — | — | — | — |

---

### p12n-manual-approval-gate

**Upstream:** openshift-pipelines/manual-approval-gate
**Commits:** 0 (0 shown, 0 dependabot hidden)
**Range:** [`2ec8c3ba8921`](https://github.com/openshift-pipelines/manual-approval-gate/commit/2ec8c3ba892105570223a7ddbb03e157b38f4114) → [`2ec8c3ba8921`](https://github.com/openshift-pipelines/manual-approval-gate/commit/2ec8c3ba892105570223a7ddbb03e157b38f4114)


---

### p12n-multicluster-proxy-aae

**Upstream:** openshift-pipelines/multicluster-proxy-aae
**Commits:** 3 (3 shown, 0 dependabot hidden)
**Range:** [`1d53e3aee05f`](https://github.com/openshift-pipelines/multicluster-proxy-aae/commit/1d53e3aee05f7969c25f3a803c74d45027d3e579) → [`3b3f3e39f69d`](https://github.com/openshift-pipelines/multicluster-proxy-aae/commit/3b3f3e39f69dd2b37a7d18323e858732672f51c9)

| SHA | Message | Author | Date | PR | Original PR |
|-----|---------|--------|------|----|----|
| [60279bf58d13](https://github.com/openshift-pipelines/multicluster-proxy-aae/commit/60279bf58d1374eac39182be76ca7ec4f37d137d) | Update permissions for secret | savitaashture | 2026-10-05 | [38](https://github.com/openshift-pipelines/multicluster-proxy-aae/pull/38) | — |
| [ed397a36c6e7](https://github.com/openshift-pipelines/multicluster-proxy-aae/commit/ed397a36c6e733cafe705afee0023f8c5439da0a) | Fix per request Kubernetes client construction | savitaashture | 2026-10-05 | [38](https://github.com/openshift-pipelines/multicluster-proxy-aae/pull/38) | — |
| [3b3f3e39f69d](https://github.com/openshift-pipelines/multicluster-proxy-aae/commit/3b3f3e39f69dd2b37a7d18323e858732672f51c9) | fix import issue | sashture | 2026-10-05 | [39](https://github.com/openshift-pipelines/multicluster-proxy-aae/pull/39) | — |

---

### p12n-opc

**Upstream:** openshift-pipelines/opc
**Commits:** 0 (0 shown, 0 dependabot hidden)
**Range:** [`e65413f6c2db`](https://github.com/openshift-pipelines/opc/commit/e65413f6c2dbae111b31c455e9271b737c32145a) → [`e65413f6c2db`](https://github.com/openshift-pipelines/opc/commit/e65413f6c2dbae111b31c455e9271b737c32145a)


---

### p12n-syncer-service

**Upstream:** openshift-pipelines/syncer-service
**Commits:** 1 (1 shown, 0 dependabot hidden)
**Range:** [`391de387baba`](https://github.com/openshift-pipelines/syncer-service/commit/391de387babafdd591355e23a264eb058d815d0b) → [`cb63fda4a640`](https://github.com/openshift-pipelines/syncer-service/commit/cb63fda4a64098d2c967d118baa50ac95a53ae4c)

| SHA | Message | Author | Date | PR | Original PR | Jira | RN Type | RN Status | RN Text |
|-----|---------|--------|------|----|----|------|---------|-----------|---------|
| [cb63fda4a640](https://github.com/openshift-pipelines/syncer-service/commit/cb63fda4a64098d2c967d118baa50ac95a53ae4c) | fix: reconcile existing spoke secrets | Vibhav Bobade | 2026-10-05 | [71](https://github.com/openshift-pipelines/syncer-service/pull/71) | [112](https://github.com/openshift-pipelines/syncer-service/pull/112) | [SRVKP-13803](https://redhat.atlassian.net/browse/SRVKP-13803) | — | — | — |

---

### p12n-tekton-caches

**Upstream:** openshift-pipelines/tekton-caches
**Commits:** 0 (0 shown, 0 dependabot hidden)
**Range:** [`768eb9ad4462`](https://github.com/openshift-pipelines/tekton-caches/commit/768eb9ad44621a4c899e93e490bc0b1aef5b8865) → [`768eb9ad4462`](https://github.com/openshift-pipelines/tekton-caches/commit/768eb9ad44621a4c899e93e490bc0b1aef5b8865)


---

### pac-downstream

**Upstream:** tektoncd/pipelines-as-code
**Commits:** 1 (1 shown, 0 dependabot hidden)
**Range:** [`8141609d9b9d`](https://github.com/tektoncd/pipelines-as-code/commit/8141609d9b9d7e071c6a80949e343e33e8416823) → [`f5ded2f5b344`](https://github.com/tektoncd/pipelines-as-code/commit/f5ded2f5b3442e812030debbd1e142ea1509b561)

| SHA | Message | Author | Date | PR | Original PR | Jira | RN Type | RN Status | RN Text |
|-----|---------|--------|------|----|----|------|---------|-----------|---------|
| [f5ded2f5b344](https://github.com/tektoncd/pipelines-as-code/commit/f5ded2f5b3442e812030debbd1e142ea1509b561) | chore(deps): update x/crypto and grpc | Akshay Pant | 2026-10-14 | [3030](https://github.com/tektoncd/pipelines-as-code/pull/3030) | — | [SRVKP-14870](https://redhat.atlassian.net/browse/SRVKP-14870)<br>[SRVKP-14897](https://redhat.atlassian.net/browse/SRVKP-14897)<br>[SRVKP-14899](https://redhat.atlassian.net/browse/SRVKP-14899)<br>[SRVKP-14900](https://redhat.atlassian.net/browse/SRVKP-14900) | Release Note Not Required<br>—<br>—<br>— | —<br>—<br>—<br>— | —<br>—<br>—<br>— |

---

### tektoncd-chains

**Upstream:** tektoncd/chains
**Commits:** 11 (2 shown, 9 dependabot hidden)
**Range:** [`f9855bb7cba0`](https://github.com/tektoncd/chains/commit/f9855bb7cba0eb3bccfe672f84c410efc22f1ac3) → [`1834b4a186c9`](https://github.com/tektoncd/chains/commit/1834b4a186c9e3436c17e7d90d6b26d698264fb2)

| SHA | Message | Author | Date | PR | Original PR |
|-----|---------|--------|------|----|----|
| [393a60232f53](https://github.com/tektoncd/chains/commit/393a60232f532fc008d9e11f51f6a46948596adf) | bumping go to 1.26.6 to fix govcloud cves (#2002) | tekton-robot | 2026-09-24 | [1977](https://github.com/tektoncd/chains/pull/1977) | — |
| [d4413daf63b2](https://github.com/tektoncd/chains/commit/d4413daf63b284138b5ce9baa080546fd8153cf1) | fix: update google.golang.org/grpc to v1.83.1 — CVE-2026-84304 (#2039) | khelil | 2026-10-08 | [1977](https://github.com/tektoncd/chains/pull/1977) | — |

---

### tektoncd-cli

**Upstream:** tektoncd/cli
**Commits:** 9 (3 shown, 6 dependabot hidden)
**Range:** [`bc483c9ddff8`](https://github.com/tektoncd/cli/commit/bc483c9ddff876b72073e6787c6a1a91d132d0d2) → [`b804a182f54a`](https://github.com/tektoncd/cli/commit/b804a182f54a2e4f7dea0378b8bb79cfcaed104d)

| SHA | Message | Author | Date | PR | Original PR | Jira | RN Type | RN Status | RN Text |
|-----|---------|--------|------|----|----|------|---------|-----------|---------|
| [765bd6c9b922](https://github.com/tektoncd/cli/commit/765bd6c9b922c393b6ffec4d560bc587181e192f) | fix: use latest Tekton releases and fix ConfigMap timing | divyansh42 | 2026-10-06 | [3322](https://github.com/tektoncd/cli/pull/3322) | — | — | — | — | — |
| [9ec89e5b788a](https://github.com/tektoncd/cli/commit/9ec89e5b788a185c649c24c1c619aa1888706330) | fix(hub): restrict tkn hub install resources fetching | pratap0007 | 2026-10-06 | [3320](https://github.com/tektoncd/cli/pull/3320) | — | — | — | — | — |
| [57bc2598ebd7](https://github.com/tektoncd/cli/commit/57bc2598ebd740cdf58b2d2de3d66a2f71d54679) | fix: update google.golang.org/grpc to v1.83.1 — CVE-2026-84304 | diagrawa | 2026-10-08 | [3331](https://github.com/tektoncd/cli/pull/3331) | — | [SRVKP-14894](https://redhat.atlassian.net/browse/SRVKP-14894) | — | — | — |

---

### tektoncd-git-clone

**Upstream:** tektoncd-catalog/git-clone
**Commits:** 3 (2 shown, 1 dependabot hidden)
**Range:** [`e14615d80fc8`](https://github.com/tektoncd-catalog/git-clone/commit/e14615d80fc881ccbe450b1bd02d18162473443d) → [`e3aadb907515`](https://github.com/tektoncd-catalog/git-clone/commit/e3aadb907515c4e3e2e0cca9de0c51b3a319cadc)

| SHA | Message | Author | Date | PR | Original PR |
|-----|---------|--------|------|----|----|
| [a3772f4338bc](https://github.com/tektoncd-catalog/git-clone/commit/a3772f4338bcc6f85bd4e021349bd11d1b0e097b) | chore: bump version to v1.7.1 | Vincent Demeester | 2026-10-07 | [158](https://github.com/tektoncd-catalog/git-clone/pull/158) | — |
| [22c16b5394da](https://github.com/tektoncd-catalog/git-clone/commit/22c16b5394daae0cf1ad1dfff6721c66d887c09c) | ci: add Artifact Hub lint checks | Vincent Demeester | 2026-10-07 | [158](https://github.com/tektoncd-catalog/git-clone/pull/158) | — |

---

### tektoncd-pipeline

**Upstream:** tektoncd/pipeline
**Commits:** 16 (0 shown, 16 dependabot hidden)
**Range:** [`8f9940e28516`](https://github.com/tektoncd/pipeline/commit/8f9940e28516ce6c1788536272e81a087e97107e) → [`e69720ef6e36`](https://github.com/tektoncd/pipeline/commit/e69720ef6e3608e2e752400f40e7a8864dc4f422)

_All commits are from dependabot._


---

### tektoncd-pruner

**Upstream:** tektoncd/pruner
**Commits:** 1 (1 shown, 0 dependabot hidden)
**Range:** [`86e2e8e7deaa`](https://github.com/tektoncd/pruner/commit/86e2e8e7deaa85c091f8f04179fb40c72e90810b) → [`483fd0ba021e`](https://github.com/tektoncd/pruner/commit/483fd0ba021e873812a70b37859b47c280aa6222)

| SHA | Message | Author | Date | PR | Original PR |
|-----|---------|--------|------|----|----|
| [483fd0ba021e](https://github.com/tektoncd/pruner/commit/483fd0ba021e873812a70b37859b47c280aa6222) | fix: pin builder image | Shubham Bhardwaj | 2026-10-14 | [461](https://github.com/tektoncd/pruner/pull/461) | — |

---

### tektoncd-results

**Upstream:** tektoncd/results
**Commits:** 1 (1 shown, 0 dependabot hidden)
**Range:** [`9518ad13db9c`](https://github.com/tektoncd/results/commit/9518ad13db9c0c5d81502a5ebef0fd7a2bdb2939) → [`cde17241a0d6`](https://github.com/tektoncd/results/commit/cde17241a0d6695e28363977c3ad8cf2f11dbca7)

| SHA | Message | Author | Date | PR | Original PR | Jira | RN Type | RN Status | RN Text |
|-----|---------|--------|------|----|----|------|---------|-----------|---------|
| [cde17241a0d6](https://github.com/tektoncd/results/commit/cde17241a0d6695e28363977c3ad8cf2f11dbca7) | fix: update grpc to v1.83.1 and x/crypto to v0.56.0 | adityavshinde | 2026-10-15 | [1509](https://github.com/tektoncd/results/pull/1509) | — | [SRVKP-14871](https://redhat.atlassian.net/browse/SRVKP-14871)<br>[SRVKP-14872](https://redhat.atlassian.net/browse/SRVKP-14872)<br>[SRVKP-14878](https://redhat.atlassian.net/browse/SRVKP-14878) | —<br>—<br>— | —<br>—<br>— | —<br>—<br>— |

---

### tektoncd-triggers

**Upstream:** tektoncd/triggers
**Commits:** 1 (1 shown, 0 dependabot hidden)
**Range:** [`266f9010cd6b`](https://github.com/tektoncd/triggers/commit/266f9010cd6b72dabdcf4d56b55b6e1201095a87) → [`9f74bff94034`](https://github.com/tektoncd/triggers/commit/9f74bff9403412cbbce96dcb9c44b75b832bb5e6)

| SHA | Message | Author | Date | PR | Original PR |
|-----|---------|--------|------|----|----|
| [9f74bff94034](https://github.com/tektoncd/triggers/commit/9f74bff9403412cbbce96dcb9c44b75b832bb5e6) | cherry-pick add automatic GitHub draft release creation from main | Vincent Demeester | 2026-10-07 | [2173](https://github.com/tektoncd/triggers/pull/2173) | — |

---

### tekton-kueue

**Upstream:** tektoncd/tekton-kueue
**Commits:** 5 (5 shown, 0 dependabot hidden)
**Range:** [`2f0ab5fd4ecd`](https://github.com/tektoncd/tekton-kueue/commit/2f0ab5fd4ecdb700612a1d9a0052ac5cf21dbe44) → [`674dd6371a3e`](https://github.com/tektoncd/tekton-kueue/commit/674dd6371a3e8017ed89341e9d01eb7f6f10e411)

| SHA | Message | Author | Date | PR | Original PR | Jira | RN Type | RN Status | RN Text |
|-----|---------|--------|------|----|----|------|---------|-----------|---------|
| [211245eba98e](https://github.com/tektoncd/tekton-kueue/commit/211245eba98e50440ce12256c4cb1c3bf0691b01) | fix: bump go directive to 1.26.6 for stdlib CVE fixes | Khurram Baig | 2026-09-28 | [547](https://github.com/tektoncd/tekton-kueue/pull/547) | — | [SRVKP-13466](https://redhat.atlassian.net/browse/SRVKP-13466)<br>[SRVKP-13483](https://redhat.atlassian.net/browse/SRVKP-13483) | —<br>— | —<br>— | —<br>— |
| [5fb101fe961e](https://github.com/tektoncd/tekton-kueue/commit/5fb101fe961e7386e6b2fb5cc9a4e3b97474a7cc) | fix: backport release tooling to v0.4 | Vibhav Bobade | 2026-09-30 | [542](https://github.com/tektoncd/tekton-kueue/pull/542) | — | [SRVKP-13561](https://redhat.atlassian.net/browse/SRVKP-13561) | Release Note Not Required | — | — |
| [2448e76bc019](https://github.com/tektoncd/tekton-kueue/commit/2448e76bc0193d2d3450757b90373cd787429a2e) | Sync release branch ownership | Vibhav Bobade | 2026-10-01 | [545](https://github.com/tektoncd/tekton-kueue/pull/545) | — | [SRVKP-13258](https://redhat.atlassian.net/browse/SRVKP-13258)<br>[SRVKP-13561](https://redhat.atlassian.net/browse/SRVKP-13561) | —<br>Release Note Not Required | —<br>— | —<br>— |
| [5a0637f520d6](https://github.com/tektoncd/tekton-kueue/commit/5a0637f520d6cc5cdd47c93ec928f93c0b2eab1d) | KFLUXINFRA-4254 Add agent files defense CI checks | Gal Levi | 2026-10-11 | [563](https://github.com/tektoncd/tekton-kueue/pull/563) | [548](https://github.com/tektoncd/tekton-kueue/pull/548) | — | — | — | — |
| [674dd6371a3e](https://github.com/tektoncd/tekton-kueue/commit/674dd6371a3e8017ed89341e9d01eb7f6f10e411) | Merge pull request #563 from glevi-rh/backport/agent-files-defense-release-v0.4.x | Gal Levi | 2026-10-11 | [563](https://github.com/tektoncd/tekton-kueue/pull/563) | — | — | — | — | — |

---

