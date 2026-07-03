# Release Checklist: 1.22.4

**Generated:** 2026-07-02 11:01 IST
**Release branch:** release-v1.22.x
**Release tag:** 1.22.4
**Code freeze:** false
**Component monitor:** https://openshift-pipelines.github.io/hack/
**Konflux UI:** https://konflux-ui.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com

| # | Step | Status | Details |
|---|------|--------|---------|
| 1 | Release version in hack config | DONE | `release-tag: 1.22.4` matches target version |
| 2 | Hack release-manager PR | DONE | [#827](https://github.com/openshift-pipelines/hack/pull/827) merged 2026-06-19 10:29 IST |
| 3 | Konflux configuration PR | DONE | [#828](https://github.com/openshift-pipelines/hack/pull/828) merged 2026-06-19 10:50 IST |
| 3b | Cluster config matches hack repo | DONE | 12 applications, all present on cluster |
| 4 | ReleasePlanAdmission | DONE | 8 RPAs found (core/bundle/fbc/cdn × stage/prod) |
| 5 | Pyxis configuration | DONE | Existing config sufficient for patch release |
| 6 | Component PRs on release branches | DONE | 19/19 components merged |
| 7 | Upstream source sync | DONE | [#24524](https://github.com/openshift-pipelines/operator/pull/24524) merged 2026-06-30 22:17 IST |
| 8 | Nudge PRs | IN PROGRESS | 1 open: [#24820](https://github.com/openshift-pipelines/operator/pull/24820) — bundle nudge, CI passing, needs `lgtm` label |
| 9 | Operator project.yaml version | DONE | `versions.current: 1.22.4` ✓ |
| 10 | Build validation (SHA comparison) | **STALE** | Operator 21 commits behind HEAD in core snapshot — see details below |

## Build Validation Details

**Core snapshot:** `openshift-pipelines-core-1-22-20260702-034604-000` (2026-07-02 09:16 IST)

| Downstream Repo | HEAD SHA | Snapshot SHA | Status |
|-----------------|----------|-------------|--------|
| openshift-pipelines/tektoncd-pipeline | [`6b33b998e515`](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/6b33b998e5151b664f440ac1ae156c3336c57302) | [`6b33b998e515`](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/6b33b998e5151b664f440ac1ae156c3336c57302) | CURRENT |
| openshift-pipelines/tektoncd-triggers | [`8abaa965abcc`](https://github.com/openshift-pipelines/tektoncd-triggers/commit/8abaa965abccd8cdfdea788ff381681f00194fa0) | [`8abaa965abcc`](https://github.com/openshift-pipelines/tektoncd-triggers/commit/8abaa965abccd8cdfdea788ff381681f00194fa0) | CURRENT |
| openshift-pipelines/tektoncd-chains | [`5ce5347c8544`](https://github.com/openshift-pipelines/tektoncd-chains/commit/5ce5347c8544de34ddc9832921e4d3095f9dfc38) | [`5ce5347c8544`](https://github.com/openshift-pipelines/tektoncd-chains/commit/5ce5347c8544de34ddc9832921e4d3095f9dfc38) | CURRENT |
| openshift-pipelines/tektoncd-results | [`23029ec34482`](https://github.com/openshift-pipelines/tektoncd-results/commit/23029ec344825b3f9bb60a11f8b9fd3cd893bbbe) | [`23029ec34482`](https://github.com/openshift-pipelines/tektoncd-results/commit/23029ec344825b3f9bb60a11f8b9fd3cd893bbbe) | CURRENT |
| openshift-pipelines/tektoncd-hub | [`83d350ed657e`](https://github.com/openshift-pipelines/tektoncd-hub/commit/83d350ed657ef344b81918fca4b9a9270667677d) | [`83d350ed657e`](https://github.com/openshift-pipelines/tektoncd-hub/commit/83d350ed657ef344b81918fca4b9a9270667677d) | CURRENT |
| openshift-pipelines/tektoncd-pruner | [`68bfea16b601`](https://github.com/openshift-pipelines/tektoncd-pruner/commit/68bfea16b60110e1dc6658b615bfe2300d6f630e) | [`68bfea16b601`](https://github.com/openshift-pipelines/tektoncd-pruner/commit/68bfea16b60110e1dc6658b615bfe2300d6f630e) | CURRENT |
| openshift-pipelines/tektoncd-cli | [`3f9d353d6e2b`](https://github.com/openshift-pipelines/tektoncd-cli/commit/3f9d353d6e2bb547b64809d77a8dd17ca8212a78) | [`3f9d353d6e2b`](https://github.com/openshift-pipelines/tektoncd-cli/commit/3f9d353d6e2bb547b64809d77a8dd17ca8212a78) | CURRENT |
| openshift-pipelines/tektoncd-git-clone | [`e07eb2b683f3`](https://github.com/openshift-pipelines/tektoncd-git-clone/commit/e07eb2b683f3c43cc47a7bb091a6ed4b44113a95) | [`e07eb2b683f3`](https://github.com/openshift-pipelines/tektoncd-git-clone/commit/e07eb2b683f3c43cc47a7bb091a6ed4b44113a95) | CURRENT |
| openshift-pipelines/pac-downstream | [`0837eab48f06`](https://github.com/openshift-pipelines/pac-downstream/commit/0837eab48f0638d821941b802354a91cf21ce2b3) | [`0837eab48f06`](https://github.com/openshift-pipelines/pac-downstream/commit/0837eab48f0638d821941b802354a91cf21ce2b3) | CURRENT |
| openshift-pipelines/p12n-console-plugin | [`11c1ea728bb0`](https://github.com/openshift-pipelines/p12n-console-plugin/commit/11c1ea728bb084d4e58f202001b2fc2aaf2cd913) | [`11c1ea728bb0`](https://github.com/openshift-pipelines/p12n-console-plugin/commit/11c1ea728bb084d4e58f202001b2fc2aaf2cd913) | CURRENT |
| openshift-pipelines/p12n-console-plugin-pf5 | [`5f871b8eb87f`](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/commit/5f871b8eb87fe7c7dcadb5294ea29414839da4db) | [`5f871b8eb87f`](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/commit/5f871b8eb87fe7c7dcadb5294ea29414839da4db) | CURRENT |
| openshift-pipelines/p12n-manual-approval-gate | [`b9e652dce488`](https://github.com/openshift-pipelines/p12n-manual-approval-gate/commit/b9e652dce4888b6bcecaa83855d26bcfd41fe19b) | [`b9e652dce488`](https://github.com/openshift-pipelines/p12n-manual-approval-gate/commit/b9e652dce4888b6bcecaa83855d26bcfd41fe19b) | CURRENT |
| openshift-pipelines/p12n-multicluster-proxy-aae | [`86475aae6abe`](https://github.com/openshift-pipelines/p12n-multicluster-proxy-aae/commit/86475aae6abe26a17fe54541cb7c42d131f890c3) | [`86475aae6abe`](https://github.com/openshift-pipelines/p12n-multicluster-proxy-aae/commit/86475aae6abe26a17fe54541cb7c42d131f890c3) | CURRENT |
| openshift-pipelines/p12n-opc | [`75c5770be7ad`](https://github.com/openshift-pipelines/p12n-opc/commit/75c5770be7ada0cdfdc2de4f76dc44e86c48faa5) | [`75c5770be7ad`](https://github.com/openshift-pipelines/p12n-opc/commit/75c5770be7ada0cdfdc2de4f76dc44e86c48faa5) | CURRENT |
| openshift-pipelines/p12n-syncer-service | [`b029f0528ca8`](https://github.com/openshift-pipelines/p12n-syncer-service/commit/b029f0528ca8ca179001d2ded1d806bd3fad1c67) | [`b029f0528ca8`](https://github.com/openshift-pipelines/p12n-syncer-service/commit/b029f0528ca8ca179001d2ded1d806bd3fad1c67) | CURRENT |
| openshift-pipelines/p12n-tekton-caches | [`35d13ca95e59`](https://github.com/openshift-pipelines/p12n-tekton-caches/commit/35d13ca95e59b3278264a3d9b6546afc890662df) | [`35d13ca95e59`](https://github.com/openshift-pipelines/p12n-tekton-caches/commit/35d13ca95e59b3278264a3d9b6546afc890662df) | CURRENT |
| openshift-pipelines/serve-tkn-cli | [`a510d8905687`](https://github.com/openshift-pipelines/serve-tkn-cli/commit/a510d8905687a1be3e3d03ce7771bd8b2011b0ee) | [`a510d8905687`](https://github.com/openshift-pipelines/serve-tkn-cli/commit/a510d8905687a1be3e3d03ce7771bd8b2011b0ee) | CURRENT |
| openshift-pipelines/tekton-kueue | [`9e2bc49b0eb2`](https://github.com/openshift-pipelines/tekton-kueue/commit/9e2bc49b0eb29f028d54778c0e3f1379049723fc) | [`9e2bc49b0eb2`](https://github.com/openshift-pipelines/tekton-kueue/commit/9e2bc49b0eb29f028d54778c0e3f1379049723fc) | CURRENT |
| openshift-pipelines/operator | [`92006961c26a`](https://github.com/openshift-pipelines/operator/commit/92006961c26aadb484d4ff2f93f9620fa6c1db19) | [`e122ff116bf8`](https://github.com/openshift-pipelines/operator/commit/e122ff116bf8f9c1d20c4ca895ad946db4986e5a) / [`b7ec4f688997`](https://github.com/openshift-pipelines/operator/commit/b7ec4f688997d43159c2cf83176c26125b9b7842) | **STALE + SPLIT** (21 / 6 commits behind) |

**Bundle snapshot:** `openshift-pipelines-bundle-1-22-20260702-045606-000` (2026-07-02 10:26 IST)
- Operator revision: [`92006961c26a`](https://github.com/openshift-pipelines/operator/commit/92006961c26aadb484d4ff2f93f9620fa6c1db19) — **CURRENT** ✓ (matches operator HEAD)

### Operator commit gap (core snapshot → HEAD)

The operator HEAD [`92006961c26a`](https://github.com/openshift-pipelines/operator/commit/92006961c26aadb484d4ff2f93f9620fa6c1db19) is 21 commits ahead of the older snapshot SHA. Key commits in the gap:

| SHA | Message |
|-----|---------|
| [`d44dbac53d86`](https://github.com/openshift-pipelines/operator/commit/d44dbac53d86) | Revert nudge: operator-1-22-operator |
| [`ff91db419190`](https://github.com/openshift-pipelines/operator/commit/ff91db419190) | Revert nudge: operator-1-22-bundle |
| [`880d5806218b`](https://github.com/openshift-pipelines/operator/commit/880d5806218b) | Add .placeholder for index rebuild |
| 9× | `[bot:release-v1.22.x] update konflux configuration` |
| [`0bfa3e2270d4`](https://github.com/openshift-pipelines/operator/commit/0bfa3e2270d4) | Rebuild 4.19 FBC |
| [`92006961c26a`](https://github.com/openshift-pipelines/operator/commit/92006961c26a) | [bot] Update CSV |

### Open nudge PR

[#24820](https://github.com/openshift-pipelines/operator/pull/24820) — `chore(deps): update operator-1-22-bundle to 47302ed` (CI passing, needs `lgtm` label for Prow merge)

## Next Action

**Step 10: Trigger operator image rebuild and wait for convergence**

The core snapshot operator images are 21 commits behind HEAD. The gap includes Konflux config updates, FBC rebuilds, reverts, and a CSV update. No blocking PRs — the operator HEAD is stable (latest commit is a CSV update).

1. **Ensure nudge PR [#24820](https://github.com/openshift-pipelines/operator/pull/24820) merges**
   - CI is passing, needs `lgtm` label
   - This will update the bundle image SHA in `project.yaml`

2. **Trigger image rebuilds** (after nudge merges):
   - Go to: https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml
   - Parameters: version: `1.22`, repo: `operator`

3. **Wait for convergence:**
   - Operator images rebuild → nudge PRs update `project.yaml` → HEAD moves → converges
   - Monitor: https://openshift-pipelines.github.io/hack/

4. **Re-run `/release-checklist 1.22.4`** to verify all SHAs match in a new core snapshot

## Remaining Steps

| # | Step |
|---|------|
| 11 | OLM config and index images |
| 12 | Code freeze |
| 13 | Stage release (core → bundle → index) |
| 14 | QA handover |
| 15 | Production release |
| 16 | Advisory |
