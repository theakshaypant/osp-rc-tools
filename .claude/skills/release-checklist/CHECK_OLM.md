# Check OLM Config and Code Freeze

Verify OLM catalog configuration, index image builds, and code freeze status.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `CODE_FREEZE` (from release config), `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`

**Report ownership:** Steps 11, 12. Detail section: `## OLM and Code Freeze Details`.

**Formatting:** All PR numbers must be rendered as markdown links. All commit SHAs as links. All timestamps as absolute local time.

**Early stop:** Check steps 11, 12 in order. If any step requires action, return immediately with the blocking step's status and action details.

**Constraints:**
- Konflux cluster: **READ-ONLY** (`oc get`/`kubectl get` only, never `apply`/`create`/`delete`)

## Step 11: Check OLM render and index images

Note: The olm/config.yaml bundle version check is in step 9c (CHECK_BUILDS). By this point, the bundle version should already be correct.

### 11a. Check render-olm-catalog workflow

The `render-olm-catalog` workflow runs daily at 1 AM UTC and dispatches to all release branches. Check recent runs:
```bash
gh run list --repo openshift-pipelines/operator \
  --workflow=render-olm-catalog.yaml \
  --limit 10 \
  --json databaseId,headBranch,status,conclusion,createdAt,displayTitle
```

Show all `createdAt` timestamps as absolute local time. A successful run means the catalog was generated and `catalog.json` was updated for each supported OCP version.

### 11b. Check index images via Konflux releases

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

**Return:** Status for steps 11 and 12 with details. Include latest workflow run time, index release status, and code freeze state. After producing results, read and follow the report update instructions in `REPORT.md` to write/update `reports/checklist-${VERSION}.md`.
