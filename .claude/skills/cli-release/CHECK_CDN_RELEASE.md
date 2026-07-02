# Check CDN Production Release

Verify the CDN production release status on Konflux.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `KONFLUX_NS` (`tekton-ecosystem-tenant`), `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT`

**Formatting:** All timestamps as absolute local time with timezone.

**Constraints:**
- Konflux cluster: **READ-ONLY** (`oc get`/`kubectl get` only, never `apply`/`create`/`delete`)

## Step 6: Check CDN production release

Stage release of the binaries requires manual product version configuration in stage CDN, so go directly to production release while keeping the `invisible` flag set to `true` in the product version YAML.

### Check for existing CDN releases

```bash
oc get releases -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify 2>&1 | grep -E "${MM_DASHED}.*(cdn-prod)"
```

Report all release timestamps as absolute local time.

### Get the latest snapshot

To find the latest snapshot for the core application:

```bash
oc get snapshot -n tekton-ecosystem-tenant \
  --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
  --insecure-skip-tls-verify \
  -o jsonpath="{range .items[*]} {.metadata.creationTimestamp}{'\t'} {.metadata.labels.pac\.test\.appstudio\.openshift\.io\/event-type}{'\t'} {.metadata.name} {'\n'}{end}" \
  --sort-by=.metadata.creationTimestamp \
  -l "pac.test.appstudio.openshift.io/event-type=push,appstudio.openshift.io/application=openshift-pipelines-core-${MM_DASHED}" \
  2>/dev/null | tail -5
```

Report the latest snapshot name and creation timestamp (absolute local time).

### Check release status

If a CDN production release exists:
- Check status: look for `Succeeded` vs `Failed` in the release output
- Report the release name, status, and timestamp

If the release Succeeded → **DONE**.

If the release Failed:
```
NEXT ACTION: Retry CDN production release with a newer snapshot.

The previous CDN production release failed. Create a new Release with the latest snapshot.
```

### If no CDN production release found

```
NEXT ACTION: Create CDN production release.

1. Get the latest core snapshot (last entry from the sorted list above):
   SNAPSHOT: ${LATEST_SNAPSHOT_NAME}

2. Create a release YAML file (e.g. cdn-release.yaml):

apiVersion: appstudio.redhat.com/v1alpha1
kind: Release
metadata:
  labels:
    appstudio.openshift.io/application: openshift-pipelines-core-${MM_DASHED}
  generateName: openshift-pipelines-${MM_DASHED}-core-cdn-prod-release-
  namespace: tekton-ecosystem-tenant
spec:
  data:
  gracePeriodDays: 10
  releasePlan: openshift-pipelines-${MM_DASHED}-core-cdn-prod
  snapshot: ${LATEST_SNAPSHOT_NAME}

3. Apply on Konflux:
   oc create -f cdn-release.yaml \
     --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
     --insecure-skip-tls-verify

4. Monitor the release:
   oc get releases -n tekton-ecosystem-tenant \
     --server="$KONFLUX_SERVER" --token="$KONFLUX_TOKEN" \
     --insecure-skip-tls-verify 2>&1 | grep "${MM_DASHED}.*cdn-prod"

5. After the release succeeds, update the product version YAML to set
   invisible: false (if ready for public visibility).
```

If `KONFLUX_SERVER` or `KONFLUX_TOKEN` are not set, mark as **MANUAL** and provide the instructions above without the server/token flags.

**Return:** Status for step 6 with release details and snapshot information.
