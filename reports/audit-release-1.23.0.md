# Release Audit: 1.23.0

**Source:** `reports/release_1.23.0.json`<br>
**Total commits:** 1308 across 16 components<br>
**Already linked (by Go tool):** 145 · **Newly linked:** 117 · **Unlinked:** 395 · **Dependabot (skipped):** 651<br>

## Index

| Component | Commits | Jira | Unsynced | Unmatched Jiras |
|-----------|---------|------|----------|-----------------|
| [git-init](#git-init) | error | — | — | — |
| [opc](#opc) | 3 | 3/3 | 0 | 2 |
| [console-plugin-pf5](#console-plugin-pf5) | 19 | 18/19 | 0 | 54 |
| [tekton-kueue](#tekton-kueue) | 2 | 1/2 | 0 | 0 |
| [pipelines-as-code](#pipelines-as-code) | 230 | 96/230 | 0 | 14 |
| [tektoncd-hub](#tektoncd-hub) | 5 | 1/5 | 0 | 0 |
| [tektoncd-pruner](#tektoncd-pruner) | 162 | 24/162 | 0 | 0 |
| [tektoncd-results](#tektoncd-results) | 116 | 10/116 | 0 | 1 |
| [operator](#operator) | 161 | 50/161 | 0 | 7 |
| [manual-approval-gate](#manual-approval-gate) | 2 | 1/2 | 0 | 1 |
| [tektoncd-chains](#tektoncd-chains) | 145 | 8/145 | 0 | 1 |
| [tektoncd-cli](#tektoncd-cli) | 71 | 7/71 | 0 | 1 |
| [tektoncd-pipeline](#tektoncd-pipeline) | 327 | 28/327 | 0 | 9 |
| [tektoncd-triggers](#tektoncd-triggers) | 60 | 11/60 | 0 | 0 |
| [syncer-service](#syncer-service) | 1 | 0/1 | 0 | 0 |
| [multicluster-proxy-aae](#multicluster-proxy-aae) | 0 | 0/0 | 0 | 0 |
| [tekton-assist](#tekton-assist) | error | — | — | — |
| [tekton-caches](#tekton-caches) | 4 | 4/4 | 0 | 0 |

---


<details>
<summary><h2>git-init</h2> — 0 commits, 0 linked</summary>


**Error:** failed to get head SHA at from-date: failed to list head file commits for openshift-pipelines/git-init: GET https://api.github.com/repos/openshift-pipelines/git-init/commits?path=head&per_page=1&sha=release-v1.22.x&until=2026-05-15T07%3A57%3A06Z: 404 Not Found []


</details>

---


<details>
<summary><h2>opc</h2> — 3 commits, 3 linked</summary>


**Upstream:** openshift-pipelines/opc · **Commits:** 3 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| fd19a79 | Update opc version 1.23.0 | Khurram Baig | [#458](https://github.com/openshift-pipelines/opc/pull/458) | [SRVKP-12016](https://redhat.atlassian.net/browse/SRVKP-12016) *(keyword)* | **Release Notes Not Required:** Internal version bump to 1.23.0. *(generated)* |
| bc04cca | add replace for deprecated hub repo | Shiv Verma | [#462](https://github.com/openshift-pipelines/opc/pull/462) | [SRVKP-11950](https://redhat.atlassian.net/browse/SRVKP-11950) *(keyword)* | **Deprecated Functionality:** The opc CLI now uses the openshift-pipelines/hub fork in place of the deprecated *(generated)* |
| b9be146 | Bump to Pipelines, PAC, Triggers and Chains LTS | Khurram Baig | [#466](https://github.com/openshift-pipelines/opc/pull/466) | [SRVKP-12016](https://redhat.atlassian.net/browse/SRVKP-12016) *(keyword)* | **Release Notes Not Required:** Internal dependency update bumping Pipelines as Code to 0.48.0, Tekton CLI to 0. *(generated)* |


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-10246](https://redhat.atlassian.net/browse/SRVKP-10246) | OCP Compatibility – Refactor Modals & Tables | Closed |
| [SRVKP-10245](https://redhat.atlassian.net/browse/SRVKP-10245) | OCP 4.22 Readiness – PatternFly 6 & React@18 Upgrade | Backlog |


</details>

---


<details>
<summary><h2>console-plugin-pf5</h2> — 19 commits, 18 linked</summary>


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


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-11715](https://redhat.atlassian.net/browse/SRVKP-11715) | Translation PR for all migrations in master branch | Release Pending |
| [SRVKP-11711](https://redhat.atlassian.net/browse/SRVKP-11711) | Upgrade to console sdk 4.22.0-prerelease.3 | Release Pending |
| [SRVKP-11708](https://redhat.atlassian.net/browse/SRVKP-11708) | CSS Collision: pipelines-plugin overrides .ocs-quick-search-modal styles in Console Topology | Closed |
| [SRVKP-11566](https://redhat.atlassian.net/browse/SRVKP-11566) | Replace <LoadingBox/> with Loading Component to remove PF 5 elements | Closed |
| [SRVKP-11565](https://redhat.atlassian.net/browse/SRVKP-11565) | Pipeline topology not working in Pipelinerun details page for Multicluster re-run | Closed |
| [SRVKP-11564](https://redhat.atlassian.net/browse/SRVKP-11564) | LogViewer not showing up for a failed task in multicluster | Closed |
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
| [SRVKP-11001](https://redhat.atlassian.net/browse/SRVKP-11001) | Yarn4 CI operator ( release repo ) Issue about name conflict for main_ocp_4.22 branch in the release repo config | Closed |
| [SRVKP-10953](https://redhat.atlassian.net/browse/SRVKP-10953) | Fix the issue of failed CI job (frontend) for main_ocp_4.22 | Release Pending |
| [SRVKP-10307](https://redhat.atlassian.net/browse/SRVKP-10307) | Upgrade react-redux | Release Pending |
| [SRVKP-10306](https://redhat.atlassian.net/browse/SRVKP-10306) | Upgrade react router to react router 7 | Release Pending |
| [SRVKP-10080](https://redhat.atlassian.net/browse/SRVKP-10080) | Fix User and Group tags CSS  in Approval task details page | Release Pending |
| [SRVKP-10077](https://redhat.atlassian.net/browse/SRVKP-10077) | Streaming event icon not appearing | Release Pending |
| [SRVKP-10076](https://redhat.atlassian.net/browse/SRVKP-10076) | Spacing between Pipelines and samples and snippet  | Release Pending |
| [SRVKP-10075](https://redhat.atlassian.net/browse/SRVKP-10075) | Alignment in Overview Page for all cards ( tekton result + prometheus ) | Release Pending |
| [SRVKP-10072](https://redhat.atlassian.net/browse/SRVKP-10072) | Deprecate modal.ts  | Release Pending |
| [SRVKP-10069](https://redhat.atlassian.net/browse/SRVKP-10069) | Fix Task ResourceIcon color to be consistent with other CRDs and hierarchy in TT and CTB | Release Pending |
| [SRVKP-10037](https://redhat.atlassian.net/browse/SRVKP-10037) | Pipeline builder topology not showing correct status when PLR is Rerun from the Action dropdown | Release Pending |
| [SRVKP-9983](https://redhat.atlassian.net/browse/SRVKP-9983) | Fix log viewer UI in PF6 | Release Pending |
| [SRVKP-9918](https://redhat.atlassian.net/browse/SRVKP-9918) | Deprecate CreateModalLauncher and use useOverlay hook for pipeline builder modal | Release Pending |
| [SRVKP-9789](https://redhat.atlassian.net/browse/SRVKP-9789) | List style on the Topology side panel for the Pipelines resource is not consistent with PF5 | Release Pending |
| [SRVKP-9788](https://redhat.atlassian.net/browse/SRVKP-9788) | Input field styling on start pipelines and dropdown style on the Start pipeline modal is not consistent with PF5 | Release Pending |
| [SRVKP-9787](https://redhat.atlassian.net/browse/SRVKP-9787) | PipelineRun details page sub header is broken, and the log viewer style is not in sync with PF5 | Release Pending |
| [SRVKP-9786](https://redhat.atlassian.net/browse/SRVKP-9786) | Fix the Dropdown Actions | Release Pending |
| [SRVKP-9785](https://redhat.atlassian.net/browse/SRVKP-9785) | Fix Pipeline Builder Modal Styles | Release Pending |
| [SRVKP-9784](https://redhat.atlassian.net/browse/SRVKP-9784) | Fix Border Styles in PF6  | Release Pending |
| [SRVKP-9783](https://redhat.atlassian.net/browse/SRVKP-9783) | Fix - Ensure Consistent Font Size | Release Pending |
| [SRVKP-9713](https://redhat.atlassian.net/browse/SRVKP-9713) | Migrate remaining plugin list/detail pages to ConsoleDataView and LazyActionMenu | Closed |
| [SRVKP-9712](https://redhat.atlassian.net/browse/SRVKP-9712) | Update all VirtualisedTable list to data-view for EventListener and Trigger Template tabs | Release Pending |
| [SRVKP-9711](https://redhat.atlassian.net/browse/SRVKP-9711) | Deprecate react-helmet and use DocumentTitle instead | Release Pending |
| [SRVKP-9710](https://redhat.atlassian.net/browse/SRVKP-9710) | Migrate Task and TaskRun List Views to DataView (React 18) | Release Pending |
| [SRVKP-9709](https://redhat.atlassian.net/browse/SRVKP-9709) | Update VirtualisedTable list to ConsoleDataView for repositories tab | Release Pending |
| [SRVKP-9627](https://redhat.atlassian.net/browse/SRVKP-9627) | Replace useModal with use useOverlay hook and add native PF6 Modal components for Start pipeline, Add Trigger, Remove Trigger and Approval modals | Release Pending |
| [SRVKP-9626](https://redhat.atlassian.net/browse/SRVKP-9626) | Update all VirtualisedTable list to data-view for pipeline and pipeline run tabs | Release Pending |
| [SRVKP-9625](https://redhat.atlassian.net/browse/SRVKP-9625) | Update page and section headings for Pipelines, Tasks and Triggers Page | Release Pending |
| [SRVKP-9392](https://redhat.atlassian.net/browse/SRVKP-9392) | Review and update variable and class names. | Release Pending |
| [SRVKP-9391](https://redhat.atlassian.net/browse/SRVKP-9391) | Remove all CSS overrides. | Release Pending |
| [SRVKP-9390](https://redhat.atlassian.net/browse/SRVKP-9390) | Update dependencies to use PatternFly 6 instead of PatternFly 5 | Release Pending |
| [SRVKP-9389](https://redhat.atlassian.net/browse/SRVKP-9389) | Run PF6 codemods suite | Release Pending |
| [SRVKP-9303](https://redhat.atlassian.net/browse/SRVKP-9303) | Discovery on the upgrade to PF6 | Release Pending |
| [SRVKP-8292](https://redhat.atlassian.net/browse/SRVKP-8292) | Cancelled/failed pipelineRun did not trigger finally tasks | Release Pending |


</details>

---


<details>
<summary><h2>tekton-kueue</h2> — 2 commits, 1 linked</summary>


**Upstream:** konflux-ci/tekton-kueue · **Commits:** 2 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 096a0e8 | fix(deps): update kubernetes packages to v0.35.5 | red-hat-konflux[bot] | [#430](https://github.com/konflux-ci/tekton-kueue/pull/430) | — | — |
| ec8c653 | cve: Fixes following CVEs (#466) | Pramod Bindal | [#430](https://github.com/konflux-ci/tekton-kueue/pull/430) | [SRVKP-11951](https://redhat.atlassian.net/browse/SRVKP-11951) *(unmatched)* | — |


</details>

---


<details>
<summary><h2>pipelines-as-code</h2> — 230 commits, 96 linked, 15 dependabot</summary>


**Upstream:** tektoncd/pipelines-as-code · **Commits:** 230 · **Unsynced:** 0 · **Dependabot:** 15

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
| 059c01a | refactor(github): consolidate JWT generation | Akshay Pant | [#2541](https://github.com/tektoncd/pipelines-as-code/pull/2541) | [SRVKP-10952](https://redhat.atlassian.net/browse/SRVKP-10952) | **Enhancement:** Consolidates GitHub JWT token generation into a single code path for better main *(generated)* |
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
| 900e837 | chore: add release-notes skill | Chmouel Boudjnah | [#2568](https://github.com/tektoncd/pipelines-as-code/pull/2568) | — | — |
| bb03a79 | feat: overhaul README and add link checker | Chmouel Boudjnah | [#2578](https://github.com/tektoncd/pipelines-as-code/pull/2578) | — | — |
| 3ab1ff8 | chore: remove PR CI and Jira skill automation | Chmouel Boudjnah | [#2581](https://github.com/tektoncd/pipelines-as-code/pull/2581) | — | — |
| 85ed765 | docs: simplify the pull request template | Chmouel Boudjnah | [#2581](https://github.com/tektoncd/pipelines-as-code/pull/2581) | — | — |
| d9a6599 | fix: add second gitlab token and group for e2e tests | Chmouel Boudjnah | [#2584](https://github.com/tektoncd/pipelines-as-code/pull/2584) | — | — |
| d7b0258 | Fix stale k8s.io/client-go version in go.mod | David Simansky | [#2585](https://github.com/tektoncd/pipelines-as-code/pull/2585) | [SRVKP-11437](https://redhat.atlassian.net/browse/SRVKP-11437) *(keyword)* | **Bug Fix:** Fixes a stale k8s.io/client-go dependency version in go.mod. *(generated)* |
| f973266 | chore: remove a bunch go.mod replace directives | Chmouel Boudjnah | [#2586](https://github.com/tektoncd/pipelines-as-code/pull/2586) | [SRVKP-12242](https://redhat.atlassian.net/browse/SRVKP-12242) *(keyword)* | **Release Notes Not Required:** Internal: remove a bunch go.mod replace directives *(generated)* |
| d3c9611 | fix: disable documentation link checker | Chmouel Boudjnah | [#2587](https://github.com/tektoncd/pipelines-as-code/pull/2587) | — | — |
| 8cc1747 | fix: fix unittest failure when validating gh url | Chmouel Boudjnah | [#2593](https://github.com/tektoncd/pipelines-as-code/pull/2593) | [SRVKP-10943](https://redhat.atlassian.net/browse/SRVKP-10943) *(keyword)* | **Bug Fix:** Fixes a unit test failure in GitHub URL validation. *(generated)* |
| 45aa2f7 | fix: Update /ok-to-test status only for Forgejo | Zaki Shaikh | [#2571](https://github.com/tektoncd/pipelines-as-code/pull/2571) | [SRVKP-11138](https://redhat.atlassian.net/browse/SRVKP-11138) | **Bug Fix:** The /ok-to-test status update is now restricted to Forgejo only, preventing inco *(generated)* |
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
| 8854274 | fix: remove unused secrets/delete permission from controller | Chmouel Boudjnah | [#2744](https://github.com/tektoncd/pipelines-as-code/pull/2744) | [SRVKP-10952](https://redhat.atlassian.net/browse/SRVKP-10952) *(keyword)* | **Enhancement:** Removes unused secrets/delete RBAC permission from the controller, following the *(generated)* |
| 223e39c | fix(ci): parse JSON test output for Slack notifications | Akshay Pant | [#2753](https://github.com/tektoncd/pipelines-as-code/pull/2753) | [SRVKP-10661](https://redhat.atlassian.net/browse/SRVKP-10661) *(keyword)* | **Release Notes Not Required:** Internal: parse JSON test output for Slack notifications *(generated)* |
| 2c03760 | fix(security): redact query string from incoming webhook log | Shubham Bhardwaj | [#2754](https://github.com/tektoncd/pipelines-as-code/pull/2754) | [SRVKP-10609](https://redhat.atlassian.net/browse/SRVKP-10609) *(keyword)* | **Bug Fix:** Redacts query string parameters from incoming webhook log entries, preventing se *(generated)* |
| bd262aa | chore: update incoming webhook legacy params deprecation mes | Zaki Shaikh | [#2757](https://github.com/tektoncd/pipelines-as-code/pull/2757) | [SRVKP-11511](https://redhat.atlassian.net/browse/SRVKP-11511) *(keyword)* | **Release Notes Not Required:** Internal: update incoming webhook legacy params deprecation message *(generated)* |
| 0017828 | fix(github): scope App token to triggering repo | Akshay Pant | [#2705](https://github.com/tektoncd/pipelines-as-code/pull/2705) | [SRVKP-12241](https://redhat.atlassian.net/browse/SRVKP-12241) | **Bug Fix:** Scopes GitHub App installation tokens to only the repository that triggered the  *(generated)* |
| ee5d9b0 | fix(resolve): deep-copy cached resources before inlining | Akshay Pant | [#2705](https://github.com/tektoncd/pipelines-as-code/pull/2705) | [SRVKP-12241](https://redhat.atlassian.net/browse/SRVKP-12241) | **Bug Fix:** Fixes a deep-copy issue where cached pipeline resources were mutated during inli *(generated)* |
| 402d5c7 | fix: prevent GitHub Enterprise header hijacking in app token | Chmouel Boudjnah | [#2759](https://github.com/tektoncd/pipelines-as-code/pull/2759) | [SRVKP-12241](https://redhat.atlassian.net/browse/SRVKP-12241) *(keyword)* | **Bug Fix:** Prevents GitHub Enterprise header hijacking in app token requests, fixing a secu *(generated)* |
| 0b6c487 | Release yaml generated from https://github.com/tektoncd/pipe | Pipelines as Code CI Robot | — | — | — |


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-12538](https://redhat.atlassian.net/browse/SRVKP-12538) | QA- [Regression] Relative task references in Pipeline broken | Closed |
| [SRVKP-12421](https://redhat.atlassian.net/browse/SRVKP-12421) | CLONE -  Upstream release notes for PAC | New |
| [SRVKP-11932](https://redhat.atlassian.net/browse/SRVKP-11932) | Delete Repository CR status field and do clean up in the PaC repo | Tasking and Estimation |
| [SRVKP-11931](https://redhat.atlassian.net/browse/SRVKP-11931) | Deprecate Repository CR PipelineRun status field | Backlog |
| [SRVKP-11549](https://redhat.atlassian.net/browse/SRVKP-11549) | AI analysis comment duplicated on each PR event instead of updating existing comment | In Progress |
| [SRVKP-11512](https://redhat.atlassian.net/browse/SRVKP-11512) | Add comment ownership validation to prevent editing other users' comments | Release Pending |
| [SRVKP-11464](https://redhat.atlassian.net/browse/SRVKP-11464) | Support 'max_completion_tokens' for newer OpenAI-compatible models in PAC LLM | Closed |
| [SRVKP-11101](https://redhat.atlassian.net/browse/SRVKP-11101) | [Testing] CLONE - Validate Repository CR URL path shape in admission webhook to prevent token scoping bypass | Closed |
| [SRVKP-10969](https://redhat.atlassian.net/browse/SRVKP-10969) | Adopting Pipelines-as-Code into Tekton Project | Closed |
| [SRVKP-10967](https://redhat.atlassian.net/browse/SRVKP-10967) | [Testing] Add comment ownership validation to prevent editing other users' comments | To Do |
| [SRVKP-10910](https://redhat.atlassian.net/browse/SRVKP-10910) | Automate GitHub App Token Extension for Private Forks | Tasking and Estimation |
| [SRVKP-10413](https://redhat.atlassian.net/browse/SRVKP-10413) | Migrate from using gitea sdk to forgejo sdk | Closed |
| [SRVKP-10266](https://redhat.atlassian.net/browse/SRVKP-10266) | Pipelines as Code integration with Forgejo | Dev Complete |
| [SRVKP-9663](https://redhat.atlassian.net/browse/SRVKP-9663) | Forgejo support for Pipelines as Code (PaC) | Closed |


</details>

---


<details>
<summary><h2>tektoncd-hub</h2> — 5 commits, 1 linked</summary>


**Upstream:** openshift-pipelines/hub · **Commits:** 5 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| c79258f | bump UI and swagger packages to fix the CVEs | Shiv Verma | — | — | **CVE:** Fixes CVE-2026-4800 and CVE-2026-29074 by updating the Hub UI and Swagger packag *(generated)* |
| dd532ff | feat(cli): deprecate tkn hub commands for Tekton Hub support | divyansh42 | — | [SRVKP-11950](https://redhat.atlassian.net/browse/SRVKP-11950) *(keyword)* | **Deprecated Functionality:** The tkn hub CLI commands that depend on Tekton Hub (check-upgrade, downgrade, ge *(generated)* |
| e22c048 | bump UI and swagger packages to fix the CVEs | Shiv Verma | — | — | **CVE:** Updates Hub UI and Swagger packages to address CVE vulnerabilities. *(generated)* |
| 5d7d39c | bump go version to v1.25.8 and update dockerfile | Shiv Verma | — | — | **Release Notes Not Required** *(generated)* |
| a84deeb | fix CVE vulnerablities in the UI and Swagger component | Shiv Verma | — | — | **CVE:** Fixes CVE vulnerabilities in the Hub UI and Swagger component. *(generated)* |


</details>

---


<details>
<summary><h2>tektoncd-pruner</h2> — 162 commits, 24 linked, 118 dependabot</summary>


**Upstream:** tektoncd/pruner · **Commits:** 162 · **Unsynced:** 0 · **Dependabot:** 118

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 9524430 | feat: introduce new function to expose validation logic on s | Anitha Natarajan | [#57](https://github.com/tektoncd/pruner/pull/57) | [SRVKP-9431](https://redhat.atlassian.net/browse/SRVKP-9431) | **Release Notes Not Required** *(generated)* |
| bef0990 | fix: add missing unit test cases | Anitha Natarajan | [#59](https://github.com/tektoncd/pruner/pull/59) | [SRVKP-9431](https://redhat.atlassian.net/browse/SRVKP-9431) | **Release Notes Not Required** *(generated)* |
| c9eae3c | feat: 0.3.2 retract and doc udpates | Anitha Natarajan | [#60](https://github.com/tektoncd/pruner/pull/60) | [SRVKP-9431](https://redhat.atlassian.net/browse/SRVKP-9431) | **Release Notes Not Required** *(generated)* |
| bd99f38 | Fix errors and improve formatting in release cheat sheet | Shubham Bhardwaj | [#87](https://github.com/tektoncd/pruner/pull/87) | [SRVKP-9463](https://redhat.atlassian.net/browse/SRVKP-9463) | **Release Notes Not Required** *(generated)* |
| e56420f | fix: resolve yamllint errors and warnings in workflows | Shubham Bhardwaj | [#86](https://github.com/tektoncd/pruner/pull/86) | [SRVKP-9463](https://redhat.atlassian.net/browse/SRVKP-9463) | **Release Notes Not Required** *(generated)* |
| 37446ec | .github/workflows: use CHATOPS_TOKEN for coverage comments | Vincent Demeester | [#92](https://github.com/tektoncd/pruner/pull/92) | — | — |
| b38434b | Add releases.md documenting release policy and current relea | Shubham Bhardwaj | [#91](https://github.com/tektoncd/pruner/pull/91) | [SRVKP-9473](https://redhat.atlassian.net/browse/SRVKP-9473) | **Release Notes Not Required** *(generated)* |
| 7787766 | feat: add cherry-pick slash command workflow | Vincent Demeester | [#109](https://github.com/tektoncd/pruner/pull/109) | — | — |
| 1a2dc82 | Remove tekton-* namespace exclusion from pruner | Shubham Bhardwaj | [#108](https://github.com/tektoncd/pruner/pull/108) | [SRVKP-9677](https://redhat.atlassian.net/browse/SRVKP-9677) | **Bug Fix:** Allow tekton-* namespaces to be processed by the pruner, except tekton-pipelines *(extracted)* |
| f85a120 | Removes the container resource requests and limits from the  | Shubham Bhardwaj | [#106](https://github.com/tektoncd/pruner/pull/106) | [SRVKP-9676](https://redhat.atlassian.net/browse/SRVKP-9676) | **Enhancement:** The pruner controller deployment no longer specifies container resource requests *(generated)* |
| d1773de | chore: migrate retest workflow to use plumbing reusable work | Vincent Demeester | [#110](https://github.com/tektoncd/pruner/pull/110) | — | — |
| b9e852c | fix: Update E2E tests for tekton-* namespace validation chan | Shubham Bhardwaj | [#111](https://github.com/tektoncd/pruner/pull/111) | [SRVKP-9677](https://redhat.atlassian.net/browse/SRVKP-9677) *(keyword)* | **Release Notes Not Required** *(generated)* |
| b5863e3 | Enable GitHub merge queue | Anitha Natarajan | [#112](https://github.com/tektoncd/pruner/pull/112) | — | — |
| 7685fa4 | ci: Remove version tip override from ko setup | Vincent Demeester | [#120](https://github.com/tektoncd/pruner/pull/120) | — | — |
| 39f2fa0 | chore: update workflows for dependabot and migrate to infra. | Shubham Bhardwaj | [#128](https://github.com/tektoncd/pruner/pull/128) | — | — |
| 8bb5cee | chore: add omitempty to YAML/JSON struct tags | Shubham Bhardwaj | [#127](https://github.com/tektoncd/pruner/pull/127) | [SRVKP-9968](https://redhat.atlassian.net/browse/SRVKP-9968) | **Enhancement:** Namespace config no longer shows null values for unset fields in TektonPruner gl *(extracted)* |
| 790b98d | Update pipelines-lts matrix to latest LTS versions | Shubham Bhardwaj | [#135](https://github.com/tektoncd/pruner/pull/135) | — | — |
| 8feec18 | fix: add JSON struct tags to fix TektonConfig status seriali | Shubham Bhardwaj | [#141](https://github.com/tektoncd/pruner/pull/141) | [SRVKP-9968](https://redhat.atlassian.net/browse/SRVKP-9968) | **Release Notes Not Required** *(extracted)* |
| c2e7a05 | chore: Update release documentation for v0.3.4 | Shubham Bhardwaj | [#138](https://github.com/tektoncd/pruner/pull/138) | — | — |
| aa31556 | fix: selector-based pruning for label and annotation selecto | Shubham Bhardwaj | [#142](https://github.com/tektoncd/pruner/pull/142) | [SRVKP-10028](https://redhat.atlassian.net/browse/SRVKP-10028) | **Bug Fix:** Selector-based pruning now correctly groups PipelineRuns by ConfigMap-defined la *(extracted)* |
| f452a90 | Chore: update release deatils and remove unused workflow | Anitha Natarajan | [#146](https://github.com/tektoncd/pruner/pull/146) | — | — |
| fa797de | docs: Fix documentation inconsistencies and incorrect YAML s | Shubham Bhardwaj | [#151](https://github.com/tektoncd/pruner/pull/151) | [SRVKP-7225](https://redhat.atlassian.net/browse/SRVKP-7225) | **Release Notes Not Required** *(extracted)* |
| cfef219 | chore(ci): update cherry-pick workflow to fix multi-commit P | Vincent Demeester | [#162](https://github.com/tektoncd/pruner/pull/162) | — | — |
| 7e15ff4 | docs: remove tekton-* and add tekton-operator | Aditya Shinde | [#170](https://github.com/tektoncd/pruner/pull/170) | [SRVKP-9677](https://redhat.atlassian.net/browse/SRVKP-9677) *(keyword)* | **Release Notes Not Required** *(extracted)* |
| 84ade14 | docs: improve punctuation | Ankur Sinha | [#171](https://github.com/tektoncd/pruner/pull/171) | — | — |
| 3a48be1 | chore(ci): include retest in slash workflow | Anitha Natarajan | [#172](https://github.com/tektoncd/pruner/pull/172) | — | — |
| a94b31c | ci: add unified CI summary job | Vincent Demeester | [#182](https://github.com/tektoncd/pruner/pull/182) | — | — |
| 69ea00b | fix: invoke e2e test workflow from the ci workflow | Anitha Natarajan | [#189](https://github.com/tektoncd/pruner/pull/189) | — | — |
| 514bbf7 | fix: doc typos | Naomi Gelman | [#183](https://github.com/tektoncd/pruner/pull/183) | — | — |
| d3c6f36 | chore: move architecture diagram to root | Shubham Bhardwaj | [#188](https://github.com/tektoncd/pruner/pull/188) | — | — |
| 9e45a7c | chore(deps): replace go.opentelemetry.io/otel/sdk v1.39.0 wi | Andrew Thorp | [#165](https://github.com/tektoncd/pruner/pull/165) | — | — |
| 4970aaf | Move inactive approvers/reviewers to alumni | Vincent Demeester | [#202](https://github.com/tektoncd/pruner/pull/202) | [SRVKP-7687](https://redhat.atlassian.net/browse/SRVKP-7687) | **Release Notes Not Required** *(extracted)* |
| 379afd0 | Bump pipeline version & knative.dev/pkg version | Shubham Bhardwaj | [#220](https://github.com/tektoncd/pruner/pull/220) | [SRVKP-9325](https://redhat.atlassian.net/browse/SRVKP-9325) | **Release Notes Not Required** *(extracted)* |
| 83c2eb2 | docs: clean up observability config and update metrics docum | Shubham Bhardwaj | [#221](https://github.com/tektoncd/pruner/pull/221) | — | — |
| af2c740 | Bump knative.dev/pkg to adopt centralized WEBHOOK_* TLS conf | Jawed khelil | [#231](https://github.com/tektoncd/pruner/pull/231) | [SRVKP-9644](https://redhat.atlassian.net/browse/SRVKP-9644) *(keyword)* | **Enhancement:** The pruner now supports centralized WEBHOOK_* TLS configuration from knative.dev *(generated)* |
| d488086 | add scorecard workflow | Tyler Auerbeck | [#226](https://github.com/tektoncd/pruner/pull/226) | — | — |
| 18ffdcc | fix cherry-pick failure | Anitha Natarajan | [#236](https://github.com/tektoncd/pruner/pull/236) | — | — |
| f0eb10d | chore: add zizmor workflow and fix GitHub Actions security f | Shubham Bhardwaj | [#262](https://github.com/tektoncd/pruner/pull/262) | [SRVKP-12065](https://redhat.atlassian.net/browse/SRVKP-12065) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 8e317b9 | tekton: automate releases with Pipelines-as-Code | Shubham Bhardwaj | [#246](https://github.com/tektoncd/pruner/pull/246) | [SRVKP-11528](https://redhat.atlassian.net/browse/SRVKP-11528) | **Release Notes Not Required** *(generated)* |
| a01f51e | fix: update release pipeline for oci-ci-cd cluster compatibi | Shubham Bhardwaj | [#284](https://github.com/tektoncd/pruner/pull/284) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 1dacc8b | fix: add e2e build tag to skip e2e tests in release pipeline | Shubham Bhardwaj | [#289](https://github.com/tektoncd/pruner/pull/289) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 62c4f57 | ci: add updated pipelines version and k8s version | Shubham Bhardwaj | [#290](https://github.com/tektoncd/pruner/pull/290) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 2d108c1 | feat: add draft release and Chains signing to release pipeli | Shubham Bhardwaj | [#296](https://github.com/tektoncd/pruner/pull/296) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 10ac09c | ci: diable automatic dependency review workflow | Shubham Bhardwaj | [#297](https://github.com/tektoncd/pruner/pull/297) | [SRVKP-11622](https://redhat.atlassian.net/browse/SRVKP-11622) *(keyword)* | **Release Notes Not Required** *(generated)* |


</details>

---


<details>
<summary><h2>tektoncd-results</h2> — 116 commits, 10 linked, 88 dependabot</summary>


**Upstream:** tektoncd/results · **Commits:** 116 · **Unsynced:** 0 · **Dependabot:** 88

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 1e1a0ec | feature: update release cheatsheet and pipeline (#1176) | Anitha Natarajan | [#1176](https://github.com/tektoncd/results/pull/1176) | — | — |
| 2b493c4 | Nominate divyansh42 as results approver and reviewer | Vincent Demeester | [#1247](https://github.com/tektoncd/results/pull/1247) | — | — |
| 7ac3dc4 | ci: Add CI summary fan-in job to presubmit workflow | Vincent Demeester | [#1225](https://github.com/tektoncd/results/pull/1225) | — | — |
| 1907cb6 | CI: Remove manual PR author check from presubmit workflow | Khurram Baig | [#1197](https://github.com/tektoncd/results/pull/1197) | — | — |
| fcf7dd4 | Improve finalizer error visibility | Emil Natan | [#1252](https://github.com/tektoncd/results/pull/1252) | — | — |
| 8f3629e | Change all occurences of GCS buckets with OCI buckets | adityavshinde | [#1257](https://github.com/tektoncd/results/pull/1257) | [SRVKP-11090](https://redhat.atlassian.net/browse/SRVKP-11090) | **Release Notes Not Required** *(generated)* |
| 1bc1349 | Revert OCI buckets with GCS buckets which was mistakenly rep | adityavshinde | [#1257](https://github.com/tektoncd/results/pull/1257) | [SRVKP-11090](https://redhat.atlassian.net/browse/SRVKP-11090) | **Release Notes Not Required** *(generated)* |
| f1ea6a8 | Revert changed Dockerfile | adityavshinde | [#1257](https://github.com/tektoncd/results/pull/1257) | [SRVKP-11090](https://redhat.atlassian.net/browse/SRVKP-11090) | **Release Notes Not Required** *(generated)* |
| 9dabb77 | Move inactive approvers and reviewers to alumni | Vincent Demeester | [#1246](https://github.com/tektoncd/results/pull/1246) | — | — |
| a4eacc0 | fix(watcher): fix TestController for client-go v0.35 | Khurram Baig | [#1281](https://github.com/tektoncd/results/pull/1281) | — | — |
| 007212d | chore: update deprecated grpc_middleware components | Tyler Auerbeck | [#1274](https://github.com/tektoncd/results/pull/1274) | — | — |
| af03e66 | ci: fix GitHub Actions security issues found by zizmor | ab-ghosh | [#1293](https://github.com/tektoncd/results/pull/1293) | — | — |
| b5ff73d | ci: add zizmor GitHub Actions security analysis | ab-ghosh | [#1293](https://github.com/tektoncd/results/pull/1293) | — | — |
| 4362545 | ci: fix remaining zizmor findings (template injection) | ab-ghosh | [#1293](https://github.com/tektoncd/results/pull/1293) | — | — |
| 2b40f90 | chore(deps): Update dependencies for OpenTelemetry migration | divyansh42 | [#1249](https://github.com/tektoncd/results/pull/1249) | [SRVKP-8533](https://redhat.atlassian.net/browse/SRVKP-8533) | **Release Notes Not Required** *(generated)* |
| 9b6bf42 | feat(metrics): Migrate from OpenCensus to OpenTelemetry | divyansh42 | [#1249](https://github.com/tektoncd/results/pull/1249) | [SRVKP-8533](https://redhat.atlassian.net/browse/SRVKP-8533) | **Enhancement:** Migrates Results Watcher metrics from OpenCensus to OpenTelemetry. The observabi *(extracted)* |
| b3706b2 | chore(cli): deprecate legacy commands and flags | divyansh42 | [#1301](https://github.com/tektoncd/results/pull/1301) | [SRVKP-11434](https://redhat.atlassian.net/browse/SRVKP-11434) | **Deprecated Functionality:** Deprecates legacy CLI commands (`result`, `records`, `logs`) and several global   |
| c73b3fb | fix(config): trim whitespace from host, token, and API path  | divyansh42 | [#1314](https://github.com/tektoncd/results/pull/1314) | — | — |
| 3d4e84e | Bump Kind k8s version to 1.35 | Emil Natan | [#1319](https://github.com/tektoncd/results/pull/1319) | — | — |
| 3f99f11 | perf(cli): use exact match for describe/logs command | divyansh42 | [#1283](https://github.com/tektoncd/results/pull/1283) | [SRVKP-11380](https://redhat.atlassian.net/browse/SRVKP-11380) | **Enhancement:** Enhacement: Improved performance for `describe` and `logs` commands with optimiz  |
| 8a63de5 | fix(logs): use ListRecords for dev CLI log listing | divyansh42 | [#1306](https://github.com/tektoncd/results/pull/1306) | [SRVKP-11301](https://redhat.atlassian.net/browse/SRVKP-11301) | **Bug Fix:** Fix `logs list` command failing with "unknown service tekton.results.v1alpha2.Lo  |
| 2e08db1 | Fix typos in API docs | Alan Greene | [#1337](https://github.com/tektoncd/results/pull/1337) | — | — |
| 589e0cd | feat(logs): added new logql config and documentation | Florian Thiévent | [#1210](https://github.com/tektoncd/results/pull/1210) | — | — |
| ac20152 | Fix race condition in the TestController | Emil Natan | [#1351](https://github.com/tektoncd/results/pull/1351) | — | — |
| cc6cc17 | Add agentic workflows context files | Emil Natan | [#1345](https://github.com/tektoncd/results/pull/1345) | [SRVKP-11817](https://redhat.atlassian.net/browse/SRVKP-11817) | **Release Notes Not Required**  |
| 208e14a | Add support for storing CustomRuns | Emil Natan | [#1320](https://github.com/tektoncd/results/pull/1320) | [SRVKP-11855](https://redhat.atlassian.net/browse/SRVKP-11855) *(keyword)* | **Feature:** The Results Watcher now watches and stores CustomRun resources (tekton.dev/v1bet *(extracted)* |
| d3dc464 | chore: switch base image to ghcr.io/tektoncd/plumbing/static | Vincent Demeester | [#1357](https://github.com/tektoncd/results/pull/1357) | — | — |
| ecca90d | Fix release pipelines | Emil Natan | [#1359](https://github.com/tektoncd/results/pull/1359) | — | — |


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-10269](https://redhat.atlassian.net/browse/SRVKP-10269) | Extend the list of managed resources | Refinement |


</details>

---


<details>
<summary><h2>operator</h2> — 161 commits, 50 linked, 71 dependabot</summary>


**Upstream:** tektoncd/operator · **Commits:** 161 · **Unsynced:** 0 · **Dependabot:** 71

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| fb4751e | chore: update the prechech pipeline version | Anitha Natarajan | [#3261](https://github.com/tektoncd/operator/pull/3261) | — | — |
| 077462a | Move inactive approvers to alumni | Vincent Demeester | [#3267](https://github.com/tektoncd/operator/pull/3267) | — | — |
| a07ab6d | Add TektonScheduler and TektonMulticlusterProxyAAE CRDs to H | mbpavan | [#3265](https://github.com/tektoncd/operator/pull/3265) | — | — |
| d7304ec | Nominate pratap0007 as operator approver | Vincent Demeester | [#3268](https://github.com/tektoncd/operator/pull/3268) | — | — |
| 9f82475 | fix: pass httpClient to PAC SyncConfig | Shiv Verma | [#3264](https://github.com/tektoncd/operator/pull/3264) | — | — |
| b5bb818 | chore: add pramod as approver for operator | Anitha Natarajan | [#3278](https://github.com/tektoncd/operator/pull/3278) | — | — |
| e8451e5 | fix readme and bundle script update | mbpavan | [#3277](https://github.com/tektoncd/operator/pull/3277) | [SRVKP-11045](https://redhat.atlassian.net/browse/SRVKP-11045) | **Release Notes Not Required** *(extracted)* |
| 4f93b09 | chore: bump hub from v1.23.7 to v1.23.8 | Vincent Demeester | [#3279](https://github.com/tektoncd/operator/pull/3279) | — | — |
| 9e20b9c | Fix TektonInstallerSet deadlock when resources have deletion | Jawed khelil | [#3217](https://github.com/tektoncd/operator/pull/3217) | [SRVKP-10858](https://redhat.atlassian.net/browse/SRVKP-10858) | **Release Note Not Required** *(extracted)* |
| 4a89460 | fix some minor issues | mbpavan | [#3282](https://github.com/tektoncd/operator/pull/3282) | [SRVKP-11089](https://redhat.atlassian.net/browse/SRVKP-11089) | **Bug Fix** *(extracted)* |
| 1fa6522 | chore: bump component versions | Vincent Demeester | [#3281](https://github.com/tektoncd/operator/pull/3281) | [SRVKP-12488](https://redhat.atlassian.net/browse/SRVKP-12488) *(keyword)* | **Release Notes Not Required:** ** [v1.10.0...v1.10.1](https://github.com/tektoncd/pipeline/compare/v1.10.0...v1 *(extracted)* |
| 122e2a3 | chore: bump pipeline from v1.10.1 to v1.10.2 | Vincent Demeester | [#3291](https://github.com/tektoncd/operator/pull/3291) | — | — |
| dd3b702 | Fix authorization field in pipeline_console_plugin.yaml | Khurram | [#3293](https://github.com/tektoncd/operator/pull/3293) | [SRVKP-11107](https://redhat.atlassian.net/browse/SRVKP-11107) | **Release Notes Not Required** *(extracted)* |
| 2419609 | chore(deps): bump google.golang.org/grpc from 1.79.1 to 1.79 | Anitha Natarajan | [#3302](https://github.com/tektoncd/operator/pull/3302) | — | — |
| 4a11bc1 | fix-cherry-pick-failures | Anitha Natarajan | [#3323](https://github.com/tektoncd/operator/pull/3323) | — | — |
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
| d8985d0 | fix: remove YAML document separator from .tekton files | Vincent Demeester | [#3342](https://github.com/tektoncd/operator/pull/3342) | — | — |
| 5c00da3 | update subject for tekton-scheduler-role to tekton-operator | Pramod Bindal | [#3361](https://github.com/tektoncd/operator/pull/3361) | — | — |
| ef45264 | fix: use correct fully qualified image names in release pipe | Vincent Demeester | [#3362](https://github.com/tektoncd/operator/pull/3362) | — | — |
| 71ed26f | Add Pipelines-as-Code on Kubernetes and resolve openshift-pi | mbpavan | [#3337](https://github.com/tektoncd/operator/pull/3337) | [SRVKP-11429](https://redhat.atlassian.net/browse/SRVKP-11429) | **Feature:** The Tekton Operator can now install and reconcile Pipelines-as-Code (PAC) on van *(extracted)* |
| 2a8ae75 | Use dedicated release ServiceAccount in PAC release pipeline | Vincent Demeester | [#3364](https://github.com/tektoncd/operator/pull/3364) | — | — |
| 8eebadf | chore: bump pipeline from v1.11.0 to v1.11.1 | Vincent Demeester | [#3365](https://github.com/tektoncd/operator/pull/3365) | — | — |
| ac736e5 | fix: use source-subpath/release-subpath params for create-dr | Vincent Demeester | [#3366](https://github.com/tektoncd/operator/pull/3366) | — | — |
| 4f23998 | chore: bump chains from v0.26.2 to v0.26.3 | Vincent Demeester | [#3367](https://github.com/tektoncd/operator/pull/3367) | — | — |
| 1efe55b | fix(release): correct devel version rewriting in publish-ope | Anitha Natarajan | [#3370](https://github.com/tektoncd/operator/pull/3370) | [SRVKP-12011](https://redhat.atlassian.net/browse/SRVKP-12011) *(unmatched)* | — |
| f028c33 | chore: bump hub from v1.23.9 to v1.23.10 | Vincent Demeester | [#3375](https://github.com/tektoncd/operator/pull/3375) | — | — |
| 1f7ac47 | Bump dependencies for OpenTelemetry migration | Shubham Bhardwaj | [#3332](https://github.com/tektoncd/operator/pull/3332) | [SRVKP-9320](https://redhat.atlassian.net/browse/SRVKP-9320) | **Enhancement:** Migrated operator metrics from OpenCensus to OpenTelemetry.  ACTION REQUIRED:  1 *(extracted)* |
| 6df3de1 | Update generated code for new dependencies | Shubham Bhardwaj | [#3332](https://github.com/tektoncd/operator/pull/3332) | [SRVKP-9320](https://redhat.atlassian.net/browse/SRVKP-9320) | **Enhancement:** Migrated operator metrics from OpenCensus to OpenTelemetry.  ACTION REQUIRED:  1 *(extracted)* |
| 161a1c4 | Migrate metrics from OpenCensus to OpenTelemetry | Shubham Bhardwaj | [#3332](https://github.com/tektoncd/operator/pull/3332) | [SRVKP-9320](https://redhat.atlassian.net/browse/SRVKP-9320) | **Enhancement:** Migrated operator metrics from OpenCensus to OpenTelemetry.  ACTION REQUIRED:  1 *(extracted)* |
| e3b4e28 | Remove unused custom operator metrics | Shubham Bhardwaj | [#3332](https://github.com/tektoncd/operator/pull/3332) | [SRVKP-9320](https://redhat.atlassian.net/browse/SRVKP-9320) | **Enhancement:** Migrated operator metrics from OpenCensus to OpenTelemetry.  ACTION REQUIRED:  1 *(extracted)* |
| 70798fa | chore: bump pipeline from v1.11.1 to v1.12.0 | Vincent Demeester | [#3395](https://github.com/tektoncd/operator/pull/3395) | — | — |
| 4db29a4 | Feat: Console Plugin Image should be picked conditionally ba | Pramod Bindal | [#3386](https://github.com/tektoncd/operator/pull/3386) | [SRVKP-11927](https://redhat.atlassian.net/browse/SRVKP-11927) | **Enhancement:** Added support for conditionally picking the console-plugin image on openshift. O *(extracted)* |
| 6a1cb9e | pin base image | Anitha Natarajan | [#3396](https://github.com/tektoncd/operator/pull/3396) | — | — |
| f558e68 | fix: update the koExtraArgs | Anitha Natarajan | [#3400](https://github.com/tektoncd/operator/pull/3400) | — | — |
| cecf39b | fix: missing subjects in attestation and capture UUID instea | Anitha Natarajan | [#3402](https://github.com/tektoncd/operator/pull/3402) | — | — |
| 694775d | chore: bump component versions | Vincent Demeester | [#3404](https://github.com/tektoncd/operator/pull/3404) | [SRVKP-12488](https://redhat.atlassian.net/browse/SRVKP-12488) *(keyword)* | **Release Notes Not Required:** ** [v0.45.0...v0.46.0](https://github.com/tektoncd/pipelines-as-code/compare/v0. *(extracted)* |
| 7c96af4 | Generate OpenAPI v3 schemas for all CRDs using controller-ge | Alessandro | [#3340](https://github.com/tektoncd/operator/pull/3340) | [SRVKP-12477](https://redhat.atlassian.net/browse/SRVKP-12477) *(keyword)* | **Enhancement:** CRD manifests now include full OpenAPI v3 schemas generated for all operator.tek *(extracted)* |
| c79fc55 | fix: helm automation with correct version | Anitha Natarajan | [#3405](https://github.com/tektoncd/operator/pull/3405) | — | — |
| a9cb7e4 | fix(cve): update Go to 1.25.10 to address stdlib security vu | khelil | [#3409](https://github.com/tektoncd/operator/pull/3409) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Security fix: update Go from 1.25.8 to 1.25.10 to address 10 standard library vu *(extracted)* |
| f2f55b4 |  fix(tekton-results): use passthrough TLS termination for ro | divyansh42 | [#3425](https://github.com/tektoncd/operator/pull/3425) | [SRVKP-11697](https://redhat.atlassian.net/browse/SRVKP-11697) | **Bug Fix:** Tekton Results API route now uses passthrough TLS termination by default, enabli *(extracted)* |
| 5dfea2e | fix: update-tektoncd-task-versions workflow to resolve CI fa | Shubham Bhardwaj | [#3426](https://github.com/tektoncd/operator/pull/3426) | [SRVKP-12022](https://redhat.atlassian.net/browse/SRVKP-12022) | **Release Notes Not Required** *(extracted)* |
| 345f8a3 | fix: harden patch-release workflow against script injection | Vincent Demeester | [#3424](https://github.com/tektoncd/operator/pull/3424) | — | — |
| ba5b11d | update task-maven to 0.4.1 | ab-ghosh | [#3415](https://github.com/tektoncd/operator/pull/3415) | [SRVKP-11954](https://redhat.atlassian.net/browse/SRVKP-11954) | **Release Notes Not Required** *(extracted)* |
| 4fec297 | feat(tls): make central TLS opt-out and enable ML-KEM for co | Jawed khelil | [#3416](https://github.com/tektoncd/operator/pull/3416) | [SRVKP-11957](https://redhat.atlassian.net/browse/SRVKP-11957) | **Feature:** Central TLS configuration is now enabled by default on OpenShift. The `enableCen *(extracted)* |
| bcfa313 | fix: update SendCloudEventsForRuns default and mark deprecat | Anitha Natarajan | [#3418](https://github.com/tektoncd/operator/pull/3418) | — | — |
| e1484f9 | refactor(tls): extract SetupAPIServerTLSWatch into occommon  | Jawed khelil | [#3406](https://github.com/tektoncd/operator/pull/3406) | [SRVKP-9613](https://redhat.atlassian.net/browse/SRVKP-9613) | **Feature:** On OpenShift, the tekton-operator-webhook and tekton-operator-proxy-webhook now  *(extracted)* |
| 534fc37 | feat(webhook/tls): read OpenShift APIServer TLS profile at s | Jawed khelil | [#3406](https://github.com/tektoncd/operator/pull/3406) | [SRVKP-9613](https://redhat.atlassian.net/browse/SRVKP-9613) | **Feature:** On OpenShift, the tekton-operator-webhook and tekton-operator-proxy-webhook now  *(extracted)* |
| 4975b71 | feat(proxy-webhook/tls): read OpenShift APIServer TLS profil | Jawed khelil | [#3406](https://github.com/tektoncd/operator/pull/3406) | [SRVKP-9613](https://redhat.atlassian.net/browse/SRVKP-9613) | **Feature:** On OpenShift, the tekton-operator-webhook and tekton-operator-proxy-webhook now  *(extracted)* |
| 6888701 | feat(tls): inject centrally managed TLS config into tekton-p | Jawed khelil | [#3383](https://github.com/tektoncd/operator/pull/3383) | [SRVKP-9614](https://redhat.atlassian.net/browse/SRVKP-9614) | **Feature:** The tekton-pipelines-webhook now inherits TLS configuration (minimum version and *(extracted)* |
| a9c8470 | feat(tls): inject centrally managed TLS config into tekton-t | Jawed khelil | [#3384](https://github.com/tektoncd/operator/pull/3384) | [SRVKP-9615](https://redhat.atlassian.net/browse/SRVKP-9615) | **Feature:** The tekton-triggers-webhook and tekton-triggers-core-interceptor  now inherits T *(extracted)* |
| d43fca7 | feat(tls): inject TLS env vars into triggers core intercepto | Jawed khelil | [#3384](https://github.com/tektoncd/operator/pull/3384) | [SRVKP-9615](https://redhat.atlassian.net/browse/SRVKP-9615) | **Feature:** The tekton-triggers-webhook and tekton-triggers-core-interceptor  now inherits T *(extracted)* |
| 273cd27 | docs(install): document that default namespace is unsupporte | Matt Van Horn | [#3440](https://github.com/tektoncd/operator/pull/3440) | [SRVKP-12017](https://redhat.atlassian.net/browse/SRVKP-12017) *(PR body)* | **Release Notes Not Required** *(extracted)* |
| a2af438 | docs(install): fix broken OpenShift namespaces link | Matt Van Horn | [#3440](https://github.com/tektoncd/operator/pull/3440) | [SRVKP-12017](https://redhat.atlassian.net/browse/SRVKP-12017) *(PR body)* | **Release Notes Not Required** *(extracted)* |
| e363e1a | fix(cve): GO-2026-5019/GO-2026-5018 - update golang.org/x/cr | Jawed Khelil | [#3451](https://github.com/tektoncd/operator/pull/3451) | [SRVKP-6935](https://redhat.atlassian.net/browse/SRVKP-6935) *(keyword)* | **CVE:** Updates golang.org/x/crypto from v0.50.0 to v0.52.0 to address GO-2026-5019 and  *(generated)* |
| a4e8dba | chore: fix failing unit tests during upgrade | Anitha Natarajan | [#3432](https://github.com/tektoncd/operator/pull/3432) | — | — |
| 4573cf5 | Add centrally managed TLS configuration for console-plugin n | Jawed khelil | [#3218](https://github.com/tektoncd/operator/pull/3218) | [SRVKP-9632](https://redhat.atlassian.net/browse/SRVKP-9632) | **Feature:** The console-plugin nginx server now inherits TLS settings from the centrally man *(extracted)* |
| 82150f0 | chore(deps): update postgres-15 image for Openshift | Andrew Thorp | [#3455](https://github.com/tektoncd/operator/pull/3455) | [SRVKP-8836](https://redhat.atlassian.net/browse/SRVKP-8836) *(keyword)* | **Release Notes Not Required:** Update images digests for Openshift images *(extracted)* |
| 1c1be97 | chore(deps): Upgrade image SHAs for Openshift images | Andrew Thorp | [#3455](https://github.com/tektoncd/operator/pull/3455) | [SRVKP-3348](https://redhat.atlassian.net/browse/SRVKP-3348) *(keyword)* | **Release Notes Not Required:** Update images digests for Openshift images *(extracted)* |
| a8699d8 | feat(tls): inject centrally managed TLS config into pipeline | Jawed khelil | [#3385](https://github.com/tektoncd/operator/pull/3385) | [SRVKP-9616](https://redhat.atlassian.net/browse/SRVKP-9616) | **Feature:** On OpenShift, the `pipelines-as-code-webhook` deployment now automatically inher *(extracted)* |
| c5b1051 | fix(cve): GO-2026-5026 - update golang.org/x/net v0.53.0 → v | Jawed Khelil | [#3452](https://github.com/tektoncd/operator/pull/3452) | [SRVKP-6924](https://redhat.atlassian.net/browse/SRVKP-6924) *(keyword)* | **CVE:** Updates golang.org/x/net from v0.53.0 to v0.55.0 to address GO-2026-5026 securit *(generated)* |
| 77a2f42 | feat(tls): inject centrally managed TLS config into pruner w | Shubham Bhardwaj | [#3453](https://github.com/tektoncd/operator/pull/3453) | [SRVKP-9623](https://redhat.atlassian.net/browse/SRVKP-9623) *(keyword)* | **Feature:** Central TLS configuration is now injected into the tekton-pruner-webhook on Open *(extracted)* |
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
| ffdc96b | fix(helm): add openshift-pipelines.org RBAC to kubernetes ch | Jawed khelil | [#3579](https://github.com/tektoncd/operator/pull/3579) | [SRVKP-2270](https://redhat.atlassian.net/browse/SRVKP-2270) *(keyword)* | **Bug Fix:** Fix ManualApprovalGate installation failure on Kubernetes when using the Helm ch *(extracted)* |
| f0f903c | Cleanup manual CRDs and use genearted CRDs in kustomize | Pramod Bindal | [#3580](https://github.com/tektoncd/operator/pull/3580) | [SRVKP-8082](https://redhat.atlassian.net/browse/SRVKP-8082) *(keyword)* | **Release Notes Not Required** *(extracted)* |
| 5c9be32 | Cleanup manual CRDs and use genearted CRDs in kustomize | Pramod Bindal | [#3580](https://github.com/tektoncd/operator/pull/3580) | — | — |


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-12577](https://redhat.atlassian.net/browse/SRVKP-12577) | [OSP 1.23.0] TektonHub CR status.version field is empty causing version verification test to fail | New |
| [SRVKP-12576](https://redhat.atlassian.net/browse/SRVKP-12576) | [OSP 1.23.0] tkn and tkn-pac version mismatch in ecosystem tests - env config expects 0.45/0.48 but cluster has 0.41/0.42 | To Do |
| [SRVKP-11929](https://redhat.atlassian.net/browse/SRVKP-11929) | [Reopen] Default TekonConfig is created when it should not | Tasking and Estimation |
| [SRVKP-10928](https://redhat.atlassian.net/browse/SRVKP-10928) | Add self-healing for CA bundle configmaps in user namespaces | Closed |
| [SRVKP-10469](https://redhat.atlassian.net/browse/SRVKP-10469) | [operator] Fetch TLS config and expose to components  | Closed |
| [SRVKP-10254](https://redhat.atlassian.net/browse/SRVKP-10254) | Fetch TLS Settings from a centrallised location | In Progress |
| [SRVKP-7899](https://redhat.atlassian.net/browse/SRVKP-7899) | Migrate Tekton components from OpenCensus to OpenTelemetry for observability and tracing. | Dev Complete |


</details>

---


<details>
<summary><h2>manual-approval-gate</h2> — 2 commits, 1 linked</summary>


**Upstream:** openshift-pipelines/manual-approval-gate · **Commits:** 2 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 5932c2f | Update go and knative version | divyansh42 | [#965](https://github.com/openshift-pipelines/manual-approval-gate/pull/965) | [SRVKP-11514](https://redhat.atlassian.net/browse/SRVKP-11514) | **Enhancement:** ```
  Updated dependencies including Kubernetes (v0.35.3), Tekton Pipelines (v1.  |
| edb306f | Update go version in workflows and remove konflux related fi | divyansh42 | [#967](https://github.com/openshift-pipelines/manual-approval-gate/pull/967) | — | **Release Notes Not Required**  |


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-12484](https://redhat.atlassian.net/browse/SRVKP-12484) | Manual Approval Gate Component Missing from Deployment in OpenShift Pipelines 1.23.0 | On QA |


</details>

---


<details>
<summary><h2>tektoncd-chains</h2> — 145 commits, 8 linked, 97 dependabot</summary>


**Upstream:** tektoncd/chains · **Commits:** 145 · **Unsynced:** 0 · **Dependabot:** 97

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 21dfd1e | add release v0.26.0 to release documentation (#1468) | khelil | [#1468](https://github.com/tektoncd/chains/pull/1468) | — | — |
| 37297be | fix:improve e2e test failure analysis with chains controller | Anitha Natarajan | [#1470](https://github.com/tektoncd/chains/pull/1470) | — | — |
| ec6157c | add anithapriyanatarajan as maintainer (#1473) | khelil | [#1473](https://github.com/tektoncd/chains/pull/1473) | — | — |
| 3754523 | Add GitHub Actions workflow for build and unit tests (#1446) | Alan Greene | [#1446](https://github.com/tektoncd/chains/pull/1446) | — | — |
| 59c9aa9 | docs: fix cosign repo reference (#1434) | Emmanuel Ferdman | [#1434](https://github.com/tektoncd/chains/pull/1434) | — | — |
| 61ead14 | .github/workflows: use CHATOPS_TOKEN for coverage comments ( | Vincent Demeester | [#1480](https://github.com/tektoncd/chains/pull/1480) | — | — |
| 52e0f7c | feat: add retest workflow using plumbing reusable workflow ( | Vincent Demeester | [#1488](https://github.com/tektoncd/chains/pull/1488) | — | — |
| 93b3806 | ci: consolidate and modernize GitHub Actions workflows (#149 | Vincent Demeester | [#1490](https://github.com/tektoncd/chains/pull/1490) | — | — |
| 965df1d | feat: add cherry-pick slash command workflow (#1487) | Vincent Demeester | [#1487](https://github.com/tektoncd/chains/pull/1487) | — | — |
| 295ad99 | Refactor the annotations management (#1422) | Emil Natan | [#1422](https://github.com/tektoncd/chains/pull/1422) | — | — |
| 503e560 | chore: replace gcs storage path with infra.tekton.dev (#1491 | Anitha Natarajan | [#1491](https://github.com/tektoncd/chains/pull/1491) | — | — |
| f1f3796 | chore(release-pipeline): improve release cheat sheet structu | Shubham Bhardwaj | [#1489](https://github.com/tektoncd/chains/pull/1489) | — | — |
| 128c252 | ci: fix and modernize e2e workflow step ordering (#1492) | Vincent Demeester | [#1492](https://github.com/tektoncd/chains/pull/1492) | — | — |
| 2032c9a | chore: fix some minor issues in comments | socialsister | [#1403](https://github.com/tektoncd/chains/pull/1403) | — | — |
| 8763405 | fix: microshift e2e test failures on merge (#1500) | Anitha Natarajan | [#1500](https://github.com/tektoncd/chains/pull/1500) | — | — |
| 920e627 | Remove GHCR migration notice from the readme | Alan Greene | [#1499](https://github.com/tektoncd/chains/pull/1499) | — | — |
| 5e01cd9 | docs: separate command and output in tutorial | ab-ghosh | [#1498](https://github.com/tektoncd/chains/pull/1498) | — | — |
| 4b42e7f | fix: replace hard codede go ref from mod file (#1501) | Anitha Natarajan | [#1501](https://github.com/tektoncd/chains/pull/1501) | — | — |
| 29046c2 | .github/workflows: Add comment for dependabot to pick up upd | Shubham Bhardwaj | [#1502](https://github.com/tektoncd/chains/pull/1502) | — | — |
| b9d06f1 | feat(oci): support insecure OCI registry (#1374) | l-qing | [#1374](https://github.com/tektoncd/chains/pull/1374) | — | — |
| f07f988 | update golangci-lint version that supports go1.25 (#1508) | Anitha Natarajan | [#1508](https://github.com/tektoncd/chains/pull/1508) | — | — |
| 8bd204d | Fix golangci-lint action step for large diff (#1510) | Anitha Natarajan | [#1510](https://github.com/tektoncd/chains/pull/1510) | — | — |
| 0e98b57 | bump fulcio and other dependencies (#1506) | khelil | [#1506](https://github.com/tektoncd/chains/pull/1506) | — | — |
| 09000a3 | test: add unit tests for pubsub backend constructor and kafk | Naomi Gelman | [#1511](https://github.com/tektoncd/chains/pull/1511) | — | — |
| 99e2f1c | fix: update slash workflow to refer latest version that incl | Anitha Natarajan | [#1522](https://github.com/tektoncd/chains/pull/1522) | — | — |
| 3c6b8dd | chore(ci): update cherry-pick workflow to fix multi-commit P | Vincent Demeester | [#1539](https://github.com/tektoncd/chains/pull/1539) | — | — |
| 60afd54 | add Vincent as approver (#1546) | Anitha Natarajan | [#1546](https://github.com/tektoncd/chains/pull/1546) | — | — |
| 202f8e6 | Fix- Update Docdb storage logic (issue #1178) (#1505) | Naomi Gelman | [#1505](https://github.com/tektoncd/chains/pull/1505) | [SRVKP-10473](https://redhat.atlassian.net/browse/SRVKP-10473) | **Bug Fix:** The DocDB backend watcher now correctly detects filesystem changes under MongoSe *(extracted)* |
| 43a8448 | chore(docs update): update release cheatsheet, pipeline and  | Anitha Natarajan | [#1549](https://github.com/tektoncd/chains/pull/1549) | [SRVKP-10747](https://redhat.atlassian.net/browse/SRVKP-10747) | **Release Notes Not Required:** Internal release documentation and pipeline configuration updates for the Chains *(generated)* |
| 43281fc | fix: doc typo (#1561) | Naomi Gelman | [#1561](https://github.com/tektoncd/chains/pull/1561) | — | — |
| 9086de7 | ci: Add CI summary fan-in check (#1566) | Vincent Demeester | [#1566](https://github.com/tektoncd/chains/pull/1566) | — | — |
| df1288b | Move inactive approvers to emeritus (#1576) | Vincent Demeester | [#1576](https://github.com/tektoncd/chains/pull/1576) | — | — |
| eff5aad | docs: Add MongoDB storage tutorial (#1557) | Naomi Gelman | [#1557](https://github.com/tektoncd/chains/pull/1557) | — | — |
| 53834cb | feat(metrics): Migrate from OpenCensus to OpenTelemetry (#15 | Anitha Natarajan | [#1550](https://github.com/tektoncd/chains/pull/1550) | [SRVKP-9322](https://redhat.atlassian.net/browse/SRVKP-9322) | **Feature:** Chains metrics implementation migrates from the deprecated OpenCensus library to *(extracted)* |
| 31743e6 | Bump knative.dev/pkg to adopt centralized WEBHOOK_* TLS conf | khelil | [#1602](https://github.com/tektoncd/chains/pull/1602) | [SRVKP-9322](https://redhat.atlassian.net/browse/SRVKP-9322) *(keyword)* | **Enhancement:** Chains adopts centralized WEBHOOK_* TLS configuration from knative.dev/pkg, enab *(generated)* |
| b3f28ec | Ci/zizmor security fixes (#1612) | Abhishek Ghosh | [#1612](https://github.com/tektoncd/chains/pull/1612) | — | — |
| fc64ff4 | Nominate Emil and Abhishek are reviewers for chains repo (#1 | Anitha Natarajan | [#1621](https://github.com/tektoncd/chains/pull/1621) | — | — |
| f1137de | Update example performance doc (#1636) | Emil Natan | [#1636](https://github.com/tektoncd/chains/pull/1636) | — | — |
| 7290445 | Handle signing OCI artifacts in *ARTIFACT_OUTPUTS (#1578) | Brad Beck | [#1578](https://github.com/tektoncd/chains/pull/1578) | — | — |
| 51dd101 | Enable Server-Side Apply for finalizers management (#1635) | Emil Natan | [#1635](https://github.com/tektoncd/chains/pull/1635) | [SRVKP-8542](https://redhat.atlassian.net/browse/SRVKP-8542) | **Enhancement:** Chains now uses Kubernetes Server-Side Apply for managing finalizers on TaskRuns *(generated)* |
| 8182379 | chore: update doc for buildType (#1647) | Anitha Natarajan | [#1647](https://github.com/tektoncd/chains/pull/1647) | — | — |
| 12e6db0 | chore: bump tektoncd/plumbing cherry-pick workflow (#1673) | Abhishek Ghosh | [#1673](https://github.com/tektoncd/chains/pull/1673) | — | — |
| 2e5f961 | fix: add GHA based nightly workflow for chains (#1634) | Anitha Natarajan | [#1634](https://github.com/tektoncd/chains/pull/1634) | — | — |
| 6db27b1 | chore: update CI e2e matrix for k8s and pipelines versions ( | Abhishek Ghosh | [#1674](https://github.com/tektoncd/chains/pull/1674) | — | — |
| 4f9ad41 | chore: update chains maintainers (#1679) | Anitha Natarajan | [#1679](https://github.com/tektoncd/chains/pull/1679) | — | — |
| 4bb6198 | tekton: automate releases with Pipelines-as-Code (#1656) | Abhishek Ghosh | [#1656](https://github.com/tektoncd/chains/pull/1656) | [SRVKP-11937](https://redhat.atlassian.net/browse/SRVKP-11937) | **Release Notes Not Required:** Chains releases are now automated using Pipelines-as-Code, replacing manual tkn  *(generated)* |
| 3fbaa94 | Fix duplicate .att/.sig OCI layers for same-digest type hint | Abhishek Ghosh | [#1601](https://github.com/tektoncd/chains/pull/1601) | [SRVKP-11363](https://redhat.atlassian.net/browse/SRVKP-11363) | **Bug Fix:** Chains no longer produces duplicate .att and .sig OCI layers when a Task declare *(extracted)* |
| 020a29a | Add migration cleanup for SSA finalizers (#1699) | tekton-robot | [#1699](https://github.com/tektoncd/chains/pull/1699) | [SRVKP-12226](https://redhat.atlassian.net/browse/SRVKP-12226) *(keyword)* | **Bug Fix:** Adds migration cleanup logic to remove stale finalizers from in-flight objects c *(generated)* |


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-11484](https://redhat.atlassian.net/browse/SRVKP-11484) | Embed Rekor bundle in OCI attestation layer annotations when transparency is enabled | Code Review |


</details>

---


<details>
<summary><h2>tektoncd-cli</h2> — 71 commits, 7 linked, 52 dependabot</summary>


**Upstream:** tektoncd/cli · **Commits:** 71 · **Unsynced:** 0 · **Dependabot:** 52

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| ed60a80 | chore: referencing the setup-kind script in the plumbing rep | Shubham Bhardwaj | [#2590](https://github.com/tektoncd/cli/pull/2590) | — | — |
| 6c7d68a | Nominate divyansh42 as cli approver | Vincent Demeester | [#2761](https://github.com/tektoncd/cli/pull/2761) | — | — |
| f588be9 | Move piyush-garg to alumni | Vincent Demeester | [#2760](https://github.com/tektoncd/cli/pull/2760) | — | — |
| 698d2ea | update goreleaser version to compatible version | Shiv Verma | [#2748](https://github.com/tektoncd/cli/pull/2748) | — | — |
| 57f5665 | Update CLI docs for v0.44.0 release | Shiv Verma | [#2748](https://github.com/tektoncd/cli/pull/2748) | — | — |
| bde8531 | Add CI summary fan-in job to presubmit CI | Vincent Demeester | [#2741](https://github.com/tektoncd/cli/pull/2741) | — | — |
| 536f99c | feat: add cherry-pick command workflow | Vincent Demeester | [#2682](https://github.com/tektoncd/cli/pull/2682) | — | — |
| 7843065 | Change all occurences of GCS buckets with OCI buckets | adityavshinde | [#2768](https://github.com/tektoncd/cli/pull/2768) | [SRVKP-11090](https://redhat.atlassian.net/browse/SRVKP-11090) | **Release Notes Not Required** *(generated)* |
| 025ac7b | Add long flag --display-name to display the log of pipeliner | changjun | [#2450](https://github.com/tektoncd/cli/pull/2450) | — | — |
| f40a19b | refactor(logs): 将display-name参数重命名为long | changjun | [#2450](https://github.com/tektoncd/cli/pull/2450) | — | — |
| f0807af | feat: add describe subcommand in customrun command | Shiv Verma | [#2712](https://github.com/tektoncd/cli/pull/2712) | [SRVKP-11855](https://redhat.atlassian.net/browse/SRVKP-11855) *(keyword)* | **Feature:** The tkn CLI now supports a describe subcommand for CustomRuns, allowing users to *(extracted)* |
| 3d05ba3 | Remove all the occurences of chains subcommand | adityavshinde | [#2769](https://github.com/tektoncd/cli/pull/2769) | [SRVKP-11155](https://redhat.atlassian.net/browse/SRVKP-11155) | **Removed Functionality:** The chains subcommand has been removed from the tkn CLI. Users should use Tekton *(generated)* |
| 4a210b4 | Update knative and components version | Khurram Baig | [#2788](https://github.com/tektoncd/cli/pull/2788) | [SRVKP-11481](https://redhat.atlassian.net/browse/SRVKP-11481) | **Release Notes Not Required** *(generated)* |
| 5b4413c | Remove Metrics Exporter Logger due to OTEL migration | Khurram Baig | [#2788](https://github.com/tektoncd/cli/pull/2788) | [SRVKP-11481](https://redhat.atlassian.net/browse/SRVKP-11481) | **Release Notes Not Required** *(generated)* |
| 6a29abb | Migrate tracing from OpenCensus to OpenTelemetry | Khurram Baig | [#2799](https://github.com/tektoncd/cli/pull/2799) | [SRVKP-10252](https://redhat.atlassian.net/browse/SRVKP-10252) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 5f37794 | (deps) Bump go version to 1.25.8 to fix CVE-2026-25679 | divyansh42 | [#2803](https://github.com/tektoncd/cli/pull/2803) | — | — |
| 960f517 | Strip Go symbol table from release binaries | alliasgher | [#2843](https://github.com/tektoncd/cli/pull/2843) | — | — |
| 00c055a | Update tekton hub to 1.24.0 | divyansh42 | [#2845](https://github.com/tektoncd/cli/pull/2845) | [SRVKP-11950](https://redhat.atlassian.net/browse/SRVKP-11950) *(keyword)* | **Deprecated Functionality:** Tekton Hub support in the tkn CLI is deprecated in favor of Artifact Hub. Severa *(extracted)* |
| ce6edf3 | New version v0.45.0 | Shiv Verma | — | — | — |


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-11433](https://redhat.atlassian.net/browse/SRVKP-11433) | [Multi Cluster] Displaying blank status for Pending pipelineRuns across CLI & UI | Tasking and Estimation |


</details>

---


<details>
<summary><h2>tektoncd-pipeline</h2> — 327 commits, 28 linked, 173 dependabot</summary>


**Upstream:** tektoncd/pipeline · **Commits:** 327 · **Unsynced:** 0 · **Dependabot:** 173

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 12c973b | build(deps): bump sigstore/sigstore from 1.9.5 to 1.10.4 | Vincent Demeester | [#9331](https://github.com/tektoncd/pipeline/pull/9331) | — | — |
| c68c8d0 | build(deps): bump github.com/tektoncd/pipeline to v1.7.0 in  | Vincent Demeester | [#9329](https://github.com/tektoncd/pipeline/pull/9329) | — | — |
| f001341 | Fix wait-task-beta test infrastructure for k8s compatibility | Vincent Demeester | [#9329](https://github.com/tektoncd/pipeline/pull/9329) | — | — |
| b49da8a | Set KUBERNETES_MIN_VERSION for wait-task-beta controller | Vincent Demeester | [#9329](https://github.com/tektoncd/pipeline/pull/9329) | — | — |
| 7783005 | docs: update releases.md for v1.9.0 LTS | Vincent Demeester | [#9339](https://github.com/tektoncd/pipeline/pull/9339) | — | — |
| d9df4b3 | docs: Document roadmap project board workflows and best prac | Vincent Demeester | [#9311](https://github.com/tektoncd/pipeline/pull/9311) | — | — |
| 05b05ba | chore: remove redundant word | wmypku | [#9009](https://github.com/tektoncd/pipeline/pull/9009) | — | — |
| 6cd9158 | Update docs for changes in apiVersion v1 | Johan Kok | [#9042](https://github.com/tektoncd/pipeline/pull/9042) | — | — |
| 37afb99 | build(deps): bump opentelemetry exporter packages to v1.39.0 | Vincent Demeester | [#9332](https://github.com/tektoncd/pipeline/pull/9332) | — | — |
| 6ae1296 | fix(ci): simplify e2e test health status result | Anitha Natarajan | [#9361](https://github.com/tektoncd/pipeline/pull/9361) | — | — |
| abfa29d | fix: Align cache configstore with framework implementation | Seunghun Shin | [#9282](https://github.com/tektoncd/pipeline/pull/9282) | — | — |
| 6206ff7 | ci: add /rebase slash command workflow | Vibhav Bobade | [#9375](https://github.com/tektoncd/pipeline/pull/9375) | — | — |
| d52f146 | fix(pipelines): allow pipeline param defaults to use non-par | Andrew Thorp | [#9386](https://github.com/tektoncd/pipeline/pull/9386) | [SRVKP-10840](https://redhat.atlassian.net/browse/SRVKP-10840) | **Bug Fix:** Fixed a bug which caused PipelineRun validation to fail when a pipeline paramete *(extracted)* |
| 2df9fbe | feat: Add SHA-256 support for Git resolver revision validati | 7h3-3mp7y-m4n | [#9278](https://github.com/tektoncd/pipeline/pull/9278) | — | — |
| 579ccaf | ci: replace e2e-only fan-in with unified CI summary job | Vincent Demeester | [#9394](https://github.com/tektoncd/pipeline/pull/9394) | — | — |
| 31a2e31 | tekton: update plumbing ref to include full image references | Vincent Demeester | [#9399](https://github.com/tektoncd/pipeline/pull/9399) | — | — |
| 5d50dad | ci: Update cherry-pick command to latest plumbing | Vincent Demeester | [#9400](https://github.com/tektoncd/pipeline/pull/9400) | — | — |
| 0c3608f | fix: Remove redundant shortNames from ResolutionRequest CRD | Vincent Demeester | [#9398](https://github.com/tektoncd/pipeline/pull/9398) | — | — |
| c0bef5f | tekton: update plumbing ref to latest commit | Vincent Demeester | [#9406](https://github.com/tektoncd/pipeline/pull/9406) | — | — |
| 5d0e97e | docs: clarify flag availability across controller binaries | Lucas Aguiar | [#9390](https://github.com/tektoncd/pipeline/pull/9390) | — | — |
| 4d10d23 | Assess several new gosec findings | Sascha Schwarze | [#9405](https://github.com/tektoncd/pipeline/pull/9405) | — | — |
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
| b0a3970 | Enable Gitea e2e tests in CI | ab-ghosh | [#9442](https://github.com/tektoncd/pipeline/pull/9442) | [SRVKP-10929](https://redhat.atlassian.net/browse/SRVKP-10929) | **Release Notes Not Required** *(extracted)* |
| 4648b3e | Update release-cheat-sheet.md and release pipeline | Naomi Gelman | [#9415](https://github.com/tektoncd/pipeline/pull/9415) | — | — |
| cd0046c | fix: revert mistaken metadata changes in resolvers config-ob | Khurram Baig | [#9468](https://github.com/tektoncd/pipeline/pull/9468) | — | — |
| e712e56 | chore: bump knative.dev/pkg to main and k8s libs to 0.35.1 | Vincent Demeester | [#9470](https://github.com/tektoncd/pipeline/pull/9470) | [SRVKP-9318](https://redhat.atlassian.net/browse/SRVKP-9318) *(unmatched)* | — |
| 9ad8e4b | chore: regenerate code after knative/pkg and k8s bump | Vincent Demeester | [#9470](https://github.com/tektoncd/pipeline/pull/9470) | [SRVKP-9318](https://redhat.atlassian.net/browse/SRVKP-9318) *(unmatched)* | — |
| 520a2ff | docs: replace 'coming soon' with tkn bundle link in taskruns | Chinonso Nwakudu | [#9509](https://github.com/tektoncd/pipeline/pull/9509) | — | — |
| d1379f8 | fix: propagate PipelineRun tasks/finally timeout to child Ta | Shubham Bhardwaj | [#9419](https://github.com/tektoncd/pipeline/pull/9419) | [SRVKP-8963](https://redhat.atlassian.net/browse/SRVKP-8963) | **Bug Fix:** PipelineRun tasks/finally timeouts are now correctly propagated to child TaskRun *(generated)* |
| 5c74658 | Move inactive approvers to alumni | Vincent Demeester | [#9518](https://github.com/tektoncd/pipeline/pull/9518) | — | — |
| bf12c8d | Nominate khrm and aThorp96 as pipeline approvers | Vincent Demeester | [#9519](https://github.com/tektoncd/pipeline/pull/9519) | — | — |
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
| 961388f | fix: prevent path traversal in git resolver pathInRepo param | Vincent Demeester | — | — | — |
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
| e8f7a28 | docs: add 4 undocumented metrics to docs/metrics.md | Chinonso Nwakudu | [#9512](https://github.com/tektoncd/pipeline/pull/9512) | — | — |
| c408cbd | docs: fix broken internal markdown links | Goutham-AR | [#9507](https://github.com/tektoncd/pipeline/pull/9507) | — | — |
| 97c62a5 | doc: Fix broken Tekton Bundles example link in taskruns.md | sahilleth | [#9462](https://github.com/tektoncd/pipeline/pull/9462) | — | — |
| 64a339b | doc: Clarify scope of auth documentation | sahilleth | [#9461](https://github.com/tektoncd/pipeline/pull/9461) | — | — |
| 838e536 | Fix: Add useHttpPath to support multiple Git credentials on  | ab-ghosh | [#9143](https://github.com/tektoncd/pipeline/pull/9143) | [SRVKP-9280](https://redhat.atlassian.net/browse/SRVKP-9280) | **Bug Fix:** Tekton now supports multiple Git credentials on the same host by using useHttpPa *(generated)* |
| 0bfd458 | docs: add README files for pipelinerun and taskrun examples | anirudh242 | [#9505](https://github.com/tektoncd/pipeline/pull/9505) | — | — |
| 1ea381e | docs: remove file listings and reorder README sections based | anirudh242 | [#9505](https://github.com/tektoncd/pipeline/pull/9505) | — | — |
| 418dd19 | fix: Upgrade Gitea test infrastructure from v1.17.1 to lates | Sri Vignesh | [#9568](https://github.com/tektoncd/pipeline/pull/9568) | — | — |
| 5f15773 | fix: resolve cache race with singleflight | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) | [SRVKP-10234](https://redhat.atlassian.net/browse/SRVKP-10234) *(keyword)* | **Bug Fix:** Concurrent resolver cache requests for the same resource are now deduplicated us *(generated)* |
| 5df99a0 | fix: remove corrupted cache entries on type error | Stanislav Jakuschevskij | [#9365](https://github.com/tektoncd/pipeline/pull/9365) | [SRVKP-10234](https://redhat.atlassian.net/browse/SRVKP-10234) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 546de63 | Bump knative.dev/pkg to adopt centralized WEBHOOK_* TLS conf | Jawed khelil | [#9466](https://github.com/tektoncd/pipeline/pull/9466) | [SRVKP-9614](https://redhat.atlassian.net/browse/SRVKP-9614) | **Feature:** The Tekton webhook now supports centrally managed TLS configuration via WEBHOOK_ *(generated)* |
| 24766ef | fix: record metrics for cancelled PipelineRuns | Khurram Baig | [#9658](https://github.com/tektoncd/pipeline/pull/9658) | [SRVKP-11099](https://redhat.atlassian.net/browse/SRVKP-11099) | **Bug Fix:** Cancelled PipelineRuns are now correctly recorded in the tekton_pipelines_contro *(extracted)* |
| 7e3e7cc | Update stale comment about storing TaskSpec in status | Andrea Frittoli | [#9661](https://github.com/tektoncd/pipeline/pull/9661) | — | — |
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
| efca098 | ci: remove compromised tj-actions/changed-files dependency | Vincent Demeester | [#9713](https://github.com/tektoncd/pipeline/pull/9713) | — | — |
| 61e0304 | docs: update releases.md for v1.11.0 | Vincent Demeester | [#9685](https://github.com/tektoncd/pipeline/pull/9685) | — | — |
| 7ea863e | fix: pin registry image and relax log-based cache assertion | Vincent Demeester | [#9761](https://github.com/tektoncd/pipeline/pull/9761) | — | — |
| 721cc83 | fix: surface specific TaskRun failure reasons for pod evicti | Vibhav Bobade | [#9368](https://github.com/tektoncd/pipeline/pull/9368) | — | — |
| 1d50b07 | fix: surface clear errors when completed tasks miss referenc | Vibhav Bobade | [#9529](https://github.com/tektoncd/pipeline/pull/9529) | — | — |
| ffd686e | Rework the events controller cache | Andrea Frittoli | [#6914](https://github.com/tektoncd/pipeline/pull/6914) | — | — |
| be4a4bc | docs: update metrics.md to reflect OpenTelemetry migration | Shubham Bhardwaj | [#9641](https://github.com/tektoncd/pipeline/pull/9641) | — | — |
| 5b6ec37 | [TEP-0137] Add TaskRun notifications controller | Andrea Frittoli | [#9674](https://github.com/tektoncd/pipeline/pull/9674) | — | — |
| 8546d5f | Remove aws-sdk-go-v2 ECR replace directives | Vincent Demeester | [#9773](https://github.com/tektoncd/pipeline/pull/9773) | — | — |
| 8e53012 | docs(examples): remove stale v1beta1 references from example | Sakthi Harish | [#9672](https://github.com/tektoncd/pipeline/pull/9672) | — | — |
| 9f9db26 | test(e2e): fix flaky matrix and retry example tests | Ogulcan Aydogan | [#9655](https://github.com/tektoncd/pipeline/pull/9655) | — | — |
| 770a6f7 | perf: use maps.Equal instead of reflect.DeepEqual for label/ | Ogulcan Aydogan | [#9778](https://github.com/tektoncd/pipeline/pull/9778) | — | — |
| 7e12074 | fix(e2e): replace ubuntu with busybox in dind-sidecar test D | Vibhav Bobade | [#9806](https://github.com/tektoncd/pipeline/pull/9806) | — | — |
| 5eeec51 | tekton: add draft release creation to release pipeline | Vincent Demeester | [#9420](https://github.com/tektoncd/pipeline/pull/9420) | — | — |
| 229ed41 | tekton: address review feedback on draft release tasks | Vincent Demeester | [#9420](https://github.com/tektoncd/pipeline/pull/9420) | — | — |
| 21a9083 | Regenerate code after k8s.io/code-generator 0.35.3 bump | Vincent Demeester | [#9644](https://github.com/tektoncd/pipeline/pull/9644) | — | — |
| 3fd3dac | [TEP-0137] Add PipelineRun notifications controller | Andrea Frittoli | [#9677](https://github.com/tektoncd/pipeline/pull/9677) | — | — |
| 6287db3 | Simplify TestEmit by removing table-driven test structure | Andrea Frittoli | [#9677](https://github.com/tektoncd/pipeline/pull/9677) | — | — |
| e964bce | Security: reject system API token with user-controlled serve | Vincent Demeester | — | — | — |
| 0967cc8 | fix: prevent git argument injection via revision parameter ( | Vincent Demeester | — | — | — |
| c8cc73a | fix: normalize VolumeMount paths before /tekton/ restriction | Vincent Demeester | — | — | — |
| b890560 | fix: strip resolver prefixes and use non-capturing group for | Vincent Demeester | — | — | — |
| d4ba8ca | fix: limit HTTP resolver response body size to prevent OOM D | Vincent Demeester | — | — | — |
| 37f9657 | fix: trim whitespace from source URI before pattern matching | Vincent Demeester | — | — | — |
| 033ce8e | build(deps): bump k8s.io dependencies from 0.35.2 to 0.35.4 | Vincent Demeester | [#9848](https://github.com/tektoncd/pipeline/pull/9848) | — | — |
| cadf9f7 | docs: add bundle resolver configuration options default valu | Gonçalo Marques | [#9772](https://github.com/tektoncd/pipeline/pull/9772) | — | — |
| 85ff854 | ci: Automate Dependabot configuration generation | Vincent Demeester | [#9188](https://github.com/tektoncd/pipeline/pull/9188) | — | — |
| 42b30d4 | fix: address zizmor findings in dependabot-regen workflow | Vincent Demeester | [#9188](https://github.com/tektoncd/pipeline/pull/9188) | — | — |
| 0b4440d | [TEP-0137] Deprecate send-cloudevents-for-runs feature flag | Andrea Frittoli | [#9774](https://github.com/tektoncd/pipeline/pull/9774) | — | — |
| fec1273 | perf: reduce reconcile churn for completed PipelineRuns | Vincent Demeester | [#9706](https://github.com/tektoncd/pipeline/pull/9706) | — | — |
| 8709b51 | perf: remove unnecessary SetDefaults from TaskRun done path | Vincent Demeester | [#9706](https://github.com/tektoncd/pipeline/pull/9706) | — | — |
| d635dcb | test: add e2e test for TaskRun pending status | Vibhav Bobade | [#9681](https://github.com/tektoncd/pipeline/pull/9681) | — | — |
| 0301423 | fix(resolvers): validate data is Tekton object in resolver f | Andrew Thorp | [#9963](https://github.com/tektoncd/pipeline/pull/9963) | — | — |
| 513e0e0 | Fix gen-crd-api-reference-docs require to use fetchable vers | Vincent Demeester | [#10001](https://github.com/tektoncd/pipeline/pull/10001) | — | — |
| b150ab2 | fix(pipelinerun): use generateName for anonymous pipeline la | Andrew Thorp | [#10079](https://github.com/tektoncd/pipeline/pull/10079) | — | — |
| 9c0a33a | fix: truncate affinity assistant volume names to 63 characte | Vincent Demeester | [#10137](https://github.com/tektoncd/pipeline/pull/10137) | — | — |
| 803d921 | Fix PipelineRun premature failure when TaskRun recovers afte | ab-ghosh | [#10161](https://github.com/tektoncd/pipeline/pull/10161) | — | — |
| a7e231d | fix: prevent controller CPU variant from leaking into platfo | Vincent Demeester | [#10164](https://github.com/tektoncd/pipeline/pull/10164) | — | — |
| 7121dd9 | fix: TaskRun stuck in Running when init container is OOMKill | Vincent Demeester | [#10186](https://github.com/tektoncd/pipeline/pull/10186) | — | — |
| 64ec216 | fix: add automated draft release support to release pipeline | Vincent Demeester | [#10216](https://github.com/tektoncd/pipeline/pull/10216) | — | — |
| 829429b | fix(resolvers): Allow ResolutionRequests to resolve all Tekt | Andrew Thorp | [#10252](https://github.com/tektoncd/pipeline/pull/10252) | — | — |
| 03e3bae | fix: bump Go for CVEs | Vibhav Bobade | [#10340](https://github.com/tektoncd/pipeline/pull/10340) | — | — |


### Unmatched Jiras

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-11432](https://redhat.atlassian.net/browse/SRVKP-11432) | PVC auto-cleanup annotation propagates to child PipelineRuns in Pipeline-in-Pipeline | Code Review |
| [SRVKP-11199](https://redhat.atlassian.net/browse/SRVKP-11199) | MultiCluster \| Unable to Cancel/Stop a pipeline run  | Closed |
| [SRVKP-10736](https://redhat.atlassian.net/browse/SRVKP-10736) | Reproduce this number of PipelineRuns count mismatch using Kind | Dev Complete |
| [SRVKP-10630](https://redhat.atlassian.net/browse/SRVKP-10630) | Test Automation - Add Release Tests | In Progress |
| [SRVKP-10457](https://redhat.atlassian.net/browse/SRVKP-10457) | PipelineRun intermittently fails and skips downstream tasks while a TaskRun is still Running under high concurrency | Closed |
| [SRVKP-10262](https://redhat.atlassian.net/browse/SRVKP-10262) | Define project/release resource to tie up related PipelineRuns up | Refinement |
| [SRVKP-10066](https://redhat.atlassian.net/browse/SRVKP-10066) | Add namespace label to `tekton_pipelines_controller_pipelinerun_total` metric | In Progress |
| [SRVKP-10038](https://redhat.atlassian.net/browse/SRVKP-10038) | Logs should contain pipeline or task names  | Dev Complete |
| [SRVKP-1801](https://redhat.atlassian.net/browse/SRVKP-1801) | Pipelines don't work with resource quota | Dev Complete |


</details>

---


<details>
<summary><h2>tektoncd-triggers</h2> — 60 commits, 11 linked, 37 dependabot</summary>


**Upstream:** tektoncd/triggers · **Commits:** 60 · **Unsynced:** 0 · **Dependabot:** 37

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 6ac1c13 | Change precheck in release pipeline to OCI infra | Khurram Baig | [#1947](https://github.com/tektoncd/triggers/pull/1947) | [SRVKP-9533](https://redhat.atlassian.net/browse/SRVKP-9533) *(keyword)* | **Release Notes Not Required** *(generated)* |
| f61d6b1 | Update releases.md for v0.35.0 | Khurram Baig | [#1952](https://github.com/tektoncd/triggers/pull/1952) | — | — |
| 5501e44 | Consolidate CI workflows for build, lint, and e2e tests | Shubham Bhardwaj | [#1957](https://github.com/tektoncd/triggers/pull/1957) | — | — |
| bb8ddfc | Update Release Cheat Sheet for release-draft-oci pipeline | Khurram Baig | [#1948](https://github.com/tektoncd/triggers/pull/1948) | [SRVKP-9533](https://redhat.atlassian.net/browse/SRVKP-9533) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 88493b0 | fix: doc typos | Naomi Gelman | [#1953](https://github.com/tektoncd/triggers/pull/1953) | — | — |
| 7f047e1 | Bump go version to 1.25.6 and golang ci version to 2.8.0 | Shubham Bhardwaj | [#1966](https://github.com/tektoncd/triggers/pull/1966) | — | — |
| d73587e | Move inactive approvers to emeritus | Vincent Demeester | [#1965](https://github.com/tektoncd/triggers/pull/1965) | — | — |
| 5f58ccb | Fix e2e failure due to owners file change | Khurram Baig | [#1970](https://github.com/tektoncd/triggers/pull/1970) | — | — |
| e8b2f37 | cleanup: replace GCS release URLs with infra.tekton.dev | Ankur Sinha | [#1973](https://github.com/tektoncd/triggers/pull/1973) | [SRVKP-11096](https://redhat.atlassian.net/browse/SRVKP-11096) | **Release Note Not Required** *(generated)* |
| 009871a | Fix port conflict between event-listener and knative health  | Shubham Bhardwaj | [#1934](https://github.com/tektoncd/triggers/pull/1934) | [SRVKP-9324](https://redhat.atlassian.net/browse/SRVKP-9324) | **Bug Fix:** Fixes a port conflict between the event-listener and Knative health probes that  *(generated)* |
| b6eed12 | Bump Knative and Pipeline to latest | Shubham Bhardwaj | [#1934](https://github.com/tektoncd/triggers/pull/1934) | [SRVKP-9324](https://redhat.atlassian.net/browse/SRVKP-9324) | **Enhancement:** Updates Knative and Pipeline dependencies to latest versions to support the Open *(generated)* |
| 3ff1d1d | Migrate from opencensus to Opentelemetry | Shubham Bhardwaj | [#1934](https://github.com/tektoncd/triggers/pull/1934) | [SRVKP-9324](https://redhat.atlassian.net/browse/SRVKP-9324) | **Feature:** Triggers metrics implementation is migrated from the deprecated OpenCensus libra *(extracted)* |
| b293c22 | Update observability docs and config for OpenTelemetry migra | Shubham Bhardwaj | [#1934](https://github.com/tektoncd/triggers/pull/1934) | [SRVKP-9324](https://redhat.atlassian.net/browse/SRVKP-9324) | **Enhancement:** Updates observability documentation and configuration to reflect the OpenTelemet *(generated)* |
| 04257df | Bump the all group across 1 directory with 7 updates | Khurram Baig | [#1985](https://github.com/tektoncd/triggers/pull/1985) | — | — |
| f00609d | Updated generated files for newer k8s version | Khurram Baig | [#1985](https://github.com/tektoncd/triggers/pull/1985) | — | — |
| 1aad8a1 | Bump tektoncd/pipeline to v1.11.0 and update vendor dependen | Khurram Baig | [#1986](https://github.com/tektoncd/triggers/pull/1986) | — | — |
| 0c827bc | Update release ko image to go1.25 | Khurram Baig | [#1997](https://github.com/tektoncd/triggers/pull/1997) | [SRVKP-2980](https://redhat.atlassian.net/browse/SRVKP-2980) | **Release Notes Not Required:** NA *(generated)* |
| 117a7d9 | ci: Fix zizmor security findings in GitHub Actions | Shubham Bhardwaj | [#1998](https://github.com/tektoncd/triggers/pull/1998) | — | — |
| 9f011db | Fix intermittent panic in Test_UpdateCACertToClusterIntercep | Khurram Baig | [#2000](https://github.com/tektoncd/triggers/pull/2000) | — | — |
| 6e57af9 | Fix curl command to follow redirects for release file | Khurram | [#2017](https://github.com/tektoncd/triggers/pull/2017) | [SRVKP-9533](https://redhat.atlassian.net/browse/SRVKP-9533) *(keyword)* | **Release Notes Not Required** *(generated)* |
| bf0a492 | Change release pipeline to use 'release-draft-oci' | Khurram | [#2018](https://github.com/tektoncd/triggers/pull/2018) | [SRVKP-9533](https://redhat.atlassian.net/browse/SRVKP-9533) *(keyword)* | **Release Notes Not Required** *(generated)* |
| 1d36b80 | Bump the pipelines dependency and other deps | Khurram Baig | [#2020](https://github.com/tektoncd/triggers/pull/2020) | — | — |
| bf95a60 | Add TLS security profile support for core interceptors | Jawed khelil | [#2019](https://github.com/tektoncd/triggers/pull/2019) | [SRVKP-12012](https://redhat.atlassian.net/browse/SRVKP-12012) | **Feature:** Core interceptors now honor the TLS security profile injected by the Tekton oper *(extracted)* |


</details>

---


<details>
<summary><h2>syncer-service</h2> — 1 commits, 0 linked</summary>


**Upstream:** openshift-pipelines/syncer-service · **Commits:** 1 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 30b70d1 | update pipeline version to v1.9.3 | Pramod Bindal | [#71](https://github.com/openshift-pipelines/syncer-service/pull/71) | — | — |


</details>

---


<details>
<summary><h2>tekton-assist</h2> — 0 commits, 0 linked</summary>


**Error:** failed to get head SHA at from-date: failed to list head file commits for openshift-pipelines/p12n-tekton-assist: GET https://api.github.com/repos/openshift-pipelines/p12n-tekton-assist/commits?path=head&per_page=1&sha=release-v1.22.x&until=2026-05-15T07%3A57%3A06Z: 404 Not Found []


</details>

---


<details>
<summary><h2>tekton-caches</h2> — 4 commits, 4 linked</summary>


**Upstream:** openshift-pipelines/tekton-caches · **Commits:** 4 · **Unsynced:** 0

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| 2450f24 | Security: Fix CVE-2026-34986 (go-jose/go-jose/v4) | Pramod Bindal | [#729](https://github.com/openshift-pipelines/tekton-caches/pull/729) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Fixes CVE-2026-34986 by upgrading go-jose/go-jose/v4 to v4.1.4, addressing a vul *(generated)* |
| 2a64502 | bump otel and pipeline version for CVE fixes | Pramod Bindal | [#736](https://github.com/openshift-pipelines/tekton-caches/pull/736) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Updates OpenTelemetry and Tekton Pipeline dependency versions to incorporate ups *(generated)* |
| 5d6871a | cve: Fix GHSA-xmrv-pmrh-hhx2 | Pramod Bindal | [#739](https://github.com/openshift-pipelines/tekton-caches/pull/739) | [SRVKP-12342](https://redhat.atlassian.net/browse/SRVKP-12342) *(keyword)* | **CVE:** Fixes GHSA-xmrv-pmrh-hhx2 by upgrading github.com/aws/aws-sdk-go-v2 from v1.41.2 *(generated)* |
| 768eb9a | Add Agentic Contexts and CI Workflows | Pramod Bindal | [#739](https://github.com/openshift-pipelines/tekton-caches/pull/739) | [SRVKP-12024](https://redhat.atlassian.net/browse/SRVKP-12024) *(keyword)* | **Release Notes Not Required** *(generated)* |


</details>

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

Tickets with no Jira component label or no matching release component.

| Key | Summary | Status |
|-----|---------|--------|
| [SRVKP-12426](https://redhat.atlassian.net/browse/SRVKP-12426) | CLONE -  Upstream release notes for hub | New |
| [SRVKP-12425](https://redhat.atlassian.net/browse/SRVKP-12425) | CLONE -  Upstream release notes for Operator | New |
| [SRVKP-12422](https://redhat.atlassian.net/browse/SRVKP-12422) | CLONE -  Upstream release notes for Pipelines | New |
| [SRVKP-12418](https://redhat.atlassian.net/browse/SRVKP-12418) | Upstream release notes for 1.23 | New |
| [SRVKP-12416](https://redhat.atlassian.net/browse/SRVKP-12416) | Release Captain Work 1.23 | Dev Complete |
| [SRVKP-12185](https://redhat.atlassian.net/browse/SRVKP-12185) | In all triggers & sub pages, when searching by label, auto suggestions are not populated | Closed |
| [SRVKP-12184](https://redhat.atlassian.net/browse/SRVKP-12184) | In tasks/taskruns page, when searching by label, auto suggestions are not populated | Closed |
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
| [SRVKP-11640](https://redhat.atlassian.net/browse/SRVKP-11640) | OpenShift 5.0 Operator Compatibility Requirements for Pipelines | Refinement |
| [SRVKP-11451](https://redhat.atlassian.net/browse/SRVKP-11451) | Q2 OSP Roadmap workshop 2026 Presentation - Child 1 | Closed |
| [SRVKP-11445](https://redhat.atlassian.net/browse/SRVKP-11445) | Openshift Pipelines Release 1.23.0 | Backlog |
| [SRVKP-11167](https://redhat.atlassian.net/browse/SRVKP-11167) | Create a custom filter component as the console sdk has not added @patternfly/react-data-view to shared modules in the pre-release version 4.22.0-prerelease.1 | Release Pending |
| [SRVKP-11136](https://redhat.atlassian.net/browse/SRVKP-11136) | Document Guidance on Scaling Tekton Controllers for High-Volume Workloads | Tasking and Estimation |
| [SRVKP-11123](https://redhat.atlassian.net/browse/SRVKP-11123) | CI Failure in React 18 Branch Build and updating dependencies in package.json | Release Pending |
| [SRVKP-11115](https://redhat.atlassian.net/browse/SRVKP-11115) | [QE - 1.23 ] Validate the 1.22 OSP image as 1.23 OSP for OCP 4.21 and below | Closed |
| [SRVKP-11098](https://redhat.atlassian.net/browse/SRVKP-11098) | R&D: Schedule PipelineRuns on Hub and Spoke clusters based on availability using MultiKueue | Closed |
| [SRVKP-11051](https://redhat.atlassian.net/browse/SRVKP-11051) | Adopt Pipelines-as-Code into the Tekton Project | In Progress |
| [SRVKP-11023](https://redhat.atlassian.net/browse/SRVKP-11023) | Tasks for deprecating Tekton Hub | Backlog |
| [SRVKP-11003](https://redhat.atlassian.net/browse/SRVKP-11003) | console-plugin branch readiness for react 18 | Release Pending |
| [SRVKP-10964](https://redhat.atlassian.net/browse/SRVKP-10964) | Migrate CI Credentials to HashiCorp Vault Cloud | Closed |
| [SRVKP-10963](https://redhat.atlassian.net/browse/SRVKP-10963) | Migrate CI Credentials to HashiCorp Vault Cloud | Closed |
| [SRVKP-10962](https://redhat.atlassian.net/browse/SRVKP-10962) | Migrate CI Credentials to HashiCorp Vault Cloud | Closed |
| [SRVKP-10907](https://redhat.atlassian.net/browse/SRVKP-10907) | Test Openshift Pipelines on RHCOS 9 and 10 for OCP 4.22 | Closed |
| [SRVKP-10806](https://redhat.atlassian.net/browse/SRVKP-10806) | Add JAVA_VERSION parameter to Maven Task for Java version selection | Closed |
| [SRVKP-10574](https://redhat.atlassian.net/browse/SRVKP-10574) | [Pre-Stage] [release testing] Testing in disconnected env P | To Do |
| [SRVKP-10573](https://redhat.atlassian.net/browse/SRVKP-10573) | [Pre-Stage] [release testing] Cluster upgrade tests | To Do |
| [SRVKP-10572](https://redhat.atlassian.net/browse/SRVKP-10572) | [Pre-Stage] [release testing] update ci-config.yaml | Closed |
| [SRVKP-10571](https://redhat.atlassian.net/browse/SRVKP-10571) | [Pre-Stage] [release testing] Documentation review of Release Notes  | To Do |
| [SRVKP-10570](https://redhat.atlassian.net/browse/SRVKP-10570) | [Post-Release] [release testing] Multiarch testing - ARM64 | To Do |
| [SRVKP-10569](https://redhat.atlassian.net/browse/SRVKP-10569) | [Stage] [release testing] Operator Upgrade  | To Do |
| [SRVKP-10568](https://redhat.atlassian.net/browse/SRVKP-10568) | [Stage] [release testing] Acceptance Tests  | In Progress |
| [SRVKP-10567](https://redhat.atlassian.net/browse/SRVKP-10567) | [Pre-Stage] [release testing] TKN/OPC tests | To Do |
| [SRVKP-10566](https://redhat.atlassian.net/browse/SRVKP-10566) | [Post-Release] [release testing] Post-release verify operator installation | To Do |
| [SRVKP-10565](https://redhat.atlassian.net/browse/SRVKP-10565) | [Pre-Stage] [release testing] Testing in disconnected env x86_64 | To Do |
| [SRVKP-10564](https://redhat.atlassian.net/browse/SRVKP-10564) | [Pre-Stage] [release testing] create release branch in release-tests | Closed |
| [SRVKP-10563](https://redhat.atlassian.net/browse/SRVKP-10563) | [Pre-Stage] [release testing] Documentation review For Feature Testing | To Do |
| [SRVKP-10562](https://redhat.atlassian.net/browse/SRVKP-10562) | [Pre-Stage] [release testing] Devconsole UI manual tests | To Do |
| [SRVKP-10561](https://redhat.atlassian.net/browse/SRVKP-10561) | [Stage] [release testing] Verification of CLI binary SHA/signature on Mac and Windows | To Do |
| [SRVKP-10560](https://redhat.atlassian.net/browse/SRVKP-10560) | [Pre-Stage] [release testing] Feature testing | In Progress |
| [SRVKP-10559](https://redhat.atlassian.net/browse/SRVKP-10559) | [Pre-Stage] [release testing] Testing in disconnected env Z  | To Do |
| [SRVKP-10558](https://redhat.atlassian.net/browse/SRVKP-10558) | [Post-Release] [release testing] Multiarch testing - IBM Z | To Do |
| [SRVKP-10557](https://redhat.atlassian.net/browse/SRVKP-10557) | [Pre-Stage] [release testing] Operator upgrade tests | In Progress |
| [SRVKP-10556](https://redhat.atlassian.net/browse/SRVKP-10556) | [Pre-Stage] [release testing] Test on FIPS cluster | To Do |
| [SRVKP-10555](https://redhat.atlassian.net/browse/SRVKP-10555) | [Post release] [release testing] Verify binaries on mirror.openshift.com | To Do |
| [SRVKP-10554](https://redhat.atlassian.net/browse/SRVKP-10554) | [Pre-Stage] [release testing] Bug verification | In Progress |
| [SRVKP-10553](https://redhat.atlassian.net/browse/SRVKP-10553) | [Post-Release] [release testing] Multiarch testing - IBM Power | To Do |
| [SRVKP-10552](https://redhat.atlassian.net/browse/SRVKP-10552) | [Pre-Stage] [release testing] Hub manual tests | To Do |
| [SRVKP-10551](https://redhat.atlassian.net/browse/SRVKP-10551) | Test Openshift Pipelines 1.23.0 | In Progress |
| [SRVKP-10303](https://redhat.atlassian.net/browse/SRVKP-10303) | Migrate CI Credentials to HashiCorp Vault Cloud | Closed |
| [SRVKP-10296](https://redhat.atlassian.net/browse/SRVKP-10296) | Releasing RPMs via Konflux | Closed |
| [SRVKP-9722](https://redhat.atlassian.net/browse/SRVKP-9722) | Performance regression testing for Pipelines 1.23.0 | Tasking and Estimation |
| [SRVKP-9721](https://redhat.atlassian.net/browse/SRVKP-9721) | Release OpenShift Pipelines CLI for Pipelines 1.23.0 | Tasking and Estimation |
| [SRVKP-9720](https://redhat.atlassian.net/browse/SRVKP-9720) | Release OpenShift Pipelines Operator for Pipelines 1.23.0 | Tasking and Estimation |
| [SRVKP-9447](https://redhat.atlassian.net/browse/SRVKP-9447) | Migration Guide and Documentation | Tasking and Estimation |
| [SRVKP-9114](https://redhat.atlassian.net/browse/SRVKP-9114) | Testing for the epic | Closed |
| [SRVKP-8541](https://redhat.atlassian.net/browse/SRVKP-8541) | Enable SSA finalizer management Results | Code Review |

---
