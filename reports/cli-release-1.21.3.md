# CLI Release: 1.21.3

**Generated:** 2026-07-09 15:56 IST
**Release branch:** release-v1.21.x

| # | Step | Status | Details |
|---|------|--------|---------|
| 1 | OPC version.json | DONE | opc: 1.21.3, all upstream on expected branches |
| 2 | p12n-opc sync | DONE | version.json match |
| 3 | serve-tkn-cli submodules | DONE | All 3 submodules CURRENT |
| 4 | Product version config | DONE | versionName: 1.21.3, releaseDate: 2026-07-09 |
| 5 | CDN RP/RPA | DONE | RPAs + RPs exist (cluster names use `-rp` suffix) |
| 6 | CDN production release | DONE | Succeeded 2026-07-09 15:27 IST |

## Step Details

### Step 1: OPC version.json

| Component | version.json | Latest upstream | Status |
|-----------|-------------|-----------------|--------|
| pac | 0.39.7 | 0.49.0 | REVIEW (pinned to release-v0.39.x) |
| tkn | 0.43.2 | 0.45.0 | REVIEW (pinned to release-v0.43.2) |
| results | 0.17.2 | 0.19.0 | REVIEW (pinned to release-v0.17.x) |
| manualapprovalgate | 0.7.0 | 0.9.0 | REVIEW (pinned to release-v0.7.0) |
| assist | 0.1.1 | 0.1.1 | CURRENT |
| opc | 1.21.3 | 1.21.3 (target) | CURRENT |

### Step 2: p12n-opc sync

- **OPC HEAD:** [b705f0fcdf32](https://github.com/openshift-pipelines/opc/commit/b705f0fcdf326ffbea6be481b9c84d656b0663d9)
- **p12n-opc HEAD:** [d6f0567f7dce](https://github.com/openshift-pipelines/p12n-opc/commit/d6f0567f7dce37aaf650115efbfe27b732bfd9f0)
- **version.json:** Match
- **Latest sync PR:** [#272](https://github.com/openshift-pipelines/p12n-opc/pull/272) MERGED 2026-07-09 12:15 IST

### Step 3: serve-tkn-cli submodules

| Submodule | Repo | Branch | Submodule SHA | Branch HEAD | Status |
|-----------|------|--------|---------------|-------------|--------|
| sources/cli | tektoncd/cli | release-v0.43.2 | [5818ca48c61e](https://github.com/tektoncd/cli/commit/5818ca48c61ee4aec3d61f6fda2c270da1bc32d6) | [5818ca48c61e](https://github.com/tektoncd/cli/commit/5818ca48c61ee4aec3d61f6fda2c270da1bc32d6) | CURRENT |
| sources/opc | openshift-pipelines/opc | release-v1.21.x | [b705f0fcdf32](https://github.com/openshift-pipelines/opc/commit/b705f0fcdf326ffbea6be481b9c84d656b0663d9) | [b705f0fcdf32](https://github.com/openshift-pipelines/opc/commit/b705f0fcdf326ffbea6be481b9c84d656b0663d9) | CURRENT |
| sources/pac | openshift-pipelines/pipelines-as-code | release-v0.39.x | [6acde9dbaa6d](https://github.com/openshift-pipelines/pipelines-as-code/commit/6acde9dbaa6da67cfbd67b14b9e0c0a2043c32bb) | [6acde9dbaa6d](https://github.com/openshift-pipelines/pipelines-as-code/commit/6acde9dbaa6da67cfbd67b14b9e0c0a2043c32bb) | CURRENT |

### Step 4: Product version config

- **versionName:** 1.21.3
- **ga:** false
- **hidden:** false
- **invisible:** false
- **releaseDate:** 2026-07-09

### Step 5: CDN RP/RPA

| Resource | Type | Location | Status |
|----------|------|----------|--------|
| openshift-pipelines-1-21-core-cdn-prod.yaml | RPA | GitLab config/ | FOUND |
| openshift-pipelines-1-21-core-cdn-stage.yaml | RPA | GitLab config/ | FOUND |
| openshift-pipelines-1-21-core-cdn-prod-rp | ReleasePlan | Cluster | FOUND |
| openshift-pipelines-1-21-core-cdn-stage-rp | ReleasePlan | Cluster | FOUND |

### Step 6: CDN production release

- **Release:** `openshift-pipelines-1-21-core-cdn-prod-release-s4bm6`
- **Status:** Succeeded
- **Completed:** 2026-07-09 15:27 IST
- **Snapshot:** `openshift-pipelines-core-1-21-20260706-124700-000-at`
- **ReleasePlan:** `openshift-pipelines-1-21-core-cdn-prod-rp`
