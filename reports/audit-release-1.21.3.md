# Release Audit: 1.21.3

**Source:** `reports/release_1.21.3.json`<br>
**Total commits:** 342 across 14 components<br>
**Already linked (by Go tool):** 23 · **Newly linked:** 26 · **Unlinked:** 293 · **Dependabot (skipped):** 60<br>

## Index

| Component | Commits | Jira | Unsynced | Unmatched Jiras |
|-----------|---------|------|----------|-----------------|
| [tektoncd-pipeline](#tektoncd-pipeline) | 30 | 11/30 | 0 | 0 |
| [console-plugin](#console-plugin) | 183 | 20/183 | 0 | 1 |
| [operator](#operator) | 11 | 2/11 | 0 | 0 |
| [manual-approval-gate](#manual-approval-gate) | 3 | 2/3 | 0 | 0 |
| [pipelines-as-code](#pipelines-as-code) | 16 | 2/16 | 0 | 0 |
| [tektoncd-cli](#tektoncd-cli) | 7 | 0/7 | 0 | 0 |
| [console-plugin-pf5](#console-plugin-pf5) | — | — | — | — |
| [tektoncd-triggers](#tektoncd-triggers) | 0 | 0/0 | 0 | 0 |
| [git-init](#git-init) | 56 | 0/56 | 0 | 0 |
| [tekton-caches](#tekton-caches) | 9 | 5/9 | 0 | 0 |
| [tektoncd-hub](#tektoncd-hub) | 9 | 1/9 | 0 | 0 |
| [tektoncd-results](#tektoncd-results) | 10 | 2/10 | 0 | 0 |
| [tektoncd-pruner](#tektoncd-pruner) | 3 | 0/3 | 0 | 0 |
| [tektoncd-chains](#tektoncd-chains) | 5 | 4/5 | 0 | 0 |

---

<details>
<summary><h2>tektoncd-pipeline</h2> — 30 commits, 11 linked, 28 dependabot</summary>

**Upstream:** tektoncd/pipeline · **Commits:** 30 · **Unsynced:** 0 · **Dependabot:** 28

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| b728378 | Remove deprecated `// +build` directive from most files | Vincent Demeester | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) | [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) *(PR body)* | **Type:** Release Notes Not Required<br>**Text:** Remove deprecated Go build directive syntax from source files. *(generated)* |
| c1810c5 | test: implement parallel/serial test categorization system | Vincent Demeester | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) | [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) *(PR body)* | **Type:** Release Notes Not Required<br>**Text:** Add parallel and serial test categorization system for e2e test execution. *(generated)* |
| f7f9584 | refactor: add clock injection to cache for testing | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) | [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) *(PR body)* | **Type:** Release Notes Not Required<br>**Text:** Add clock injection to resolver cache for improved test isolation. *(generated)* |
| c639f88 | fix: Align cache configstore with framework implementation | Seunghun Shin | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) | [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) *(PR body)* | **Type:** Bug Fix<br>**Text:** Fix resolver cache race condition causing duplicate upstream pulls under concurrent load. *(extracted)* |
| 48765b2 | fix: resolve cache race with singleflight | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) | [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) *(PR body)* | **Type:** Bug Fix<br>**Text:** Deduplicate concurrent resolver cache requests with singleflight to prevent cache stampede and reduce upstream API calls (e.g., GitLab 429 throttling). *(extracted)* |
| 192a6c1 | fix: remove corrupted cache entries on type error | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) | [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) *(PR body)* | **Type:** Bug Fix<br>**Text:** Remove corrupted cache entries on type assertion errors to prevent stale data from blocking subsequent resolution attempts. *(generated)* |
| 6bbcfb1 | fix: use TextParser struct directly for prometheus/common v0.62.0 compatibility | Vibhav Bobade | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) | [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) *(PR body)* | **Type:** Bug Fix<br>**Text:** Fix compatibility with prometheus/common v0.62.0 by using TextParser struct directly. *(generated)* |
| 6e0504b | fix: add missing @test:execution annotation to e2e test | Vibhav Bobade | [#9365](https://github.com/tektoncd/pipeline/pull/9365) → [#9665](https://github.com/tektoncd/pipeline/pull/9665) | [KONFLUX-6727](https://redhat.atlassian.net/browse/KONFLUX-6727) *(PR body)* | **Type:** Release Notes Not Required<br>**Text:** Add missing test execution annotation to e2e test. *(generated)* |
| b2e34cf | fix: bump Go to 1.24.13 to fix CVE-2025-61728, CVE-2025-61726, CVE-2025-61729 | Vibhav Bobade | [#9734](https://github.com/tektoncd/pipeline/pull/9734) | — | — |
| 00055c0 | fix: pin registry image and relax log-based cache assertion | Vincent Demeester | [#9734](https://github.com/tektoncd/pipeline/pull/9734) | — | — |
| d855289 | fix(e2e): improve dind-sidecar probe configuration for reliability | Vincent Demeester | [#9734](https://github.com/tektoncd/pipeline/pull/9734) | — | — |
| 6765156 | fix(e2e): replace ubuntu with busybox in dind-sidecar test Dockerfile | Vibhav Bobade | [#9734](https://github.com/tektoncd/pipeline/pull/9734) | — | — |
| 3f60a5d | Security: reject system API token with user-controlled serverURL | Vincent Demeester | [#10350](https://github.com/tektoncd/pipeline/pull/10350) | — | — |
| 060e18c | fix: prevent git argument injection via revision parameter (GHSA-94jr-7pqp-xhcq) | Vincent Demeester | [#10350](https://github.com/tektoncd/pipeline/pull/10350) | — | — |
| 3090efa | fix: normalize VolumeMount paths before /tekton/ restriction check | Vincent Demeester | [#10350](https://github.com/tektoncd/pipeline/pull/10350) | — | — |
| 0ea0634 | fix: strip resolver prefixes and use non-capturing group for pattern anchoring | Vincent Demeester | [#10350](https://github.com/tektoncd/pipeline/pull/10350) | — | — |
| c9eef11 | fix: limit HTTP resolver response body size to prevent OOM DoS | Vincent Demeester | [#10350](https://github.com/tektoncd/pipeline/pull/10350) | — | — |
| 43cc433 | fix: trim whitespace from source URI before pattern matching | Vincent Demeester | [#10350](https://github.com/tektoncd/pipeline/pull/10350) | — | — |
| 99f5732 | fix: set GOTOOLCHAIN=auto in publish task for Go version compatibility | Vincent Demeester | [#10350](https://github.com/tektoncd/pipeline/pull/10350) | — | — |
| 215d4ce | fix: convert pod latency metric to histogram and remove pod label | Vibhav Bobade | [#9530](https://github.com/tektoncd/pipeline/pull/9530) → [#10109](https://github.com/tektoncd/pipeline/pull/10109) | [SRVKP-6762](https://redhat.atlassian.net/browse/SRVKP-6762) *(keyword)* | **Type:** Bug Fix<br>**Text:** The tekton_pipelines_controller_taskruns_pod_latency_milliseconds metric has been converted from a Gauge to a Histogram and the pod label has been removed. Dashboards or alerts referencing this metric will need to be updated to use histogram_quantile() instead of direct value queries. *(extracted)* |
| 5fa8e84 | docs: add pod latency histogram metric to metrics.md | Vibhav Bobade | [#9530](https://github.com/tektoncd/pipeline/pull/9530) → [#10109](https://github.com/tektoncd/pipeline/pull/10109) | [SRVKP-6762](https://redhat.atlassian.net/browse/SRVKP-6762) *(keyword)* | **Type:** Release Notes Not Required<br>**Text:** Add documentation for the new pod latency histogram metric to metrics reference. *(generated)* |
| 0140b68 | [cherry-pick: release-v1.6.x] fix: bump google.golang.org/grpc to 1.79.3 (CVE... | Vibhav Bobade | [#9909](https://github.com/tektoncd/pipeline/pull/9909) | [SRVKP-12033](https://redhat.atlassian.net/browse/SRVKP-12033) *(keyword)* | **Type:** CVE<br>**Text:** Bump google.golang.org/grpc from 1.75.0 to 1.80.0 to fix CVE-2026-33186 (gRPC-Go authorization bypass via missing leading slash in :path header). *(extracted)* |
| 0a30367 | fix(pipelinerun): use generateName for anonymous pipeline label | Andrew Thorp | [#9826](https://github.com/tektoncd/pipeline/pull/9826) → [#10107](https://github.com/tektoncd/pipeline/pull/10107) | — | — |
| 6854423 | fix: resolve goroutine leak from unbuffered channels in resolver reconcilers | Shubham Bhardwaj | [#10098](https://github.com/tektoncd/pipeline/pull/10098) → [#10111](https://github.com/tektoncd/pipeline/pull/10111) | — | — |
| da7e868 | fix(resolvers): validate data is Tekton object in resolver framework | Andrew Thorp | [#9588](https://github.com/tektoncd/pipeline/pull/9588) → [#9962](https://github.com/tektoncd/pipeline/pull/9962) | — | — |
| 4ff64cf | fix: prevent controller CPU variant from leaking into platform commands | Vincent Demeester | [#10077](https://github.com/tektoncd/pipeline/pull/10077) → [#10162](https://github.com/tektoncd/pipeline/pull/10162) | — | — |
| 64ca547 | fix: add automated draft release support to release pipeline | Vincent Demeester | [#10203](https://github.com/tektoncd/pipeline/pull/10203) → [#10213](https://github.com/tektoncd/pipeline/pull/10213) | — | — |
| 8375924 | fix(resolvers): Allow ResolutionRequests to resolve all Tekton kinds | Andrew Thorp | [#10242](https://github.com/tektoncd/pipeline/pull/10242) → [#10254](https://github.com/tektoncd/pipeline/pull/10254) | — | — |
| f53bb02 | build: bump go directive to 1.25 | Vincent Demeester | [#10028](https://github.com/tektoncd/pipeline/pull/10028) | — | — |
| b54dcda | fix: bump Go for CVEs | Vibhav Bobade | [#10341](https://github.com/tektoncd/pipeline/pull/10341) | — | — |

</details>

<details>
<summary><h2>console-plugin</h2> — 183 commits, 20 linked, 0 dependabot</summary>

**Upstream:** openshift-pipelines/console-plugin · **Commits:** 183 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 89018d7 | Fix: Reset submit error on YAML editor changes | Anwesha Palit | — | — | — |
| f6b07dd | Add anwesha-palit-redhat as approver and arvindk-softwaredev as reviewer | Vikram Raj | — | — | — |
| 8ebcf41 | fix: updated field name for capturing next page token in tkr | Anwesha Palit | — | — | — |
| c381612 | feat: update filter from last weeks to last week | Anwesha Palit | — | — | — |
| ddda0af | Merge pull request #818 from anwesha-palit-redhat/fix/SRVKP-9428 | anwesha-palit-redhat | — | — | — |
| 904e9c5 | Merge pull request #815 from anwesha-palit-redhat/fix/SRVKP-9397 | anwesha-palit-redhat | — | — | — |
| 2706d9e | Add console.navigation/section for Pipelines | Vikram Raj | — | — | — |
| 72723a0 | updated loading component to use spinner and used it in pipelineoverview page | Arvind | — | — | — |
| 5a34e36 | Merge pull request #825 from arvindk-softwaredev/fix/SRVKP-9436 | anwesha-palit-redhat | — | — | — |
| 008b2d7 | feat: added abort controller, increased API timeout and added placeholder err... | Anwesha Palit | — | — | — |
| fe98eb9 | fix: preserve error state when namespace is all-namespace | Anwesha Palit | — | — | — |
| 3f98d39 | Merge pull request #824 from anwesha-palit-redhat/fix/SRVKP-9427 | anwesha-palit-redhat | — | — | — |
| 3a7d0b7 | Merge pull request #828 from anwesha-palit-redhat/fix/SRVKP-9450 | anwesha-palit-redhat | — | — | — |
| 549980b | fix: updated condition to show error message instead of the loader after api ... | Anwesha Palit | — | — | — |
| 28b7765 | feat: added redux state for pipeline overview filter to persist | Anwesha Palit | — | — | — |
| b1ac353 | Merge pull request #833 from anwesha-palit-redhat/fix/SRVKP-9591 | anwesha-palit-redhat | — | — | — |
| ec1aa2f | feat: added error banner and loader for average duration in pipeline metrics | Anwesha Palit | — | — | — |
| 3d3f04f | pipeline overviewpage error css fix | Arvind | — | — | — |
| 332b1f6 | Merge pull request #832 from arvindk-softwaredev/fix/SRVKP-9470 | anwesha-palit-redhat | — | — | — |
| 8567c4c | feat: upgrading to PF6 - updated dependencies and ran codemodes suites | Anwesha Palit | — | — | — |
| d57e3ee | Merge pull request #836 from anwesha-palit-redhat/feat/SRVKP-SRVKP-9675 | anwesha-palit-redhat | — | — | — |
| 44d666d | Merge pull request #845 from anwesha-palit-redhat/feat/SRVKP-9273 | anwesha-palit-redhat | — | — | — |
| 389583d | Merge pull request #834 from anwesha-palit-redhat/feat/SRVKP-9607 | anwesha-palit-redhat | — | — | — |
| d8d244d | PF6 Page Heading fixes and refactoring | Arvind | — | — | — |
| 21ec2dc | fix: css fix for sub header issue | Anwesha Palit | — | — | — |
| 458bacc | Merge pull request #850 from anwesha-palit-redhat/SRVKP-9787 | openshift-merge-bot[bot] | — | — | — |
| 2ffc6e2 | Merge pull request #849 from arvindk-softwaredev/feature/SRVKP-9625 | anwesha-palit-redhat | — | — | — |
| b91e0bb | feat: start pipeline modal fixes and replacing useModal with useOverlay with ... | Anwesha Palit | — | — | — |
| 473ec34 | SRVKP-9675: Add Error Banner and Loading Spinner to Pipeline Metrics Page for... | anwesha-palit-redhat | — | — | — |
| 123e816 | Merge pull request #834 from anwesha-palit-redhat/feat/SRVKP-9607 | anwesha-palit-redhat | — | — | — |
| 541fd49 | Merge pull request #851 from anwesha-palit-redhat/feat/SRVKP-9627 | openshift-merge-bot[bot] | — | — | — |
| 4f8773c | deprecated errormodal and reset modal changes for PipelineBuilder | Anwesha Palit | — | — | — |
| 813677d | Merge pull request #854 from arvindk-softwaredev/feat/SRVKP-9918 | Arvind Krishnamurthy | — | — | — |
| d6536ce | Merge pull request #852 from anwesha-palit-redhat/cherry-pick-SRVKP-9605-SRVK... | anwesha-palit-redhat | — | — | — |
| d119fbf | feat: Replaced the action menu with LazyActionMenu | Anwesha Palit | — | — | — |
| e70e242 | Merge pull request #856 from anwesha-palit-redhat/feat/SRVKP-9786 | Arvind Krishnamurthy | — | — | — |
| 13e1912 | feat: CSS fix for log viewer and nav items | Anwesha Palit | — | — | — |
| b0da77c | translations for pipeline overview page label changes | Arvind | — | — | — |
| 2e4b9a0 | Merge pull request #855 from arvindk-softwaredev/fix/SRVKP-10032 | anwesha-palit-redhat | — | — | — |
| dba4553 | translations for pipeline overview page label changes | Arvind | — | — | — |
| 5f38376 | Merge pull request #858 from openshift-cherrypick-robot/cherry-pick-855-to-ma... | anwesha-palit-redhat | — | — | — |
| 4ba2184 | Merge pull request #857 from anwesha-palit-redhat/feat/SRVKP-9983 | Arvind Krishnamurthy | — | — | — |
| ae1338b | PF6 changes for QuickSearch and PipelineBuilder | Arvind | — | — | — |
| f7b885c | Merge pull request #860 from arvindk-softwaredev/feature/SRVKP-9785 | anwesha-palit-redhat | — | — | — |
| c262024 | fix: ResourceDropdown.tsx update to support React functional component and PF... | Anwesha Palit | — | — | — |
| 5795b94 | css fix for pause play icon in event streaming | Anwesha Palit | — | — | — |
| f2825c5 | css fix for eventlistener triggers | Arvind | — | — | — |
| 6aeeecf | deprecated modal.tsx | Arvind | — | — | — |
| 8dab896 | Merge pull request #865 from anwesha-palit-redhat/fix/SRVKP-10077 | Arvind Krishnamurthy | — | — | — |
| 1844f5a | Merge pull request #861 from anwesha-palit-redhat/fix/SRVKP-10078 | Arvind Krishnamurthy | — | — | — |
| 6373725 | Merge pull request #868 from arvindk-softwaredev/feature/SRVKP-10069 | anwesha-palit-redhat | — | — | — |
| 81e4367 | Merge pull request #870 from arvindk-softwaredev/feature/SRVKP-10072 | anwesha-palit-redhat | — | — | — |
| ca3aa91 | added metadata for CRDs | Arvind | — | — | — |
| c03fbf8 | fix: css fix for logviewer to preserve whitespace | Anwesha Palit | — | — | — |
| 31f1a7c | Merge pull request #867 from arvindk-softwaredev/fix/SRVKP-10310 | Arvind Krishnamurthy | — | — | — |
| 318e71c | added metadata for CRDs | Arvind | — | — | — |
| 8ed3dab | Merge pull request #873 from openshift-cherrypick-robot/cherry-pick-867-to-ma... | Arvind Krishnamurthy | — | — | — |
| c373d1e | Merge pull request #866 from anwesha-palit-redhat/fix/SRVKP-10035 | anwesha-palit-redhat | — | — | — |
| c9ae510 | css spacing fix | Arvind | — | — | — |
| d9de09c | deprecate react-helemt and use DocumentTitle | Arvind | — | — | — |
| a87376c | Merge pull request #880 from arvindk-softwaredev/feature/SRVKP-10076 | Arvind Krishnamurthy | — | — | — |
| 3cacd32 | feat: support ansi colors for logviewer | Anwesha Palit | — | — | — |
| 2d830f3 | Merge pull request #881 from anwesha-palit-redhat/feat/SRVKP-10407 | anwesha-palit-redhat | — | — | — |
| a84982d | Merge pull request #879 from arvindk-softwaredev/feature/SRVKP-9711 | Arvind Krishnamurthy | — | — | — |
| c76d7ee | feat: adding main_ocp_4.22 and PR workflow | Anwesha Palit | — | — | — |
| 95f7ea0 | Merge pull request #883 from anwesha-palit-redhat/feat/SRVKP-10470 | anwesha-palit-redhat | — | — | — |
| 2c38f55 | SRVKP-10439: Yarn 4 migration | Anwesha Palit | — | — | — |
| f749ade | SRVKP-10439: updating ci-operator.yaml to fix build errors in release PR | Anwesha Palit | — | — | — |
| d0d194d | SRVKP-10439: updating ci-operator.yaml to fix build errors in release PR | Anwesha Palit | — | — | — |
| a583384 | Merge branch 'main' into test_yarn4 | anwesha-palit-redhat | — | — | — |
| 48c084f | Merge pull request #920 from openshift-pipelines/test_yarn4 | Vincent Demeester | — | — | — |
| b2c07f7 | feat/SRVKP-10439-update-ci-operator-config | Anwesha Palit | — | — | — |
| 502a8f8 | feat: adding Arvind as approver for the console-plugin repo | anwesha-palit-redhat | [#925](https://github.com/openshift-pipelines/console-plugin/pull/925) | — | — |
| acefa43 | Merge pull request #925 from openshift-pipelines/anwesha-palit-redhat-patch-1 | anwesha-palit-redhat | [#925](https://github.com/openshift-pipelines/console-plugin/pull/925) | — | — |
| 7aa5097 | feat: adding logs for multicluster | Anwesha Palit | — | — | — |
| 3137a44 | Merge pull request #926 from anwesha-palit-redhat/feat/SRVKP-8974-MC-Changes | Arvind Krishnamurthy | — | — | — |
| 3d3f4df | Update: docker-build-ta location and nginx, node images | Anwesha Palit | — | — | — |
| fbb43ab | Merge pull request #935 from anwesha-palit-redhat/update_location_docker-buil... | Arvind Krishnamurthy | — | — | — |
| c6cdb5c | Yarn 4 multi-arch support for main branch | Anwesha Palit | — | — | — |
| 14aa0da | Merge pull request #937 from anwesha-palit-redhat/yarn4-support-for-multi-arc... | anwesha-palit-redhat | — | — | — |
| 39a5efc | multicluster changes to use managedBy field | Arvind | — | — | — |
| f9d72c9 | Merge pull request #944 from arvindk-softwaredev/feat/SRVKP-10808 | anwesha-palit-redhat | — | — | — |
| 9ecef9c | feat: adding icon for multicluster plr in hub cluster | Anwesha Palit | — | — | — |
| f994ff7 | Merge pull request #934 from anwesha-palit-redhat/feat/SRVKP-10807 | Arvind Krishnamurthy | — | — | — |
| 1bae704 | feat: addting translations for multi cluster logs | Anwesha Palit | — | — | — |
| f2586d7 | Merge pull request #929 from anwesha-palit-redhat/feat/SRVKP-10785 | Arvind Krishnamurthy | — | — | — |
| 5d4951f | CVE changes | Anwesha Palit | — | — | — |
| eb9be4b | Merge pull request #945 from anwesha-palit-redhat/CVE-main | Arvind Krishnamurthy | — | — | — |
| 3ce8240 | updaed hub cluster detection to use multi-cluster-role | Arvind | — | — | — |
| f057021 | Merge pull request #959 from arvindk-softwaredev/feat/SRVKP-10786 | openshift-merge-bot[bot] | — | — | — |
| a2d6b8c | feat:CVEs SRVKP-10058 SRVKP-9700 and yarn 4 migration with node20 support | Anwesha Palit | — | — | — |
| 049dd9e | Merge pull request #955 from arvindk-softwaredev/feat/SRVKP-10743 | anwesha-palit-redhat | — | — | — |
| e686c3b | fix konflux build | Anwesha Palit | — | — | — |
| 931f7d8 | Merge pull request #963 from anwesha-palit-redhat/fix-konflux/SRVKP-10743 | anwesha-palit-redhat | — | — | — |
| e104fb0 | fix: marking tests as broken | Anwesha Palit | — | — | — |
| f5c5a3b | Merge pull request #964 from anwesha-palit-redhat/fix/SRVKP-10743 | Arvind Krishnamurthy | — | — | — |
| 64b6e5f | fix: marking tests as broken | Anwesha Palit | — | — | — |
| 46b5516 | Merge pull request #965 from arvindk-softwaredev/fix/SRVKP-10743 | openshift-merge-bot[bot] | — | — | — |
| bcee6cd | fix: added missing translation to fix CI job | Anwesha Palit | — | — | — |
| 3cc46dd | Merge pull request #972 from anwesha-palit-redhat/fix/SRVKP-10953 | Arvind Krishnamurthy | — | — | — |
| 0daf439 | approver list tag css fix | Anwesha Palit | — | — | — |
| 4646e73 | Merge pull request #974 from arvindk-softwaredev/feat/SRVKP-10080 | Anwesha Palit | — | — | — |
| 6572ad0 | delete: removing .konflux and .tekton dir from upstream console-plugin | Anwesha Palit | — | — | — |
| 051af8d | [bot:main] update konflux configuration | openshift-pipelines-bot | — | — | — |
| 0b8374d | Merge pull request #975 from anwesha-palit-redhat/delete/SRVKP-10961 | openshift-merge-bot[bot] | — | — | — |
| c386391 | renamed main_ocp_4.22 branch to master | Arvind | — | — | — |
| 173c1b6 | Merge pull request #995 from arvindk-softwaredev/feat/SRVKP-11001 | Anwesha Palit | — | — | — |
| 9e29790 | removing all .konflux and .tekton files from main_ocp_4.22 | Anwesha Palit | — | — | — |
| 2f6d8e1 | Merge pull request #992 from anwesha-palit-redhat/cherry-pick/SRVKP-10961 | openshift-merge-bot[bot] | — | — | — |
| f9bcbd9 | react 18 readiness | Anwesha Palit | — | — | — |
| 14a239b | Merge pull request #999 from anwesha-palit-redhat/feat/SRVKP-11003 | Arvind Krishnamurthy | — | — | — |
| 660064f | fix: ci job for docker image is failing | Anwesha Palit | — | — | — |
| b0fa093 | Merge pull request #1001 from anwesha-palit-redhat/feat/SRVKP-11123 | openshift-merge-bot[bot] | — | — | — |
| f5c64d0 | feat: adding custom filter for ConsoleDataView | Anwesha Palit | — | — | — |
| c64fba5 | Merge pull request #1003 from anwesha-palit-redhat/feat/SRVKP-11167 | openshift-merge-bot[bot] | — | — | — |
| 3fde794 | migrate to ConsoleDataView for PL and PLR list | Arvind | — | — | — |
| 0b44ca5 | Merge pull request #1011 from arvindk-softwaredev/feat/SRVKP-9626 | openshift-merge-bot[bot] | — | — | — |
| 018b5a6 | feat: Add consoledataview for tasks and taskruns | Anwesha Palit | — | — | — |
| c387d0b | Merge pull request #1010 from anwesha-palit-redhat/feat/SRVKP-9710 | openshift-merge-bot[bot] | — | — | — |
| 41f2204 | Merge branch 'main' into master-with-main | Anwesha Palit | — | — | — |
| 5701cf0 | refactor pipelines-overview page to use PatternFly v6 utilities | shverma | — | — | — |
| a777dbd | Merge pull request #869 from pratap0007/overview-pr-status-card | openshift-merge-bot[bot] | — | — | — |
| a273bd7 | Delete PipelineMetrics.scss file and Change some files for responsiveness and... | adityavshinde | — | — | — |
| 54f34c7 | fix: remove nested scrollbars and align metric cards for tekton-results | Ankur Sinha | — | — | — |
| 26e11bf | Merge pull request #1004 from adityavshinde/fix/SRVKP-10871 | openshift-merge-bot[bot] | — | — | — |
| eff032d | Merge pull request #1005 from ankrsinha/fix/SRVKP-10073 | openshift-merge-bot[bot] | — | — | — |
| df7f858 | Merge pull request #1014 from anwesha-palit-redhat/master-with-main | openshift-merge-bot[bot] | — | — | — |
| 9b20e29 | fix: updated the pf-v5 classnames to pf-v6 | Anwesha Palit | — | — | — |
| 56db337 | fix: updated the loader logic so that the taskrun details page is visible for... | Anwesha Palit | — | — | — |
| 1100e96 | fix: remove sidebar list bullets for PipelineRuns and Triggers on the Topolog... | Ankur Sinha | — | — | — |
| ed537a7 | Merge pull request #1017 from anwesha-palit-redhat/SRVKP-11190 | openshift-merge-bot[bot] | — | — | — |
| afcc58f | fix for fetching TRs of specific PLR | Arvind | — | — | — |
| 6b4c118 | Merge pull request #1016 from anwesha-palit-redhat/fix/SRVKP-11192 | openshift-merge-bot[bot] | — | — | — |
| 8e91e58 | fix for archived plr loader logic | Anwesha Palit | — | — | — |
| e7290ff | Merge pull request #1021 from arvindk-softwaredev/fix/SRVKP-11393 | openshift-merge-bot[bot] | — | — | — |
| 81d30ff | Merge pull request #1023 from arvindk-softwaredev/fix/SRVKP-11391 | openshift-merge-bot[bot] | — | — | — |
| 9612642 | SRVKP-11398: bumped console sdk packages to 4.22.0-prerelease.2 (#1025) | Arvind Krishnamurthy | — | — | — |
| 281b929 | feat: upgrading react-router | Anwesha Palit | — | — | — |
| b48ed50 | Merge pull request #1024 from anwesha-palit-redhat/feat/SRVKP-10306 | openshift-merge-bot[bot] | — | — | — |
| 4d31b7e | Merge pull request #1019 from ankrsinha/fix/SRVKP-9789 | openshift-merge-bot[bot] | — | — | — |
| 55c5362 | fix: pipeline topology task status fix and update default | Anwesha Palit | — | — | — |
| a635dae | Merge pull request #1027 from anwesha-palit-redhat/fix/SRVKP-10037 | openshift-merge-bot[bot] | — | — | — |
| 64b6dad | conditinally navigate to details page on PLR re-run action | Arvind | — | — | — |
| 9da4653 | Merge pull request #1035 from arvindk-softwaredev/feat/SRVKP-11394 | openshift-merge-bot[bot] | — | — | — |
| 7f8ba75 | feat: updating repository tabs with consoledataview | Anwesha Palit | — | — | — |
| 700f888 | Merge pull request #1036 from anwesha-palit-redhat/feat/SRVKP-9709 | openshift-merge-bot[bot] | — | — | — |
| a22d4b9 | Improve UI and layout for pipeline overview loading states | Aditya Shinde | — | — | — |
| 4dc6fca | Merge pull request #1018 from adityavshinde/fix/SRVKP-11379 | openshift-merge-bot[bot] | — | — | — |
| d23f435 | Improve UI and layout for pipeline overview loading states | Aditya Shinde | — | — | — |
| a9caa8b | Merge pull request #1022 from adityavshinde/fix/SRVKP-11383 | openshift-merge-bot[bot] | — | — | — |
| 06a86f8 | refactor: refactored useTaskRuns.ts | Anwesha Palit | — | — | — |
| 66b56e1 | Merge pull request #1037 from anwesha-palit-redhat/fix/SRVKP-11395 | openshift-merge-bot[bot] | — | — | — |
| a4623e2 | updated consoledataview for eventlisteners and triggertemplates tab | Arvind | — | — | — |
| cf80d0a | Merge pull request #1040 from arvindk-softwaredev/feat/SRVKP-9712 | openshift-merge-bot[bot] | — | — | — |
| a3b9291 | fixed color for EPL icon | Arvind | — | — | — |
| dc45cbc | Merge pull request #1045 from arvindk-softwaredev/fix/SRVKP-11400 | openshift-merge-bot[bot] | — | — | — |
| e6761e7 | feat: removed LoadingInline wrapper and updated the LoadingBox component | Anwesha Palit | [#1042](https://github.com/openshift-pipelines/console-plugin/pull/1042) | [SRVKP-11566](https://redhat.atlassian.net/browse/SRVKP-11566) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| bbdb251 | Merge pull request #1042 from anwesha-palit-redhat/feat/SRVKP-11566 | openshift-merge-bot[bot] | [#1042](https://github.com/openshift-pipelines/console-plugin/pull/1042) | [SRVKP-11566](https://redhat.atlassian.net/browse/SRVKP-11566) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 66544f6 | feat: added consoledataview and modified styles | Anwesha Palit | [#1041](https://github.com/openshift-pipelines/console-plugin/pull/1041) | [SRVKP-11399](https://redhat.atlassian.net/browse/SRVKP-11399) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 993182c | Merge pull request #1041 from anwesha-palit-redhat/feat/SRVKP-11399 | openshift-merge-bot[bot] | [#1041](https://github.com/openshift-pipelines/console-plugin/pull/1041) | [SRVKP-11399](https://redhat.atlassian.net/browse/SRVKP-11399) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| e08957b | feat: adding consoledataview and lazyactionmenu for the remaining resources | Anwesha Palit | [#1046](https://github.com/openshift-pipelines/console-plugin/pull/1046) | [SRVKP-9713](https://redhat.atlassian.net/browse/SRVKP-9713) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| ae1bb11 | fix: added key for same plr names so that they rerender everytime a new plr i... | Anwesha Palit | [#1048](https://github.com/openshift-pipelines/console-plugin/pull/1048) | [SRVKP-11565](https://redhat.atlassian.net/browse/SRVKP-11565) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 1c78bea | Merge pull request #1048 from anwesha-palit-redhat/fix/SRVKP-11565 | openshift-merge-bot[bot] | [#1048](https://github.com/openshift-pipelines/console-plugin/pull/1048) | [SRVKP-11565](https://redhat.atlassian.net/browse/SRVKP-11565) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 1e35a4c | upgraded console sdk to 4.22.0-prerelease.3 | Arvind | [#1053](https://github.com/openshift-pipelines/console-plugin/pull/1053) | [SRVKP-11711](https://redhat.atlassian.net/browse/SRVKP-11711) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| d99e842 | Merge pull request #1046 from anwesha-palit-redhat/feat/SRVKP-9713 | openshift-merge-bot[bot] | [#1046](https://github.com/openshift-pipelines/console-plugin/pull/1046) | [SRVKP-9713](https://redhat.atlassian.net/browse/SRVKP-9713) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| dd67cb1 | migrate ApprovalList to ConsoleDataView | Arvind | [#1047](https://github.com/openshift-pipelines/console-plugin/pull/1047) | [SRVKP-11401](https://redhat.atlassian.net/browse/SRVKP-11401) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 02229d1 | updated customData type and options for DataViewFilter | Arvind | [#1047](https://github.com/openshift-pipelines/console-plugin/pull/1047) | [SRVKP-11401](https://redhat.atlassian.net/browse/SRVKP-11401) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| dcdc747 | Merge pull request #1047 from arvindk-softwaredev/feat/SRVKP-11401 | openshift-merge-bot[bot] | [#1047](https://github.com/openshift-pipelines/console-plugin/pull/1047) | [SRVKP-11401](https://redhat.atlassian.net/browse/SRVKP-11401) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 27b5419 | translations for master branch | Arvind | [#1052](https://github.com/openshift-pipelines/console-plugin/pull/1052) | [SRVKP-11715](https://redhat.atlassian.net/browse/SRVKP-11715) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 61f01b9 | Merge pull request #1052 from arvindk-softwaredev/feat/SRVKP-11715 | openshift-merge-bot[bot] | [#1052](https://github.com/openshift-pipelines/console-plugin/pull/1052) | [SRVKP-11715](https://redhat.atlassian.net/browse/SRVKP-11715) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 4cdecf1 | Merge pull request #1053 from arvindk-softwaredev/feat/SRVKP-11711 | openshift-merge-bot[bot] | [#1053](https://github.com/openshift-pipelines/console-plugin/pull/1053) | [SRVKP-11711](https://redhat.atlassian.net/browse/SRVKP-11711) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| c27daa5 | fix: renamed classnames to avoid classname collisions | Anwesha Palit | [#1049](https://github.com/openshift-pipelines/console-plugin/pull/1049) | [SRVKP-11708](https://redhat.atlassian.net/browse/SRVKP-11708) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 13fc29d | Merge pull request #1049 from anwesha-palit-redhat/fix/SRVKP-11708 | openshift-merge-bot[bot] | [#1049](https://github.com/openshift-pipelines/console-plugin/pull/1049) | [SRVKP-11708](https://redhat.atlassian.net/browse/SRVKP-11708) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 23ca364 | SRVKP-8292: Unify Stop and Cancel PipelineRun visibility logic [release-v1.23... | OpenShift Cherrypick Robot | [#1062](https://github.com/openshift-pipelines/console-plugin/pull/1062) → [#1063](https://github.com/openshift-pipelines/console-plugin/pull/1063) | — | — |
| 51a05da | SRVKP-11951: Updating modules in master branch to address CVEs [release-v1.23... | OpenShift Cherrypick Robot | [#1064](https://github.com/openshift-pipelines/console-plugin/pull/1064) → [#1065](https://github.com/openshift-pipelines/console-plugin/pull/1065) | [SRVKP-11951](https://redhat.atlassian.net/browse/SRVKP-11951) | **Type:** Release Note Not Required<br>**Text:**  *(Jira)* |
| 8baa026 | fix: redirection logic fix to resolver tasks | Anwesha Palit | [#1076](https://github.com/openshift-pipelines/console-plugin/pull/1076) → [#1079](https://github.com/openshift-pipelines/console-plugin/pull/1079) | — | — |
| 70d7102 | added flag for ApprovalTask tab display in PipelineRunDetails page | Arvind | [#1082](https://github.com/openshift-pipelines/console-plugin/pull/1082) → [#1086](https://github.com/openshift-pipelines/console-plugin/pull/1086) | — | — |
| 076a37a | Merge pull request #1079 from openshift-cherrypick-robot/cherry-pick-1076-to-... | openshift-merge-bot[bot] | [#1076](https://github.com/openshift-pipelines/console-plugin/pull/1076) → [#1079](https://github.com/openshift-pipelines/console-plugin/pull/1079) | — | — |
| c39fced | Merge pull request #1086 from openshift-cherrypick-robot/cherry-pick-1082-to-... | Arvind Krishnamurthy | [#1082](https://github.com/openshift-pipelines/console-plugin/pull/1082) → [#1086](https://github.com/openshift-pipelines/console-plugin/pull/1086) | — | — |
| 62664eb | Add release branch pattern to workflow triggers (#1111) | Anwesha Palit | [#1111](https://github.com/openshift-pipelines/console-plugin/pull/1111) | — | — |
| 83539f4 | Update OWNERS to remove certain approvers and reviewers (#1113) | Anwesha Palit | [#1113](https://github.com/openshift-pipelines/console-plugin/pull/1113) | — | — |
| f9699c5 | feat: add CVE fix skill with dependency analysis script | Ankur Sinha | [#1136](https://github.com/openshift-pipelines/console-plugin/pull/1136) → [#1140](https://github.com/openshift-pipelines/console-plugin/pull/1140) | [SRVKP-12265](https://redhat.atlassian.net/browse/SRVKP-12265) | — |
| 22161b2 | Merge pull request #1140 from openshift-cherrypick-robot/cherry-pick-1136-to-... | openshift-merge-bot[bot] | [#1136](https://github.com/openshift-pipelines/console-plugin/pull/1136) → [#1140](https://github.com/openshift-pipelines/console-plugin/pull/1140) | [SRVKP-12265](https://redhat.atlassian.net/browse/SRVKP-12265) | — |

### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-9448](https://redhat.atlassian.net/browse/SRVKP-9448) | [UI] Current status overlaps in Approvals list view | Dev Complete |

</details>

<details>
<summary><h2>operator</h2> — 11 commits, 2 linked, 8 dependabot</summary>

**Upstream:** tektoncd/operator · **Commits:** 11 · **Unsynced:** 0 · **Dependabot:** 8

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 774d04b | fix(cve): update Go to 1.25.10 to address stdlib security vulnerabilities | Jawed Khelil | [#3411](https://github.com/tektoncd/operator/pull/3411) | [SRVKP-11355](https://redhat.atlassian.net/browse/SRVKP-11355) *(PR body)* | **Type:** CVE<br>**Text:** Updates Go from 1.25.0 to 1.25.10 to address 10 stdlib security vulnerabilities including fixes for crypto/x509, net/http, and mime packages. *(generated)* |
| fefa49e | fix(cve): CVE-2026-33814, CVE-2026-27141 - bump golang.org/x/net to v0.53.0 | Jawed Khelil | [#3436](https://github.com/tektoncd/operator/pull/3436) | — | — |
| b095a0f | update hub version to v1.23.10 | Shiv Verma | [#3457](https://github.com/tektoncd/operator/pull/3457) | — | — |
| ab348d7 | chore(deps): update Openshift images | Andrew Thorp | [#3458](https://github.com/tektoncd/operator/pull/3458) | — | — |
| 6fd20a3 | update the hub version to v1.23.11 | Shiv Verma | [#3465](https://github.com/tektoncd/operator/pull/3465) | — | — |
| 9d408fc | Feat: Console Plugin Image should be picked conditionally based OCP Version | Pramod Bindal | [#3386](https://github.com/tektoncd/operator/pull/3386) → [#3467](https://github.com/tektoncd/operator/pull/3467) | [SRVKP-11927](https://redhat.atlassian.net/browse/SRVKP-11927) | **Type:** Enhancement<br>**Text:** Added support for conditionally picking the console-plugin image on openshift.
OCP version older than 4.22 pick the legacy console plugin
Newer version pick the last console plugin image *(extracted)* |
| 48b496c | fix(cve): Update Go 1.25.10 → 1.25.11 to fix stdlib CVEs | Jawed Khelil | [#3470](https://github.com/tektoncd/operator/pull/3470) | — | — |
| d2771d5 | chore(deps): update openshift images sha | Shiv Verma | [#3515](https://github.com/tektoncd/operator/pull/3515) | — | — |
| 3fe2d61 | bump(deps): bump PAC, Pipeline and Hub version | Shiv Verma | [#3513](https://github.com/tektoncd/operator/pull/3513) | — | — |
| 8ef8dd0 | refactor: simplify update-image-sha.sh script | Shiv Verma | [#3516](https://github.com/tektoncd/operator/pull/3516) → [#3586](https://github.com/tektoncd/operator/pull/3586) | — | — |
| 84c7e7f | ci(e2e): fix ko >=v0.19 SBOM push failure on plain-HTTP registry | Jawed khelil | [#3621](https://github.com/tektoncd/operator/pull/3621) → [#3642](https://github.com/tektoncd/operator/pull/3642) | — | — |

</details>

<details>
<summary><h2>manual-approval-gate</h2> — 3 commits, 2 linked, 0 dependabot</summary>

**Upstream:** openshift-pipelines/manual-approval-gate · **Commits:** 3 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 2399dc6 | [Release-v0.7.0] Fix CVE-2026-33211: bump tektoncd/pipeline to v1.0.2 | infernus01 | [#980](https://github.com/openshift-pipelines/manual-approval-gate/pull/980) | — | — |
| 2f91dd6 | fix(cve): CVE-2026-33186 - google.golang.org/grpc (release-v0.7.0) | Divyanshu Agrawal | [#988](https://github.com/openshift-pipelines/manual-approval-gate/pull/988) | [SRVKP-11741](https://redhat.atlassian.net/browse/SRVKP-11741) *(PR body)* | **Type:** CVE<br>**Text:** Fixes CVE-2026-33186 (High severity) by updating google.golang.org/grpc to v1.79.3, resolving an authorization bypass vulnerability where a missing leading slash in the :path header could allow attackers to bypass gRPC server authorization checks. *(generated)* |
| f9ccb26 | fix(cve): CVE-2026-39821 + CVE-2026-33814 - Upgrade Go 1.24.0 → 1.25.11, x/ne... | Divyanshu Agrawal | [#1000](https://github.com/openshift-pipelines/manual-approval-gate/pull/1000) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **Type:** CVE<br>**Text:** Fixes CVE-2026-39821 and CVE-2026-33814 by upgrading golang.org/x/net to v0.55.0 and Go to 1.25.11. This resolves an ASCII-only Punycode label rejection bypass in x/net/idna and an HTTP/2 infinite loop caused by bad SETTINGS_MAX_FRAME_SIZE values. *(generated)* |

</details>

<details>
<summary><h2>pipelines-as-code</h2> — 16 commits, 2 linked, 0 dependabot</summary>

**Upstream:** tektoncd/pipelines-as-code · **Commits:** 16 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 6d5125b | chore(deps): update grpc to v1.79.3 | Akshay Pant | [#2730](https://github.com/tektoncd/pipelines-as-code/pull/2730) | — | — |
| 652421a | fix(deps): update go-jose to fix GHSA-78h2-9frx-2jm8 | Akshay Pant | [#2730](https://github.com/tektoncd/pipelines-as-code/pull/2730) | — | — |
| 68d2f42 | perf(informer): add TransformFuncs to reduce cache memory usage | Akshay Pant | [#2730](https://github.com/tektoncd/pipelines-as-code/pull/2730) | [SRVKP-11061](https://redhat.atlassian.net/browse/SRVKP-11061) | — |
| 632c9a5 | chore(deps): bump tektoncd/pipeline to v1.6.2 | Akshay Pant | [#2730](https://github.com/tektoncd/pipelines-as-code/pull/2730) | — | — |
| 95ae7a8 | fix(reconciler): skip watcher status updates | Chmouel Boudjnah | [#2730](https://github.com/tektoncd/pipelines-as-code/pull/2730) | — | — |
| d1f9603 | fix(ci): skip TLS verification for gosmee client in e2e tests | Zaki Shaikh | — | — | — |
| e0c4a11 | fix(security): backport app token safeguards | Chmouel Boudjnah | [#2762](https://github.com/tektoncd/pipelines-as-code/pull/2762) | [SRVKP-12241](https://redhat.atlassian.net/browse/SRVKP-12241), [SRVKP-12207](https://redhat.atlassian.net/browse/SRVKP-12207)| **Type:** CVE<br>**Text:** GitHub App installation tokens are now scoped to the triggering repository, preventing unauthorized access to private repositories via remote task resolution. *(generated)* |
| 591a50f | build(lint): pin golangci toolchain | Chmouel Boudjnah | [#2762](https://github.com/tektoncd/pipelines-as-code/pull/2762) | — | — |
| 64673d3 | chore(deps): bump go-jose v4 to v4.1.4 | Chmouel Boudjnah | [#2762](https://github.com/tektoncd/pipelines-as-code/pull/2762) | — | — |
| c85286e | fix(security): redact query string from incoming webhook log | Shubham Bhardwaj | [#2762](https://github.com/tektoncd/pipelines-as-code/pull/2762) | — | — |
| 402ae2c | ci: rewrite e2e script for main's matrix | Akshay Pant | [#2762](https://github.com/tektoncd/pipelines-as-code/pull/2762) | — | — |
| f5e54cd | fix(gitea): add nil-safety and Forgejo compat | Akshay Pant | [#2762](https://github.com/tektoncd/pipelines-as-code/pull/2762) | — | — |
| af48fc8 | test: rename bitbucket DC env vars to match main | Akshay Pant | [#2762](https://github.com/tektoncd/pipelines-as-code/pull/2762) | — | — |
| 27f3ec4 | Release yaml generated from https://github.com/tektoncd/pipelines-as-code/com... | Pipelines as Code CI Robot | — | — | — |
| aec194e | fix(webhook): prevent duplicate repository CR on trailing slash | Zaki Shaikh | [#2772](https://github.com/tektoncd/pipelines-as-code/pull/2772) | [SRVKP-11295](https://redhat.atlassian.net/browse/SRVKP-11295) *(keyword)* | **Type:** Bug Fix<br>**Text:** Pipelines as Code webhook no longer creates duplicate Repository custom resources when the repository URL contains a trailing slash. *(generated)* |
| 6acde9d | Release yaml generated from https://github.com/tektoncd/pipelines-as-code/com... | Pipelines as Code CI Robot | — | — | — |

</details>

<details>
<summary><h2>tektoncd-cli</h2> — 7 commits, 0 linked, 0 dependabot</summary>

**Upstream:** tektoncd/cli · **Commits:** 7 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| a9857c4 | Add CI summary job to the github workflows | divyansh42 | [#2912](https://github.com/tektoncd/cli/pull/2912) | — | — |
| bf5ea18 | Update grpc to v1.80.0 to fix CVE-2026-33186 | divyansh42 | [#2912](https://github.com/tektoncd/cli/pull/2912) | — | — |
| e931087 | Bump pipelines to 1.6.1 to fix CVE CVE-2026-33211 | divyansh42 | [#2912](https://github.com/tektoncd/cli/pull/2912) | — | — |
| 14893d7 | fix(cve): CVE-2026-34986 - go-jose/go-jose/v4 | Divyanshu Agrawal | [#2912](https://github.com/tektoncd/cli/pull/2912) | — | — |
| cf1e983 | fix(cve): CVE-2026-24051 - bump go.opentelemetry.io/otel to v1.40.0 [release-... | divyansh42 | [#2912](https://github.com/tektoncd/cli/pull/2912) | — | — |
| ca98f6a | fix(cve): CVE-2026-40938, CVE-2026-40161 - github.com/tektoncd/pipeline | divyansh42 | [#2912](https://github.com/tektoncd/cli/pull/2912) | — | — |
| db2b9a1 | New version v0.43.2 | Shiv Verma | [#2912](https://github.com/tektoncd/cli/pull/2912) | — | — |

</details>

<details>
<summary><h2>console-plugin-pf5</h2> — new component — not present in old snapshot</summary>

**Error:** new component — not present in old snapshot

</details>

<details>
<summary><h2>tektoncd-triggers</h2> — 0 commits, 0 linked, 0 dependabot</summary>

**Upstream:** tektoncd/triggers · **Commits:** 0 · **Unsynced:** 0 · **Dependabot:** 0

</details>

<details>
<summary><h2>git-init</h2> — 56 commits, 0 linked, 21 dependabot</summary>

**Upstream:** tektoncd-catalog/git-clone · **Commits:** 56 · **Unsynced:** 0 · **Dependabot:** 21

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| feed9a5 | Add user-friendly error messages for git-clone failures (enabled by default) | Chmouel Boudjnah | [#104](https://github.com/tektoncd-catalog/git-clone/pull/104) | — | — |
| 61d5b52 | fix: golangci-lint fixes | Chmouel Boudjnah | [#104](https://github.com/tektoncd-catalog/git-clone/pull/104) | — | — |
| 75241a8 | docs: rewrite README with proper project documentation | Vincent Demeester | [#106](https://github.com/tektoncd-catalog/git-clone/pull/106) | — | — |
| 5954a23 | chore: update Go to 1.25 and tektoncd/pipeline to v1.12.0 LTS | Vincent Demeester | [#107](https://github.com/tektoncd-catalog/git-clone/pull/107) | — | — |
| afe2eb7 | chore: drop tektoncd/pipeline dependency | Vincent Demeester | [#108](https://github.com/tektoncd-catalog/git-clone/pull/108) | — | — |
| f12396a | ci: fix release workflow failures | Vincent Demeester | [#109](https://github.com/tektoncd-catalog/git-clone/pull/109) | — | — |
| aad26bf | chore: bump version to v1.3.0 | Vincent Demeester | — | — | — |
| f7bbf68 | fix: replace kodata LICENSE symlink with copy | Vincent Demeester | — | — | — |
| f8e4874 | chore: add Artifact Hub metadata and catalog annotations | Vincent Demeester | — | — | — |
| a35d224 | chore: set Artifact Hub repository ID | Vincent Demeester | — | — | — |
| 6e3f196 | docs: add Artifact Hub badge to README | Vincent Demeester | — | — | — |
| 520317d | chore: add Artifact Hub annotations for richer package listing | Vincent Demeester | — | — | — |
| 2d927c2 | fix: update task API version from v1beta1 to v1 | Vincent Demeester | — | — | — |
| 6a00bdd | chore: bump version to v1.4.0 | Vincent Demeester | — | — | — |
| 81189e7 | chore: bump minVersion, add PR template, fix CI tests | Vincent Demeester | [#110](https://github.com/tektoncd-catalog/git-clone/pull/110) | — | — |
| 0ffda8e | fix: use chainguard/git base image for releases | Vincent Demeester | [#112](https://github.com/tektoncd-catalog/git-clone/pull/112) | — | — |
| 54d4b67 | ci: add e2e tests with Kind cluster | Vincent Demeester | [#111](https://github.com/tektoncd-catalog/git-clone/pull/111) | — | — |
| e35ff2c | feat: publish Tekton Bundle on release | Vincent Demeester | [#115](https://github.com/tektoncd-catalog/git-clone/pull/115) | — | — |
| b5e1f95 | feat: add custom base image with git-lfs, openssh, uid 65532 | Vincent Demeester | [#116](https://github.com/tektoncd-catalog/git-clone/pull/116) | — | — |
| 768795f | chore: add release script | Vincent Demeester | [#120](https://github.com/tektoncd-catalog/git-clone/pull/120) | — | — |
| 4011f7e | chore: bump version to v1.5.0 | Vincent Demeester | — | — | — |
| 7d3edaf | chore: add task signing public key | Vincent Demeester | — | — | — |
| 2e43d8a | feat: add git-clone StepAction, generated from Task | Vincent Demeester | [#121](https://github.com/tektoncd-catalog/git-clone/pull/121) | — | — |
| bf7d27e | docs: update READMEs for Artifact Hub and StepAction | Vincent Demeester | [#121](https://github.com/tektoncd-catalog/git-clone/pull/121) | — | — |
| 8be4f21 | feat: stepaction tests, bundle publishing, release integration | Vincent Demeester | [#121](https://github.com/tektoncd-catalog/git-clone/pull/121) | — | — |
| 72fde74 | docs: add parameters and results tables to StepAction README | Vincent Demeester | [#121](https://github.com/tektoncd-catalog/git-clone/pull/121) | — | — |
| cea7dec | chore: add task signing to release script | Vincent Demeester | — | — | — |
| 97ba0b1 | fix: release script backtick stripping in changelog normalization | Vincent Demeester | — | — | — |
| 1a89de8 | chore: bump version to v1.6.0 | Vincent Demeester | — | — | — |
| cf24a58 | chore: add Artifact Hub metadata for StepAction repository | Vincent Demeester | — | — | — |
| de06b44 | chore: set Artifact Hub repository ID for StepAction | Vincent Demeester | — | — | — |
| b756d38 | fix: strip spurious resources: {} from signed task YAML | Vincent Demeester | — | — | — |
| e9cb6de | fix: strip spurious resources: {} from signed task YAMLs | Vincent Demeester | — | — | — |
| e082053 | fix: release script version detection and signing | Vincent Demeester | — | — | — |
| d8879cc | fix: release script signing and bundle workflow | Vincent Demeester | — | — | — |
| bd41a2c | chore: bump version to v1.6.1 | Vincent Demeester | — | — | — |
| 744f8ad | fix: rewrite StepAction generator with pyyaml | Vincent Demeester | — | — | — |
| d20cd45 | fix: reset version to v1.5.0 for v1.6.0 re-release | Vincent Demeester | — | — | — |
| 59b4802 | chore: bump version to v1.6.0 | Vincent Demeester | — | — | — |
| 54d6e0b | fix: quote special characters in artifacthub changelog annotation | Vincent Demeester | — | — | — |
| 0caea16 | fix: reset version for v1.6.0 re-release (pyyaml generator) | Vincent Demeester | — | — | — |
| dd1d541 | chore: bump version to v1.6.0 | Vincent Demeester | — | — | — |
| 9c96850 | fix: sanitize special YAML chars in changelog descriptions | Vincent Demeester | — | — | — |
| b5b6070 | chore: bump version to v1.6.0 | Vincent Demeester | — | — | — |
| 303b9f6 | fix: quote descriptions containing colons in changelog | Vincent Demeester | — | — | — |
| 4be0731 | chore: bump version to v1.6.0 | Vincent Demeester | — | — | — |
| 59623d0 | fix: always quote descriptions in AH changelog | Vincent Demeester | — | — | — |
| 8b26e19 | chore: bump version to v1.6.0 | Vincent Demeester | — | — | — |
| 05651c4 | ci: build base image for multi-architecture | Vincent Demeester | — | — | — |
| c166e56 | chore: bump version to v1.7.0 | Vincent Demeester | — | — | — |
| 8f19a96 | chore: drop Trusted Resources signature from release | Vincent Demeester | — | — | — |
| 577fd3d | docs: add DEVELOPMENT.md, CONTRIBUTING.md, and AGENTS.md | Vincent Demeester | [#126](https://github.com/tektoncd-catalog/git-clone/pull/126) | — | — |
| d364aa7 | docs: address review feedback | Vincent Demeester | [#126](https://github.com/tektoncd-catalog/git-clone/pull/126) | — | — |
| a962c33 | ci: make e2e tests more robust against flakes | Vincent Demeester | [#127](https://github.com/tektoncd-catalog/git-clone/pull/127) | — | — |
| e03aee4 | docs: add Artifact Hub StepActions badge and browse links | Vincent Demeester | — | — | — |
| 27c670b | chore: group dependabot updates and fix gomod path | Vincent Demeester | [#137](https://github.com/tektoncd-catalog/git-clone/pull/137) | — | — |

</details>

<details>
<summary><h2>tekton-caches</h2> — 9 commits, 5 linked, 0 dependabot</summary>

**Upstream:** openshift-pipelines/tekton-caches · **Commits:** 9 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 80d513d | [bot:release-v0.3.x] update konflux configuration | openshift-pipelines-bot | [#708](https://github.com/openshift-pipelines/tekton-caches/pull/708) | — | — |
| 6f8e051 | Fix CVE: CVE-2026-33186 openshift-pipelines/pipelines-cache-rhel9: gRPC-Go: A... | Pramod Bindal | [#725](https://github.com/openshift-pipelines/tekton-caches/pull/725) | [SRVKP-12040](https://redhat.atlassian.net/browse/SRVKP-12040) *(keyword)* | **Type:** CVE<br>**Text:** Fixes CVE-2026-33186, a gRPC-Go vulnerability that allows authorization bypass due to improper HTTP/2 path validation in the Tekton Caches component. *(generated)* |
| 27ef971 | Cleanup konflux | Pramod Bindal | [#725](https://github.com/openshift-pipelines/tekton-caches/pull/725) | [SRVKP-12040](https://redhat.atlassian.net/browse/SRVKP-12040) *(keyword)* | **Type:** Release Notes Not Required<br>**Text:** Internal Konflux build configuration cleanup. *(generated)* |
| 0d3e9df | Upgrade Pipelines | Pramod Bindal | [#725](https://github.com/openshift-pipelines/tekton-caches/pull/725) | [SRVKP-12040](https://redhat.atlassian.net/browse/SRVKP-12040) *(keyword)* | **Type:** Release Notes Not Required<br>**Text:** Internal pipeline dependency upgrade as part of CVE remediation. *(generated)* |
| 208a38a | Fix lint issues | Pramod Bindal | [#725](https://github.com/openshift-pipelines/tekton-caches/pull/725) | [SRVKP-12040](https://redhat.atlassian.net/browse/SRVKP-12040) *(keyword)* | **Type:** Release Notes Not Required<br>**Text:** Internal lint issue fixes. *(generated)* |
| 2450f24 | Security: Fix CVE-2026-34986 (go-jose/go-jose/v4) | Pramod Bindal | [#729](https://github.com/openshift-pipelines/tekton-caches/pull/729) | — | — |
| 2a64502 | bump otel and pipeline version for CVE fixes | Pramod Bindal | [#736](https://github.com/openshift-pipelines/tekton-caches/pull/736) | — | — |
| 5d6871a | cve: Fix GHSA-xmrv-pmrh-hhx2 | Pramod Bindal | [#739](https://github.com/openshift-pipelines/tekton-caches/pull/739) | — | — |
| 768eb9a | Add Agentic Contexts and CI Workflows | Pramod Bindal | [#739](https://github.com/openshift-pipelines/tekton-caches/pull/739) | [SRVKP-11823](https://redhat.atlassian.net/browse/SRVKP-11823) *(keyword)* | **Type:** Release Notes Not Required<br>**Text:** Adds agentic context files and CI workflows to the tekton-caches repository for improved developer tooling. *(generated)* |

</details>

<details>
<summary><h2>tektoncd-hub</h2> — 9 commits, 1 linked, 0 dependabot</summary>

**Upstream:** openshift-pipelines/hub · **Commits:** 9 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 74023f0 | bump google.golang.org/grpc to version 1.79.3 | Shiv Verma | — | [SRVKP-12037](https://redhat.atlassian.net/browse/SRVKP-12037) *(keyword)* | **Type:** CVE<br>**Text:** Updates the gRPC-Go dependency to version 1.79.3 to fix an authorization bypass vulnerability caused by a missing leading slash in the :path header. *(generated)* |
| eb086e1 | bump tektoncd/pipeline version 1.9.2 | Shiv Verma | — | — | — |
| c79258f | bump UI and swagger packages to fix the CVEs | Shiv Verma | — | — | — |
| 382b33a | bump tektoncd/pipeline version 1.9.3 to fix CVE-2026-40938 | Shiv Verma | — | — | — |
| 4f994d2 | bump  github.com/jackc/pgx/v5 to version 5.9.0 to fix CVE-2026-33816 | Shiv Verma | — | — | — |
| 1aa9735 | bump golang.org/x/net and golang.org/x/crypto package | Shiv Verma | — | — | — |
| 2a203c0 | bump github.com/go-jose/go-jose to v4.1.4 | Shiv Verma | — | — | — |
| f04ddb0 | update go version to 1.25.10 | Shiv Verma | — | — | — |
| f7d1f0b | update the api and db dockerfile | Shiv Verma | — | — | — |

</details>

<details>
<summary><h2>tektoncd-results</h2> — 10 commits, 2 linked, 0 dependabot</summary>

**Upstream:** tektoncd/results · **Commits:** 10 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| bd87be2 | Bump google.golang.org/grpc to v1.80.0 | Emil Natan | [#1308](https://github.com/tektoncd/results/pull/1308) | — | — |
| 00afe86 | Merge pull request #1308 from enarha/update-release-v0.17.x | Emil Natan | [#1308](https://github.com/tektoncd/results/pull/1308) | — | — |
| 4f55175 | Update tektoncd/pipeline and go-jose modules | Emil Natan | [#1323](https://github.com/tektoncd/results/pull/1323) | — | — |
| d9a8596 | Merge pull request #1323 from enarha/update-pipeline-release-v0.17.x | Emil Natan | [#1323](https://github.com/tektoncd/results/pull/1323) | — | — |
| 508cc32 | [v0.17.x] Fix CVE issues | Emil Natan | [#1346](https://github.com/tektoncd/results/pull/1346) | — | — |
| 56d6a87 | Bump golangci-lint version to support Go 1.25 | Emil Natan | [#1346](https://github.com/tektoncd/results/pull/1346) | — | — |
| 306ee40 | Fix race condition in the TestController | Emil Natan | [#1346](https://github.com/tektoncd/results/pull/1346) | — | — |
| 4afeda9 | Merge pull request #1346 from enarha/update-1.21-cves-release-v0.17.x | Emil Natan | [#1346](https://github.com/tektoncd/results/pull/1346) | — | — |
| 8cda1d8 | perf(cli): use exact match for describe/logs command | divyansh42 | [#1283](https://github.com/tektoncd/results/pull/1283) → [#1336](https://github.com/tektoncd/results/pull/1336) | [SRVKP-11380](https://redhat.atlassian.net/browse/SRVKP-11380) | **Type:** Enhancement<br>**Text:** Enhacement: Improved performance for `describe` and `logs` commands with optimized filtering; when multiple runs share the same name, the most recent one is returned (use `--uid` to  target a specific run)  *(extracted)* |
| ae090bc | Merge pull request #1336 from divyansh42/cherrypick-11380-0.17 | Emil Natan | [#1283](https://github.com/tektoncd/results/pull/1283) → [#1336](https://github.com/tektoncd/results/pull/1336) | [SRVKP-11380](https://redhat.atlassian.net/browse/SRVKP-11380) | **Type:** Enhancement<br>**Text:** Enhacement: Improved performance for `describe` and `logs` commands with optimized filtering; when multiple runs share the same name, the most recent one is returned (use `--uid` to  target a specific run)  *(extracted)* |

</details>

<details>
<summary><h2>tektoncd-pruner</h2> — 3 commits, 0 linked, 0 dependabot</summary>

**Upstream:** tektoncd/pruner · **Commits:** 3 · **Unsynced:** 0 · **Dependabot:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 18903cd | (CVE-2026-40938) PR to fix update pipelines to v1.6.2 | Naomi Gelman | [#267](https://github.com/tektoncd/pruner/pull/267) | — | — |
| a32f9cb | fix(security): patch three OpenTelemetry Go CVEs on release-v0.3.x | Shubham Bhardwaj | [#310](https://github.com/tektoncd/pruner/pull/310) | — | — |
| 0715ff1 | fix(cve): bump OTel SDK to v1.43.0 to fix GHSA-9h8m-3fm2-qjrq and GHSA-hfvc-g... | Shubham Bhardwaj | [#311](https://github.com/tektoncd/pruner/pull/311) | — | — |

</details>

<details>
<summary><h2>tektoncd-chains</h2> — 5 commits, 4 linked, 3 dependabot</summary>

**Upstream:** tektoncd/chains · **Commits:** 5 · **Unsynced:** 0 · **Dependabot:** 3

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 7d7d316 | fix ci workflow to include CI summary (#1633) | Anitha Natarajan | [#1633](https://github.com/tektoncd/chains/pull/1633) | — | — |
| 868c359 | fix: CVE-2026-34986, CVE-2026-33211, & CVE-2026-33186 (#1630) | Shubham Bhardwaj | [#1630](https://github.com/tektoncd/chains/pull/1630) | [SRVKP-12344](https://redhat.atlassian.net/browse/SRVKP-12344) *(keyword)* | **Type:** CVE<br>**Text:** Fixes CVE-2026-34986, CVE-2026-33211, and CVE-2026-33186 by updating vulnerable Go dependencies in the Chains controller. *(generated)* |
| a51457d | fix: update pipelines to v1.6.2 (#1649) | Anitha Natarajan | [#1649](https://github.com/tektoncd/chains/pull/1649) | [SRVKP-12344](https://redhat.atlassian.net/browse/SRVKP-12344) *(keyword)* | **Type:** CVE<br>**Text:** Fixes CVE-2026-40938 and CVE-2026-40161 by updating the Tekton Pipelines dependency to v1.6.2 in the Chains controller. *(generated)* |
| dbad7be | fix(cve): CVE-2026-33814 - golang.org/x/net HTTP/2 infinite loop (#1670) | khelil | [#1670](https://github.com/tektoncd/chains/pull/1670) | [SRVKP-12344](https://redhat.atlassian.net/browse/SRVKP-12344) *(keyword)* | **Type:** CVE<br>**Text:** Fixes CVE-2026-33814, a high-severity vulnerability in golang.org/x/net that could cause an infinite loop in HTTP/2 transport, leading to unbounded CPU consumption. *(generated)* |
| 881f5e9 | fix(deps): bump Go modules to address CVEs (#1703) | khelil | [#1703](https://github.com/tektoncd/chains/pull/1703) | [SRVKP-12344](https://redhat.atlassian.net/browse/SRVKP-12344) *(keyword)* | **Type:** CVE<br>**Text:** Remediates multiple critical and important CVEs by updating Go dependencies including Tekton Pipeline, golang.org/x/crypto, golang.org/x/net, OpenTelemetry, Sigstore, and other modules in the Chains controller. *(generated)* |

</details>

## Unmatched Jiras

| Key | Summary | Status | Components |
|-----|---------|--------|------------|
| [SRVKP-12589](https://redhat.atlassian.net/browse/SRVKP-12589) | CLI testing | Backlog | — |
| [SRVKP-12588](https://redhat.atlassian.net/browse/SRVKP-12588) | UI Regression testing | Backlog | — |
| [SRVKP-12587](https://redhat.atlassian.net/browse/SRVKP-12587) | Acceptance test | Backlog | QA |
| [SRVKP-12586](https://redhat.atlassian.net/browse/SRVKP-12586) | Upgrade test | Backlog | QA |
| [SRVKP-12585](https://redhat.atlassian.net/browse/SRVKP-12585) | Test Openshift Pipelines 1.21.3 | Backlog | QA |
| [SRVKP-12574](https://redhat.atlassian.net/browse/SRVKP-12574) | OSP 1.21.3 patch release | New | — |
| [SRVKP-12552](https://redhat.atlassian.net/browse/SRVKP-12552) | OSP 1.21.3 release configuration  | In Progress | — |

## Summary

| Component | Commits | Linked (tool) | Linked (audit) | Unlinked | Has RN | Generated RN | Dependabot |
|-----------|---------|---------------|----------------|----------|--------|--------------|------------|
| tektoncd-pipeline | 30 | 0 | 11 | 19 | 11 | 7 | 28 |
| console-plugin | 183 | 20 | 0 | 163 | 18 | 0 | 0 |
| operator | 11 | 1 | 1 | 9 | 2 | 1 | 8 |
| manual-approval-gate | 3 | 0 | 2 | 1 | 2 | 2 | 0 |
| pipelines-as-code | 16 | 0 | 2 | 14 | 2 | 2 | 0 |
| tektoncd-cli | 7 | 0 | 0 | 7 | 0 | 0 | 0 |
| console-plugin-pf5 | — | — | — | — | — | — | — |
| tektoncd-triggers | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| git-init | 56 | 0 | 0 | 56 | 0 | 0 | 21 |
| tekton-caches | 9 | 0 | 5 | 4 | 5 | 5 | 0 |
| tektoncd-hub | 9 | 0 | 1 | 8 | 1 | 1 | 0 |
| tektoncd-results | 10 | 2 | 0 | 8 | 2 | 0 | 0 |
| tektoncd-pruner | 3 | 0 | 0 | 3 | 0 | 0 | 0 |
| tektoncd-chains | 5 | 0 | 4 | 1 | 4 | 4 | 3 |
| **Total** | **342** | **23** | **26** | **293** | | | **60** |
