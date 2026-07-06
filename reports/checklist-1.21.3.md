# Release Checklist: 1.21.3

**Generated:** 2026-07-06 17:21 IST
**Release branch:** release-v1.21.x
**Release tag:** 1.21.3
**Code freeze:** true
**Component monitor:** https://openshift-pipelines.github.io/hack/
**Konflux UI:** https://konflux-ui.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com

| # | Step | Status | Details |
|---|------|--------|---------|
| 1 | Release version in hack repo | DONE | release-tag: 1.21.3 matches target |
| 2 | Hack release-manager PR | DONE | [#759](https://github.com/openshift-pipelines/hack/pull/759) merged 2026-06-01 13:24 IST |
| 3 | Konflux config PR | DONE | [#822](https://github.com/openshift-pipelines/hack/pull/822) merged 2026-06-17 20:27 IST |
| 3b | Konflux cluster config | DONE | 11/11 hack apps have matching cluster apps; core 16 dirs → 37 components, bundle 1 dir → 1 component |
| 4 | ReleasePlanAdmission | DONE | 8 RPAs found (core/bundle/fbc/cdn × stage+prod) |
| 5 | Pyxis configuration | DONE | Config exists under products/openshift-pipelines/ |
| 6 | Component PRs on release branches | DONE | 14/14 active components merged; operator has additional open [#24926](https://github.com/openshift-pipelines/operator/pull/24926) |
| 7 | Upstream source sync | DONE | [#24800](https://github.com/openshift-pipelines/operator/pull/24800) merged 2026-07-02 10:01 IST |
| 8 | Nudge PRs (component SHA updates) | IN PROGRESS | 97 merged, 1 open: [#24958](https://github.com/openshift-pipelines/operator/pull/24958) (bundle nudge, failing CI — index pipelines + missing lgtm) |
| 9 | Operator project.yaml version | DONE | `versions.current: 1.21.3` matches target |
| 9b | Image rebuilds triggered | DONE | 4 rebuild runs failed for operator (409); builds triggered via push events and nudge cycle completed |
| 10 | Build validation (SHA comparison) | STALE | All repos have 1-2 Konflux config commits after latest snapshot; no source changes — only `.tekton/` updates |
| 11 | OLM config and index images | **ACTION NEEDED** | olm/config.yaml missing 1.21.3 bundle entry; staging render produced empty catalogs |
| 12 | Code freeze | DONE | code-freeze: true |
| 13 | Stage release | **BLOCKED** | Core stage succeeded; bundle stage succeeded; index stage blocked — empty catalogs from missing olm/config.yaml 1.21.3 entry |

## Konflux Config Details

**Step 3:** Konflux config PR [#822](https://github.com/openshift-pipelines/hack/pull/822) merged 2026-06-17 20:27 IST. Two earlier PRs also merged ([#811](https://github.com/openshift-pipelines/hack/pull/811), [#803](https://github.com/openshift-pipelines/hack/pull/803)).

**Step 3b:** Cluster configuration validated against hack repo.

| Application | Hack Repo Dirs | Cluster Components | Status |
|-------------|---------------|--------------------|--------|
| core | 16 dirs | 37 (multi-image) | OK |
| bundle | 1 dir | 1 | OK |
| index-4.14 | 1 dir | 1 | OK |
| index-4.16 | 1 dir | 1 | OK |
| index-4.18 | 1 dir | 1 | OK |
| index-4.19 | 1 dir | 1 | OK |
| index-4.20 | 1 dir | 1 | OK |
| index-4.21 | 1 dir | 1 | OK |
| index-4.22 | 1 dir | 1 | OK |
| index-4.23 | 1 dir | 1 | OK |
| index-5.0 | 1 dir | 1 | OK |

Cluster has extra `openshift-pipelines-index-4-17-1-21` (not in hack repo — likely from a previous config).

Independently-versioned components correctly excluded from 1.21:
- syncer-service (min-version: 1.22)
- multicluster-proxy-aae (min-version: 1.22)
- tekton-kueue (min-version: 1.22)
- tekton-assist (min-version: 5.0)

**Step 4:** ReleasePlanAdmission files found in konflux-release-data:
- openshift-pipelines-1-21-core-stage.yaml / prod.yaml
- openshift-pipelines-1-21-bundle-stage.yaml / prod.yaml
- openshift-pipelines-1-21-fbc-stage.yaml / prod.yaml
- openshift-pipelines-1-21-core-cdn-stage.yaml / prod.yaml

**Step 5:** Pyxis config exists under `products/openshift-pipelines/` (patch release — no new components to add).

## Component Details

**Step 6:** 14/14 active component PRs merged on release-v1.21.x (4 components excluded via min-version > 1.21):

| Component | Repo | PR | Merged |
|-----------|------|----|--------|
| git-init | tektoncd-git-clone | [#850](https://github.com/openshift-pipelines/tektoncd-git-clone/pull/850) | 2026-07-03 |
| manual-approval-gate | p12n-manual-approval-gate | [#332](https://github.com/openshift-pipelines/p12n-manual-approval-gate/pull/332) | 2026-07-03 |
| pipelines-as-code | pac-downstream | [#1611](https://github.com/openshift-pipelines/pac-downstream/pull/1611) | 2026-07-03 |
| tekton-caches | p12n-tekton-caches | [#236](https://github.com/openshift-pipelines/p12n-tekton-caches/pull/236) | 2026-07-03 |
| tektoncd-chains | tektoncd-chains | [#1176](https://github.com/openshift-pipelines/tektoncd-chains/pull/1176) | 2026-07-03 |
| tektoncd-cli | tektoncd-cli | [#1315](https://github.com/openshift-pipelines/tektoncd-cli/pull/1315) | 2026-07-03 |
| tektoncd-hub | tektoncd-hub | [#1919](https://github.com/openshift-pipelines/tektoncd-hub/pull/1919) | 2026-07-03 |
| tektoncd-pipeline | tektoncd-pipeline | [#2198](https://github.com/openshift-pipelines/tektoncd-pipeline/pull/2198) | 2026-07-03 |
| tektoncd-pruner | tektoncd-pruner | [#884](https://github.com/openshift-pipelines/tektoncd-pruner/pull/884) | 2026-07-03 |
| tektoncd-results | tektoncd-results | [#1188](https://github.com/openshift-pipelines/tektoncd-results/pull/1188) | 2026-07-03 |
| tektoncd-triggers | tektoncd-triggers | [#1288](https://github.com/openshift-pipelines/tektoncd-triggers/pull/1288) | 2026-07-03 |
| operator | operator | [#24890](https://github.com/openshift-pipelines/operator/pull/24890) | 2026-07-02 |
| console-plugin | p12n-console-plugin | [#318](https://github.com/openshift-pipelines/p12n-console-plugin/pull/318) | 2026-07-03 |
| console-plugin-pf5 | p12n-console-plugin-pf5 | [#143](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/pull/143) | 2026-07-03 |

Operator has additional open PR [#24926](https://github.com/openshift-pipelines/operator/pull/24926) (created 2026-07-03 16:54 IST, from latest hack push).

**Step 7:** Upstream source sync PR [#24800](https://github.com/openshift-pipelines/operator/pull/24800) merged 2026-07-02 10:01 IST (synced tektoncd/operator release-v0.78.x).

**Step 8:** 97 nudge PRs merged. 1 open: [#24958](https://github.com/openshift-pipelines/operator/pull/24958) (`chore(deps): update operator-1-21-bundle to e8ee421`, created 2026-07-03 18:36 IST). CI status: all index pipeline runs failing, missing `lgtm` label. This is from the post-stage build cycle and does not block the current stage release.

## Build Validation Details

**Step 9:** Operator project.yaml `versions.current: 1.21.3` — matches target.

**Step 9b:** 4 image rebuild runs on 2026-07-02 all failed for the operator repo (409 Conflict):
- [Run 28591841841](https://github.com/openshift-pipelines/hack/actions/runs/28591841841) at 2026-07-02 18:28 IST — failure
- [Run 28586743804](https://github.com/openshift-pipelines/hack/actions/runs/28586743804) at 2026-07-02 17:02 IST — failure
- [Run 28583912600](https://github.com/openshift-pipelines/hack/actions/runs/28583912600) at 2026-07-02 16:10 IST — failure
- [Run 28582549948](https://github.com/openshift-pipelines/hack/actions/runs/28582549948) at 2026-07-02 15:45 IST — failure

Builds were triggered via Konflux config PR merges and push events instead. Nudge cycle completed.

**Step 10:** Core snapshot `openshift-pipelines-core-1-21-20260704-120600-000` (created 2026-07-04 17:49 IST).

Current build state is **STALE** — all repos have 1-2 commits after the snapshot, but ALL are automated Konflux config updates (`[bot:release-v1.21.x] update konflux configuration`). No source code changes since the stage-released snapshots.

The stage releases (core from 2026-07-02, bundle from 2026-07-03) remain valid — only `.tekton/` pipeline config changes have landed since.

**Operator cycle:** 15 commits between core snapshot and HEAD — all automated (nudge PRs, CSV updates, catalog JSONs, Konflux config). Open nudge PR [#24958](https://github.com/openshift-pipelines/operator/pull/24958) (bundle) — from post-stage build cycle.

**Bundle snapshot:** `openshift-pipelines-bundle-1-21-20260704-125217-000` (latest). Multiple bundle snapshots on 2026-07-04 from the new build cycle.

## OLM and Code Freeze Details

**Step 11: ACTION NEEDED — olm/config.yaml missing 1.21.3 bundle entry.**

**11a — olm/config.yaml** on `release-v1.21.x`:
- Package: `openshift-pipelines-operator-rh`, defaultChannel: `latest`
- Bundles in upgrade chain: **1.18.0, 1.21.2** (missing 1.21.3!)
- The 1.21.2 image ref in config points to staging registry SHA `663c937...` (from 1.21.2 release cycle)

**11b — render-olm-catalog workflow (staging, [run 28766226849](https://github.com/openshift-pipelines/operator/actions/runs/28766226849)):**
- Workflow reported **success**, but `opm render` failed silently for every OCP version
- Root cause: `update-olm-bundle.sh staging` found only the 1.21.2 bundle entry, rewrote its image ref to `registry.stage.redhat.io/openshift-pipelines/pipelines-operator-bundle@sha256:663c937...`
- That image does NOT exist in the staging registry → `opm render` returned "not found" for all versions
- Result: **empty catalog.json files** committed and merged via [#24957](https://github.com/openshift-pipelines/operator/pull/24957) (commit [1d9e6fd3bab8](https://github.com/openshift-pipelines/operator/commit/1d9e6fd3bab8) — 0 additions, 30K+ deletions)

**11c — Index pipeline failures:**
- All 7 index push pipelines failed after PR #24957 merged (empty catalogs → `validate-fbc` FAILURE → `fbc-target-index-pruning-check` failed)
- This is NOT a Konflux infrastructure issue — it's caused by the empty catalogs

**Bundle stage release verification:**
- Konflux bundle stage release **Succeeded** for snapshot `openshift-pipelines-bundle-1-21-20260703-130237-000`
- That snapshot contains image SHA `f3b64690...` (different from `663c937...` used by the render)
- The staged bundle image IS in the staging registry, but the render never used it because olm/config.yaml doesn't have a 1.21.3 entry

**Step 12:** Code freeze is `true` in the release config.

## Release Status Details

**Step 13:** Stage release — core and bundle succeeded, index NOT DONE.

| Application | ReleasePlan | Snapshot | Status | Time |
|-------------|------------|----------|--------|------|
| core | `core-1-21-stage-rp` | `openshift-pipelines-core-1-21-20260702-114559-000-lv` | **Succeeded** | 2026-07-03 |
| bundle | `bundle-1-21-stage-rp` | `openshift-pipelines-bundle-1-21-20260703-130237-000` | **Succeeded** | 2026-07-03 |
| bundle (earlier) | `bundle-1-21-stage-rp` | `openshift-pipelines-bundle-1-21-20260703-125508-000` | Failed | 2026-07-03 |
| bundle (earliest) | `bundle-1-21-stage-rp` | `openshift-pipelines-bundle-1-21-20260701-143248-000` | Failed | 2026-07-01 |
| index | — | — | **NOT DONE** | — |

CSV PRs (all merged):
- [#24956](https://github.com/openshift-pipelines/operator/pull/24956) merged 2026-07-03 18:32 IST (staging CSV)
- [#24888](https://github.com/openshift-pipelines/operator/pull/24888), [#24883](https://github.com/openshift-pipelines/operator/pull/24883), [#24850](https://github.com/openshift-pipelines/operator/pull/24850), [#24847](https://github.com/openshift-pipelines/operator/pull/24847) (earlier nudge-triggered CSVs)

## Next Action

**Step 11 — olm/config.yaml must be updated to add 1.21.3 before render-olm-catalog can work.**

The render-olm-catalog workflow for staging produced **empty catalog.json files** because:
1. `olm/config.yaml` only lists bundles up to 1.21.2 — no 1.21.3 entry
2. The script tried to render the 1.21.2 bundle from the staging registry (SHA `663c937...`) → "not found"
3. `opm render` failed silently → empty catalogs committed via [#24957](https://github.com/openshift-pipelines/operator/pull/24957)
4. All 7 index push pipelines failed because they tried to build from empty catalogs

**Core and bundle stage releases are confirmed valid:**
- Core: `openshift-pipelines-core-1-21-20260702-114559-000-lv` — **Succeeded**
- Bundle: `openshift-pipelines-bundle-1-21-20260703-130237-000` (SHA `f3b64690...`) — **Succeeded**

**Steps to fix:**

1. **Update olm/config.yaml** on `release-v1.21.x` — add 1.21.3 to the bundles list:
   ```yaml
   bundles:
     - version: 1.18.0
       tag: 1.18.0
     - version: 1.21.2
       image: <1.21.2 bundle image>
     - version: 1.21.3
   ```

2. **Re-run render-olm-catalog** with `environment: staging`:
   - https://github.com/openshift-pipelines/operator/actions/workflows/render-olm-catalog.yaml
   - Parameters: Branch: `release-v1.21.x`, Environment: `staging`
   - This time, `update-olm-bundle.sh` should pick up the 1.21.3 bundle and use the staged image (SHA `f3b64690...`)

3. **Merge the new catalog PR** — verify catalog.json files are NOT empty this time

4. **Wait for index push pipelines** to succeed with the populated catalogs

5. **Apply index stage release** using the new index snapshots

## Remaining Steps

- 14 — QA handover
- 15 — Production release
- 16 — Advisory
