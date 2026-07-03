# Release Checklist: 1.21.3

**Generated:** 2026-07-02 19:20 IST
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
| 3b | Konflux cluster config | DONE | 12/12 apps match, 10/12 have components (index-4.23 and index-5.0 empty in hack repo too) |
| 4 | ReleasePlanAdmission | DONE | 8 RPAs found (core/bundle/fbc/cdn × stage+prod) |
| 5 | Pyxis configuration | DONE | Config exists under products/openshift-pipelines/ |
| 6 | Component PRs on release branches | DONE | 17/17 components merged |
| 7 | Upstream source sync | DONE | [#24800](https://github.com/openshift-pipelines/operator/pull/24800) merged 2026-07-02 10:01 IST |
| 8 | Nudge PRs (component SHA updates) | DONE | 98 merged, 0 open ([#24844](https://github.com/openshift-pipelines/operator/pull/24844) closed, replaced by [#24889](https://github.com/openshift-pipelines/operator/pull/24889)) |
| 9 | Operator project.yaml version | DONE | `versions.current: 1.21.3` matches target |
| 9b | Image rebuilds triggered | DONE | Workflow failed for operator (409 — expected); operator rebuilds via push events instead |
| 10 | Build validation (SHA comparison) | DONE | 15/16 repos CURRENT, operator CYCLE OK (nudge-only commits; path filters prevent re-trigger) |
| 11 | OLM config and index images | DONE | render-olm-catalog succeeded; index releases for OCP 4.16–4.22 all succeeded |
| 12 | Code freeze | DONE | code-freeze: true |
| 13 | Stage release | **ACTION NEEDED** | Core stage succeeded (old snapshot); bundle stage failed; needs fresh release with 2026-07-02 snapshots |

## Konflux Config Details

**Step 3:** Konflux config PR [#822](https://github.com/openshift-pipelines/hack/pull/822) merged 2026-06-17 20:27 IST. Two earlier PRs also merged ([#811](https://github.com/openshift-pipelines/hack/pull/811), [#803](https://github.com/openshift-pipelines/hack/pull/803)).

**Step 3b:** Cluster configuration validated against hack repo.

| Application | Hack Repo Dirs | Cluster Components | Status |
|-------------|---------------|--------------------|--------|
| core | 16 dirs | 37 (multi-image) | OK |
| bundle | 1 dir | 1 | OK |
| index-4.14 | 1 dir | 1 | OK |
| index-4.16 | 1 dir | 1 | OK |
| index-4.17 | 1 dir | 1 | OK |
| index-4.18 | 1 dir | 1 | OK |
| index-4.19 | 1 dir | 1 | OK |
| index-4.20 | 1 dir | 1 | OK |
| index-4.21 | 1 dir | 1 | OK |
| index-4.22 | 1 dir | 1 | OK |
| index-4.23 | 0 dirs | 0 | OK (no component dirs in hack repo) |
| index-5.0 | 0 dirs | 0 | OK (no component dirs in hack repo) |

Core components on cluster (37 total): all 16 hack repo component directories have matching cluster components. Each directory produces 1-8 image components (e.g. tektoncd-pipeline → 8 components: controller, entrypoint, events, nop, resolvers, sidecarlogresults, webhook, workingdirinit).

Independently-versioned components correctly excluded from 1.21:
- syncer-service (min-version: 1.22) → has own app `syncer-service-0-1`
- multicluster-proxy-aae (min-version: 1.22) → has own app `multicluster-proxy-aae-0-1`
- tekton-kueue (min-version: 1.22) → excluded
- tekton-assist (min-version: 5.0) → excluded

**Step 4:** ReleasePlanAdmission files found in konflux-release-data:
- openshift-pipelines-1-21-core-stage.yaml / prod.yaml
- openshift-pipelines-1-21-bundle-stage.yaml / prod.yaml
- openshift-pipelines-1-21-fbc-stage.yaml / prod.yaml
- openshift-pipelines-1-21-core-cdn-stage.yaml / prod.yaml

**Step 5:** Pyxis config exists under `products/openshift-pipelines/` (patch release — no new components to add).

## Component Details

**Step 6:** All 17 component PRs merged on release-v1.21.x (tekton-kueue excluded — min-version: 1.22):

| Component | Repo | PR | Merged |
|-----------|------|----|--------|
| git-init | tektoncd-git-clone | [#838](https://github.com/openshift-pipelines/tektoncd-git-clone/pull/838) | 2026-07-01 |
| manual-approval-gate | p12n-manual-approval-gate | [#321](https://github.com/openshift-pipelines/p12n-manual-approval-gate/pull/321) | 2026-07-01 |
| pipelines-as-code | pac-downstream | [#1603](https://github.com/openshift-pipelines/pac-downstream/pull/1603) | 2026-07-01 |
| tekton-assist | p12n-tekton-assist | [#159](https://github.com/openshift-pipelines/p12n-tekton-assist/pull/159) | 2026-06-26 |
| tekton-caches | p12n-tekton-caches | [#229](https://github.com/openshift-pipelines/p12n-tekton-caches/pull/229) | 2026-06-18 |
| tektoncd-chains | tektoncd-chains | [#1162](https://github.com/openshift-pipelines/tektoncd-chains/pull/1162) | 2026-06-26 |
| tektoncd-cli | tektoncd-cli | [#1299](https://github.com/openshift-pipelines/tektoncd-cli/pull/1299) | 2026-06-18 |
| tektoncd-hub | tektoncd-hub | [#1894](https://github.com/openshift-pipelines/tektoncd-hub/pull/1894) | 2026-06-23 |
| tektoncd-pipeline | tektoncd-pipeline | [#2155](https://github.com/openshift-pipelines/tektoncd-pipeline/pull/2155) | 2026-06-18 |
| tektoncd-pruner | tektoncd-pruner | [#860](https://github.com/openshift-pipelines/tektoncd-pruner/pull/860) | 2026-06-18 |
| tektoncd-results | tektoncd-results | [#1170](https://github.com/openshift-pipelines/tektoncd-results/pull/1170) | 2026-06-18 |
| tektoncd-triggers | tektoncd-triggers | [#1271](https://github.com/openshift-pipelines/tektoncd-triggers/pull/1271) | 2026-06-18 |
| syncer-service | p12n-syncer-service | [#179](https://github.com/openshift-pipelines/p12n-syncer-service/pull/179) | 2026-06-03 |
| multicluster-proxy-aae | p12n-multicluster-proxy-aae | [#178](https://github.com/openshift-pipelines/p12n-multicluster-proxy-aae/pull/178) | 2026-06-03 |
| operator | operator | [#23616](https://github.com/openshift-pipelines/operator/pull/23616) | 2026-06-18 |
| console-plugin | p12n-console-plugin | [#286](https://github.com/openshift-pipelines/p12n-console-plugin/pull/286) | 2026-06-18 |
| console-plugin-pf5 | p12n-console-plugin-pf5 | [#125](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/pull/125) | 2026-06-18 |

**Step 7:** Upstream source sync PR [#24800](https://github.com/openshift-pipelines/operator/pull/24800) merged 2026-07-02 10:01 IST (synced tektoncd/operator release-v0.78.x).

**Step 8:** All nudge PRs merged (98 total). PR [#24844](https://github.com/openshift-pipelines/operator/pull/24844) was closed (conflicts) and replaced by [#24889](https://github.com/openshift-pipelines/operator/pull/24889) (merged 2026-07-02 18:08 IST).

## Build Validation Details

**Step 9:** Operator project.yaml `versions.current: 1.21.3` — matches target.

**Step 9b:** 3 image rebuild runs on 2026-07-02 all failed for the operator repo (409 Conflict):
- [Run 28586743804](https://github.com/openshift-pipelines/hack/actions/runs/28586743804) at 2026-07-02 17:02 IST — failure
- [Run 28583912600](https://github.com/openshift-pipelines/hack/actions/runs/28583912600) at 2026-07-02 16:10 IST — failure
- [Run 28582549948](https://github.com/openshift-pipelines/hack/actions/runs/28582549948) at 2026-07-02 15:45 IST — failure

Latest upstream sync [#24800](https://github.com/openshift-pipelines/operator/pull/24800) merged 2026-07-02 10:01 IST. Rebuilds were triggered AFTER the sync but failed. User pushed an empty commit to trigger Konflux directly. Operator push builds succeeded on Konflux UI (operator, proxy, webhook all from [f48b26f34624](https://github.com/openshift-pipelines/operator/commit/f48b26f3462459c90c7431d36d360b05f4647825)). Release tests still running at time of report.

**Step 10:** Core snapshot `openshift-pipelines-core-1-21-20260702-114559-000-lv` (created 2026-07-02 17:25 IST):

| Downstream Repo | HEAD SHA | Snapshot SHA | Status |
|-----------------|----------|-------------|--------|
| openshift-pipelines/operator | [cf8d09bbbe02](https://github.com/openshift-pipelines/operator/commit/cf8d09bbbe02cf32b9b19b8a3b8e81e60f00b1a5) | [f48b26f34624](https://github.com/openshift-pipelines/operator/commit/f48b26f3462459c90c7431d36d360b05f4647825) | **CYCLE OK** |
| openshift-pipelines/p12n-console-plugin | [aa8dec80f83c](https://github.com/openshift-pipelines/p12n-console-plugin/commit/aa8dec80f83cdd7d42c7b2ba77ad8bc486442a28) | [aa8dec80f83c](https://github.com/openshift-pipelines/p12n-console-plugin/commit/aa8dec80f83cdd7d42c7b2ba77ad8bc486442a28) | CURRENT |
| openshift-pipelines/p12n-console-plugin-pf5 | [52f0c5fa05b5](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/commit/52f0c5fa05b5edfa3bc7e13b12c79d3793400078) | [52f0c5fa05b5](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/commit/52f0c5fa05b5edfa3bc7e13b12c79d3793400078) | CURRENT |
| openshift-pipelines/p12n-manual-approval-gate | [3d799f812374](https://github.com/openshift-pipelines/p12n-manual-approval-gate/commit/3d799f812374e5a7dcf2cfd666e34ec2b5c7e61d) | [3d799f812374](https://github.com/openshift-pipelines/p12n-manual-approval-gate/commit/3d799f812374e5a7dcf2cfd666e34ec2b5c7e61d) | CURRENT |
| openshift-pipelines/p12n-opc | [7ce8d4640de5](https://github.com/openshift-pipelines/p12n-opc/commit/7ce8d4640de55fcd7e59b5708a6cfb39d3a4df36) | [7ce8d4640de5](https://github.com/openshift-pipelines/p12n-opc/commit/7ce8d4640de55fcd7e59b5708a6cfb39d3a4df36) | CURRENT |
| openshift-pipelines/p12n-tekton-caches | [12d6f0326764](https://github.com/openshift-pipelines/p12n-tekton-caches/commit/12d6f0326764016f24a54763bbab4d49d5d7628d) | [12d6f0326764](https://github.com/openshift-pipelines/p12n-tekton-caches/commit/12d6f0326764016f24a54763bbab4d49d5d7628d) | CURRENT |
| openshift-pipelines/pac-downstream | [27c5b87b7737](https://github.com/openshift-pipelines/pac-downstream/commit/27c5b87b7737caf86aabad7da6c79291e028c4c0) | [27c5b87b7737](https://github.com/openshift-pipelines/pac-downstream/commit/27c5b87b7737caf86aabad7da6c79291e028c4c0) | CURRENT |
| openshift-pipelines/serve-tkn-cli | [09410521ff8b](https://github.com/openshift-pipelines/serve-tkn-cli/commit/09410521ff8b2f7502dbb259b77702e5e5988fd8) | [09410521ff8b](https://github.com/openshift-pipelines/serve-tkn-cli/commit/09410521ff8b2f7502dbb259b77702e5e5988fd8) | CURRENT |
| openshift-pipelines/tektoncd-chains | [b214adea8d96](https://github.com/openshift-pipelines/tektoncd-chains/commit/b214adea8d96350f5014dc985db05abe94eca4c8) | [b214adea8d96](https://github.com/openshift-pipelines/tektoncd-chains/commit/b214adea8d96350f5014dc985db05abe94eca4c8) | CURRENT |
| openshift-pipelines/tektoncd-cli | [62a2bfa60512](https://github.com/openshift-pipelines/tektoncd-cli/commit/62a2bfa6051236a3a188790b281246e47d833860) | [62a2bfa60512](https://github.com/openshift-pipelines/tektoncd-cli/commit/62a2bfa6051236a3a188790b281246e47d833860) | CURRENT |
| openshift-pipelines/tektoncd-git-clone | [d58c1da8401d](https://github.com/openshift-pipelines/tektoncd-git-clone/commit/d58c1da8401d8c2ba0aec3a62dd0bd3a77bcd429) | [d58c1da8401d](https://github.com/openshift-pipelines/tektoncd-git-clone/commit/d58c1da8401d8c2ba0aec3a62dd0bd3a77bcd429) | CURRENT |
| openshift-pipelines/tektoncd-hub | [9765ea32aaf7](https://github.com/openshift-pipelines/tektoncd-hub/commit/9765ea32aaf7c9da252c1048167e13fd1508f23d) | [9765ea32aaf7](https://github.com/openshift-pipelines/tektoncd-hub/commit/9765ea32aaf7c9da252c1048167e13fd1508f23d) | CURRENT |
| openshift-pipelines/tektoncd-pipeline | [4046f59d500b](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/4046f59d500b95fe917c50dc8576d61924232e48) | [4046f59d500b](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/4046f59d500b95fe917c50dc8576d61924232e48) | CURRENT |
| openshift-pipelines/tektoncd-pruner | [dce8205765f6](https://github.com/openshift-pipelines/tektoncd-pruner/commit/dce8205765f6b8f214c10b0f08ae4f325354d5e7) | [dce8205765f6](https://github.com/openshift-pipelines/tektoncd-pruner/commit/dce8205765f6b8f214c10b0f08ae4f325354d5e7) | CURRENT |
| openshift-pipelines/tektoncd-results | [e6e23bd7b856](https://github.com/openshift-pipelines/tektoncd-results/commit/e6e23bd7b856380d4d99002f4f08178e5db4fff8) | [e6e23bd7b856](https://github.com/openshift-pipelines/tektoncd-results/commit/e6e23bd7b856380d4d99002f4f08178e5db4fff8) | CURRENT |
| openshift-pipelines/tektoncd-triggers | [eec3c389a062](https://github.com/openshift-pipelines/tektoncd-triggers/commit/eec3c389a062083c0603ad852e618252ff6c36d3) | [eec3c389a062](https://github.com/openshift-pipelines/tektoncd-triggers/commit/eec3c389a062083c0603ad852e618252ff6c36d3) | CURRENT |

**Operator cycle check:** 6 commits between core snapshot and HEAD — all automated:
- [60deae70502c](https://github.com/openshift-pipelines/operator/commit/60deae70502c) — chore(deps): update operator-1-21-proxy to 1798c5f
- [3ef81248f1a6](https://github.com/openshift-pipelines/operator/commit/3ef81248f1a6) — chore(deps): update operator-1-21-webhook to 2ee41ee
- [4c36224d5e7e](https://github.com/openshift-pipelines/operator/commit/4c36224d5e7e) — chore(deps): update operator-1-21-operator to 577b632
- [aba23565352e](https://github.com/openshift-pipelines/operator/commit/aba23565352e) — [bot:release-v1.21.x] Update generate CSV
- [ecfb70f1dbeb](https://github.com/openshift-pipelines/operator/commit/ecfb70f1dbeb) — chore(deps): update operator-1-21-bundle to 663c937
- [cf8d09bbbe02](https://github.com/openshift-pipelines/operator/commit/cf8d09bbbe02) — chore: auto-generate OCP catalog JSONs

No open nudge PRs. Operator is **CYCLE OK** — this is NOT an infinite loop. Push pipelines use CEL path filters: operator/proxy/webhook only trigger on `upstream/`, `.konflux/patches/`, `.konflux/rpms/` changes. Nudge PRs only change `project.yaml`, which is excluded from all path filters. The build flow is one-directional:
1. Source changes → operator builds → nudge PRs update `project.yaml` → **no rebuild**
2. CSV generation → `.konflux/olm-catalog/bundle/` changes → bundle build triggers
3. Catalog JSON generation → `.konflux/olm-catalog/index/` changes → index builds trigger

**Bundle snapshot:** `openshift-pipelines-bundle-1-21-20260702-120548-000` (created 2026-07-02 17:41 IST):
- Bundle operator revision: [aba23565352e](https://github.com/openshift-pipelines/operator/commit/aba23565352e)
- Operator HEAD: [cf8d09bbbe02](https://github.com/openshift-pipelines/operator/commit/cf8d09bbbe02)
- 2 commits behind: bundle nudge + catalog JSONs (both automated, don't trigger operator rebuild)

## OLM and Code Freeze Details

**Step 11:** OLM configuration verified.

**11a — olm/config.yaml** on `release-v1.21.x`:
- Package: `openshift-pipelines-operator-rh`, defaultChannel: `latest`
- Bundles in upgrade chain: 1.18.0, 1.21.2
- MinVersion: v4.14 → 1.14.0, v4.16–v4.21 → 1.15.0, v4.22 → 1.21.0, v4.23+ → 1.22.0 (excluded)

**11b — render-olm-catalog workflow:**
- Latest run on `release-v1.21.x`: 2026-07-02 18:08 IST — **success**
- Title: `chore: auto-generate OCP catalog JSONs`

**11c — Index releases** (all Succeeded):

| OCP Version | Latest Snapshot | Created |
|-------------|----------------|---------|
| 4.16 | `openshift-pipelines-index-4-16-1-21-20260702-123830-000` | 2026-07-02 18:08 IST |
| 4.18 | `openshift-pipelines-index-4-18-1-21-20260702-123831-000` | 2026-07-02 18:08 IST |
| 4.19 | `openshift-pipelines-index-4-19-1-21-20260702-123830-000` | 2026-07-02 18:08 IST |
| 4.20 | `openshift-pipelines-index-4-20-1-21-20260702-123831-000` | 2026-07-02 18:08 IST |
| 4.21 | `openshift-pipelines-index-4-21-1-21-20260702-123830-000` | 2026-07-02 18:08 IST |
| 4.22 | `openshift-pipelines-index-4-22-1-21-20260702-123830-000` | 2026-07-02 18:08 IST |

**Step 12:** Code freeze is `true` in the release config.

## Release Status Details

**Step 13:** Stage release — needs fresh release with latest snapshots.

Previous stage release activity (from 2026-07-01, before latest build cycle):

| Application | ReleasePlan | Snapshot | Status | Time |
|-------------|------------|----------|--------|------|
| core | `core-1-21-stage-rp` | `openshift-pipelines-core-1-21-20260701-140700-000-4q` | Succeeded | ~2026-07-01 19:37 IST |
| bundle | `bundle-1-21-stage-rp` | `openshift-pipelines-bundle-1-21-20260701-143248-000` | **Failed** | ~2026-07-01 20:02 IST |

Latest available snapshots (from 2026-07-02 build cycle):
- Core: `openshift-pipelines-core-1-21-20260702-114559-000-lv` (2026-07-02 17:15 IST)
- Bundle: `openshift-pipelines-bundle-1-21-20260702-120548-000` (2026-07-02 17:35 IST)
- Index (4.22 sample): `openshift-pipelines-index-4-22-1-21-20260702-123830-000` (2026-07-02 18:08 IST)

CSV PRs (all merged): [#24888](https://github.com/openshift-pipelines/operator/pull/24888), [#24883](https://github.com/openshift-pipelines/operator/pull/24883), [#24850](https://github.com/openshift-pipelines/operator/pull/24850), [#24847](https://github.com/openshift-pipelines/operator/pull/24847), [#24834](https://github.com/openshift-pipelines/operator/pull/24834)

## Next Action

**Step 13 — Start stage release with latest snapshots.**

The previous stage releases used older snapshots from 2026-07-01 (and the bundle stage failed). A fresh release is needed with the 2026-07-02 snapshots.

Stage release has 3 sub-steps (IN ORDER):

1. **CORE APPLICATION:**
   - Use snapshot: `openshift-pipelines-core-1-21-20260702-114559-000-lv`
   - Update release.yaml with the snapshot
   - Apply release.yaml on Konflux cluster
   - Wait for managed pipelinerun to succeed

2. **BUNDLE APPLICATION:**
   - Trigger operator-update-images workflow:
     https://github.com/openshift-pipelines/operator/actions/workflows/operator-update-images.yaml
     Parameters: Branch: `release-v1.21.x`, Environment: `staging`
   - Review the CSV PR — verify ALL images point to stage registry (no quay.io refs)
   - Merge the PR
   - Wait for bundle image to build
   - Use latest bundle snapshot, update release.yaml, apply on Konflux
   - Wait for managed pipelinerun to succeed

3. **INDEX APPLICATION:**
   - Trigger render-olm-catalog workflow:
     https://github.com/openshift-pipelines/operator/actions/workflows/render-olm-catalog.yaml
     Parameters: Branch: `release-v1.21.x`, Environment: `staging`
   - Use latest index snapshot, update release.yaml, apply on Konflux
   - Wait for managed pipelinerun to succeed

**IMPORTANT:** Note the snapshot IDs used — you need them for production release.

## Remaining Steps

- 14 — QA handover
- 15 — Production release
- 16 — Advisory
