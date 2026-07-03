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

- If a rebuild run exists AND its `createdAt` is **after** the latest upstream sync merge time AND its `conclusion` is `success` → **DONE** (rebuilds were triggered after the latest changes)
- If a rebuild run exists AND its `createdAt` is **after** the latest upstream sync merge time BUT its `conclusion` is `failure` → **FAILED** (rebuilds were attempted but failed — see diagnostics below)
- If a rebuild run exists but its `createdAt` is **before** the latest upstream sync merge time → **ACTION NEEDED** (changes landed after the last rebuild)
- If no rebuild run exists for this version → **ACTION NEEDED**

If the rebuild ran after the sync but **failed**, inspect the failed run:

```bash
gh run view ${FAILED_RUN_ID} --repo openshift-pipelines/hack --log-failed 2>&1 | head -80
```

Look for the failing repo and error message (e.g. `FAILED operator (release-v1.21.x)`, `Error: 409 Conflict`). If multiple recent rebuild runs all failed:

```bash
gh run list --repo openshift-pipelines/hack \
  --workflow=trigger-image-rebuilds.yaml \
  --limit 20 \
  --json databaseId,status,conclusion,createdAt,displayTitle
```

Filter for runs targeting `${MAJOR_MINOR}` and list all failed runs with their timestamps as absolute local time.

Report the failure with:
- Run ID and link: `https://github.com/openshift-pipelines/hack/actions/runs/${FAILED_RUN_ID}`
- Creation time as absolute local time
- The failing repo name extracted from the log
- The error message

If rebuilds failed:
```
NEXT ACTION: Rebuild failed — use alternative methods.

The trigger-image-rebuilds workflow failed ${FAIL_COUNT} time(s) for ${FAILING_REPO}.
Error: ${ERROR_MESSAGE}
Run: https://github.com/openshift-pipelines/hack/actions/runs/${FAILED_RUN_ID}

Option 1: Retry the workflow targeting just the failing repo:
  Go to: https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml
  Click "Run workflow" with:
    - version: ${MAJOR_MINOR}
    - repo: ${FAILING_REPO}

Option 2: Push an empty commit to trigger Konflux directly:
  git clone --branch ${RELEASE_BRANCH} --depth 1 \
    git@github.com:openshift-pipelines/${FAILING_REPO}.git /tmp/${FAILING_REPO}-rebuild
  cd /tmp/${FAILING_REPO}-rebuild
  git commit --allow-empty -m "chore: trigger rebuild for ${VERSION}"
  git push origin ${RELEASE_BRANCH}

Option 3: Rerun a PipelineRun from the Konflux UI:
  1. Go to: https://konflux-ui.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com
  2. Navigate to the application (e.g. openshift-pipelines-core-${MM_DASHED})
  3. Find the most recent push PipelineRun for ${FAILING_REPO} components
  4. Click the kebab menu → "Rerun"

After rebuilding, wait for new snapshots and re-run the checklist.
```

If not done (rebuild before sync or no rebuild):
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
- **Mismatch:** snapshot revision != HEAD → build may be **stale** (see operator cycle check below)
- **Split:** same repo has multiple revisions in the snapshot → some components rebuilt, others didn't

Report a table with linked SHAs:
```
| Downstream Repo | HEAD SHA | Snapshot SHA | Status |
|-----------------|----------|-------------|--------|
| openshift-pipelines/tektoncd-pipeline | [f9d8e9cb9c85](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/FULL) | [f9d8e9cb9c85](https://github.com/openshift-pipelines/tektoncd-pipeline/commit/FULL) | CURRENT |
| openshift-pipelines/operator | [b628089785fc](https://github.com/openshift-pipelines/operator/commit/FULL) | [6b5649fd1234](https://github.com/openshift-pipelines/operator/commit/FULL) | CYCLE OK |
```

### Operator cycle check

The **operator repo** is a special case. After operator components build, nudge PRs land image SHA updates back into the operator repo, moving HEAD forward. This means the core snapshot will always be behind the operator HEAD — this is expected behavior, not staleness.

**Why this is NOT an infinite loop:** All push pipelines in the operator repo use CEL path filters (`on-cel-expression`) that restrict which file changes trigger builds:

| Pipeline type | Triggers only on changes to |
|--------------|----------------------------|
| operator, proxy, webhook | `upstream/***`, `.konflux/patches/***`, `.konflux/rpms/***`, specific Dockerfiles, `.tekton/*-push.yaml` |
| bundle | `.konflux/olm-catalog/bundle/***`, `bundle.Dockerfile`, `.tekton/*-bundle-push.yaml` |
| index | `.konflux/olm-catalog/index/***`, index Dockerfiles, `.tekton/*-index-*-push.yaml` |

Nudge PRs only change `project.yaml` — which is **not in any path filter**. So when nudge PRs merge, no operator/proxy/webhook push pipeline triggers. The build flow is one-directional:

1. Source changes (`upstream/`, `.konflux/patches/`) → operator/proxy/webhook builds trigger
2. Builds complete → nudge PRs update `project.yaml` → **no rebuild** (path filter excludes it)
3. CSV generation updates `.konflux/olm-catalog/bundle/` → **bundle build triggers**
4. Catalog JSON generation updates `.konflux/olm-catalog/index/` → **index builds trigger**

To verify the path filters, check the push PipelineRun definitions:
```bash
for f in operator proxy webhook bundle; do
  echo "=== ${f} push ==="
  gh api repos/openshift-pipelines/operator/contents/.tekton/operator-${MM_DASHED}-${f}-push.yaml?ref=${RELEASE_BRANCH} \
    --jq '.content' | base64 -d | grep -A5 'on-cel-expression'
done
```

When the operator is the only mismatched repo, check whether the commits between the snapshot SHA and HEAD are all automated:

```bash
gh api repos/openshift-pipelines/operator/compare/${SNAPSHOT_SHA}...${HEAD_SHA} \
  --jq '.commits[] | "\(.sha[:12]) \(.commit.message | split("\n")[0])"'
```

Classify each commit:
- **Automated:** message starts with `chore(deps):`, `[bot:`, or `chore: auto-generate` — these are nudge PRs, CSV regeneration, or catalog JSON updates
- **Manual:** any other commit — a human-authored change that needs a rebuild

Also check for open nudge PRs:
```bash
gh pr list --repo openshift-pipelines/operator \
  --base ${RELEASE_BRANCH} --label konflux-nudge \
  --state open --json number,title
```

Decision:
- If ALL commits are automated AND no nudge PRs are open → **CYCLE OK** (the build cycle completed normally; validate via bundle snapshot in step 10d)
- If ALL commits are automated BUT nudge PRs are still open → **IN PROGRESS** (cycle is still running; wait for nudge PRs to merge)
- If any non-automated commits exist → genuinely **STALE** (needs rebuild)

Report the commits between snapshot and HEAD as a list with linked SHAs:
```
Commits between snapshot and HEAD on operator (${COUNT} commits, all automated):
- [SHORT_SHA](URL) — chore(deps): update operator-1-21-proxy to ...
- [SHORT_SHA](URL) — chore: auto-generate OCP catalog JSONs
```

When the operator is CYCLE OK, the bundle snapshot (step 10d) becomes the primary validation gate.

### Non-operator stale repos

If any components other than the operator are STALE or SPLIT, generate the next-action guidance **in dependency order**. Triggering image rebuilds before downstream repos are in sync wastes builds — the resulting images will be stale as soon as pending changes land.

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
    - repo: ${STALE_REPO} (target just the stale repo, or leave empty for all)

  If the workflow fails (e.g. 409 errors), use alternative methods:
  a) Push an empty commit to the stale repo's release branch:
     git clone --branch ${RELEASE_BRANCH} --depth 1 \
       git@github.com:openshift-pipelines/${STALE_REPO}.git /tmp/${STALE_REPO}-rebuild
     cd /tmp/${STALE_REPO}-rebuild
     git commit --allow-empty -m "chore: trigger rebuild for ${VERSION}"
     git push origin ${RELEASE_BRANCH}
  b) Rerun from the Konflux UI:
     Go to: https://konflux-ui.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com
     Navigate to the application → find the latest push PipelineRun → Rerun
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

**When the operator is CYCLE OK (step 10c), the bundle snapshot is the primary validation gate:**

- If bundle revision == operator HEAD → **DONE** (full build cycle complete — all component images, nudge PRs, CSV, catalog JSONs, and bundle are current)
- If bundle revision != operator HEAD → check what's between them:
  - If only the bundle nudge commit and catalog JSON commit are missing → **IN PROGRESS** (bundle rebuild is pending or in flight; wait for it)
  - If open nudge PRs exist → **IN PROGRESS** (cycle still running; nudge PRs must merge before the bundle rebuilds)
  - Otherwise → **STALE** (bundle needs rebuilding)

Check for open nudge PRs that indicate the cycle is still running:
```bash
gh pr list --repo openshift-pipelines/operator \
  --base ${RELEASE_BRANCH} --label konflux-nudge \
  --state open --json number,title
```

If in progress:
```
Step 10 — Build cycle in progress.

The operator components have rebuilt but the cycle is still completing:
- Open nudge PRs: ${OPEN_NUDGE_COUNT} (must merge before bundle rebuilds)
- Bundle snapshot revision: [SHORT](URL) (behind operator HEAD [SHORT](URL))

Wait for nudge PRs to merge and a new bundle snapshot to appear, then re-run the checklist.
```

### 10e. Diagnose missing nudge PRs after builds

If builds succeeded on the Konflux UI but no new snapshots or nudge PRs appeared, run this diagnostic checklist:

**10e.1: Verify the build was a push pipeline (not PR pipeline)**

Only push pipelines create snapshots and trigger nudges. PR pipelines run tests but do not produce release artifacts.

Check recent PipelineRuns — push pipelines contain `on-push` in the name; PR pipelines contain `on-pull-request`:
```bash
oc get pipelineruns -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "operator-${MM_DASHED}.*on-push" | tail -5
```

**10e.2: Check the build's source commit (Reference SHA)**

Even if a push pipeline ran, it may have built from the wrong commit:
```bash
oc get pipelinerun ${RECENT_PUSH_PLR} -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -o jsonpath='{.spec.params[?(@.name=="revision")].value}' 2>/dev/null
```

If the PipelineRun's revision matches the old snapshot SHA (not HEAD), the build ran from a stale commit. A new push event is needed — use the empty commit method from step 9b.

**10e.3: Check for new snapshots**

```bash
oc get snapshots -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "core-${MM_DASHED}" | tail -5
```

Compare the latest snapshot's creation timestamp against the rebuild time. If no new snapshot appeared after the rebuild, the integration tests may still be running or may have failed.

**10e.4: Check integration test results**

Nudge PRs only fire after a successful integration test on the snapshot. Check the latest snapshot's test annotations:
```bash
oc get snapshot ${LATEST_SNAPSHOT} -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -o jsonpath='{.metadata.annotations}' 2>/dev/null | python3 -c "
import sys, json
annotations = json.load(sys.stdin)
for k, v in sorted(annotations.items()):
    if 'test' in k.lower() or 'result' in k.lower():
        print(f'{k}: {v}')
"
```

Report findings:
```
DIAGNOSTIC: Builds succeeded but no nudge PRs appeared.

- Pipeline type: ${PUSH_OR_PR} (only on-push creates nudges)
- Build source commit: ${PLR_REVISION} (expected: ${HEAD_SHA})
- Latest snapshot: ${LATEST_SNAPSHOT} (created: ${SNAPSHOT_TIME})
- Integration tests: ${PASS_OR_FAIL}

Root cause: ${DIAGNOSIS}
```

Possible root causes:
- **PR pipeline, not push:** Wait for the PR to merge, or push an empty commit to trigger a push pipeline.
- **Built from stale commit:** Push an empty commit to trigger a new build from HEAD.
- **Integration test failed:** Check test logs and fix issues.
- **Tests still running:** Wait for integration tests to complete.

**Return:** Status for steps 9, 9b, and 10 with operator project.yaml version check, SHA comparison table (with operator cycle status), bundle snapshot validation, and any diagnostic findings. After producing results, read and follow the report update instructions in `REPORT.md` to write/update `reports/checklist-${VERSION}.md`.
