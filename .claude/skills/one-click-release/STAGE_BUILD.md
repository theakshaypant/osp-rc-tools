# Stage 2: Build

Process release PRs, verify Konflux snapshots are current, merge nudge PRs, render OLM catalog, verify FBC builds, and set code freeze. Steps run sequentially — stop at the first step that requires action (unless the user approves execution).

**Pipeline source:** `one-click-release/pipelines/openshift-pipelines-release.yaml` — tasks `process-pull-requests`, `wait-for-core-snapshot`, `process-nudge-prs`, `wait-for-fbc-snapshot`. Also covers OLM catalog rendering and FBC index build verification from `one-click-release/README.md` Build Stage.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`, `REPORT_BASE`, `REPORT_TIMESTAMP`

**Constraints:**
- Konflux cluster: **READ-ONLY** for verify commands (`oc get`/`kubectl get` only)
- Execute commands require explicit user approval before running

**Formatting:**
- PR links: `[#NUM](https://github.com/OWNER/REPO/pull/NUM)`
- SHA links: `[SHORT](https://github.com/OWNER/REPO/commit/FULL)` (12-char short)
- Timestamps: absolute local time with timezone (e.g. `2026-07-08 14:30 IST`)

---

## Step 2.1: Process release PRs

Mirrors pipeline task `osp-github-pr-processor` with label `hack`. The pipeline searches for open PRs with release-related labels across all `openshift-pipelines` repos on the release branch, checks their CI status, and merges them when green.

### Verify

Search for open release PRs on the release branch:
```bash
gh search prs --owner openshift-pipelines \
  --base "${RELEASE_BRANCH}" \
  --state open \
  --json repository,url,title,labels \
  "label:hack,upstream,automated" 2>/dev/null
```

For each open PR, check its mergeable and CI status:
```bash
gh pr view "${PR_URL}" --json mergeable,mergeStateStatus,statusCheckRollup
```

Classify each PR:
- `mergeable: CONFLICTING` or `mergeStateStatus: DIRTY` → **CONFLICT** (cannot auto-merge)
- `mergeStateStatus: BEHIND` → **BEHIND** (needs rebase)
- CI checks with `bucket: fail` → **CI FAILING**
- CI checks with `bucket: pending` → **CI PENDING**
- All checks pass and mergeable → **READY TO MERGE**

**Collect links:** Report each PR as `repo [#NUM](URL)` with its status.

**Expected when DONE:** No open PRs found — all release PRs have been merged.

### If not done — Execute (requires approval)

For each PR that is **READY TO MERGE**:
```bash
gh pr edit "${PR_URL}" --add-label "lgtm,approved,one-click-release"
gh pr review --approve "${PR_URL}"
gh pr merge "${PR_URL}" -d -r --auto
```

For PRs that are **BEHIND**:
```bash
gh pr update-branch "${PR_URL}" --rebase
```

For PRs with **CI FAILING**: report the failing checks and skip. The user must investigate manually.

For PRs with **CI PENDING**: report that checks are still running. Re-check after a few minutes.

After merging, re-verify to confirm no open PRs remain.

---

## Step 2.2: Wait for core snapshot

Mirrors pipeline task `wait-for-snapshot` (APP=core) and script `wait-for-latest-snapshot.sh`. Verifies the latest Konflux core snapshot has the correct commit SHA for each component repo.

**Requires:** `KONFLUX_SERVER` and `KONFLUX_TOKEN`. If missing, SKIP this step.

### Verify

Get the latest core snapshot:
```bash
CORE_APP=$(oc get applications.appstudio.redhat.com -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify -o json 2>/dev/null \
  | python3 -c "
import sys, json
data = json.load(sys.stdin)
mm = '${MM_DASHED}'
apps = [i['metadata']['name'] for i in data['items'] if 'core' in i['metadata']['name'] and mm in i['metadata']['name']]
print(apps[0] if apps else '')
")

LATEST_SNAPSHOT=$(oc get snapshots -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -l "pac.test.appstudio.openshift.io/event-type=push,appstudio.openshift.io/application=${CORE_APP}" \
  --sort-by=.metadata.creationTimestamp \
  -o jsonpath='{.items[*].metadata.name}' 2>/dev/null | awk '{print $NF}')

echo "Latest core snapshot: ${LATEST_SNAPSHOT}"
```

Extract components and compare against release branch HEAD:
```bash
oc get snapshot ${LATEST_SNAPSHOT} -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -o jsonpath='{.spec.components}' 2>/dev/null | python3 -c "
import sys, json, subprocess

components = json.load(sys.stdin)
repos = {}
for c in components:
    src = c.get('source', {}).get('git', {})
    url = src.get('url', '').rstrip('.git')
    rev = src.get('revision', '?')
    repo_path = url.replace('https://github.com/', '')
    repo_name = repo_path.split('/')[-1]

    # Skip operator for core snapshot (has nudge commits ahead)
    if repo_name == 'operator':
        continue

    if repo_path not in repos:
        repos[repo_path] = {'revision': rev, 'components': []}
    repos[repo_path]['components'].append(c['name'])

for repo_path, info in sorted(repos.items()):
    print(f'{repo_path} {info[\"revision\"][:12]} ({len(info[\"components\"])} components)')
"
```

For each unique repo in the snapshot, compare against HEAD:
```bash
REMOTE_SHA=$(git ls-remote "https://github.com/${REPO_PATH}.git" "refs/heads/${RELEASE_BRANCH}" | awk '{print $1}')
```

Classify:
- Snapshot revision == HEAD → **CURRENT**
- Snapshot revision != HEAD → **STALE** (needs rebuild)
- Same repo has multiple revisions → **SPLIT** (partial rebuild)

Report a table:
```
| Repo | Snapshot SHA | HEAD SHA | Status |
|------|-------------|----------|--------|
```

**Collect links:** Report the snapshot name and creation time. For stale repos, note which need rebuilds.

**Expected when DONE:** All non-operator repos are CURRENT.

### If not done — Execute (requires approval)

For each stale repo, trigger a rebuild by pushing a placeholder commit (mirrors `rebuild_repo()` from the pipeline):
```bash
TMP_DIR=$(mktemp -d)
gh repo clone "openshift-pipelines/${STALE_REPO}" "${TMP_DIR}" -- -b "${RELEASE_BRANCH}" --depth 1 --quiet
cd "${TMP_DIR}"
mkdir -p .konflux/patches/
echo "Forced multi-component rebuild at $(date)" > .konflux/patches/.placeholder
git config user.name "One Click Release Bot"
git config user.email "one-click-release-bot@redhat.com"
git add .
git commit -m "One Click Release: build all konflux components" --quiet
git push origin "${RELEASE_BRANCH}" --quiet
cd -
rm -rf "${TMP_DIR}"
```

After pushing, wait for Konflux pipelines to complete. Check pipeline status:
```bash
gh api "repos/openshift-pipelines/${STALE_REPO}/commits/${NEW_SHA}/check-runs" \
  --jq '[.check_runs[] | select(.status == "queued" or .status == "in_progress")] | length'
```

When pipelines finish, re-verify the snapshot to confirm all repos are CURRENT.

---

## Step 2.3: Process nudge PRs

Mirrors pipeline task `osp-github-pr-processor` with label `konflux-nudge`. After core components build, Konflux creates nudge PRs in the operator repo to update image SHAs in `project.yaml`.

### Verify

Search for open nudge PRs:
```bash
gh search prs --owner openshift-pipelines \
  --base "${RELEASE_BRANCH}" \
  --state open \
  --json repository,url,title,labels \
  "label:konflux-nudge" 2>/dev/null
```

For each open nudge PR, check its CI status:
```bash
gh pr view "${PR_URL}" --json mergeable,mergeStateStatus,statusCheckRollup
```

Also check recently merged nudge PRs to report progress:
```bash
gh pr list --repo openshift-pipelines/operator \
  --base ${RELEASE_BRANCH} \
  --label konflux-nudge \
  --state merged --limit 20 \
  --json number,title,mergedAt
```

**Collect links:** Report each open nudge PR as `operator [#NUM](URL)` with its status. Include count of recently merged nudge PRs.

**Expected when DONE:** No open nudge PRs remain.

### If not done — Execute (requires approval)

For each nudge PR that is ready to merge (all CI passing, mergeable):
```bash
gh pr edit "${PR_URL}" --add-label "lgtm,approved,one-click-release"
gh pr review --approve "${PR_URL}"
gh pr merge "${PR_URL}" -d -r --auto
```

For nudge PRs with failing CI:
```bash
gh pr checks "${PR_URL}" --json name,bucket,link
```

Report the failing checks. If the failure is a known transient issue, suggest `/retest`:
```bash
gh pr comment "${PR_URL}" --body "/retest"
```

After merging, re-verify to confirm no open nudge PRs remain.

---

## Step 2.4: OLM catalog render

After nudge PRs merge and CSV is regenerated, the `render-olm-catalog` workflow renders catalog JSONs for each supported OCP version. These catalog JSON updates trigger index image builds. This step verifies the render workflow has succeeded.

**Pipeline source:** `one-click-release/README.md` Build Stage — "Update OLM config with latest bundle image, Render OLM, Build Index Images."

### Verify

Check recent `render-olm-catalog` workflow runs on the operator repo:
```bash
gh run list --repo openshift-pipelines/operator \
  --workflow=render-olm-catalog.yaml \
  --limit 10 \
  --json databaseId,headBranch,status,conclusion,createdAt,displayTitle,url
```

Filter for runs on `${RELEASE_BRANCH}` or dispatched runs (which target all release branches). Show `createdAt` timestamps as absolute local time.

Also check for auto-generated catalog JSON commits on the release branch:
```bash
gh api repos/openshift-pipelines/operator/commits?sha=${RELEASE_BRANCH}\&per_page=10 \
  --jq '.[] | select(.commit.message | test("catalog|render|OCP catalog")) | "\(.sha[:12]) \(.commit.message | split("\n")[0])"'
```

**Collect links:** Report the most recent successful workflow run URL as `[render-olm-catalog](URL)`.

**Expected when DONE:** A `render-olm-catalog` run completed successfully after the latest nudge PRs merged. Catalog JSON commits are present on the release branch.

### If not done — Execute (requires approval)

Trigger the render-olm-catalog workflow:
```bash
gh workflow run render-olm-catalog.yaml \
  --repo openshift-pipelines/operator \
  -f branch=${RELEASE_BRANCH}
```

The workflow runs daily at 1 AM UTC and dispatches to all release branches. To target just this release branch, pass the branch parameter.

After triggering, wait for the workflow to complete:
```bash
gh run list --repo openshift-pipelines/operator \
  --workflow=render-olm-catalog.yaml \
  --limit 3 \
  --json databaseId,status,conclusion,createdAt
```

Re-verify once complete.

---

## Step 2.5: Wait for FBC build

After the OLM catalog is rendered, catalog JSON changes trigger index image builds via Konflux push pipelines. Verify that bundle and index snapshots exist and are current.

**Pipeline source:** `one-click-release/scripts/wait-for-fbc-build.sh` — verifies index app snapshots the same way `wait-for-latest-snapshot.sh` verifies core.

**Requires:** `KONFLUX_SERVER` and `KONFLUX_TOKEN`. If missing, SKIP this step.

### Verify

#### 2.5a: Check bundle snapshot

```bash
BUNDLE_APP=$(oc get applications.appstudio.redhat.com -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify -o json 2>/dev/null \
  | python3 -c "
import sys, json
data = json.load(sys.stdin)
mm = '${MM_DASHED}'
apps = [i['metadata']['name'] for i in data['items'] if 'bundle' in i['metadata']['name'] and mm in i['metadata']['name']]
print(apps[0] if apps else '')
")

LATEST_BUNDLE=$(oc get snapshots -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -l "pac.test.appstudio.openshift.io/event-type=push,appstudio.openshift.io/application=${BUNDLE_APP}" \
  --sort-by=.metadata.creationTimestamp \
  -o jsonpath='{.items[*].metadata.name}' 2>/dev/null | awk '{print $NF}')

echo "Latest bundle snapshot: ${LATEST_BUNDLE}"
```

Extract bundle snapshot revision and compare against operator HEAD:
```bash
BUNDLE_REV=$(oc get snapshot ${LATEST_BUNDLE} -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -o jsonpath='{.spec.components[0].source.git.revision}' 2>/dev/null)

OPERATOR_HEAD=$(git ls-remote "https://github.com/openshift-pipelines/operator.git" \
  "refs/heads/${RELEASE_BRANCH}" | awk '{print $1}')

echo "Bundle revision: ${BUNDLE_REV:0:12}"
echo "Operator HEAD:   ${OPERATOR_HEAD:0:12}"
```

If the bundle revision doesn't match operator HEAD, check the commits between them:
```bash
gh api repos/openshift-pipelines/operator/compare/${BUNDLE_REV}...${OPERATOR_HEAD} \
  --jq '.commits[] | "\(.sha[:12]) \(.commit.message | split("\n")[0])"'
```

If all intervening commits are automated (catalog/CSV/nudge), the bundle is acceptable — the gap is from expected automated activity after the bundle was built.

#### 2.5b: Check index snapshots

```bash
INDEX_APPS=$(oc get applications.appstudio.redhat.com -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify -o json 2>/dev/null \
  | python3 -c "
import sys, json
data = json.load(sys.stdin)
mm = '${MM_DASHED}'
apps = sorted([i['metadata']['name'] for i in data['items'] if 'index' in i['metadata']['name'] and mm in i['metadata']['name']])
for a in apps:
    print(a)
")
```

For each index app, get its latest snapshot and verify:
```bash
LATEST_INDEX=$(oc get snapshots -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -l "pac.test.appstudio.openshift.io/event-type=push,appstudio.openshift.io/application=${INDEX_APP}" \
  --sort-by=.metadata.creationTimestamp \
  -o jsonpath='{.items[*].metadata.name}' 2>/dev/null | awk '{print $NF}')
```

Compare each index snapshot's operator revision against HEAD. Index apps with empty config dirs (no components for this release) will have no snapshot — this is expected.

**Collect links:** Report snapshot names and creation times. Note any stale snapshots.

**Expected when DONE:** Bundle snapshot exists (automated commit gap is acceptable). All index snapshots with non-empty config are current.

### If not done — Execute (requires approval)

For stale bundle, push placeholder to `.konflux/olm-catalog/bundle/` to trigger a bundle rebuild:
```bash
TMP_DIR=$(mktemp -d)
gh repo clone "openshift-pipelines/operator" "${TMP_DIR}" -- -b "${RELEASE_BRANCH}" --depth 1 --quiet
cd "${TMP_DIR}"
mkdir -p .konflux/olm-catalog/bundle/
echo "Forced bundle rebuild at $(date)" > .konflux/olm-catalog/bundle/.placeholder
git config user.name "One Click Release Bot"
git config user.email "one-click-release-bot@redhat.com"
git add .
git commit -m "One Click Release: rebuild bundle" --quiet
git push origin "${RELEASE_BRANCH}" --quiet
cd -
rm -rf "${TMP_DIR}"
```

For stale index images, push placeholder to `.konflux/olm-catalog/index/` to trigger index rebuilds:
```bash
TMP_DIR=$(mktemp -d)
gh repo clone "openshift-pipelines/operator" "${TMP_DIR}" -- -b "${RELEASE_BRANCH}" --depth 1 --quiet
cd "${TMP_DIR}"
mkdir -p .konflux/olm-catalog/index/
echo "Forced index rebuild at $(date)" > .konflux/olm-catalog/index/.placeholder
git config user.name "One Click Release Bot"
git config user.email "one-click-release-bot@redhat.com"
git add .
git commit -m "One Click Release: rebuild index images" --quiet
git push origin "${RELEASE_BRANCH}" --quiet
cd -
rm -rf "${TMP_DIR}"
```

After pushing, wait for Konflux pipelines to complete, then re-verify snapshots.

---

## Step 2.6: Code freeze

The `code-freeze` field in the hack release config should be set to `true` when builds are complete and index images are ready for QE testing. The `create-new-patch` workflow resets `code-freeze: false` when bumping the version, so it must be manually set back to `true`.

### Verify

```bash
CODE_FREEZE=$(gh api repos/openshift-pipelines/hack/contents/config/downstream/releases/${MAJOR_MINOR}.yaml \
  --jq '.content' | base64 -d | grep 'code-freeze:' | awk '{print $2}')
echo "code-freeze: ${CODE_FREEZE}"
```

**Collect links:** If a code freeze PR exists, report it:
```bash
gh pr list --repo openshift-pipelines/hack \
  --state all --limit 5 \
  --search "code-freeze ${MAJOR_MINOR} in:title" \
  --json number,title,state,mergedAt,url
```

**Expected when DONE:** `code-freeze: true`.

### If not done — Execute (requires approval)

Create a PR to set code freeze:
```bash
TMP_DIR=$(mktemp -d)
gh repo clone "openshift-pipelines/hack" "${TMP_DIR}" -- --depth 1 --quiet
cd "${TMP_DIR}"

sed -i 's/code-freeze: false/code-freeze: true/' config/downstream/releases/${MAJOR_MINOR}.yaml

git checkout -b "release/${VERSION}/code-freeze"
git config user.name "One Click Release Bot"
git config user.email "one-click-release-bot@redhat.com"
git add config/downstream/releases/${MAJOR_MINOR}.yaml
git commit -m "[bot:${MAJOR_MINOR}] Set code freeze for ${VERSION}"
git push origin "release/${VERSION}/code-freeze" --quiet

gh pr create --repo openshift-pipelines/hack \
  --base main \
  --head "release/${VERSION}/code-freeze" \
  --title "[bot:${MAJOR_MINOR}] Set code freeze for ${VERSION}" \
  --body "Sets code-freeze: true for ${MAJOR_MINOR} after builds are complete.

This prevents the update-sources workflow from running on the release branch." \
  --label automated

cd -
rm -rf "${TMP_DIR}"
```

After the PR is created, merge it:
```bash
PR_NUMBER=$(gh pr list --repo openshift-pipelines/hack \
  --head "release/${VERSION}/code-freeze" \
  --state open --limit 1 \
  --json number --jq '.[0].number')
gh pr merge --repo openshift-pipelines/hack ${PR_NUMBER} --rebase
```

Re-verify that `code-freeze: true` after merge.

---

## Report Output

After processing all steps, write the stage report to `${REPORT_BASE}/build/report_${REPORT_TIMESTAMP}.md`.

**Report format:**

```markdown
# Build Stage Report — ${VERSION}

**Generated:** ${REPORT_TIMESTAMP}
**Release:** ${VERSION} (${MAJOR_MINOR})
**Branch:** ${RELEASE_BRANCH}

## Summary

| Step | Title | Status | Details | Links |
|------|-------|--------|---------|-------|
| 2.1 | Process release PRs | {status} | {details} | {links} |
| 2.2 | Wait for core snapshot | {status} | {details} | {links} |
| 2.3 | Process nudge PRs | {status} | {details} | {links} |
| 2.4 | OLM catalog render | {status} | {details} | {links} |
| 2.5 | Wait for FBC build | {status} | {details} | {links} |
| 2.6 | Code freeze | {status} | {details} | {links} |

## Step Details

### Step 2.1: Process release PRs
- **Status:** {DONE | ACTION NEEDED | SKIPPED}
- **Open PRs:** {count}
- **Merged PRs:** {count}
- **Details:** {per-PR status if any open}

### Step 2.2: Wait for core snapshot
- **Status:** {DONE | ACTION NEEDED | SKIPPED}
- **Snapshot:** {name} (created {timestamp})
- **SHA comparison:**

| Repo | Snapshot SHA | HEAD SHA | Status |
|------|-------------|----------|--------|
| {repo} | [{short}]({url}) | [{short}]({url}) | {CURRENT/STALE} |

### Step 2.3: Process nudge PRs
- **Status:** {DONE | ACTION NEEDED | SKIPPED}
- **Open nudge PRs:** {count}
- **Merged nudge PRs:** {count}
- **Details:** {per-PR status if any open}

### Step 2.4: OLM catalog render
- **Status:** {DONE | ACTION NEEDED | SKIPPED}
- **Latest run:** [{workflow-name}]({url}) ({status}/{conclusion}, {timestamp})
- **Catalog commits:** {count} catalog JSON commits on release branch

### Step 2.5: Wait for FBC build
- **Status:** {DONE | ACTION NEEDED | SKIPPED}
- **Bundle snapshot:** {name} ({CURRENT/STALE/automated gap OK})
- **Index snapshots:**

| Index App | Snapshot | Status |
|-----------|----------|--------|
| {app} | {name} | {CURRENT/STALE/EXPECTED (no components)} |

### Step 2.6: Code freeze
- **Status:** {DONE | ACTION NEEDED | SKIPPED}
- **code-freeze:** {true/false}
- **PR:** hack [#{number}]({url}) — {state}

{...include details for each step checked...}

## Blocking Step

{If stopped early, show which step blocked and why. Omit if all steps DONE.}
```

**Column values:**

- **Status:** `DONE`, `ACTION NEEDED`, `SKIPPED`
- **Details:** one-line summary (e.g. `0 open, 14 merged`, `all repos CURRENT`, `code-freeze: true`)
- **Links:** all links collected for that step:
  - GitHub PRs: `repo [#NUM](URL)`
  - Snapshot names (no URL — cluster resource)
  - Multiple links separated by `, `
  - No links: `—`

Write the report file and print the path to the user:
```
Report written to: reports/${MAJOR_MINOR}/${VERSION}/build/report_${REPORT_TIMESTAMP}.md
```
