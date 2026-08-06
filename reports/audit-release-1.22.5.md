# Release Audit: 1.22.5

**Source:** upstream commits after 2026-07-02 (1.22.4 release date)<br>
**Total commits:** 82 across 12 components (5 components with no changes)<br>
**Dependabot (skipped):** 30 · **Bot/merge commits:** 14 · **Code changes:** 38<br>

## Index

| Component | Commits | Code Changes | Dependabot | Category |
|-----------|---------|-------------|------------|----------|
| [tektoncd-pipeline](#tektoncd-pipeline) | 11 | 9 | 1 | Bug fixes (ResolutionRequest, resolvers) |
| [operator](#operator) | 19 | 10 | 5 | Maintenance (image digests, version bumps) |
| [console-plugin](#console-plugin) | 12 | 8 | 0 | CVE fixes, bug fix |
| [console-plugin-pf5](#console-plugin-pf5) | 12 | 8 | 0 | CVE fixes, bug fixes |
| [tektoncd-cli](#tektoncd-cli) | 18 | 3 | 14 | CVE fix, maintenance |
| [tektoncd-chains](#tektoncd-chains) | 12 | 1 | 11 | CVE fix (symlinks) |
| [tektoncd-results](#tektoncd-results) | 4 | 2 | 0 | CVE fixes (Go stdlib, x/text) |
| [manual-approval-gate](#manual-approval-gate) | 3 | 3 | 0 | CVE fixes |
| [tektoncd-pruner](#tektoncd-pruner) | 1 | 1 | 0 | Bug fix (webhook ConfigMap) |
| [pipelines-as-code](#pipelines-as-code) | 2 | 1 | 0 | Go version bump (v0.42.3) |
| [opc](#opc) | 2 | 2 | 0 | Version bump |
| [tekton-kueue](#tekton-kueue) | 3 | 0 | 3 | Dependency updates |
| tektoncd-triggers | 0 | — | — | No changes |
| syncer-service | 0 | — | — | No changes |
| tekton-caches | 0 | — | — | No changes |
| tektoncd-hub | 0 | — | — | No changes |
| multicluster-proxy-aae | 0 | — | — | No changes |

---

<details>
<summary><h2>tektoncd-pipeline</h2> — 11 commits, 9 code changes, 1 dependabot</summary>

**Upstream:** tektoncd/pipeline · **Branch:** release-v1.9.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [1db9d78326dc](https://github.com/tektoncd/pipeline/commit/1db9d78326dc) | 07-30 | waveywaves | fix(resolutionrequest): refine lifecycle status handling | [#10494](https://github.com/tektoncd/pipeline/pull/10494) | — |
| [2c49354ac0c7](https://github.com/tektoncd/pipeline/commit/2c49354ac0c7) | 07-29 | waveywaves | fix(resolutionrequest): requeue status update conflicts | [#10494](https://github.com/tektoncd/pipeline/pull/10494) | — |
| [501eee491cde](https://github.com/tektoncd/pipeline/commit/501eee491cde) | 07-29 | waveywaves | fix(resolutionrequest): preserve resolver-written status fields | [#10494](https://github.com/tektoncd/pipeline/pull/10494) | — |
| [d47d6e15a8c3](https://github.com/tektoncd/pipeline/commit/d47d6e15a8c3) | 07-29 | vdemeester | Add missing "time" import to pipelinerun.go | [#10485](https://github.com/tektoncd/pipeline/pull/10485) | — |
| [a08816edec5e](https://github.com/tektoncd/pipeline/commit/a08816edec5e) | 07-14 | sahilleth | Fix PipelineRun stuck in ResolvingTaskRef when RR enqueue is missed | [#10485](https://github.com/tektoncd/pipeline/pull/10485) | — |
| [f54a23722153](https://github.com/tektoncd/pipeline/commit/f54a23722153) | 07-28 | waveywaves | fix(resolvers): skip non-owner reconciliation | [#10482](https://github.com/tektoncd/pipeline/pull/10482) | — |
| [ff5e82e8eb8c](https://github.com/tektoncd/pipeline/commit/ff5e82e8eb8c) | 07-28 | waveywaves | fix(resolvers): honor leader-election bucket ownership | [#10482](https://github.com/tektoncd/pipeline/pull/10482) | — |
| [a51c5c43dbd3](https://github.com/tektoncd/pipeline/commit/a51c5c43dbd3) | 03-09 | twoGiants | fix: skip re-resolution of already resolved requests | [#10477](https://github.com/tektoncd/pipeline/pull/10477) | — |
| [0b947f8ebfa1](https://github.com/tektoncd/pipeline/commit/0b947f8ebfa1) | 06-29 | khrm | ci: pin ko to v0.18.1 to avoid v0.19.0 regressions | — | — |
| [d5b0394db1ed](https://github.com/tektoncd/pipeline/commit/d5b0394db1ed) | 06-29 | khrm | fix: replace kodata LICENSE symlinks with actual files | — | — |

**Key changes:**
- ResolutionRequest lifecycle fixes — status handling, conflict requeue, field preservation ([#10494](https://github.com/tektoncd/pipeline/pull/10494))
- PipelineRun stuck in ResolvingTaskRef fix ([#10485](https://github.com/tektoncd/pipeline/pull/10485))
- Resolver leader-election ownership fixes ([#10482](https://github.com/tektoncd/pipeline/pull/10482))
- Skip re-resolution of already resolved requests ([#10477](https://github.com/tektoncd/pipeline/pull/10477))

</details>

<details>
<summary><h2>operator</h2> — 19 commits, 10 code changes, 5 dependabot</summary>

**Upstream:** tektoncd/operator · **Branch:** release-v0.79.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [3363c0cf85f2](https://github.com/tektoncd/operator/commit/3363c0cf85f2) | 07-29 | divyansh42 | chore: update third-party image digests | [#3835](https://github.com/tektoncd/operator/pull/3835) | — |
| [e3147f8bd097](https://github.com/tektoncd/operator/commit/e3147f8bd097) | 07-29 | vdemeester | chore: bump component versions | [#3829](https://github.com/tektoncd/operator/pull/3829) | — |
| [6d604ab90864](https://github.com/tektoncd/operator/commit/6d604ab90864) | 07-14 | pratap0007 | chore(operatorhub): remove tech-preview label from Results | [#3799](https://github.com/tektoncd/operator/pull/3799) | — |
| [3701c100ec8a](https://github.com/tektoncd/operator/commit/3701c100ec8a) | 07-22 | divyansh42 | chore: update third-party image digests | [#3789](https://github.com/tektoncd/operator/pull/3789) | — |
| [3f797e28cdee](https://github.com/tektoncd/operator/commit/3f797e28cdee) | 07-21 | divyansh42 | chore: update third-party image digests | [#3777](https://github.com/tektoncd/operator/pull/3777) | — |
| [b23704380d17](https://github.com/tektoncd/operator/commit/b23704380d17) | 07-20 | vdemeester | chore: bump pipelines-as-code from v0.42.2 to v0.42.3 | [#3760](https://github.com/tektoncd/operator/pull/3760) | — |
| [8bcee5808b3c](https://github.com/tektoncd/operator/commit/8bcee5808b3c) | 07-17 | vdemeester | chore: bump component versions | [#3722](https://github.com/tektoncd/operator/pull/3722) | — |
| [053dcb2bd7b3](https://github.com/tektoncd/operator/commit/053dcb2bd7b3) | 07-16 | divyansh42 | chore: update third-party image digests | [#3745](https://github.com/tektoncd/operator/pull/3745) | — |
| [07ef5d7ed698](https://github.com/tektoncd/operator/commit/07ef5d7ed698) | 07-13 | jkhelil | fix(common): apply proxy settings to StatefulSets too | [#3738](https://github.com/tektoncd/operator/pull/3738) | [SRVKP-12831](https://redhat.atlassian.net/browse/SRVKP-12831) |
| [0eea55d77b81](https://github.com/tektoncd/operator/commit/0eea55d77b81) | 07-15 | divyansh42 | chore: update third-party image digests | [#3736](https://github.com/tektoncd/operator/pull/3736) | — |
| [57958bde2684](https://github.com/tektoncd/operator/commit/57958bde2684) | 07-14 | divyansh42 | chore: update third-party image digests | [#3697](https://github.com/tektoncd/operator/pull/3697) | — |
| [7ad21530bb7b](https://github.com/tektoncd/operator/commit/7ad21530bb7b) | 07-08 | divyansh42 | chore: update third-party image digests | [#3672](https://github.com/tektoncd/operator/pull/3672) | — |
| [0fe5b6847eb7](https://github.com/tektoncd/operator/commit/0fe5b6847eb7) | 07-06 | theakshaypant | chore: bump component versions to latest patch releases | [#3662](https://github.com/tektoncd/operator/pull/3662) | — |
| [20504775668e](https://github.com/tektoncd/operator/commit/20504775668e) | 07-02 | divyansh42 | chore: update third-party image digests | [#3653](https://github.com/tektoncd/operator/pull/3653) | — |

**Key changes:**
- fix(common): apply proxy settings to StatefulSets too ([#3738](https://github.com/tektoncd/operator/pull/3738)) — [SRVKP-12831](https://redhat.atlassian.net/browse/SRVKP-12831)
- Remove tech-preview label from Results ([#3799](https://github.com/tektoncd/operator/pull/3799))
- Bump pipelines-as-code v0.42.2 → v0.42.3 ([#3760](https://github.com/tektoncd/operator/pull/3760))
- Multiple third-party image digest updates and component version bumps

</details>

<details>
<summary><h2>console-plugin</h2> — 12 commits, 8 code changes (upstream: release-v1.23.x)</summary>

**Upstream:** openshift-pipelines/console-plugin · **Branch:** release-v1.23.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [8f1b55baf2af](https://github.com/openshift-pipelines/console-plugin/commit/8f1b55baf2af) | 07-27 | arvindk-softwaredev | SRVKP-12826,SRVKP-12637,SRVKP-12856,SRVKP-12812: cve fix — js-yaml resolution + lock file refresh | [#1212](https://github.com/openshift-pipelines/console-plugin/pull/1212) | [SRVKP-12826](https://redhat.atlassian.net/browse/SRVKP-12826) |
| [34ab462b3faa](https://github.com/openshift-pipelines/console-plugin/commit/34ab462b3faa) | 07-23 | anwesha-palit-redhat | SRVKP-12879: bumped @patternfly/react-styles 6.4 → 6.6 | [#1188](https://github.com/openshift-pipelines/console-plugin/pull/1188) | [SRVKP-12879](https://redhat.atlassian.net/browse/SRVKP-12879) |
| [0b9697ddd72c](https://github.com/openshift-pipelines/console-plugin/commit/0b9697ddd72c) | 07-09 | anwesha-palit-redhat | cve: fixes for ws, webpack-dev-server, undici and form-data | [#1164](https://github.com/openshift-pipelines/console-plugin/pull/1164) | [SRVKP-12516](https://redhat.atlassian.net/browse/SRVKP-12516) |
| [b725960f01d9](https://github.com/openshift-pipelines/console-plugin/commit/b725960f01d9) | 06-10 | anwesha-palit-redhat | fix: ensure pruned plr and tr are removed without needing to manually re-render | [#1163](https://github.com/openshift-pipelines/console-plugin/pull/1163) | — |
| [43498ae99fb4](https://github.com/openshift-pipelines/console-plugin/commit/43498ae99fb4) | 06-29 | ankrsinha | fix: improve CVE analysis and resolution handling | [#1156](https://github.com/openshift-pipelines/console-plugin/pull/1156) | — |

**Jira tickets:**
- [SRVKP-12826](https://redhat.atlassian.net/browse/SRVKP-12826), [SRVKP-12637](https://redhat.atlassian.net/browse/SRVKP-12637), [SRVKP-12856](https://redhat.atlassian.net/browse/SRVKP-12856), [SRVKP-12812](https://redhat.atlassian.net/browse/SRVKP-12812) — CVE fixes (js-yaml, lock file)
- [SRVKP-12879](https://redhat.atlassian.net/browse/SRVKP-12879) — PatternFly react-styles bump (tab highlighting fix)
- [SRVKP-12516](https://redhat.atlassian.net/browse/SRVKP-12516) — CVE fixes (ws, webpack-dev-server, undici, form-data)

</details>

<details>
<summary><h2>console-plugin-pf5</h2> — 12 commits, 8 code changes (upstream: release-v1.22.x)</summary>

**Upstream:** openshift-pipelines/console-plugin · **Branch:** release-v1.22.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [9fe286e0d8db](https://github.com/openshift-pipelines/console-plugin/commit/9fe286e0d8db) | 07-28 | arvindk-softwaredev | cve fix — adding resolution for js-yaml | [#1216](https://github.com/openshift-pipelines/console-plugin/pull/1216) | [SRVKP-12864](https://redhat.atlassian.net/browse/SRVKP-12864) |
| [71b270edbf9a](https://github.com/openshift-pipelines/console-plugin/commit/71b270edbf9a) | 07-24 | arvindk-softwaredev | manual cherrypick of refresh-yarn-lockfile workflow | [#1207](https://github.com/openshift-pipelines/console-plugin/pull/1207) | — |
| [d5d3745f425c](https://github.com/openshift-pipelines/console-plugin/commit/d5d3745f425c) | 07-13 | arvindk-softwaredev | lockfile refresh to fix cves for ws, webpack-dev-server, undici and form-data | [#1171](https://github.com/openshift-pipelines/console-plugin/pull/1171) | — |
| [913cef78cae3](https://github.com/openshift-pipelines/console-plugin/commit/913cef78cae3) | 06-29 | anwesha-palit-redhat | fix: removal of caching logic for taskruns and pipelineruns | [#1165](https://github.com/openshift-pipelines/console-plugin/pull/1165) | [SRVKP-12032](https://redhat.atlassian.net/browse/SRVKP-12032) |
| [0b933df78b6d](https://github.com/openshift-pipelines/console-plugin/commit/0b933df78b6d) | 06-26 | anwesha-palit-redhat | fix: updated refreshKey to trigger tr api on tr and plr deletion | [#1165](https://github.com/openshift-pipelines/console-plugin/pull/1165) | [SRVKP-12032](https://redhat.atlassian.net/browse/SRVKP-12032) |
| [ca45f71e737a](https://github.com/openshift-pipelines/console-plugin/commit/ca45f71e737a) | 06-29 | ankrsinha | fix: improve CVE analysis and resolution handling | [#1160](https://github.com/openshift-pipelines/console-plugin/pull/1160) | [SRVKP-12532](https://redhat.atlassian.net/browse/SRVKP-12532) |
| [63f44259e1d2](https://github.com/openshift-pipelines/console-plugin/commit/63f44259e1d2) | 07-08 | openshift-merge-bot | cherry-pick-1136-to-release-v1.22.x | [#1139](https://github.com/openshift-pipelines/console-plugin/pull/1139) | [SRVKP-12265](https://redhat.atlassian.net/browse/SRVKP-12265) |

**Jira tickets:**
- [SRVKP-12864](https://redhat.atlassian.net/browse/SRVKP-12864) — CVE fix (js-yaml)
- [SRVKP-12032](https://redhat.atlassian.net/browse/SRVKP-12032) — Deleted Taskrun does not disappear from taskruns page
- [SRVKP-12532](https://redhat.atlassian.net/browse/SRVKP-12532) — CVE skill enhancements for console-plugin
- [SRVKP-12265](https://redhat.atlassian.net/browse/SRVKP-12265) — Add CVE remediation skill to console-plugin

</details>

<details>
<summary><h2>tektoncd-cli</h2> — 18 commits, 3 code changes, 14 dependabot</summary>

**Upstream:** tektoncd/cli · **Branch:** release-v0.44.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [573dfe633d44](https://github.com/tektoncd/cli/commit/573dfe633d44) | 07-25 | divyansh42 | fix(cve): CVE-2026-42505 — update Go stdlib to 1.25.12 | [#3066](https://github.com/tektoncd/cli/pull/3066) | — |
| [502f0e407bc9](https://github.com/tektoncd/cli/commit/502f0e407bc9) | 07-09 | divyansh42 | chore: sync OWNERS file with main | [#3015](https://github.com/tektoncd/cli/pull/3015) | — |
| [52d95059c2e5](https://github.com/tektoncd/cli/commit/52d95059c2e5) | 07-03 | vdemeester | CI: Use goinstall mode for golangci-lint action | [#2987](https://github.com/tektoncd/cli/pull/2987) | — |

**Key changes:**
- CVE-2026-42505: Go stdlib update to 1.25.12

</details>

<details>
<summary><h2>tektoncd-chains</h2> — 12 commits, 1 code change, 11 dependabot</summary>

**Upstream:** tektoncd/chains · **Branch:** release-v0.26.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [cec078dff692](https://github.com/tektoncd/chains/commit/cec078dff692) | 07-03 | tekton-robot | remove symlinks | — | — |

All other commits are dependabot dependency bumps (cosign, scaffolding, codeql-action, etc).

</details>

<details>
<summary><h2>tektoncd-results</h2> — 4 commits, 2 code changes</summary>

**Upstream:** tektoncd/results · **Branch:** release-v0.18.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [cafb640c5326](https://github.com/tektoncd/results/commit/cafb640c5326) | 07-27 | divyansh42 | fix(security): update Go stdlib to 1.25.12 and golang.org/x/text to v0.39.0 | [#1394](https://github.com/tektoncd/results/pull/1394) | — |
| [bba0857b2539](https://github.com/tektoncd/results/commit/bba0857b2539) | 07-13 | divyansh42 | fix: bump vulnerable Go deps for CVE remediation (SRVKP-12035) | [#1378](https://github.com/tektoncd/results/pull/1378) | [SRVKP-12035](https://redhat.atlassian.net/browse/SRVKP-12035) |

**Jira:** [SRVKP-12035](https://redhat.atlassian.net/browse/SRVKP-12035) — gRPC-Go authorization bypass via missing leading slash

</details>

<details>
<summary><h2>manual-approval-gate</h2> — 3 commits, all code changes</summary>

**Upstream:** openshift-pipelines/manual-approval-gate · **Branch:** release-v0.8.0

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [e45845821ac7](https://github.com/openshift-pipelines/manual-approval-gate/commit/e45845821ac7) | 07-22 | divyansh42 | fix(cve): update golang.org/x/text to v0.39.0 — CVE-2026-56852 | [#1036](https://github.com/openshift-pipelines/manual-approval-gate/pull/1036) | [SRVKP-12834](https://redhat.atlassian.net/browse/SRVKP-12834) |
| [ddd1b870ab73](https://github.com/openshift-pipelines/manual-approval-gate/commit/ddd1b870ab73) | 07-18 | divyansh42 | fix(cve): bump Go stdlib go1.25.11 → go1.25.12 — CVE-2026-42505 | [#1032](https://github.com/openshift-pipelines/manual-approval-gate/pull/1032) | — |
| [dc7a6d186cd7](https://github.com/openshift-pipelines/manual-approval-gate/commit/dc7a6d186cd7) | 06-26 | divyansh42 | fix(cve): bump tektoncd/pipeline v1.9.2 → v1.9.3 | [#1018](https://github.com/openshift-pipelines/manual-approval-gate/pull/1018) | — |

**Jira:** [SRVKP-12834](https://redhat.atlassian.net/browse/SRVKP-12834) — CVE Management (referenced in #1036)

</details>

<details>
<summary><h2>tektoncd-pruner</h2> — 1 commit</summary>

**Upstream:** tektoncd/pruner · **Branch:** release-v0.3.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [2b54ebd4133c](https://github.com/tektoncd/pruner/commit/2b54ebd4133c) | 07-09 | infernus01 | fix(webhook): allow operator-managed global ConfigMap deletion | [#354](https://github.com/tektoncd/pruner/pull/354) | — |

</details>

<details>
<summary><h2>pipelines-as-code</h2> — 2 commits (v0.42.3 release)</summary>

**Upstream:** tektoncd/pipelines-as-code · **Branch:** release-v0.42.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [fb73f8cb09dc](https://github.com/tektoncd/pipelines-as-code/commit/fb73f8cb09dc) | 07-17 | pac-robot | Release yaml generated for v0.42.3 | — | — |
| [f799bbf4c00c](https://github.com/tektoncd/pipelines-as-code/commit/f799bbf4c00c) | 07-16 | zakisk | dep(go): bump go version 1.26.4 | [#2858](https://github.com/tektoncd/pipelines-as-code/pull/2858) | — |

</details>

<details>
<summary><h2>opc</h2> — 2 commits</summary>

**Upstream:** openshift-pipelines/opc · **Branch:** release-v1.22.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [4fca3ef6b41c](https://github.com/openshift-pipelines/opc/commit/4fca3ef6b41c) | 08-03 | theakshaypant | [bot:1.22] Update OPC version to 1.22.5 | [#518](https://github.com/openshift-pipelines/opc/pull/518) | — |
| [72754e28fec6](https://github.com/openshift-pipelines/opc/commit/72754e28fec6) | 07-26 | github-actions[bot] | Update component versions for release-v1.22.x | [#499](https://github.com/openshift-pipelines/opc/pull/499) | — |

</details>

<details>
<summary><h2>tekton-kueue</h2> — 3 commits (all renovate bot)</summary>

**Upstream:** konflux-ci/tekton-kueue · **Branch:** release-v0.3.x

| SHA | Date | Author | Message | PR | Jira |
|-----|------|--------|---------|----|------|
| [289d3191a0fb](https://github.com/konflux-ci/tekton-kueue/commit/289d3191a0fb) | 07-29 | red-hat-konflux[bot] | fix(deps): update module go-logr/logr to v1.4.4 | [#515](https://github.com/tektoncd/tekton-kueue/pull/515) | — |
| [86e18dec0875](https://github.com/konflux-ci/tekton-kueue/commit/86e18dec0875) | 07-29 | red-hat-konflux[bot] | chore(deps): update golang.org/x/exp digest | [#514](https://github.com/tektoncd/tekton-kueue/pull/514) | — |
| [408204f9dbd7](https://github.com/konflux-ci/tekton-kueue/commit/408204f9dbd7) | 07-29 | red-hat-konflux[bot] | chore(deps): update prometheus/procfs to v0.21.1 | [#518](https://github.com/tektoncd/tekton-kueue/pull/518) | — |

</details>

---

## Restricted Jiras (referenced in commits/PRs but not accessible)

The following SRVKP tickets were found in commit messages or PR bodies but returned "Issue does not exist or you do not have permission to see it" from the Jira API:

| Key | Source | Component |
|-----|--------|-----------|
| [SRVKP-12826](https://redhat.atlassian.net/browse/SRVKP-12826) | commit message (console-plugin [#1212](https://github.com/openshift-pipelines/console-plugin/pull/1212)) | console-plugin |
| [SRVKP-12637](https://redhat.atlassian.net/browse/SRVKP-12637) | commit message (console-plugin [#1212](https://github.com/openshift-pipelines/console-plugin/pull/1212)) | console-plugin |
| [SRVKP-12856](https://redhat.atlassian.net/browse/SRVKP-12856) | commit message (console-plugin [#1212](https://github.com/openshift-pipelines/console-plugin/pull/1212)) | console-plugin |
| [SRVKP-12812](https://redhat.atlassian.net/browse/SRVKP-12812) | commit message (console-plugin [#1212](https://github.com/openshift-pipelines/console-plugin/pull/1212)) | console-plugin |
| [SRVKP-12516](https://redhat.atlassian.net/browse/SRVKP-12516) | PR body (console-plugin [#1164](https://github.com/openshift-pipelines/console-plugin/pull/1164)) | console-plugin |
| [SRVKP-12864](https://redhat.atlassian.net/browse/SRVKP-12864) | PR body (console-plugin-pf5 [#1216](https://github.com/openshift-pipelines/console-plugin/pull/1216)) | console-plugin-pf5 |

## Unmatched Jiras (fixVersion = Pipelines 1.22.5)

| Key | Summary | Status | Component |
|-----|---------|--------|-----------|
| [SRVKP-13085](https://redhat.atlassian.net/browse/SRVKP-13085) | Test Openshift Pipelines 1.22.5 | New | QA |
| [SRVKP-12879](https://redhat.atlassian.net/browse/SRVKP-12879) | OCP 4.22 — Incorrect tab highlighted in Overview page when switching Per pipeline/Per Repository | Release Pending | UI |
| [SRVKP-12111](https://redhat.atlassian.net/browse/SRVKP-12111) | OSP UI redirects to incorrect namespace/page for resolver Tasks causing misleading 404 | Release Pending | QA, UI |

**Note:** SRVKP-12879 is addressed by console-plugin [#1188](https://github.com/openshift-pipelines/console-plugin/pull/1188) (PatternFly react-styles bump). SRVKP-12111 was fixed in console-plugin [#1078](https://github.com/openshift-pipelines/console-plugin/pull/1078) (release-v1.22.x) — predates the 1.22.4→1.22.5 window.

## Summary

| Component | Total | Code | Dependabot/Bot | Jiras | Highlights |
|-----------|-------|------|----------------|-------|------------|
| tektoncd-pipeline | 11 | 9 | 1 | 0 | ResolutionRequest + resolver fixes |
| operator | 19 | 10 | 5 | 1 | Proxy StatefulSet fix (SRVKP-12831), Results label, PAC bump |
| console-plugin | 12 | 8 | 0 | 3 | CVE fixes (SRVKP-12826, 12516), tab fix (SRVKP-12879) |
| console-plugin-pf5 | 12 | 8 | 0 | 4 | CVE fixes (SRVKP-12864), TR/PLR fix (SRVKP-12032) |
| tektoncd-cli | 18 | 3 | 14 | 0 | CVE-2026-42505 Go stdlib |
| tektoncd-chains | 12 | 1 | 11 | 0 | Symlink removal |
| tektoncd-results | 4 | 2 | 0 | 1 | CVE fixes (SRVKP-12035) |
| manual-approval-gate | 3 | 3 | 0 | 1 | CVE fixes (SRVKP-12834) |
| tektoncd-pruner | 1 | 1 | 0 | 0 | Webhook ConfigMap deletion fix |
| pipelines-as-code | 2 | 1 | 0 | 0 | v0.42.3 Go bump |
| opc | 2 | 2 | 0 | 0 | Version bump to 1.22.5 |
| tekton-kueue | 3 | 0 | 3 | 0 | Renovate dep updates |
| **Total** | **99** | **48** | **34** | **10** | |
