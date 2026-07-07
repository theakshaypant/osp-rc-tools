#!/bin/bash
# Copy 1.21.3 stage index images to quay.io/openshift-pipeline
#
# Snapshots used:
#   4.14: openshift-pipelines-index-4-14-1-21-20260707-064556-000
#   4.16: openshift-pipelines-index-4-16-1-21-20260707-064555-000
#   4.18: openshift-pipelines-index-4-18-1-21-20260707-064554-000
#   4.20: openshift-pipelines-index-4-20-1-21-20260707-064555-000
#   4.21: openshift-pipelines-index-4-21-1-21-20260707-072615-000
#   4.22: openshift-pipelines-index-4-22-1-21-20260707-064555-000

set -euo pipefail

# 4.14 (IIB 1174739)
skopeo copy --all docker://registry-proxy.engineering.redhat.com/rh-osbs/iib@sha256:b6f8f04cd4e754fe6563931a8790ab5adc740a9193341569825469ebac10ab38 docker://quay.io/openshift-pipeline/pipelines-index-4.14:v1.21.3-stage --preserve-digests

# 4.16 (IIB 1174743)
skopeo copy --all docker://registry-proxy.engineering.redhat.com/rh-osbs/iib@sha256:9dd569be030dd22c9ffae7b27e2baad581598052b51adee4a59f37770aa4c63d docker://quay.io/openshift-pipeline/pipelines-index-4.16:v1.21.3-stage --preserve-digests

# 4.18 (IIB 1174742)
skopeo copy --all docker://registry-proxy.engineering.redhat.com/rh-osbs/iib@sha256:3c3b059f6b4ab0af74ec28d825240b87171d619763dd4e45933372370f82e1ae docker://quay.io/openshift-pipeline/pipelines-index-4.18:v1.21.3-stage --preserve-digests

# 4.20 (IIB 1174740)
skopeo copy --all docker://registry-proxy.engineering.redhat.com/rh-osbs/iib@sha256:55cf1e3302e7ac7ca1eeaca1eb898e48da9f4217d4c2db487a3d72d9edd6c9a5 docker://quay.io/openshift-pipeline/pipelines-index-4.20:v1.21.3-stage --preserve-digests

# 4.21 (IIB 1174758)
skopeo copy --all docker://registry-proxy.engineering.redhat.com/rh-osbs/iib@sha256:366b6b1fb3f13ff81bf9d00a798c67fba399266cb4ed54dfef67eef0efdadf91 docker://quay.io/openshift-pipeline/pipelines-index-4.21:v1.21.3-stage --preserve-digests

# 4.22 (IIB 1174741)
skopeo copy --all docker://registry-proxy.engineering.redhat.com/rh-osbs/iib@sha256:5ffb6149e20004bf653c232b43183e10284d038e87dcaa7b1895b747baa9842b docker://quay.io/openshift-pipeline/pipelines-index-4.22:v1.21.3-stage --preserve-digests

echo "Done — all index images copied."
