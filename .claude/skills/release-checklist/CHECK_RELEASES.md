# Check Stage Release, QA Handover, Production Release, and Advisory

Verify stage and production release status on Konflux, QA handover readiness, and advisory completion.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`

**Constraints:**
- Konflux cluster: **READ-ONLY** (`oc get`/`kubectl get` only, never `apply`/`create`/`delete`)

## Step 12: Check stage release

Stage release involves three sub-steps done IN ORDER: Core, Bundle, Index.

Check for stage releases in `tekton-ecosystem-tenant` (**READ-ONLY**):
```bash
oc get releases -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep -E "${MM_DASHED}.*(stage-rp)"
```

Look for releases using `*-stage-rp` release plans:
- `openshift-pipelines-core-${MM_DASHED}-stage-rp` — core stage
- `openshift-pipelines-bundle-${MM_DASHED}-stage-rp` — bundle stage

Check status: Succeeded vs Failed. A Failed stage release is a blocker — retry with a newer snapshot.

Report all release timestamps as absolute local time.

**12a. Core application (stage):**
Use the latest core snapshot. Update `release.yaml` with the snapshot and apply on Konflux cluster. Wait for managed pipelinerun to succeed.

**12b. Bundle application (stage):**
The `operator-update-images` workflow updates the CSV to point to the staging registry. It creates a PR:
- Branch: `actions/update/operator-update-images-${RELEASE_BRANCH}`
- Labels: `automated`, `lgtm`, `approved`
- Title: `[bot:${RELEASE_BRANCH}] Update generate Cluster Service Version (CSV)`

Check for this PR:
```bash
gh pr list --repo openshift-pipelines/operator \
  --head "actions/update/operator-update-images-${RELEASE_BRANCH}" \
  --state all --limit 5 \
  --json number,title,state,mergedAt
```

The workflow doesn't encode the environment in the PR title/branch — multiple runs (devel/staging/production) use the same branch. Check the workflow run inputs to distinguish:
```bash
gh run list --repo openshift-pipelines/operator \
  --workflow=operator-update-images.yaml \
  --limit 10 \
  --json databaseId,headBranch,status,conclusion,createdAt,displayTitle
```

Also check the latest bundle snapshots:
```bash
oc get snapshots -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep "bundle-${MM_DASHED}" | tail -3
```

If no stage release activity found:
```
NEXT ACTION: Start stage release.

Stage release has 3 sub-steps (IN ORDER):

1. CORE APPLICATION:
   - Use the latest core snapshot
   - Update release.yaml with the selected snapshot
   - Apply release.yaml on Konflux cluster
   - Wait for managed pipelinerun to succeed

2. BUNDLE APPLICATION:
   - Trigger operator-update-images workflow:
     https://github.com/openshift-pipelines/operator/actions/workflows/operator-update-images.yaml
     Parameters:
       Branch: ${RELEASE_BRANCH}
       Environment: staging
   - Review the CSV PR — verify ALL images point to stage registry (no quay.io refs)
   - Merge the PR
   - Wait for bundle image to build
   - Use latest bundle snapshot, update release.yaml, apply on Konflux
   - Wait for managed pipelinerun to succeed

3. INDEX APPLICATION:
   - Trigger render-olm-catalog workflow:
     https://github.com/openshift-pipelines/operator/actions/workflows/render-olm-catalog.yaml
     Parameters:
       Branch: ${RELEASE_BRANCH}
       Environment: staging
   - Use latest index snapshot, update release.yaml, apply on Konflux
   - Wait for managed pipelinerun to succeed

IMPORTANT: Note the snapshot IDs used — you need them for production release.
```

## Step 13: QA handover (stage)

When handing over the stage build to QA:
- Attach the release-test report
- Share index images with the QE team

Check if the audit report exists and has been enriched:
```bash
ls reports/audit-release-${VERSION}.md 2>/dev/null
```

If the report exists → **DONE** (report ready to share).
If not → suggest running `/audit-release` first.

```
NEXT ACTION: Hand over stage build to QA.

1. Ensure the audit report is generated:
   - reports/release_${VERSION}.json (Go tool output)
   - reports/audit-release-${VERSION}.md (enriched report from /audit-release)

2. Share with QE:
   - Stage index images
   - Release test results
   - Audit report with commit/Jira analysis
```

## Step 14: Check production release

Production release follows the same 3-step pattern as stage.

Check for production releases in `tekton-ecosystem-tenant` (**READ-ONLY**):
```bash
oc get releases -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep -E "${MM_DASHED}.*(prod-rp|cdn-prod)"
```

Report all release timestamps as absolute local time.

**IMPORTANT differences from stage:**
- **Core:** REUSE the same core snapshot from stage (do NOT use a different one)
- **Bundle:** Generate a NEW bundle (do NOT reuse the stage bundle snapshot)
  - Trigger `operator-update-images` with `Environment: production`
  - Verify the CSV PR — ALL images must point to production registry
- **Index:** Trigger `render-olm-catalog` with `Environment: production`, use latest index snapshot

If there was ANY change to `project.yaml` since stage, follow the bundle release steps from scratch.

Check for production CSV PRs:
```bash
gh pr list --repo openshift-pipelines/operator \
  --head "actions/update/operator-update-images-${RELEASE_BRANCH}" \
  --state all --limit 5 \
  --json number,title,state,mergedAt,createdAt
```

If no production activity:
```
NEXT ACTION: Start production release.

Same 3 sub-steps as stage, but:
- Core: REUSE the same core snapshot used for stage
- Bundle: Trigger operator-update-images with Environment: production
  - Verify CSV PR — ALL images must point to production registry
- Index: Trigger render-olm-catalog with Environment: production
```

## Step 15: Send advisory

This is a manual step after production release is complete and QA has verified the build.

```
MANUAL: Send advisory for ${VERSION}.

Coordinate with the release team to publish the advisory.
```

**Return:** Status for steps 12-15 with details. Include Konflux release status table (Application, ReleasePlan, Status, Timestamp), CSV PR status, and production readiness.
