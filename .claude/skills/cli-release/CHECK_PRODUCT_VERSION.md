# Check Product Version Configuration

Verify that the product version configuration exists in the konflux-release-data GitLab repo.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `IS_PATCH`, `GITLAB_URL`, `GITLAB_TOKEN`, `TZ_FMT`

**Formatting:** All GitLab MR numbers must be rendered as markdown links. All timestamps as absolute local time.

**Constraints:**
- GitLab: **READ-ONLY** (`GET` requests only, never `POST`/`PUT`/`DELETE`)

## Step 4: Check product version YAML

The product version configuration lives in the `konflux-release-data` GitLab repo at:
```
data/external/developer-portal/openshift-pipelines/${VERSION}.yaml
```

### Fetch the file

```bash
curl -s --header "PRIVATE-TOKEN: $GITLAB_TOKEN" \
  "$GITLAB_URL/api/v4/projects/releng%2Fkonflux-release-data/repository/files/data%2Fexternal%2Fdeveloper-portal%2Fopenshift-pipelines%2F${VERSION}.yaml/raw?ref=main" \
  2>/dev/null
```

### Validate the content

If the file exists, parse and verify:

```yaml
---
versionName: "${VERSION}"
ga: false
termsAndConditions: "Anonymous Download"
hidden: false
invisible: false
releaseDate: "2026-03-01"
```

Check:
- `versionName` matches `VERSION`
- `ga` — report the value (user decides correctness)
- `invisible` — for initial CDN release this should be `true` to prevent public visibility before QA sign-off. Report the current value.
- `releaseDate` — report the value

### Present results

If the file exists with correct `versionName` → **DONE**. Report all field values.

If the file does not exist (404 response):
```
NEXT ACTION: Add product version configuration to konflux-release-data.

Create a new file at:
  data/external/developer-portal/openshift-pipelines/${VERSION}.yaml

With content:
---
versionName: "${VERSION}"
ga: false
termsAndConditions: "Anonymous Download"
hidden: false
invisible: true
releaseDate: "${TODAY}"

Notes:
- ga: Set to false for patch versions. Set to true for minor releases and
  patches on the latest minor release.
- invisible: Set to true initially. Change to false after QA sign-off.
- releaseDate: Set to the actual release date.

Reference MR: https://gitlab.cee.redhat.com/releng/konflux-release-data/-/merge_requests/14753

Repository: https://gitlab.cee.redhat.com/releng/konflux-release-data

IMPORTANT: The product version PR must be merged BEFORE creating the RP/RPA
in the next step. ReleasePlanAdmission cannot reference a version that is
not yet defined.
```

**Return:** Status for step 4 with configuration details.
