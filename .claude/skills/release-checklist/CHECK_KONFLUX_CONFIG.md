# Check Konflux Configuration

Verify Konflux config PRs are merged, configuration is applied on the cluster, ReleasePlanAdmission exists, and Pyxis config is in place.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `GITLAB_URL`, `GITLAB_TOKEN`, `TZ_FMT`

**Report ownership:** Steps 3, 3b, 4, 5. Detail section: `## Konflux Config Details`.

**Formatting:** All PR numbers must be rendered as markdown links `[#NUM](https://github.com/openshift-pipelines/hack/pull/NUM)`. All commit SHAs as `[SHORT](https://github.com/OWNER/REPO/commit/FULL)`. All timestamps as absolute local time.

**Early stop:** If any step in this group requires action, return immediately with the blocking step's status and action details. Do not check subsequent steps.

**Constraints:**
- Konflux cluster: **READ-ONLY** (`oc get`/`kubectl get` only, never `apply`/`create`/`delete`)
- GitLab: **READ-ONLY** (`GET` requests only, never `POST`/`PUT`/`DELETE`)

## Step 3: Check Konflux configuration PR

After the hack release-manager PR merges, the `generate-konflux` workflow runs automatically (triggered by push to main). It creates:

1. **A PR in the hack repo itself** for Konflux config updates:
   - Branch: `actions/update/hack-update-konflux-main-${MAJOR_MINOR}`
   - Labels: `automated`, `hack`
   - Title: `[bot:${MAJOR_MINOR}] Update Generated Konflux Config`

2. **PRs in each downstream component repo** for PAC/pipeline configuration:
   - Branch: `hack/openshift-pipelines-core/${RELEASE_BRANCH}`
   - Labels: `hack`, `automated`
   - Title: `[bot:hack/openshift-pipelines-core/${RELEASE_BRANCH}] update konflux configuration`

Check the hack Konflux config PR:
```bash
gh pr list --repo openshift-pipelines/hack \
  --head "actions/update/hack-update-konflux-main-${MAJOR_MINOR}" \
  --state all --limit 3 \
  --json number,title,state,mergedAt
```

- If merged → configuration was generated. Show merge time as absolute local time.
- If open → needs to be merged (the `merge-hack-pull-requests` workflow runs hourly and auto-merges PRs with the `hack` label)

## Step 3b: Verify Konflux configuration on the cluster matches the hack repo

If `KONFLUX_SERVER` and `KONFLUX_TOKEN` are set, verify the configuration on the cluster matches what the hack repo generated. Use **READ-ONLY** commands only.

### 3b.1: Compare Applications

Get the list of expected applications from the hack repo:
```bash
gh api repos/openshift-pipelines/hack/contents/.konflux/openshift-pipelines/${MM_DASHED} \
  --jq '[.[] | select(.type == "dir") | .name] | sort[]'
```
This gives the expected application names (e.g. `openshift-pipelines-core`, `openshift-pipelines-bundle`, `openshift-pipelines-index-4.17`).

Get the actual applications on the cluster:
```bash
oc get applications.appstudio.redhat.com -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify -o json 2>/dev/null | python3 -c "
import sys, json
data = json.load(sys.stdin)
apps = sorted([item['metadata']['name'] for item in data['items'] if '${MM_DASHED}' in item['metadata']['name']])
for a in apps:
    print(a)
"
```

Compare: each hack repo directory `openshift-pipelines-X-Y.Z` should have a corresponding cluster application `openshift-pipelines-X-Y-Z-${MM_DASHED}` (dots in OCP versions become dashes). Report any missing applications.

### 3b.2: Compare Components per Application

For each application in the hack repo, list the expected component subdirectories:
```bash
gh api repos/openshift-pipelines/hack/contents/.konflux/openshift-pipelines/${MM_DASHED}/${APP_DIR} \
  --jq '[.[] | select(.type == "dir") | .name] | sort[]'
```

Get the actual components on the cluster for that application:
```bash
oc get components -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify -o json 2>/dev/null | python3 -c "
import sys, json
data = json.load(sys.stdin)
comps = sorted([item['metadata']['name'] for item in data['items']
    if item['spec'].get('application','') == '${CLUSTER_APP_NAME}'])
for c in comps:
    print(c)
"
```

Verify:
- Every component directory in the hack repo has a matching component on the cluster
- Every component on the cluster is assigned to the correct application (not misassigned to a different app)
- No applications exist on the cluster with zero components (empty apps indicate config was partially applied)

### 3b.3: Check independently-versioned components

Some components (e.g. `syncer-service`, `multicluster-proxy-aae`) live under separate top-level `.konflux/` directories with their own version schemes. Also, some components in the release config have a `min-version` field in their repo config that excludes them from older releases.

For each component in the release config that has no match in `.konflux/openshift-pipelines/${MM_DASHED}/`:
```bash
gh api repos/openshift-pipelines/hack/contents/config/downstream/repos/${COMPONENT}.yaml \
  --jq '.content' | base64 -d
```

Check for `min-version` — if `min-version` > `MAJOR_MINOR`, the component is correctly excluded.
Check if the component has its own `.konflux/${COMPONENT}/` directory with a separate version — verify that application exists on the cluster.

### 3b.4: Report comparison results

Present results as a table:

```
| Application | Hack Repo Components | Cluster Components | Status |
|-------------|---------------------|--------------------|--------|
| core        | 16 dirs             | 38 (multi-image)   | OK     |
| bundle      | 1 dir               | 1                  | OK     |
| index-4.17  | 1 dir               | 0 (misassigned)    | DRIFT  |
```

- **OK** — cluster matches hack repo
- **DRIFT** — component exists but is assigned to the wrong application, or application has no components
- **MISSING** — expected application or component does not exist on the cluster

If all OK → **DONE**.
If DRIFT or MISSING found → **ACTION NEEDED**:

```
NEXT ACTION: Re-apply Konflux configuration on the cluster.

The cluster configuration has drifted from the hack repo. Affected:
${LIST_OF_DRIFT_OR_MISSING_ITEMS}

1. Get cluster token: https://oauth-openshift.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com/oauth/token/display
2. Login:
   oc login --token=<TOKEN> --server=https://api.kflux-prd-rh02.0fk9.p1.openshiftapps.com:6443
3. From a local checkout of the hack repo (main branch), apply the drifted application(s).

   **Important:** The `.konflux/` directories contain `role.yaml` (RoleBindings) and `service-account.yaml` (ServiceAccounts) which require elevated RBAC permissions. These are one-time setup resources already applied by cluster admins. Exclude them to avoid Forbidden errors:

   Apply a specific drifted application:
   find .konflux/openshift-pipelines/${MM_DASHED}/${DRIFTED_APP_DIR}/ -name '*.yaml' \
     ! -name 'role.yaml' ! -name 'service-account.yaml' \
     -exec kubectl apply -f {} +

   Or re-apply all config for this release at once:
   find .konflux/openshift-pipelines/${MM_DASHED}/ -name '*.yaml' \
     ! -name 'role.yaml' ! -name 'service-account.yaml' \
     -exec kubectl apply -f {} +

   Hack repo paths:
   - Per-application: .konflux/openshift-pipelines/${MM_DASHED}/${APP_DIR}/
   - All for release: .konflux/openshift-pipelines/${MM_DASHED}/
```

Always include the specific `.konflux/` paths for each drifted or missing application in the action guidance.

If `KONFLUX_SERVER` or `KONFLUX_TOKEN` are not set, mark as MANUAL and print the instructions above.

**IMPORTANT:** Do NOT merge the component PRs in downstream repos before applying the Konflux config on the cluster — the pipelines triggered by those PRs will fail without the config.

## Step 4: Check ReleasePlanAdmission

For stage/production releases, RPA must exist in konflux-release-data. Check using GitLab API (**READ-ONLY**):

```bash
curl -s --header "PRIVATE-TOKEN: $GITLAB_TOKEN" \
  "$GITLAB_URL/api/v4/projects/releng%2Fkonflux-release-data/repository/tree?path=config/kflux-prd-rh02.0fk9.p1/product/ReleasePlanAdmission/tekton-ecosystem&ref=main&per_page=100" \
  2>/dev/null | python3 -c "
import sys, json
data = json.load(sys.stdin)
if isinstance(data, list):
    matches = [f['name'] for f in data if '${MM_DASHED}' in f['name']]
    for m in matches:
        print(m)
    if not matches:
        print('NO_RPA_FOUND')
"
```

Expected RPAs for a release: core (stage+prod), bundle (stage+prod), fbc (stage+prod), cdn (stage+prod).

If RPA files exist for this version → **DONE**.
If not found:
```
NEXT ACTION: Copy ReleasePlanAdmission from hack repo to konflux-release-data.

The RPAs are generated in the hack repo under .konflux/. Copy them to:
https://gitlab.cee.redhat.com/releng/konflux-release-data

Example MR: https://gitlab.cee.redhat.com/releng/konflux-release-data/-/merge_requests/10083/diffs
```

If `GITLAB_TOKEN` is not set, try the GitHub mirror:
```bash
gh search code "openshift-pipelines" --repo redhat-appstudio/konflux-release-data --json path --limit 20 2>/dev/null | grep -i "${MAJOR_MINOR}" || true
```

## Step 5: Check Pyxis configuration (if new components)

Only needed if new component images are being shipped with this release. Check GitLab (**READ-ONLY**):

```bash
curl -s --header "PRIVATE-TOKEN: $GITLAB_TOKEN" \
  "$GITLAB_URL/api/v4/projects/releng%2Fpyxis-repo-configs/search?scope=blobs&search=openshift-pipelines&per_page=5" \
  2>/dev/null
```

If config exists under `products/openshift-pipelines/` → **DONE** (for patch releases, existing config is sufficient).
If not found or new components need adding:
```
MANUAL (if applicable): Copy Pyxis configuration to:
https://gitlab.cee.redhat.com/releng/pyxis-repo-configs/

Check hack repo for generated Pyxis config.
```

**Return:** Status for steps 3, 3b, 4, 5 with details. After producing results, read and follow the report update instructions in `REPORT.md` to write/update `reports/checklist-${VERSION}.md`.
