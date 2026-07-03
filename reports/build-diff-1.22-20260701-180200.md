# Build Diff: 1.22

**Old snapshot:** openshift-pipelines-core-1-22-20260630-164711-000-pi
**New snapshot:** openshift-pipelines-core-1-22-20260701-115006-000-wd
**Generated:** 2026-07-01 18:02 IST

## Summary

| Upstream Repo | Upstream Branch | Old Head | New Head | Upstream Commits | Jiras |
|---------------|-----------------|----------|----------|------------------|-------|
| tektoncd-catalog/git-clone | main | c3498c6... | 523d1fe... | 1 (+7 dependabot) | 0 |

4 repos rebuilt with no code changes (downstream bot nudges only): p12n-manual-approval-gate, pac-downstream, tektoncd-hub, tektoncd-pipeline.

---

## tektoncd-catalog/git-clone (1 upstream commit + 7 dependabot)

**Downstream:** openshift-pipelines/tektoncd-git-clone
**Upstream branch:** main
**Upstream head:** c3498c6 → 523d1fe

| SHA | Message | Author | Date | PR | Jira |
|-----|---------|--------|------|----|------|
| 27c670b | chore: group dependabot updates and fix gomod path | Vincent Demeester | 2026-07-01 | [#137](https://github.com/tektoncd-catalog/git-clone/pull/137) | — |

*7 dependabot commits omitted (dependency bumps for actions/setup-go, imjasonh/setup-crane, ko-build/setup-ko, alpine/git, goreleaser/goreleaser-action, actions/checkout).*

---

## Downstream-only Rebuilds (no code changes)

The following repos had downstream SHA changes but **no upstream changes** and **no human downstream commits** — only bot nudge/rebuild triggers:

| Downstream Repo | Old SHA | New SHA | Cause |
|-----------------|---------|---------|-------|
| openshift-pipelines/p12n-manual-approval-gate | f25d0d7 | b9e652d | bot: update konflux configuration |
| openshift-pipelines/pac-downstream | bd2efb6 | 0837eab | bot: update konflux configuration |
| openshift-pipelines/tektoncd-hub | 410f970 | 83d350e | bot: One Click Release + Konflux nudges |
| openshift-pipelines/tektoncd-pipeline | 1774b2f | 6b33b99 | bot: upstream sync (no new upstream commits) |

---

## Unchanged Components (14)

operator, p12n-console-plugin, p12n-console-plugin-pf5, p12n-multicluster-proxy-aae, p12n-opc, p12n-syncer-service, p12n-tekton-caches, serve-tkn-cl, tekton-kueue, tektoncd-chains, tektoncd-cl, tektoncd-pruner, tektoncd-results, tektoncd-triggers
