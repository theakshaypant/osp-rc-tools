# Release Checklist: 1.22.4

**Generated:** 2026-07-01 13:58 IST
**Release branch:** release-v1.22.x
**Release tag:** 1.22.4
**Code freeze:** false
**Component monitor:** https://openshift-pipelines.github.io/hack/
**Konflux UI:** https://konflux-ui.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com

| # | Step | Status | Details |
|---|------|--------|---------|
| 1 | Generate patch in hack | DONE | release-tag: 1.22.4 |
| 2 | Hack release PR merged | DONE | PR [#827](https://github.com/openshift-pipelines/hack/pull/827) merged 2026-06-19 10:29 IST |
| 3 | Konflux config PR merged | DONE | PR [#828](https://github.com/openshift-pipelines/hack/pull/828) merged 2026-06-19 10:50 IST |
| 3b | Konflux config applied | DONE | Applications and components confirmed in `tekton-ecosystem-tenant` |
| 4 | ReleasePlanAdmission | DONE | 8 RPA files on GitLab (core/bundle/fbc, stage+prod) |
| 5 | Pyxis configuration | DONE | Config exists under `products/openshift-pipelines/` |
| 6 | Component PRs merged | PARTIAL | 16/19 merged, 3 open (see below) |
| 7 | Upstream sources synced | DONE | PR [#24524](https://github.com/openshift-pipelines/operator/pull/24524) merged 2026-06-30 22:17 IST |
| 8 | Component builds & nudges | PARTIAL | 99 merged, 1 open (see below) |
| 9 | Build validation (SHA) | ISSUES | 3 repos stale/split — operator, tektoncd-hub, tektoncd-pipeline |
| 10 | OLM config & index images | DONE | olm/config.yaml has 1.22.4; index images releasing successfully |
| 11 | Code freeze set | NOT DONE | code-freeze: false |
| 12 | Stage release | ISSUES | Core+Bundle stage: Succeeded; CDN stage: mixed; **Index 4-21 stage: FAILING** |
| 13 | QA handover | NOT DONE | |
| 14 | Production release | NOT DONE | No production releases found on Konflux |
| 15 | Advisory | NOT DONE | |

---

## Open Component PRs (Step 6)

3 Konflux configuration PRs are still open. They have the `hack` label and should auto-merge via `merge-hack-pull-requests` (runs hourly):

| Downstream Repo | PR | Component |
|-----------------|-----|-----------|
| p12n-manual-approval-gate | [#319](https://github.com/openshift-pipelines/p12n-manual-approval-gate/pull/319) | manual-approval-gate |
| pac-downstream | [#1600](https://github.com/openshift-pipelines/pac-downstream/pull/1600) | pipelines-as-code |
| tektoncd-git-clone | [#834](https://github.com/openshift-pipelines/tektoncd-git-clone/pull/834) | git-init |

Previous rounds of component PRs for these repos were already merged. These 3 should not block the release pipeline.

## Open Nudge PR (Step 8)

PR [#24614](https://github.com/openshift-pipelines/operator/pull/24614) — `chore(deps): update operator-1-22-bundle to 66533a6` — created **2026-07-01 13:09 IST** (today).

**CI status:** All Konflux index pipelines **pass**. Tide reports: "Not mergeable. Needs lgtm label." This is a fresh PR — may need time for auto-labeling.

## Build Validation — SHA Comparison (Step 9)

Latest core snapshot: `openshift-pipelines-core-1-22-20260701-071148-000-ry` (created 2026-07-01 12:41 IST)

| Downstream Repo | HEAD SHA | Snapshot SHA | Status |
|-----------------|----------|-------------|--------|
| openshift-pipelines/operator | 6547a2082730 | a8570514dd8f (3) | **STALE** |
| openshift-pipelines/p12n-console-plugin | 11c1ea728bb0 | 11c1ea728bb0 | CURRENT |
| openshift-pipelines/p12n-console-plugin-pf5 | 5f871b8eb87f | 5f871b8eb87f | CURRENT |
| openshift-pipelines/p12n-manual-approval-gate | f25d0d7642f0 | f25d0d7642f0 | CURRENT |
| openshift-pipelines/p12n-multicluster-proxy-aae | 86475aae6abe | 86475aae6abe | CURRENT |
| openshift-pipelines/p12n-opc | 75c5770be7ad | 75c5770be7ad | CURRENT |
| openshift-pipelines/p12n-syncer-service | b029f0528ca8 | b029f0528ca8 | CURRENT |
| openshift-pipelines/p12n-tekton-caches | 35d13ca95e59 | 35d13ca95e59 | CURRENT |
| openshift-pipelines/pac-downstream | bd2efb6bac78 | bd2efb6bac78 (4) | CURRENT |
| openshift-pipelines/serve-tkn-cl | — | a510d8905687 | N/A (no release branch) |
| openshift-pipelines/tekton-kueue | 9e2bc49b0eb2 | 9e2bc49b0eb2 | CURRENT |
| openshift-pipelines/tektoncd-chains | 5ce5347c8544 | 5ce5347c8544 | CURRENT |
| openshift-pipelines/tektoncd-cl | — | 3f9d353d6e2b | N/A (no release branch) |
| openshift-pipelines/tektoncd-git-clone | f4ea3d89f9db | f4ea3d89f9db | CURRENT |
| openshift-pipelines/tektoncd-hub | ff21b02885dc | 410f9709ee1d (2), ff21b02885dc (1) | **SPLIT** |
| openshift-pipelines/tektoncd-pipeline | 6b33b998e515 | 6b33b998e515 (7), 1774b2fdc8f3 (1) | **SPLIT** |
| openshift-pipelines/tektoncd-pruner | 68bfea16b601 | 68bfea16b601 | CURRENT |
| openshift-pipelines/tektoncd-results | 23029ec34482 | 23029ec34482 | CURRENT |
| openshift-pipelines/tektoncd-triggers | 8abaa965abcc | 8abaa965abcc | CURRENT |

**3 repos need attention:**

1. **operator** — Snapshot built from `a8570514dd8f`, HEAD is `6547a2082730`. Stale build; all 3 components need rebuild.
2. **tektoncd-hub** — HEAD is `ff21b02885dc`. 1 component at HEAD, 2 at older rev `410f9709ee1d`. Split build; 2 components need rebuild.
3. **tektoncd-pipeline** — HEAD is `6b33b998e515`. 7 components at HEAD, 1 at older rev `1774b2fdc8f3`. Split build; 1 component needs rebuild.

**Bundle snapshot:** `openshift-pipelines-bundle-1-22-20260701-075342-000` (created 2026-07-01 13:23 IST)
Bundle operator revision: `6547a2082730` — **matches operator HEAD** (bundle is current).

**Action:** Trigger image rebuilds for stale core components via:
https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml

## OLM Config (Step 10)

`olm/config.yaml` on `release-v1.22.x` has bundles up to **1.22.4** — **DONE**.

Index images releasing successfully — latest index releases on Konflux at 2026-07-01 11:10 IST.

## Konflux Stage Release (Step 12)

**Core stage:** Succeeded (multiple rounds, latest on 2026-06-30)

**Bundle stage:** Succeeded (latest on 2026-06-30, older failed on 2026-06-24 then retried successfully)

**CDN stage:** Mixed results:
| Release | Status |
|---------|--------|
| cdn-stage-release-ttczk | **Succeeded** |
| cdn-stage-release-zxqjw | Failed |
| cdn-stage-release-z8dd7 | Failed |
| cdn-stage-release-z8dd7-rerun | Failed |
| cdn-stage-release-snl6n | Failed |

**Index stage:** All OCP versions Succeeded **except index-4-21**:
- index-4-21 stage: **FAILING** — 3 attempts all failed (latest rerun failed)
- All other versions (4.14–4.20, 4.22–4.23, 5.0): Succeeded

**Blocker:** Index 4-21 stage release is failing. Investigate the failed release:
- `openshift-pipelines-index-4-21-1-22-stage-release-djwwk` and its reruns all failed.

## Cluster Access Notes

The Konflux token for user `akpant` in namespace `tekton-ecosystem-tenant`:
- **Can:** `oc get applications`, `oc get releases`, `oc get snapshots`
- **Cannot:** `oc get components`, `oc get namespaces` (Forbidden)

---

## Next Actions (Priority Order)

1. **Fix index-4-21 stage release** (Step 12) — 3 attempts all failed. Investigate the release failure. Try with a newer snapshot: `openshift-pipelines-index-4-21-1-22-20260701-054031-000` (regular release succeeded).

2. **Rebuild stale/split core components** (Step 9) — 3 repos have builds from outdated commits: operator (all 3 components), tektoncd-hub (2 components), tektoncd-pipeline (1 component). Trigger rebuilds via [trigger-image-rebuilds](https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml).

3. **Monitor nudge PR #24614** (Step 8) — Created today, all pipelines pass. Waiting for lgtm label. Should auto-merge soon.

4. **Set code freeze** (Step 11) — Update `code-freeze: true` in [hack release config](https://github.com/openshift-pipelines/hack/blob/main/config/downstream/releases/1.22.yaml) when ready to hand off to QE.

5. **Wait for component PRs to auto-merge** (Step 6) — 3 open PRs should merge via hourly `merge-hack-pull-requests`.

6. **QA handover** (Step 13) — After stage issues resolved and code freeze set.

## Remaining Steps After Blockers

```
Fix index-4-21 stage → Rebuild stale core components → New nudge PRs → Code freeze → QA handover → Prod release → Advisory
```
