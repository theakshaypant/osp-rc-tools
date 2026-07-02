# Check Build Validation (SHA Comparison)

Validate that Konflux builds originate from the correct (latest) commit SHA on each downstream repo's release branch. A build from a stale commit means the component needs rebuilding.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`

**Report ownership:** Steps 9, 9b, 10. Detail section: `## Build Validation Details`.

**Formatting:** All commit SHAs in the comparison table must be rendered as markdown links `[SHORT](https://github.com/OWNER/REPO/commit/FULL)` using the first 12 characters for SHORT. Both the HEAD SHA and Snapshot SHA columns should contain links. All timestamps as absolute local time.

**Constraints:**
- Konflux cluster: **READ-ONLY** (`oc get`/`kubectl get` only, never `apply`/`create`/`delete`)

## Step 9: Verify operator project.yaml version

The operator's `project.yaml` on the release branch contains the bundle version under `versions.current`. This **must** match `VERSION` before the final image rebuild — otherwise the operator image ships with the wrong version string.

The `create-new-patch` workflow only bumps `release-tag` in the hack repo config. It does NOT update the operator's `project.yaml`. This is a separate step.

```bash
gh api repos/openshift-pipelines/operator/contents/project.yaml?ref=${RELEASE_BRANCH} \
  --jq '.content' | base64 -d | head -5
```

Check the `versions.current` field:
- If `current: ${VERSION}` → **DONE**
- If `current` shows a different version (e.g. previous patch) → **ACTION NEEDED**

If not done:
```
NEXT ACTION: Update operator project.yaml version.

The operator's project.yaml on ${RELEASE_BRANCH} shows current: ${CURRENT_VERSION},
but it must be ${VERSION} before images are built.

Create a PR on ${RELEASE_BRANCH} to update project.yaml:
  versions:
    current: ${VERSION}
    previous: ${PREVIOUS_VERSION}

This must be done BEFORE triggering the final image rebuild.
```

**This step must be completed before build validation (step 10).** If the version is wrong, do not proceed to build validation — any images built now would have the wrong version.

## Step 9b: Check if image rebuilds were triggered after latest changes

Before validating SHAs, check whether the `trigger-image-rebuilds` workflow was run after the latest upstream sync and operator version update landed. If no rebuild was triggered, builds are guaranteed stale — skip SHA validation and guide the user to trigger rebuilds.

### 9b.1: Get the latest upstream sync merge time

```bash
gh pr list --repo openshift-pipelines/operator \
  --head "actions/update/sources-${RELEASE_BRANCH}" \
  --state merged --limit 1 \
  --json mergedAt --jq '.[0].mergedAt'
```

### 9b.2: Get the latest trigger-image-rebuilds run for this version

```bash
gh run list --repo openshift-pipelines/hack \
  --workflow=trigger-image-rebuilds.yaml \
  --limit 20 \
  --json databaseId,status,conclusion,createdAt,displayTitle
```

Filter for runs that mention `${MAJOR_MINOR}` in the display title or inputs. Get the most recent one's `createdAt` timestamp.

### 9b.3: Compare timestamps

- If a rebuild run exists AND its `createdAt` is **after** the latest upstream sync merge time → **DONE** (rebuilds were triggered after the latest changes)
- If a rebuild run exists but its `createdAt` is **before** the latest upstream sync merge time → **ACTION NEEDED** (changes landed after the last rebuild)
- If no rebuild run exists for this version → **ACTION NEEDED**

If not done:
```
NEXT ACTION: Trigger image rebuilds.

The latest upstream sync ([#PR_NUMBER](URL)) merged at ${SYNC_MERGE_TIME},
but the most recent image rebuild for ${MAJOR_MINOR} was at ${REBUILD_TIME} (before the sync).

Trigger a rebuild:
  Go to: https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml
  Click "Run workflow" with:
    - version: ${MAJOR_MINOR}
    - repo: (leave empty to rebuild all)

Wait for the rebuild to complete, then re-run the checklist.
```

If done, proceed to step 10.

## Step 10: Validate Konflux builds against downstream repo commits

This is the critical validation step. Builds are only valid if they originate from the correct (latest) commit SHA on each downstream repo's release branch.

### 10a. Get the latest core snapshot from Konflux

```bash
LATEST_SNAPSHOT=$(oc get snapshots -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "core-${MM_DASHED}" | tail -1 | awk '{print $1}')
```

Report the snapshot name and its creation time as absolute local time.

### 10b. Extract per-component (repo, revision) from the snapshot

```bash
oc get snapshot ${LATEST_SNAPSHOT} -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify -o jsonpath='{.spec.components}' 2>/dev/null | python3 -c "
import sys, json
components = json.load(sys.stdin)
repos = {}
for c in components:
    src = c.get('source', {}).get('git', {})
    url = src.get('url', '').rstrip('.git').replace('https://github.com/', '')
    rev = src.get('revision', '?')
    if url not in repos:
        repos[url] = {}
    if rev not in repos[url]:
        repos[url][rev] = []
    repos[url][rev].append(c['name'])
for repo, revs in sorted(repos.items()):
    for rev, comps in revs.items():
        print(f'{repo} {rev[:12]} ({len(comps)} components)')
"
```

### 10c. Compare against downstream repo HEADs

For each unique repo URL in the snapshot, get the HEAD SHA of the release branch:
```bash
gh api repos/${REPO}/commits/${RELEASE_BRANCH} --jq '.sha'
```

Compare:
- **Match:** snapshot revision == HEAD → build is current
- **Mismatch:** snapshot revision != HEAD → build is **stale**, needs rebuild
- **Split:** same repo has multiple revisions in the snapshot → some components rebuilt, others didn't

Report a table with linked SHAs:
```
| Downstream Repo | HEAD SHA | Snapshot SHA | Status |
|-----------------|----------|-------------|--------|
| openshift-pipelines/tektoncd-pipeline | [f9d8e9cb9c85](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/FULL) | [f9d8e9cb9c85](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/FULL) | CURRENT |
| openshift-pipelines/operator | [b628089785fc](https://github.com/openshift-pipelines/operator/commit/FULL) | [6b5649fd1234](https://github.com/openshift-pipelines/operator/commit/FULL) | STALE |
```

If any components are STALE or SPLIT, generate the next-action guidance **in dependency order**. Triggering image rebuilds before downstream repos are in sync wastes builds — the resulting images will be stale as soon as pending changes land.

Check for upstream blockers first:
1. **Open update-sources PRs** on the operator (step 7) — if one is open, it must merge first to land upstream sync.
2. **Stuck nudge PRs** (step 8) — if any nudge PR has failing CI, it must be fixed before rebuilds.
3. **Pending downstream repo changes** — for each STALE/SPLIT repo, check what commits exist between the snapshot revision and HEAD. If they are Konflux config updates, Renovate digest updates, or upstream syncs, these must build and their nudge PRs must merge into the operator before triggering rebuilds.
4. **Verify operator project.yaml version** (step 9) — `versions.current` must match `VERSION` before the final rebuild.
5. **Trigger image rebuilds** — only after steps 1–4 are resolved and the operator HEAD is stable.
6. **Verify new snapshots** — wait for new core + bundle snapshots where all SHAs match, then re-run the checklist.

```
ACTION NEEDED: Sync downstream repos, then rebuild.

Step 1: Merge any open update-sources PR (upstream sync must land first)
Step 2: Fix any stuck nudge PRs with failing CI
Step 3: Wait for pending downstream repo changes to build and nudge PRs to merge
Step 4: Verify operator project.yaml version matches ${VERSION}
Step 5: Trigger image rebuilds (only after operator HEAD is stable):
  Go to: https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml
  Parameters:
    - version: ${MAJOR_MINOR}
    - repo: (leave empty to rebuild all)
Step 6: Verify new core + bundle snapshots match all downstream HEADs
```

### 10d. Check bundle snapshot separately

```bash
LATEST_BUNDLE=$(oc get snapshots -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "bundle-${MM_DASHED}" | tail -1 | awk '{print $1}')
```

```bash
oc get snapshot ${LATEST_BUNDLE} -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify -o jsonpath='{.spec.components[0].source.git.revision}' 2>/dev/null
```

Compare the bundle snapshot's operator revision against the operator repo HEAD. If they differ, the bundle was built from a stale operator commit. Report the bundle snapshot name and its creation time as absolute local time.

**Return:** Status for steps 9 and 10 with operator project.yaml version check, SHA comparison table, list of stale/split repos, and bundle snapshot status. After producing results, read and follow the report update instructions in `REPORT.md` to write/update `reports/checklist-${VERSION}.md`.
