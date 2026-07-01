# Check Build Validation (SHA Comparison)

Validate that Konflux builds originate from the correct (latest) commit SHA on each downstream repo's release branch. A build from a stale commit means the component needs rebuilding.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`

**Constraints:**
- Konflux cluster: **READ-ONLY** (`oc get`/`kubectl get` only, never `apply`/`create`/`delete`)

## Step 9: Validate Konflux builds against downstream repo commits

This is the critical validation step. Builds are only valid if they originate from the correct (latest) commit SHA on each downstream repo's release branch.

### 9a. Get the latest core snapshot from Konflux

```bash
LATEST_SNAPSHOT=$(oc get snapshots -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "core-${MM_DASHED}" | tail -1 | awk '{print $1}')
```

Report the snapshot name and its creation time as absolute local time.

### 9b. Extract per-component (repo, revision) from the snapshot

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

### 9c. Compare against downstream repo HEADs

For each unique repo URL in the snapshot, get the HEAD SHA of the release branch:
```bash
gh api repos/${REPO}/commits/${RELEASE_BRANCH} --jq '.sha'
```

Compare:
- **Match:** snapshot revision == HEAD → build is current
- **Mismatch:** snapshot revision != HEAD → build is **stale**, needs rebuild
- **Split:** same repo has multiple revisions in the snapshot → some components rebuilt, others didn't

Report a table:
```
| Downstream Repo | HEAD SHA | Snapshot SHA | Status |
|-----------------|----------|-------------|--------|
| openshift-pipelines/tektoncd-pipeline | f9d8e9cb... | f9d8e9cb... | CURRENT |
| openshift-pipelines/operator | b6280897... | 6b5649fd... | STALE |
| openshift-pipelines/pac-downstream | 217ca7d8... | 217ca7d8... (3), adb1dea6... (1) | SPLIT |
```

If any components are STALE or SPLIT:
```
ACTION NEEDED: Trigger image rebuilds for stale components.

Go to: https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml
Parameters:
  - version: ${MAJOR_MINOR}
  - repo: ${STALE_REPO_NAME} (or leave empty to rebuild all)

After rebuilds complete, new nudge PRs will update the operator's project.yaml.
```

### 9d. Check bundle snapshot separately

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

**Return:** Status for step 9 with the SHA comparison table, list of stale/split repos, and bundle snapshot status.
