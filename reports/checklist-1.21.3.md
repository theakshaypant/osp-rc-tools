# Release Checklist: 1.21.3

**Generated:** 2026-07-07 10:30 IST
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
| 3b | Konflux cluster config | DONE | 11/11 hack apps match cluster; core 16 dirs → 37 components, bundle 1 dir → 1 component; index-4.23 and index-5.0 have empty config dirs (no components expected for 1.21) |
| 4 | ReleasePlanAdmission | DONE | ReleasePlans confirmed on cluster (core/bundle/index × stage+prod+cdn) |
| 5 | Pyxis configuration | DONE | Patch release — existing config sufficient |
| 6 | Component PRs on release branches | DONE | 14/14 active components merged (incl. operator [#24926](https://github.com/openshift-pipelines/operator/pull/24926) merged 2026-07-06 16:39 IST) |
| 6b | OLM bundle version | DONE | olm/config.yaml last bundle entry is 1.21.3 |
| 7 | Upstream source sync | DONE | [#25135](https://github.com/openshift-pipelines/operator/pull/25135) merged 2026-07-06 18:16 IST |
| 8 | Nudge PRs (component SHA updates) | DONE | 200+ merged, 0 open (incl. bundle nudge [#25147](https://github.com/openshift-pipelines/operator/pull/25147) merged 2026-07-07 07:38 IST) |
| 9 | Operator project.yaml version | DONE | `versions.current: 1.21.3` matches target |
| 9b | Image rebuilds after sync | DONE | Sync PR merge triggered Konflux push pipelines; core snapshot created 2026-07-06 18:25 IST |
| 10 | Build validation (SHA comparison) | DONE | Core: 15/15 repos CURRENT, operator CYCLE OK. Bundle: [e169148a29ac](https://github.com/openshift-pipelines/operator/commit/e169148a29ac274cb81707b0e6a704b16521632f) + 2 automated commits to HEAD — cycle complete |
| 11 | OLM render and index images | DONE | render-olm-catalog succeeded 2026-07-07 07:38 IST; catalog JSONs auto-generated |
| 12 | Code freeze | DONE | code-freeze: true |
| 13 | Stage release | **ACTION NEEDED** | Core: Succeeded. Bundle: Succeeded. Index: 6/7 succeeded, index-4-19 outstanding |

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
| index-4.23 | 0 dirs (empty) | 0 | OK |
| index-5.0 | 0 dirs (empty) | 0 | OK |

Cluster has extra `openshift-pipelines-index-4-17-1-21` (not in hack repo — leftover from previous config).

Independently-versioned components correctly excluded from 1.21:
- syncer-service (min-version: 1.22)
- multicluster-proxy-aae (min-version: 1.22)
- tekton-kueue (min-version: 1.22)
- tekton-assist (min-version: 5.0)

**Step 4:** ReleasePlans confirmed on Konflux cluster for 1.21: core (stage+prod+cdn), bundle (stage+prod), all index apps (stage+prod). GitLab unreachable — RPA verification via cluster ReleasePlans.

**Step 5:** Pyxis config — patch release, no new components to add. GitLab unreachable for direct verification.

## Component Details

**Step 6:** 14/14 active component PRs merged on release-v1.21.x (4 components excluded via min-version > 1.21):

| Component | Repo | PR | Merged |
|-----------|------|----|--------|
| console-plugin-pf5 | p12n-console-plugin-pf5 | [#143](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/pull/143) | 2026-07-03 |
| console-plugin | p12n-console-plugin | [#318](https://github.com/openshift-pipelines/p12n-console-plugin/pull/318) | 2026-07-03 |
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
| operator | operator | [#24926](https://github.com/openshift-pipelines/operator/pull/24926) | 2026-07-06 |

**Step 6b:** olm/config.yaml on release-v1.21.x has `version: 1.21.3` as the last bundle entry — correct.

**Step 7:** Upstream source sync PR [#25135](https://github.com/openshift-pipelines/operator/pull/25135) merged 2026-07-06 18:16 IST (synced tektoncd/operator release-v0.78.x to [1f5d2204f7f7](https://github.com/openshift-pipelines/operator/commit/1f5d2204f7f79ddbdb19450440f146d405fd2c2c)).

**Step 8:** 200+ nudge PRs merged, 0 open. Bundle nudge [#25147](https://github.com/openshift-pipelines/operator/pull/25147) merged 2026-07-07 07:38 IST (updated index components with new bundle image SHA).

## Build Validation Details

**Step 9:** Operator project.yaml `versions.current: 1.21.3` — matches target.

**Step 9b:** Sync PR merge to release-v1.21.x triggered Konflux push pipelines automatically — core snapshot `openshift-pipelines-core-1-21-20260706-124700-000-m7` created 2026-07-06 18:25 IST (9 minutes after sync merge).

**Step 10:** Core snapshot `openshift-pipelines-core-1-21-20260706-124700-000-m7` (created 2026-07-06 18:25 IST).

**Core SHA comparison (15 non-operator repos all CURRENT):**

| Downstream Repo | HEAD SHA | Snapshot SHA | Status |
|-----------------|----------|-------------|--------|
| openshift-pipelines/p12n-console-plugin | [c6c8eb76ab9b](https://github.com/openshift-pipelines/p12n-console-plugin/commit/c6c8eb76ab9b2cafc5752f1446338d6fd8944548) | [c6c8eb76ab9b](https://github.com/openshift-pipelines/p12n-console-plugin/commit/c6c8eb76ab9b2cafc5752f1446338d6fd8944548) | CURRENT |
| openshift-pipelines/p12n-console-plugin-pf5 | [ba85d3c72470](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/commit/ba85d3c7247038b4e3a0d0a013eb029b300e2ba7) | [ba85d3c72470](https://github.com/openshift-pipelines/p12n-console-plugin-pf5/commit/ba85d3c7247038b4e3a0d0a013eb029b300e2ba7) | CURRENT |
| openshift-pipelines/p12n-manual-approval-gate | [e25d64a2cfc4](https://github.com/openshift-pipelines/p12n-manual-approval-gate/commit/e25d64a2cfc4a941af0111f4e02096783370f02b) | [e25d64a2cfc4](https://github.com/openshift-pipelines/p12n-manual-approval-gate/commit/e25d64a2cfc4a941af0111f4e02096783370f02b) | CURRENT |
| openshift-pipelines/p12n-opc | [938f533cf53c](https://github.com/openshift-pipelines/p12n-opc/commit/938f533cf53c62b99879ea7579d2c0658077e96b) | [938f533cf53c](https://github.com/openshift-pipelines/p12n-opc/commit/938f533cf53c62b99879ea7579d2c0658077e96b) | CURRENT |
| openshift-pipelines/p12n-tekton-caches | [a582ebcda617](https://github.com/openshift-pipelines/p12n-tekton-caches/commit/a582ebcda617b1e812ff99c90ca99ea9734077c2) | [a582ebcda617](https://github.com/openshift-pipelines/p12n-tekton-caches/commit/a582ebcda617b1e812ff99c90ca99ea9734077c2) | CURRENT |
| openshift-pipelines/pac-downstream | [cadb77a6ee6e](https://github.com/openshift-pipelines/pac-downstream/commit/cadb77a6ee6ec3f134d3fbf6ee17bad39c777cd6) | [cadb77a6ee6e](https://github.com/openshift-pipelines/pac-downstream/commit/cadb77a6ee6ec3f134d3fbf6ee17bad39c777cd6) | CURRENT |
| openshift-pipelines/serve-tkn-cli | [b01eedb194a8](https://github.com/openshift-pipelines/serve-tkn-cli/commit/b01eedb194a8ccb0ed1f5d2b3d6e3bf8d5772d19) | [b01eedb194a8](https://github.com/openshift-pipelines/serve-tkn-cli/commit/b01eedb194a8ccb0ed1f5d2b3d6e3bf8d5772d19) | CURRENT |
| openshift-pipelines/tektoncd-chains | [6149030b735e](https://github.com/openshift-pipelines/tektoncd-chains/commit/6149030b735e2e0574706eb1e480f37932eb0f05) | [6149030b735e](https://github.com/openshift-pipelines/tektoncd-chains/commit/6149030b735e2e0574706eb1e480f37932eb0f05) | CURRENT |
| openshift-pipelines/tektoncd-cli | [dd1b37cfe2d2](https://github.com/openshift-pipelines/tektoncd-cli/commit/dd1b37cfe2d2ebca8a624e0485b83d4616f3d20f) | [dd1b37cfe2d2](https://github.com/openshift-pipelines/tektoncd-cli/commit/dd1b37cfe2d2ebca8a624e0485b83d4616f3d20f) | CURRENT |
| openshift-pipelines/tektoncd-git-clone | [cca25ea35394](https://github.com/openshift-pipelines/tektoncd-git-clone/commit/cca25ea35394a0493dca1fc895334a5912308e70) | [cca25ea35394](https://github.com/openshift-pipelines/tektoncd-git-clone/commit/cca25ea35394a0493dca1fc895334a5912308e70) | CURRENT |
| openshift-pipelines/tektoncd-hub | [1109073ab8ce](https://github.com/openshift-pipelines/tektoncd-hub/commit/1109073ab8ce4c5348fd2ec7f5dca017e59cda08) | [1109073ab8ce](https://github.com/openshift-pipelines/tektoncd-hub/commit/1109073ab8ce4c5348fd2ec7f5dca017e59cda08) | CURRENT |
| openshift-pipelines/tektoncd-pipeline | [caa67fad818a](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/caa67fad818a02244b22a94ecfef3fc02d4170ea) | [caa67fad818a](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/caa67fad818a02244b22a94ecfef3fc02d4170ea) | CURRENT |
| openshift-pipelines/tektoncd-pruner | [5d09825b776b](https://github.com/openshift-pipelines/tektoncd-pruner/commit/5d09825b776b870d9bcb4f54f286cede781968bb) | [5d09825b776b](https://github.com/openshift-pipelines/tektoncd-pruner/commit/5d09825b776b870d9bcb4f54f286cede781968bb) | CURRENT |
| openshift-pipelines/tektoncd-results | [989e3e461472](https://github.com/openshift-pipelines/tektoncd-results/commit/989e3e4614721a5b1b0ee4d7b0cee14f69f9e8ac) | [989e3e461472](https://github.com/openshift-pipelines/tektoncd-results/commit/989e3e4614721a5b1b0ee4d7b0cee14f69f9e8ac) | CURRENT |
| openshift-pipelines/tektoncd-triggers | [a4018bc83f71](https://github.com/openshift-pipelines/tektoncd-triggers/commit/a4018bc83f716c553764a4c22b24fe67007a3409) | [a4018bc83f71](https://github.com/openshift-pipelines/tektoncd-triggers/commit/a4018bc83f716c553764a4c22b24fe67007a3409) | CURRENT |

**Operator cycle (CYCLE OK):**

Operator HEAD: [070d2518d5af](https://github.com/openshift-pipelines/operator/commit/070d2518d5afe7a9ce8e49eb294c254e0a37d971). Core snapshot has SPLIT revisions (operator+webhook at [41bba78bb0b5](https://github.com/openshift-pipelines/operator/commit/41bba78bb0b52f8b8c1804abd805a24bdedeca1d), proxy at [a54dd435cbb7](https://github.com/openshift-pipelines/operator/commit/a54dd435cbb7bd183a2bec4f65c48c965b537c66)).

9 commits between newer core snapshot SHA and HEAD — all automated:
- [8d18b53c7220](https://github.com/openshift-pipelines/operator/commit/8d18b53c7220) — [bot:release-v1.21.x] Update catalog
- [f42a9513a285](https://github.com/openshift-pipelines/operator/commit/f42a9513a285) — chore(deps): update operator-1-21-webhook to 2113a93
- [6d9e43fce883](https://github.com/openshift-pipelines/operator/commit/6d9e43fce883) — chore(deps): update operator-1-21-operator to f9ed912
- [b23dcee221f3](https://github.com/openshift-pipelines/operator/commit/b23dcee221f3) — chore(deps): update operator-1-21-proxy to e6e567a
- [9da6a08c71e3](https://github.com/openshift-pipelines/operator/commit/9da6a08c71e3) — update bundle should update the last version instead of adding new version (manual — but only changed scripts/config not in build path filters)
- [e3896f7c52bb](https://github.com/openshift-pipelines/operator/commit/e3896f7c52bb) — [bot:release-v1.21.x] Update catalog
- [e169148a29ac](https://github.com/openshift-pipelines/operator/commit/e169148a29ac) — Clean up annotations.yaml by removing blank line
- [c375dfa51cbd](https://github.com/openshift-pipelines/operator/commit/c375dfa51cbd) — chore(deps): update operator-1-21-bundle to cb69d29
- [070d2518d5af](https://github.com/openshift-pipelines/operator/commit/070d2518d5af) — chore: auto-generate OCP catalog JSONs

No open nudge PRs. **CYCLE OK.**

**Bundle snapshot:** `openshift-pipelines-bundle-1-21-20260707-011035-000` built from [e169148a29ac](https://github.com/openshift-pipelines/operator/commit/e169148a29ac274cb81707b0e6a704b16521632f). 2 automated commits to HEAD (bundle nudge + catalog JSONs) — neither in bundle path filters. **Cycle complete.**

## OLM and Code Freeze Details

**Step 11a — render-olm-catalog:** Latest run on `release-v1.21.x` succeeded at 2026-07-07 07:38 IST ("chore: auto-generate OCP catalog JSONs"). Catalog JSONs regenerated after bundle nudge merge.

**Step 11b — Index snapshots:** New index snapshots created for all OCP versions (4-14 through 4-22) at ~2026-07-07 07:52 IST after catalog regeneration.

**Step 12:** Code freeze is `true` in the release config.

## Release Status Details

**Step 13:** Stage release status:

| Application | ReleasePlan | Status |
|-------------|------------|--------|
| core | `core-1-21-stage-rp` | **Succeeded** |
| bundle | `bundle-1-21-stage-rp` | **Succeeded** |
| index-4-14 | `index-4-14-1-21-stage-rp` | **Succeeded** |
| index-4-16 | `index-4-16-1-21-stage-rp` | **Succeeded** |
| index-4-18 | `index-4-18-1-21-stage-rp` | **Succeeded** |
| index-4-19 | `index-4-19-1-21-stage-rp` | **Outstanding** |
| index-4-20 | `index-4-20-1-21-stage-rp` | **Succeeded** |
| index-4-21 | `index-4-21-1-21-stage-rp` | **Succeeded** |
| index-4-22 | `index-4-22-1-21-stage-rp` | **Succeeded** |

**Root cause of earlier index failures:** The initial index stage releases (20260707-020855/020856 snapshots) all failed with `fbc-fips-check-oci-ta` because `olm/config.yaml` referenced the bundle at `registry.stage.redhat.io`, which was not yet available — the index builds triggered before the bundle stage release completed. After the bundle was released to staging, new index builds passed and stage releases succeeded for all indices except 4-19.

## Next Action

**Step 13 — Complete index-4-19 stage release.**

All stage releases have succeeded except index-4-19. Apply the index-4-19 stage release using the latest index-4-19 snapshot:

```bash
oc apply -f release-1.21.3-stage-index-4-19.yaml \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  -n tekton-ecosystem-tenant
```

Once index-4-19 succeeds, step 13 is complete. Proceed to step 14 (QA handover).

## Remaining Steps

- 14 — QA handover
- 15 — Production release
- 16 — Advisory
