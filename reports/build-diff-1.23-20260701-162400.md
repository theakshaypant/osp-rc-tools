# Build Diff: 1.23

**Old snapshot:** `openshift-pipelines-core-1-23-20260630-064911-000`
**New snapshot:** `openshift-pipelines-core-1-23-20260701-071141-000`
**Generated:** 2026-07-01 16:24 IST

## Summary

| Upstream Repo | Branch | Old Head | New Head | Upstream Commits | Jiras |
|---------------|--------|----------|----------|------------------|-------|
| tektoncd/operator | release-v0.80.x | 5c9be32 | daffb82 | 4 (2 human, 1 CI, 1 dependabot) | 1 |
| tektoncd/pipeline | release-v1.12.x | 5a36e12 | c041e3b | 2 (dependabot only) | 0 |
| tektoncd/pruner | release-v0.4.x | (same) | (same) | 0 (downstream-only) | 0 |

---

## tektoncd/operator (4 upstream commits)

**Downstream:** openshift-pipelines/operator
**Upstream branch:** release-v0.80.x
**Upstream head:** `5c9be32` → `daffb82`

| SHA | Message | Author | Date | PR | Jira |
|-----|---------|--------|------|----|------|
| 6fa49fe | fix(manualapprovalgate): inject TLS env vars into webhook | Jawed khelil | 2026-06-26 | [#3584](https://github.com/tektoncd/operator/pull/3584) | [SRVKP-12613](https://issues.redhat.com/browse/SRVKP-12613) |
| 5025382 | refactor: simplify update-image-sha.sh script | Shiv Verma | 2026-06-21 | [#3585](https://github.com/tektoncd/operator/pull/3585) | — |
| 25d6da7 | ci(e2e): fix ko >=v0.19 SBOM push failure on plain-HTTP registry | Jawed khelil | 2026-06-30 | [#3629](https://github.com/tektoncd/operator/pull/3629) | — |

*+ 1 dependabot commit (bump github-actions group)*

**Jira details:**
- [SRVKP-12613](https://issues.redhat.com/browse/SRVKP-12613) — Manual Approval Gate webhook missing TLS environment variables — **Dev Complete** *(Jira linked via original PR [#3582](https://github.com/tektoncd/operator/pull/3582))*

### Downstream-only changes

| SHA | Message | Author | Date | PR | Jira |
|-----|---------|--------|------|----|------|
| 9c834d2 | Cleanup 4.17 | Khurram Baig | 2026-06-26 | [#24294](https://github.com/openshift-pipelines/operator/pull/24294) | — |
| 7554dee | Update min OSP version on OCP 4.21 to OSP 1.15 | Pramod Bindal | 2026-06-29 | [#24521](https://github.com/openshift-pipelines/operator/pull/24521) | — |

---

## tektoncd/pipeline (0 human upstream commits)

**Downstream:** openshift-pipelines/tektoncd-pipeline
**Upstream branch:** release-v1.12.x
**Upstream head:** `5a36e12` → `c041e3b`

*2 dependabot commits only (dependency bumps)*

**Note:** Only 2 of 8 downstream components (webhook, nop) rebuilt with the new upstream head. The remaining 6 components are still at the old revision — split build in progress.

---

## tektoncd/pruner (downstream-only)

**Downstream:** openshift-pipelines/tektoncd-pruner
**Upstream branch:** release-v0.4.x
**Upstream head:** `10ac09c` (unchanged)

No upstream changes. The downstream SHA changed because the pruner-controller component was rebuilt and now matches the pruner-webhook revision (previously split, now converged).

---

## Unchanged Components (16)

p12n-console-plugin, p12n-console-plugin-pf5, p12n-manual-approval-gate, p12n-multicluster-proxy-aae, p12n-opc, p12n-syncer-service, p12n-tekton-caches, pac-downstream, serve-tkn-cl, tekton-kueue, tektoncd-chains, tektoncd-cl, tektoncd-git-clone, tektoncd-hub, tektoncd-results, tektoncd-triggers
