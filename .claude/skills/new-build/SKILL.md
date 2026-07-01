---
description: Verify build health and generate a diff report for a new build to share with QE
---

# New Build

<objective>
When a dev fix lands on a release branch and a new build needs to go to QE, this skill:
1. Finds the baseline snapshot (what QE last tested — from the last successful stage release)
2. Finds the latest snapshot (what the new build would be)
3. Checks build pipeline health (component PRs, upstream sync, nudges, SHA validation)
4. Diffs the two snapshots to show exactly what upstream commits and Jiras changed

Invoked as: `/new-build <version>` or `/new-build <version> <baseline-snapshot>` (when no stage release exists)

Example: `/new-build 1.22.4` or `/new-build 1.23.0 openshift-pipelines-core-1-23-20260628-071141-000`

**External systems** (credentials in `.env`):
- GitHub: `gh` CLI (read-only for PRs and commits)
- Konflux cluster: `KONFLUX_SERVER` + `KONFLUX_TOKEN` — **READ-ONLY** (`oc get`/`kubectl get` only)
- Jira: `JIRA_URL` + `JIRA_EMAIL` + `JIRA_TOKEN`

**Konflux namespace:** All operations target `tekton-ecosystem-tenant`.

**Jira linking rule:** Only Jiras found via the "Git Pull Request" custom field (`cf[10875]`) count. No keyword search, no body parsing, no inference.
</objective>

<process>
<step name="setup">
**Parse arguments and initialize.**

Parse the version argument (required) and optional baseline snapshot name.

```bash
source .env
```

Verify `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `JIRA_URL`, `JIRA_EMAIL`, `JIRA_TOKEN` are set.

Derive:
- `MAJOR_MINOR` — e.g. `1.22`
- `MM_DASHED` — e.g. `1-22`
- `VERSION` — e.g. `1.22.4`
- `RELEASE_BRANCH` — `release-v${MAJOR_MINOR}.x`
- `KONFLUX_NS` — always `tekton-ecosystem-tenant`

**Time formatting:**
```bash
TZ_FMT='%Y-%m-%d %H:%M %Z'
```

All timestamps MUST use absolute local time. **Never** use relative time expressions.

Fetch the release config from hack repo:
```bash
RELEASE_CFG=$(gh api repos/openshift-pipelines/hack/contents/config/downstream/releases/${MAJOR_MINOR}.yaml --jq '.content' | base64 -d)
```

Parse `COMPONENTS` list from the `branches` map.
</step>

<step name="find_baseline">
**Find the baseline snapshot — what QE currently has.**

If a baseline snapshot name was provided as second argument, use it directly. Verify it exists:
```bash
oc get snapshot ${BASELINE_SNAP} -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify -o name 2>/dev/null
```

Otherwise, auto-detect from the last successful core stage release:
```bash
oc get releases -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 \
  | grep "core-${MM_DASHED}" | grep "stage-rp " | grep "Succeeded" | tail -1
```

The snapshot name is in column 2 of the output. Extract it.

**If no succeeded stage release exists AND no baseline was provided:**
```
No baseline found. No successful core stage release exists for ${MAJOR_MINOR}.

Provide the baseline snapshot name as second argument:
  /new-build ${VERSION} <baseline-snapshot-name>

To find available snapshots:
  oc get snapshots -n tekton-ecosystem-tenant | grep "core-${MM_DASHED}"
```
Stop here.
</step>

<step name="find_latest">
**Find the latest snapshot — what the new build would be.**

```bash
LATEST_SNAP=$(oc get snapshots -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 \
  | grep "core-${MM_DASHED}" | tail -1 | awk '{print $1}')
```

If `LATEST_SNAP` equals `BASELINE_SNAP`, report:
```
No new build available. The latest snapshot is the same as the baseline.

Latest snapshot: ${LATEST_SNAP}
Baseline: ${BASELINE_SNAP}
```
Stop here.

Report both snapshot names and their creation timestamps (absolute local time).
</step>

<step name="check_health">
**Check build pipeline health.**

Read and follow the instructions in `CHECK_BUILD_HEALTH.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `COMPONENTS`, `LATEST_SNAP`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`.

Health issues are **warnings, not blockers** — the diff is always generated even if some components are stale or PRs are open.
</step>

<step name="diff">
**Diff baseline vs latest snapshot to show what changed.**

Read and follow the instructions in `../build-diff/DIFF_SNAPSHOTS.md` to identify changed repos.

Pass `SNAP_OLD=BASELINE_SNAP`, `SNAP_NEW=LATEST_SNAP`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`.

Then read and follow `../build-diff/RESOLVE_COMMITS.md` to resolve upstream commits and Jiras for each changed repo.

Pass the changed repos list, `TZ_FMT`, and `MAJOR_MINOR`.
</step>

<step name="report">
**Generate the new build report.**

Write to `reports/new-build-${VERSION}-${YYYYMMDD-HHMMSS}.md`:

```markdown
# New Build Report: ${VERSION}

**Generated:** ${timestamp}
**Baseline:** ${BASELINE_SNAP} (from stage release: ${release_name} OR user-provided)
**Latest:** ${LATEST_SNAP}
**Release branch:** ${RELEASE_BRANCH}

## Build Health

| Check | Status | Details |
|-------|--------|---------|
| Component PRs (step 6) | DONE/PARTIAL | N/M merged, K open |
| Upstream sync (step 7) | DONE/IN PROGRESS | PR status |
| Nudge PRs (step 8) | DONE/PARTIAL | N merged, K open |
| SHA validation (step 9) | OK/ISSUES | N repos stale/split |

${health_warnings_if_any}

## What Changed (upstream)

| Upstream Repo | Branch | Commits | Jiras |
|---------------|--------|---------|-------|
| tektoncd/repo | branch | N | N |

---

### upstream-repo (N commits)

**Downstream:** openshift-pipelines/repo
**Upstream branch:** release-vX.Y.x
**Upstream head:** old_sha → new_sha

| SHA | Message | Author | Date | PR | Jira |
|-----|---------|--------|------|----|------|
...

### Downstream-only changes (if any)
...

---

## Unchanged Components (N)

repo-a, repo-b, repo-c, ...
```

**Column rules:**
- **SHA**: first 7 characters
- **PR**: `[#N](url)` when found, `—` when missing
- **Jira**: `[KEY](browse-url)` when found via `cf[10875]`, `—` otherwise
- **Date**: absolute local time, date only (YYYY-MM-DD)
- **Skip dependabot commits** — count in summary line

Print the summary to the conversation after writing the report.
</step>
</process>

<success_criteria>
- [ ] Baseline snapshot auto-detected from last successful stage release, or accepted from user
- [ ] Latest snapshot found from Konflux
- [ ] Build health checks run (steps 6-9) with clear status table
- [ ] Health issues are warnings, not blockers — diff always generated
- [ ] Changed repos identified via snapshot component SHA comparison
- [ ] Upstream commits resolved via `head` file + GitHub compare API
- [ ] Jiras found via `cf[10875]` field only — no guessing
- [ ] Downstream-only changes identified separately
- [ ] All timestamps use absolute local time with timezone
- [ ] Report written to `reports/new-build-${VERSION}-${timestamp}.md`
</success_criteria>
