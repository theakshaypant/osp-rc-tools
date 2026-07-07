# Jiras for OSP 1.21.3

## tektoncd-pipeline

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 1 | KONFLUX-6727 | Cache fixes, test categorization, build directives | https://redhat.atlassian.net/browse/KONFLUX-6727 | no |
| 2 | SRVKP-6762 | Pod latency metric conversion from Gauge to Histogram | https://redhat.atlassian.net/browse/SRVKP-6762 | no |
| 3 | SRVKP-12033 | Bump google.golang.org/grpc to fix CVE-2026-33186 | https://redhat.atlassian.net/browse/SRVKP-12033 | no rn |

## console-plugin

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 4 | SRVKP-9448 | [UI] Current status overlaps in Approvals list view | https://redhat.atlassian.net/browse/SRVKP-9448 | |
| 5 | SRVKP-9713 | Adding consoledataview and lazyactionmenu for remaining resources | https://redhat.atlassian.net/browse/SRVKP-9713 | |
| 6 | SRVKP-11399 | Added consoledataview and modified styles | https://redhat.atlassian.net/browse/SRVKP-11399 | |
| 7 | SRVKP-11401 | Migrate ApprovalList to ConsoleDataView | https://redhat.atlassian.net/browse/SRVKP-11401 | |
| 8 | SRVKP-11565 | Fix for same PLR names so they rerender when new PLR is added | https://redhat.atlassian.net/browse/SRVKP-11565 | |
| 9 | SRVKP-11566 | Removed LoadingInline wrapper and updated LoadingBox component | https://redhat.atlassian.net/browse/SRVKP-11566 | |
| 10 | SRVKP-11708 | Renamed classnames to avoid classname collisions | https://redhat.atlassian.net/browse/SRVKP-11708 | |
| 11 | SRVKP-11711 | Upgraded console SDK to 4.22.0-prerelease.3 | https://redhat.atlassian.net/browse/SRVKP-11711 | |
| 12 | SRVKP-11715 | Translations for master branch | https://redhat.atlassian.net/browse/SRVKP-11715 | |
| 13 | SRVKP-11951 | Updating modules in master branch to address CVEs | https://redhat.atlassian.net/browse/SRVKP-11951 | |
| 14 | SRVKP-12265 | Add CVE fix skill with dependency analysis script | https://redhat.atlassian.net/browse/SRVKP-12265 | |

## operator

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 15 | SRVKP-11355 | Update Go to 1.25.10 to address stdlib security vulnerabilities | https://redhat.atlassian.net/browse/SRVKP-11355 | no rn |
| 16 | SRVKP-11927 | Console Plugin Image should be picked conditionally based on OCP Version | https://redhat.atlassian.net/browse/SRVKP-11927 | needed |

## manual-approval-gate

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 17 | SRVKP-11741 | CVE-2026-33186 - google.golang.org/grpc fix | https://redhat.atlassian.net/browse/SRVKP-11741 | no rn |
| 18 | SRVKP-12342 | CVE-2026-39821 + CVE-2026-33814 - Upgrade Go and x/net | https://redhat.atlassian.net/browse/SRVKP-12342 | no rn |

## pipelines-as-code

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 19 | SRVKP-11061 | Add TransformFuncs to reduce cache memory usage | https://redhat.atlassian.net/browse/SRVKP-11061 | yes |
| 20 | SRVKP-11295 | Prevent duplicate repository CR on trailing slash | https://redhat.atlassian.net/browse/SRVKP-11295 | yes |
| 21 | SRVKP-12207 | Security: backport app token safeguards | https://redhat.atlassian.net/browse/SRVKP-12207 | yes |
| 22 | SRVKP-12241 | Security: backport app token safeguards | https://redhat.atlassian.net/browse/SRVKP-12241 | yes |

## tekton-caches

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 23 | SRVKP-11823 | Add Agentic Contexts and CI Workflows | https://redhat.atlassian.net/browse/SRVKP-11823 | no rn |
| 24 | SRVKP-12040 | CVE-2026-33186 gRPC-Go authorization bypass fix | https://redhat.atlassian.net/browse/SRVKP-12040 | no rn |

## tektoncd-hub

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 25 | SRVKP-12037 | Bump google.golang.org/grpc to 1.79.3 | https://redhat.atlassian.net/browse/SRVKP-12037 | no rn |

## tektoncd-results

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 26 | SRVKP-11380 | Improved performance for describe and logs commands with optimized filtering | https://redhat.atlassian.net/browse/SRVKP-11380 | confirm with Divyanshu - no resp |

## tektoncd-chains

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 27 | SRVKP-12344 | Multiple CVE fixes (CVE-2026-34986, CVE-2026-33211, CVE-2026-33186) | https://redhat.atlassian.net/browse/SRVKP-12344 | no rn |

## Unmatched Jiras

| # | Jira ID | Title | URL | Notes |
|---|---------|-------|-----|-------|
| 28 | SRVKP-12552 | OSP 1.21.3 release configuration | https://redhat.atlassian.net/browse/SRVKP-12552 | |
| 29 | SRVKP-12574 | OSP 1.21.3 patch release | https://redhat.atlassian.net/browse/SRVKP-12574 | |
| 30 | SRVKP-12585 | Test Openshift Pipelines 1.21.3 | https://redhat.atlassian.net/browse/SRVKP-12585 | |
| 31 | SRVKP-12586 | Upgrade test | https://redhat.atlassian.net/browse/SRVKP-12586 | |
| 32 | SRVKP-12587 | Acceptance test | https://redhat.atlassian.net/browse/SRVKP-12587 | |
| 33 | SRVKP-12588 | UI Regression testing | https://redhat.atlassian.net/browse/SRVKP-12588 | |
| 34 | SRVKP-12589 | CLI testing | https://redhat.atlassian.net/browse/SRVKP-12589 | |
