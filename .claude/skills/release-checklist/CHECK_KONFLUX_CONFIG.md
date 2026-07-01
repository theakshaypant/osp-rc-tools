# Check Konflux Configuration

Verify Konflux config PRs are merged, configuration is applied on the cluster, ReleasePlanAdmission exists, and Pyxis config is in place.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `GITLAB_URL`, `GITLAB_TOKEN`, `TZ_FMT`

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

## Step 3b: Verify Konflux configuration is applied on the cluster

If `KONFLUX_SERVER` and `KONFLUX_TOKEN` are set, verify the configuration was applied by checking for Application and Component resources in `tekton-ecosystem-tenant`. Use **READ-ONLY** commands only:

```bash
oc get applications.appstudio.redhat.com -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "${MM_DASHED}"
```

```bash
oc get components -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "${MM_DASHED}" | head -10
```

If applications and components exist for this version → **DONE** (Konflux config applied).
If not found → **ACTION NEEDED**:

```
NEXT ACTION: Apply Konflux configuration on the cluster.

1. Get cluster token: https://oauth-openshift.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com/oauth/token/display
2. Login:
   oc login --token=<TOKEN> --server=https://api.kflux-prd-rh02.0fk9.p1.openshiftapps.com:6443
3. Apply config from the hack repo:
   kubectl apply -Rf .konflux/
```

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

**Return:** Status for steps 3, 3b, 4, 5 with details.
