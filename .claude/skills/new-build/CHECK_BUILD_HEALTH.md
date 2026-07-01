# Check Build Health

Verify the build pipeline is healthy before generating a new build diff. This reuses the logic from release-checklist steps 6-9.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `COMPONENTS` (list from release config), `LATEST_SNAP` (latest snapshot name), `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`

**Constraints:**
- Konflux cluster: **READ-ONLY** (`oc get`/`kubectl get` only)
- Health issues are **warnings, not blockers** — report them but don't stop the diff

## Check 1: Component PRs (Step 6)

Check for open component PRs with the `hack` label on downstream repos:

For each component in `COMPONENTS`, resolve its downstream repo:
```bash
gh api repos/openshift-pipelines/hack/contents/config/downstream/repos/${COMPONENT}.yaml \
  --jq '.content' | base64 -d | grep '^repo:' | awk '{print $2}'
```

Then check for open PRs:
```bash
gh pr list --repo openshift-pipelines/${REPO} \
  --head "hack/openshift-pipelines-core/${RELEASE_BRANCH}" \
  --state open --limit 3 \
  --json number,title,state,createdAt
```

Report: `N/M merged`. If any open, list them as warnings.

## Check 2: Upstream Sync (Step 7)

Check for open update-sources PRs:
```bash
gh pr list --repo openshift-pipelines/operator \
  --head "actions/update/sources-${RELEASE_BRANCH}" \
  --state all --limit 3 \
  --json number,title,state,mergedAt,createdAt
```

- Merged recently → DONE
- Open → IN PROGRESS (show creation time)
- None → check if update-sources ran recently

## Check 3: Nudge PRs (Step 8)

Check for open nudge PRs:
```bash
gh pr list --repo openshift-pipelines/operator \
  --base ${RELEASE_BRANCH} \
  --label konflux-nudge \
  --state open --limit 10 \
  --json number,title,createdAt
```

If any open, check their CI status:
```bash
gh pr checks ${PR_NUMBER} --repo openshift-pipelines/operator
```

Report: `N merged, K open`. List open ones with CI status as warnings.

## Check 4: SHA Validation (Step 9)

Compare the latest snapshot's component SHAs against downstream repo HEADs.

Extract per-component (repo, revision) from the latest snapshot:
```bash
oc get snapshot ${LATEST_SNAP} -n tekton-ecosystem-tenant \
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

For each unique repo, compare against HEAD:
```bash
gh api repos/${REPO}/commits/${RELEASE_BRANCH} --jq '.sha'
```

Report:
- **CURRENT**: snapshot revision == HEAD
- **STALE**: snapshot revision != HEAD → warning
- **SPLIT**: same repo has multiple revisions → warning

Repos without a release branch (e.g. `serve-tkn-cl`, `tektoncd-cl`) → N/A.

## Output

Return a health status table:

| Check | Status | Details |
|-------|--------|---------|
| Component PRs | DONE/PARTIAL | N/M merged, K open |
| Upstream sync | DONE/IN PROGRESS | details |
| Nudge PRs | DONE/PARTIAL | N merged, K open |
| SHA validation | OK/ISSUES | N repos stale, M split |

Plus any warnings with details (open PRs, stale repos, failing CI).
