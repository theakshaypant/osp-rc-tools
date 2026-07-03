#!/usr/bin/env bash
# Extract per-component (repo, revision) from a Konflux snapshot.
#
# Prerequisites:
#   oc login --token=<TOKEN> --server=https://api.kflux-prd-rh02.0fk9.p1.openshiftapps.com:6443
#
# Usage:
#   ./scripts/snapshot-components.sh <snapshot-name>
#
# Example:
#   ./scripts/snapshot-components.sh openshift-pipelines-core-1-22-20260702-034604-000

set -euo pipefail

SNAPSHOT="${1:?Usage: $0 <snapshot-name>}"
NS="tekton-ecosystem-tenant"

CREATED=$(oc get snapshot "$SNAPSHOT" -n "$NS" \
  --insecure-skip-tls-verify \
  -o jsonpath='{.metadata.creationTimestamp}' 2>/dev/null)

echo "Snapshot: $SNAPSHOT"
echo "Created:  $(date -d "$CREATED" '+%Y-%m-%d %H:%M %Z')"
echo

oc get snapshot "$SNAPSHOT" -n "$NS" \
  --insecure-skip-tls-verify \
  -o json 2>/dev/null | jq -r '
  [.spec.components[]
   | {
       repo: (.source.git.url // "" | gsub("\\.git$";"") | gsub("^https://github.com/";"")),
       rev: (.source.git.revision // "?")
     }]
  | group_by(.repo)
  | sort_by(.[0].repo)[]
  | group_by(.rev)[]
  | "\(.[0].repo)\t\(.[0].rev[:12])\t\(length)"
' | column -t -s $'\t' -N "Repo,Revision,Components"
