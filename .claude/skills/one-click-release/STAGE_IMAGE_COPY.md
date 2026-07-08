# Stage 3: Image Copy (temporary)

Copy FBC-built index images from Konflux snapshots to quay.io for QE testing. This is a temporary stage standing in for the Test stage — it makes index images accessible on quay.io so QE can test them.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`, `REPORT_BASE`, `REPORT_TIMESTAMP`

**Constraints:**
- Konflux cluster: **READ-ONLY** for verify commands (`oc get`/`kubectl get` only)
- `skopeo` must be available on the local machine
- Execute commands require explicit user approval before running

**Formatting:**
- SHA links: `[SHORT](https://github.com/OWNER/REPO/commit/FULL)` (12-char short)
- Timestamps: absolute local time with timezone (e.g. `2026-07-08 14:30 IST`)

---

## Step 3.1: Extract index image digests

Get the containerImage reference from each index snapshot. These are the FBC-built index images produced by Konflux.

**Requires:** `KONFLUX_SERVER` and `KONFLUX_TOKEN`. If missing, SKIP this step.

### Verify

List all index applications for this release:
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

For each index app, get its latest push snapshot and extract the containerImage:
```bash
LATEST_SNAPSHOT=$(oc get snapshots -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -l "pac.test.appstudio.openshift.io/event-type=push,appstudio.openshift.io/application=${INDEX_APP}" \
  --sort-by=.metadata.creationTimestamp \
  -o jsonpath='{.items[*].metadata.name}' 2>/dev/null | awk '{print $NF}')

# Extract containerImage from snapshot
oc get snapshot ${LATEST_SNAPSHOT} -n ${KONFLUX_NS} \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -o jsonpath='{.spec.components[0].containerImage}' 2>/dev/null
```

Index apps with no snapshot (no components for this release, e.g. index-4-17) or empty config (e.g. index-4-23, index-5-0) are skipped.

Derive the OCP version from the app name to construct the quay.io target tag:
- `openshift-pipelines-index-4-14-1-21` → OCP version `4.14` → `quay.io/openshift-pipeline/pipelines-index-4.14:v${VERSION}-stage`

Report a table:
```
| Index App | Snapshot | Container Image | OCP Version | Quay Target |
|-----------|----------|-----------------|-------------|-------------|
```

**Collect links:** Report snapshot names and containerImage references.

**Expected when DONE:** All index apps with snapshots have containerImage references extracted.

### If not done

If a snapshot exists but has no containerImage, the FBC build may not have completed. Wait for the build pipeline to finish (check step 2.5).

---

## Step 3.2: Copy index images to quay.io

Copy each index image from the Konflux build registry to quay.io for QE access.

**Requires:** `skopeo` installed locally. If not available, generate a copy script instead.

### Verify

Check if the images already exist on quay.io:
```bash
skopeo inspect --no-tags docker://quay.io/openshift-pipeline/pipelines-index-${OCP_VERSION}:v${VERSION}-stage 2>/dev/null
```

If the image exists and the digest matches the source → **DONE** for that index.

Report a table:
```
| OCP Version | Quay Tag | Status |
|-------------|----------|--------|
```

**Expected when DONE:** All index images exist on quay.io with correct digests.

### If not done — Execute (requires approval)

For each index image that needs copying, generate and run skopeo copy commands:
```bash
skopeo copy --all \
  docker://${CONTAINER_IMAGE} \
  docker://quay.io/openshift-pipeline/pipelines-index-${OCP_VERSION}:v${VERSION}-stage \
  --preserve-digests
```

If `skopeo` is not available, generate a shell script at `scripts/copy-index-images-${VERSION}-stage.sh`:
```bash
#!/bin/bash
# Copy ${VERSION} stage index images to quay.io/openshift-pipeline
#
# Generated: ${REPORT_TIMESTAMP}
# Snapshots used:
${SNAPSHOT_COMMENTS}

set -euo pipefail

${SKOPEO_COMMANDS}

echo "Done — all index images copied."
```

Make the script executable:
```bash
chmod +x scripts/copy-index-images-${VERSION}-stage.sh
```

Show the user the script path and how to run it:
```
Copy script written to: scripts/copy-index-images-${VERSION}-stage.sh

Run it:
  ./scripts/copy-index-images-${VERSION}-stage.sh

Or run individual commands above.
```

After copying, re-verify that images exist on quay.io.

---

## Report Output

After processing all steps, write the stage report to `${REPORT_BASE}/image-copy/report_${REPORT_TIMESTAMP}.md`.

**Report format:**

```markdown
# Image Copy Stage Report — ${VERSION}

**Generated:** ${REPORT_TIMESTAMP}
**Release:** ${VERSION} (${MAJOR_MINOR})
**Branch:** ${RELEASE_BRANCH}

## Summary

| Step | Title | Status | Details | Links |
|------|-------|--------|---------|-------|
| 3.1 | Extract index image digests | {status} | {details} | {links} |
| 3.2 | Copy index images to quay.io | {status} | {details} | {links} |

## Step Details

### Step 3.1: Extract index image digests
- **Status:** {DONE | ACTION NEEDED | SKIPPED}
- **Index images:**

| Index App | Snapshot | Container Image | OCP Version |
|-----------|----------|-----------------|-------------|
| {app} | {snapshot} | {containerImage} | {ocp_version} |

### Step 3.2: Copy index images to quay.io
- **Status:** {DONE | ACTION NEEDED | SKIPPED}
- **Images copied:**

| OCP Version | Source | Quay Target | Status |
|-------------|--------|-------------|--------|
| {ocp_version} | {containerImage} | quay.io/openshift-pipeline/pipelines-index-{ocp_version}:v{VERSION}-stage | {DONE/PENDING} |

- **Copy script:** {path if generated}

## Blocking Step

{If stopped early, show which step blocked and why. Omit if all steps DONE.}
```

Write the report file and print the path to the user:
```
Report written to: reports/${MAJOR_MINOR}/${VERSION}/image-copy/report_${REPORT_TIMESTAMP}.md
```
