# Release Checklist: 1.21.3

**Generated:** 2026-07-01 13:53 IST
**Release branch:** release-v1.21.x
**Release tag:** 1.21.3
**Code freeze:** false
**Component monitor:** https://openshift-pipelines.github.io/hack/
**Konflux UI:** https://konflux-ui.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com

| # | Step | Status | Details |
|---|------|--------|---------|
| 1 | Generate patch in hack | DONE | release-tag: 1.21.3 |
| 2 | Hack release PR merged | DONE | PR [#759](https://github.com/openshift-pipelines/hack/pull/759) merged 2026-06-01 13:24 IST |
| 3 | Konflux config PR merged | DONE | PR [#822](https://github.com/openshift-pipelines/hack/pull/822) merged 2026-06-17 20:27 IST |
| 3b | Konflux config applied | DONE | Applications and components confirmed in `tekton-ecosystem-tenant` |
| 4 | ReleasePlanAdmission | DONE | 8 RPA files on GitLab (core/bundle/fbc, stage+prod) |
| 5 | Pyxis configuration | DONE | Config exists under `products/openshift-pipelines/` |
| 6 | Component PRs merged | PARTIAL | 15/18 merged, 3 open (see below) |
| 7 | Upstream sources synced | IN PROGRESS | PR [#24522](https://github.com/openshift-pipelines/operator/pull/24522) open since 2026-06-30 09:21 IST |
| 8 | Component builds & nudges | PARTIAL | 99 merged, 1 open with failing CI (see below) |
| 9 | Build validation (SHA) | ISSUES | 4 repos stale/split — operator, tektoncd-hub, pac-downstream, tektoncd-pipeline |
| 10 | OLM config & index images | PARTIAL | olm/config.yaml only has 1.21.2 — needs 1.21.3 entry; index images releasing successfully |
| 11 | Code freeze set | NOT DONE | code-freeze: false |
| 12 | Stage release | BLOCKED | Core stage: Succeeded; **Bundle stage: FAILED** (2026-06-24) |
| 13 | QA handover | NOT DONE | No audit report |
| 14 | Production release | NOT DONE | No production releases found on Konflux |
| 15 | Advisory | NOT DONE | |

---

## Open Component PRs (Step 6)

3 Konflux configuration PRs are still open. They have the `hack` label and should auto-merge via `merge-hack-pull-requests` (runs hourly):

| Downstream Repo | PR | Component |
|-----------------|-----|-----------|
| p12n-manual-approval-gate | [#321](https://github.com/openshift-pipelines/p12n-manual-approval-gate/pull/321) | manual-approval-gate |
| pac-downstream | [#1603](https://github.com/openshift-pipelines/pac-downstream/pull/1603) | pipelines-as-code |
| tektoncd-git-clone | [#838](https://github.com/openshift-pipelines/tektoncd-git-clone/pull/838) | git-init |

Previous rounds of component PRs for these repos were already merged. These 3 should not block the release pipeline.

## Stale Nudge PR (Step 8)

PR [#24110](https://github.com/openshift-pipelines/operator/pull/24110) — `chore(deps): update operator-1-21-bundle to 200bb43` — has been open since **2026-06-25 10:47 IST**.

**CI status:** `operator-1-21-index-4.17-on-pull-request` is **failing** on Konflux; other index pipelines pass. Tide reports: "Not mergeable. Needs lgtm label."

- [Failing pipeline](https://konflux-ui.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com/ns/tekton-ecosystem-tenant/pipelinerun/operator-1-21-index-4.17-on-pull-request-8jq6g)

## Build Validation — SHA Comparison (Step 9)

Latest core snapshot: `openshift-pipelines-core-1-21-20260701-071144-000-jy` (created 2026-07-01 12:41 IST)

| Downstream Repo | HEAD SHA | Snapshot SHA | Status |
|-----------------|----------|-------------|--------|
| openshift-pipelines/operator | b628089785fc | 6b5649fdcf76 | **STALE** |
| openshift-pipelines/p12n-console-plugin | 57173fc21b34 | 57173fc21b34 | CURRENT |
| openshift-pipelines/p12n-console-plugin-pf5 | a8ef0e7a6008 | a8ef0e7a6008 | CURRENT |
| openshift-pipelines/p12n-manual-approval-gate | 547a998c61d8 | 547a998c61d8 | CURRENT |
| openshift-pipelines/p12n-opc | 4c36a275bc7e | 4c36a275bc7e | CURRENT |
| openshift-pipelines/p12n-tekton-caches | c133f78a73e8 | c133f78a73e8 | CURRENT |
| openshift-pipelines/pac-downstream | 217ca7d880e4 | 217ca7d880e4 (3), adb1dea652a0 (1) | **SPLIT** |
| openshift-pipelines/serve-tkn-cl | — | ec69a1a0746a | N/A (no release branch) |
| openshift-pipelines/tektoncd-chains | 632264e486ff | 632264e486ff | CURRENT |
| openshift-pipelines/tektoncd-cl | — | cb0260208302 | N/A (no release branch) |
| openshift-pipelines/tektoncd-git-clone | 04b1073a1f58 | 04b1073a1f58 | CURRENT |
| openshift-pipelines/tektoncd-hub | c96b01a40075 | 2f22460f7634 | **STALE** |
| openshift-pipelines/tektoncd-pipeline | f9d8e9cb9c85 | f9d8e9cb9c85 (5), 3abf98b861dd (3) | **SPLIT** |
| openshift-pipelines/tektoncd-pruner | 60d04d49e992 | 60d04d49e992 | CURRENT |
| openshift-pipelines/tektoncd-results | 7a07a991fe48 | 7a07a991fe48 | CURRENT |
| openshift-pipelines/tektoncd-triggers | ff8e22447c67 | ff8e22447c67 | CURRENT |

**4 repos need attention:**

1. **operator** — Snapshot built from `6b5649fdcf76`, HEAD is `b628089785fc`. Stale build; needs image rebuild.
2. **tektoncd-hub** — Snapshot built from `2f22460f7634`, HEAD is `c96b01a40075`. Stale build; needs image rebuild.
3. **pac-downstream** — 3 components at HEAD (`217ca7d880e4`), but controller component at older rev `adb1dea652a0`. Split build; controller needs rebuild.
4. **tektoncd-pipeline** — 5 components at HEAD (`f9d8e9cb9c85`), but 3 components (controller, workingdirinit, webhook) at older rev `3abf98b861dd`. Split build; 3 components need rebuild.

**Bundle snapshot:** `openshift-pipelines-bundle-1-21-20260701-074746-000` (created 2026-07-01 13:17 IST)
Bundle operator revision: `b628089785fc` — **matches operator HEAD** (bundle is current).

**Action:** Trigger image rebuilds for stale core components via:
https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml

## OLM Config (Step 10)

`olm/config.yaml` on `release-v1.21.x` currently lists bundles up to **1.21.2** only:

```yaml
bundles:
  - version: 1.18.0
  - version: 1.21.2
    image: quay.io/openshift-pipeline/pipelines-operator-bundle@sha256:a0ed0f...
```

The 1.21.3 bundle entry needs to be added.

**Index images** are being released successfully — latest index releases on Konflux at 2026-07-01 11:55 IST (OCP 4.19–4.22).

**render-olm-catalog** workflow: latest 1.21 run at 2026-07-01 13:22 IST (succeeded, then action_required on re-run).

## Konflux Stage Release (Step 12)

Releases in `tekton-ecosystem-tenant` namespace:

| Application | ReleasePlan | Status | Snapshot |
|------------|-------------|--------|----------|
| Core (stage) | `openshift-pipelines-core-1-21-stage-rp` | **Succeeded** | core-1-21-20260624-073901-000 |
| Core (stage) | `openshift-pipelines-core-1-21-stage-rp` | **Succeeded** | core-1-21-20260624-061626-000-mj |
| Bundle (stage) | `openshift-pipelines-bundle-1-21-stage-rp` | **FAILED** | bundle-1-21-20260624-080940-000 |

**The bundle stage release failed on 2026-06-24.** This is a blocker. A new bundle stage release needs to be created using a recent bundle snapshot (latest: `openshift-pipelines-bundle-1-21-20260701-074746-000`, created 2026-07-01 13:17 IST).

Regular (non-stage) releases continue to succeed:
- Latest core release: Succeeded at 2026-07-01 12:41 IST
- Latest bundle release: Succeeded at 2026-07-01 13:17 IST
- Index releases: Succeeded at 2026-07-01 11:55 IST (OCP 4.19–4.22)

## ReleasePlanAdmission (Step 4)

GitLab path: `config/kflux-prd-rh02.0fk9.p1/product/ReleasePlanAdmission/tekton-ecosystem/`

Files present for 1.21:
- `openshift-pipelines-1-21-core-prod.yaml`
- `openshift-pipelines-1-21-core-stage.yaml`
- `openshift-pipelines-1-21-core-cdn-prod.yaml`
- `openshift-pipelines-1-21-core-cdn-stage.yaml`
- `openshift-pipelines-1-21-bundle-prod.yaml`
- `openshift-pipelines-1-21-bundle-stage.yaml`
- `openshift-pipelines-1-21-fbc-prod.yaml`
- `openshift-pipelines-1-21-fbc-stage.yaml`

## Cluster Access Notes

The Konflux token for user `akpant` in namespace `tekton-ecosystem-tenant`:
- **Can:** `oc get applications`, `oc get releases`, `oc get snapshots`
- **Cannot:** `oc get components`, `oc get namespaces` (Forbidden)

---

## Next Actions (Priority Order)

1. **Rebuild stale/split core components** (Step 9) — 4 repos have builds from outdated commits: operator, tektoncd-hub, pac-downstream (controller), tektoncd-pipeline (3 components). Trigger rebuilds via [trigger-image-rebuilds](https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml).

2. **Investigate stale nudge PR #24110** (Step 8) — Open since 2026-06-25 10:47 IST with failing `operator-1-21-index-4.17` pipeline. Tide says "Needs lgtm label." May need the PR rebased or the pipeline issue resolved.

3. **Retry bundle stage release** (Step 12) — Failed on 2026-06-24. After core component rebuilds land, retry with a new bundle snapshot. This is the primary release blocker.

4. **Update olm/config.yaml** (Step 10) — Add 1.21.3 bundle entry when the bundle image SHA is confirmed.

5. **Set code freeze** (Step 11) — Update `code-freeze: true` in [hack release config](https://github.com/openshift-pipelines/hack/blob/main/config/downstream/releases/1.21.yaml) when ready to hand off to QE.

6. **Monitor upstream sync PR #24522** (Step 7) — Open since 2026-06-30 09:21 IST. Auto-merges when CI passes.

7. **Wait for component PRs to auto-merge** (Step 6) — 3 open PRs should merge via hourly `merge-hack-pull-requests`.

## Remaining Steps After Blockers

```
Rebuild stale core components → New nudge PRs → Retry bundle stage → OLM config update → Code freeze → QA handover → Prod release → Advisory
```
