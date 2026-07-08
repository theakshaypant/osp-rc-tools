#!/bin/bash
# Copy 1.21.3 FBC index images from Konflux snapshots to quay.io/openshift-pipeline
#
# Source: Konflux FBC snapshots (2026-07-08)
# These are the Konflux-built FBC images, replacing the old IIB-based copies.
#
# Snapshots used:
#   4.14: openshift-pipelines-index-4-14-1-21-20260708-043315-000
#   4.16: openshift-pipelines-index-4-16-1-21-20260708-043315-000
#   4.18: openshift-pipelines-index-4-18-1-21-20260708-043315-000
#   4.19: openshift-pipelines-index-4-19-1-21-20260708-043314-000
#   4.20: openshift-pipelines-index-4-20-1-21-20260708-043315-000
#   4.21: openshift-pipelines-index-4-21-1-21-20260708-043315-000
#   4.22: openshift-pipelines-index-4-22-1-21-20260708-043315-000
#
# Skipped (no snapshot): 4.17 (no components), 4.23 (empty config), 5.0 (empty config)

set -euo pipefail

# 4.14
skopeo copy --all docker://quay.io/redhat-user-workloads/tekton-ecosystem-tenant/pipelines-index-4.14@sha256:b672c2c40054fa6eb0c11fbea9d7db64d81ad82236ad5f6ee2a1b559b7ddf09a docker://quay.io/openshift-pipeline/pipelines-index-4.14:v1.21.3-stage --preserve-digests

# 4.16
skopeo copy --all docker://quay.io/redhat-user-workloads/tekton-ecosystem-tenant/pipelines-index-4.16@sha256:ef4d71480d67c36ade70f43df3e9e7be1fbb4646848c91afdc75718b524734c1 docker://quay.io/openshift-pipeline/pipelines-index-4.16:v1.21.3-stage --preserve-digests

# 4.18
skopeo copy --all docker://quay.io/redhat-user-workloads/tekton-ecosystem-tenant/pipelines-index-4.18@sha256:c20de853f1c73804fba0e20461e076b45994c55a0448754ae28ba64105833c69 docker://quay.io/openshift-pipeline/pipelines-index-4.18:v1.21.3-stage --preserve-digests

# 4.19
skopeo copy --all docker://quay.io/redhat-user-workloads/tekton-ecosystem-tenant/pipelines-index-4.19@sha256:3779ef00b0ca541fac9c024295be8ff62da2243f8ff7debfdcedff05b069ec9b docker://quay.io/openshift-pipeline/pipelines-index-4.19:v1.21.3-stage --preserve-digests

# 4.20
skopeo copy --all docker://quay.io/redhat-user-workloads/tekton-ecosystem-tenant/pipelines-index-4.20@sha256:84cc28c5dae3417082d6db21be3bd4a76ee18a9e40474a0f9d02dfaa8785e0a0 docker://quay.io/openshift-pipeline/pipelines-index-4.20:v1.21.3-stage --preserve-digests

# 4.21
skopeo copy --all docker://quay.io/redhat-user-workloads/tekton-ecosystem-tenant/pipelines-index-4.21@sha256:95b0f109041a7d008ffef43ef56a542e1bc60639d9d4d9be09baf2a60b3ac35c docker://quay.io/openshift-pipeline/pipelines-index-4.21:v1.21.3-stage --preserve-digests

# 4.22
skopeo copy --all docker://quay.io/redhat-user-workloads/tekton-ecosystem-tenant/pipelines-index-4.22@sha256:abc6e8e6dd86c89d9b0ceb8437bc7c69db7ab9ae54a305183aacd6e7d4ec139f docker://quay.io/openshift-pipeline/pipelines-index-4.22:v1.21.3-stage --preserve-digests

echo "Done — all 7 FBC index images copied to quay.io."
