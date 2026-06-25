# Release Audit: 1.23.0

**Source:** `reports/release_1.23.0.json`<br>
**Total commits:** 1308 across 16 components<br>
**Already linked (by Go tool):** 145 · **Newly linked:** 117 · **Unlinked:** 395 · **Dependabot (skipped):** 651<br>

---

## git-init

**Error:** failed to get head SHA at from-date: failed to list head file commits for openshift-pipelines/git-init: GET https://api.github.com/repos/openshift-pipelines/git-init/commits?path=head&per_page=1&sha=release-v1.22.x&until=2026-05-15T07%3A57%3A06Z: 404 Not Found []

---

## opc

**Upstream:** openshift-pipelines/opc · **Commits:** 3 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| fd19a79 | Update opc version 1.23.0 | Khurram Baig | [#458](https://github.com/openshift-pipelines/opc/pull/458) | [SRVKP-12016](https://redhat.atlassian.net/browse/SRVKP-12016) *(keyword)* | **Release Notes Not Required:** Internal version bump to 1.23.0. *(generated)* |
| bc04cca | add replace for deprecated hub repo | Shiv Verma | [#462](https://github.com/openshift-pipelines/opc/pull/462) | [SRVKP-11950](https://redhat.atlassian.net/browse/SRVKP-11950) *(keyword)* | **Deprecated Functionality:** The opc CLI now uses the openshift-pipelines/hub fork in place of the deprecated *(generated)* |
| b9be146 | Bump to Pipelines, PAC, Triggers and Chains LTS | Khurram Baig | [#466](https://github.com/openshift-pipelines/opc/pull/466) | [SRVKP-12016](https://redhat.atlassian.net/browse/SRVKP-12016) *(keyword)* | **Release Notes Not Required:** Internal dependency update bumping Pipelines as Code to 0.48.0, Tekton CLI to 0. *(generated)* |

---

## console-plugin-pf5

**Upstream:** openshift-pipelines/console-plugin · **Commits:** 19 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| e282b00 | updaed hub cluster detection to use multi-cluster-role | Arvind | — | [SRVKP-10786](https://redhat.atlassian.net/browse/SRVKP-10786) *(PR body)* | **Enhancement:** Hub cluster detection now uses the TektonConfig scheduler configuration instead  *(generated)* |
| 6b19686 | Merge pull request #960 from openshift-cherrypick-robot/cher | anwesha-palit-redhat | — | [SRVKP-10786](https://redhat.atlassian.net/browse/SRVKP-10786) *(PR body)* | **Enhancement:** Hub cluster detection now uses the TektonConfig scheduler configuration instead  *(generated)* |
| e3c13ae | delete: removing .konflux and .tekton dir from upstream cons | Anwesha Palit | — | [SRVKP-10961](https://redhat.atlassian.net/browse/SRVKP-10961) *(PR body)* | **Release Notes Not Required:** Internal build configuration cleanup; no user-facing changes. *(generated)* |
| 82222fd | Merge pull request #991 from openshift-cherrypick-robot/cher | Anwesha Palit | — | [SRVKP-10961](https://redhat.atlassian.net/browse/SRVKP-10961) *(PR body)* | **Release Notes Not Required:** Internal build configuration cleanup; no user-facing changes. *(generated)* |
| 5b0316e | [release-v1.22.x] SRVKP-11533: included missing z-index for  | OpenShift Cherrypick Robot | — | [SRVKP-11533](https://redhat.atlassian.net/browse/SRVKP-11533) *(PR body)* | **Bug Fix:** Fixes a rendering issue where the Task sidebar could appear behind other UI elem *(generated)* |
| 5004f49 | SRVKP-11699: updated versions in main branch for CVEs [relea | OpenShift Cherrypick Robot | [#1071](https://github.com/openshift-pipelines/console-plugin/pull/1071) | [SRVKP-11699](https://redhat.atlassian.net/browse/SRVKP-11699) | **CVE:** Updates dependency versions to address known CVE vulnerabilities in the console  *(generated)* |
| f44dc45 | fix: redirection logic fix to resolver tasks | Anwesha Palit | [#1078](https://github.com/openshift-pipelines/console-plugin/pull/1078) | [SRVKP-12111](https://redhat.atlassian.net/browse/SRVKP-12111) *(PR body)* | **Bug Fix:** Fixes incorrect task links on the Pipeline details page for cluster resolver tas *(generated)* |
| f41b9b8 | added flag for ApprovalTask tab display in PipelineRunDetail | Arvind | [#1085](https://github.com/openshift-pipelines/console-plugin/pull/1085) | [SRVKP-11698](https://redhat.atlassian.net/browse/SRVKP-11698) *(PR body)* | **Bug Fix:** The Approvals tab in PipelineRun details is now conditionally displayed based on *(generated)* |
| ff73742 | Merge pull request #1078 from openshift-cherrypick-robot/che | openshift-merge-bot[bot] | [#1078](https://github.com/openshift-pipelines/console-plugin/pull/1078) | [SRVKP-12111](https://redhat.atlassian.net/browse/SRVKP-12111) *(PR body)* | **Bug Fix:** Fixes incorrect task links on the Pipeline details page for cluster resolver tas *(generated)* |
| 3b72aef | Merge pull request #1085 from openshift-cherrypick-robot/che | Arvind Krishnamurthy | [#1085](https://github.com/openshift-pipelines/console-plugin/pull/1085) | [SRVKP-11698](https://redhat.atlassian.net/browse/SRVKP-11698) *(PR body)* | **Bug Fix:** The Approvals tab in PipelineRun details is now conditionally displayed based on *(generated)* |
| e0bf688 | Updating PR for github workflow and also updating owner file | Anwesha Palit | [#1106](https://github.com/openshift-pipelines/console-plugin/pull/1106) | — | — |
| 54110e3 | fix: resolve protobufjs dependency via resolution | Ankur Sinha | [#1116](https://github.com/openshift-pipelines/console-plugin/pull/1116) | [SRVKP-11776](https://redhat.atlassian.net/browse/SRVKP-11776) *(PR body)* | **CVE:** Resolves CVE-2026-41242 in the protobufjs transitive dependency by enforcing a p *(generated)* |
| b91caaa | Remove position: fixed to solve overlapping current status i | Aditya Shinde | [#1120](https://github.com/openshift-pipelines/console-plugin/pull/1120) | [SRVKP-9448](https://redhat.atlassian.net/browse/SRVKP-9448) *(PR body)* | **Bug Fix:** The "Current Status" column in the Approvals list now renders correctly instead  *(extracted)* |
| f643a3e | Merge pull request #1120 from openshift-cherrypick-robot/che | openshift-merge-bot[bot] | [#1120](https://github.com/openshift-pipelines/console-plugin/pull/1120) | [SRVKP-9448](https://redhat.atlassian.net/browse/SRVKP-9448) *(PR body)* | **Bug Fix:** The "Current Status" column in the Approvals list now renders correctly instead  *(extracted)* |
| caf42cc | Merge pull request #1116 from ankrsinha/release-v1.22.x | openshift-merge-bot[bot] | [#1116](https://github.com/openshift-pipelines/console-plugin/pull/1116) | [SRVKP-11776](https://redhat.atlassian.net/browse/SRVKP-11776) *(PR body)* | **CVE:** Resolves CVE-2026-41242 in the protobufjs transitive dependency by enforcing a p *(generated)* |
| b6e3e2f | trigger rebuild for rebuilding base image for CVE-2026-45186 | Anwesha Palit | [#1128](https://github.com/openshift-pipelines/console-plugin/pull/1128) | [SRVKP-12362](https://redhat.atlassian.net/browse/SRVKP-12362) *(PR body)* | **CVE:** Triggers a base image rebuild to address CVE-2026-45186. *(generated)* |
| 26c1e45 | Merge pull request #1128 from anwesha-palit-redhat/release-v | Arvind Krishnamurthy | [#1128](https://github.com/openshift-pipelines/console-plugin/pull/1128) | [SRVKP-12362](https://redhat.atlassian.net/browse/SRVKP-12362) *(PR body)* | **CVE:** Triggers a base image rebuild to address CVE-2026-45186. *(generated)* |
| 59c8128 | fix mag toast redirection for non admin | Arvind | [#1131](https://github.com/openshift-pipelines/console-plugin/pull/1131) | [SRVKP-9451](https://redhat.atlassian.net/browse/SRVKP-9451) | **Known Issue:** Fixes the Manual Approval Gate toast notification redirection for non-admin user *(generated)* |
| 17a5fc8 | Merge pull request #1131 from openshift-cherrypick-robot/che | openshift-merge-bot[bot] | [#1131](https://github.com/openshift-pipelines/console-plugin/pull/1131) | [SRVKP-9451](https://redhat.atlassian.net/browse/SRVKP-9451) | **Known Issue:** Fixes the Manual Approval Gate toast notification redirection for non-admin user *(generated)* |

---

## tekton-kueue

**Upstream:** konflux-ci/tekton-kueue · **Commits:** 2 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 096a0e8 | fix(deps): update kubernetes packages to v0.35.5 | red-hat-konflux[bot] | [#430](https://github.com/konflux-ci/tekton-kueue/pull/430) | — | — |
| ec8c653 | cve: Fixes following CVEs (#466) | Pramod Bindal | [#430](https://github.com/konflux-ci/tekton-kueue/pull/430) | [SRVKP-11951](https://redhat.atlassian.net/browse/SRVKP-11951) *(unmatched)* | — |

---

## pipelines-as-code

**Upstream:** tektoncd/pipelines-as-code · **Commits:** 230 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| bfcbb1b | chore: re-enable tests for Bitbucket Data Center | Chmouel Boudjnah | [#2497](https://github.com/tektoncd/pipelines-as-code/pull/2497) | [SRVKP-9638](https://redhat.atlassian.net/browse/SRVKP-9638) *(keyword)* | **Release Notes Not Required:** Internal: re-enable tests for Bitbucket Data Center *(generated)* |
| 561d5e9 | feat: Move more tests to GHE | Chmouel Boudjnah | [#2497](https://github.com/tektoncd/pipelines-as-code/pull/2497) | — | — |
| 41c6ace | refactor(providers): constantize provider StatusOpt conclusi | Andrew Thorp | [#2384](https://github.com/tektoncd/pipelines-as-code/pull/2384) | — | — |
| efe2e0e | fix: Correct type for status options | Chmouel Boudjnah | [#2499](https://github.com/tektoncd/pipelines-as-code/pull/2499) | — | — |
| 4d950e7 | fix(providers): restrict comment editing to PAC-owned commen | Akshay Pant | [#2488](https://github.com/tektoncd/pipelines-as-code/pull/2488) | [SRVKP-10857](https://redhat.atlassian.net/browse/SRVKP-10857) | **Bug Fix:** Pipelines as Code now restricts comment editing to only comments owned by PAC, p *(generated)* |
| 106b201 | feat: Allow custom user agent for Forgejo/Gitea | Chmouel Boudjnah | [#2494](https://github.com/tektoncd/pipelines-as-code/pull/2494) | [SRVKP-10579](https://redhat.atlassian.net/browse/SRVKP-10579) | **Feature:** Forgejo and Gitea providers now support configuring a custom User-Agent string f *(generated)* |
| fc4ee2f | feat: Add task to generate AI release notes | Chmouel Boudjnah | [#2500](https://github.com/tektoncd/pipelines-as-code/pull/2500) | — | — |
| 1b043c5 | fix: set PipelineURL for cached pipelines to resolve relativ | Shubham Mathur | [#2416](https://github.com/tektoncd/pipelines-as-code/pull/2416) | [SRVKP-10604](https://redhat.atlassian.net/browse/SRVKP-10604) | **Bug Fix:** Fixes relative task path resolution for cached pipelines by correctly setting th *(generated)* |
| c2c43a9 | test(github): add DirectWebhook support in CRD helper | Akshay Pant | [#2506](https://github.com/tektoncd/pipelines-as-code/pull/2506) | — | — |
| 59a480c | fix(labels): use positive check for InstallationID | Akshay Pant | [#2506](https://github.com/tektoncd/pipelines-as-code/pull/2506) | — | — |
| f97499b | ci: Pin GitHub Actions to commit SHAs | Vincent Demeester | [#2507](https://github.com/tektoncd/pipelines-as-code/pull/2507) | — | — |
| a227d9f | chore: Add second GitHub application ID to e2e tests | Chmouel Boudjnah | [#2508](https://github.com/tektoncd/pipelines-as-code/pull/2508) | — | — |
| e48931a | fix(gitops-commands): parse branch vs tag by 'tag:' | Shubham Mathur | [#2505](https://github.com/tektoncd/pipelines-as-code/pull/2505) | [SRVKP-10915](https://redhat.atlassian.net/browse/SRVKP-10915) | **Bug Fix:** Gitops commands now correctly distinguish between branch and tag references usin *(generated)* |
| e5302e0 | feat: Add comment strategy support for Foregejo | Chmouel Boudjnah | [#2503](https://github.com/tektoncd/pipelines-as-code/pull/2503) | [SRVKP-10900](https://redhat.atlassian.net/browse/SRVKP-10900) | **Feature:** Comment strategy support is now available for the Forgejo provider, allowing con *(generated)* |
| ec55549 | docs: update bitbucket cloud docs for api tokens | Zaki Shaikh | [#2501](https://github.com/tektoncd/pipelines-as-code/pull/2501) | [SRVKP-10621](https://redhat.atlassian.net/browse/SRVKP-10621) *(unmatched)* | — |
| 8b14fcd | feat: Move back to startpaac for E2E testing | Chmouel Boudjnah | [#2509](https://github.com/tektoncd/pipelines-as-code/pull/2509) | — | — |
| 3015674 | fix(gitea): preserve source_url on retest comment reruns | Chmouel Boudjnah | [#2502](https://github.com/tektoncd/pipelines-as-code/pull/2502) | [SRVKP-10575](https://redhat.atlassian.net/browse/SRVKP-10575) | **Bug Fix:** Fixes source_url preservation on Gitea/Forgejo when rerunning pipelines via rete *(generated)* |
| 435b6e4 | fix: Avoid webhook feedback loop on no-ops comment events (# | Chmouel Boudjnah | [#2504](https://github.com/tektoncd/pipelines-as-code/pull/2504) | [SRVKP-10912](https://redhat.atlassian.net/browse/SRVKP-10912) | **Bug Fix:** Prevents webhook feedback loops caused by no-op comment events that could trigge *(generated)* |
| 70d1ff1 | fix(e2e): add SinceSeconds support to pod log helpers | Zaki Shaikh | [#2515](https://github.com/tektoncd/pipelines-as-code/pull/2515) | — | — |
| d4b5b97 | chore: increase sinceSeconds in Git Clone e2e test | Zaki Shaikh | [#2516](https://github.com/tektoncd/pipelines-as-code/pull/2516) | — | — |
| c057f26 | test: split github_ghe job into 3 parallel chunks | Zaki Shaikh | [#2517](https://github.com/tektoncd/pipelines-as-code/pull/2517) | — | — |
| bd1176c | chore: remove nightly from TestGithubGHEPullRequestGitCloneT | Zaki Shaikh | [#2518](https://github.com/tektoncd/pipelines-as-code/pull/2518) | — | — |
| 8a2e0e4 | chore: remove GitHub comment dedup workaround | Chmouel Boudjnah | [#2511](https://github.com/tektoncd/pipelines-as-code/pull/2511) | [SRVKP-10938](https://redhat.atlassian.net/browse/SRVKP-10938) | **Release Notes Not Required:** Internal: remove GitHub comment dedup workaround *(generated)* |
| 8fa0777 | chore: sync PR template types with linter config | Zaki Shaikh | [#2522](https://github.com/tektoncd/pipelines-as-code/pull/2522) | — | — |
| 160f7a4 | chore(deps): bump actions/upload-artifact from 6.0.0 to 7.0. | dependabot[bot] | [#2521](https://github.com/tektoncd/pipelines-as-code/pull/2521) | — | skipped |
| e83a17f | test: Implement file based configuration for test environmen | Chmouel Boudjnah | [#2512](https://github.com/tektoncd/pipelines-as-code/pull/2512) | — | — |
| f055cdf | chore: remove all commit prefix checks in pr template | Zaki Shaikh | [#2525](https://github.com/tektoncd/pipelines-as-code/pull/2525) | — | — |
| eb6cd3d | test: use dynamic projects and smee for GitLab E2E | Chmouel Boudjnah | [#2524](https://github.com/tektoncd/pipelines-as-code/pull/2524) | — | — |
| 95bb453 | fix(github): restrict comment editing to PAC-owned comments | Akshay Pant | [#2487](https://github.com/tektoncd/pipelines-as-code/pull/2487) | [SRVKP-10857](https://redhat.atlassian.net/browse/SRVKP-10857) | **Bug Fix:** Restricts GitHub comment editing to only PAC-owned comments, preventing unintend *(generated)* |
| 093ef51 | test: move BadYaml e2e test from GitHub to GitLab | Chmouel Boudjnah | [#2526](https://github.com/tektoncd/pipelines-as-code/pull/2526) | — | — |
| 88f59ae | docs: correct artifacthub API url in docs | Zaki Shaikh | [#2530](https://github.com/tektoncd/pipelines-as-code/pull/2530) | — | — |
| 05529c9 | fix(gitlab): workaround the gitlab diff API limit when neces | Andrew Thorp | [#2482](https://github.com/tektoncd/pipelines-as-code/pull/2482) | [SRVKP-10677](https://redhat.atlassian.net/browse/SRVKP-10677) | **Bug Fix:** Works around the GitLab diff API size limit to handle large merge requests with  *(generated)* |
| 34f1bfb | docs: migrate documentation from hugo-book to Hextra theme | Chmouel Boudjnah | [#2532](https://github.com/tektoncd/pipelines-as-code/pull/2532) | [SRVKP-11068](https://redhat.atlassian.net/browse/SRVKP-11068) *(keyword)* | **Release Notes Not Required:** Internal: migrate documentation from hugo-book to Hextra theme *(generated)* |
| 62c8fad | feat: add documentation link checking and auto-fix tooling | Chmouel Boudjnah | [#2532](https://github.com/tektoncd/pipelines-as-code/pull/2532) | — | — |
| e13b345 | chore: update CLAUDE.md reference | Chmouel Boudjnah | [#2532](https://github.com/tektoncd/pipelines-as-code/pull/2532) | — | — |
| efb7f77 | docs: Restructure configuration settings documentation | Chmouel Boudjnah | [#2535](https://github.com/tektoncd/pipelines-as-code/pull/2535) | — | — |
| b0864df | docs: add project logo to readme | Chmouel Boudjnah | [#2535](https://github.com/tektoncd/pipelines-as-code/pull/2535) | — | — |
| 94f0777 | fix: use correct os name for mac in htmltest | Zaki Shaikh | [#2536](https://github.com/tektoncd/pipelines-as-code/pull/2536) | — | — |
| 2cbc52f | fix: improve message when /retest has nothing to retest | Shubham Mathur | [#2399](https://github.com/tektoncd/pipelines-as-code/pull/2399) | [SRVKP-10373](https://redhat.atlassian.net/browse/SRVKP-10373) | **Bug Fix:** Improves the error message when /retest finds no matching PipelineRuns to rerun. *(generated)* |
| ccefffa | fix: improve message when /retest has nothing to retest | Shubham Mathur | [#2399](https://github.com/tektoncd/pipelines-as-code/pull/2399) | [SRVKP-10373](https://redhat.atlassian.net/browse/SRVKP-10373) | **Bug Fix:** Improves the error message when /retest finds no matching PipelineRuns to rerun. *(generated)* |
| f8814d8 | ci: use steps syntax for cache-fetch result reference | Akshay Pant | [#2537](https://github.com/tektoncd/pipelines-as-code/pull/2537) | — | — |
| 00e0d26 | test: add GHE webhook dynamic repo creation | Akshay Pant | [#2529](https://github.com/tektoncd/pipelines-as-code/pull/2529) | — | — |
| b78a425 | style: fix golangci-lint errors | Akshay Pant | [#2540](https://github.com/tektoncd/pipelines-as-code/pull/2540) | — | — |
| 95e429d | refactor(test): restructure GitHub GHE test setup | Akshay Pant | [#2540](https://github.com/tektoncd/pipelines-as-code/pull/2540) | — | — |
| ba3299c | chore: don't set VERSION in gitignore | Chmouel Boudjnah | [#2538](https://github.com/tektoncd/pipelines-as-code/pull/2538) | — | — |
| 78f0f37 | test(github): prevent nil panic in GHE test teardown | Akshay Pant | [#2542](https://github.com/tektoncd/pipelines-as-code/pull/2542) | — | — |
| fbbf735 | test: rename GHE webhook tests for consistency | Akshay Pant | [#2542](https://github.com/tektoncd/pipelines-as-code/pull/2542) | — | — |
| 39b1aaa | fix(forgejo): bump forgejo-sdk to bypass Anubis AI protectio | Martin Basti | [#2544](https://github.com/tektoncd/pipelines-as-code/pull/2544) | [SRVKP-11012](https://redhat.atlassian.net/browse/SRVKP-11012) | **Bug Fix:** Bumps the Forgejo SDK to a version that bypasses Anubis AI protection, restoring *(generated)* |
| b065d8f | fix: update team creation to match new Forgejo SDK requireme | Chmouel Boudjnah | [#2544](https://github.com/tektoncd/pipelines-as-code/pull/2544) | [SRVKP-11012](https://redhat.atlassian.net/browse/SRVKP-11012) | **Bug Fix:** Updates team creation logic to comply with new Forgejo SDK requirements. *(generated)* |
| b216b05 | chore(doc-update): Clarify App ID location in GitHub App det | Shubham Mathur | [#2546](https://github.com/tektoncd/pipelines-as-code/pull/2546) | — | — |
| 77d1de2 | test: add e2e for GitLab comment update strategy | Chmouel Boudjnah | [#2531](https://github.com/tektoncd/pipelines-as-code/pull/2531) | — | — |
| 7c38f08 | test(gitlab): fix CEL expression in update strategy test | Akshay Pant | [#2547](https://github.com/tektoncd/pipelines-as-code/pull/2547) | — | — |
| 46ec8a2 | fix(lint): resolve golangci-lint issues | Chmouel Boudjnah | [#2551](https://github.com/tektoncd/pipelines-as-code/pull/2551) | — | — |
| 09bd363 | fix: use correct installation id in ghe pull request test | Zaki Shaikh | [#2548](https://github.com/tektoncd/pipelines-as-code/pull/2548) | — | — |
| 0dd3851 | chore(deps): bump actions/setup-go from 6.2.0 to 6.3.0 | dependabot[bot] | [#2520](https://github.com/tektoncd/pipelines-as-code/pull/2520) | — | skipped |
| 059c01a | refactor(github): consolidate JWT generation | Akshay Pant | [#2541](https://github.com/tektoncd/pipelines-as-code/pull/2541) | [SRVKP-10952](https://redhat.atlassian.net/browse/SRVKP-10952) | **Enhancement:** Consolidates GitHub JWT token generation into a single code path for better main *(generated)* |
| 10dd0b6 | chore(deps): bump actions/download-artifact from 7.0.0 to 8. | dependabot[bot] | [#2519](https://github.com/tektoncd/pipelines-as-code/pull/2519) | — | skipped |
| c49e6bc | fix: update e2e workflow secrets configuration | Chmouel Boudjnah | [#2555](https://github.com/tektoncd/pipelines-as-code/pull/2555) | — | — |
| a676852 | fix: pin actions/checkout to a specific hash | Chmouel Boudjnah | [#2556](https://github.com/tektoncd/pipelines-as-code/pull/2556) | — | — |
| 43e030f | test(ghe webhook): correct PLR count in regex marker test | Akshay Pant | [#2558](https://github.com/tektoncd/pipelines-as-code/pull/2558) | — | — |
| eff5928 | feat: cache changed files in Gitea provider | Chmouel Boudjnah | [#2552](https://github.com/tektoncd/pipelines-as-code/pull/2552) | [SRVKP-10944](https://redhat.atlassian.net/browse/SRVKP-10944) | **Feature:** Gitea provider now caches the list of changed files per pull request, reducing A *(generated)* |
| 5f5f4b3 | chore: migrate to tektoncd org | Chmouel Boudjnah | [#2569](https://github.com/tektoncd/pipelines-as-code/pull/2569) | [SRVKP-11068](https://redhat.atlassian.net/browse/SRVKP-11068) *(keyword)* | **Release Notes Not Required:** Internal: migrate to tektoncd org *(generated)* |
| 539e5ea | test: add documentation to TestGiteaCancelRun | Chmouel Boudjnah | [#2569](https://github.com/tektoncd/pipelines-as-code/pull/2569) | — | — |
| a21ae43 | chore: remove unused PipelineRuns in pac | Chmouel Boudjnah | [#2569](https://github.com/tektoncd/pipelines-as-code/pull/2569) | — | — |
| 0441231 | fix: add http suffix in provider url for forgejo | Zaki Shaikh | [#2573](https://github.com/tektoncd/pipelines-as-code/pull/2573) | [SRVKP-10575](https://redhat.atlassian.net/browse/SRVKP-10575) *(keyword)* | **Bug Fix:** Fixes Forgejo provider URL construction by adding the required HTTP suffix. *(generated)* |
| c6397d3 | fix: add http suffix in provider url for forgejo | Zaki Shaikh | [#2575](https://github.com/tektoncd/pipelines-as-code/pull/2575) | [SRVKP-10575](https://redhat.atlassian.net/browse/SRVKP-10575) *(keyword)* | **Bug Fix:** Fixes Forgejo provider URL construction by adding the required HTTP suffix. *(generated)* |
| de692ba | fix(github): detect GHE instances for URL validation | Akshay Pant | [#2514](https://github.com/tektoncd/pipelines-as-code/pull/2514) | [SRVKP-10943](https://redhat.atlassian.net/browse/SRVKP-10943) | **Bug Fix:** Improves GitHub Enterprise detection by validating URLs to correctly identify GH *(generated)* |
| 3eb7d0f | docs: Add Tekton branding to documentation footer | Chmouel Boudjnah | [#2570](https://github.com/tektoncd/pipelines-as-code/pull/2570) | — | — |
| 78232e2 | feat: migrate docs deployment to custom domain | Chmouel Boudjnah | [#2570](https://github.com/tektoncd/pipelines-as-code/pull/2570) | — | — |
| 6af8b34 | chore(deps): bump actions/download-artifact from 8.0.0 to 8. | dependabot[bot] | [#2574](https://github.com/tektoncd/pipelines-as-code/pull/2574) | — | skipped |
| 900e837 | chore: add release-notes skill | Chmouel Boudjnah | [#2568](https://github.com/tektoncd/pipelines-as-code/pull/2568) | — | — |
| bb03a79 | feat: overhaul README and add link checker | Chmouel Boudjnah | [#2578](https://github.com/tektoncd/pipelines-as-code/pull/2578) | — | — |
| 3ab1ff8 | chore: remove PR CI and Jira skill automation | Chmouel Boudjnah | [#2581](https://github.com/tektoncd/pipelines-as-code/pull/2581) | — | — |
| 85ed765 | docs: simplify the pull request template | Chmouel Boudjnah | [#2581](https://github.com/tektoncd/pipelines-as-code/pull/2581) | — | — |
| d9a6599 | fix: add second gitlab token and group for e2e tests | Chmouel Boudjnah | [#2584](https://github.com/tektoncd/pipelines-as-code/pull/2584) | — | — |
| d7b0258 | Fix stale k8s.io/client-go version in go.mod | David Simansky | [#2585](https://github.com/tektoncd/pipelines-as-code/pull/2585) | [SRVKP-11437](https://redhat.atlassian.net/browse/SRVKP-11437) *(keyword)* | **Bug Fix:** Fixes a stale k8s.io/client-go dependency version in go.mod. *(generated)* |
| f973266 | chore: remove a bunch go.mod replace directives | Chmouel Boudjnah | [#2586](https://github.com/tektoncd/pipelines-as-code/pull/2586) | [SRVKP-12242](https://redhat.atlassian.net/browse/SRVKP-12242) *(keyword)* | **Release Notes Not Required:** Internal: remove a bunch go.mod replace directives *(generated)* |
| d3c9611 | fix: disable documentation link checker | Chmouel Boudjnah | [#2587](https://github.com/tektoncd/pipelines-as-code/pull/2587) | — | — |
| cf25914 | chore(deps): bump google.golang.org/grpc from 1.78.0 to 1.79 | dependabot[bot] | [#2591](https://github.com/tektoncd/pipelines-as-code/pull/2591) | — | skipped |
| 8cc1747 | fix: fix unittest failure when validating gh url | Chmouel Boudjnah | [#2593](https://github.com/tektoncd/pipelines-as-code/pull/2593) | [SRVKP-10943](https://redhat.atlassian.net/browse/SRVKP-10943) *(keyword)* | **Bug Fix:** Fixes a unit test failure in GitHub URL validation. *(generated)* |
| 45aa2f7 | fix: Update /ok-to-test status only for Forgejo | Zaki Shaikh | [#2571](https://github.com/tektoncd/pipelines-as-code/pull/2571) | [SRVKP-11138](https://redhat.atlassian.net/browse/SRVKP-11138) | **Bug Fix:** The /ok-to-test status update is now restricted to Forgejo only, preventing inco *(generated)* |
| ae69110 | chore(deps): bump actions/cache from 5.0.3 to 5.0.4 | dependabot[bot] | [#2600](https://github.com/tektoncd/pipelines-as-code/pull/2600) | — | skipped |
| 30fddf4 | fix(gitlab): fallback to commit status on /retest | Chmouel Boudjnah | [#2583](https://github.com/tektoncd/pipelines-as-code/pull/2583) | [SRVKP-11112](https://redhat.atlassian.net/browse/SRVKP-11112) | **Bug Fix:** GitLab provider now falls back to commit status API when /retest is used, fixing *(generated)* |
| 0127492 | docs(operations): add profiling guide for PAC components | Akshay Pant | [#2602](https://github.com/tektoncd/pipelines-as-code/pull/2602) | — | — |
| de28dca | chore: only build the container when needed | Chmouel Boudjnah | [#2588](https://github.com/tektoncd/pipelines-as-code/pull/2588) | — | — |
| baaf5e1 | refactor: migrate from OpenCensus to OpenTelemetry | Zaki Shaikh | [#2567](https://github.com/tektoncd/pipelines-as-code/pull/2567) | [SRVKP-8544](https://redhat.atlassian.net/browse/SRVKP-8544) | **Enhancement:** Migrates Pipelines as Code observability from OpenCensus to OpenTelemetry, align *(generated)* |
| d49a8e0 | test: use /test TestGiteaReTestAll (and rename it) | Chmouel Boudjnah | [#2589](https://github.com/tektoncd/pipelines-as-code/pull/2589) | — | — |
| 317754c | test: run e2e on go.mod, go.sum and modules.txt | Zaki Shaikh | [#2609](https://github.com/tektoncd/pipelines-as-code/pull/2609) | — | — |
| 7e6df69 | feat: update gendocs to handle old version | Chmouel Boudjnah | [#2610](https://github.com/tektoncd/pipelines-as-code/pull/2610) | — | — |
| d11f7aa | test: remove nightly tests completely | Zaki Shaikh | [#2553](https://github.com/tektoncd/pipelines-as-code/pull/2553) | — | — |
| 0e9d2a7 | fix: enforce TEST_GITHUB_SECOND_WEBHOOK_SECRET existence | Zaki Shaikh | [#2553](https://github.com/tektoncd/pipelines-as-code/pull/2553) | — | — |
| 71d5588 | fix: add TEST_GITHUB_SECOND_APPLICATION_ID variable for ghe | Zaki Shaikh | [#2553](https://github.com/tektoncd/pipelines-as-code/pull/2553) | — | — |
| 870b15b | feat(github): add configurable GitOps command prefix | Zaki Shaikh | [#2443](https://github.com/tektoncd/pipelines-as-code/pull/2443) | [SRVKP-7197](https://redhat.atlassian.net/browse/SRVKP-7197) | **Feature:** Adds support for a configurable GitOps command prefix, allowing administrators t *(generated)* |
| e6e217b | fix: use dynamic prefix in retest error | Zaki Shaikh | [#2443](https://github.com/tektoncd/pipelines-as-code/pull/2443) | [SRVKP-7197](https://redhat.atlassian.net/browse/SRVKP-7197) | **Bug Fix:** Updates error messages to use the dynamically configured GitOps command prefix i *(generated)* |
| e40416c | feat: integrate testrr for CI test reporting | Chmouel Boudjnah | [#2606](https://github.com/tektoncd/pipelines-as-code/pull/2606) | — | — |
| c2b69be | chore(deps): bump github.com/tektoncd/pipeline from 1.10.0 t | dependabot[bot] | [#2607](https://github.com/tektoncd/pipelines-as-code/pull/2607) | — | skipped |
| 182d181 | fix(e2e): find controller in installation namespaces | Akshay Pant | [#2603](https://github.com/tektoncd/pipelines-as-code/pull/2603) | — | — |
| 1c44713 | chore: update dependencies | Chmouel Boudjnah | [#2614](https://github.com/tektoncd/pipelines-as-code/pull/2614) | — | — |
| 56a3a61 | chore: update dependencies and documentation linting rules | Chmouel Boudjnah | [#2617](https://github.com/tektoncd/pipelines-as-code/pull/2617) | — | — |
| 95a5344 | fix: update URLs after move to @pipelines-as-code | Chmouel Boudjnah | [#2619](https://github.com/tektoncd/pipelines-as-code/pull/2619) | [SRVKP-11068](https://redhat.atlassian.net/browse/SRVKP-11068) *(keyword)* | **Bug Fix:** Updates internal URLs to reflect the move to the tektoncd GitHub organization. *(generated)* |
| b044d7f | fix: update stepaction url in test data | Chmouel Boudjnah | [#2620](https://github.com/tektoncd/pipelines-as-code/pull/2620) | [SRVKP-11068](https://redhat.atlassian.net/browse/SRVKP-11068) *(keyword)* | **Bug Fix:** Updates StepAction URLs in test data to reflect the new tektoncd organization. *(generated)* |
| e5877db | fix(test): 404 retry for member propagation lag | Chmouel Boudjnah | [#2620](https://github.com/tektoncd/pipelines-as-code/pull/2620) | [SRVKP-11068](https://redhat.atlassian.net/browse/SRVKP-11068) *(keyword)* | **Release Notes Not Required:** Internal: 404 retry for member propagation lag *(generated)* |
| 3872461 | chore: use gotestsum for running tests | Chmouel Boudjnah | [#2620](https://github.com/tektoncd/pipelines-as-code/pull/2620) | — | — |
| 1d04509 | test(e2e): allow external PRs via ok-to-test label | Zaki Shaikh | [#2611](https://github.com/tektoncd/pipelines-as-code/pull/2611) | — | — |
| 0faad24 | fix(adapter): eliminate data race on shared event field | Chmouel Boudjnah | [#2590](https://github.com/tektoncd/pipelines-as-code/pull/2590) | [SRVKP-10092](https://redhat.atlassian.net/browse/SRVKP-10092) *(keyword)* | **Bug Fix:** Fixes a data race condition on the shared event field in the adapter component. *(generated)* |
| 11e7bc2 | fix: TestGHEPullRequestGitopsCommentCancel race | Chmouel Boudjnah | [#2598](https://github.com/tektoncd/pipelines-as-code/pull/2598) | — | — |
| 33c6fae | ci(e2e): trigger e2e on .github/scripts changes | Zaki Shaikh | [#2623](https://github.com/tektoncd/pipelines-as-code/pull/2623) | — | — |
| e61bf66 | fix(ci): remove ok-to-test label immediately after check | Zaki Shaikh | [#2622](https://github.com/tektoncd/pipelines-as-code/pull/2622) | — | — |
| a65895c | chore: add .gemini configuration file for ignroePaths | Zaki Shaikh | [#2618](https://github.com/tektoncd/pipelines-as-code/pull/2618) | — | — |
| 1ce8d6f | fix(e2e): correct team slugs for permission check | Zaki Shaikh | [#2625](https://github.com/tektoncd/pipelines-as-code/pull/2625) | — | — |
| 66bd8c2 | tests: enforce conventions for go Test | Chmouel Boudjnah | [#2624](https://github.com/tektoncd/pipelines-as-code/pull/2624) | — | — |
| 2ea44d9 | chore: use pipelines-as-code org for pac-boussole | Zaki Shaikh | [#2626](https://github.com/tektoncd/pipelines-as-code/pull/2626) | — | — |
| 091f843 | fix(openai): use max_completion_tokens param | Akshay Pant | [#2615](https://github.com/tektoncd/pipelines-as-code/pull/2615) | [SRVKP-11465](https://redhat.atlassian.net/browse/SRVKP-11465) | **Bug Fix:** Uses the max_completion_tokens parameter for OpenAI API calls, fixing compatibil *(generated)* |
| f8e4343 | fix(llm): default analysis to failed pipelines only | Akshay Pant | [#2615](https://github.com/tektoncd/pipelines-as-code/pull/2615) | [SRVKP-11465](https://redhat.atlassian.net/browse/SRVKP-11465) | **Bug Fix:** LLM-based pipeline analysis now defaults to analyzing only failed pipelines, red *(generated)* |
| 02a9f6a | chore: replace real time.Sleep with clockwork in unit tests | Chmouel Boudjnah | [#2627](https://github.com/tektoncd/pipelines-as-code/pull/2627) | — | — |
| 439e536 | refactor: move secret creation to reconciler | Zaki Shaikh | [#2510](https://github.com/tektoncd/pipelines-as-code/pull/2510) | [SRVKP-11482](https://redhat.atlassian.net/browse/SRVKP-11482) *(keyword)* | **Enhancement:** Moves secret creation logic into the reconciler for better lifecycle management  *(generated)* |
| ece5326 | fix(resolve): restore relative task path resolution for repo | Akshay Pant | [#2554](https://github.com/tektoncd/pipelines-as-code/pull/2554) | [SRVKP-12575](https://redhat.atlassian.net/browse/SRVKP-12575) | **Bug Fix:** Restores relative task path resolution for repository paths, fixing a regression *(generated)* |
| b8ed140 | feat: Implement GraphQL batch fetching for .tekton | Chmouel Boudjnah | [#2423](https://github.com/tektoncd/pipelines-as-code/pull/2423) | [SRVKP-11470](https://redhat.atlassian.net/browse/SRVKP-11470) | **Feature:** Implements GraphQL batch fetching for .tekton directory contents, significantly  *(generated)* |
| 92b37cf | perf(github): cache getPullRequest result in Provider | Akshay Pant | [#2621](https://github.com/tektoncd/pipelines-as-code/pull/2621) | [SRVKP-11121](https://redhat.atlassian.net/browse/SRVKP-11121) | **Enhancement:** Caches getPullRequest results in the GitHub provider to avoid redundant API call *(generated)* |
| ef35fcc | test: use gotestsum for unit test if available | Chmouel Boudjnah | [#2631](https://github.com/tektoncd/pipelines-as-code/pull/2631) | — | — |
| f2b8933 | test: fix test execution and improve assertions | Chmouel Boudjnah | [#2631](https://github.com/tektoncd/pipelines-as-code/pull/2631) | — | — |
| dfab680 | ci(e2e): use ok-to-test action for permission checks | Zaki Shaikh | [#2628](https://github.com/tektoncd/pipelines-as-code/pull/2628) | — | — |
| f1e1fb9 | ci: update ok-to-test action commit SHA | Zaki Shaikh | [#2634](https://github.com/tektoncd/pipelines-as-code/pull/2634) | — | — |
| bdcbf41 | fix(gh): fix re-run button on Pull Requests | Chmouel Boudjnah | [#2597](https://github.com/tektoncd/pipelines-as-code/pull/2597) | [SRVKP-11161](https://redhat.atlassian.net/browse/SRVKP-11161) | **Bug Fix:** Fixes the GitHub re-run button on pull requests to correctly trigger new Pipelin *(generated)* |
| ef85098 | fix: watcher secret creation log and osc link | Chmouel Boudjnah | [#2637](https://github.com/tektoncd/pipelines-as-code/pull/2637) | [SRVKP-11482](https://redhat.atlassian.net/browse/SRVKP-11482) *(keyword)* | **Bug Fix:** Improves watcher secret creation logging and adds OpenShift Console link for bet *(generated)* |
| 010acfd | chore(deps): bump github.com/go-jose/go-jose/v4 from 4.1.3 t | dependabot[bot] | [#2639](https://github.com/tektoncd/pipelines-as-code/pull/2639) | — | skipped |
| a2f56e8 | chore(deps): bump github.com/go-jose/go-jose/v3 from 3.0.4 t | dependabot[bot] | [#2638](https://github.com/tektoncd/pipelines-as-code/pull/2638) | — | skipped |
| 6779bd3 | chore(deps): bump actions/setup-go from 6.3.0 to 6.4.0 | dependabot[bot] | [#2641](https://github.com/tektoncd/pipelines-as-code/pull/2641) | — | skipped |
| cb4296f | chore(deps): bump jaxxstorm/action-install-gh-release | dependabot[bot] | [#2640](https://github.com/tektoncd/pipelines-as-code/pull/2640) | — | skipped |
| a8db555 | fix: PR close condition in e2e tests | Zaki Shaikh | [#2645](https://github.com/tektoncd/pipelines-as-code/pull/2645) | — | — |
| 8c1e41e | docs: rearrange events in gh app docs | Zaki Shaikh | [#2648](https://github.com/tektoncd/pipelines-as-code/pull/2648) | — | — |
| d78ef4e | ci(gha): fix zizmor security findings in workflows | Akshay Pant | [#2632](https://github.com/tektoncd/pipelines-as-code/pull/2632) | — | — |
| a387d41 | fix: populate DefaultBranch for incoming webhooks | Vincent Demeester | [#2647](https://github.com/tektoncd/pipelines-as-code/pull/2647) | [SRVKP-11511](https://redhat.atlassian.net/browse/SRVKP-11511) | **Bug Fix:** Fixes DefaultBranch population for incoming webhook payloads, resolving failures *(generated)* |
| 501ba3b | fix(ci): allow actions token write permission on PR | Zaki Shaikh | [#2650](https://github.com/tektoncd/pipelines-as-code/pull/2650) | — | — |
| 92a3273 | fix(ci): pass smee URL env vars to e2e test runner step | Zaki Shaikh | [#2651](https://github.com/tektoncd/pipelines-as-code/pull/2651) | — | — |
| 69c1197 | chore: Update Go and third-party dependencies | Chmouel Boudjnah | [#2656](https://github.com/tektoncd/pipelines-as-code/pull/2656) | — | — |
| 4c8cf58 | docs: update release notes documentation URL format | Chmouel Boudjnah | [#2654](https://github.com/tektoncd/pipelines-as-code/pull/2654) | — | — |
| 421f96c | docs(forgejo): fix incorrect webhook signature validation cl | Chmouel Boudjnah | [#2658](https://github.com/tektoncd/pipelines-as-code/pull/2658) | [SRVKP-11165](https://redhat.atlassian.net/browse/SRVKP-11165) *(keyword)* | **Release Notes Not Required:** Internal: fix incorrect webhook signature validation claim *(generated)* |
| ddd649e | chore: use GH_TOKEN secert in permission check step | Zaki Shaikh | [#2660](https://github.com/tektoncd/pipelines-as-code/pull/2660) | — | — |
| 2ada0e4 | test: remove outdated tests and stabilize ordering | Chmouel Boudjnah | [#2649](https://github.com/tektoncd/pipelines-as-code/pull/2649) | — | — |
| 14a8e6b | chore: replace e2e env variables from server to data center | Zaki Shaikh | [#2666](https://github.com/tektoncd/pipelines-as-code/pull/2666) | [SRVKP-9638](https://redhat.atlassian.net/browse/SRVKP-9638) *(keyword)* | **Release Notes Not Required:** Internal: replace e2e env variables from server to data center *(generated)* |
| 3c89639 | fix(gitlab): update /ok-to-test status to success for GitLab | Zaki Shaikh | [#2642](https://github.com/tektoncd/pipelines-as-code/pull/2642) | [SRVKP-11156](https://redhat.atlassian.net/browse/SRVKP-11156) | **Bug Fix:** Updates /ok-to-test status to success for GitLab, fixing stuck pending status af *(generated)* |
| 20688f1 | fix(github): guard nil response and cap comment pagination i | Chmouel Boudjnah | [#2663](https://github.com/tektoncd/pipelines-as-code/pull/2663) | [SRVKP-11532](https://redhat.atlassian.net/browse/SRVKP-11532) *(keyword)* | **Bug Fix:** Guards against nil API responses and caps comment pagination in ACL permission c *(generated)* |
| 0f22da7 | docs: use bitbucket datacenter in docs | Zaki Shaikh | [#2668](https://github.com/tektoncd/pipelines-as-code/pull/2668) | [SRVKP-9638](https://redhat.atlassian.net/browse/SRVKP-9638) *(keyword)* | **Release Notes Not Required:** Internal: use bitbucket datacenter in docs *(generated)* |
| 4c43e22 | test: replace bb server to bb data center in test config | Zaki Shaikh | [#2668](https://github.com/tektoncd/pipelines-as-code/pull/2668) | [SRVKP-9638](https://redhat.atlassian.net/browse/SRVKP-9638) *(keyword)* | **Release Notes Not Required:** Internal: replace bb server to bb data center in test config *(generated)* |
| 6a4a5b0 | feat: implement GetCommitStatuses on forgejo | Chmouel Boudjnah | [#2659](https://github.com/tektoncd/pipelines-as-code/pull/2659) | [SRVKP-11529](https://redhat.atlassian.net/browse/SRVKP-11529) | **Feature:** Implements GetCommitStatuses on the Forgejo provider, enabling commit status ret *(generated)* |
| d407897 | chore: use provider wide funcs | Zaki Shaikh | [#2672](https://github.com/tektoncd/pipelines-as-code/pull/2672) | — | — |
| 3169e83 | feat: allow configuration of gotestsum output format | Chmouel Boudjnah | [#2675](https://github.com/tektoncd/pipelines-as-code/pull/2675) | — | — |
| 658e9d4 | fix(gitlab): map skipped status correctly | Akshay Pant | [#2676](https://github.com/tektoncd/pipelines-as-code/pull/2676) | [SRVKP-11540](https://redhat.atlassian.net/browse/SRVKP-11540) | **Bug Fix:** Correctly maps the skipped pipeline status in GitLab, preventing incorrect statu *(generated)* |
| 738cc02 | chore: Add github step summaries for e2e test suites | Chmouel Boudjnah | [#2670](https://github.com/tektoncd/pipelines-as-code/pull/2670) | — | — |
| 4899355 | chore(deps): bump actions/upload-artifact from 7.0.0 to 7.0. | dependabot[bot] | [#2678](https://github.com/tektoncd/pipelines-as-code/pull/2678) | — | skipped |
| 7440c7d | ci: update tar extraction for new zizmor archive | Akshay Pant | [#2679](https://github.com/tektoncd/pipelines-as-code/pull/2679) | — | — |
| e1ee25f | refactor(llm): simplify and restructure the LLM package | Chmouel Boudjnah | [#2673](https://github.com/tektoncd/pipelines-as-code/pull/2673) | [SRVKP-11539](https://redhat.atlassian.net/browse/SRVKP-11539) | **Enhancement:** Simplifies and restructures the LLM package for better maintainability and exten *(generated)* |
| 2dc12b5 | chore: update default LLM models | Chmouel Boudjnah | [#2673](https://github.com/tektoncd/pipelines-as-code/pull/2673) | [SRVKP-11539](https://redhat.atlassian.net/browse/SRVKP-11539) | **Release Notes Not Required:** Internal: update default LLM models *(generated)* |
| de6de63 | chore: return early from detect for edited comments | Zaki Shaikh | [#2674](https://github.com/tektoncd/pipelines-as-code/pull/2674) | — | — |
| a6e4aab | fix restrict same repo ACL perm to trusted context | Chmouel Boudjnah | [#2665](https://github.com/tektoncd/pipelines-as-code/pull/2665) | [SRVKP-11527](https://redhat.atlassian.net/browse/SRVKP-11527) *(keyword)* | **Bug Fix:** Restricts same-repo ACL permission checks to trusted execution contexts only. *(generated)* |
| c9be9d6 | fix(webhook): prevent duplicate repository CR on trailing sl | Zaki Shaikh | [#2683](https://github.com/tektoncd/pipelines-as-code/pull/2683) | [SRVKP-11295](https://redhat.atlassian.net/browse/SRVKP-11295) | **Bug Fix:** Prevents duplicate Repository CR creation when the webhook URL contains a traili *(generated)* |
| 29725a3 | chore: add reviewers and approvers in OWNERS | Zaki Shaikh | [#2686](https://github.com/tektoncd/pipelines-as-code/pull/2686) | — | — |
| 7bbf595 | docs: fix the quickstart bootstrap command | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 1c75525 | docs: describe the LLM trigger default accurately | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 96864eb | docs: clarify bootstrap github-app behavior | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| cb4d67c | docs: include all CLI commands in the index | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 6bf01c8 | docs: match shell completion docs to the command | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 0df5a1d | docs: show the correct repository delete syntax | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| d8904e5 | docs: mention Forgejo in the installation overview | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 3ace25d | docs: clean up formatting and typos | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| c110dcb | docs: standardize the Homebrew install command | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 8c05b81 | chore: Add symbolic link for agent skills | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| e7e9a61 | docs: correct token scoping prerequisites | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 32d0765 | docs: fix default Artifact Hub API URL | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| b8c1241 | docs: correct comment strategy field path | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| f8b0a96 | docs: replace nonexistent policy.team_ids reference | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 46c84dc | docs: align on_cel examples with CEL context | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 565f857 | docs: document custom console namespace URL | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 0384b59 | docs: add gitops_command_prefix to settings ref | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 9fc27b5 | docs: fix Forgejo webhook and hub URL drift | Chmouel Boudjnah | [#2684](https://github.com/tektoncd/pipelines-as-code/pull/2684) | — | — |
| 80f9be1 | fix: update repository URLs to upstream tektoncd | Chmouel Boudjnah | [#2691](https://github.com/tektoncd/pipelines-as-code/pull/2691) | [SRVKP-11068](https://redhat.atlassian.net/browse/SRVKP-11068) *(keyword)* | **Bug Fix:** Updates repository URLs to the upstream tektoncd organization. *(generated)* |
| bd9f468 | feat: enable tracing for webhooks and PipelineRun execution | Josiah England | [#2605](https://github.com/tektoncd/pipelines-as-code/pull/2605) | [SRVKP-8544](https://redhat.atlassian.net/browse/SRVKP-8544) | **Feature:** Enables distributed tracing for webhook processing and PipelineRun execution, pr *(generated)* |
| dded990 | refactor: move  client setup to its own package | Zaki Shaikh | [#2682](https://github.com/tektoncd/pipelines-as-code/pull/2682) | [SRVKP-11557](https://redhat.atlassian.net/browse/SRVKP-11557) | **Enhancement:** Moves the provider client setup into its own package for better code organizatio *(generated)* |
| 532790b | fix(github-webhook): /ok-to-test is not triggering CI on PRs | Zaki Shaikh | [#2682](https://github.com/tektoncd/pipelines-as-code/pull/2682) | [SRVKP-11557](https://redhat.atlassian.net/browse/SRVKP-11557) | **Bug Fix:** Fixes /ok-to-test not triggering CI on pull requests when using GitHub webhooks. *(generated)* |
| 8892691 | perf(informer): add TransformFuncs to reduce cache memory us | Akshay Pant | [#2667](https://github.com/tektoncd/pipelines-as-code/pull/2667) | [SRVKP-8544](https://redhat.atlassian.net/browse/SRVKP-8544) *(keyword)* | **Enhancement:** Adds TransformFuncs to Kubernetes informers to strip unnecessary metadata from c *(generated)* |
| ed2592b | docs(informer): add cache optimization documentation | Akshay Pant | [#2667](https://github.com/tektoncd/pipelines-as-code/pull/2667) | — | — |
| c51c085 | feat(SRVKP-11530): enable recursive .tekton dir retrieval fo | Zaki Shaikh | [#2694](https://github.com/tektoncd/pipelines-as-code/pull/2694) | [SRVKP-11530](https://redhat.atlassian.net/browse/SRVKP-11530) | **Feature:** Enables recursive retrieval of .tekton directory contents for the Forgejo provid *(generated)* |
| 9c4e9cb | fix: fix TestGithubGHEPullRequestOkToTest | Zaki Shaikh | [#2697](https://github.com/tektoncd/pipelines-as-code/pull/2697) | — | — |
| 66add4d | chore(deps): bump actions/cache from 5.0.4 to 5.0.5 | dependabot[bot] | [#2693](https://github.com/tektoncd/pipelines-as-code/pull/2693) | — | skipped |
| 61937fa | chore(deps): bump github.com/tektoncd/pipeline from 1.11.0 t | dependabot[bot] | [#2698](https://github.com/tektoncd/pipelines-as-code/pull/2698) | — | skipped |
| 6b69972 | fix: improve unexpected event filtering in tests | Chmouel Boudjnah | [#2688](https://github.com/tektoncd/pipelines-as-code/pull/2688) | — | — |
| 5c09648 | feat(cel): handle Forgejo headers in CLI provider auto-detec | Shubham Bhardwaj | [#2700](https://github.com/tektoncd/pipelines-as-code/pull/2700) | [SRVKP-11638](https://redhat.atlassian.net/browse/SRVKP-11638) | **Feature:** Adds Forgejo header detection to the CEL-based CLI provider auto-detection, impr *(generated)* |
| 47ee286 | fix(github): use provided target ref in GetFileInsideRepo | Akshay Pant | [#2696](https://github.com/tektoncd/pipelines-as-code/pull/2696) | [SRVKP-11527](https://redhat.atlassian.net/browse/SRVKP-11527) | **Bug Fix:** Uses the provided target ref instead of the default branch when fetching files f *(generated)* |
| 68c5f1d | fix(bitbucket-cloud): resolve CEL expression failure on push | Zaki Shaikh | [#2704](https://github.com/tektoncd/pipelines-as-code/pull/2704) | [SRVKP-9635](https://redhat.atlassian.net/browse/SRVKP-9635) | **Bug Fix:** Resolves CEL expression evaluation failure on Bitbucket Cloud push events. *(generated)* |
| 7c46558 | perf(github): cache check-run lookups with retry | Akshay Pant | [#2669](https://github.com/tektoncd/pipelines-as-code/pull/2669) | [SRVKP-11568](https://redhat.atlassian.net/browse/SRVKP-11568) | **Enhancement:** Caches check-run lookups with retry logic in the GitHub provider, reducing API c *(generated)* |
| 88e8d39 | chore: Use GHE fro github enterprise instead of second | Zaki Shaikh | [#2707](https://github.com/tektoncd/pipelines-as-code/pull/2707) | — | — |
| 357d970 | Revert "chore: Use GHE fro github enterprise instead of seco | Zaki Shaikh | [#2710](https://github.com/tektoncd/pipelines-as-code/pull/2710) | — | — |
| 0948d88 | docs: note about e2e permission in bb cloud docs | Zaki Shaikh | [#2703](https://github.com/tektoncd/pipelines-as-code/pull/2703) | — | — |
| e951778 | fix: fix skip-install option to bootstrap command | Chmouel Boudjnah | [#2709](https://github.com/tektoncd/pipelines-as-code/pull/2709) | [SRVKP-9062](https://redhat.atlassian.net/browse/SRVKP-9062) *(keyword)* | **Bug Fix:** Fixes the --skip-install option in the bootstrap command to correctly skip the P *(generated)* |
| b0bcbac |   fix(gitlab): pin commit statuses to same pipeline | ab-ghosh | [#2671](https://github.com/tektoncd/pipelines-as-code/pull/2671) | [SRVKP-11437](https://redhat.atlassian.net/browse/SRVKP-11437) | **Bug Fix:** Pins GitLab commit statuses to the same pipeline, preventing build status from a *(generated)* |
| baf49e3 | dep: update go-github dependency | Zaki Shaikh | [#2714](https://github.com/tektoncd/pipelines-as-code/pull/2714) | — | — |
| c37a213 | fix: Use pull request number from issue comment payload | Zaki Shaikh | [#2711](https://github.com/tektoncd/pipelines-as-code/pull/2711) | [SRVKP-10662](https://redhat.atlassian.net/browse/SRVKP-10662) *(keyword)* | **Bug Fix:** Uses the pull request number from the issue comment payload instead of making a  *(generated)* |
| eda5331 | fix(bitbucket-cloud): fix status key handling for build stat | Zaki Shaikh | [#2702](https://github.com/tektoncd/pipelines-as-code/pull/2702) | [SRVKP-11710](https://redhat.atlassian.net/browse/SRVKP-11710) | **Bug Fix:** Fixes status key handling for Bitbucket Cloud build statuses, resolving incorrec *(generated)* |
| 74b0d28 | fix(github-webhook): clear pending check on ok-to-test | Zaki Shaikh | [#2706](https://github.com/tektoncd/pipelines-as-code/pull/2706) | [SRVKP-11789](https://redhat.atlassian.net/browse/SRVKP-11789) | **Bug Fix:** Clears pending check status when /ok-to-test is used on GitHub webhooks. *(generated)* |
| c615efb | chore(github): emit event when API rate limit exceeded | Akshay Pant | [#2715](https://github.com/tektoncd/pipelines-as-code/pull/2715) | [SRVKP-10951](https://redhat.atlassian.net/browse/SRVKP-10951) | **Release Notes Not Required:** Internal: emit event when API rate limit exceeded *(generated)* |
| 4b1d7f7 | fix(opscomments): skip key=value args as PR names | Akshay Pant | [#2712](https://github.com/tektoncd/pipelines-as-code/pull/2712) | [SRVKP-10945](https://redhat.atlassian.net/browse/SRVKP-10945) | **Bug Fix:** Fixes ops comment argument parsing to skip key=value style arguments when lookin *(generated)* |
| 861a507 | docs(on-comment): warn against built-in commands | Akshay Pant | [#2712](https://github.com/tektoncd/pipelines-as-code/pull/2712) | [SRVKP-10945](https://redhat.atlassian.net/browse/SRVKP-10945) | **Release Notes Not Required:** Internal: warn against built-in commands *(generated)* |
| 6522de8 | fix: label values when dot are present | Francesco Ilario | [#2724](https://github.com/tektoncd/pipelines-as-code/pull/2724) | [SRVKP-11710](https://redhat.atlassian.net/browse/SRVKP-11710) *(keyword)* | **Bug Fix:** Fixes label value handling when dots are present, ensuring valid Kubernetes labe *(generated)* |
| 25b9307 | add sanification of the input and improve logic | Francesco Ilario | [#2724](https://github.com/tektoncd/pipelines-as-code/pull/2724) | [SRVKP-11710](https://redhat.atlassian.net/browse/SRVKP-11710) *(keyword)* | **Bug Fix:** Adds input sanitization and improves label value normalization logic. *(generated)* |
| e75942f | fix: move from RFC1123 to LabelValue | Francesco Ilario | [#2724](https://github.com/tektoncd/pipelines-as-code/pull/2724) | [SRVKP-11710](https://redhat.atlassian.net/browse/SRVKP-11710) *(keyword)* | **Bug Fix:** Moves label value validation from RFC1123 to Kubernetes LabelValue format for co *(generated)* |
| 5226968 | fix: linter complaints | Francesco Ilario | [#2724](https://github.com/tektoncd/pipelines-as-code/pull/2724) | [SRVKP-11710](https://redhat.atlassian.net/browse/SRVKP-11710) *(keyword)* | **Bug Fix:** Fixes linter complaints in the label sanitization code. *(generated)* |
| a43ae55 | fix: remove unneeded check | Francesco Ilario | [#2724](https://github.com/tektoncd/pipelines-as-code/pull/2724) | [SRVKP-11710](https://redhat.atlassian.net/browse/SRVKP-11710) *(keyword)* | **Bug Fix:** Removes an unneeded validation check in the label sanitization code. *(generated)* |
| f05bcd0 | fix: point to the right doc in formatting | Francesco Ilario | [#2731](https://github.com/tektoncd/pipelines-as-code/pull/2731) | — | — |
| f841d2d | fix(bitbucket-datacenter): detect changes on merged PR push | Zaki Shaikh | [#2719](https://github.com/tektoncd/pipelines-as-code/pull/2719) | [SRVKP-9638](https://redhat.atlassian.net/browse/SRVKP-9638) | **Bug Fix:** Detects file changes on merged pull request push events in Bitbucket Data Center *(generated)* |
| 445941b | feat(cel): enable string and list extension functions in CEL | Zaki Shaikh | [#2725](https://github.com/tektoncd/pipelines-as-code/pull/2725) | [SRVKP-11940](https://redhat.atlassian.net/browse/SRVKP-11940) | **Feature:** Enables string and list extension functions in CEL expressions, providing more p *(generated)* |
| 9db1fec | chore: update golangci linter configuration | Chmouel Boudjnah | [#2736](https://github.com/tektoncd/pipelines-as-code/pull/2736) | — | — |
| 2b9a6f1 | ci: update golangci-lint to v2.12.2 | Akshay Pant | [#2737](https://github.com/tektoncd/pipelines-as-code/pull/2737) | — | — |
| 4c7b0e0 | fix(reconciler): skip watcher status updates | Chmouel Boudjnah | [#2735](https://github.com/tektoncd/pipelines-as-code/pull/2735) | [SRVKP-4264](https://redhat.atlassian.net/browse/SRVKP-4264) *(keyword)* | **Bug Fix:** Skips watcher status updates when not needed, reducing unnecessary API calls. *(generated)* |
| 529a725 | fix(gitlab): post MR comment on inaccessible fork | Akshay Pant | [#2739](https://github.com/tektoncd/pipelines-as-code/pull/2739) | [SRVKP-10477](https://redhat.atlassian.net/browse/SRVKP-10477) | **Bug Fix:** Posts a merge request comment on GitLab when the fork is inaccessible, providing *(generated)* |
| 32820cb | feat(gitea): implement GetTaskURI for remote task resolution | Akshay Pant | [#2732](https://github.com/tektoncd/pipelines-as-code/pull/2732) | [SRVKP-10610](https://redhat.atlassian.net/browse/SRVKP-10610) | **Feature:** Implements GetTaskURI for the Gitea/Forgejo provider, enabling remote task resol *(generated)* |
| 69fa323 | fix(release): preserve dots in image tags for version tag pu | Chmouel Boudjnah | [#2742](https://github.com/tektoncd/pipelines-as-code/pull/2742) | [SRVKP-11068](https://redhat.atlassian.net/browse/SRVKP-11068) *(keyword)* | **Bug Fix:** Preserves dots in image tags for version tag pushes, fixing incorrect tag format *(generated)* |
| 4dac4d6 | feat: Add deprecation warnings for Tekton Hub integration | Chmouel Boudjnah | [#2746](https://github.com/tektoncd/pipelines-as-code/pull/2746) | [SRVKP-12187](https://redhat.atlassian.net/browse/SRVKP-12187) | **Feature:** Adds deprecation warnings when users configure Tekton Hub integration, guiding m *(generated)* |
| 0bb2f82 | feat: add TLS configuration support | Zaki Shaikh | [#2738](https://github.com/tektoncd/pipelines-as-code/pull/2738) | [SRVKP-9681](https://redhat.atlassian.net/browse/SRVKP-9681) | **Feature:** Adds TLS configuration support for the Pipelines as Code controller, enabling se *(generated)* |
| 3811249 | chore(deps): bump knative/eventing to v0.49.0 | Akshay Pant | [#2720](https://github.com/tektoncd/pipelines-as-code/pull/2720) | [SRVKP-11514](https://redhat.atlassian.net/browse/SRVKP-11514) *(keyword)* | **Release Notes Not Required:** Internal: bump knative/eventing to v0.49.0 *(generated)* |
| 67cfa52 | docs(profiling): update guide for OTel migration | Akshay Pant | [#2720](https://github.com/tektoncd/pipelines-as-code/pull/2720) | [SRVKP-8544](https://redhat.atlassian.net/browse/SRVKP-8544) *(keyword)* | **Release Notes Not Required:** Internal: update guide for OTel migration *(generated)* |
| 4d0454b | chore(deps): bump mxschmitt/action-tmate from 3.23 to 3.24 | dependabot[bot] | [#2750](https://github.com/tektoncd/pipelines-as-code/pull/2750) | — | skipped |
| 8854274 | fix: remove unused secrets/delete permission from controller | Chmouel Boudjnah | [#2744](https://github.com/tektoncd/pipelines-as-code/pull/2744) | [SRVKP-10952](https://redhat.atlassian.net/browse/SRVKP-10952) *(keyword)* | **Enhancement:** Removes unused secrets/delete RBAC permission from the controller, following the *(generated)* |
| 223e39c | fix(ci): parse JSON test output for Slack notifications | Akshay Pant | [#2753](https://github.com/tektoncd/pipelines-as-code/pull/2753) | [SRVKP-10661](https://redhat.atlassian.net/browse/SRVKP-10661) *(keyword)* | **Release Notes Not Required:** Internal: parse JSON test output for Slack notifications *(generated)* |
| 2c03760 | fix(security): redact query string from incoming webhook log | Shubham Bhardwaj | [#2754](https://github.com/tektoncd/pipelines-as-code/pull/2754) | [SRVKP-10609](https://redhat.atlassian.net/browse/SRVKP-10609) *(keyword)* | **Bug Fix:** Redacts query string parameters from incoming webhook log entries, preventing se *(generated)* |
| bd262aa | chore: update incoming webhook legacy params deprecation mes | Zaki Shaikh | [#2757](https://github.com/tektoncd/pipelines-as-code/pull/2757) | [SRVKP-11511](https://redhat.atlassian.net/browse/SRVKP-11511) *(keyword)* | **Release Notes Not Required:** Internal: update incoming webhook legacy params deprecation message *(generated)* |
| 0017828 | fix(github): scope App token to triggering repo | Akshay Pant | [#2705](https://github.com/tektoncd/pipelines-as-code/pull/2705) | [SRVKP-12241](https://redhat.atlassian.net/browse/SRVKP-12241) | **Bug Fix:** Scopes GitHub App installation tokens to only the repository that triggered the  *(generated)* |
| ee5d9b0 | fix(resolve): deep-copy cached resources before inlining | Akshay Pant | [#2705](https://github.com/tektoncd/pipelines-as-code/pull/2705) | [SRVKP-12241](https://redhat.atlassian.net/browse/SRVKP-12241) | **Bug Fix:** Fixes a deep-copy issue where cached pipeline resources were mutated during inli *(generated)* |
| 402d5c7 | fix: prevent GitHub Enterprise header hijacking in app token | Chmouel Boudjnah | [#2759](https://github.com/tektoncd/pipelines-as-code/pull/2759) | [SRVKP-12241](https://redhat.atlassian.net/browse/SRVKP-12241) *(keyword)* | **Bug Fix:** Prevents GitHub Enterprise header hijacking in app token requests, fixing a secu *(generated)* |
| 0b6c487 | Release yaml generated from https://github.com/tektoncd/pipe | Pipelines as Code CI Robot | — | — | — |

---

## tektoncd-hub

**Upstream:** openshift-pipelines/hub · **Commits:** 5 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| c79258f | bump UI and swagger packages to fix the CVEs | Shiv Verma | — | — | **CVE:** Fixes CVE-2026-4800 and CVE-2026-29074 by updating the Hub UI and Swagger packag *(generated)* |
| dd532ff | feat(cli): deprecate tkn hub commands for Tekton Hub support | divyansh42 | — | [SRVKP-11950](https://redhat.atlassian.net/browse/SRVKP-11950) *(keyword)* | **Deprecated Functionality:** The tkn hub CLI commands that depend on Tekton Hub (check-upgrade, downgrade, ge *(generated)* |
| e22c048 | bump UI and swagger packages to fix the CVEs | Shiv Verma | — | — | **CVE:** Updates Hub UI and Swagger packages to address CVE vulnerabilities. *(generated)* |
| 5d7d39c | bump go version to v1.25.8 and update dockerfile | Shiv Verma | — | — | **Release Notes Not Required** *(generated)* |
| a84deeb | fix CVE vulnerablities in the UI and Swagger component | Shiv Verma | — | — | **CVE:** Fixes CVE vulnerabilities in the Hub UI and Swagger component. *(generated)* |

---

## tektoncd-pruner

**Upstream:** tektoncd/pruner · **Commits:** 162 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| efbedba | chore(deps): bump go.uber.org/zap from 1.27.0 to 1.27.1 | dependabot[bot] | [#50](https://github.com/tektoncd/pruner/pull/50) | — | skipped |
| 8f46c8d | chore(deps): bump github/codeql-action from 4.31.2 to 4.31.4 | dependabot[bot] | [#53](https://github.com/tektoncd/pruner/pull/53) | — | skipped |
| ba21946 | chore(deps): bump actions/checkout from 5.0.0 to 6.0.0 | dependabot[bot] | [#52](https://github.com/tektoncd/pruner/pull/52) | — | skipped |
| 825d41d | chore(deps): bump golangci/golangci-lint-action from 9.0.0 t | dependabot[bot] | [#54](https://github.com/tektoncd/pruner/pull/54) | — | skipped |
| 682e4a2 | chore(deps): bump actions/dependency-review-action from 4.8. | dependabot[bot] | [#55](https://github.com/tektoncd/pruner/pull/55) | — | skipped |
| f6142f7 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#56](https://github.com/tektoncd/pruner/pull/56) | — | skipped |
| da5b765 | chore(deps): bump tj-actions/changed-files | dependabot[bot] | [#51](https://github.com/tektoncd/pruner/pull/51) | — | skipped |
| 9524430 | feat: introduce new function to expose validation logic on s | Anitha Natarajan | [#57](https://github.com/tektoncd/pruner/pull/57) | [SRVKP-9431](https://redhat.atlassian.net/browse/SRVKP-9431) | **Release Notes Not Required** *(generated)* |
| bef0990 | fix: add missing unit test cases | Anitha Natarajan | [#59](https://github.com/tektoncd/pruner/pull/59) | [SRVKP-9431](https://redhat.atlassian.net/browse/SRVKP-9431) | **Release Notes Not Required** *(generated)* |
| c9eae3c | feat: 0.3.2 retract and doc udpates | Anitha Natarajan | [#60](https://github.com/tektoncd/pruner/pull/60) | [SRVKP-9431](https://redhat.atlassian.net/browse/SRVKP-9431) | **Release Notes Not Required** *(generated)* |
| a63bfe9 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#73](https://github.com/tektoncd/pruner/pull/73) | — | skipped |
| b7de037 | chore(deps): bump softprops/action-gh-release from 2.4.2 to  | dependabot[bot] | [#75](https://github.com/tektoncd/pruner/pull/75) | — | skipped |
| d536ae2 | chore(deps): bump chainguard-dev/actions from 1.5.8 to 1.5.1 | dependabot[bot] | [#76](https://github.com/tektoncd/pruner/pull/76) | — | skipped |
| a9cae00 | chore(deps): bump actions/setup-go from 6.0.0 to 6.1.0 | dependabot[bot] | [#78](https://github.com/tektoncd/pruner/pull/78) | — | skipped |
| da1eb4a | chore(deps): bump tj-actions/changed-files | dependabot[bot] | [#79](https://github.com/tektoncd/pruner/pull/79) | — | skipped |
| b3ada94 | chore(deps): bump github/codeql-action from 4.31.4 to 4.31.5 | dependabot[bot] | [#77](https://github.com/tektoncd/pruner/pull/77) | — | skipped |
| bd99f38 | Fix errors and improve formatting in release cheat sheet | Shubham Bhardwaj | [#87](https://github.com/tektoncd/pruner/pull/87) | [SRVKP-9463](https://redhat.atlassian.net/browse/SRVKP-9463) | **Release Notes Not Required** *(generated)* |
| e56420f | fix: resolve yamllint errors and warnings in workflows | Shubham Bhardwaj | [#86](https://github.com/tektoncd/pruner/pull/86) | [SRVKP-9463](https://redhat.atlassian.net/browse/SRVKP-9463) | **Release Notes Not Required** *(generated)* |
| 37446ec | .github/workflows: use CHATOPS_TOKEN for coverage comments | Vincent Demeester | [#92](https://github.com/tektoncd/pruner/pull/92) | — | — |
| b38434b | Add releases.md documenting release policy and current relea | Shubham Bhardwaj | [#91](https://github.com/tektoncd/pruner/pull/91) | [SRVKP-9473](https://redhat.atlassian.net/browse/SRVKP-9473) | **Release Notes Not Required** *(generated)* |
| b0fd554 | chore(deps): bump github.com/tektoncd/pipeline from 1.6.0 to | dependabot[bot] | [#94](https://github.com/tektoncd/pruner/pull/94) | — | skipped |
| 93cc047 | chore(deps): bump step-security/harden-runner from 2.13.2 to | dependabot[bot] | [#95](https://github.com/tektoncd/pruner/pull/95) | — | skipped |
| fca7eaf | chore(deps): bump golangci/golangci-lint-action from 9.1.0 t | dependabot[bot] | [#98](https://github.com/tektoncd/pruner/pull/98) | — | skipped |
| 34edf35 | chore(deps): bump actions/checkout from 6.0.0 to 6.0.1 | dependabot[bot] | [#99](https://github.com/tektoncd/pruner/pull/99) | — | skipped |
| 7787766 | feat: add cherry-pick slash command workflow | Vincent Demeester | [#109](https://github.com/tektoncd/pruner/pull/109) | — | — |
| 1a2dc82 | Remove tekton-* namespace exclusion from pruner | Shubham Bhardwaj | [#108](https://github.com/tektoncd/pruner/pull/108) | [SRVKP-9677](https://redhat.atlassian.net/browse/SRVKP-9677) | **Bug Fix:** Allow tekton-* namespaces to be processed by the pruner, except tekton-pipelines *(extracted)* |
| f85a120 | Removes the container resource requests and limits from the  | Shubham Bhardwaj | [#106](https://github.com/tektoncd/pruner/pull/106) | [SRVKP-9676](https://redhat.atlassian.net/browse/SRVKP-9676) | **Enhancement:** The pruner controller deployment no longer specifies container resource requests *(generated)* |
| d1773de | chore: migrate retest workflow to use plumbing reusable work | Vincent Demeester | [#110](https://github.com/tektoncd/pruner/pull/110) | — | — |
| ea85281 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#96](https://github.com/tektoncd/pruner/pull/96) | — | skipped |
| 9aff622 | chore(deps): bump github/codeql-action from 4.31.5 to 4.31.7 | dependabot[bot] | [#97](https://github.com/tektoncd/pruner/pull/97) | — | skipped |
| cb88558 | chore(deps): bump go.opentelemetry.io/otel/metric from 1.38. | dependabot[bot] | [#101](https://github.com/tektoncd/pruner/pull/101) | — | skipped |
| b9e852c | fix: Update E2E tests for tekton-* namespace validation chan | Shubham Bhardwaj | [#111](https://github.com/tektoncd/pruner/pull/111) | [SRVKP-9677](https://redhat.atlassian.net/browse/SRVKP-9677) *(keyword)* | **Release Notes Not Required** *(generated)* |
| b5863e3 | Enable GitHub merge queue | Anitha Natarajan | [#112](https://github.com/tektoncd/pruner/pull/112) | — | — |
| 7685fa4 | ci: Remove version tip override from ko setup | Vincent Demeester | [#120](https://github.com/tektoncd/pruner/pull/120) | — | — |
| e3d3e95 | chore(deps): bump step-security/harden-runner from 2.13.3 to | dependabot[bot] | [#117](https://github.com/tektoncd/pruner/pull/117) | — | skipped |
| 5847668 | chore(deps): bump tj-actions/changed-files | dependabot[bot] | [#118](https://github.com/tektoncd/pruner/pull/118) | — | skipped |
| f4bc715 | chore(deps): bump github/codeql-action from 4.31.7 to 4.31.8 | dependabot[bot] | [#119](https://github.com/tektoncd/pruner/pull/119) | — | skipped |
| d6337bf | chore(deps): bump actions/upload-artifact from 5.0.0 to 6.0. | dependabot[bot] | [#114](https://github.com/tektoncd/pruner/pull/114) | — | skipped |
| f0bf449 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#116](https://github.com/tektoncd/pruner/pull/116) | — | skipped |
| be6fa45 | chore(deps): bump actions/cache from 4.3.0 to 5.0.1 | dependabot[bot] | [#115](https://github.com/tektoncd/pruner/pull/115) | — | skipped |
| 012ad58 | chore(deps): bump the all group in /tekton with 3 updates | dependabot[bot] | [#130](https://github.com/tektoncd/pruner/pull/130) | — | skipped |
| 39f2fa0 | chore: update workflows for dependabot and migrate to infra. | Shubham Bhardwaj | [#128](https://github.com/tektoncd/pruner/pull/128) | — | — |
| 8bb5cee | chore: add omitempty to YAML/JSON struct tags | Shubham Bhardwaj | [#127](https://github.com/tektoncd/pruner/pull/127) | [SRVKP-9968](https://redhat.atlassian.net/browse/SRVKP-9968) | **Enhancement:** Namespace config no longer shows null values for unset fields in TektonPruner gl *(extracted)* |
| d394469 | chore(deps): bump peter-evans/slash-command-dispatch from 5. | dependabot[bot] | [#133](https://github.com/tektoncd/pruner/pull/133) | — | skipped |
| 975ffcf | chore(deps): bump github/codeql-action from 4.31.8 to 4.31.9 | dependabot[bot] | [#134](https://github.com/tektoncd/pruner/pull/134) | — | skipped |
| 790b98d | Update pipelines-lts matrix to latest LTS versions | Shubham Bhardwaj | [#135](https://github.com/tektoncd/pruner/pull/135) | — | — |
| e4a3513 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#140](https://github.com/tektoncd/pruner/pull/140) | — | skipped |
| 8feec18 | fix: add JSON struct tags to fix TektonConfig status seriali | Shubham Bhardwaj | [#141](https://github.com/tektoncd/pruner/pull/141) | [SRVKP-9968](https://redhat.atlassian.net/browse/SRVKP-9968) | **Release Notes Not Required** *(extracted)* |
| bef2175 | chore(deps): bump the all group in /tekton with 2 updates | dependabot[bot] | [#143](https://github.com/tektoncd/pruner/pull/143) | — | skipped |
| c2e7a05 | chore: Update release documentation for v0.3.4 | Shubham Bhardwaj | [#138](https://github.com/tektoncd/pruner/pull/138) | — | — |
| aa31556 | fix: selector-based pruning for label and annotation selecto | Shubham Bhardwaj | [#142](https://github.com/tektoncd/pruner/pull/142) | [SRVKP-10028](https://redhat.atlassian.net/browse/SRVKP-10028) | **Bug Fix:** Selector-based pruning now correctly groups PipelineRuns by ConfigMap-defined la *(extracted)* |
| f452a90 | Chore: update release deatils and remove unused workflow | Anitha Natarajan | [#146](https://github.com/tektoncd/pruner/pull/146) | — | — |
| 2f7b72f | chore(deps): bump chainguard-dev/actions from 1.5.10 to 1.5. | dependabot[bot] | [#147](https://github.com/tektoncd/pruner/pull/147) | — | skipped |
| fe03d80 | chore(deps): bump the all group in /tekton with 2 updates | dependabot[bot] | [#148](https://github.com/tektoncd/pruner/pull/148) | — | skipped |
| 61e4c25 | chore(deps): bump actions/cache from 5.0.1 to 5.0.2 | dependabot[bot] | [#152](https://github.com/tektoncd/pruner/pull/152) | — | skipped |
| 7d5655e | chore(deps): bump chainguard-dev/actions from 1.5.11 to 1.5. | dependabot[bot] | [#153](https://github.com/tektoncd/pruner/pull/153) | — | skipped |
| 5a0fd41 | chore(deps): bump the all group in /tekton with 2 updates | dependabot[bot] | [#154](https://github.com/tektoncd/pruner/pull/154) | — | skipped |
| 6f41634 | chore(deps): bump github/codeql-action from 4.31.9 to 4.31.1 | dependabot[bot] | [#155](https://github.com/tektoncd/pruner/pull/155) | — | skipped |
| ad78d5b | chore(deps): bump actions/setup-go from 6.1.0 to 6.2.0 | dependabot[bot] | [#156](https://github.com/tektoncd/pruner/pull/156) | — | skipped |
| fa797de | docs: Fix documentation inconsistencies and incorrect YAML s | Shubham Bhardwaj | [#151](https://github.com/tektoncd/pruner/pull/151) | [SRVKP-7225](https://redhat.atlassian.net/browse/SRVKP-7225) | **Release Notes Not Required** *(extracted)* |
| bc2b207 | chore(deps): bump chainguard-dev/actions from 1.5.12 to 1.5. | dependabot[bot] | [#157](https://github.com/tektoncd/pruner/pull/157) | — | skipped |
| 15f3ed2 | chore(deps): bump the all group in /tekton with 3 updates | dependabot[bot] | [#158](https://github.com/tektoncd/pruner/pull/158) | — | skipped |
| b0816d1 | chore(deps): bump actions/checkout from 6.0.1 to 6.0.2 | dependabot[bot] | [#159](https://github.com/tektoncd/pruner/pull/159) | — | skipped |
| 6be3e37 | chore(deps): bump github/codeql-action from 4.31.10 to 4.31. | dependabot[bot] | [#160](https://github.com/tektoncd/pruner/pull/160) | — | skipped |
| 8934034 | chore(deps): bump step-security/harden-runner from 2.14.0 to | dependabot[bot] | [#161](https://github.com/tektoncd/pruner/pull/161) | — | skipped |
| cfef219 | chore(ci): update cherry-pick workflow to fix multi-commit P | Vincent Demeester | [#162](https://github.com/tektoncd/pruner/pull/162) | — | — |
| fa2dd03 | chore(deps): bump github/codeql-action from 4.31.11 to 4.32. | dependabot[bot] | [#163](https://github.com/tektoncd/pruner/pull/163) | — | skipped |
| 355ebeb | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#164](https://github.com/tektoncd/pruner/pull/164) | — | skipped |
| 9e54164 | chore(deps): bump actions/cache from 5.0.2 to 5.0.3 | dependabot[bot] | [#166](https://github.com/tektoncd/pruner/pull/166) | — | skipped |
| a2a8f07 | chore(deps): bump chainguard-dev/actions from 1.5.13 to 1.5. | dependabot[bot] | [#167](https://github.com/tektoncd/pruner/pull/167) | — | skipped |
| ff8d389 | chore(deps): bump go.opentelemetry.io/otel from 1.39.0 to 1. | dependabot[bot] | [#169](https://github.com/tektoncd/pruner/pull/169) | — | skipped |
| 7e15ff4 | docs: remove tekton-* and add tekton-operator | Aditya Shinde | [#170](https://github.com/tektoncd/pruner/pull/170) | [SRVKP-9677](https://redhat.atlassian.net/browse/SRVKP-9677) *(keyword)* | **Release Notes Not Required** *(extracted)* |
| 84ade14 | docs: improve punctuation | Ankur Sinha | [#171](https://github.com/tektoncd/pruner/pull/171) | — | — |
| 3a48be1 | chore(ci): include retest in slash workflow | Anitha Natarajan | [#172](https://github.com/tektoncd/pruner/pull/172) | — | — |
| f83b550 | chore(deps): bump github/codeql-action from 4.32.0 to 4.32.2 | dependabot[bot] | [#173](https://github.com/tektoncd/pruner/pull/173) | — | skipped |
| d0a5bda | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#174](https://github.com/tektoncd/pruner/pull/174) | — | skipped |
| 112bf05 | chore(deps): bump chainguard-dev/actions from 1.5.14 to 1.5. | dependabot[bot] | [#175](https://github.com/tektoncd/pruner/pull/175) | — | skipped |
| 89cedab | chore(deps): bump step-security/harden-runner from 2.14.1 to | dependabot[bot] | [#176](https://github.com/tektoncd/pruner/pull/176) | — | skipped |
| a94b31c | ci: add unified CI summary job | Vincent Demeester | [#182](https://github.com/tektoncd/pruner/pull/182) | — | — |
| 69ea00b | fix: invoke e2e test workflow from the ci workflow | Anitha Natarajan | [#189](https://github.com/tektoncd/pruner/pull/189) | — | — |
| 514bbf7 | fix: doc typos | Naomi Gelman | [#183](https://github.com/tektoncd/pruner/pull/183) | — | — |
| d3c6f36 | chore: move architecture diagram to root | Shubham Bhardwaj | [#188](https://github.com/tektoncd/pruner/pull/188) | — | — |
| 612d841 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#179](https://github.com/tektoncd/pruner/pull/179) | — | skipped |
| 4777812 | chore(deps): bump chainguard-dev/actions from 1.5.16 to 1.6. | dependabot[bot] | [#185](https://github.com/tektoncd/pruner/pull/185) | — | skipped |
| 9c3863e | chore(deps): bump github/codeql-action from 4.32.2 to 4.32.4 | dependabot[bot] | [#186](https://github.com/tektoncd/pruner/pull/186) | — | skipped |
| f4cf4e7 | chore(deps): bump actions/dependency-review-action from 4.8. | dependabot[bot] | [#187](https://github.com/tektoncd/pruner/pull/187) | — | skipped |
| 704989b | chore(deps): bump tj-actions/changed-files from 47.0.1 to 47 | dependabot[bot] | [#184](https://github.com/tektoncd/pruner/pull/184) | — | skipped |
| bd36c18 | chore(deps): bump github.com/tektoncd/pipeline from 1.7.0 to | dependabot[bot] | [#165](https://github.com/tektoncd/pruner/pull/165) | — | skipped |
| 9e45a7c | chore(deps): replace go.opentelemetry.io/otel/sdk v1.39.0 wi | Andrew Thorp | [#165](https://github.com/tektoncd/pruner/pull/165) | — | — |
| 6d82a7e | chore(deps): bump github.com/tektoncd/pipeline from 1.9.0 to | dependabot[bot] | [#190](https://github.com/tektoncd/pruner/pull/190) | — | skipped |
| 79ae7c3 | chore(deps): bump actions/setup-go from 6.2.0 to 6.3.0 | dependabot[bot] | [#193](https://github.com/tektoncd/pruner/pull/193) | — | skipped |
| d0bd4f7 | chore(deps): bump chainguard-dev/actions from 1.6.3 to 1.6.5 | dependabot[bot] | [#196](https://github.com/tektoncd/pruner/pull/196) | — | skipped |
| 19b8b17 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#194](https://github.com/tektoncd/pruner/pull/194) | — | skipped |
| 5b29bc1 | chore(deps): bump actions/upload-artifact from 6.0.0 to 7.0. | dependabot[bot] | [#192](https://github.com/tektoncd/pruner/pull/192) | — | skipped |
| 7860fcb | chore(deps): bump step-security/harden-runner from 2.14.2 to | dependabot[bot] | [#195](https://github.com/tektoncd/pruner/pull/195) | — | skipped |
| 4970aaf | Move inactive approvers/reviewers to alumni | Vincent Demeester | [#202](https://github.com/tektoncd/pruner/pull/202) | [SRVKP-7687](https://redhat.atlassian.net/browse/SRVKP-7687) | **Release Notes Not Required** *(extracted)* |
| 1245588 | chore(deps): bump go.opentelemetry.io/otel/metric from 1.40. | dependabot[bot] | [#197](https://github.com/tektoncd/pruner/pull/197) | — | skipped |
| c27d693 | chore(deps): bump github.com/docker/cli | dependabot[bot] | [#200](https://github.com/tektoncd/pruner/pull/200) | — | skipped |
| f65ecd8 | chore(deps): bump actions/dependency-review-action from 4.8. | dependabot[bot] | [#207](https://github.com/tektoncd/pruner/pull/207) | — | skipped |
| 49522aa | chore(deps): bump tj-actions/changed-files from 47.0.4 to 47 | dependabot[bot] | [#209](https://github.com/tektoncd/pruner/pull/209) | — | skipped |
| 20f8ede | chore(deps): bump github/codeql-action from 4.32.4 to 4.32.6 | dependabot[bot] | [#205](https://github.com/tektoncd/pruner/pull/205) | — | skipped |
| a79b109 | chore(deps): bump chainguard-dev/actions from 1.6.5 to 1.6.7 | dependabot[bot] | [#210](https://github.com/tektoncd/pruner/pull/210) | — | skipped |
| abac1dc | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#204](https://github.com/tektoncd/pruner/pull/204) | — | skipped |
| b67a2c4 | chore(deps): bump step-security/harden-runner from 2.15.0 to | dependabot[bot] | [#203](https://github.com/tektoncd/pruner/pull/203) | — | skipped |
| 629fcb9 | chore(deps): bump fgrosse/go-coverage-report from 1.2.0 to 1 | dependabot[bot] | [#211](https://github.com/tektoncd/pruner/pull/211) | — | skipped |
| b285dd9 | chore(deps): bump step-security/harden-runner from 2.15.1 to | dependabot[bot] | [#214](https://github.com/tektoncd/pruner/pull/214) | — | skipped |
| c980e67 | chore(deps): bump chainguard-dev/actions from 1.6.7 to 1.6.8 | dependabot[bot] | [#212](https://github.com/tektoncd/pruner/pull/212) | — | skipped |
| 7868875 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#213](https://github.com/tektoncd/pruner/pull/213) | — | skipped |
| 5e421ca | chore(deps): bump go.opentelemetry.io/otel/metric from 1.41. | dependabot[bot] | [#206](https://github.com/tektoncd/pruner/pull/206) | — | skipped |
| 85d2096 | chore(deps): bump google.golang.org/grpc from 1.77.0 to 1.79 | dependabot[bot] | [#215](https://github.com/tektoncd/pruner/pull/215) | — | skipped |
| 4fd54ab | chore(deps): bump github/codeql-action from 4.32.6 to 4.34.1 | dependabot[bot] | [#216](https://github.com/tektoncd/pruner/pull/216) | — | skipped |
| 06c65c5 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#217](https://github.com/tektoncd/pruner/pull/217) | — | skipped |
| 6d325ba | chore(deps): bump actions/cache from 5.0.3 to 5.0.4 | dependabot[bot] | [#218](https://github.com/tektoncd/pruner/pull/218) | — | skipped |
| 7549b17 | chore(deps): bump chainguard-dev/actions from 1.6.8 to 1.6.9 | dependabot[bot] | [#219](https://github.com/tektoncd/pruner/pull/219) | — | skipped |
| 379afd0 | Bump pipeline version & knative.dev/pkg version | Shubham Bhardwaj | [#220](https://github.com/tektoncd/pruner/pull/220) | [SRVKP-9325](https://redhat.atlassian.net/browse/SRVKP-9325) | **Release Notes Not Required** *(extracted)* |
| 83c2eb2 | docs: clean up observability config and update metrics docum | Shubham Bhardwaj | [#221](https://github.com/tektoncd/pruner/pull/221) | — | — |
| bf82dcb | chore(deps): bump k8s.io/api from 0.35.2 to 0.35.3 | dependabot[bot] | [#225](https://github.com/tektoncd/pruner/pull/225) | — | skipped |
| 317bf5a | chore(deps): bump k8s.io/client-go from 0.35.2 to 0.35.3 | dependabot[bot] | [#224](https://github.com/tektoncd/pruner/pull/224) | — | skipped |
| 454def6 | chore(deps): bump k8s.io/code-generator from 0.35.2 to 0.35. | dependabot[bot] | [#222](https://github.com/tektoncd/pruner/pull/222) | — | skipped |
| af2c740 | Bump knative.dev/pkg to adopt centralized WEBHOOK_* TLS conf | Jawed khelil | [#231](https://github.com/tektoncd/pruner/pull/231) | [SRVKP-9644](https://redhat.atlassian.net/browse/SRVKP-9644) *(keyword)* | **Enhancement:** The pruner now supports centralized WEBHOOK_* TLS configuration from knative.dev *(generated)* |
| 06a1fb3 | chore(deps): bump chainguard-dev/actions from 1.6.9 to 1.6.1 | dependabot[bot] | [#227](https://github.com/tektoncd/pruner/pull/227) | — | skipped |
| 1005f06 | chore(deps): bump actions/setup-go from 6.3.0 to 6.4.0 | dependabot[bot] | [#228](https://github.com/tektoncd/pruner/pull/228) | — | skipped |
| d99c25a | chore(deps): bump github/codeql-action from 4.34.1 to 4.35.1 | dependabot[bot] | [#229](https://github.com/tektoncd/pruner/pull/229) | — | skipped |
| 7cdc053 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#230](https://github.com/tektoncd/pruner/pull/230) | — | skipped |
| abf13a2 | chore(deps): bump github.com/tektoncd/pipeline from 1.10.2 t | dependabot[bot] | [#234](https://github.com/tektoncd/pruner/pull/234) | — | skipped |
| d488086 | add scorecard workflow | Tyler Auerbeck | [#226](https://github.com/tektoncd/pruner/pull/226) | — | — |
| 18ffdcc | fix cherry-pick failure | Anitha Natarajan | [#236](https://github.com/tektoncd/pruner/pull/236) | — | — |
| b51ee26 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#237](https://github.com/tektoncd/pruner/pull/237) | — | skipped |
| d4ca138 | chore(deps): bump go.opentelemetry.io/otel/metric from 1.42. | dependabot[bot] | [#238](https://github.com/tektoncd/pruner/pull/238) | — | skipped |
| c550eef | chore(deps): bump github/codeql-action from 4.34.1 to 4.35.1 | dependabot[bot] | [#239](https://github.com/tektoncd/pruner/pull/239) | — | skipped |
| e456d67 | chore(deps): bump step-security/harden-runner from 2.16.0 to | dependabot[bot] | [#241](https://github.com/tektoncd/pruner/pull/241) | — | skipped |
| 7f9658b | chore(deps): bump chainguard-dev/actions from 1.6.11 to 1.6. | dependabot[bot] | [#242](https://github.com/tektoncd/pruner/pull/242) | — | skipped |
| 309e1ec | chore(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#245](https://github.com/tektoncd/pruner/pull/245) | — | skipped |
| deb45c7 | chore(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#243](https://github.com/tektoncd/pruner/pull/243) | — | skipped |
| 38139ba | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#249](https://github.com/tektoncd/pruner/pull/249) | — | skipped |
| c336722 | chore(deps): bump step-security/harden-runner from 2.16.1 to | dependabot[bot] | [#247](https://github.com/tektoncd/pruner/pull/247) | — | skipped |
| 1f0999f | chore(deps): bump chainguard-dev/actions from 1.6.13 to 1.6. | dependabot[bot] | [#248](https://github.com/tektoncd/pruner/pull/248) | — | skipped |
| 586ce85 | chore(deps): bump actions/upload-artifact from 7.0.0 to 7.0. | dependabot[bot] | [#250](https://github.com/tektoncd/pruner/pull/250) | — | skipped |
| 1eb087e | chore(deps): bump k8s.io/client-go from 0.35.3 to 0.35.4 | dependabot[bot] | [#252](https://github.com/tektoncd/pruner/pull/252) | — | skipped |
| 748c0bb | chore(deps): bump k8s.io/code-generator from 0.35.3 to 0.35. | dependabot[bot] | [#254](https://github.com/tektoncd/pruner/pull/254) | — | skipped |
| cfd5864 | chore(deps): bump chainguard-dev/actions from 1.6.14 to 1.6. | dependabot[bot] | [#255](https://github.com/tektoncd/pruner/pull/255) | — | skipped |
| 47014e2 | chore(deps): bump actions/cache from 5.0.4 to 5.0.5 | dependabot[bot] | [#260](https://github.com/tektoncd/pruner/pull/260) | — | skipped |
| 9d3ca12 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#256](https://github.com/tektoncd/pruner/pull/256) | — | skipped |
| cd674fd | chore(deps): bump step-security/harden-runner from 2.17.0 to | dependabot[bot] | [#258](https://github.com/tektoncd/pruner/pull/258) | — | skipped |
| 7ad14a6 | chore(deps): bump github/codeql-action from 4.35.1 to 4.35.2 | dependabot[bot] | [#259](https://github.com/tektoncd/pruner/pull/259) | — | skipped |
| 84d335d | chore(deps): bump tj-actions/changed-files from 47.0.5 to 47 | dependabot[bot] | [#257](https://github.com/tektoncd/pruner/pull/257) | — | skipped |
| 3282774 | chore(deps): bump github.com/tektoncd/pipeline from 1.11.0 t | dependabot[bot] | [#261](https://github.com/tektoncd/pruner/pull/261) | — | skipped |
| d224832 | chore(deps): bump chainguard-dev/actions from 1.6.15 to 1.6. | dependabot[bot] | [#264](https://github.com/tektoncd/pruner/pull/264) | — | skipped |
| d3302cd | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#265](https://github.com/tektoncd/pruner/pull/265) | — | skipped |
| 528d612 | chore(deps): bump go.uber.org/zap from 1.27.1 to 1.28.0 | dependabot[bot] | [#266](https://github.com/tektoncd/pruner/pull/266) | — | skipped |
| f0eb10d | chore: add zizmor workflow and fix GitHub Actions security f | Shubham Bhardwaj | [#262](https://github.com/tektoncd/pruner/pull/262) | [SRVKP-12065](https://redhat.atlassian.net/browse/SRVKP-12065) *(keyword)* | **Release Notes Not Required** *(generated)* |
| c7396a3 | chore(deps): bump chainguard-dev/actions from 1.6.16 to 1.6. | dependabot[bot] | [#272](https://github.com/tektoncd/pruner/pull/272) | — | skipped |
| 3d07863 | chore(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#268](https://github.com/tektoncd/pruner/pull/268) | — | skipped |
| 563a1b4 | chore(deps): bump step-security/harden-runner from 2.19.0 to | dependabot[bot] | [#271](https://github.com/tektoncd/pruner/pull/271) | — | skipped |
| 6f2c3cc | chore(deps): bump github.com/tektoncd/pipeline from 1.11.1 t | dependabot[bot] | [#270](https://github.com/tektoncd/pruner/pull/270) | — | skipped |
| 87d476f | chore(deps): bump github/codeql-action from 4.35.2 to 4.35.3 | dependabot[bot] | [#269](https://github.com/tektoncd/pruner/pull/269) | — | skipped |
| 8e317b9 | tekton: automate releases with Pipelines-as-Code | Shubham Bhardwaj | [#246](https://github.com/tektoncd/pruner/pull/246) | [SRVKP-11528](https://redhat.atlassian.net/browse/SRVKP-11528) | **Release Notes Not Required** *(generated)* |
| a01f51e | fix: update release pipeline for oci-ci-cd cluster compatibi | Shubham Bhardwaj | [#284](https://github.com/tektoncd/pruner/pull/284) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 1dacc8b | fix: add e2e build tag to skip e2e tests in release pipeline | Shubham Bhardwaj | [#289](https://github.com/tektoncd/pruner/pull/289) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 62c4f57 | ci: add updated pipelines version and k8s version | Shubham Bhardwaj | [#290](https://github.com/tektoncd/pruner/pull/290) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 2d108c1 | feat: add draft release and Chains signing to release pipeli | Shubham Bhardwaj | [#296](https://github.com/tektoncd/pruner/pull/296) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 10ac09c | ci: diable automatic dependency review workflow | Shubham Bhardwaj | [#297](https://github.com/tektoncd/pruner/pull/297) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |

---

## tektoncd-results

**Upstream:** tektoncd/results · **Commits:** 116 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 34feca6 | Bump golang.org/x/oauth2 from 0.34.0 to 0.35.0 | dependabot[bot] | [#1199](https://github.com/tektoncd/results/pull/1199) | — | skipped |
| 4207025 | Bump github.com/tektoncd/cli from 0.41.1 to 0.43.0 | dependabot[bot] | [#1200](https://github.com/tektoncd/results/pull/1200) | — | skipped |
| 53ae6b3 | Bump github.com/aws/aws-sdk-go-v2/config from 1.32.2 to 1.32 | dependabot[bot] | [#1202](https://github.com/tektoncd/results/pull/1202) | — | skipped |
| da7af72 | Bump github.com/mattn/go-sqlite3 from 1.14.32 to 1.14.33 | dependabot[bot] | [#1201](https://github.com/tektoncd/results/pull/1201) | — | skipped |
| 47b3dcd | Bump github.com/grpc-ecosystem/grpc-gateway/v2 from 2.27.3 t | dependabot[bot] | [#1203](https://github.com/tektoncd/results/pull/1203) | — | skipped |
| 989ee98 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.93.0 to  | dependabot[bot] | [#1205](https://github.com/tektoncd/results/pull/1205) | — | skipped |
| 81d0aa2 | Bump k8s.io/client-go from 0.34.2 to 0.34.3 | dependabot[bot] | [#1207](https://github.com/tektoncd/results/pull/1207) | — | skipped |
| be8673b | Bump google.golang.org/api from 0.265.0 to 0.266.0 | dependabot[bot] | [#1209](https://github.com/tektoncd/results/pull/1209) | — | skipped |
| 0231e46 | Bump github.com/grpc-ecosystem/grpc-gateway/v2 from 2.27.7 t | dependabot[bot] | [#1206](https://github.com/tektoncd/results/pull/1206) | — | skipped |
| c636bed | Bump golang.org/x/net from 0.49.0 to 0.50.0 | dependabot[bot] | [#1211](https://github.com/tektoncd/results/pull/1211) | — | skipped |
| 3983e49 | Bump github.com/spf13/cobra from 1.10.1 to 1.10.2 | dependabot[bot] | [#1212](https://github.com/tektoncd/results/pull/1212) | — | skipped |
| ec562fe | Bump github.com/grpc-ecosystem/go-grpc-middleware/v2 from 2. | dependabot[bot] | [#1213](https://github.com/tektoncd/results/pull/1213) | — | skipped |
| 485ede0 | Bump k8s.io/apiserver from 0.32.11 to 0.32.12 | dependabot[bot] | [#1214](https://github.com/tektoncd/results/pull/1214) | — | skipped |
| 87ce923 | Bump google.golang.org/grpc from 1.78.0 to 1.79.1 | dependabot[bot] | [#1216](https://github.com/tektoncd/results/pull/1216) | — | skipped |
| fb0c378 | Bump github.com/aws/aws-sdk-go-v2/credentials from 1.19.7 to | dependabot[bot] | [#1217](https://github.com/tektoncd/results/pull/1217) | — | skipped |
| 80cd080 | Bump github.com/grpc-ecosystem/grpc-gateway/v2 from 2.27.8 t | dependabot[bot] | [#1219](https://github.com/tektoncd/results/pull/1219) | — | skipped |
| 5254a97 | Bump google.golang.org/api from 0.266.0 to 0.267.0 | dependabot[bot] | [#1220](https://github.com/tektoncd/results/pull/1220) | — | skipped |
| f03c4dd | Bump github.com/aws/aws-sdk-go-v2/config from 1.32.7 to 1.32 | dependabot[bot] | [#1218](https://github.com/tektoncd/results/pull/1218) | — | skipped |
| 13a1a07 | Bump github.com/aws/aws-sdk-go-v2/config from 1.32.8 to 1.32 | dependabot[bot] | [#1221](https://github.com/tektoncd/results/pull/1221) | — | skipped |
| e755414 | Bump filippo.io/edwards25519 from 1.1.0 to 1.1.1 | dependabot[bot] | [#1223](https://github.com/tektoncd/results/pull/1223) | — | skipped |
| 7f6d2f5 | Bump github.com/tektoncd/pipeline from 1.9.0 to 1.9.1 (#1224 | dependabot[bot] | [#1224](https://github.com/tektoncd/results/pull/1224) | — | skipped |
| 1e1a0ec | feature: update release cheatsheet and pipeline (#1176) | Anitha Natarajan | [#1176](https://github.com/tektoncd/results/pull/1176) | — | — |
| 78c39a2 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.96.0 to  | dependabot[bot] | [#1226](https://github.com/tektoncd/results/pull/1226) | — | skipped |
| 17bb946 | Bump google.golang.org/api from 0.267.0 to 0.268.0 | dependabot[bot] | [#1230](https://github.com/tektoncd/results/pull/1230) | — | skipped |
| b490d2c | Bump github.com/tektoncd/cli from 0.43.0 to 0.44.0 | dependabot[bot] | [#1227](https://github.com/tektoncd/results/pull/1227) | — | skipped |
| 088a8b1 | Bump github.com/aws/aws-sdk-go-v2/config from 1.32.9 to 1.32 | dependabot[bot] | [#1228](https://github.com/tektoncd/results/pull/1228) | — | skipped |
| 2bc5103 | Bump google.golang.org/api from 0.268.0 to 0.269.0 | dependabot[bot] | [#1232](https://github.com/tektoncd/results/pull/1232) | — | skipped |
| 660e736 | Bump k8s.io/client-go from 0.34.4 to 0.34.5 (#1237) | dependabot[bot] | [#1237](https://github.com/tektoncd/results/pull/1237) | — | skipped |
| 754f3b4 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.96.1 to  | dependabot[bot] | [#1235](https://github.com/tektoncd/results/pull/1235) | — | skipped |
| 5985828 | Bump golang.org/x/net from 0.50.0 to 0.51.0 | dependabot[bot] | [#1234](https://github.com/tektoncd/results/pull/1234) | — | skipped |
| 09d567a | Bump go.opentelemetry.io/otel/sdk from 1.39.0 to 1.40.0 | dependabot[bot] | [#1238](https://github.com/tektoncd/results/pull/1238) | — | skipped |
| bfa382f | Bump gocloud.dev from 0.44.0 to 0.45.0 | dependabot[bot] | [#1239](https://github.com/tektoncd/results/pull/1239) | — | skipped |
| 259477b | Bump k8s.io/apiserver from 0.32.12 to 0.32.13 (#1240) | dependabot[bot] | [#1240](https://github.com/tektoncd/results/pull/1240) | — | skipped |
| 753e219 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.96.2 to  | dependabot[bot] | [#1241](https://github.com/tektoncd/results/pull/1241) | — | skipped |
| e72d401 | Bump github.com/aws/aws-sdk-go-v2/config from 1.32.10 to 1.3 | dependabot[bot] | [#1242](https://github.com/tektoncd/results/pull/1242) | — | skipped |
| 1183335 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.96.3 to  | dependabot[bot] | [#1245](https://github.com/tektoncd/results/pull/1245) | — | skipped |
| 1d4f1c7 | Bump google.golang.org/grpc from 1.79.1 to 1.79.2 | dependabot[bot] | [#1248](https://github.com/tektoncd/results/pull/1248) | — | skipped |
| 2a7e236 | Bump golang.org/x/oauth2 from 0.35.0 to 0.36.0 | dependabot[bot] | [#1250](https://github.com/tektoncd/results/pull/1250) | — | skipped |
| d1e371b | Bump google.golang.org/api from 0.269.0 to 0.270.0 | dependabot[bot] | [#1251](https://github.com/tektoncd/results/pull/1251) | — | skipped |
| 2b493c4 | Nominate divyansh42 as results approver and reviewer | Vincent Demeester | [#1247](https://github.com/tektoncd/results/pull/1247) | — | — |
| 7ac3dc4 | ci: Add CI summary fan-in job to presubmit workflow | Vincent Demeester | [#1225](https://github.com/tektoncd/results/pull/1225) | — | — |
| 1907cb6 | CI: Remove manual PR author check from presubmit workflow | Khurram Baig | [#1197](https://github.com/tektoncd/results/pull/1197) | — | — |
| 511602c | Bump google.golang.org/api from 0.270.0 to 0.271.0 | dependabot[bot] | [#1253](https://github.com/tektoncd/results/pull/1253) | — | skipped |
| fcf7dd4 | Improve finalizer error visibility | Emil Natan | [#1252](https://github.com/tektoncd/results/pull/1252) | — | — |
| 876ccee | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.96.4 to  | dependabot[bot] | [#1255](https://github.com/tektoncd/results/pull/1255) | — | skipped |
| 439af7d | Bump golang.org/x/net from 0.51.0 to 0.52.0 | dependabot[bot] | [#1254](https://github.com/tektoncd/results/pull/1254) | — | skipped |
| 8f3629e | Change all occurences of GCS buckets with OCI buckets | adityavshinde | [#1257](https://github.com/tektoncd/results/pull/1257) | [SRVKP-11090](https://redhat.atlassian.net/browse/SRVKP-11090) | **Release Notes Not Required** *(generated)* |
| 1bc1349 | Revert OCI buckets with GCS buckets which was mistakenly rep | adityavshinde | [#1257](https://github.com/tektoncd/results/pull/1257) | [SRVKP-11090](https://redhat.atlassian.net/browse/SRVKP-11090) | **Release Notes Not Required** *(generated)* |
| f1ea6a8 | Revert changed Dockerfile | adityavshinde | [#1257](https://github.com/tektoncd/results/pull/1257) | [SRVKP-11090](https://redhat.atlassian.net/browse/SRVKP-11090) | **Release Notes Not Required** *(generated)* |
| d0c82f6 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.97.0 to  | dependabot[bot] | [#1259](https://github.com/tektoncd/results/pull/1259) | — | skipped |
| 6313aa8 | Bump github.com/aws/aws-sdk-go-v2/config from 1.32.11 to 1.3 | dependabot[bot] | [#1260](https://github.com/tektoncd/results/pull/1260) | — | skipped |
| 1a3ba97 | Bump google.golang.org/api from 0.271.0 to 0.272.0 (#1264) | dependabot[bot] | [#1264](https://github.com/tektoncd/results/pull/1264) | — | skipped |
| 21490b9 | Bump github.com/mattn/go-sqlite3 from 1.14.34 to 1.14.37 | dependabot[bot] | [#1263](https://github.com/tektoncd/results/pull/1263) | — | skipped |
| 9dabb77 | Move inactive approvers and reviewers to alumni | Vincent Demeester | [#1246](https://github.com/tektoncd/results/pull/1246) | — | — |
| d5ba834 | Bump google.golang.org/grpc from 1.79.2 to 1.79.3 | dependabot[bot] | [#1266](https://github.com/tektoncd/results/pull/1266) | — | skipped |
| 8e189eb | Bump k8s.io/apimachinery from 0.34.5 to 0.34.6 | dependabot[bot] | [#1267](https://github.com/tektoncd/results/pull/1267) | — | skipped |
| 72a9644 | Bump k8s.io/client-go from 0.34.5 to 0.34.6 | dependabot[bot] | [#1269](https://github.com/tektoncd/results/pull/1269) | — | skipped |
| 0e0cd1e | Bump github.com/fatih/color from 1.18.0 to 1.19.0 | dependabot[bot] | [#1270](https://github.com/tektoncd/results/pull/1270) | — | skipped |
| 88f69a2 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.97.1 to  | dependabot[bot] | [#1271](https://github.com/tektoncd/results/pull/1271) | — | skipped |
| 822614c | Bump github.com/jackc/pgx/v5 from 5.8.0 to 5.9.1 | dependabot[bot] | [#1272](https://github.com/tektoncd/results/pull/1272) | — | skipped |
| e047a25 | Bump github.com/tektoncd/pipeline from 1.9.1 to 1.9.2 | dependabot[bot] | [#1273](https://github.com/tektoncd/results/pull/1273) | — | skipped |
| d150199 | Bump github.com/aws/aws-sdk-go-v2/config from 1.32.12 to 1.3 | dependabot[bot] | [#1276](https://github.com/tektoncd/results/pull/1276) | — | skipped |
| 65e32b8 | Bump google.golang.org/api from 0.272.0 to 0.273.0 | dependabot[bot] | [#1277](https://github.com/tektoncd/results/pull/1277) | — | skipped |
| e06f7f0 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.97.2 to  | dependabot[bot] | [#1279](https://github.com/tektoncd/results/pull/1279) | — | skipped |
| a4eacc0 | fix(watcher): fix TestController for client-go v0.35 | Khurram Baig | [#1281](https://github.com/tektoncd/results/pull/1281) | — | — |
| 007212d | chore: update deprecated grpc_middleware components | Tyler Auerbeck | [#1274](https://github.com/tektoncd/results/pull/1274) | — | — |
| e4f7918 | Bump github.com/mattn/go-sqlite3 from 1.14.37 to 1.14.38 | dependabot[bot] | [#1282](https://github.com/tektoncd/results/pull/1282) | — | skipped |
| 9efac70 | Bump github.com/aws/aws-sdk-go-v2/service/s3 from 1.97.3 to  | dependabot[bot] | [#1284](https://github.com/tektoncd/results/pull/1284) | — | skipped |
| 3e25269 | Bump google.golang.org/api from 0.273.0 to 0.273.1 | dependabot[bot] | [#1286](https://github.com/tektoncd/results/pull/1286) | — | skipped |
| e8fef26 | Bump google.golang.org/grpc from 1.79.3 to 1.80.0 | dependabot[bot] | [#1287](https://github.com/tektoncd/results/pull/1287) | — | skipped |
| 768e0b3 | chore(deps): bump github.com/go-jose/go-jose/v4 from 4.1.3 t | dependabot[bot] | [#1291](https://github.com/tektoncd/results/pull/1291) | — | skipped |
| 5881c19 | chore(deps): bump github.com/mattn/go-sqlite3 from 1.14.38 t | dependabot[bot] | [#1292](https://github.com/tektoncd/results/pull/1292) | — | skipped |
| c9c8fc6 | chore(deps): bump google.golang.org/api from 0.273.1 to 0.27 | dependabot[bot] | [#1288](https://github.com/tektoncd/results/pull/1288) | — | skipped |
| c9adc97 | chore(deps): bump github.com/aws/aws-sdk-go-v2/config (#1290 | dependabot[bot] | [#1290](https://github.com/tektoncd/results/pull/1290) | — | skipped |
| a7e6d6c | chore(deps): bump go.opentelemetry.io/otel/sdk from 1.42.0 t | dependabot[bot] | [#1300](https://github.com/tektoncd/results/pull/1300) | — | skipped |
| 028658d | chore(deps): bump github.com/google/cel-go from 0.27.0 to 0. | dependabot[bot] | [#1298](https://github.com/tektoncd/results/pull/1298) | — | skipped |
| cd36c69 | chore(deps): bump golang.org/x/net from 0.52.0 to 0.53.0 | dependabot[bot] | [#1302](https://github.com/tektoncd/results/pull/1302) | — | skipped |
| 428d2ba | chore(deps): bump google.golang.org/api from 0.274.0 to 0.27 | dependabot[bot] | [#1297](https://github.com/tektoncd/results/pull/1297) | — | skipped |
| 2d06648 | chore(deps): bump github.com/tektoncd/cli from 0.44.0 to 0.4 | dependabot[bot] | [#1296](https://github.com/tektoncd/results/pull/1296) | — | skipped |
| af03e66 | ci: fix GitHub Actions security issues found by zizmor | ab-ghosh | [#1293](https://github.com/tektoncd/results/pull/1293) | — | — |
| b5ff73d | ci: add zizmor GitHub Actions security analysis | ab-ghosh | [#1293](https://github.com/tektoncd/results/pull/1293) | — | — |
| 4362545 | ci: fix remaining zizmor findings (template injection) | ab-ghosh | [#1293](https://github.com/tektoncd/results/pull/1293) | — | — |
| 273820c | chore(deps): bump k8s.io/client-go from 0.34.6 to 0.34.7 | dependabot[bot] | [#1305](https://github.com/tektoncd/results/pull/1305) | — | skipped |
| 2b40f90 | chore(deps): Update dependencies for OpenTelemetry migration | divyansh42 | [#1249](https://github.com/tektoncd/results/pull/1249) | [SRVKP-8533](https://redhat.atlassian.net/browse/SRVKP-8533) | **Release Notes Not Required** *(generated)* |
| 9b6bf42 | feat(metrics): Migrate from OpenCensus to OpenTelemetry | divyansh42 | [#1249](https://github.com/tektoncd/results/pull/1249) | [SRVKP-8533](https://redhat.atlassian.net/browse/SRVKP-8533) | **Enhancement:** Migrates Results Watcher metrics from OpenCensus to OpenTelemetry. The observabi *(extracted)* |
| c7facf8 | chore(deps): bump github.com/aws/aws-sdk-go-v2/service/s3 | dependabot[bot] | [#1311](https://github.com/tektoncd/results/pull/1311) | — | skipped |
| 3980384 | chore(deps): bump github.com/grpc-ecosystem/grpc-gateway/v2 | dependabot[bot] | [#1310](https://github.com/tektoncd/results/pull/1310) | — | skipped |
| baac5b4 | chore(deps): bump github.com/tektoncd/pipeline from 1.11.0 t | dependabot[bot] | [#1312](https://github.com/tektoncd/results/pull/1312) | — | skipped |
| b3706b2 | chore(cli): deprecate legacy commands and flags | divyansh42 | [#1301](https://github.com/tektoncd/results/pull/1301) | [SRVKP-11434](https://redhat.atlassian.net/browse/SRVKP-11434) | **Deprecated Functionality:** Deprecates legacy CLI commands (`result`, `records`, `logs`) and several global   |
| 3883355 | chore(deps): bump google.golang.org/api from 0.274.0 to 0.27 | dependabot[bot] | [#1303](https://github.com/tektoncd/results/pull/1303) | — | skipped |
| c73b3fb | fix(config): trim whitespace from host, token, and API path  | divyansh42 | [#1314](https://github.com/tektoncd/results/pull/1314) | — | — |
| 3d4e84e | Bump Kind k8s version to 1.35 | Emil Natan | [#1319](https://github.com/tektoncd/results/pull/1319) | — | — |
| 3f99f11 | perf(cli): use exact match for describe/logs command | divyansh42 | [#1283](https://github.com/tektoncd/results/pull/1283) | [SRVKP-11380](https://redhat.atlassian.net/browse/SRVKP-11380) | **Enhancement:** Enhacement: Improved performance for `describe` and `logs` commands with optimiz  |
| 8a63de5 | fix(logs): use ListRecords for dev CLI log listing | divyansh42 | [#1306](https://github.com/tektoncd/results/pull/1306) | [SRVKP-11301](https://redhat.atlassian.net/browse/SRVKP-11301) | **Bug Fix:** Fix `logs list` command failing with "unknown service tekton.results.v1alpha2.Lo  |
| 9f3c6f4 | chore(deps): bump k8s.io/client-go from 0.35.3 to 0.35.4 | dependabot[bot] | [#1315](https://github.com/tektoncd/results/pull/1315) | — | skipped |
| dfcdfea | chore(deps): bump github.com/jackc/pgx/v5 from 5.9.1 to 5.9. | dependabot[bot] | [#1309](https://github.com/tektoncd/results/pull/1309) | — | skipped |
| cfb50ff | chore(deps): bump github.com/mattn/go-sqlite3 from 1.14.40 t | dependabot[bot] | [#1326](https://github.com/tektoncd/results/pull/1326) | — | skipped |
| dcc2e89 | chore(deps): bump github.com/aws/aws-sdk-go-v2/credentials | dependabot[bot] | [#1325](https://github.com/tektoncd/results/pull/1325) | — | skipped |
| 6369b8d | chore(deps): bump github.com/tidwall/gjson from 1.18.0 to 1. | dependabot[bot] | [#1328](https://github.com/tektoncd/results/pull/1328) | — | skipped |
| 3c4158d | chore(deps): bump github.com/tektoncd/cli | dependabot[bot] | [#1333](https://github.com/tektoncd/results/pull/1333) | — | skipped |
| 2e08db1 | Fix typos in API docs | Alan Greene | [#1337](https://github.com/tektoncd/results/pull/1337) | — | — |
| 22ace8d | chore(deps): bump github.com/aws/aws-sdk-go-v2/config | dependabot[bot] | [#1330](https://github.com/tektoncd/results/pull/1330) | — | skipped |
| 48a721a | chore(deps): bump k8s.io/api from 0.35.4 to 0.35.5 | dependabot[bot] | [#1334](https://github.com/tektoncd/results/pull/1334) | — | skipped |
| 8f13082 | chore(deps): bump golang.org/x/net from 0.53.0 to 0.54.0 | dependabot[bot] | [#1338](https://github.com/tektoncd/results/pull/1338) | — | skipped |
| bc26da4 | chore(deps): bump k8s.io/apiserver from 0.35.4 to 0.35.5 | dependabot[bot] | [#1332](https://github.com/tektoncd/results/pull/1332) | — | skipped |
| 589e0cd | feat(logs): added new logql config and documentation | Florian Thiévent | [#1210](https://github.com/tektoncd/results/pull/1210) | — | — |
| 2c65dea | chore(deps): bump github.com/google/cel-go from 0.28.0 to 0. | dependabot[bot] | [#1344](https://github.com/tektoncd/results/pull/1344) | — | skipped |
| 4e3a518 | chore(deps): bump github.com/aws/aws-sdk-go-v2/config | dependabot[bot] | [#1340](https://github.com/tektoncd/results/pull/1340) | — | skipped |
| 9c5c31f | chore(deps): bump go.opentelemetry.io/otel/metric from 1.43. | dependabot[bot] | [#1341](https://github.com/tektoncd/results/pull/1341) | — | skipped |
| ac20152 | Fix race condition in the TestController | Emil Natan | [#1351](https://github.com/tektoncd/results/pull/1351) | — | — |
| e804b0e | chore(deps): bump github.com/aws/aws-sdk-go-v2 from 1.41.7 t | dependabot[bot] | [#1347](https://github.com/tektoncd/results/pull/1347) | — | skipped |
| 19f8019 | chore(deps): bump google.golang.org/api from 0.276.0 to 0.28 | dependabot[bot] | [#1349](https://github.com/tektoncd/results/pull/1349) | — | skipped |
| cc6cc17 | Add agentic workflows context files | Emil Natan | [#1345](https://github.com/tektoncd/results/pull/1345) | [SRVKP-11817](https://redhat.atlassian.net/browse/SRVKP-11817) | **Release Notes Not Required**  |
| 208e14a | Add support for storing CustomRuns | Emil Natan | [#1320](https://github.com/tektoncd/results/pull/1320) | [SRVKP-11855](https://redhat.atlassian.net/browse/SRVKP-11855) *(keyword)* | **Feature:** The Results Watcher now watches and stores CustomRun resources (tekton.dev/v1bet *(extracted)* |
| d3dc464 | chore: switch base image to ghcr.io/tektoncd/plumbing/static | Vincent Demeester | [#1357](https://github.com/tektoncd/results/pull/1357) | — | — |
| ecca90d | Fix release pipelines | Emil Natan | [#1359](https://github.com/tektoncd/results/pull/1359) | — | — |

---

## operator

**Upstream:** tektoncd/operator · **Commits:** 161 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 1015555 | chore(deps): bump step-security/harden-runner from 2.14.2 to | dependabot[bot] | [#3257](https://github.com/tektoncd/operator/pull/3257) | — | skipped |
| fb4751e | chore: update the prechech pipeline version | Anitha Natarajan | [#3261](https://github.com/tektoncd/operator/pull/3261) | — | — |
| 077462a | Move inactive approvers to alumni | Vincent Demeester | [#3267](https://github.com/tektoncd/operator/pull/3267) | — | — |
| a07ab6d | Add TektonScheduler and TektonMulticlusterProxyAAE CRDs to H | mbpavan | [#3265](https://github.com/tektoncd/operator/pull/3265) | — | — |
| 65db10f | chore(deps): bump github.com/docker/cli | dependabot[bot] | [#3263](https://github.com/tektoncd/operator/pull/3263) | — | skipped |
| d7304ec | Nominate pratap0007 as operator approver | Vincent Demeester | [#3268](https://github.com/tektoncd/operator/pull/3268) | — | — |
| 4a40071 | chore(deps): bump github.com/openshift-pipelines/pipelines-a | dependabot[bot] | [#3264](https://github.com/tektoncd/operator/pull/3264) | — | skipped |
| 9f82475 | fix: pass httpClient to PAC SyncConfig | Shiv Verma | [#3264](https://github.com/tektoncd/operator/pull/3264) | — | — |
| c72bc75 | chore(deps): bump github.com/konflux-ci/tekton-kueue | dependabot[bot] | [#3214](https://github.com/tektoncd/operator/pull/3214) | — | skipped |
| 362d123 | chore(deps): bump chainguard-dev/actions from 1.6.5 to 1.6.7 | dependabot[bot] | [#3271](https://github.com/tektoncd/operator/pull/3271) | — | skipped |
| 28d0ffa | chore(deps): bump step-security/harden-runner from 2.15.0 to | dependabot[bot] | [#3272](https://github.com/tektoncd/operator/pull/3272) | — | skipped |
| 7a6f2de | chore(deps): bump github/codeql-action from 4.32.5 to 4.32.6 | dependabot[bot] | [#3273](https://github.com/tektoncd/operator/pull/3273) | — | skipped |
| 8a9a95b | chore(deps): bump golang.org/x/sync from 0.19.0 to 0.20.0 | dependabot[bot] | [#3274](https://github.com/tektoncd/operator/pull/3274) | — | skipped |
| b5bb818 | chore: add pramod as approver for operator | Anitha Natarajan | [#3278](https://github.com/tektoncd/operator/pull/3278) | — | — |
| e8451e5 | fix readme and bundle script update | mbpavan | [#3277](https://github.com/tektoncd/operator/pull/3277) | [SRVKP-11045](https://redhat.atlassian.net/browse/SRVKP-11045) | **Release Notes Not Required** *(extracted)* |
| 4f93b09 | chore: bump hub from v1.23.7 to v1.23.8 | Vincent Demeester | [#3279](https://github.com/tektoncd/operator/pull/3279) | — | — |
| 9e20b9c | Fix TektonInstallerSet deadlock when resources have deletion | Jawed khelil | [#3217](https://github.com/tektoncd/operator/pull/3217) | [SRVKP-10858](https://redhat.atlassian.net/browse/SRVKP-10858) | **Release Note Not Required** *(extracted)* |
| 4a89460 | fix some minor issues | mbpavan | [#3282](https://github.com/tektoncd/operator/pull/3282) | [SRVKP-11089](https://redhat.atlassian.net/browse/SRVKP-11089) | **Bug Fix** *(extracted)* |
| ab0476c | chore(deps): bump step-security/harden-runner from 2.15.1 to | dependabot[bot] | [#3283](https://github.com/tektoncd/operator/pull/3283) | — | skipped |
| 8f72852 | chore(deps): bump github/codeql-action from 4.32.6 to 4.33.0 | dependabot[bot] | [#3284](https://github.com/tektoncd/operator/pull/3284) | — | skipped |
| 1fa6522 | chore: bump component versions | Vincent Demeester | [#3281](https://github.com/tektoncd/operator/pull/3281) | [SRVKP-12488](https://redhat.atlassian.net/browse/SRVKP-12488) *(keyword)* | **Release Notes Not Required:** ** [v1.10.0...v1.10.1](https://github.com/tektoncd/pipeline/compare/v1.10.0...v1 *(extracted)* |
| 9bde28a | chore(deps): bump fgrosse/go-coverage-report from 1.2.0 to 1 | dependabot[bot] | [#3285](https://github.com/tektoncd/operator/pull/3285) | — | skipped |
| 7090a73 | chore(deps): bump golang.org/x/mod from 0.33.0 to 0.34.0 | dependabot[bot] | [#3289](https://github.com/tektoncd/operator/pull/3289) | — | skipped |
| a68f5cc | chore(deps): bump chainguard-dev/actions from 1.6.7 to 1.6.8 | dependabot[bot] | [#3286](https://github.com/tektoncd/operator/pull/3286) | — | skipped |
| 6fd473d | chore(deps): bump github.com/cert-manager/cert-manager | dependabot[bot] | [#3287](https://github.com/tektoncd/operator/pull/3287) | — | skipped |
| 6ef5497 | chore(deps): bump github.com/openshift-pipelines/pipelines-a | dependabot[bot] | [#3288](https://github.com/tektoncd/operator/pull/3288) | — | skipped |
| 122e2a3 | chore: bump pipeline from v1.10.1 to v1.10.2 | Vincent Demeester | [#3291](https://github.com/tektoncd/operator/pull/3291) | — | — |
| dd3b702 | Fix authorization field in pipeline_console_plugin.yaml | Khurram | [#3293](https://github.com/tektoncd/operator/pull/3293) | [SRVKP-11107](https://redhat.atlassian.net/browse/SRVKP-11107) | **Release Notes Not Required** *(extracted)* |
| bdc151d | chore(deps): bump chainguard-dev/actions from 1.6.8 to 1.6.9 | dependabot[bot] | [#3295](https://github.com/tektoncd/operator/pull/3295) | — | skipped |
| 35319ea | chore(deps): bump github/codeql-action from 4.33.0 to 4.34.1 | dependabot[bot] | [#3296](https://github.com/tektoncd/operator/pull/3296) | — | skipped |
| 836566c | chore(deps): bump actions/cache from 5.0.3 to 5.0.4 | dependabot[bot] | [#3297](https://github.com/tektoncd/operator/pull/3297) | — | skipped |
| 7567ea8 | chore(deps): bump github.com/tektoncd/pipeline from 1.9.1 to | dependabot[bot] | [#3299](https://github.com/tektoncd/operator/pull/3299) | — | skipped |
| 2419609 | chore(deps): bump google.golang.org/grpc from 1.79.1 to 1.79 | Anitha Natarajan | [#3302](https://github.com/tektoncd/operator/pull/3302) | — | — |
| 206fdc0 | chore(deps): bump actions/setup-go from 6.3.0 to 6.4.0 | dependabot[bot] | [#3312](https://github.com/tektoncd/operator/pull/3312) | — | skipped |
| 4a11bc1 | fix-cherry-pick-failures | Anitha Natarajan | [#3323](https://github.com/tektoncd/operator/pull/3323) | — | — |
| b72ff08 | chore(deps): bump azure/setup-helm from 4.3.1 to 5.0.0 | dependabot[bot] | [#3313](https://github.com/tektoncd/operator/pull/3313) | — | skipped |
| 4db3f86 | chore(deps): bump github/codeql-action from 4.34.1 to 4.35.1 | dependabot[bot] | [#3314](https://github.com/tektoncd/operator/pull/3314) | — | skipped |
| 5f8f5fb | chore(deps): bump chainguard-dev/actions from 1.6.9 to 1.6.1 | dependabot[bot] | [#3315](https://github.com/tektoncd/operator/pull/3315) | — | skipped |
| 59d9bed | chore(deps): bump github.com/cert-manager/cert-manager | dependabot[bot] | [#3316](https://github.com/tektoncd/operator/pull/3316) | — | skipped |
| 65eda6c | chore: bump component versions | Vincent Demeester | [#3317](https://github.com/tektoncd/operator/pull/3317) | [SRVKP-12488](https://redhat.atlassian.net/browse/SRVKP-12488) *(keyword)* | **Release Notes Not Required:** ** [v0.66.0...v0.67.0](https://github.com/tektoncd/dashboard/compare/v0.66.0...v *(extracted)* |
| 5e5f323 | ci: fix GitHub Actions security issues found by zizmor | Vincent Demeester | [#3324](https://github.com/tektoncd/operator/pull/3324) | — | — |
| 4f9fdfa | ci: fix remaining zizmor findings and add zizmor CI check | Vincent Demeester | [#3324](https://github.com/tektoncd/operator/pull/3324) | — | — |
| 6df6dfb | Fix: Syncer Service Transformer | Pramod Bindal | [#3326](https://github.com/tektoncd/operator/pull/3326) | [SRVKP-11461](https://redhat.atlassian.net/browse/SRVKP-11461) | **Release Notes Not Required** *(extracted)* |
| 1169a40 | tekton: automate releases with Pipelines-as-Code | Vincent Demeester | [#3322](https://github.com/tektoncd/operator/pull/3322) | [SRVKP-11937](https://redhat.atlassian.net/browse/SRVKP-11937) *(keyword)* | **Release Notes Not Required** *(extracted)* |
| dafd94c | tekton: clarify push event only fires on branch creation | Vincent Demeester | [#3322](https://github.com/tektoncd/operator/pull/3322) | — | — |
| 43d2c0d | fix: include branch filter in CEL expression for release tri | Vincent Demeester | [#3335](https://github.com/tektoncd/operator/pull/3335) | — | — |
| 66f8f28 | fix(cve): Fix CVE-2026-34986 (go-jose DoS) | Jawed khelil | [#3333](https://github.com/tektoncd/operator/pull/3333) | [SRVKP-11500](https://redhat.atlassian.net/browse/SRVKP-11500) *(PR body)* | **CVE:** Bump go-jose/v4 from v4.0.5 to v4.1.4 to fix CVE-2026-34986 (DoS via crafted JWE *(extracted)* |
| 3c00e60 | fix: fix CEL expression and remove release-right-meow SA | Vincent Demeester | [#3336](https://github.com/tektoncd/operator/pull/3336) | — | — |
| f0a780e | chore: bump component versions | Vincent Demeester | [#3330](https://github.com/tektoncd/operator/pull/3330) | [SRVKP-12488](https://redhat.atlassian.net/browse/SRVKP-12488) *(keyword)* | **Release Notes Not Required:** ** [v0.44.0...v0.45.0](https://github.com/openshift-pipelines/pipelines-as-code/ *(extracted)* |
| 243ea3b | Add centralized TLS configuration infrastructure | Emil Natan | [#3225](https://github.com/tektoncd/operator/pull/3225) | [SRVKP-12006](https://redhat.atlassian.net/browse/SRVKP-12006) *(keyword)* | **Feature:** Add support for centralized TLS configuration on OpenShift. When `enableCentralT *(extracted)* |
| 18df1de | Enable centralized TLS configuration for TektonResult | Emil Natan | [#3225](https://github.com/tektoncd/operator/pull/3225) | [SRVKP-11697](https://redhat.atlassian.net/browse/SRVKP-11697) *(keyword)* | **Feature:** Add support for centralized TLS configuration on OpenShift. When `enableCentralT *(extracted)* |
| fff9269 | Add testing plan for central TLS functionality | Emil Natan | [#3225](https://github.com/tektoncd/operator/pull/3225) | [SRVKP-12006](https://redhat.atlassian.net/browse/SRVKP-12006) *(keyword)* | **Feature:** Add support for centralized TLS configuration on OpenShift. When `enableCentralT *(extracted)* |
| 9a475cd | Set DeploymentEnvVarKubernetesMinVersion as common transform | Pramod Bindal | [#3339](https://github.com/tektoncd/operator/pull/3339) | [SRVKP-7276](https://redhat.atlassian.net/browse/SRVKP-7276) *(keyword)* | **Release Notes Not Required** *(extracted)* |
| 56c1e67 | chore(deps): bump go.opentelemetry.io/otel/sdk from 1.40.0 t | dependabot[bot] | [#3338](https://github.com/tektoncd/operator/pull/3338) | — | skipped |
| b7c1ff8 | chore(deps): bump chainguard-dev/actions from 1.6.11 to 1.6. | dependabot[bot] | [#3345](https://github.com/tektoncd/operator/pull/3345) | — | skipped |
| 04de724 | chore(deps): bump github.com/sigstore/timestamp-authority/v2 | dependabot[bot] | [#3346](https://github.com/tektoncd/operator/pull/3346) | — | skipped |
| 6d8e0ed | chore(deps): bump step-security/harden-runner from 2.16.0 to | dependabot[bot] | [#3344](https://github.com/tektoncd/operator/pull/3344) | — | skipped |
| d8985d0 | fix: remove YAML document separator from .tekton files | Vincent Demeester | [#3342](https://github.com/tektoncd/operator/pull/3342) | — | — |
| c154af5 | chore(deps): bump github.com/cert-manager/cert-manager | dependabot[bot] | [#3343](https://github.com/tektoncd/operator/pull/3343) | — | skipped |
| 69f0541 | chore(deps): bump actions/cache from 5.0.4 to 5.0.5 | dependabot[bot] | [#3353](https://github.com/tektoncd/operator/pull/3353) | — | skipped |
| 649c104 | chore(deps): bump actions/upload-artifact from 7.0.0 to 7.0. | dependabot[bot] | [#3354](https://github.com/tektoncd/operator/pull/3354) | — | skipped |
| 1e483d3 | chore(deps): bump peter-evans/create-pull-request from 8.1.0 | dependabot[bot] | [#3355](https://github.com/tektoncd/operator/pull/3355) | — | skipped |
| d6595a7 | chore(deps): bump golang.org/x/mod from 0.34.0 to 0.35.0 | dependabot[bot] | [#3356](https://github.com/tektoncd/operator/pull/3356) | — | skipped |
| 322ddc3 | chore(deps): bump github.com/cert-manager/cert-manager | dependabot[bot] | [#3357](https://github.com/tektoncd/operator/pull/3357) | — | skipped |
| 1375778 | chore(deps): bump github.com/sigstore/cosign/v2 from 2.6.2 t | dependabot[bot] | [#3358](https://github.com/tektoncd/operator/pull/3358) | — | skipped |
| 1826275 | chore(deps): bump zizmorcore/zizmor-action from 0.5.2 to 0.5 | dependabot[bot] | [#3359](https://github.com/tektoncd/operator/pull/3359) | — | skipped |
| 5c00da3 | update subject for tekton-scheduler-role to tekton-operator | Pramod Bindal | [#3361](https://github.com/tektoncd/operator/pull/3361) | — | — |
| ef45264 | fix: use correct fully qualified image names in release pipe | Vincent Demeester | [#3362](https://github.com/tektoncd/operator/pull/3362) | — | — |
| 71ed26f | Add Pipelines-as-Code on Kubernetes and resolve openshift-pi | mbpavan | [#3337](https://github.com/tektoncd/operator/pull/3337) | [SRVKP-11429](https://redhat.atlassian.net/browse/SRVKP-11429) | **Feature:** The Tekton Operator can now install and reconcile Pipelines-as-Code (PAC) on van *(extracted)* |
| 2a8ae75 | Use dedicated release ServiceAccount in PAC release pipeline | Vincent Demeester | [#3364](https://github.com/tektoncd/operator/pull/3364) | — | — |
| 8eebadf | chore: bump pipeline from v1.11.0 to v1.11.1 | Vincent Demeester | [#3365](https://github.com/tektoncd/operator/pull/3365) | — | — |
| ac736e5 | fix: use source-subpath/release-subpath params for create-dr | Vincent Demeester | [#3366](https://github.com/tektoncd/operator/pull/3366) | — | — |
| 4f23998 | chore: bump chains from v0.26.2 to v0.26.3 | Vincent Demeester | [#3367](https://github.com/tektoncd/operator/pull/3367) | — | — |
| 1efe55b | fix(release): correct devel version rewriting in publish-ope | Anitha Natarajan | [#3370](https://github.com/tektoncd/operator/pull/3370) | [SRVKP-12011](https://redhat.atlassian.net/browse/SRVKP-12011) *(unmatched)* | — |
| a86e469 | chore(deps): bump step-security/harden-runner from 2.17.0 to | dependabot[bot] | [#3374](https://github.com/tektoncd/operator/pull/3374) | — | skipped |
| f028c33 | chore: bump hub from v1.23.9 to v1.23.10 | Vincent Demeester | [#3375](https://github.com/tektoncd/operator/pull/3375) | — | — |
| d0c685f | chore(deps): bump chainguard-dev/actions from 1.6.14 to 1.6. | dependabot[bot] | [#3373](https://github.com/tektoncd/operator/pull/3373) | — | skipped |
| af76c6f | chore(deps): bump github/codeql-action from 4.35.1 to 4.35.2 | dependabot[bot] | [#3372](https://github.com/tektoncd/operator/pull/3372) | — | skipped |
| 1f7ac47 | Bump dependencies for OpenTelemetry migration | Shubham Bhardwaj | [#3332](https://github.com/tektoncd/operator/pull/3332) | [SRVKP-9320](https://redhat.atlassian.net/browse/SRVKP-9320) | **Enhancement:** Migrated operator metrics from OpenCensus to OpenTelemetry.  ACTION REQUIRED:  1 *(extracted)* |
| 6df3de1 | Update generated code for new dependencies | Shubham Bhardwaj | [#3332](https://github.com/tektoncd/operator/pull/3332) | [SRVKP-9320](https://redhat.atlassian.net/browse/SRVKP-9320) | **Enhancement:** Migrated operator metrics from OpenCensus to OpenTelemetry.  ACTION REQUIRED:  1 *(extracted)* |
| 161a1c4 | Migrate metrics from OpenCensus to OpenTelemetry | Shubham Bhardwaj | [#3332](https://github.com/tektoncd/operator/pull/3332) | [SRVKP-9320](https://redhat.atlassian.net/browse/SRVKP-9320) | **Enhancement:** Migrated operator metrics from OpenCensus to OpenTelemetry.  ACTION REQUIRED:  1 *(extracted)* |
| e3b4e28 | Remove unused custom operator metrics | Shubham Bhardwaj | [#3332](https://github.com/tektoncd/operator/pull/3332) | [SRVKP-9320](https://redhat.atlassian.net/browse/SRVKP-9320) | **Enhancement:** Migrated operator metrics from OpenCensus to OpenTelemetry.  ACTION REQUIRED:  1 *(extracted)* |
| 33020d9 | chore(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#3378](https://github.com/tektoncd/operator/pull/3378) | — | skipped |
| e4eaabd | chore(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#3379](https://github.com/tektoncd/operator/pull/3379) | — | skipped |
| 70798fa | chore: bump pipeline from v1.11.1 to v1.12.0 | Vincent Demeester | [#3395](https://github.com/tektoncd/operator/pull/3395) | — | — |
| ae37659 | chore(deps): bump chainguard-dev/actions from 1.6.15 to 1.6. | dependabot[bot] | [#3389](https://github.com/tektoncd/operator/pull/3389) | — | skipped |
| a887458 | chore(deps): bump k8s.io/apiextensions-apiserver from 0.35.3 | dependabot[bot] | [#3391](https://github.com/tektoncd/operator/pull/3391) | — | skipped |
| 616106e | chore(deps): bump github.com/tektoncd/pipeline from 1.11.0 t | dependabot[bot] | [#3392](https://github.com/tektoncd/operator/pull/3392) | — | skipped |
| 4db29a4 | Feat: Console Plugin Image should be picked conditionally ba | Pramod Bindal | [#3386](https://github.com/tektoncd/operator/pull/3386) | [SRVKP-11927](https://redhat.atlassian.net/browse/SRVKP-11927) | **Enhancement:** Added support for conditionally picking the console-plugin image on openshift. O *(extracted)* |
| 0680028 | chore(deps): bump github.com/openshift-pipelines/pipelines-a | dependabot[bot] | [#3393](https://github.com/tektoncd/operator/pull/3393) | — | skipped |
| 6a1cb9e | pin base image | Anitha Natarajan | [#3396](https://github.com/tektoncd/operator/pull/3396) | — | — |
| f558e68 | fix: update the koExtraArgs | Anitha Natarajan | [#3400](https://github.com/tektoncd/operator/pull/3400) | — | — |
| cecf39b | fix: missing subjects in attestation and capture UUID instea | Anitha Natarajan | [#3402](https://github.com/tektoncd/operator/pull/3402) | — | — |
| 694775d | chore: bump component versions | Vincent Demeester | [#3404](https://github.com/tektoncd/operator/pull/3404) | [SRVKP-12488](https://redhat.atlassian.net/browse/SRVKP-12488) *(keyword)* | **Release Notes Not Required:** ** [v0.45.0...v0.46.0](https://github.com/tektoncd/pipelines-as-code/compare/v0. *(extracted)* |
| 7c96af4 | Generate OpenAPI v3 schemas for all CRDs using controller-ge | Alessandro | [#3340](https://github.com/tektoncd/operator/pull/3340) | [SRVKP-12477](https://redhat.atlassian.net/browse/SRVKP-12477) *(keyword)* | **Enhancement:** CRD manifests now include full OpenAPI v3 schemas generated for all operator.tek *(extracted)* |
| c79fc55 | fix: helm automation with correct version | Anitha Natarajan | [#3405](https://github.com/tektoncd/operator/pull/3405) | — | — |
| ae5d04c | chore(deps): bump github.com/in-toto/in-toto-golang from 0.9 | dependabot[bot] | [#3414](https://github.com/tektoncd/operator/pull/3414) | — | skipped |
| a9cb7e4 | fix(cve): update Go to 1.25.10 to address stdlib security vu | khelil | [#3409](https://github.com/tektoncd/operator/pull/3409) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Security fix: update Go from 1.25.8 to 1.25.10 to address 10 standard library vu *(extracted)* |
| ecc167c | chore(deps): bump chainguard-dev/actions from 1.6.16 to 1.6. | dependabot[bot] | [#3417](https://github.com/tektoncd/operator/pull/3417) | — | skipped |
| 5c2c596 | chore(deps): bump step-security/harden-runner from 2.19.0 to | dependabot[bot] | [#3421](https://github.com/tektoncd/operator/pull/3421) | — | skipped |
| 0acc7e9 | chore(deps): bump github/codeql-action from 4.35.2 to 4.35.3 | dependabot[bot] | [#3420](https://github.com/tektoncd/operator/pull/3420) | — | skipped |
| 499b5bf | chore(deps): bump go.uber.org/zap from 1.27.1 to 1.28.0 | dependabot[bot] | [#3419](https://github.com/tektoncd/operator/pull/3419) | — | skipped |
| f2f55b4 |  fix(tekton-results): use passthrough TLS termination for ro | divyansh42 | [#3425](https://github.com/tektoncd/operator/pull/3425) | [SRVKP-11697](https://redhat.atlassian.net/browse/SRVKP-11697) | **Bug Fix:** Tekton Results API route now uses passthrough TLS termination by default, enabli *(extracted)* |
| 5dfea2e | fix: update-tektoncd-task-versions workflow to resolve CI fa | Shubham Bhardwaj | [#3426](https://github.com/tektoncd/operator/pull/3426) | [SRVKP-12022](https://redhat.atlassian.net/browse/SRVKP-12022) | **Release Notes Not Required** *(extracted)* |
| 345f8a3 | fix: harden patch-release workflow against script injection | Vincent Demeester | [#3424](https://github.com/tektoncd/operator/pull/3424) | — | — |
| ba5b11d | update task-maven to 0.4.1 | ab-ghosh | [#3415](https://github.com/tektoncd/operator/pull/3415) | [SRVKP-11954](https://redhat.atlassian.net/browse/SRVKP-11954) | **Release Notes Not Required** *(extracted)* |
| 4fec297 | feat(tls): make central TLS opt-out and enable ML-KEM for co | Jawed khelil | [#3416](https://github.com/tektoncd/operator/pull/3416) | [SRVKP-11957](https://redhat.atlassian.net/browse/SRVKP-11957) | **Feature:** Central TLS configuration is now enabled by default on OpenShift. The `enableCen *(extracted)* |
| 82a21f8 | chore(deps): bump k8s.io/apimachinery from 0.35.4 to 0.35.5 | dependabot[bot] | [#3431](https://github.com/tektoncd/operator/pull/3431) | — | skipped |
| 8237c05 | chore(deps): bump golang.org/x/mod from 0.35.0 to 0.36.0 | dependabot[bot] | [#3433](https://github.com/tektoncd/operator/pull/3433) | — | skipped |
| 5900b24 | chore(deps): bump github/codeql-action from 4.35.3 to 4.35.4 | dependabot[bot] | [#3434](https://github.com/tektoncd/operator/pull/3434) | — | skipped |
| 50bc752 | chore(deps): bump github.com/tektoncd/pipeline from 1.11.1 t | dependabot[bot] | [#3418](https://github.com/tektoncd/operator/pull/3418) | — | skipped |
| bcfa313 | fix: update SendCloudEventsForRuns default and mark deprecat | Anitha Natarajan | [#3418](https://github.com/tektoncd/operator/pull/3418) | — | — |
| e1484f9 | refactor(tls): extract SetupAPIServerTLSWatch into occommon  | Jawed khelil | [#3406](https://github.com/tektoncd/operator/pull/3406) | [SRVKP-9613](https://redhat.atlassian.net/browse/SRVKP-9613) | **Feature:** On OpenShift, the tekton-operator-webhook and tekton-operator-proxy-webhook now  *(extracted)* |
| 534fc37 | feat(webhook/tls): read OpenShift APIServer TLS profile at s | Jawed khelil | [#3406](https://github.com/tektoncd/operator/pull/3406) | [SRVKP-9613](https://redhat.atlassian.net/browse/SRVKP-9613) | **Feature:** On OpenShift, the tekton-operator-webhook and tekton-operator-proxy-webhook now  *(extracted)* |
| 4975b71 | feat(proxy-webhook/tls): read OpenShift APIServer TLS profil | Jawed khelil | [#3406](https://github.com/tektoncd/operator/pull/3406) | [SRVKP-9613](https://redhat.atlassian.net/browse/SRVKP-9613) | **Feature:** On OpenShift, the tekton-operator-webhook and tekton-operator-proxy-webhook now  *(extracted)* |
| 6888701 | feat(tls): inject centrally managed TLS config into tekton-p | Jawed khelil | [#3383](https://github.com/tektoncd/operator/pull/3383) | [SRVKP-9614](https://redhat.atlassian.net/browse/SRVKP-9614) | **Feature:** The tekton-pipelines-webhook now inherits TLS configuration (minimum version and *(extracted)* |
| a9c8470 | feat(tls): inject centrally managed TLS config into tekton-t | Jawed khelil | [#3384](https://github.com/tektoncd/operator/pull/3384) | [SRVKP-9615](https://redhat.atlassian.net/browse/SRVKP-9615) | **Feature:** The tekton-triggers-webhook and tekton-triggers-core-interceptor  now inherits T *(extracted)* |
| d43fca7 | feat(tls): inject TLS env vars into triggers core intercepto | Jawed khelil | [#3384](https://github.com/tektoncd/operator/pull/3384) | [SRVKP-9615](https://redhat.atlassian.net/browse/SRVKP-9615) | **Feature:** The tekton-triggers-webhook and tekton-triggers-core-interceptor  now inherits T *(extracted)* |
| 273cd27 | docs(install): document that default namespace is unsupporte | Matt Van Horn | [#3440](https://github.com/tektoncd/operator/pull/3440) | [SRVKP-12017](https://redhat.atlassian.net/browse/SRVKP-12017) *(PR body)* | **Release Notes Not Required** *(extracted)* |
| a2af438 | docs(install): fix broken OpenShift namespaces link | Matt Van Horn | [#3440](https://github.com/tektoncd/operator/pull/3440) | [SRVKP-12017](https://redhat.atlassian.net/browse/SRVKP-12017) *(PR body)* | **Release Notes Not Required** *(extracted)* |
| a9fc312 | chore(deps): bump step-security/harden-runner from 2.19.1 to | dependabot[bot] | [#3443](https://github.com/tektoncd/operator/pull/3443) | — | skipped |
| 77711c7 | chore(deps): bump zizmorcore/zizmor-action from 0.5.3 to 0.5 | dependabot[bot] | [#3444](https://github.com/tektoncd/operator/pull/3444) | — | skipped |
| 65611f0 | chore(deps): bump github.com/konflux-ci/tekton-kueue from 0. | dependabot[bot] | [#3445](https://github.com/tektoncd/operator/pull/3445) | — | skipped |
| 8ce8743 | chore(deps): bump github.com/tektoncd/pruner from 0.3.5 to 0 | dependabot[bot] | [#3446](https://github.com/tektoncd/operator/pull/3446) | — | skipped |
| af2546d | chore(deps): bump k8s.io/apiextensions-apiserver from 0.35.4 | dependabot[bot] | [#3447](https://github.com/tektoncd/operator/pull/3447) | — | skipped |
| e363e1a | fix(cve): GO-2026-5019/GO-2026-5018 - update golang.org/x/cr | Jawed Khelil | [#3451](https://github.com/tektoncd/operator/pull/3451) | [SRVKP-6935](https://redhat.atlassian.net/browse/SRVKP-6935) *(keyword)* | **CVE:** Updates golang.org/x/crypto from v0.50.0 to v0.52.0 to address GO-2026-5019 and  *(generated)* |
| 2ef7cc9 | chore(deps): bump github/codeql-action from 4.35.4 to 4.35.5 | dependabot[bot] | [#3449](https://github.com/tektoncd/operator/pull/3449) | — | skipped |
| d42a211 | chore(deps): bump github.com/openshift-pipelines/pipelines-a | dependabot[bot] | [#3432](https://github.com/tektoncd/operator/pull/3432) | — | skipped |
| a4e8dba | chore: fix failing unit tests during upgrade | Anitha Natarajan | [#3432](https://github.com/tektoncd/operator/pull/3432) | — | — |
| 4573cf5 | Add centrally managed TLS configuration for console-plugin n | Jawed khelil | [#3218](https://github.com/tektoncd/operator/pull/3218) | [SRVKP-9632](https://redhat.atlassian.net/browse/SRVKP-9632) | **Feature:** The console-plugin nginx server now inherits TLS settings from the centrally man *(extracted)* |
| 82150f0 | chore(deps): update postgres-15 image for Openshift | Andrew Thorp | [#3455](https://github.com/tektoncd/operator/pull/3455) | [SRVKP-8836](https://redhat.atlassian.net/browse/SRVKP-8836) *(keyword)* | **Release Notes Not Required:** Update images digests for Openshift images *(extracted)* |
| 1c1be97 | chore(deps): Upgrade image SHAs for Openshift images | Andrew Thorp | [#3455](https://github.com/tektoncd/operator/pull/3455) | [SRVKP-3348](https://redhat.atlassian.net/browse/SRVKP-3348) *(keyword)* | **Release Notes Not Required:** Update images digests for Openshift images *(extracted)* |
| a8699d8 | feat(tls): inject centrally managed TLS config into pipeline | Jawed khelil | [#3385](https://github.com/tektoncd/operator/pull/3385) | [SRVKP-9616](https://redhat.atlassian.net/browse/SRVKP-9616) | **Feature:** On OpenShift, the `pipelines-as-code-webhook` deployment now automatically inher *(extracted)* |
| c5b1051 | fix(cve): GO-2026-5026 - update golang.org/x/net v0.53.0 → v | Jawed Khelil | [#3452](https://github.com/tektoncd/operator/pull/3452) | [SRVKP-6924](https://redhat.atlassian.net/browse/SRVKP-6924) *(keyword)* | **CVE:** Updates golang.org/x/net from v0.53.0 to v0.55.0 to address GO-2026-5026 securit *(generated)* |
| 77a2f42 | feat(tls): inject centrally managed TLS config into pruner w | Shubham Bhardwaj | [#3453](https://github.com/tektoncd/operator/pull/3453) | [SRVKP-9623](https://redhat.atlassian.net/browse/SRVKP-9623) *(keyword)* | **Feature:** Central TLS configuration is now injected into the tekton-pruner-webhook on Open *(extracted)* |
| cdd2fcf | chore(deps): bump step-security/harden-runner from 2.19.3 to | dependabot[bot] | [#3459](https://github.com/tektoncd/operator/pull/3459) | — | skipped |
| 34125bb | chore(deps): bump github/codeql-action from 4.35.5 to 4.36.0 | dependabot[bot] | [#3462](https://github.com/tektoncd/operator/pull/3462) | — | skipped |
| de1cae0 | chore(deps): bump github.com/tektoncd/triggers from 0.35.0 t | dependabot[bot] | [#3461](https://github.com/tektoncd/operator/pull/3461) | — | skipped |
| b4b1a7f | chore(deps): bump golangci/golangci-lint-action from 9.2.0 t | dependabot[bot] | [#3463](https://github.com/tektoncd/operator/pull/3463) | — | skipped |
| cdcccd7 | chore(deps): bump github.com/openshift-pipelines/pipelines-a | dependabot[bot] | [#3460](https://github.com/tektoncd/operator/pull/3460) | — | skipped |
| e91f3b3 | fix(cve): Update Go 1.25.10 → 1.25.11 to fix stdlib CVEs | Jawed Khelil | [#3468](https://github.com/tektoncd/operator/pull/3468) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Updates Go from 1.25.10 to 1.25.11 to fix stdlib CVEs. *(generated)* |
| 4431a67 | chore(Makefile): align golangci-lint with CI and drop unused | Jawed khelil | [#3387](https://github.com/tektoncd/operator/pull/3387) | [SRVKP-11820](https://redhat.atlassian.net/browse/SRVKP-11820) | **Release Notes Not Required** *(generated)* |
| d453bcf | docs: add AGENTS.md and agent skills | Jawed khelil | [#3387](https://github.com/tektoncd/operator/pull/3387) | [SRVKP-11820](https://redhat.atlassian.net/browse/SRVKP-11820) | **Release Notes Not Required** *(generated)* |
| c44c0c8 | chore: bump component versions | Jawed khelil | [#3464](https://github.com/tektoncd/operator/pull/3464) | [SRVKP-12488](https://redhat.atlassian.net/browse/SRVKP-12488) *(keyword)* | **Release Notes Not Required:** ** [v0.3.5...v0.4.0](https://github.com/tektoncd/pruner/compare/v0.3.5...v0.4.0) *(extracted)* |
| f3dc5d2 | fix(hub): ensure PVCs are applied on every reconcile cycle | Jawed khelil | [#3464](https://github.com/tektoncd/operator/pull/3464) | — | — |
| 980306f | test(e2e): update addon installer set names for resolver mig | Jawed khelil | [#3464](https://github.com/tektoncd/operator/pull/3464) | — | — |
| 62864a7 | fix(openshift): delete operator webhooks before namespace cl | Jawed khelil | [#3472](https://github.com/tektoncd/operator/pull/3472) | [SRVKP-12010](https://redhat.atlassian.net/browse/SRVKP-12010) | **Bug Fix:** Bug fix: OpenShift Pipelines operator now uninstalls cleanly via the Console UI. *(extracted)* |
| ca2b95e | bump chains to 0.27.1 | Jawed khelil | [#3475](https://github.com/tektoncd/operator/pull/3475) | — | — |
| a30fb53 | fix(helm): fix appVersion prefix and k8s min version handlin | Jawed khelil | [#3487](https://github.com/tektoncd/operator/pull/3487) | — | — |
| b344392 | [v0.80.x] bump results to v0.19.0 | Khurram Baig | [#3498](https://github.com/tektoncd/operator/pull/3498) | — | — |
| 335c26e | feat(helm): publish chart to OCI registry on release | Jawed khelil | [#3511](https://github.com/tektoncd/operator/pull/3511) | — | — |
| b35002f | Merge multikueueclusters permission to scheduler role | Pramod Bindal | [#3521](https://github.com/tektoncd/operator/pull/3521) | — | — |
| 6ae9dae | fix: options field should be optional in all components | Pramod Bindal | [#3528](https://github.com/tektoncd/operator/pull/3528) | — | — |
| ff6ad6c | chore(deps): bump chainguard-dev/actions from 1.6.19 to 1.6. | dependabot[bot] | [#3534](https://github.com/tektoncd/operator/pull/3534) | — | skipped |
| ec650cf | chore(deps): bump actions/checkout from 6.0.2 to 6.0.3 | dependabot[bot] | [#3536](https://github.com/tektoncd/operator/pull/3536) | — | skipped |
| 20b0c41 | chore(deps): bump k8s.io/apiextensions-apiserver from 0.35.5 | dependabot[bot] | [#3550](https://github.com/tektoncd/operator/pull/3550) | — | skipped |
| fa89ed8 | chore(deps): bump github/codeql-action from 4.36.0 to 4.36.2 | dependabot[bot] | [#3540](https://github.com/tektoncd/operator/pull/3540) | — | skipped |
| e31b202 | chore(deps): bump github.com/tektoncd/pipeline from 1.12.0 t | dependabot[bot] | [#3552](https://github.com/tektoncd/operator/pull/3552) | — | skipped |
| ffdc96b | fix(helm): add openshift-pipelines.org RBAC to kubernetes ch | Jawed khelil | [#3579](https://github.com/tektoncd/operator/pull/3579) | [SRVKP-2270](https://redhat.atlassian.net/browse/SRVKP-2270) *(keyword)* | **Bug Fix:** Fix ManualApprovalGate installation failure on Kubernetes when using the Helm ch *(extracted)* |
| f0f903c | Cleanup manual CRDs and use genearted CRDs in kustomize | Pramod Bindal | [#3580](https://github.com/tektoncd/operator/pull/3580) | [SRVKP-8082](https://redhat.atlassian.net/browse/SRVKP-8082) *(keyword)* | **Release Notes Not Required** *(extracted)* |
| 5c9be32 | Cleanup manual CRDs and use genearted CRDs in kustomize | Pramod Bindal | [#3580](https://github.com/tektoncd/operator/pull/3580) | — | — |

---

## manual-approval-gate

**Upstream:** openshift-pipelines/manual-approval-gate · **Commits:** 2 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 5932c2f | Update go and knative version | divyansh42 | [#965](https://github.com/openshift-pipelines/manual-approval-gate/pull/965) | [SRVKP-11514](https://redhat.atlassian.net/browse/SRVKP-11514) | **Enhancement:** ```  Updated dependencies including Kubernetes (v0.35.3), Tekton Pipelines (v1.  |
| edb306f | Update go version in workflows and remove konflux related fi | divyansh42 | [#967](https://github.com/openshift-pipelines/manual-approval-gate/pull/967) | — | **Release Notes Not Required**  |

---

## tektoncd-chains

**Upstream:** tektoncd/chains · **Commits:** 145 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 21dfd1e | add release v0.26.0 to release documentation (#1468) | khelil | [#1468](https://github.com/tektoncd/chains/pull/1468) | — | — |
| 37297be | fix:improve e2e test failure analysis with chains controller | Anitha Natarajan | [#1470](https://github.com/tektoncd/chains/pull/1470) | — | — |
| ec6157c | add anithapriyanatarajan as maintainer (#1473) | khelil | [#1473](https://github.com/tektoncd/chains/pull/1473) | — | — |
| 3754523 | Add GitHub Actions workflow for build and unit tests (#1446) | Alan Greene | [#1446](https://github.com/tektoncd/chains/pull/1446) | — | — |
| 59c9aa9 | docs: fix cosign repo reference (#1434) | Emmanuel Ferdman | [#1434](https://github.com/tektoncd/chains/pull/1434) | — | — |
| 61ead14 | .github/workflows: use CHATOPS_TOKEN for coverage comments ( | Vincent Demeester | [#1480](https://github.com/tektoncd/chains/pull/1480) | — | — |
| 46ec5dc | Bump chainguard-dev/actions | dependabot[bot] | [#1462](https://github.com/tektoncd/chains/pull/1462) | — | skipped |
| 52e0f7c | feat: add retest workflow using plumbing reusable workflow ( | Vincent Demeester | [#1488](https://github.com/tektoncd/chains/pull/1488) | — | — |
| 93b3806 | ci: consolidate and modernize GitHub Actions workflows (#149 | Vincent Demeester | [#1490](https://github.com/tektoncd/chains/pull/1490) | — | — |
| 965df1d | feat: add cherry-pick slash command workflow (#1487) | Vincent Demeester | [#1487](https://github.com/tektoncd/chains/pull/1487) | — | — |
| 6f41388 | Bump github.com/in-toto/go-witness from 0.8.6 to 0.9.1 (#146 | dependabot[bot] | [#1460](https://github.com/tektoncd/chains/pull/1460) | — | skipped |
| 295ad99 | Refactor the annotations management (#1422) | Emil Natan | [#1422](https://github.com/tektoncd/chains/pull/1422) | — | — |
| 503e560 | chore: replace gcs storage path with infra.tekton.dev (#1491 | Anitha Natarajan | [#1491](https://github.com/tektoncd/chains/pull/1491) | — | — |
| 11b7a58 | Bump actions/checkout from 4.2.2 to 6.0.1 (#1483) | dependabot[bot] | [#1483](https://github.com/tektoncd/chains/pull/1483) | — | skipped |
| a82ca07 | Bump github/codeql-action from 3.30.7 to 4.31.7 (#1486) | dependabot[bot] | [#1486](https://github.com/tektoncd/chains/pull/1486) | — | skipped |
| f1f3796 | chore(release-pipeline): improve release cheat sheet structu | Shubham Bhardwaj | [#1489](https://github.com/tektoncd/chains/pull/1489) | — | — |
| 128c252 | ci: fix and modernize e2e workflow step ordering (#1492) | Vincent Demeester | [#1492](https://github.com/tektoncd/chains/pull/1492) | — | — |
| 4bf9c5a | Bump golang.org/x/crypto from 0.42.0 to 0.45.0 (#1477) | dependabot[bot] | [#1477](https://github.com/tektoncd/chains/pull/1477) | — | skipped |
| e5d569a | Bump actions/setup-go from 5.0.2 to 6.1.0 (#1484) | dependabot[bot] | [#1484](https://github.com/tektoncd/chains/pull/1484) | — | skipped |
| 222884e | chore(deps): bump sigstore/scaffolding from 0.7.25 to 0.7.33 | dependabot[bot] | [#1494](https://github.com/tektoncd/chains/pull/1494) | — | skipped |
| abfe82f | chore(deps): bump step-security/harden-runner from 2.13.1 to | dependabot[bot] | [#1495](https://github.com/tektoncd/chains/pull/1495) | — | skipped |
| 3767a6a | chore(deps): bump actions/upload-artifact from 4.6.2 to 6.0. | dependabot[bot] | [#1496](https://github.com/tektoncd/chains/pull/1496) | — | skipped |
| 46909c8 | chore(deps): bump chainguard-dev/actions from 1.5.7 to 1.5.1 | dependabot[bot] | [#1497](https://github.com/tektoncd/chains/pull/1497) | — | skipped |
| 2032c9a | chore: fix some minor issues in comments | socialsister | [#1403](https://github.com/tektoncd/chains/pull/1403) | — | — |
| 8763405 | fix: microshift e2e test failures on merge (#1500) | Anitha Natarajan | [#1500](https://github.com/tektoncd/chains/pull/1500) | — | — |
| 920e627 | Remove GHCR migration notice from the readme | Alan Greene | [#1499](https://github.com/tektoncd/chains/pull/1499) | — | — |
| 5e01cd9 | docs: separate command and output in tutorial | ab-ghosh | [#1498](https://github.com/tektoncd/chains/pull/1498) | — | — |
| 4b42e7f | fix: replace hard codede go ref from mod file (#1501) | Anitha Natarajan | [#1501](https://github.com/tektoncd/chains/pull/1501) | — | — |
| 29046c2 | .github/workflows: Add comment for dependabot to pick up upd | Shubham Bhardwaj | [#1502](https://github.com/tektoncd/chains/pull/1502) | — | — |
| 7688248 | chore(deps): bump github/codeql-action from 4.31.8 to 4.31.9 | dependabot[bot] | [#1504](https://github.com/tektoncd/chains/pull/1504) | — | skipped |
| b9d06f1 | feat(oci): support insecure OCI registry (#1374) | l-qing | [#1374](https://github.com/tektoncd/chains/pull/1374) | — | — |
| 9b6d193 | chore(deps): bump the all group across 1 directory with 19 u | dependabot[bot] | [#1503](https://github.com/tektoncd/chains/pull/1503) | — | skipped |
| f07f988 | update golangci-lint version that supports go1.25 (#1508) | Anitha Natarajan | [#1508](https://github.com/tektoncd/chains/pull/1508) | — | — |
| 8bd204d | Fix golangci-lint action step for large diff (#1510) | Anitha Natarajan | [#1510](https://github.com/tektoncd/chains/pull/1510) | — | — |
| 0e98b57 | bump fulcio and other dependencies (#1506) | khelil | [#1506](https://github.com/tektoncd/chains/pull/1506) | — | — |
| 09000a3 | test: add unit tests for pubsub backend constructor and kafk | Naomi Gelman | [#1511](https://github.com/tektoncd/chains/pull/1511) | — | — |
| aec00d5 | chore(deps): bump the all group across 1 directory with 5 up | dependabot[bot] | [#1512](https://github.com/tektoncd/chains/pull/1512) | — | skipped |
| 5abdfb6 | chore(deps): bump github/codeql-action from 4.31.9 to 4.31.1 | dependabot[bot] | [#1513](https://github.com/tektoncd/chains/pull/1513) | — | skipped |
| a9205f6 | chore(deps): bump the all group across 1 directory with 2 up | dependabot[bot] | [#1518](https://github.com/tektoncd/chains/pull/1518) | — | skipped |
| 2feaeb6 | chore(deps): bump chainguard-dev/actions from 1.5.10 to 1.5. | dependabot[bot] | [#1514](https://github.com/tektoncd/chains/pull/1514) | — | skipped |
| f1a581f | chore(deps): bump github.com/sigstore/fulcio from 1.8.4 to 1 | dependabot[bot] | [#1516](https://github.com/tektoncd/chains/pull/1516) | — | skipped |
| 8cab516 | chore(deps): bump chainguard-dev/actions from 1.5.11 to 1.5. | dependabot[bot] | [#1520](https://github.com/tektoncd/chains/pull/1520) | — | skipped |
| 056611e | chore(deps): bump actions/setup-go from 6.1.0 to 6.2.0 (#151 | dependabot[bot] | [#1519](https://github.com/tektoncd/chains/pull/1519) | — | skipped |
| 99e2f1c | fix: update slash workflow to refer latest version that incl | Anitha Natarajan | [#1522](https://github.com/tektoncd/chains/pull/1522) | — | — |
| 9309310 | chore(deps): bump github.com/theupdateframework/go-tuf/v2 (# | dependabot[bot] | [#1525](https://github.com/tektoncd/chains/pull/1525) | — | skipped |
| e745f76 | chore(deps): bump github.com/sigstore/rekor from 1.4.3 to 1. | dependabot[bot] | [#1526](https://github.com/tektoncd/chains/pull/1526) | — | skipped |
| f1f32a4 | chore(deps): bump the all group with 6 updates (#1530) | dependabot[bot] | [#1530](https://github.com/tektoncd/chains/pull/1530) | — | skipped |
| 2f1eeca | chore(deps): bump github/codeql-action from 4.31.10 to 4.31. | dependabot[bot] | [#1531](https://github.com/tektoncd/chains/pull/1531) | — | skipped |
| 103697a | chore(deps): bump chainguard-dev/actions from 1.5.12 to 1.5. | dependabot[bot] | [#1532](https://github.com/tektoncd/chains/pull/1532) | — | skipped |
| 8282f99 | chore(deps): bump step-security/harden-runner from 2.14.0 to | dependabot[bot] | [#1533](https://github.com/tektoncd/chains/pull/1533) | — | skipped |
| af8a19d | chore(deps): bump actions/checkout from 6.0.1 to 6.0.2 (#153 | dependabot[bot] | [#1534](https://github.com/tektoncd/chains/pull/1534) | — | skipped |
| 6817618 | chore(deps): bump github.com/theupdateframework/go-tuf/v2 (# | dependabot[bot] | [#1535](https://github.com/tektoncd/chains/pull/1535) | — | skipped |
| 3c6b8dd | chore(ci): update cherry-pick workflow to fix multi-commit P | Vincent Demeester | [#1539](https://github.com/tektoncd/chains/pull/1539) | — | — |
| dc310e4 | chore(deps): bump cloud.google.com/go/storage in the all gro | dependabot[bot] | [#1541](https://github.com/tektoncd/chains/pull/1541) | — | skipped |
| 0da47bf | chore(deps): bump github/codeql-action from 4.31.11 to 4.32. | dependabot[bot] | [#1545](https://github.com/tektoncd/chains/pull/1545) | — | skipped |
| 2944597 | chore(deps): bump github.com/tektoncd/pipeline in the all gr | dependabot[bot] | [#1543](https://github.com/tektoncd/chains/pull/1543) | — | skipped |
| dc030cd | chore(deps): bump chainguard-dev/actions from 1.5.13 to 1.5. | dependabot[bot] | [#1544](https://github.com/tektoncd/chains/pull/1544) | — | skipped |
| 60afd54 | add Vincent as approver (#1546) | Anitha Natarajan | [#1546](https://github.com/tektoncd/chains/pull/1546) | — | — |
| 202f8e6 | Fix- Update Docdb storage logic (issue #1178) (#1505) | Naomi Gelman | [#1505](https://github.com/tektoncd/chains/pull/1505) | [SRVKP-10473](https://redhat.atlassian.net/browse/SRVKP-10473) | **Bug Fix:** The DocDB backend watcher now correctly detects filesystem changes under MongoSe *(extracted)* |
| 43a8448 | chore(docs update): update release cheatsheet, pipeline and  | Anitha Natarajan | [#1549](https://github.com/tektoncd/chains/pull/1549) | [SRVKP-10747](https://redhat.atlassian.net/browse/SRVKP-10747) | **Release Notes Not Required:** Internal release documentation and pipeline configuration updates for the Chains *(generated)* |
| 2889bdb | chore(deps): bump step-security/harden-runner from 2.14.1 to | dependabot[bot] | [#1551](https://github.com/tektoncd/chains/pull/1551) | — | skipped |
| 7293b71 | chore(deps): bump github/codeql-action from 4.32.1 to 4.32.2 | dependabot[bot] | [#1552](https://github.com/tektoncd/chains/pull/1552) | — | skipped |
| 4cdd90b | chore(deps): bump the all group across 1 directory with 7 up | dependabot[bot] | [#1556](https://github.com/tektoncd/chains/pull/1556) | — | skipped |
| 42d81ad | chore(deps): bump github/codeql-action from 4.32.2 to 4.32.3 | dependabot[bot] | [#1559](https://github.com/tektoncd/chains/pull/1559) | — | skipped |
| a4ed4a2 | chore(deps): bump chainguard-dev/actions from 1.5.14 to 1.6. | dependabot[bot] | [#1558](https://github.com/tektoncd/chains/pull/1558) | — | skipped |
| 43281fc | fix: doc typo (#1561) | Naomi Gelman | [#1561](https://github.com/tektoncd/chains/pull/1561) | — | — |
| 9d61b89 | chore(deps): bump the all group with 2 updates (#1563) | dependabot[bot] | [#1563](https://github.com/tektoncd/chains/pull/1563) | — | skipped |
| 9086de7 | ci: Add CI summary fan-in check (#1566) | Vincent Demeester | [#1566](https://github.com/tektoncd/chains/pull/1566) | — | — |
| a58d477 | chore(deps): bump filippo.io/edwards25519 from 1.1.0 to 1.1. | dependabot[bot] | [#1562](https://github.com/tektoncd/chains/pull/1562) | — | skipped |
| 5956a3d | chore(deps): bump chainguard-dev/actions from 1.6.1 to 1.6.3 | dependabot[bot] | [#1564](https://github.com/tektoncd/chains/pull/1564) | — | skipped |
| a890075 | chore(deps): bump github/codeql-action from 4.32.3 to 4.32.4 | dependabot[bot] | [#1565](https://github.com/tektoncd/chains/pull/1565) | — | skipped |
| 9b124c4 | chore(deps): bump github/codeql-action from 4.32.4 to 4.32.5 | dependabot[bot] | [#1574](https://github.com/tektoncd/chains/pull/1574) | — | skipped |
| da96d11 | chore(deps): bump the all group across 1 directory with 6 up | dependabot[bot] | [#1577](https://github.com/tektoncd/chains/pull/1577) | — | skipped |
| 97be44f | chore(deps): bump chainguard-dev/actions from 1.6.3 to 1.6.5 | dependabot[bot] | [#1571](https://github.com/tektoncd/chains/pull/1571) | — | skipped |
| df1288b | Move inactive approvers to emeritus (#1576) | Vincent Demeester | [#1576](https://github.com/tektoncd/chains/pull/1576) | — | — |
| ec0c951 | chore(deps): bump step-security/harden-runner from 2.14.2 to | dependabot[bot] | [#1570](https://github.com/tektoncd/chains/pull/1570) | — | skipped |
| 748ab47 | chore(deps): bump actions/upload-artifact from 6.0.0 to 7.0. | dependabot[bot] | [#1573](https://github.com/tektoncd/chains/pull/1573) | — | skipped |
| fb1f5ce | chore(deps): bump step-security/harden-runner from 2.15.0 to | dependabot[bot] | [#1579](https://github.com/tektoncd/chains/pull/1579) | — | skipped |
| 61a2d36 | chore(deps): bump chainguard-dev/actions from 1.6.5 to 1.6.7 | dependabot[bot] | [#1580](https://github.com/tektoncd/chains/pull/1580) | — | skipped |
| 837af9a | chore(deps): bump github/codeql-action from 4.32.5 to 4.32.6 | dependabot[bot] | [#1581](https://github.com/tektoncd/chains/pull/1581) | — | skipped |
| 031bf54 | chore(deps): bump the all group across 1 directory with 2 up | dependabot[bot] | [#1584](https://github.com/tektoncd/chains/pull/1584) | — | skipped |
| 42a1539 | chore(deps): bump cloud.google.com/go/storage in the all gro | dependabot[bot] | [#1585](https://github.com/tektoncd/chains/pull/1585) | — | skipped |
| ebe22e1 | chore(deps): bump chainguard-dev/actions from 1.6.7 to 1.6.8 | dependabot[bot] | [#1588](https://github.com/tektoncd/chains/pull/1588) | — | skipped |
| 5507533 | chore(deps): bump github/codeql-action from 4.32.6 to 4.33.0 | dependabot[bot] | [#1589](https://github.com/tektoncd/chains/pull/1589) | — | skipped |
| 280f991 | chore(deps): bump step-security/harden-runner from 2.15.1 to | dependabot[bot] | [#1586](https://github.com/tektoncd/chains/pull/1586) | — | skipped |
| f699aa4 | chore(deps): bump fgrosse/go-coverage-report from 1.2.0 to 1 | dependabot[bot] | [#1587](https://github.com/tektoncd/chains/pull/1587) | — | skipped |
| 25b9538 | chore(deps): bump google.golang.org/grpc from 1.79.2 to 1.79 | dependabot[bot] | [#1591](https://github.com/tektoncd/chains/pull/1591) | — | skipped |
| 84c79b1 | chore(deps): bump the all group across 1 directory with 6 up | dependabot[bot] | [#1593](https://github.com/tektoncd/chains/pull/1593) | — | skipped |
| 4010cc2 | chore(deps): bump chainguard-dev/actions from 1.6.8 to 1.6.9 | dependabot[bot] | [#1595](https://github.com/tektoncd/chains/pull/1595) | — | skipped |
| 8e8b95a | chore(deps): bump github/codeql-action from 4.33.0 to 4.34.1 | dependabot[bot] | [#1594](https://github.com/tektoncd/chains/pull/1594) | — | skipped |
| eff5aad | docs: Add MongoDB storage tutorial (#1557) | Naomi Gelman | [#1557](https://github.com/tektoncd/chains/pull/1557) | — | — |
| 53834cb | feat(metrics): Migrate from OpenCensus to OpenTelemetry (#15 | Anitha Natarajan | [#1550](https://github.com/tektoncd/chains/pull/1550) | [SRVKP-9322](https://redhat.atlassian.net/browse/SRVKP-9322) | **Feature:** Chains metrics implementation migrates from the deprecated OpenCensus library to *(extracted)* |
| 31743e6 | Bump knative.dev/pkg to adopt centralized WEBHOOK_* TLS conf | khelil | [#1602](https://github.com/tektoncd/chains/pull/1602) | [SRVKP-9322](https://redhat.atlassian.net/browse/SRVKP-9322) *(keyword)* | **Enhancement:** Chains adopts centralized WEBHOOK_* TLS configuration from knative.dev/pkg, enab *(generated)* |
| 52c98f7 | chore(deps): bump the all group across 1 directory with 7 up | dependabot[bot] | [#1608](https://github.com/tektoncd/chains/pull/1608) | — | skipped |
| 49d3db7 | chore(deps): bump github/codeql-action from 4.34.1 to 4.35.1 | dependabot[bot] | [#1607](https://github.com/tektoncd/chains/pull/1607) | — | skipped |
| 8676a34 | chore(deps): bump chainguard-dev/actions from 1.6.9 to 1.6.1 | dependabot[bot] | [#1613](https://github.com/tektoncd/chains/pull/1613) | — | skipped |
| 58bca7a | chore(deps): bump step-security/harden-runner from 2.16.0 to | dependabot[bot] | [#1614](https://github.com/tektoncd/chains/pull/1614) | — | skipped |
| 9609a3a | chore(deps): bump go.opentelemetry.io/otel/sdk from 1.42.0 t | dependabot[bot] | [#1619](https://github.com/tektoncd/chains/pull/1619) | — | skipped |
| 7d40d1b | chore(deps): bump the all group across 1 directory with 3 up | dependabot[bot] | [#1620](https://github.com/tektoncd/chains/pull/1620) | — | skipped |
| b3f28ec | Ci/zizmor security fixes (#1612) | Abhishek Ghosh | [#1612](https://github.com/tektoncd/chains/pull/1612) | — | — |
| daa3a1c | chore(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#1617](https://github.com/tektoncd/chains/pull/1617) | — | skipped |
| 82216b0 | chore(deps): bump github.com/go-jose/go-jose/v4 from 4.1.3 t | dependabot[bot] | [#1609](https://github.com/tektoncd/chains/pull/1609) | — | skipped |
| fc64ff4 | Nominate Emil and Abhishek are reviewers for chains repo (#1 | Anitha Natarajan | [#1621](https://github.com/tektoncd/chains/pull/1621) | — | — |
| 30ee356 | chore(deps): bump actions/upload-artifact from 7.0.0 to 7.0. | dependabot[bot] | [#1624](https://github.com/tektoncd/chains/pull/1624) | — | skipped |
| af6b652 | chore(deps): bump actions/setup-go from 6.2.0 to 6.4.0 (#160 | dependabot[bot] | [#1605](https://github.com/tektoncd/chains/pull/1605) | — | skipped |
| e9a1241 | chore(deps): bump chainguard-dev/actions from 1.6.13 to 1.6. | dependabot[bot] | [#1626](https://github.com/tektoncd/chains/pull/1626) | — | skipped |
| ab69d32 | chore(deps): bump github.com/sigstore/timestamp-authority/v2 | dependabot[bot] | [#1627](https://github.com/tektoncd/chains/pull/1627) | — | skipped |
| c38dd2b | chore(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#1618](https://github.com/tektoncd/chains/pull/1618) | — | skipped |
| 164e705 | chore(deps): bump step-security/harden-runner from 2.16.1 to | dependabot[bot] | [#1625](https://github.com/tektoncd/chains/pull/1625) | — | skipped |
| f1137de | Update example performance doc (#1636) | Emil Natan | [#1636](https://github.com/tektoncd/chains/pull/1636) | — | — |
| 9c2c3be | chore(deps): bump step-security/harden-runner from 2.17.0 to | dependabot[bot] | [#1639](https://github.com/tektoncd/chains/pull/1639) | — | skipped |
| ab6d265 | chore(deps): bump zizmorcore/zizmor-action from 0.5.2 to 0.5 | dependabot[bot] | [#1640](https://github.com/tektoncd/chains/pull/1640) | — | skipped |
| 6c9e971 | chore(deps): bump github/codeql-action from 4.35.1 to 4.35.2 | dependabot[bot] | [#1641](https://github.com/tektoncd/chains/pull/1641) | — | skipped |
| 91008ba | chore(deps): bump chainguard-dev/actions from 1.6.14 to 1.6. | dependabot[bot] | [#1642](https://github.com/tektoncd/chains/pull/1642) | — | skipped |
| 05d8255 | chore(deps): bump the all group across 1 directory with 8 up | dependabot[bot] | [#1637](https://github.com/tektoncd/chains/pull/1637) | — | skipped |
| 92649d0 | chore(deps): bump github.com/tektoncd/pipeline from 1.11.0 t | dependabot[bot] | [#1644](https://github.com/tektoncd/chains/pull/1644) | — | skipped |
| 7290445 | Handle signing OCI artifacts in *ARTIFACT_OUTPUTS (#1578) | Brad Beck | [#1578](https://github.com/tektoncd/chains/pull/1578) | — | — |
| 51dd101 | Enable Server-Side Apply for finalizers management (#1635) | Emil Natan | [#1635](https://github.com/tektoncd/chains/pull/1635) | [SRVKP-8542](https://redhat.atlassian.net/browse/SRVKP-8542) | **Enhancement:** Chains now uses Kubernetes Server-Side Apply for managing finalizers on TaskRuns *(generated)* |
| 2c46f1d | chore(deps): bump the all group with 4 updates (#1645) | dependabot[bot] | [#1645](https://github.com/tektoncd/chains/pull/1645) | — | skipped |
| 2ad9ffe | chore(deps): bump chainguard-dev/actions from 1.6.15 to 1.6. | dependabot[bot] | [#1646](https://github.com/tektoncd/chains/pull/1646) | — | skipped |
| 8182379 | chore: update doc for buildType (#1647) | Anitha Natarajan | [#1647](https://github.com/tektoncd/chains/pull/1647) | — | — |
| 99bcc3b | chore(deps): bump go.uber.org/zap from 1.27.1 to 1.28.0 in t | dependabot[bot] | [#1648](https://github.com/tektoncd/chains/pull/1648) | — | skipped |
| aaffc24 | chore(deps): bump github.com/fsnotify/fsnotify in the all gr | dependabot[bot] | [#1650](https://github.com/tektoncd/chains/pull/1650) | — | skipped |
| 334df04 | chore(deps): bump step-security/harden-runner from 2.19.0 to | dependabot[bot] | [#1655](https://github.com/tektoncd/chains/pull/1655) | — | skipped |
| 94dd4dd | chore(deps): bump the all group with 4 updates (#1652) | dependabot[bot] | [#1652](https://github.com/tektoncd/chains/pull/1652) | — | skipped |
| a53b480 | chore(deps): bump chainguard-dev/actions from 1.6.16 to 1.6. | dependabot[bot] | [#1653](https://github.com/tektoncd/chains/pull/1653) | — | skipped |
| 3ec0961 | chore(deps): bump golang.org/x/crypto in the all group (#165 | dependabot[bot] | [#1657](https://github.com/tektoncd/chains/pull/1657) | — | skipped |
| 83e3566 | chore(deps): bump github/codeql-action from 4.35.2 to 4.35.4 | dependabot[bot] | [#1658](https://github.com/tektoncd/chains/pull/1658) | — | skipped |
| 04470c3 | chore(deps): bump the all group with 4 updates (#1671) | dependabot[bot] | [#1671](https://github.com/tektoncd/chains/pull/1671) | — | skipped |
| 01c72e0 | chore(deps): bump the all group with 6 updates (#1672) | dependabot[bot] | [#1672](https://github.com/tektoncd/chains/pull/1672) | — | skipped |
| 12e6db0 | chore: bump tektoncd/plumbing cherry-pick workflow (#1673) | Abhishek Ghosh | [#1673](https://github.com/tektoncd/chains/pull/1673) | — | — |
| 2e5f961 | fix: add GHA based nightly workflow for chains (#1634) | Anitha Natarajan | [#1634](https://github.com/tektoncd/chains/pull/1634) | — | — |
| 6db27b1 | chore: update CI e2e matrix for k8s and pipelines versions ( | Abhishek Ghosh | [#1674](https://github.com/tektoncd/chains/pull/1674) | — | — |
| f1df872 | chore(deps): bump zizmorcore/zizmor-action from 0.5.3 to 0.5 | dependabot[bot] | [#1676](https://github.com/tektoncd/chains/pull/1676) | — | skipped |
| e1a7227 | chore(deps): bump step-security/harden-runner from 2.19.1 to | dependabot[bot] | [#1677](https://github.com/tektoncd/chains/pull/1677) | — | skipped |
| 7a12f22 | chore(deps): bump github/codeql-action from 4.35.4 to 4.35.5 | dependabot[bot] | [#1678](https://github.com/tektoncd/chains/pull/1678) | — | skipped |
| 4f9ad41 | chore: update chains maintainers (#1679) | Anitha Natarajan | [#1679](https://github.com/tektoncd/chains/pull/1679) | — | — |
| 4bb6198 | tekton: automate releases with Pipelines-as-Code (#1656) | Abhishek Ghosh | [#1656](https://github.com/tektoncd/chains/pull/1656) | [SRVKP-11937](https://redhat.atlassian.net/browse/SRVKP-11937) | **Release Notes Not Required:** Chains releases are now automated using Pipelines-as-Code, replacing manual tkn  *(generated)* |
| 0835671 | chore(deps): bump github/codeql-action from 4.35.5 to 4.36.0 | dependabot[bot] | [#1681](https://github.com/tektoncd/chains/pull/1681) | — | skipped |
| 9169ea0 | chore(deps): bump step-security/harden-runner from 2.19.3 to | dependabot[bot] | [#1682](https://github.com/tektoncd/chains/pull/1682) | — | skipped |
| 5f38adb | chore(deps): bump golangci/golangci-lint-action from 9.2.0 t | dependabot[bot] | [#1683](https://github.com/tektoncd/chains/pull/1683) | — | skipped |
| 3fbaa94 | Fix duplicate .att/.sig OCI layers for same-digest type hint | Abhishek Ghosh | [#1601](https://github.com/tektoncd/chains/pull/1601) | [SRVKP-11363](https://redhat.atlassian.net/browse/SRVKP-11363) | **Bug Fix:** Chains no longer produces duplicate .att and .sig OCI layers when a Task declare *(extracted)* |
| 020a29a | Add migration cleanup for SSA finalizers (#1699) | tekton-robot | [#1699](https://github.com/tektoncd/chains/pull/1699) | [SRVKP-12226](https://redhat.atlassian.net/browse/SRVKP-12226) *(keyword)* | **Bug Fix:** Adds migration cleanup logic to remove stale finalizers from in-flight objects c *(generated)* |
| 69f62a9 | chore(deps): bump github/codeql-action from 4.36.0 to 4.36.2 | dependabot[bot] | [#1718](https://github.com/tektoncd/chains/pull/1718) | — | skipped |
| 80338e0 | chore(deps): bump actions/checkout from 6.0.2 to 6.0.3 (#172 | dependabot[bot] | [#1718](https://github.com/tektoncd/chains/pull/1718) | — | skipped |

---

## tektoncd-cli

**Upstream:** tektoncd/cli · **Commits:** 71 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 8c1b774 | Bump chainguard-dev/actions from 1.6.1 to 1.6.3 | dependabot[bot] | [#2742](https://github.com/tektoncd/cli/pull/2742) | — | skipped |
| 245267f | Bump go.opentelemetry.io/otel/sdk from 1.39.0 to 1.40.0 | dependabot[bot] | [#2753](https://github.com/tektoncd/cli/pull/2753) | — | skipped |
| ed60a80 | chore: referencing the setup-kind script in the plumbing rep | Shubham Bhardwaj | [#2590](https://github.com/tektoncd/cli/pull/2590) | — | — |
| 6c7d68a | Nominate divyansh42 as cli approver | Vincent Demeester | [#2761](https://github.com/tektoncd/cli/pull/2761) | — | — |
| f588be9 | Move piyush-garg to alumni | Vincent Demeester | [#2760](https://github.com/tektoncd/cli/pull/2760) | — | — |
| 698d2ea | update goreleaser version to compatible version | Shiv Verma | [#2748](https://github.com/tektoncd/cli/pull/2748) | — | — |
| 57f5665 | Update CLI docs for v0.44.0 release | Shiv Verma | [#2748](https://github.com/tektoncd/cli/pull/2748) | — | — |
| 799e37b | Bump github/codeql-action from 4.32.3 to 4.32.5 | dependabot[bot] | [#2755](https://github.com/tektoncd/cli/pull/2755) | — | skipped |
| 59fbf67 | Bump step-security/harden-runner from 2.14.2 to 2.15.0 | dependabot[bot] | [#2757](https://github.com/tektoncd/cli/pull/2757) | — | skipped |
| 2dc3904 | Bump actions/setup-go from 6.2.0 to 6.3.0 | dependabot[bot] | [#2754](https://github.com/tektoncd/cli/pull/2754) | — | skipped |
| 3c1b260 | Bump github.com/google/go-containerregistry from 0.21.0 to 0 | dependabot[bot] | [#2759](https://github.com/tektoncd/cli/pull/2759) | — | skipped |
| 101f47e | Bump github.com/docker/cli in the go-docker-dependencies gro | dependabot[bot] | [#2763](https://github.com/tektoncd/cli/pull/2763) | — | skipped |
| ec284c2 | Bump golang.org/x/term from 0.40.0 to 0.41.0 | dependabot[bot] | [#2766](https://github.com/tektoncd/cli/pull/2766) | — | skipped |
| db16aff | Bump github.com/golangci/golangci-lint/v2 in /tools | dependabot[bot] | [#2765](https://github.com/tektoncd/cli/pull/2765) | — | skipped |
| bde8531 | Add CI summary fan-in job to presubmit CI | Vincent Demeester | [#2741](https://github.com/tektoncd/cli/pull/2741) | — | — |
| 536f99c | feat: add cherry-pick command workflow | Vincent Demeester | [#2682](https://github.com/tektoncd/cli/pull/2682) | — | — |
| 6af5d24 | Bump chainguard-dev/actions from 1.6.3 to 1.6.5 | dependabot[bot] | [#2756](https://github.com/tektoncd/cli/pull/2756) | — | skipped |
| adc8405 | Bump actions/upload-artifact from 6.0.0 to 7.0.0 | dependabot[bot] | [#2758](https://github.com/tektoncd/cli/pull/2758) | — | skipped |
| 858e61d | Bump the go-k8s-dependencies group with 3 updates | dependabot[bot] | [#2752](https://github.com/tektoncd/cli/pull/2752) | — | skipped |
| 7843065 | Change all occurences of GCS buckets with OCI buckets | adityavshinde | [#2768](https://github.com/tektoncd/cli/pull/2768) | [SRVKP-11090](https://redhat.atlassian.net/browse/SRVKP-11090) | **Release Notes Not Required** *(generated)* |
| 025ac7b | Add long flag --display-name to display the log of pipeliner | changjun | [#2450](https://github.com/tektoncd/cli/pull/2450) | — | — |
| f40a19b | refactor(logs): 将display-name参数重命名为long | changjun | [#2450](https://github.com/tektoncd/cli/pull/2450) | — | — |
| f0807af | feat: add describe subcommand in customrun command | Shiv Verma | [#2712](https://github.com/tektoncd/cli/pull/2712) | [SRVKP-11855](https://redhat.atlassian.net/browse/SRVKP-11855) *(keyword)* | **Feature:** The tkn CLI now supports a describe subcommand for CustomRuns, allowing users to *(extracted)* |
| 3d05ba3 | Remove all the occurences of chains subcommand | adityavshinde | [#2769](https://github.com/tektoncd/cli/pull/2769) | [SRVKP-11155](https://redhat.atlassian.net/browse/SRVKP-11155) | **Removed Functionality:** The chains subcommand has been removed from the tkn CLI. Users should use Tekton *(generated)* |
| 100b7cd | Bump google.golang.org/grpc from 1.78.0 to 1.79.3 | dependabot[bot] | [#2775](https://github.com/tektoncd/cli/pull/2775) | — | skipped |
| 2182852 | Bump github.com/tektoncd/pipeline from 1.9.1 to 1.9.2 | dependabot[bot] | [#2781](https://github.com/tektoncd/cli/pull/2781) | — | skipped |
| 76115cf | Bump github.com/go-jose/go-jose/v4 from 4.1.3 to 4.1.4 | dependabot[bot] | [#2789](https://github.com/tektoncd/cli/pull/2789) | — | skipped |
| 4a210b4 | Update knative and components version | Khurram Baig | [#2788](https://github.com/tektoncd/cli/pull/2788) | [SRVKP-11481](https://redhat.atlassian.net/browse/SRVKP-11481) | **Release Notes Not Required** *(generated)* |
| 5b4413c | Remove Metrics Exporter Logger due to OTEL migration | Khurram Baig | [#2788](https://github.com/tektoncd/cli/pull/2788) | [SRVKP-11481](https://redhat.atlassian.net/browse/SRVKP-11481) | **Release Notes Not Required** *(generated)* |
| 6a29abb | Migrate tracing from OpenCensus to OpenTelemetry | Khurram Baig | [#2799](https://github.com/tektoncd/cli/pull/2799) | [SRVKP-10252](https://redhat.atlassian.net/browse/SRVKP-10252) *(keyword)* | **Release Notes Not Required** *(generated)* |
| bcfc8ab | Bump go.opentelemetry.io/otel/sdk from 1.42.0 to 1.43.0 | dependabot[bot] | [#2800](https://github.com/tektoncd/cli/pull/2800) | — | skipped |
| 53ca94f | Bump github.com/docker/cli in the go-docker-dependencies gro | dependabot[bot] | [#2783](https://github.com/tektoncd/cli/pull/2783) | — | skipped |
| fc0694b | Bump go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlpt | dependabot[bot] | [#2801](https://github.com/tektoncd/cli/pull/2801) | — | skipped |
| e2355ad | Bump github.com/letsencrypt/boulder from 0.20251110.0 to 0.2 | dependabot[bot] | [#2796](https://github.com/tektoncd/cli/pull/2796) | — | skipped |
| 5f37794 | (deps) Bump go version to 1.25.8 to fix CVE-2026-25679 | divyansh42 | [#2803](https://github.com/tektoncd/cli/pull/2803) | — | — |
| cfd05cd | Bump github.com/google/go-containerregistry from 0.21.3 to 0 | dependabot[bot] | [#2797](https://github.com/tektoncd/cli/pull/2797) | — | skipped |
| 79f37b3 | Bump github.com/sigstore/cosign/v2 from 2.6.2 to 2.6.3 | dependabot[bot] | [#2798](https://github.com/tektoncd/cli/pull/2798) | — | skipped |
| cab32dd | Bump step-security/harden-runner from 2.15.0 to 2.17.0 | dependabot[bot] | [#2808](https://github.com/tektoncd/cli/pull/2808) | — | skipped |
| 785d6dd | Bump github.com/fatih/color from 1.18.0 to 1.19.0 | dependabot[bot] | [#2777](https://github.com/tektoncd/cli/pull/2777) | — | skipped |
| ed54c4c | Bump github.com/golangci/golangci-lint/v2 in /tools | dependabot[bot] | [#2778](https://github.com/tektoncd/cli/pull/2778) | — | skipped |
| 906c538 | Bump github/codeql-action from 4.32.5 to 4.35.1 | dependabot[bot] | [#2784](https://github.com/tektoncd/cli/pull/2784) | — | skipped |
| 0f918ee | Bump actions/setup-go from 6.3.0 to 6.4.0 | dependabot[bot] | [#2786](https://github.com/tektoncd/cli/pull/2786) | — | skipped |
| 17bbc68 | Bump github.com/sigstore/sigstore from 1.10.4 to 1.10.5 | dependabot[bot] | [#2805](https://github.com/tektoncd/cli/pull/2805) | — | skipped |
| 613adb7 | Bump github.com/sigstore/timestamp-authority/v2 from 2.0.3 t | dependabot[bot] | [#2809](https://github.com/tektoncd/cli/pull/2809) | — | skipped |
| 515d460 | Bump go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlp | dependabot[bot] | [#2802](https://github.com/tektoncd/cli/pull/2802) | — | skipped |
| 49ee12b | Bump actions/upload-artifact from 7.0.0 to 7.0.1 | dependabot[bot] | [#2816](https://github.com/tektoncd/cli/pull/2816) | — | skipped |
| 02d7881 | Bump step-security/harden-runner from 2.17.0 to 2.19.0 | dependabot[bot] | [#2815](https://github.com/tektoncd/cli/pull/2815) | — | skipped |
| e884dbc | Bump the go-k8s-dependencies group with 3 updates | dependabot[bot] | [#2811](https://github.com/tektoncd/cli/pull/2811) | — | skipped |
| 118b494 | Bump github.com/letsencrypt/boulder from 0.20260406.0 to 0.2 | dependabot[bot] | [#2812](https://github.com/tektoncd/cli/pull/2812) | — | skipped |
| b49dc88 | Bump github.com/google/go-containerregistry from 0.21.4 to 0 | dependabot[bot] | [#2813](https://github.com/tektoncd/cli/pull/2813) | — | skipped |
| d52be7f | Bump chainguard-dev/actions from 1.6.5 to 1.6.15 | dependabot[bot] | [#2814](https://github.com/tektoncd/cli/pull/2814) | — | skipped |
| 4e26784 | Bump github/codeql-action from 4.35.1 to 4.35.2 | dependabot[bot] | [#2817](https://github.com/tektoncd/cli/pull/2817) | — | skipped |
| adf725f | Bump github.com/docker/cli in the go-docker-dependencies gro | dependabot[bot] | [#2820](https://github.com/tektoncd/cli/pull/2820) | — | skipped |
| ae79995 | Bump github.com/tektoncd/pipeline from 1.11.0 to 1.11.1 | dependabot[bot] | [#2823](https://github.com/tektoncd/cli/pull/2823) | — | skipped |
| 573fba8 | Bump github.com/tektoncd/chains from 0.26.2 to 0.26.3 | dependabot[bot] | [#2827](https://github.com/tektoncd/cli/pull/2827) | — | skipped |
| 03c73ce | Bump chainguard-dev/actions from 1.6.15 to 1.6.16 | dependabot[bot] | [#2829](https://github.com/tektoncd/cli/pull/2829) | — | skipped |
| 66bba9d | Bump go.uber.org/zap from 1.27.1 to 1.28.0 | dependabot[bot] | [#2830](https://github.com/tektoncd/cli/pull/2830) | — | skipped |
| 7b4168b | Bump github.com/letsencrypt/boulder from 0.20260413.0 to 0.2 | dependabot[bot] | [#2821](https://github.com/tektoncd/cli/pull/2821) | — | skipped |
| 546e113 | Bump github.com/tektoncd/pipeline from 1.11.1 to 1.12.0 | dependabot[bot] | [#2836](https://github.com/tektoncd/cli/pull/2836) | — | skipped |
| a15a74b | Bump google.golang.org/grpc from 1.80.0 to 1.81.0 | dependabot[bot] | [#2837](https://github.com/tektoncd/cli/pull/2837) | — | skipped |
| 8d4599a | Bump chainguard-dev/actions from 1.6.16 to 1.6.19 | dependabot[bot] | [#2839](https://github.com/tektoncd/cli/pull/2839) | — | skipped |
| 32a465e | Bump github.com/golangci/golangci-lint/v2 in /tools | dependabot[bot] | [#2834](https://github.com/tektoncd/cli/pull/2834) | — | skipped |
| 87c35de | Bump step-security/harden-runner from 2.19.0 to 2.19.1 | dependabot[bot] | [#2840](https://github.com/tektoncd/cli/pull/2840) | — | skipped |
| 960f517 | Strip Go symbol table from release binaries | alliasgher | [#2843](https://github.com/tektoncd/cli/pull/2843) | — | — |
| afa025b | Bump github/codeql-action from 4.35.2 to 4.35.3 | dependabot[bot] | [#2838](https://github.com/tektoncd/cli/pull/2838) | — | skipped |
| 2aedc78 | Bump github.com/golangci/golangci-lint/v2 in /tools | dependabot[bot] | [#2842](https://github.com/tektoncd/cli/pull/2842) | — | skipped |
| a7d2c84 | Bump github.com/in-toto/in-toto-golang from 0.10.0 to 0.11.0 | dependabot[bot] | [#2844](https://github.com/tektoncd/cli/pull/2844) | — | skipped |
| 26317d4 | Bump github.com/docker/cli | dependabot[bot] | [#2833](https://github.com/tektoncd/cli/pull/2833) | — | skipped |
| 00c055a | Update tekton hub to 1.24.0 | divyansh42 | [#2845](https://github.com/tektoncd/cli/pull/2845) | [SRVKP-11950](https://redhat.atlassian.net/browse/SRVKP-11950) *(keyword)* | **Deprecated Functionality:** Tekton Hub support in the tkn CLI is deprecated in favor of Artifact Hub. Severa *(extracted)* |
| ce8b7d0 | Bump golang.org/x/crypto from 0.50.0 to 0.51.0 | dependabot[bot] | [#2847](https://github.com/tektoncd/cli/pull/2847) | — | skipped |
| ce6edf3 | New version v0.45.0 | Shiv Verma | — | — | — |

---

## tektoncd-pipeline

**Upstream:** tektoncd/pipeline · **Commits:** 327 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 12c973b | build(deps): bump sigstore/sigstore from 1.9.5 to 1.10.4 | Vincent Demeester | [#9331](https://github.com/tektoncd/pipeline/pull/9331) | — | — |
| c68c8d0 | build(deps): bump github.com/tektoncd/pipeline to v1.7.0 in  | Vincent Demeester | [#9329](https://github.com/tektoncd/pipeline/pull/9329) | — | — |
| f001341 | Fix wait-task-beta test infrastructure for k8s compatibility | Vincent Demeester | [#9329](https://github.com/tektoncd/pipeline/pull/9329) | — | — |
| b49da8a | Set KUBERNETES_MIN_VERSION for wait-task-beta controller | Vincent Demeester | [#9329](https://github.com/tektoncd/pipeline/pull/9329) | — | — |
| d2882c6 | build(deps): bump google.golang.org/grpc from 1.77.0 to 1.78 | dependabot[bot] | [#9337](https://github.com/tektoncd/pipeline/pull/9337) | — | skipped |
| 8db0978 | build(deps): bump github.com/spiffe/spire-api-sdk from 1.14. | dependabot[bot] | [#9336](https://github.com/tektoncd/pipeline/pull/9336) | — | skipped |
| 7783005 | docs: update releases.md for v1.9.0 LTS | Vincent Demeester | [#9339](https://github.com/tektoncd/pipeline/pull/9339) | — | — |
| d9df4b3 | docs: Document roadmap project board workflows and best prac | Vincent Demeester | [#9311](https://github.com/tektoncd/pipeline/pull/9311) | — | — |
| ada716e | build(deps): bump github.com/tektoncd/pipeline | dependabot[bot] | [#9340](https://github.com/tektoncd/pipeline/pull/9340) | — | skipped |
| 03724cc | build(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#9345](https://github.com/tektoncd/pipeline/pull/9345) | — | skipped |
| 5fe8b44 | build(deps): bump actions/cache from 4.2.3 to 5.0.3 | dependabot[bot] | [#9348](https://github.com/tektoncd/pipeline/pull/9348) | — | skipped |
| c2921a1 | build(deps): bump github/codeql-action from 4.32.0 to 4.32.1 | dependabot[bot] | [#9350](https://github.com/tektoncd/pipeline/pull/9350) | — | skipped |
| e315849 | build(deps): bump chainguard-dev/actions from 1.5.13 to 1.5. | dependabot[bot] | [#9351](https://github.com/tektoncd/pipeline/pull/9351) | — | skipped |
| 2fea6ed | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9352](https://github.com/tektoncd/pipeline/pull/9352) | — | skipped |
| 05b05ba | chore: remove redundant word | wmypku | [#9009](https://github.com/tektoncd/pipeline/pull/9009) | — | — |
| 6cd9158 | Update docs for changes in apiVersion v1 | Johan Kok | [#9042](https://github.com/tektoncd/pipeline/pull/9042) | — | — |
| 37afb99 | build(deps): bump opentelemetry exporter packages to v1.39.0 | Vincent Demeester | [#9332](https://github.com/tektoncd/pipeline/pull/9332) | — | — |
| 6ae1296 | fix(ci): simplify e2e test health status result | Anitha Natarajan | [#9361](https://github.com/tektoncd/pipeline/pull/9361) | — | — |
| abfa29d | fix: Align cache configstore with framework implementation | Seunghun Shin | [#9282](https://github.com/tektoncd/pipeline/pull/9282) | — | — |
| a386fa3 | build(deps): bump golang.org/x/crypto from 0.47.0 to 0.48.0 | dependabot[bot] | [#9369](https://github.com/tektoncd/pipeline/pull/9369) | — | skipped |
| 5696458 | build(deps): bump chainguard-dev/actions from 1.5.14 to 1.5. | dependabot[bot] | [#9370](https://github.com/tektoncd/pipeline/pull/9370) | — | skipped |
| 635e378 | build(deps): bump tj-actions/changed-files from 47.0.1 to 47 | dependabot[bot] | [#9371](https://github.com/tektoncd/pipeline/pull/9371) | — | skipped |
| 3790175 | build(deps): bump step-security/harden-runner from 2.14.1 to | dependabot[bot] | [#9372](https://github.com/tektoncd/pipeline/pull/9372) | — | skipped |
| a5c1907 | build(deps): bump github/codeql-action from 4.32.1 to 4.32.2 | dependabot[bot] | [#9374](https://github.com/tektoncd/pipeline/pull/9374) | — | skipped |
| 6206ff7 | ci: add /rebase slash command workflow | Vibhav Bobade | [#9375](https://github.com/tektoncd/pipeline/pull/9375) | — | — |
| 2cef3c7 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9373](https://github.com/tektoncd/pipeline/pull/9373) | — | skipped |
| fedb5f7 | build(deps): bump k8s.io/apimachinery | dependabot[bot] | [#9380](https://github.com/tektoncd/pipeline/pull/9380) | — | skipped |
| b91c6fa | build(deps): bump k8s.io/apimachinery from 0.33.7 to 0.33.8 | dependabot[bot] | [#9384](https://github.com/tektoncd/pipeline/pull/9384) | — | skipped |
| d52f146 | fix(pipelines): allow pipeline param defaults to use non-par | Andrew Thorp | [#9386](https://github.com/tektoncd/pipeline/pull/9386) | [SRVKP-10840](https://redhat.atlassian.net/browse/SRVKP-10840) | **Bug Fix:** Fixed a bug which caused PipelineRun validation to fail when a pipeline paramete *(extracted)* |
| 92df9ee | build(deps): bump k8s.io/api in /test/custom-task-ctrls/wait | dependabot[bot] | [#9381](https://github.com/tektoncd/pipeline/pull/9381) | — | skipped |
| c8f8633 | build(deps): bump k8s.io/client-go from 0.32.11 to 0.32.12 | dependabot[bot] | [#9383](https://github.com/tektoncd/pipeline/pull/9383) | — | skipped |
| 7c92f82 | build(deps): bump google.golang.org/grpc from 1.78.0 to 1.79 | dependabot[bot] | [#9389](https://github.com/tektoncd/pipeline/pull/9389) | — | skipped |
| 30d2462 | build(deps): bump k8s.io/code-generator from 0.32.11 to 0.32 | dependabot[bot] | [#9388](https://github.com/tektoncd/pipeline/pull/9388) | — | skipped |
| 2df9fbe | feat: Add SHA-256 support for Git resolver revision validati | 7h3-3mp7y-m4n | [#9278](https://github.com/tektoncd/pipeline/pull/9278) | — | — |
| c808c3d | build(deps): bump k8s.io/apiextensions-apiserver from 0.32.1 | dependabot[bot] | [#9385](https://github.com/tektoncd/pipeline/pull/9385) | — | skipped |
| 579ccaf | ci: replace e2e-only fan-in with unified CI summary job | Vincent Demeester | [#9394](https://github.com/tektoncd/pipeline/pull/9394) | — | — |
| 597194e | build(deps): bump google.golang.org/grpc from 1.79.0 to 1.79 | dependabot[bot] | [#9392](https://github.com/tektoncd/pipeline/pull/9392) | — | skipped |
| 99a7d11 | build(deps): bump github.com/jenkins-x/go-scm from 1.15.16 t | dependabot[bot] | [#9391](https://github.com/tektoncd/pipeline/pull/9391) | — | skipped |
| 294e143 | build(deps): bump chainguard-dev/actions from 1.5.16 to 1.6. | dependabot[bot] | [#9395](https://github.com/tektoncd/pipeline/pull/9395) | — | skipped |
| 31a2e31 | tekton: update plumbing ref to include full image references | Vincent Demeester | [#9399](https://github.com/tektoncd/pipeline/pull/9399) | — | — |
| 5d50dad | ci: Update cherry-pick command to latest plumbing | Vincent Demeester | [#9400](https://github.com/tektoncd/pipeline/pull/9400) | — | — |
| 0c3608f | fix: Remove redundant shortNames from ResolutionRequest CRD | Vincent Demeester | [#9398](https://github.com/tektoncd/pipeline/pull/9398) | — | — |
| c0bef5f | tekton: update plumbing ref to latest commit | Vincent Demeester | [#9406](https://github.com/tektoncd/pipeline/pull/9406) | — | — |
| 5d0e97e | docs: clarify flag availability across controller binaries | Lucas Aguiar | [#9390](https://github.com/tektoncd/pipeline/pull/9390) | — | — |
| 5954b37 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9397](https://github.com/tektoncd/pipeline/pull/9397) | — | skipped |
| 3f11e1a | build(deps): bump github/codeql-action from 4.32.2 to 4.32.3 | dependabot[bot] | [#9396](https://github.com/tektoncd/pipeline/pull/9396) | — | skipped |
| e7755b3 | build(deps): bump github.com/tektoncd/pipeline | dependabot[bot] | [#9417](https://github.com/tektoncd/pipeline/pull/9417) | — | skipped |
| a3e1e37 | build(deps): bump k8s.io/client-go | dependabot[bot] | [#9382](https://github.com/tektoncd/pipeline/pull/9382) | — | skipped |
| becdc2a | build(deps): bump github.com/google/go-containerregistry | dependabot[bot] | [#9418](https://github.com/tektoncd/pipeline/pull/9418) | — | skipped |
| a245af3 | build(deps): bump tj-actions/changed-files from 47.0.2 to 47 | dependabot[bot] | [#9429](https://github.com/tektoncd/pipeline/pull/9429) | — | skipped |
| 8b62740 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9430](https://github.com/tektoncd/pipeline/pull/9430) | — | skipped |
| 71065cd | build(deps): bump github/codeql-action from 4.32.3 to 4.32.4 | dependabot[bot] | [#9431](https://github.com/tektoncd/pipeline/pull/9431) | — | skipped |
| c281693 | build(deps): bump chainguard-dev/actions from 1.6.1 to 1.6.4 | dependabot[bot] | [#9427](https://github.com/tektoncd/pipeline/pull/9427) | — | skipped |
| a9f0a88 | build(deps): bump actions/dependency-review-action from 4.8. | dependabot[bot] | [#9428](https://github.com/tektoncd/pipeline/pull/9428) | — | skipped |
| f759308 | build(deps): bump golang.org/x/crypto in /test/resolver-with | dependabot[bot] | [#9426](https://github.com/tektoncd/pipeline/pull/9426) | — | skipped |
| 4d10d23 | Assess several new gosec findings | Sascha Schwarze | [#9405](https://github.com/tektoncd/pipeline/pull/9405) | — | — |
| 8e4f6d6 | build(deps): bump github.com/sigstore/sigstore | dependabot[bot] | [#9425](https://github.com/tektoncd/pipeline/pull/9425) | — | skipped |
| e174ae9 | build(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#9363](https://github.com/tektoncd/pipeline/pull/9363) | — | skipped |
| cb43282 | build(deps): bump github.com/google/go-containerregistry | dependabot[bot] | [#9433](https://github.com/tektoncd/pipeline/pull/9433) | — | skipped |
| 7a2e474 | Move v0.68 LTS to End of Life releases | Vincent Demeester | [#9434](https://github.com/tektoncd/pipeline/pull/9434) | — | — |
| 5b46423 | Bump Knative/pkg to latest | Khurram Baig | [#9043](https://github.com/tektoncd/pipeline/pull/9043) | [SRVKP-9321](https://redhat.atlassian.net/browse/SRVKP-9321) | **Release Note Not Required** *(extracted)* |
| 81d71ac | Update generated code and CRDs | Khurram Baig | [#9043](https://github.com/tektoncd/pipeline/pull/9043) | [SRVKP-9321](https://redhat.atlassian.net/browse/SRVKP-9321) | **Release Note Not Required** *(extracted)* |
| 17d93c7 | fix: update unit test checksums | Khurram Baig | [#9043](https://github.com/tektoncd/pipeline/pull/9043) | [SRVKP-9321](https://redhat.atlassian.net/browse/SRVKP-9321) | **Release Note Not Required** *(extracted)* |
| 2ce78df | refactor: use observability package for metrics config name | Khurram Baig | [#9043](https://github.com/tektoncd/pipeline/pull/9043) | [SRVKP-9321](https://redhat.atlassian.net/browse/SRVKP-9321) | **Release Note Not Required** *(extracted)* |
| c8d3f4b | Remove old metrics code from watcher | Khurram Baig | [#9043](https://github.com/tektoncd/pipeline/pull/9043) | [SRVKP-9321](https://redhat.atlassian.net/browse/SRVKP-9321) | **Release Note Not Required** *(extracted)* |
| d679848 | test: Drop OpenCensus based metric exporter initialisation | Khurram Baig | [#9043](https://github.com/tektoncd/pipeline/pull/9043) | [SRVKP-9321](https://redhat.atlassian.net/browse/SRVKP-9321) | **Release Note Not Required** *(extracted)* |
| 72c59cd | feat: Migrate PipelineRun and TaskRun metrics to OpenTelemet | Khurram Baig | [#9043](https://github.com/tektoncd/pipeline/pull/9043) | [SRVKP-9321](https://redhat.atlassian.net/browse/SRVKP-9321) | **Release Note Not Required** *(extracted)* |
| a434edc | fix: include results and artifacts from failed tasks in pipe | Vibhav Bobade | [#9367](https://github.com/tektoncd/pipeline/pull/9367) | [SRVKP-1513](https://redhat.atlassian.net/browse/SRVKP-1513) *(unmatched)* | — |
| a5a9c6f | accept featureFlags.EnableTektonOCIBundles to fix unknown fi | Niklas Siltakorpi | [#8996](https://github.com/tektoncd/pipeline/pull/8996) | — | — |
| ebf64f3 | add a explanation comment and fix casing in the yaml tag | Niklas Siltakorpi | [#8996](https://github.com/tektoncd/pipeline/pull/8996) | — | — |
| a27b0dc | Fix gofmt and tag alignment | Niklas Siltakorpi | [#8996](https://github.com/tektoncd/pipeline/pull/8996) | — | — |
| 9db88e0 | chore: regenerate CRDs with field descriptions | Vibhav Bobade | [#8996](https://github.com/tektoncd/pipeline/pull/8996) | — | — |
| 15b0933 | docs: update releases.md for v1.10.0 | Vincent Demeester | [#9448](https://github.com/tektoncd/pipeline/pull/9448) | — | — |
| 1656835 | build(deps): bump k8s.io/client-go from 0.34.3 to 0.34.4 | dependabot[bot] | [#9447](https://github.com/tektoncd/pipeline/pull/9447) | — | skipped |
| 94775c6 | build(deps): bump go.opentelemetry.io/otel/trace from 1.39.0 | dependabot[bot] | [#9445](https://github.com/tektoncd/pipeline/pull/9445) | — | skipped |
| b0a3970 | Enable Gitea e2e tests in CI | ab-ghosh | [#9442](https://github.com/tektoncd/pipeline/pull/9442) | [SRVKP-10929](https://redhat.atlassian.net/browse/SRVKP-10929) | **Release Notes Not Required** *(extracted)* |
| 4648b3e | Update release-cheat-sheet.md and release pipeline | Naomi Gelman | [#9415](https://github.com/tektoncd/pipeline/pull/9415) | — | — |
| 82fc97b | build(deps): bump github.com/tektoncd/pipeline | dependabot[bot] | [#9453](https://github.com/tektoncd/pipeline/pull/9453) | — | skipped |
| 1280f9b | build(deps): bump k8s.io/apiextensions-apiserver from 0.34.3 | dependabot[bot] | [#9455](https://github.com/tektoncd/pipeline/pull/9455) | — | skipped |
| cd0046c | fix: revert mistaken metadata changes in resolvers config-ob | Khurram Baig | [#9468](https://github.com/tektoncd/pipeline/pull/9468) | — | — |
| e712e56 | chore: bump knative.dev/pkg to main and k8s libs to 0.35.1 | Vincent Demeester | [#9470](https://github.com/tektoncd/pipeline/pull/9470) | [SRVKP-9318](https://redhat.atlassian.net/browse/SRVKP-9318) *(unmatched)* | — |
| 9ad8e4b | chore: regenerate code after knative/pkg and k8s bump | Vincent Demeester | [#9470](https://github.com/tektoncd/pipeline/pull/9470) | [SRVKP-9318](https://redhat.atlassian.net/browse/SRVKP-9318) *(unmatched)* | — |
| ab95bd9 | build(deps): bump chainguard-dev/actions from 1.6.4 to 1.6.5 | dependabot[bot] | [#9479](https://github.com/tektoncd/pipeline/pull/9479) | — | skipped |
| 503e0e2 | build(deps): bump actions/setup-go from 6.2.0 to 6.3.0 | dependabot[bot] | [#9480](https://github.com/tektoncd/pipeline/pull/9480) | — | skipped |
| b09c411 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9483](https://github.com/tektoncd/pipeline/pull/9483) | — | skipped |
| 4fa7d28 | build(deps): bump k8s.io/api in /test/custom-task-ctrls/wait | dependabot[bot] | [#9472](https://github.com/tektoncd/pipeline/pull/9472) | — | skipped |
| bfe5e42 | build(deps): bump github/codeql-action from 4.32.4 to 4.32.5 | dependabot[bot] | [#9482](https://github.com/tektoncd/pipeline/pull/9482) | — | skipped |
| 3904b7b | build(deps): bump actions/upload-artifact from 6.0.0 to 7.0. | dependabot[bot] | [#9478](https://github.com/tektoncd/pipeline/pull/9478) | — | skipped |
| 374728c | build(deps): bump k8s.io/apimachinery from 0.35.1 to 0.35.2 | dependabot[bot] | [#9476](https://github.com/tektoncd/pipeline/pull/9476) | — | skipped |
| 558559e | build(deps): bump k8s.io/code-generator from 0.35.1 to 0.35. | dependabot[bot] | [#9473](https://github.com/tektoncd/pipeline/pull/9473) | — | skipped |
| 42a6789 | build(deps): bump go.opentelemetry.io/otel/metric from 1.40. | dependabot[bot] | [#9477](https://github.com/tektoncd/pipeline/pull/9477) | — | skipped |
| 9ad6179 | build(deps): bump step-security/harden-runner from 2.14.2 to | dependabot[bot] | [#9481](https://github.com/tektoncd/pipeline/pull/9481) | — | skipped |
| 520a2ff | docs: replace 'coming soon' with tkn bundle link in taskruns | Chinonso Nwakudu | [#9509](https://github.com/tektoncd/pipeline/pull/9509) | — | — |
| 330ad67 | build(deps): bump k8s.io/client-go | dependabot[bot] | [#9475](https://github.com/tektoncd/pipeline/pull/9475) | — | skipped |
| d1379f8 | fix: propagate PipelineRun tasks/finally timeout to child Ta | Shubham Bhardwaj | [#9419](https://github.com/tektoncd/pipeline/pull/9419) | [SRVKP-8963](https://redhat.atlassian.net/browse/SRVKP-8963) | **Bug Fix:** PipelineRun tasks/finally timeouts are now correctly propagated to child TaskRun *(generated)* |
| 5c74658 | Move inactive approvers to alumni | Vincent Demeester | [#9518](https://github.com/tektoncd/pipeline/pull/9518) | — | — |
| b17d57f | build(deps): bump k8s.io/apiextensions-apiserver from 0.35.1 | dependabot[bot] | [#9487](https://github.com/tektoncd/pipeline/pull/9487) | — | skipped |
| bf12c8d | Nominate khrm and aThorp96 as pipeline approvers | Vincent Demeester | [#9519](https://github.com/tektoncd/pipeline/pull/9519) | — | — |
| c2d66eb | build(deps): bump step-security/harden-runner from 2.15.0 to | dependabot[bot] | [#9542](https://github.com/tektoncd/pipeline/pull/9542) | — | skipped |
| 2d4e464 | build(deps): bump actions/dependency-review-action from 4.8. | dependabot[bot] | [#9536](https://github.com/tektoncd/pipeline/pull/9536) | — | skipped |
| 98a0611 | build(deps): bump github/codeql-action from 4.32.5 to 4.32.6 | dependabot[bot] | [#9538](https://github.com/tektoncd/pipeline/pull/9538) | — | skipped |
| d9b4f2b | build(deps): bump chainguard-dev/actions from 1.6.5 to 1.6.7 | dependabot[bot] | [#9539](https://github.com/tektoncd/pipeline/pull/9539) | — | skipped |
| a6c785a | build(deps): bump tj-actions/changed-files from 47.0.4 to 47 | dependabot[bot] | [#9540](https://github.com/tektoncd/pipeline/pull/9540) | — | skipped |
| c53b01c | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9541](https://github.com/tektoncd/pipeline/pull/9541) | — | skipped |
| 063e9b9 | docs(auth): clean stale TODO | 比泽尔 | [#9504](https://github.com/tektoncd/pipeline/pull/9504) | — | — |
| 9d7743e | Update tj-actions/changed-files version comment to v47.0.5 | Vincent Demeester | [#9552](https://github.com/tektoncd/pipeline/pull/9552) | — | — |
| 933142f | Remove opencensus dependency from test files | Khurram Baig | [#9553](https://github.com/tektoncd/pipeline/pull/9553) | — | — |
| 04c0c7f | fix: update default tracing endpoint to oltp protobuf endpoi | Vibhav Bobade | [#9141](https://github.com/tektoncd/pipeline/pull/9141) | — | — |
| 5c26041 | Fix running_taskruns metric overcounting TaskRuns with no co | Teleфоnчиk | [#9485](https://github.com/tektoncd/pipeline/pull/9485) | — | — |
| a5f56db | fix: imports | Teleфоnчиk | [#9485](https://github.com/tektoncd/pipeline/pull/9485) | — | — |
| 2ebc5cf | fix: guard nil succeeded condition in pipelinerun metrics | Teleфоnчиk | [#9485](https://github.com/tektoncd/pipeline/pull/9485) | — | — |
| 5eead3f | Fix panic in GenerateDeterministicNameFromSpec with long res | Vincent Demeester | — | — | — |
| a8fc41b | fix: use hashed volume names to prevent credential volume na | Vibhav Bobade | [#9528](https://github.com/tektoncd/pipeline/pull/9528) | [SRVKP-6798](https://redhat.atlassian.net/browse/SRVKP-6798) *(keyword)* | **Bug Fix:** Credential volume names are now generated using deterministic hashing to prevent *(generated)* |
| cf5b657 | fix: shorten credential volume name prefix and use full FNV- | Vibhav Bobade | [#9528](https://github.com/tektoncd/pipeline/pull/9528) | [SRVKP-6798](https://redhat.atlassian.net/browse/SRVKP-6798) *(keyword)* | **Release Notes Not Required** *(generated)* |
| f7d4c4f | fix: convert pod latency metric to histogram and remove pod  | Vibhav Bobade | [#9530](https://github.com/tektoncd/pipeline/pull/9530) | [SRVKP-6762](https://redhat.atlassian.net/browse/SRVKP-6762) *(keyword)* | **Bug Fix:** The tekton_pipelines_controller_taskruns_pod_latency_milliseconds metric is now  *(generated)* |
| 4cae928 | docs: add pod latency histogram metric to metrics.md | Vibhav Bobade | [#9530](https://github.com/tektoncd/pipeline/pull/9530) | [SRVKP-6762](https://redhat.atlassian.net/browse/SRVKP-6762) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 41db6d8 | [TEP-0137] Restructure customrun event controller | Andrea Frittoli | [#6889](https://github.com/tektoncd/pipeline/pull/6889) | — | — |
| 0a348fc | [TEP-0137] Add missing test coverage for PR 1 changes | Andrea Frittoli | [#6889](https://github.com/tektoncd/pipeline/pull/6889) | — | — |
| 25dde8a | build(deps): bump github.com/google/go-containerregistry | dependabot[bot] | [#9548](https://github.com/tektoncd/pipeline/pull/9548) | — | skipped |
| 961388f | fix: prevent path traversal in git resolver pathInRepo param | Vincent Demeester | — | — | — |
| 468107d | build(deps): bump google.golang.org/grpc from 1.79.1 to 1.79 | dependabot[bot] | [#9547](https://github.com/tektoncd/pipeline/pull/9547) | — | skipped |
| 6296083 | build(deps): bump go.opentelemetry.io/otel/trace from 1.41.0 | dependabot[bot] | [#9549](https://github.com/tektoncd/pipeline/pull/9549) | — | skipped |
| 1c876f3 | build(deps): bump chainguard-dev/actions from 1.6.7 to 1.6.8 | dependabot[bot] | [#9583](https://github.com/tektoncd/pipeline/pull/9583) | — | skipped |
| 9207d82 | build(deps): bump step-security/harden-runner from 2.15.1 to | dependabot[bot] | [#9584](https://github.com/tektoncd/pipeline/pull/9584) | — | skipped |
| 52d706a | build(deps): bump fgrosse/go-coverage-report from 1.2.0 to 1 | dependabot[bot] | [#9585](https://github.com/tektoncd/pipeline/pull/9585) | — | skipped |
| e699c97 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9587](https://github.com/tektoncd/pipeline/pull/9587) | — | skipped |
| 079dee2 | build(deps): bump golang.org/x/sync from 0.19.0 to 0.20.0 | dependabot[bot] | [#9611](https://github.com/tektoncd/pipeline/pull/9611) | — | skipped |
| 799ce5f | build(deps): bump github/codeql-action from 4.32.6 to 4.33.0 | dependabot[bot] | [#9586](https://github.com/tektoncd/pipeline/pull/9586) | — | skipped |
| 3bbd955 | Re-enable pipeline-api.md generation | Andrea Frittoli | [#9604](https://github.com/tektoncd/pipeline/pull/9604) | — | — |
| 4fe69bd | fix: Pin Bash container images to digests in examples/v1/tas | Aiswarya R | [#9610](https://github.com/tektoncd/pipeline/pull/9610) | — | — |
| c1e5237 | fix: Pin alpine container images to digests in examples/v1/t | Aiswarya R | [#9610](https://github.com/tektoncd/pipeline/pull/9610) | — | — |
| d537d17 | fix: Pin busybox container images to digests in examples/v1/ | Aiswarya R | [#9610](https://github.com/tektoncd/pipeline/pull/9610) | — | — |
| 15a50eb | Add explicit permissions blocks to workflows missing them | Parameshwaran Krishnasamy | [#9562](https://github.com/tektoncd/pipeline/pull/9562) | — | — |
| 9d3a656 | docs: update releases.md with security patch releases | Vincent Demeester | [#9616](https://github.com/tektoncd/pipeline/pull/9616) | — | — |
| 61edf0a | fix: Pin Ubuntu container images to digests in examples/v1/t | Aiswarya R | [#9607](https://github.com/tektoncd/pipeline/pull/9607) | — | — |
| c620809 | fix: Pin alpine-git-nonroot container images to digests in e | Aiswarya R | [#9614](https://github.com/tektoncd/pipeline/pull/9614) | — | — |
| b989228 | fix: Pin alpine/git container images to digests in examples/ | Aiswarya R | [#9614](https://github.com/tektoncd/pipeline/pull/9614) | — | — |
| 63f7f0e | fix: Pin cgr.dev/chainguard/busybox container images to dige | Aiswarya R | [#9614](https://github.com/tektoncd/pipeline/pull/9614) | — | — |
| 6553ba5 | fix: Pin nop container images to digests in examples/v1/task | Aiswarya R | [#9614](https://github.com/tektoncd/pipeline/pull/9614) | — | — |
| 2c39871 | perf(pipelinerun): hoist VerificationPolicy list out of per- | Ankur Sinha | [#9601](https://github.com/tektoncd/pipeline/pull/9601) | [SRVKP-11151](https://redhat.atlassian.net/browse/SRVKP-11151) | **Release Note Not Required** *(generated)* |
| d989f57 | Expand section on environment variables | jcjveraa | [#8620](https://github.com/tektoncd/pipeline/pull/8620) | — | — |
| f8d1e06 | add secret-env-from.yaml, correct section headers in tasks.m | jelle | [#8620](https://github.com/tektoncd/pipeline/pull/8620) | — | — |
| aaf74a8 | Enhance secret-env-from script to produce correct Basic Auth | jcjveraa | [#8620](https://github.com/tektoncd/pipeline/pull/8620) | — | — |
| 8b420ab | cleanup: replace GCS release URLs with infra.tekton.dev | Ankur Sinha | [#9569](https://github.com/tektoncd/pipeline/pull/9569) | [SRVKP-11096](https://redhat.atlassian.net/browse/SRVKP-11096) | **Release Note Not Required** *(generated)* |
| f95fc30 | fix: Pin container images to digests in examples/v1/taskruns | Aiswarya R | [#9618](https://github.com/tektoncd/pipeline/pull/9618) | — | — |
| 9adac1a | fix: Pin container images to digests in examples/v1/taskruns | Aiswarya R | [#9618](https://github.com/tektoncd/pipeline/pull/9618) | — | — |
| 24afb9b | fix: resolve context key collision and ownerRef nil panic in | Vibhav Bobade | [#9593](https://github.com/tektoncd/pipeline/pull/9593) | — | — |
| 507e16d | fix: address review feedback - use apiequality, expand tests | Vibhav Bobade | [#9593](https://github.com/tektoncd/pipeline/pull/9593) | — | — |
| c011286 | build(deps): bump github.com/google/go-containerregistry | dependabot[bot] | [#9627](https://github.com/tektoncd/pipeline/pull/9627) | — | skipped |
| 931861b | build(deps): bump google.golang.org/grpc from 1.79.2 to 1.79 | dependabot[bot] | [#9628](https://github.com/tektoncd/pipeline/pull/9628) | — | skipped |
| 24e6ce0 | build(deps): bump github.com/spiffe/spire-api-sdk from 1.14. | dependabot[bot] | [#9629](https://github.com/tektoncd/pipeline/pull/9629) | — | skipped |
| 22c300c | build(deps): bump github.com/tektoncd/pipeline | dependabot[bot] | [#9626](https://github.com/tektoncd/pipeline/pull/9626) | — | skipped |
| e8f7a28 | docs: add 4 undocumented metrics to docs/metrics.md | Chinonso Nwakudu | [#9512](https://github.com/tektoncd/pipeline/pull/9512) | — | — |
| c408cbd | docs: fix broken internal markdown links | Goutham-AR | [#9507](https://github.com/tektoncd/pipeline/pull/9507) | — | — |
| 97c62a5 | doc: Fix broken Tekton Bundles example link in taskruns.md | sahilleth | [#9462](https://github.com/tektoncd/pipeline/pull/9462) | — | — |
| 64a339b | doc: Clarify scope of auth documentation | sahilleth | [#9461](https://github.com/tektoncd/pipeline/pull/9461) | — | — |
| 838e536 | Fix: Add useHttpPath to support multiple Git credentials on  | ab-ghosh | [#9143](https://github.com/tektoncd/pipeline/pull/9143) | [SRVKP-9280](https://redhat.atlassian.net/browse/SRVKP-9280) | **Bug Fix:** Tekton now supports multiple Git credentials on the same host by using useHttpPa *(generated)* |
| 52916b4 | build(deps): bump k8s.io/api in /test/custom-task-ctrls/wait | dependabot[bot] | [#9637](https://github.com/tektoncd/pipeline/pull/9637) | — | skipped |
| 233b574 | build(deps): bump k8s.io/apimachinery from 0.35.2 to 0.35.3 | dependabot[bot] | [#9639](https://github.com/tektoncd/pipeline/pull/9639) | — | skipped |
| f2769b9 | build(deps): bump k8s.io/client-go from 0.35.2 to 0.35.3 | dependabot[bot] | [#9638](https://github.com/tektoncd/pipeline/pull/9638) | — | skipped |
| 0bfd458 | docs: add README files for pipelinerun and taskrun examples | anirudh242 | [#9505](https://github.com/tektoncd/pipeline/pull/9505) | — | — |
| 1ea381e | docs: remove file listings and reorder README sections based | anirudh242 | [#9505](https://github.com/tektoncd/pipeline/pull/9505) | — | — |
| 418dd19 | fix: Upgrade Gitea test infrastructure from v1.17.1 to lates | Sri Vignesh | [#9568](https://github.com/tektoncd/pipeline/pull/9568) | — | — |
| 91992fa | build(deps): bump actions/cache from 5.0.3 to 5.0.4 | dependabot[bot] | [#9650](https://github.com/tektoncd/pipeline/pull/9650) | — | skipped |
| 2e92eb6 | build(deps): bump github.com/spiffe/spire-api-sdk from 1.14. | dependabot[bot] | [#9648](https://github.com/tektoncd/pipeline/pull/9648) | — | skipped |
| 5534b13 | build(deps): bump chainguard-dev/actions from 1.6.8 to 1.6.9 | dependabot[bot] | [#9649](https://github.com/tektoncd/pipeline/pull/9649) | — | skipped |
| 5a67613 | build(deps): bump github/codeql-action from 4.33.0 to 4.34.1 | dependabot[bot] | [#9651](https://github.com/tektoncd/pipeline/pull/9651) | — | skipped |
| 5f15773 | fix: resolve cache race with singleflight | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) | [SRVKP-10234](https://redhat.atlassian.net/browse/SRVKP-10234) *(keyword)* | **Bug Fix:** Concurrent resolver cache requests for the same resource are now deduplicated us *(generated)* |
| 5df99a0 | fix: remove corrupted cache entries on type error | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) | [SRVKP-10234](https://redhat.atlassian.net/browse/SRVKP-10234) *(keyword)* | **Release Notes Not Required** *(generated)* |
| b7ec51f | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9652](https://github.com/tektoncd/pipeline/pull/9652) | — | skipped |
| 546de63 | Bump knative.dev/pkg to adopt centralized WEBHOOK_* TLS conf | Jawed khelil | [#9466](https://github.com/tektoncd/pipeline/pull/9466) | [SRVKP-9614](https://redhat.atlassian.net/browse/SRVKP-9614) | **Feature:** The Tekton webhook now supports centrally managed TLS configuration via WEBHOOK_ *(generated)* |
| 24766ef | fix: record metrics for cancelled PipelineRuns | Khurram Baig | [#9658](https://github.com/tektoncd/pipeline/pull/9658) | [SRVKP-11099](https://redhat.atlassian.net/browse/SRVKP-11099) | **Bug Fix:** Cancelled PipelineRuns are now correctly recorded in the tekton_pipelines_contro *(extracted)* |
| 7e3e7cc | Update stale comment about storing TaskSpec in status | Andrea Frittoli | [#9661](https://github.com/tektoncd/pipeline/pull/9661) | — | — |
| 00fda97 | build(deps): bump k8s.io/client-go | dependabot[bot] | [#9634](https://github.com/tektoncd/pipeline/pull/9634) | — | skipped |
| 273875b | fix: cluster resolver namespace access control whitespace an | Vibhav Bobade | [#9592](https://github.com/tektoncd/pipeline/pull/9592) | — | — |
| 8b5459b | fix: address review feedback - remove scope creep, add empty | Vibhav Bobade | [#9592](https://github.com/tektoncd/pipeline/pull/9592) | — | — |
| 6503458 | [#9575] Extract memberOfLookup to reduce nesting | adityavshinde | [#9596](https://github.com/tektoncd/pipeline/pull/9596) | [SRVKP-11152](https://redhat.atlassian.net/browse/SRVKP-11152) | **Release Note Not Required** *(extracted)* |
| dbb065f | fix: replace silent "default" namespace fallback with explic | Vibhav Bobade | [#9594](https://github.com/tektoncd/pipeline/pull/9594) | — | — |
| 6f60970 | fix: address review feedback - fix error label, strengthen t | Vibhav Bobade | [#9594](https://github.com/tektoncd/pipeline/pull/9594) | — | — |
| aad74ce | fix: address review feedback on GetNameAndNamespace | Vibhav Bobade | [#9594](https://github.com/tektoncd/pipeline/pull/9594) | — | — |
| 76d413c | fix: gofmt formatting in name_test.go | Vibhav Bobade | [#9594](https://github.com/tektoncd/pipeline/pull/9594) | — | — |
| e882250 | Add pending status support for TaskRun (parity with Pipeline | sahilleth | [#9464](https://github.com/tektoncd/pipeline/pull/9464) | — | — |
| 794c7c2 | hack: update pipeline api docs for TaskRunPending | Sahil Bhardwaj | [#9464](https://github.com/tektoncd/pipeline/pull/9464) | — | — |
| 72b1678 | docs: clarify timeout while TaskRun is pending | Sahil Bhardwaj | [#9464](https://github.com/tektoncd/pipeline/pull/9464) | — | — |
| 291fdd6 | tests: add pending TaskRun lifecycle transition coverage | Sahil Bhardwaj | [#9464](https://github.com/tektoncd/pipeline/pull/9464) | — | — |
| 96c066b | ci: fix GitHub Actions security issues found by zizmor | Vincent Demeester | [#9667](https://github.com/tektoncd/pipeline/pull/9667) | — | — |
| 2ba1b9f | ci: add zizmor GitHub Actions security analysis | Vincent Demeester | [#9667](https://github.com/tektoncd/pipeline/pull/9667) | — | — |
| 1b4a945 | ci: fix remaining zizmor findings (permissions, injection, a | Vincent Demeester | [#9667](https://github.com/tektoncd/pipeline/pull/9667) | — | — |
| bc5524c | fix: make step-init symlink creation idempotent | Vibhav Bobade | [#9600](https://github.com/tektoncd/pipeline/pull/9600) | — | — |
| 4b8046e | fix: handle os.Remove errors and add init idempotency test | Vibhav Bobade | [#9600](https://github.com/tektoncd/pipeline/pull/9600) | — | — |
| 6d461ee | fix: use os.IsExist instead of errors.Is per review feedback | Vibhav Bobade | [#9600](https://github.com/tektoncd/pipeline/pull/9600) | — | — |
| 0e9378b | feat: add optional PVC auto-cleanup annotation for workspace | Vincent Demeester | [#9354](https://github.com/tektoncd/pipeline/pull/9354) | [SRVKP-10219](https://redhat.atlassian.net/browse/SRVKP-10219) *(keyword)* | **Feature:** A new opt-in annotation tekton.dev/auto-cleanup-pvc enables automatic PVC cleanu *(generated)* |
| 2530f24 | Add multi-URL support and per-resolution url param to Hub Re | ab-ghosh | [#9465](https://github.com/tektoncd/pipeline/pull/9465) | [SRVKP-10955](https://redhat.atlassian.net/browse/SRVKP-10955) | **Feature:** The Hub Resolver now supports multiple hub URLs with ordered fallback via artifa *(generated)* |
| 383d57b | Fix: Add SSH Host aliases to support multiple SSH credential | ab-ghosh | [#9643](https://github.com/tektoncd/pipeline/pull/9643) | [SRVKP-7882](https://redhat.atlassian.net/browse/SRVKP-7882) | **Bug Fix:** Tekton now supports multiple SSH credentials on the same host by generating SSH  *(generated)* |
| fc5cc26 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9694](https://github.com/tektoncd/pipeline/pull/9694) | — | skipped |
| e6b37b1 | build(deps): bump github/codeql-action from 4.34.1 to 4.35.1 | dependabot[bot] | [#9691](https://github.com/tektoncd/pipeline/pull/9691) | — | skipped |
| e629f38 | build(deps): bump github.com/tektoncd/pipeline | dependabot[bot] | [#9689](https://github.com/tektoncd/pipeline/pull/9689) | — | skipped |
| a08d842 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#9690](https://github.com/tektoncd/pipeline/pull/9690) | — | skipped |
| 3cb545d | build(deps): bump chainguard-dev/actions from 1.6.9 to 1.6.1 | dependabot[bot] | [#9693](https://github.com/tektoncd/pipeline/pull/9693) | — | skipped |
| 502f3ef | build(deps): bump google.golang.org/grpc in /test/resolver-w | dependabot[bot] | [#9710](https://github.com/tektoncd/pipeline/pull/9710) | — | skipped |
| 5145e61 | build(deps): bump github.com/sigstore/sigstore from 1.10.4 t | dependabot[bot] | [#9711](https://github.com/tektoncd/pipeline/pull/9711) | — | skipped |
| efca098 | ci: remove compromised tj-actions/changed-files dependency | Vincent Demeester | [#9713](https://github.com/tektoncd/pipeline/pull/9713) | — | — |
| ab20605 | build(deps): bump google.golang.org/grpc from 1.79.3 to 1.80 | dependabot[bot] | [#9715](https://github.com/tektoncd/pipeline/pull/9715) | — | skipped |
| cee027b | build(deps): bump github.com/go-jose/go-jose/v4 from 4.1.3 t | dependabot[bot] | [#9721](https://github.com/tektoncd/pipeline/pull/9721) | — | skipped |
| b1781e9 | build(deps): bump github.com/go-jose/go-jose/v3 from 3.0.4 t | dependabot[bot] | [#9722](https://github.com/tektoncd/pipeline/pull/9722) | — | skipped |
| a185026 | build(deps): bump chainguard-dev/actions from 1.6.11 to 1.6. | dependabot[bot] | [#9729](https://github.com/tektoncd/pipeline/pull/9729) | — | skipped |
| 61e0304 | docs: update releases.md for v1.11.0 | Vincent Demeester | [#9685](https://github.com/tektoncd/pipeline/pull/9685) | — | — |
| 7ea863e | fix: pin registry image and relax log-based cache assertion | Vincent Demeester | [#9761](https://github.com/tektoncd/pipeline/pull/9761) | — | — |
| 721cc83 | fix: surface specific TaskRun failure reasons for pod evicti | Vibhav Bobade | [#9368](https://github.com/tektoncd/pipeline/pull/9368) | — | — |
| a468b41 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#9717](https://github.com/tektoncd/pipeline/pull/9717) | — | skipped |
| 1d50b07 | fix: surface clear errors when completed tasks miss referenc | Vibhav Bobade | [#9529](https://github.com/tektoncd/pipeline/pull/9529) | — | — |
| b1c496e | build(deps): bump go.opentelemetry.io/otel/sdk | dependabot[bot] | [#9756](https://github.com/tektoncd/pipeline/pull/9756) | — | skipped |
| 91c22de | build(deps): bump go.opentelemetry.io/otel/sdk from 1.42.0 t | dependabot[bot] | [#9757](https://github.com/tektoncd/pipeline/pull/9757) | — | skipped |
| ffd686e | Rework the events controller cache | Andrea Frittoli | [#6914](https://github.com/tektoncd/pipeline/pull/6914) | — | — |
| 11a11f7 | build(deps): bump step-security/harden-runner from 2.16.0 to | dependabot[bot] | [#9728](https://github.com/tektoncd/pipeline/pull/9728) | — | skipped |
| 659b95f | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9730](https://github.com/tektoncd/pipeline/pull/9730) | — | skipped |
| be4a4bc | docs: update metrics.md to reflect OpenTelemetry migration | Shubham Bhardwaj | [#9641](https://github.com/tektoncd/pipeline/pull/9641) | — | — |
| 5b6ec37 | [TEP-0137] Add TaskRun notifications controller | Andrea Frittoli | [#9674](https://github.com/tektoncd/pipeline/pull/9674) | — | — |
| 8546d5f | Remove aws-sdk-go-v2 ECR replace directives | Vincent Demeester | [#9773](https://github.com/tektoncd/pipeline/pull/9773) | — | — |
| ff33416 | build(deps): bump the all group in /tekton with 2 updates | dependabot[bot] | [#9770](https://github.com/tektoncd/pipeline/pull/9770) | — | skipped |
| c97b9ba | build(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#9777](https://github.com/tektoncd/pipeline/pull/9777) | — | skipped |
| 8e53012 | docs(examples): remove stale v1beta1 references from example | Sakthi Harish | [#9672](https://github.com/tektoncd/pipeline/pull/9672) | — | — |
| 9f9db26 | test(e2e): fix flaky matrix and retry example tests | Ogulcan Aydogan | [#9655](https://github.com/tektoncd/pipeline/pull/9655) | — | — |
| 770a6f7 | perf: use maps.Equal instead of reflect.DeepEqual for label/ | Ogulcan Aydogan | [#9778](https://github.com/tektoncd/pipeline/pull/9778) | — | — |
| 7e12074 | fix(e2e): replace ubuntu with busybox in dind-sidecar test D | Vibhav Bobade | [#9806](https://github.com/tektoncd/pipeline/pull/9806) | — | — |
| 529f3b2 | build(deps): bump github.com/moby/spdystream from 0.5.0 to 0 | dependabot[bot] | [#9818](https://github.com/tektoncd/pipeline/pull/9818) | — | skipped |
| a6d67e0 | build(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#9816](https://github.com/tektoncd/pipeline/pull/9816) | — | skipped |
| 83e6c68 | build(deps): bump k8s.io/api in /test/custom-task-ctrls/wait | dependabot[bot] | [#9809](https://github.com/tektoncd/pipeline/pull/9809) | — | skipped |
| ca04748 | build(deps): bump github.com/spiffe/spire-api-sdk from 1.14. | dependabot[bot] | [#9807](https://github.com/tektoncd/pipeline/pull/9807) | — | skipped |
| ca7fc9f | build(deps): bump actions/github-script from 8.0.0 to 9.0.0 | dependabot[bot] | [#9767](https://github.com/tektoncd/pipeline/pull/9767) | — | skipped |
| ac2cdb3 | build(deps): bump actions/cache from 5.0.4 to 5.0.5 | dependabot[bot] | [#9769](https://github.com/tektoncd/pipeline/pull/9769) | — | skipped |
| 4655015 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#9808](https://github.com/tektoncd/pipeline/pull/9808) | — | skipped |
| 959bc2e | build(deps): bump chainguard-dev/actions from 1.6.13 to 1.6. | dependabot[bot] | [#9766](https://github.com/tektoncd/pipeline/pull/9766) | — | skipped |
| 6a4d301 | build(deps): bump actions/upload-artifact from 7.0.0 to 7.0. | dependabot[bot] | [#9768](https://github.com/tektoncd/pipeline/pull/9768) | — | skipped |
| 5eeec51 | tekton: add draft release creation to release pipeline | Vincent Demeester | [#9420](https://github.com/tektoncd/pipeline/pull/9420) | — | — |
| 229ed41 | tekton: address review feedback on draft release tasks | Vincent Demeester | [#9420](https://github.com/tektoncd/pipeline/pull/9420) | — | — |
| 9c35f6a | build(deps): bump github.com/docker/cli in /test/resolver-wi | dependabot[bot] | [#9819](https://github.com/tektoncd/pipeline/pull/9819) | — | skipped |
| 7ea0953 | build(deps): bump github.com/tektoncd/pipeline | dependabot[bot] | [#9817](https://github.com/tektoncd/pipeline/pull/9817) | — | skipped |
| 3c9892c | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#9825](https://github.com/tektoncd/pipeline/pull/9825) | — | skipped |
| 8613eaf | build(deps): bump github.com/google/cel-go from 0.27.0 to 0. | dependabot[bot] | [#9824](https://github.com/tektoncd/pipeline/pull/9824) | — | skipped |
| 9a9894e | build(deps): bump k8s.io/api from 0.35.3 to 0.35.4 | dependabot[bot] | [#9823](https://github.com/tektoncd/pipeline/pull/9823) | — | skipped |
| 612a108 | build(deps): bump k8s.io/code-generator from 0.35.2 to 0.35. | dependabot[bot] | [#9644](https://github.com/tektoncd/pipeline/pull/9644) | — | skipped |
| 21a9083 | Regenerate code after k8s.io/code-generator 0.35.3 bump | Vincent Demeester | [#9644](https://github.com/tektoncd/pipeline/pull/9644) | — | — |
| 78229dd | build(deps): bump actions/setup-go from 6.3.0 to 6.4.0 | dependabot[bot] | [#9692](https://github.com/tektoncd/pipeline/pull/9692) | — | skipped |
| 3fd3dac | [TEP-0137] Add PipelineRun notifications controller | Andrea Frittoli | [#9677](https://github.com/tektoncd/pipeline/pull/9677) | — | — |
| 6287db3 | Simplify TestEmit by removing table-driven test structure | Andrea Frittoli | [#9677](https://github.com/tektoncd/pipeline/pull/9677) | — | — |
| e964bce | Security: reject system API token with user-controlled serve | Vincent Demeester | — | — | — |
| 0967cc8 | fix: prevent git argument injection via revision parameter ( | Vincent Demeester | — | — | — |
| c8cc73a | fix: normalize VolumeMount paths before /tekton/ restriction | Vincent Demeester | — | — | — |
| b890560 | fix: strip resolver prefixes and use non-capturing group for | Vincent Demeester | — | — | — |
| d4ba8ca | fix: limit HTTP resolver response body size to prevent OOM D | Vincent Demeester | — | — | — |
| 37f9657 | fix: trim whitespace from source URI before pattern matching | Vincent Demeester | — | — | — |
| 033ce8e | build(deps): bump k8s.io dependencies from 0.35.2 to 0.35.4 | Vincent Demeester | [#9848](https://github.com/tektoncd/pipeline/pull/9848) | — | — |
| 5133c85 | build(deps): bump github.com/hashicorp/go-version from 1.8.0 | dependabot[bot] | [#9836](https://github.com/tektoncd/pipeline/pull/9836) | — | skipped |
| ab829bf | build(deps): bump github.com/jenkins-x/go-scm from 1.15.17 t | dependabot[bot] | [#9834](https://github.com/tektoncd/pipeline/pull/9834) | — | skipped |
| 5e99282 | build(deps): bump chainguard-dev/actions from 1.6.14 to 1.6. | dependabot[bot] | [#9837](https://github.com/tektoncd/pipeline/pull/9837) | — | skipped |
| e56ea2e | build(deps): bump zizmorcore/zizmor-action from 0.5.2 to 0.5 | dependabot[bot] | [#9840](https://github.com/tektoncd/pipeline/pull/9840) | — | skipped |
| ae3c569 | build(deps): bump github/codeql-action from 4.35.1 to 4.35.2 | dependabot[bot] | [#9841](https://github.com/tektoncd/pipeline/pull/9841) | — | skipped |
| 2336dae | build(deps): bump step-security/harden-runner from 2.17.0 to | dependabot[bot] | [#9839](https://github.com/tektoncd/pipeline/pull/9839) | — | skipped |
| 2aeaecf | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9842](https://github.com/tektoncd/pipeline/pull/9842) | — | skipped |
| c4cb7a6 | build(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#9845](https://github.com/tektoncd/pipeline/pull/9845) | — | skipped |
| 1620f07 | build(deps): bump github.com/tektoncd/pipeline | dependabot[bot] | [#9851](https://github.com/tektoncd/pipeline/pull/9851) | — | skipped |
| cadf9f7 | docs: add bundle resolver configuration options default valu | Gonçalo Marques | [#9772](https://github.com/tektoncd/pipeline/pull/9772) | — | — |
| 85ff854 | ci: Automate Dependabot configuration generation | Vincent Demeester | [#9188](https://github.com/tektoncd/pipeline/pull/9188) | — | — |
| 42b30d4 | fix: address zizmor findings in dependabot-regen workflow | Vincent Demeester | [#9188](https://github.com/tektoncd/pipeline/pull/9188) | — | — |
| d98efb5 | build(deps): bump k8s.io/client-go | dependabot[bot] | [#9811](https://github.com/tektoncd/pipeline/pull/9811) | — | skipped |
| 0b4440d | [TEP-0137] Deprecate send-cloudevents-for-runs feature flag | Andrea Frittoli | [#9774](https://github.com/tektoncd/pipeline/pull/9774) | — | — |
| 6cf3274 | build(deps): bump github.com/spiffe/spire-api-sdk from 1.14. | dependabot[bot] | [#9925](https://github.com/tektoncd/pipeline/pull/9925) | — | skipped |
| 245b626 | build(deps): bump step-security/harden-runner from 2.13.2 to | dependabot[bot] | [#9861](https://github.com/tektoncd/pipeline/pull/9861) | — | skipped |
| 5a597bb | build(deps): bump go.uber.org/zap from 1.27.1 to 1.28.0 | dependabot[bot] | [#9926](https://github.com/tektoncd/pipeline/pull/9926) | — | skipped |
| 428471f | build(deps): bump actions/checkout from 6.0.0 to 6.0.2 | dependabot[bot] | [#9853](https://github.com/tektoncd/pipeline/pull/9853) | — | skipped |
| fec1273 | perf: reduce reconcile churn for completed PipelineRuns | Vincent Demeester | [#9706](https://github.com/tektoncd/pipeline/pull/9706) | — | — |
| 8709b51 | perf: remove unnecessary SetDefaults from TaskRun done path | Vincent Demeester | [#9706](https://github.com/tektoncd/pipeline/pull/9706) | — | — |
| d635dcb | test: add e2e test for TaskRun pending status | Vibhav Bobade | [#9681](https://github.com/tektoncd/pipeline/pull/9681) | — | — |
| 81f98f7 | build(deps): bump go.opentelemetry.io/otel/exporters/otlp/ot | dependabot[bot] | [#9846](https://github.com/tektoncd/pipeline/pull/9846) | — | skipped |
| 7798558 | build(deps): bump chainguard-dev/actions from 1.6.15 to 1.6. | dependabot[bot] | [#9938](https://github.com/tektoncd/pipeline/pull/9938) | — | skipped |
| 0301423 | fix(resolvers): validate data is Tekton object in resolver f | Andrew Thorp | [#9963](https://github.com/tektoncd/pipeline/pull/9963) | — | — |
| 0650977 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#9976](https://github.com/tektoncd/pipeline/pull/9976) | — | skipped |
| aed2bc5 | build(deps): bump github/codeql-action from 4.35.2 to 4.35.4 | dependabot[bot] | [#9994](https://github.com/tektoncd/pipeline/pull/9994) | — | skipped |
| e1054ca | build(deps): bump step-security/harden-runner from 2.19.0 to | dependabot[bot] | [#9981](https://github.com/tektoncd/pipeline/pull/9981) | — | skipped |
| a5afc76 | build(deps): bump chainguard-dev/actions from 1.6.17 to 1.6. | dependabot[bot] | [#9983](https://github.com/tektoncd/pipeline/pull/9983) | — | skipped |
| 88c9a98 | build(deps): bump chainguard/go in /tekton in the all group | dependabot[bot] | [#9993](https://github.com/tektoncd/pipeline/pull/9993) | — | skipped |
| 513e0e0 | Fix gen-crd-api-reference-docs require to use fetchable vers | Vincent Demeester | [#10001](https://github.com/tektoncd/pipeline/pull/10001) | — | — |
| f7054ce | build(deps): bump github.com/google/cel-go from 0.28.0 to 0. | dependabot[bot] | [#10017](https://github.com/tektoncd/pipeline/pull/10017) | — | skipped |
| 9ef82b8 | build(deps): bump github.com/jenkins-x/go-scm from 1.15.21 t | dependabot[bot] | [#10018](https://github.com/tektoncd/pipeline/pull/10018) | — | skipped |
| 8bfd0e3 | build(deps): bump k8s.io/client-go | dependabot[bot] | [#10023](https://github.com/tektoncd/pipeline/pull/10023) | — | skipped |
| 7aac85e | build(deps): bump step-security/harden-runner from 2.19.1 to | dependabot[bot] | [#10046](https://github.com/tektoncd/pipeline/pull/10046) | — | skipped |
| fbcfdd4 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#10041](https://github.com/tektoncd/pipeline/pull/10041) | — | skipped |
| d57b0dd | build(deps): bump k8s.io/apiextensions-apiserver from 0.35.4 | dependabot[bot] | [#10034](https://github.com/tektoncd/pipeline/pull/10034) | — | skipped |
| abad85f | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#10064](https://github.com/tektoncd/pipeline/pull/10064) | — | skipped |
| 3cd6cff | build(deps): bump github.com/google/go-containerregistry | dependabot[bot] | [#10070](https://github.com/tektoncd/pipeline/pull/10070) | — | skipped |
| 149d209 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#10065](https://github.com/tektoncd/pipeline/pull/10065) | — | skipped |
| 731018e | build(deps): bump step-security/harden-runner from 2.19.2 to | dependabot[bot] | [#10093](https://github.com/tektoncd/pipeline/pull/10093) | — | skipped |
| 6b544d8 | build(deps): bump zizmorcore/zizmor-action from 0.5.3 to 0.5 | dependabot[bot] | [#10092](https://github.com/tektoncd/pipeline/pull/10092) | — | skipped |
| 80647c5 | build(deps): bump the all group across 1 directory with 4 up | dependabot[bot] | [#10094](https://github.com/tektoncd/pipeline/pull/10094) | — | skipped |
| 7450f3b | build(deps): bump github/codeql-action from 4.35.4 to 4.35.5 | dependabot[bot] | [#10091](https://github.com/tektoncd/pipeline/pull/10091) | — | skipped |
| b150ab2 | fix(pipelinerun): use generateName for anonymous pipeline la | Andrew Thorp | [#10079](https://github.com/tektoncd/pipeline/pull/10079) | — | — |
| 62ff508 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#10061](https://github.com/tektoncd/pipeline/pull/10061) | — | skipped |
| c574505 | build(deps): bump github.com/sigstore/sigstore from 1.10.5 t | dependabot[bot] | [#10063](https://github.com/tektoncd/pipeline/pull/10063) | — | skipped |
| f3548d2 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#10062](https://github.com/tektoncd/pipeline/pull/10062) | — | skipped |
| e327252 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#10130](https://github.com/tektoncd/pipeline/pull/10130) | — | skipped |
| 808dd47 | build(deps): bump golangci/golangci-lint-action from 9.2.0 t | dependabot[bot] | [#10129](https://github.com/tektoncd/pipeline/pull/10129) | — | skipped |
| 9c0a33a | fix: truncate affinity assistant volume names to 63 characte | Vincent Demeester | [#10137](https://github.com/tektoncd/pipeline/pull/10137) | — | — |
| a9b0ad6 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#10149](https://github.com/tektoncd/pipeline/pull/10149) | — | skipped |
| c79a835 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#10146](https://github.com/tektoncd/pipeline/pull/10146) | — | skipped |
| ca0ab0f | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#10147](https://github.com/tektoncd/pipeline/pull/10147) | — | skipped |
| eec6412 | build(deps): bump github.com/spiffe/spire-api-sdk from 1.14. | dependabot[bot] | [#10148](https://github.com/tektoncd/pipeline/pull/10148) | — | skipped |
| aefd058 | build(deps): bump github.com/sigstore/sigstore/pkg/signature | dependabot[bot] | [#10150](https://github.com/tektoncd/pipeline/pull/10150) | — | skipped |
| 803d921 | Fix PipelineRun premature failure when TaskRun recovers afte | ab-ghosh | [#10161](https://github.com/tektoncd/pipeline/pull/10161) | — | — |
| a7e231d | fix: prevent controller CPU variant from leaking into platfo | Vincent Demeester | [#10164](https://github.com/tektoncd/pipeline/pull/10164) | — | — |
| 7121dd9 | fix: TaskRun stuck in Running when init container is OOMKill | Vincent Demeester | [#10186](https://github.com/tektoncd/pipeline/pull/10186) | — | — |
| 25e1258 | build(deps): bump actions/checkout from 6.0.2 to 6.0.3 | dependabot[bot] | [#10200](https://github.com/tektoncd/pipeline/pull/10200) | — | skipped |
| 1e0b0e7 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#10194](https://github.com/tektoncd/pipeline/pull/10194) | — | skipped |
| 13e5821 | build(deps): bump chainguard-dev/actions from 1.6.19 to 1.6. | dependabot[bot] | [#10199](https://github.com/tektoncd/pipeline/pull/10199) | — | skipped |
| aedbed0 | build(deps): bump github.com/sigstore/sigstore from 1.10.6 t | dependabot[bot] | [#10168](https://github.com/tektoncd/pipeline/pull/10168) | — | skipped |
| 64ec216 | fix: add automated draft release support to release pipeline | Vincent Demeester | [#10216](https://github.com/tektoncd/pipeline/pull/10216) | — | — |
| b8f43ba | build(deps): bump chainguard-dev/actions from 1.6.21 to 1.6. | dependabot[bot] | [#10229](https://github.com/tektoncd/pipeline/pull/10229) | — | skipped |
| 2046a4c | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#10234](https://github.com/tektoncd/pipeline/pull/10234) | — | skipped |
| 829429b | fix(resolvers): Allow ResolutionRequests to resolve all Tekt | Andrew Thorp | [#10252](https://github.com/tektoncd/pipeline/pull/10252) | — | — |
| f6ecc12 | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#10265](https://github.com/tektoncd/pipeline/pull/10265) | — | skipped |
| a1fc405 | build(deps): bump k8s.io/client-go from 0.35.5 to 0.35.6 | dependabot[bot] | [#10289](https://github.com/tektoncd/pipeline/pull/10289) | — | skipped |
| 6263f16 | build(deps): bump k8s.io/client-go | dependabot[bot] | [#10295](https://github.com/tektoncd/pipeline/pull/10295) | — | skipped |
| 6ba30f8 | build(deps): bump k8s.io/apiextensions-apiserver from 0.35.5 | dependabot[bot] | [#10290](https://github.com/tektoncd/pipeline/pull/10290) | — | skipped |
| d03cc6d | build(deps): bump the all group in /tekton with 4 updates | dependabot[bot] | [#10307](https://github.com/tektoncd/pipeline/pull/10307) | — | skipped |
| 63e2c70 | build(deps): bump the all group in /tekton with 3 updates | dependabot[bot] | [#10319](https://github.com/tektoncd/pipeline/pull/10319) | — | skipped |
| 3a8dfa2 | build(deps): bump github.com/google/go-containerregistry | dependabot[bot] | [#10333](https://github.com/tektoncd/pipeline/pull/10333) | — | skipped |
| 03e3bae | fix: bump Go for CVEs | Vibhav Bobade | [#10340](https://github.com/tektoncd/pipeline/pull/10340) | — | — |

---

## tektoncd-triggers

**Upstream:** tektoncd/triggers · **Commits:** 60 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 6ac1c13 | Change precheck in release pipeline to OCI infra | Khurram Baig | [#1947](https://github.com/tektoncd/triggers/pull/1947) | [SRVKP-9533](https://redhat.atlassian.net/browse/SRVKP-9533) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 94adc3b | Bump github/codeql-action from 4.32.2 to 4.32.3 | dependabot[bot] | [#1951](https://github.com/tektoncd/triggers/pull/1951) | — | skipped |
| 8332d99 | Bump the all group with 6 updates | dependabot[bot] | [#1950](https://github.com/tektoncd/triggers/pull/1950) | — | skipped |
| f61d6b1 | Update releases.md for v0.35.0 | Khurram Baig | [#1952](https://github.com/tektoncd/triggers/pull/1952) | — | — |
| 5501e44 | Consolidate CI workflows for build, lint, and e2e tests | Shubham Bhardwaj | [#1957](https://github.com/tektoncd/triggers/pull/1957) | — | — |
| 540cac8 | Bump step-security/harden-runner from 2.14.2 to 2.15.0 | dependabot[bot] | [#1963](https://github.com/tektoncd/triggers/pull/1963) | — | skipped |
| db0fd79 | Bump the all group across 1 directory with 5 updates | dependabot[bot] | [#1959](https://github.com/tektoncd/triggers/pull/1959) | — | skipped |
| fe55d0a | Bump github/codeql-action from 4.32.3 to 4.32.5 | dependabot[bot] | [#1962](https://github.com/tektoncd/triggers/pull/1962) | — | skipped |
| ba882b9 | Bump actions/upload-artifact from 6.0.0 to 7.0.0 | dependabot[bot] | [#1960](https://github.com/tektoncd/triggers/pull/1960) | — | skipped |
| 25aedd6 | Bump actions/setup-go from 6.2.0 to 6.3.0 | dependabot[bot] | [#1961](https://github.com/tektoncd/triggers/pull/1961) | — | skipped |
| bb8ddfc | Update Release Cheat Sheet for release-draft-oci pipeline | Khurram Baig | [#1948](https://github.com/tektoncd/triggers/pull/1948) | [SRVKP-9533](https://redhat.atlassian.net/browse/SRVKP-9533) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 88493b0 | fix: doc typos | Naomi Gelman | [#1953](https://github.com/tektoncd/triggers/pull/1953) | — | — |
| e03b74b | Bump github/codeql-action from 4.32.5 to 4.32.6 | dependabot[bot] | [#1969](https://github.com/tektoncd/triggers/pull/1969) | — | skipped |
| ef20b56 | Bump step-security/harden-runner from 2.15.0 to 2.15.1 | dependabot[bot] | [#1968](https://github.com/tektoncd/triggers/pull/1968) | — | skipped |
| 7f047e1 | Bump go version to 1.25.6 and golang ci version to 2.8.0 | Shubham Bhardwaj | [#1966](https://github.com/tektoncd/triggers/pull/1966) | — | — |
| 2bdab61 | Bump go.opentelemetry.io/otel/sdk from 1.39.0 to 1.40.0 | dependabot[bot] | [#1958](https://github.com/tektoncd/triggers/pull/1958) | — | skipped |
| d73587e | Move inactive approvers to emeritus | Vincent Demeester | [#1965](https://github.com/tektoncd/triggers/pull/1965) | — | — |
| 5f58ccb | Fix e2e failure due to owners file change | Khurram Baig | [#1970](https://github.com/tektoncd/triggers/pull/1970) | — | — |
| e8b2f37 | cleanup: replace GCS release URLs with infra.tekton.dev | Ankur Sinha | [#1973](https://github.com/tektoncd/triggers/pull/1973) | [SRVKP-11096](https://redhat.atlassian.net/browse/SRVKP-11096) | **Release Note Not Required** *(generated)* |
| 3ca9ea3 | Bump github/codeql-action from 4.32.6 to 4.33.0 | dependabot[bot] | [#1976](https://github.com/tektoncd/triggers/pull/1976) | — | skipped |
| 531dc83 | Bump step-security/harden-runner from 2.15.1 to 2.16.0 | dependabot[bot] | [#1974](https://github.com/tektoncd/triggers/pull/1974) | — | skipped |
| 60aefca | Bump fgrosse/go-coverage-report from 1.2.0 to 1.3.0 | dependabot[bot] | [#1975](https://github.com/tektoncd/triggers/pull/1975) | — | skipped |
| 07a072c | Bump actions/cache from 5.0.3 to 5.0.4 | dependabot[bot] | [#1979](https://github.com/tektoncd/triggers/pull/1979) | — | skipped |
| 4a09e38 | Bump github/codeql-action from 4.33.0 to 4.34.1 | dependabot[bot] | [#1980](https://github.com/tektoncd/triggers/pull/1980) | — | skipped |
| 9b51fd5 | Bump github/codeql-action from 4.34.1 to 4.35.1 | dependabot[bot] | [#1983](https://github.com/tektoncd/triggers/pull/1983) | — | skipped |
| b79968c | Bump actions/setup-go from 6.3.0 to 6.4.0 | dependabot[bot] | [#1984](https://github.com/tektoncd/triggers/pull/1984) | — | skipped |
| 009871a | Fix port conflict between event-listener and knative health  | Shubham Bhardwaj | [#1934](https://github.com/tektoncd/triggers/pull/1934) | [SRVKP-9324](https://redhat.atlassian.net/browse/SRVKP-9324) | **Bug Fix:** Fixes a port conflict between the event-listener and Knative health probes that  *(generated)* |
| b6eed12 | Bump Knative and Pipeline to latest | Shubham Bhardwaj | [#1934](https://github.com/tektoncd/triggers/pull/1934) | [SRVKP-9324](https://redhat.atlassian.net/browse/SRVKP-9324) | **Enhancement:** Updates Knative and Pipeline dependencies to latest versions to support the Open *(generated)* |
| 3ff1d1d | Migrate from opencensus to Opentelemetry | Shubham Bhardwaj | [#1934](https://github.com/tektoncd/triggers/pull/1934) | [SRVKP-9324](https://redhat.atlassian.net/browse/SRVKP-9324) | **Feature:** Triggers metrics implementation is migrated from the deprecated OpenCensus libra *(extracted)* |
| b293c22 | Update observability docs and config for OpenTelemetry migra | Shubham Bhardwaj | [#1934](https://github.com/tektoncd/triggers/pull/1934) | [SRVKP-9324](https://redhat.atlassian.net/browse/SRVKP-9324) | **Enhancement:** Updates observability documentation and configuration to reflect the OpenTelemet *(generated)* |
| 04257df | Bump the all group across 1 directory with 7 updates | Khurram Baig | [#1985](https://github.com/tektoncd/triggers/pull/1985) | — | — |
| f00609d | Updated generated files for newer k8s version | Khurram Baig | [#1985](https://github.com/tektoncd/triggers/pull/1985) | — | — |
| 1aad8a1 | Bump tektoncd/pipeline to v1.11.0 and update vendor dependen | Khurram Baig | [#1986](https://github.com/tektoncd/triggers/pull/1986) | — | — |
| fbf491d | Bump github.com/go-jose/go-jose/v3 from 3.0.4 to 3.0.5 | dependabot[bot] | [#1988](https://github.com/tektoncd/triggers/pull/1988) | — | skipped |
| 3886a08 | Bump github.com/go-jose/go-jose/v4 from 4.1.3 to 4.1.4 | dependabot[bot] | [#1989](https://github.com/tektoncd/triggers/pull/1989) | — | skipped |
| 08d437c | Bump the all group with 4 updates | dependabot[bot] | [#1990](https://github.com/tektoncd/triggers/pull/1990) | — | skipped |
| d83b644 | Bump step-security/harden-runner from 2.16.0 to 2.16.1 | dependabot[bot] | [#1991](https://github.com/tektoncd/triggers/pull/1991) | — | skipped |
| 3c42b6d | Bump go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlpt | dependabot[bot] | [#1992](https://github.com/tektoncd/triggers/pull/1992) | — | skipped |
| cbb4918 | Bump go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlp | dependabot[bot] | [#1993](https://github.com/tektoncd/triggers/pull/1993) | — | skipped |
| 32da9ab | Bump github.com/google/cel-go from 0.27.0 to 0.28.0 in the a | dependabot[bot] | [#1994](https://github.com/tektoncd/triggers/pull/1994) | — | skipped |
| fa0a680 | Bump step-security/harden-runner from 2.16.1 to 2.17.0 | dependabot[bot] | [#1995](https://github.com/tektoncd/triggers/pull/1995) | — | skipped |
| d307d94 | Bump actions/upload-artifact from 7.0.0 to 7.0.1 | dependabot[bot] | [#1996](https://github.com/tektoncd/triggers/pull/1996) | — | skipped |
| 0c827bc | Update release ko image to go1.25 | Khurram Baig | [#1997](https://github.com/tektoncd/triggers/pull/1997) | [SRVKP-2980](https://redhat.atlassian.net/browse/SRVKP-2980) | **Release Notes Not Required:** NA *(generated)* |
| 7e8b093 | Bump github.com/moby/spdystream from 0.5.0 to 0.5.1 | dependabot[bot] | [#1999](https://github.com/tektoncd/triggers/pull/1999) | — | skipped |
| 117a7d9 | ci: Fix zizmor security findings in GitHub Actions | Shubham Bhardwaj | [#1998](https://github.com/tektoncd/triggers/pull/1998) | — | — |
| 9f011db | Fix intermittent panic in Test_UpdateCACertToClusterIntercep | Khurram Baig | [#2000](https://github.com/tektoncd/triggers/pull/2000) | — | — |
| 08d1e6a | Bump the all group across 1 directory with 5 updates | dependabot[bot] | [#2001](https://github.com/tektoncd/triggers/pull/2001) | — | skipped |
| 1d2b212 | Bump step-security/harden-runner from 2.17.0 to 2.19.0 | dependabot[bot] | [#2003](https://github.com/tektoncd/triggers/pull/2003) | — | skipped |
| ba6b012 | Bump actions/cache from 5.0.4 to 5.0.5 | dependabot[bot] | [#2004](https://github.com/tektoncd/triggers/pull/2004) | — | skipped |
| 1352ff6 | Bump github/codeql-action from 4.35.1 to 4.35.2 | dependabot[bot] | [#2002](https://github.com/tektoncd/triggers/pull/2002) | — | skipped |
| d86fccf | Bump zizmorcore/zizmor-action from 0.5.2 to 0.5.3 | dependabot[bot] | [#2007](https://github.com/tektoncd/triggers/pull/2007) | — | skipped |
| f475444 | Bump github.com/tektoncd/pipeline from 1.11.0 to 1.11.1 | dependabot[bot] | [#2005](https://github.com/tektoncd/triggers/pull/2005) | — | skipped |
| f61ad46 | Bump go.uber.org/zap from 1.27.1 to 1.28.0 in the all group | dependabot[bot] | [#2016](https://github.com/tektoncd/triggers/pull/2016) | — | skipped |
| 6e57af9 | Fix curl command to follow redirects for release file | Khurram | [#2017](https://github.com/tektoncd/triggers/pull/2017) | [SRVKP-9533](https://redhat.atlassian.net/browse/SRVKP-9533) *(keyword)* | **Release Notes Not Required** *(generated)* |
| bf0a492 | Change release pipeline to use 'release-draft-oci' | Khurram | [#2018](https://github.com/tektoncd/triggers/pull/2018) | [SRVKP-9533](https://redhat.atlassian.net/browse/SRVKP-9533) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 1d36b80 | Bump the pipelines dependency and other deps | Khurram Baig | [#2020](https://github.com/tektoncd/triggers/pull/2020) | — | — |
| 01727f1 | Bump the all group across 1 directory with 7 updates | dependabot[bot] | [#2031](https://github.com/tektoncd/triggers/pull/2031) | — | skipped |
| 44efd91 | Bump github/codeql-action from 4.35.2 to 4.35.4 | dependabot[bot] | [#2032](https://github.com/tektoncd/triggers/pull/2032) | — | skipped |
| 8ed8f72 | Bump step-security/harden-runner from 2.19.0 to 2.19.3 | dependabot[bot] | [#2033](https://github.com/tektoncd/triggers/pull/2033) | — | skipped |
| bf95a60 | Add TLS security profile support for core interceptors | Jawed khelil | [#2019](https://github.com/tektoncd/triggers/pull/2019) | [SRVKP-12012](https://redhat.atlassian.net/browse/SRVKP-12012) | **Feature:** Core interceptors now honor the TLS security profile injected by the Tekton oper *(extracted)* |

---

## syncer-service

**Upstream:** openshift-pipelines/syncer-service · **Commits:** 1 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 30b70d1 | update pipeline version to v1.9.3 | Pramod Bindal | [#71](https://github.com/openshift-pipelines/syncer-service/pull/71) | — | — |

---

## tekton-assist

**Error:** failed to get head SHA at from-date: failed to list head file commits for openshift-pipelines/p12n-tekton-assist: GET https://api.github.com/repos/openshift-pipelines/p12n-tekton-assist/commits?path=head&per_page=1&sha=release-v1.22.x&until=2026-05-15T07%3A57%3A06Z: 404 Not Found []

---

## tekton-caches

**Upstream:** openshift-pipelines/tekton-caches · **Commits:** 4 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 2450f24 | Security: Fix CVE-2026-34986 (go-jose/go-jose/v4) | Pramod Bindal | [#729](https://github.com/openshift-pipelines/tekton-caches/pull/729) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Fixes CVE-2026-34986 by upgrading go-jose/go-jose/v4 to v4.1.4, addressing a vul *(generated)* |
| 2a64502 | bump otel and pipeline version for CVE fixes | Pramod Bindal | [#736](https://github.com/openshift-pipelines/tekton-caches/pull/736) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Updates OpenTelemetry and Tekton Pipeline dependency versions to incorporate ups *(generated)* |
| 5d6871a | cve: Fix GHSA-xmrv-pmrh-hhx2 | Pramod Bindal | [#739](https://github.com/openshift-pipelines/tekton-caches/pull/739) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Fixes GHSA-xmrv-pmrh-hhx2 by upgrading github.com/aws/aws-sdk-go-v2 from v1.41.2 *(generated)* |
| 768eb9a | Add Agentic Contexts and CI Workflows | Pramod Bindal | [#739](https://github.com/openshift-pipelines/tekton-caches/pull/739) | [SRVKP-12024](https://redhat.atlassian.net/browse/SRVKP-12024) *(keyword)* | **Release Notes Not Required** *(generated)* |

---

## Cross-Component Summary

| Component | Commits | Linked (tool) | Linked (audit) | Unlinked | Has RN | Generated RN | Dependabot |
|-----------|---------|---------------|----------------|----------|--------|--------------|------------|
| git-init | — | — | — | — | — | — | — |
| opc | 3 | 0 | 3 | 0 | 3 | 3 | 0 |
| console-plugin-pf5 | 19 | 3 | 15 | 1 | 18 | 16 | 0 |
| tekton-kueue | 2 | 0 | 1 | 1 | 0 | 0 | 0 |
| pipelines-as-code | 230 | 57 | 39 | 119 | 95 | 95 | 15 |
| tektoncd-hub | 5 | 0 | 1 | 4 | 5 | 5 | 0 |
| tektoncd-pruner | 162 | 15 | 9 | 20 | 24 | 2 | 118 |
| tektoncd-results | 116 | 9 | 1 | 18 | 10 | 4 | 88 |
| operator | 161 | 26 | 24 | 40 | 49 | 5 | 71 |
| manual-approval-gate | 2 | 0 | 1 | 1 | 2 | 1 | 0 |
| tektoncd-chains | 145 | 6 | 2 | 40 | 8 | 5 | 97 |
| tektoncd-cli | 71 | 4 | 3 | 12 | 7 | 4 | 52 |
| tektoncd-pipeline | 327 | 18 | 10 | 126 | 25 | 13 | 173 |
| tektoncd-triggers | 60 | 7 | 4 | 12 | 11 | 9 | 37 |
| syncer-service | 1 | 0 | 0 | 1 | 0 | 0 | 0 |
| multicluster-proxy-aae | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| tekton-assist | — | — | — | — | — | — | — |
| tekton-caches | 4 | 0 | 4 | 0 | 4 | 4 | 0 |
| **Total** | **1308** | **145** | **117** | **395** | **261** | **166** | **651** |

## Unmatched Jiras

These 172 Jira tickets have `fixVersion = 1.23.0` but could not be matched to any commit in the audited components.

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-12577](https://redhat.atlassian.net/browse/SRVKP-12577) | [OSP 1.23.0] TektonHub CR status.version field is empty causing version verification test to fail | New |
| [SRVKP-12576](https://redhat.atlassian.net/browse/SRVKP-12576) | [OSP 1.23.0] tkn and tkn-pac version mismatch in ecosystem tests - env config expects 0.45/0.48 but  | To Do |
| [SRVKP-12538](https://redhat.atlassian.net/browse/SRVKP-12538) | QA- [Regression] Relative task references in Pipeline broken | Closed |
| [SRVKP-12484](https://redhat.atlassian.net/browse/SRVKP-12484) | Manual Approval Gate Component Missing from Deployment in OpenShift Pipelines 1.23.0 | On QA |
| [SRVKP-12426](https://redhat.atlassian.net/browse/SRVKP-12426) | CLONE -  Upstream release notes for hub | New |
| [SRVKP-12425](https://redhat.atlassian.net/browse/SRVKP-12425) | CLONE -  Upstream release notes for Operator | New |
| [SRVKP-12422](https://redhat.atlassian.net/browse/SRVKP-12422) | CLONE -  Upstream release notes for Pipelines | New |
| [SRVKP-12421](https://redhat.atlassian.net/browse/SRVKP-12421) | CLONE -  Upstream release notes for PAC | New |
| [SRVKP-12418](https://redhat.atlassian.net/browse/SRVKP-12418) | Upstream release notes for 1.23 | New |
| [SRVKP-12416](https://redhat.atlassian.net/browse/SRVKP-12416) | Release Captain Work 1.23 | Dev Complete |
| [SRVKP-12185](https://redhat.atlassian.net/browse/SRVKP-12185) | In all triggers & sub pages, when searching by label, auto suggestions are not populated | Closed |
| [SRVKP-12184](https://redhat.atlassian.net/browse/SRVKP-12184) | In tasks/taskruns page, when searching by label, auto suggestions are not populated | Closed |
| [SRVKP-11932](https://redhat.atlassian.net/browse/SRVKP-11932) | Delete Repository CR status field and do clean up in the PaC repo | Tasking and Estimation |
| [SRVKP-11931](https://redhat.atlassian.net/browse/SRVKP-11931) | Deprecate Repository CR PipelineRun status field | Backlog |
| [SRVKP-11929](https://redhat.atlassian.net/browse/SRVKP-11929) | [Reopen] Default TekonConfig is created when it should not | Tasking and Estimation |
| [SRVKP-11923](https://redhat.atlassian.net/browse/SRVKP-11923) | Release Captain work 1.23 | Closed |
| [SRVKP-11888](https://redhat.atlassian.net/browse/SRVKP-11888) | [release testing] Documentation review | Tasking and Estimation |
| [SRVKP-11887](https://redhat.atlassian.net/browse/SRVKP-11887) | [release testing] Post-release verify operator installation | Tasking and Estimation |
| [SRVKP-11886](https://redhat.atlassian.net/browse/SRVKP-11886) | [release testing] TPS tests in pre-stage and stage | Tasking and Estimation |
| [SRVKP-11885](https://redhat.atlassian.net/browse/SRVKP-11885) | [release testing] Testing in disconnected env Z  | Tasking and Estimation |
| [SRVKP-11884](https://redhat.atlassian.net/browse/SRVKP-11884) | [release testing] TKN/OPC tests | Tasking and Estimation |
| [SRVKP-11883](https://redhat.atlassian.net/browse/SRVKP-11883) | [release testing] Multiarch testing - IBM Z | Tasking and Estimation |
| [SRVKP-11882](https://redhat.atlassian.net/browse/SRVKP-11882) | [release testing] Verification of CLI binary SHA/signature on Mac and Windows | Tasking and Estimation |
| [SRVKP-11881](https://redhat.atlassian.net/browse/SRVKP-11881) | [release testing] Cluster upgrade tests | Tasking and Estimation |
| [SRVKP-11880](https://redhat.atlassian.net/browse/SRVKP-11880) | [release testing] Bug verification | Tasking and Estimation |
| [SRVKP-11879](https://redhat.atlassian.net/browse/SRVKP-11879) | [release testing] Devconsole UI manual tests | Tasking and Estimation |
| [SRVKP-11878](https://redhat.atlassian.net/browse/SRVKP-11878) | [release testing] Hub manual tests | Tasking and Estimation |
| [SRVKP-11877](https://redhat.atlassian.net/browse/SRVKP-11877) | [release testing] Verify binaries on mirror.openshift.com | Tasking and Estimation |
| [SRVKP-11876](https://redhat.atlassian.net/browse/SRVKP-11876) | [release testing] create release branch in release-tests for 1.24 | Tasking and Estimation |
| [SRVKP-11875](https://redhat.atlassian.net/browse/SRVKP-11875) | [release testing] HyperShift/ROSA testing | Tasking and Estimation |
| [SRVKP-11874](https://redhat.atlassian.net/browse/SRVKP-11874) | [release testing] Testing in disconnected env x86_64 | Tasking and Estimation |
| [SRVKP-11873](https://redhat.atlassian.net/browse/SRVKP-11873) | [release testing] Multiarch testing - IBM Power | Tasking and Estimation |
| [SRVKP-11872](https://redhat.atlassian.net/browse/SRVKP-11872) | [release testing] update ci-config.yaml | Tasking and Estimation |
| [SRVKP-11871](https://redhat.atlassian.net/browse/SRVKP-11871) | [release testing] Operator upgrade tests | Tasking and Estimation |
| [SRVKP-11870](https://redhat.atlassian.net/browse/SRVKP-11870) | [release testing] Feature testing | Tasking and Estimation |
| [SRVKP-11869](https://redhat.atlassian.net/browse/SRVKP-11869) | [release testing] Test on FIPS cluster | Tasking and Estimation |
| [SRVKP-11868](https://redhat.atlassian.net/browse/SRVKP-11868) | [release testing] TKN entitlement tests | Tasking and Estimation |
| [SRVKP-11867](https://redhat.atlassian.net/browse/SRVKP-11867) | [release testing] Multiarch testing - ARM64 | Tasking and Estimation |
| [SRVKP-11866](https://redhat.atlassian.net/browse/SRVKP-11866) | [release testing] Testing in disconnected env P | Tasking and Estimation |
| [SRVKP-11840](https://redhat.atlassian.net/browse/SRVKP-11840) | Support OpenShift Compatibility via Dynamic Console Plugin Deployment | Backlog |
| [SRVKP-11827](https://redhat.atlassian.net/browse/SRVKP-11827) | Upstream Kueue and Tekton Multiclustering Contributions | Tasking and Estimation |
| [SRVKP-11826](https://redhat.atlassian.net/browse/SRVKP-11826) | ACM Pipelines Add-on for Multicluster Automation | Refinement |
| [SRVKP-11715](https://redhat.atlassian.net/browse/SRVKP-11715) | Translation PR for all migrations in master branch | Release Pending |
| [SRVKP-11711](https://redhat.atlassian.net/browse/SRVKP-11711) | Upgrade to console sdk 4.22.0-prerelease.3 | Release Pending |
| [SRVKP-11708](https://redhat.atlassian.net/browse/SRVKP-11708) | CSS Collision: pipelines-plugin overrides .ocs-quick-search-modal styles in Console Topology | Closed |
| [SRVKP-11640](https://redhat.atlassian.net/browse/SRVKP-11640) | OpenShift 5.0 Operator Compatibility Requirements for Pipelines | Refinement |
| [SRVKP-11566](https://redhat.atlassian.net/browse/SRVKP-11566) | Replace <LoadingBox/> with Loading Component to remove PF 5 elements | Closed |
| [SRVKP-11565](https://redhat.atlassian.net/browse/SRVKP-11565) | Pipeline topology not working in Pipelinerun details page for Multicluster re-run | Closed |
| [SRVKP-11564](https://redhat.atlassian.net/browse/SRVKP-11564) | LogViewer not showing up for a failed task in multicluster | Closed |
| [SRVKP-11549](https://redhat.atlassian.net/browse/SRVKP-11549) | AI analysis comment duplicated on each PR event instead of updating existing comment | In Progress |
| [SRVKP-11512](https://redhat.atlassian.net/browse/SRVKP-11512) | Add comment ownership validation to prevent editing other users' comments | Dev Complete |
| [SRVKP-11484](https://redhat.atlassian.net/browse/SRVKP-11484) | Embed Rekor bundle in OCI attestation layer annotations when transparency is enabled | Code Review |
| [SRVKP-11464](https://redhat.atlassian.net/browse/SRVKP-11464) | Support 'max_completion_tokens' for newer OpenAI-compatible models in PAC LLM | Closed |
| [SRVKP-11451](https://redhat.atlassian.net/browse/SRVKP-11451) | Q2 OSP Roadmap workshop 2026 Presentation - Child 1 | Closed |
| [SRVKP-11445](https://redhat.atlassian.net/browse/SRVKP-11445) | Openshift Pipelines Release 1.23.0 | Backlog |
| [SRVKP-11433](https://redhat.atlassian.net/browse/SRVKP-11433) | [Multi Cluster] Displaying blank status for Pending pipelineRuns across CLI & UI | Tasking and Estimation |
| [SRVKP-11432](https://redhat.atlassian.net/browse/SRVKP-11432) | PVC auto-cleanup annotation propagates to child PipelineRuns in Pipeline-in-Pipeline | Code Review |
| [SRVKP-11401](https://redhat.atlassian.net/browse/SRVKP-11401) | Update all VirtualisedTable list to consoledataview for approval tabs | Release Pending |
| [SRVKP-11400](https://redhat.atlassian.net/browse/SRVKP-11400) | Embedded Color Code Pipeline Run Details Page Dark Mode | Release Pending |
| [SRVKP-11399](https://redhat.atlassian.net/browse/SRVKP-11399) | ConsoleDataView migration for PipelineRun List in Pipeline Overview Page | Release Pending |
| [SRVKP-11398](https://redhat.atlassian.net/browse/SRVKP-11398) | Upgrade to console sdk 4.22.0-prerelease.2  | Release Pending |
| [SRVKP-11397](https://redhat.atlassian.net/browse/SRVKP-11397) | Multi cluster UI testing in master branch - 2 | Release Pending |
| [SRVKP-11396](https://redhat.atlassian.net/browse/SRVKP-11396) | Multi cluster UI testing in master branch - 1  | Release Pending |
| [SRVKP-11395](https://redhat.atlassian.net/browse/SRVKP-11395) | Refactor useTaskRuns.ts file to accommodate multi cluster changes in master branch | Closed |
| [SRVKP-11394](https://redhat.atlassian.net/browse/SRVKP-11394) | Provider actions needs to be updated for Re-run of PipelineRun | Release Pending |
| [SRVKP-11393](https://redhat.atlassian.net/browse/SRVKP-11393) | Logs tab showing all taskruns in the side nav instead of taskruns from the PLR | Release Pending |
| [SRVKP-11392](https://redhat.atlassian.net/browse/SRVKP-11392) | Log viewer not getting displayed post manual merge from main to master | Release Pending |
| [SRVKP-11391](https://redhat.atlassian.net/browse/SRVKP-11391) | Archived pipelinerun details page not showing up post manual merge from main to master | Release Pending |
| [SRVKP-11390](https://redhat.atlassian.net/browse/SRVKP-11390) | Archived taskrun details page not showing up post manual merge from main to master | Release Pending |
| [SRVKP-11383](https://redhat.atlassian.net/browse/SRVKP-11383) | Align the refresh loader on the Pipelines-Overview page ( Prometheus ) | Closed |
| [SRVKP-11379](https://redhat.atlassian.net/browse/SRVKP-11379) | Align the refresh loader on the Pipelines-Overview page ( Tekton results ) | Release Pending |
| [SRVKP-11199](https://redhat.atlassian.net/browse/SRVKP-11199) | MultiCluster \| Unable to Cancel/Stop a pipeline run  | Closed |
| [SRVKP-11167](https://redhat.atlassian.net/browse/SRVKP-11167) | Create a custom filter component as the console sdk has not added @patternfly/react-data-view to sha | Release Pending |
| [SRVKP-11136](https://redhat.atlassian.net/browse/SRVKP-11136) | Document Guidance on Scaling Tekton Controllers for High-Volume Workloads | Tasking and Estimation |
| [SRVKP-11123](https://redhat.atlassian.net/browse/SRVKP-11123) | CI Failure in React 18 Branch Build and updating dependencies in package.json | Release Pending |
| [SRVKP-11115](https://redhat.atlassian.net/browse/SRVKP-11115) | [QE - 1.23 ] Validate the 1.22 OSP image as 1.23 OSP for OCP 4.21 and below | Closed |
| [SRVKP-11101](https://redhat.atlassian.net/browse/SRVKP-11101) | [Testing] CLONE - Validate Repository CR URL path shape in admission webhook to prevent token scopin | Closed |
| [SRVKP-11098](https://redhat.atlassian.net/browse/SRVKP-11098) | R&D: Schedule PipelineRuns on Hub and Spoke clusters based on availability using MultiKueue | Closed |
| [SRVKP-11051](https://redhat.atlassian.net/browse/SRVKP-11051) | Adopt Pipelines-as-Code into the Tekton Project | In Progress |
| [SRVKP-11023](https://redhat.atlassian.net/browse/SRVKP-11023) | Tasks for deprecating Tekton Hub | Backlog |
| [SRVKP-11003](https://redhat.atlassian.net/browse/SRVKP-11003) | console-plugin branch readiness for react 18 | Release Pending |
| [SRVKP-11001](https://redhat.atlassian.net/browse/SRVKP-11001) | Yarn4 CI operator ( release repo ) Issue about name conflict for main_ocp_4.22 branch in the release | Closed |
| [SRVKP-10969](https://redhat.atlassian.net/browse/SRVKP-10969) | Adopting Pipelines-as-Code into Tekton Project | Closed |
| [SRVKP-10967](https://redhat.atlassian.net/browse/SRVKP-10967) | [Testing] Add comment ownership validation to prevent editing other users' comments | To Do |
| [SRVKP-10964](https://redhat.atlassian.net/browse/SRVKP-10964) | Migrate CI Credentials to HashiCorp Vault Cloud | Closed |
| [SRVKP-10963](https://redhat.atlassian.net/browse/SRVKP-10963) | Migrate CI Credentials to HashiCorp Vault Cloud | Closed |
| [SRVKP-10962](https://redhat.atlassian.net/browse/SRVKP-10962) | Migrate CI Credentials to HashiCorp Vault Cloud | Closed |
| [SRVKP-10953](https://redhat.atlassian.net/browse/SRVKP-10953) | Fix the issue of failed CI job (frontend) for main_ocp_4.22 | Release Pending |
| [SRVKP-10928](https://redhat.atlassian.net/browse/SRVKP-10928) | Add self-healing for CA bundle configmaps in user namespaces | Closed |
| [SRVKP-10910](https://redhat.atlassian.net/browse/SRVKP-10910) | Automate GitHub App Token Extension for Private Forks | Tasking and Estimation |
| [SRVKP-10907](https://redhat.atlassian.net/browse/SRVKP-10907) | Test Openshift Pipelines on RHCOS 9 and 10 for OCP 4.22 | Closed |
| [SRVKP-10806](https://redhat.atlassian.net/browse/SRVKP-10806) | Add JAVA_VERSION parameter to Maven Task for Java version selection | Closed |
| [SRVKP-10736](https://redhat.atlassian.net/browse/SRVKP-10736) | Reproduce this number of PipelineRuns count mismatch using Kind | Dev Complete |
| [SRVKP-10630](https://redhat.atlassian.net/browse/SRVKP-10630) | Test Automation - Add Release Tests | In Progress |
| [SRVKP-10574](https://redhat.atlassian.net/browse/SRVKP-10574) | [Pre-Stage] [release testing] Testing in disconnected env P | To Do |
| [SRVKP-10573](https://redhat.atlassian.net/browse/SRVKP-10573) | [Pre-Stage] [release testing] Cluster upgrade tests | To Do |
| [SRVKP-10572](https://redhat.atlassian.net/browse/SRVKP-10572) | [Pre-Stage] [release testing] update ci-config.yaml | In Progress |
| [SRVKP-10571](https://redhat.atlassian.net/browse/SRVKP-10571) | [Pre-Stage] [release testing] Documentation review of Release Notes  | To Do |
| [SRVKP-10570](https://redhat.atlassian.net/browse/SRVKP-10570) | [Post-Release] [release testing] Multiarch testing - ARM64 | To Do |
| [SRVKP-10569](https://redhat.atlassian.net/browse/SRVKP-10569) | [Stage] [release testing] Operator Upgrade  | In Progress |
| [SRVKP-10568](https://redhat.atlassian.net/browse/SRVKP-10568) | [Stage] [release testing] Acceptance Tests  | In Progress |
| [SRVKP-10567](https://redhat.atlassian.net/browse/SRVKP-10567) | [Pre-Stage] [release testing] TKN/OPC tests | To Do |
| [SRVKP-10566](https://redhat.atlassian.net/browse/SRVKP-10566) | [Post-Release] [release testing] Post-release verify operator installation | To Do |
| [SRVKP-10565](https://redhat.atlassian.net/browse/SRVKP-10565) | [Pre-Stage] [release testing] Testing in disconnected env x86_64 | To Do |
| [SRVKP-10564](https://redhat.atlassian.net/browse/SRVKP-10564) | [Pre-Stage] [release testing] create release branch in release-tests | To Do |
| [SRVKP-10563](https://redhat.atlassian.net/browse/SRVKP-10563) | [Pre-Stage] [release testing] Documentation review For Feature Testing | To Do |
| [SRVKP-10562](https://redhat.atlassian.net/browse/SRVKP-10562) | [Pre-Stage] [release testing] Devconsole UI manual tests | To Do |
| [SRVKP-10561](https://redhat.atlassian.net/browse/SRVKP-10561) | [Stage] [release testing] Verification of CLI binary SHA/signature on Mac and Windows | To Do |
| [SRVKP-10560](https://redhat.atlassian.net/browse/SRVKP-10560) | [Pre-Stage] [release testing] Feature testing | In Progress |
| [SRVKP-10559](https://redhat.atlassian.net/browse/SRVKP-10559) | [Pre-Stage] [release testing] Testing in disconnected env Z  | To Do |
| [SRVKP-10558](https://redhat.atlassian.net/browse/SRVKP-10558) | [Post-Release] [release testing] Multiarch testing - IBM Z | To Do |
| [SRVKP-10557](https://redhat.atlassian.net/browse/SRVKP-10557) | [Pre-Stage] [release testing] Operator upgrade tests | To Do |
| [SRVKP-10556](https://redhat.atlassian.net/browse/SRVKP-10556) | [Pre-Stage] [release testing] Test on FIPS cluster | To Do |
| [SRVKP-10555](https://redhat.atlassian.net/browse/SRVKP-10555) | [Post release] [release testing] Verify binaries on mirror.openshift.com | To Do |
| [SRVKP-10554](https://redhat.atlassian.net/browse/SRVKP-10554) | [Pre-Stage] [release testing] Bug verification | In Progress |
| [SRVKP-10553](https://redhat.atlassian.net/browse/SRVKP-10553) | [Post-Release] [release testing] Multiarch testing - IBM Power | To Do |
| [SRVKP-10552](https://redhat.atlassian.net/browse/SRVKP-10552) | [Pre-Stage] [release testing] Hub manual tests | To Do |
| [SRVKP-10551](https://redhat.atlassian.net/browse/SRVKP-10551) | Test Openshift Pipelines 1.23.0 | In Progress |
| [SRVKP-10469](https://redhat.atlassian.net/browse/SRVKP-10469) | [operator] Fetch TLS config and expose to components  | Closed |
| [SRVKP-10457](https://redhat.atlassian.net/browse/SRVKP-10457) | PipelineRun intermittently fails and skips downstream tasks while a TaskRun is still Running under h | Closed |
| [SRVKP-10413](https://redhat.atlassian.net/browse/SRVKP-10413) | Migrate from using gitea sdk to forgejo sdk | Closed |
| [SRVKP-10307](https://redhat.atlassian.net/browse/SRVKP-10307) | Upgrade react-redux | Release Pending |
| [SRVKP-10306](https://redhat.atlassian.net/browse/SRVKP-10306) | Upgrade react router to react router 7 | Release Pending |
| [SRVKP-10303](https://redhat.atlassian.net/browse/SRVKP-10303) | Migrate CI Credentials to HashiCorp Vault Cloud | Closed |
| [SRVKP-10296](https://redhat.atlassian.net/browse/SRVKP-10296) | Releasing RPMs via Konflux | Closed |
| [SRVKP-10269](https://redhat.atlassian.net/browse/SRVKP-10269) | Extend the list of managed resources | Refinement |
| [SRVKP-10266](https://redhat.atlassian.net/browse/SRVKP-10266) | Pipelines as Code integration with Forgejo | Dev Complete |
| [SRVKP-10262](https://redhat.atlassian.net/browse/SRVKP-10262) | Define project/release resource to tie up related PipelineRuns up | Refinement |
| [SRVKP-10254](https://redhat.atlassian.net/browse/SRVKP-10254) | Fetch TLS Settings from a centrallised location | In Progress |
| [SRVKP-10246](https://redhat.atlassian.net/browse/SRVKP-10246) | OCP Compatibility – Refactor Modals & Tables | Closed |
| [SRVKP-10245](https://redhat.atlassian.net/browse/SRVKP-10245) | OCP 4.22 Readiness – PatternFly 6 & React@18 Upgrade | Backlog |
| [SRVKP-10080](https://redhat.atlassian.net/browse/SRVKP-10080) | Fix User and Group tags CSS  in Approval task details page | Release Pending |
| [SRVKP-10077](https://redhat.atlassian.net/browse/SRVKP-10077) | Streaming event icon not appearing | Release Pending |
| [SRVKP-10076](https://redhat.atlassian.net/browse/SRVKP-10076) | Spacing between Pipelines and samples and snippet  | Release Pending |
| [SRVKP-10075](https://redhat.atlassian.net/browse/SRVKP-10075) | Alignment in Overview Page for all cards ( tekton result + prometheus ) | Release Pending |
| [SRVKP-10072](https://redhat.atlassian.net/browse/SRVKP-10072) | Deprecate modal.ts  | Release Pending |
| [SRVKP-10069](https://redhat.atlassian.net/browse/SRVKP-10069) | Fix Task ResourceIcon color to be consistent with other CRDs and hierarchy in TT and CTB | Release Pending |
| [SRVKP-10066](https://redhat.atlassian.net/browse/SRVKP-10066) | Add namespace label to `tekton_pipelines_controller_pipelinerun_total` metric | In Progress |
| [SRVKP-10038](https://redhat.atlassian.net/browse/SRVKP-10038) | Logs should contain pipeline or task names  | Dev Complete |
| [SRVKP-10037](https://redhat.atlassian.net/browse/SRVKP-10037) | Pipeline builder topology not showing correct status when PLR is Rerun from the Action dropdown | Release Pending |
| [SRVKP-9983](https://redhat.atlassian.net/browse/SRVKP-9983) | Fix log viewer UI in PF6 | Release Pending |
| [SRVKP-9918](https://redhat.atlassian.net/browse/SRVKP-9918) | Deprecate CreateModalLauncher and use useOverlay hook for pipeline builder modal | Release Pending |
| [SRVKP-9789](https://redhat.atlassian.net/browse/SRVKP-9789) | List style on the Topology side panel for the Pipelines resource is not consistent with PF5 | Release Pending |
| [SRVKP-9788](https://redhat.atlassian.net/browse/SRVKP-9788) | Input field styling on start pipelines and dropdown style on the Start pipeline modal is not consist | Release Pending |
| [SRVKP-9787](https://redhat.atlassian.net/browse/SRVKP-9787) | PipelineRun details page sub header is broken, and the log viewer style is not in sync with PF5 | Release Pending |
| [SRVKP-9786](https://redhat.atlassian.net/browse/SRVKP-9786) | Fix the Dropdown Actions | Release Pending |
| [SRVKP-9785](https://redhat.atlassian.net/browse/SRVKP-9785) | Fix Pipeline Builder Modal Styles | Release Pending |
| [SRVKP-9784](https://redhat.atlassian.net/browse/SRVKP-9784) | Fix Border Styles in PF6  | Release Pending |
| [SRVKP-9783](https://redhat.atlassian.net/browse/SRVKP-9783) | Fix - Ensure Consistent Font Size | Release Pending |
| [SRVKP-9722](https://redhat.atlassian.net/browse/SRVKP-9722) | Performance regression testing for Pipelines 1.23.0 | Tasking and Estimation |
| [SRVKP-9721](https://redhat.atlassian.net/browse/SRVKP-9721) | Release OpenShift Pipelines CLI for Pipelines 1.23.0 | Tasking and Estimation |
| [SRVKP-9720](https://redhat.atlassian.net/browse/SRVKP-9720) | Release OpenShift Pipelines Operator for Pipelines 1.23.0 | Tasking and Estimation |
| [SRVKP-9713](https://redhat.atlassian.net/browse/SRVKP-9713) | Migrate remaining plugin list/detail pages to ConsoleDataView and LazyActionMenu | Closed |
| [SRVKP-9712](https://redhat.atlassian.net/browse/SRVKP-9712) | Update all VirtualisedTable list to data-view for EventListener and Trigger Template tabs | Release Pending |
| [SRVKP-9711](https://redhat.atlassian.net/browse/SRVKP-9711) | Deprecate react-helmet and use DocumentTitle instead | Release Pending |
| [SRVKP-9710](https://redhat.atlassian.net/browse/SRVKP-9710) | Migrate Task and TaskRun List Views to DataView (React 18) | Release Pending |
| [SRVKP-9709](https://redhat.atlassian.net/browse/SRVKP-9709) | Update VirtualisedTable list to ConsoleDataView for repositories tab | Release Pending |
| [SRVKP-9663](https://redhat.atlassian.net/browse/SRVKP-9663) | Forgejo support for Pipelines as Code (PaC) | Closed |
| [SRVKP-9627](https://redhat.atlassian.net/browse/SRVKP-9627) | Replace useModal with use useOverlay hook and add native PF6 Modal components for Start pipeline, Ad | Release Pending |
| [SRVKP-9626](https://redhat.atlassian.net/browse/SRVKP-9626) | Update all VirtualisedTable list to data-view for pipeline and pipeline run tabs | Release Pending |
| [SRVKP-9625](https://redhat.atlassian.net/browse/SRVKP-9625) | Update page and section headings for Pipelines, Tasks and Triggers Page | Release Pending |
| [SRVKP-9447](https://redhat.atlassian.net/browse/SRVKP-9447) | Migration Guide and Documentation | Tasking and Estimation |
| [SRVKP-9392](https://redhat.atlassian.net/browse/SRVKP-9392) | Review and update variable and class names. | Release Pending |
| [SRVKP-9391](https://redhat.atlassian.net/browse/SRVKP-9391) | Remove all CSS overrides. | Release Pending |
| [SRVKP-9390](https://redhat.atlassian.net/browse/SRVKP-9390) | Update dependencies to use PatternFly 6 instead of PatternFly 5 | Release Pending |
| [SRVKP-9389](https://redhat.atlassian.net/browse/SRVKP-9389) | Run PF6 codemods suite | Release Pending |
| [SRVKP-9303](https://redhat.atlassian.net/browse/SRVKP-9303) | Discovery on the upgrade to PF6 | Release Pending |
| [SRVKP-9114](https://redhat.atlassian.net/browse/SRVKP-9114) | Testing for the epic | Closed |
| [SRVKP-8541](https://redhat.atlassian.net/browse/SRVKP-8541) | Enable SSA finalizer management Results | Code Review |
| [SRVKP-8292](https://redhat.atlassian.net/browse/SRVKP-8292) | Cancelled/failed pipelineRun did not trigger finally tasks | Release Pending |
| [SRVKP-7899](https://redhat.atlassian.net/browse/SRVKP-7899) | Migrate Tekton components from OpenCensus to OpenTelemetry for observability and tracing. | Dev Complete |
| [SRVKP-1801](https://redhat.atlassian.net/browse/SRVKP-1801) | Pipelines don't work with resource quota | Dev Complete |
