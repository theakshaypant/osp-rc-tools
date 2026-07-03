# Build Diff: 1.21

**Old snapshot:** openshift-pipelines-core-1-21-20260530-122116-000 (1.21.2 era)
**New snapshot:** openshift-pipelines-core-1-21-20260702-114559-000-lv (1.21.3)
**Generated:** 2026-07-02 20:02 IST

## Summary

| Upstream Repo | Upstream Branch | Old Head | New Head | Upstream Commits | Jiras |
|---------------|-----------------|----------|----------|------------------|-------|
| openshift-pipelines/console-plugin | release-v1.23.x | 3459458... | 2216162... | 145 (+38 bot) | 85 |
| tektoncd-catalog/git-clone | main | 5df1573... | 523d1fe... | 35 (+21 dependabot) | 0 |
| tektoncd/pipelines-as-code | release-v0.39.x | 8b79f2d... | 6acde9d... | 16 | 0 |
| tektoncd/operator | release-v0.78.x | fe2392a... | 9bcf791... | 11 (+8 dependabot) | 0 |
| tektoncd/pipeline | release-v1.6.x | 47cc14d... | b54dcda... | 11 (+13 dependabot) | 0 |
| tektoncd/results | release-v0.17.x | 8d158f8... | ae090bc... | 10 | 2 |
| openshift-pipelines/hub | release-v1.23.12 | 974a106... | f7d1f0b... | 9 | 0 |
| openshift-pipelines/tekton-caches | release-v0.3.x | 4cdb9fa... | 768eb9a... | 9 | 0 |
| tektoncd/cli | release-v0.43.2 | 9374f62... | db2b9a1... | 7 | 4 |
| tektoncd/chains | release-v0.26.x | bee7519... | 50dacd2... | 5 (+3 dependabot) | 1 |
| openshift-pipelines/manual-approval-gate | release-v0.7.0 | 47e1ce9... | f9ccb26... | 3 | 2 |
| tektoncd/pruner | release-v0.3.x | 3a1779e... | 0715ff1... | 3 | 1 |
| openshift-pipelines/opc | — | 37d7717... | b705f0f... | 2 | 0 |

**Total:** 266 upstream commits (non-bot), 83 dependabot/bot, 95 Jiras

---

## openshift-pipelines/console-plugin (183 commits, 145 non-bot)

**Downstream:** openshift-pipelines/p12n-console-plugin
**Upstream branch:** release-v1.23.x
**Upstream head:** 3459458 → 2216162

Large delta — covers ~7 months of development (Nov 2025 – Jul 2026). 145 non-bot commits, 38 bot/merge commits.

**85 linked Jiras:** SRVKP-8292, SRVKP-8974, SRVKP-9273, SRVKP-9389, SRVKP-9390, SRVKP-9391, SRVKP-9392, SRVKP-9397, SRVKP-9427, SRVKP-9428, SRVKP-9436, SRVKP-9450, SRVKP-9470, SRVKP-9591, SRVKP-9605, SRVKP-9607, SRVKP-9625, SRVKP-9626, SRVKP-9627, SRVKP-9675, SRVKP-9680, SRVKP-9700, SRVKP-9709, SRVKP-9710, SRVKP-9711, SRVKP-9712, SRVKP-9713, SRVKP-9785, SRVKP-9786, SRVKP-9787, SRVKP-9788, SRVKP-9789, SRVKP-9918, SRVKP-9983, SRVKP-10032, SRVKP-10035, SRVKP-10037, SRVKP-10058, SRVKP-10069, SRVKP-10072, SRVKP-10073, SRVKP-10076, SRVKP-10077, SRVKP-10078, SRVKP-10080, SRVKP-10306, SRVKP-10310, SRVKP-10407, SRVKP-10439, SRVKP-10470, SRVKP-10743, SRVKP-10785, SRVKP-10786, SRVKP-10807, SRVKP-10808, SRVKP-10871, SRVKP-10953, SRVKP-10961, SRVKP-11001, SRVKP-11003, SRVKP-11123, SRVKP-11167, SRVKP-11190, SRVKP-11192, SRVKP-11379, SRVKP-11383, SRVKP-11384, SRVKP-11386, SRVKP-11391, SRVKP-11393, SRVKP-11394, SRVKP-11395, SRVKP-11398, SRVKP-11399, SRVKP-11400, SRVKP-11401, SRVKP-11565, SRVKP-11566, SRVKP-11698, SRVKP-11708, SRVKP-11711, SRVKP-11715, SRVKP-11951, SRVKP-12111, SRVKP-12265

---

## tektoncd-catalog/git-clone (56 commits, 35 non-dependabot)

**Downstream:** openshift-pipelines/tektoncd-git-clone
**Upstream branch:** main
**Upstream head:** 5df1573 → 523d1fe

35 non-dependabot commits + 21 dependabot. No linked Jiras in commit messages.

---

## tektoncd/pipelines-as-code (16 commits)

**Downstream:** openshift-pipelines/pac-downstream
**Upstream branch:** release-v0.39.x
**Upstream head:** 8b79f2d → 6acde9d

3 release/security PRs covering all 16 commits:

| PR | Title | Commits |
|----|-------|---------|
| [#2730](https://github.com/tektoncd/pipelines-as-code/pull/2730) | [DNM] Release v0.39.6 | 5 |
| [#2762](https://github.com/tektoncd/pipelines-as-code/pull/2762) | [DNM] fix(security): backport v0.48.0 fixes to v0.39.x | 4 |
| [#2772](https://github.com/tektoncd/pipelines-as-code/pull/2772) | DNM: Release v0.39.7 | 1 |

1 commit (d1f9603) not associated with a PR.

No linked Jiras.

---

## tektoncd/operator (11 upstream commits + 8 dependabot)

**Downstream:** openshift-pipelines/operator
**Upstream branch:** release-v0.78.x
**Upstream head:** fe2392a → 9bcf791

11 non-dependabot upstream commits + 8 dependabot. No linked Jiras in commit messages.

---

## tektoncd/pipeline (11 upstream commits + 13 dependabot)

**Downstream:** openshift-pipelines/tektoncd-pipeline
**Upstream branch:** release-v1.6.x
**Upstream head:** 47cc14d → b54dcda

11 non-dependabot upstream commits + 13 dependabot. No linked Jiras in commit messages.

---

## tektoncd/results (10 upstream commits)

**Downstream:** openshift-pipelines/tektoncd-results
**Upstream branch:** release-v0.17.x
**Upstream head:** 8d158f8 → ae090bc

10 non-dependabot upstream commits. 2 linked Jiras:
- [SRVKP-12038](https://issues.redhat.com/browse/SRVKP-12038)
- [SRVKP-11380](https://issues.redhat.com/browse/SRVKP-11380)

---

## openshift-pipelines/hub (9 upstream commits)

**Downstream:** openshift-pipelines/tektoncd-hub
**Upstream branch:** release-v1.23.12
**Upstream head:** 974a106 → f7d1f0b

9 upstream commits. No linked Jiras.

---

## openshift-pipelines/tekton-caches (9 upstream commits)

**Downstream:** openshift-pipelines/p12n-tekton-caches
**Upstream branch:** release-v0.3.x
**Upstream head:** 4cdb9fa → 768eb9a

9 upstream commits. No linked Jiras.

---

## tektoncd/cli (7 upstream commits)

**Downstream:** openshift-pipelines/tektoncd-cli
**Upstream branch:** release-v0.43.2
**Upstream head:** 9374f62 → db2b9a1

7 upstream commits. 4 linked Jiras:
- [SRVKP-11504](https://issues.redhat.com/browse/SRVKP-11504)
- [SRVKP-12038](https://issues.redhat.com/browse/SRVKP-12038)
- [SRVKP-11739](https://issues.redhat.com/browse/SRVKP-11739)
- [SRVKP-11669](https://issues.redhat.com/browse/SRVKP-11669)

---

## tektoncd/chains (5 upstream commits + 3 dependabot)

**Downstream:** openshift-pipelines/tektoncd-chains
**Upstream branch:** release-v0.26.x
**Upstream head:** bee7519 → 50dacd2

5 non-dependabot upstream commits + 3 dependabot. 1 linked Jira:
- [SRVKP-12344](https://issues.redhat.com/browse/SRVKP-12344)

---

## openshift-pipelines/manual-approval-gate (3 upstream commits)

**Downstream:** openshift-pipelines/p12n-manual-approval-gate
**Upstream branch:** release-v0.7.0
**Upstream head:** 47e1ce9 → f9ccb26

All 3 commits are CVE fixes:

| PR | Title | Jiras |
|----|-------|-------|
| [#980](https://github.com/openshift-pipelines/manual-approval-gate/pull/980) | [Release-v0.7.0] Fix CVE-2026-33211: bump tektoncd/pipeline to v1.0.2 | — |
| [#988](https://github.com/openshift-pipelines/manual-approval-gate/pull/988) | fix(cve): CVE-2026-33186 - gRPC-Go authorization bypass | [SRVKP-11741](https://issues.redhat.com/browse/SRVKP-11741), [SRVKP-11742](https://issues.redhat.com/browse/SRVKP-11742) |
| [#1000](https://github.com/openshift-pipelines/manual-approval-gate/pull/1000) | fix(cve): CVE-2026-39821 + CVE-2026-33814 - Go 1.24.0→1.25.11, x/net bump | — |

---

## tektoncd/pruner (3 upstream commits)

**Downstream:** openshift-pipelines/tektoncd-pruner
**Upstream branch:** release-v0.3.x
**Upstream head:** 3a1779e → 0715ff1

3 upstream commits. 1 linked Jira:
- [SRVKP-12038](https://issues.redhat.com/browse/SRVKP-12038)

---

## openshift-pipelines/opc (2 upstream commits)

**Downstream:** openshift-pipelines/p12n-opc
**Upstream head:** 37d7717 → b705f0f

2 version bump commits (no Jiras):

| PR | Title |
|----|-------|
| [#438](https://github.com/openshift-pipelines/opc/pull/438) | Bump pipelines-as-code, results, cli and opc version |
| [#470](https://github.com/openshift-pipelines/opc/pull/470) | bump(deps): bump pac, pipeline and hub version |

---

## Downstream-Only Changes

Non-bot commits on downstream repos that are not upstream syncs:

| Downstream Repo | Non-Bot Commits | Notes |
|-----------------|-----------------|-------|
| openshift-pipelines/serve-tkn-cli | 3 | PRs [#254](https://github.com/openshift-pipelines/serve-tkn-cli/pull/254), [#273](https://github.com/openshift-pipelines/serve-tkn-cli/pull/273), [#308](https://github.com/openshift-pipelines/serve-tkn-cli/pull/308) |
| openshift-pipelines/p12n-console-plugin | 1 | PR [#249](https://github.com/openshift-pipelines/p12n-console-plugin/pull/249) |
| openshift-pipelines/pac-downstream | 0 | 2 "One Click Release Bot" commits (automated) |
| openshift-pipelines/tektoncd-triggers | 0 | 9 automated downstream commits only |
| openshift-pipelines/p12n-manual-approval-gate | 0 | automated only |
| openshift-pipelines/p12n-opc | 0 | automated only |
| openshift-pipelines/p12n-tekton-caches | 0 | automated only |
| openshift-pipelines/tektoncd-git-clone | 0 | automated only |

---

## Component Changes

**Added in 1.21.3:** p12n-console-plugin-pf5 (new PatternFly 5 console plugin)

**Removed in 1.21.3:** p12n-multicluster-proxy-aae, p12n-syncer-service, p12n-tekton-assist (moved to separate apps, min-version 1.22+)
