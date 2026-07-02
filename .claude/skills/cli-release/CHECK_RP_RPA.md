# Check ReleasePlan and ReleasePlanAdmission

Verify that CDN-specific ReleasePlan and ReleasePlanAdmission resources exist in the konflux-release-data GitLab repo.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `GITLAB_URL`, `GITLAB_TOKEN`, `TZ_FMT`

**Formatting:** All GitLab MR numbers must be rendered as markdown links. All timestamps as absolute local time.

**Constraints:**
- GitLab: **READ-ONLY** (`GET` requests only, never `POST`/`PUT`/`DELETE`)

## Step 5: Check RP and RPA for CDN release

CDN releases require dedicated ReleasePlanAdmission and ReleasePlan resources, separate from the image-based ones used in the main release checklist.

### Check ReleasePlanAdmission

List RPA files in GitLab and filter for CDN-related entries:

```bash
curl -s --header "PRIVATE-TOKEN: $GITLAB_TOKEN" \
  "$GITLAB_URL/api/v4/projects/releng%2Fkonflux-release-data/repository/tree?path=config/kflux-prd-rh02.0fk9.p1/product/ReleasePlanAdmission/tekton-ecosystem&ref=main&per_page=100" \
  2>/dev/null | python3 -c "
import sys, json
data = json.load(sys.stdin)
if isinstance(data, list):
    matches = [f['name'] for f in data if '${MM_DASHED}' in f['name'] and 'cdn' in f['name']]
    for m in sorted(matches):
        print(m)
    if not matches:
        print('NO_CDN_RPA_FOUND')
"
```

Expected CDN RPAs:
- `openshift-pipelines-${MM_DASHED}-core-cdn-prod.yaml`
- `openshift-pipelines-${MM_DASHED}-core-cdn-stage.yaml`

### Check ReleasePlan

List ReleasePlan files in the tenant config:

```bash
curl -s --header "PRIVATE-TOKEN: $GITLAB_TOKEN" \
  "$GITLAB_URL/api/v4/projects/releng%2Fkonflux-release-data/repository/tree?path=tenants-config/cluster/kflux-prd-rh02/tenants/tekton-ecosystem-tenant&ref=main&per_page=100" \
  2>/dev/null | python3 -c "
import sys, json
data = json.load(sys.stdin)
if isinstance(data, list):
    matches = [f['name'] for f in data if '${MM_DASHED}' in f['name'] and 'cdn' in f['name']]
    for m in sorted(matches):
        print(m)
    if not matches:
        print('NO_CDN_RP_FOUND')
"
```

Also check the auto-generated ReleasePlan files:

```bash
curl -s --header "PRIVATE-TOKEN: $GITLAB_TOKEN" \
  "$GITLAB_URL/api/v4/projects/releng%2Fkonflux-release-data/repository/tree?path=tenants-config/auto-generated/cluster/kflux-prd-rh02/tenants/tekton-ecosystem-tenant&ref=main&per_page=100" \
  2>/dev/null | python3 -c "
import sys, json
data = json.load(sys.stdin)
if isinstance(data, list):
    matches = [f['name'] for f in data if '${MM_DASHED}' in f['name'] and 'cdn' in f['name']]
    for m in sorted(matches):
        print(m)
    if not matches:
        print('NO_AUTO_GEN_CDN_RP_FOUND')
"
```

### Present results

Show a table of found RP/RPA resources:

```markdown
| Resource | Type | File | Status |
|----------|------|------|--------|
| CDN stage RPA | ReleasePlanAdmission | openshift-pipelines-${MM_DASHED}-core-cdn-stage.yaml | FOUND |
| CDN prod RPA | ReleasePlanAdmission | openshift-pipelines-${MM_DASHED}-core-cdn-prod.yaml | FOUND |
| CDN stage RP | ReleasePlan | ... | FOUND |
| CDN prod RP | ReleasePlan | ... | FOUND |
```

If ALL required RP/RPA files exist → **DONE**.

If any are missing:
```
NEXT ACTION: Create CDN ReleasePlan and ReleasePlanAdmission in konflux-release-data.

Missing resources:
${LIST_OF_MISSING_RESOURCES}

ReleasePlan and ReleasePlanAdmission are not auto-generated for CDN releases.
Copy from a previous version and update the version numbers.

Reference MR: https://gitlab.cee.redhat.com/releng/konflux-release-data/-/merge_requests/14754/diffs

Repository: https://gitlab.cee.redhat.com/releng/konflux-release-data

IMPORTANT: The product version configuration (step 4) must be merged first.
ReleasePlanAdmission cannot reference a version that is not defined.
```

If `GITLAB_TOKEN` is not set, mark as **MANUAL** and print the instructions above.

**Return:** Status for step 5 with RP/RPA details.
