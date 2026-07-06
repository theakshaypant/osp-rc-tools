# Check OLM Config and Code Freeze

Verify OLM catalog configuration, index image builds, and code freeze status.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `CODE_FREEZE` (from release config), `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`

**Report ownership:** Steps 11, 12. Detail section: `## OLM and Code Freeze Details`.

**Formatting:** All PR numbers must be rendered as markdown links. All commit SHAs as links. All timestamps as absolute local time.

**Early stop:** Check steps 11, 12 in order. If any step requires action, return immediately with the blocking step's status and action details.

**Constraints:**
- Konflux cluster: **READ-ONLY** (`oc get`/`kubectl get` only, never `apply`/`create`/`delete`)

## Step 11: Check OLM config and index images

### 11a. Check olm/config.yaml

```bash
gh api repos/openshift-pipelines/operator/contents/olm/config.yaml?ref=${RELEASE_BRANCH} \
  --jq '.content' | base64 -d
```

Parse the `bundles:` list and extract all `version:` values. The last bundle entry should be the current release version (`VERSION`). For patch releases within the same minor (e.g. `1.21.2` → `1.21.3`), the previous patch version is **replaced** — not kept alongside the new one.

- If `VERSION` is the last bundle entry → **DONE**
- If `VERSION` is **not present** → **ACTION NEEDED** — this is a **hard blocker**. Without it, `render-olm-catalog` will not render a catalog for the current release. The `update-olm-bundle.sh` script only processes bundle versions listed in this config. If the current version is missing, the render will either use a stale bundle image (which may not exist in the target registry) or skip the version entirely, producing **empty catalog.json files** that break all index builds.

If the config needs updating, create a PR on `${RELEASE_BRANCH}` that:
1. Replaces the previous patch version with `VERSION` in the `version:` field
2. Keeps the existing `image:` line unchanged (the `operator-update-images` workflow overwrites it later)
3. Removes the previous patch version entry entirely — only `VERSION` should remain (alongside any older cross-minor entries like `1.18.0`)

The `render-catalog.sh` CI check will automatically run on the PR and push additional commits that update all `catalog-template.json` files across every OCP version (v4.14+). It inserts the new version into the OLM upgrade chains and adjusts `replaces` references. It may also re-add the previous patch version entry to `olm/config.yaml` — if so, remove it again.

CI will show a `render-catalog.sh` failure for the initial commit (the image SHA is stale until the workflow updates it). This is expected and does not block merging.

```
NEXT ACTION: Update olm/config.yaml bundle version to ${VERSION}.

The olm/config.yaml on ${RELEASE_BRANCH} has version ${PREVIOUS_VERSION}
as the last bundle entry. It must be ${VERSION} for render-olm-catalog
to produce a valid catalog.

Create a PR on ${RELEASE_BRANCH} to:
1. Replace ${PREVIOUS_VERSION} with ${VERSION}, keeping the existing
   image reference (a workflow updates it later)
2. Remove the ${PREVIOUS_VERSION} entry entirely

Result should look like:
  bundles:
    - version: 1.18.0
      tag: 1.18.0
    - version: ${VERSION}
      image: <kept from ${PREVIOUS_VERSION} entry>

The render-catalog.sh CI will auto-update catalog-template.json files.
If it re-adds ${PREVIOUS_VERSION}, remove it again.

Reference: https://github.com/openshift-pipelines/operator/commit/3bc737e023a9
```

### 11b. Check render-olm-catalog workflow

The `render-olm-catalog` workflow runs daily at 1 AM UTC and dispatches to all release branches. Check recent runs:
```bash
gh run list --repo openshift-pipelines/operator \
  --workflow=render-olm-catalog.yaml \
  --limit 10 \
  --json databaseId,headBranch,status,conclusion,createdAt,displayTitle
```

Show all `createdAt` timestamps as absolute local time. A successful run means the catalog was generated and `catalog.json` was updated for each supported OCP version.

### 11c. Check index images via Konflux releases

Check for successful index releases in `tekton-ecosystem-tenant`:
```bash
oc get releases -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "index.*${MM_DASHED}" | tail -10
```

If index releases are succeeding → index images are being built. Report the most recent release time as absolute local time.

If catalog workflow succeeded and index releases exist → **DONE**.

## Step 12: Check code freeze

The `code-freeze` field in the hack release config should be set to `true` when sharing index images with QE for testing.

Note: `create-new-patch` automatically sets `code-freeze: false` when bumping the version. It must be manually set back to `true`.

Check the value from the release config fetched in setup:
- If `code-freeze: true` → **DONE**
- If `code-freeze: false` → **ACTION NEEDED**:
  ```
  NEXT ACTION: Set code freeze for ${MAJOR_MINOR}.

  Update code-freeze to true in:
  https://github.com/openshift-pipelines/hack/blob/main/config/downstream/releases/${MAJOR_MINOR}.yaml

  Create a PR to set: code-freeze: true
  ```

**Return:** Status for steps 11 and 12 with details. Include OLM config version coverage, latest workflow run time, index release status, and code freeze state. After producing results, read and follow the report update instructions in `REPORT.md` to write/update `reports/checklist-${VERSION}.md`.
