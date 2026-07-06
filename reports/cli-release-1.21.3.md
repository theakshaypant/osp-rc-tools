# CLI Release: 1.21.3

**Generated:** 2026-07-06 16:04 IST
**Release branch:** release-v1.21.x

| # | Step | Status | Details |
|---|------|--------|---------|
| 1 | OPC version.json | DONE | All upstream versions current within pinned series |
| 2 | p12n-opc sync | DONE | version.json matches OPC |
| 3 | serve-tkn-cli submodules | DONE | All submodules current (after [#317](https://github.com/openshift-pipelines/serve-tkn-cli/pull/317)) |
| 4 | Product version configuration | **ACTION NEEDED** | File not found in konflux-release-data |

## Step 1: OPC version.json

All components in `pkg/version.json` on `release-v1.21.x` are current within their pinned upstream series:

| Component | OPC version.json | Pinned series | Latest in series | Status |
|-----------|-----------------|---------------|------------------|--------|
| pac | 0.39.7 | release-v0.39.x | v0.39.7 | CURRENT |
| tkn | 0.43.2 | release-v0.43.2 | v0.43.2 | CURRENT |
| results | 0.17.2 | release-v0.17.x | v0.17.2 | CURRENT |
| manualapprovalgate | 0.7.0 | release-v0.7.0 | v0.7.0 | CURRENT |
| assist | 0.1.1 | main | v0.1.1 | CURRENT |
| opc | 1.21.3 | (target) | 1.21.3 | CURRENT |

## Step 2: p12n-opc sync

`upstream/pkg/version.json` in p12n-opc matches OPC `pkg/version.json` on `release-v1.21.x`.

- OPC HEAD: [b705f0fcdf32](https://github.com/openshift-pipelines/opc/commit/b705f0fcdf326ffbea6be481b9c84d656b0663d9)
- p12n-opc HEAD: [938f533cf53c](https://github.com/openshift-pipelines/p12n-opc/commit/938f533cf53c62b99879ea7579d2c0658077e96b)

## Step 3: serve-tkn-cli submodules

All submodules current after [#317](https://github.com/openshift-pipelines/serve-tkn-cli/pull/317) merged:

| Submodule | Repo | Branch | Submodule SHA | Branch HEAD | Status |
|-----------|------|--------|---------------|-------------|--------|
| sources/cli | tektoncd/cli | release-v0.43.2 | [5818ca48c61e](https://github.com/tektoncd/cli/commit/5818ca48c61ee4aec3d61f6fda2c270da1bc32d6) | [5818ca48c61e](https://github.com/tektoncd/cli/commit/5818ca48c61ee4aec3d61f6fda2c270da1bc32d6) | CURRENT |
| sources/opc | openshift-pipelines/opc | release-v1.21.x | [b705f0fcdf32](https://github.com/openshift-pipelines/opc/commit/b705f0fcdf326ffbea6be481b9c84d656b0663d9) | [b705f0fcdf32](https://github.com/openshift-pipelines/opc/commit/b705f0fcdf326ffbea6be481b9c84d656b0663d9) | CURRENT |
| sources/pac | openshift-pipelines/pipelines-as-code | release-v0.39.x | [6acde9dbaa6d](https://github.com/openshift-pipelines/pipelines-as-code/commit/6acde9dbaa6da67cfbd67b14b9e0c0a2043c32bb) | [6acde9dbaa6d](https://github.com/openshift-pipelines/pipelines-as-code/commit/6acde9dbaa6da67cfbd67b14b9e0c0a2043c32bb) | CURRENT |

## Next Action

**Add product version configuration to konflux-release-data.**

Create a new file at:
```
data/external/developer-portal/openshift-pipelines/1.21.3.yaml
```

With content:
```yaml
---
versionName: "1.21.3"
ga: false
termsAndConditions: "Anonymous Download"
hidden: false
invisible: true
releaseDate: "2026-07-06"
```

Notes:
- `ga`: Set to `false` for patch versions on older minor releases.
- `invisible`: Set to `true` initially. Change to `false` after QA sign-off.
- `releaseDate`: Set to the actual release date.

Reference: previous version [1.21.2.yaml](https://gitlab.cee.redhat.com/releng/konflux-release-data/-/blob/main/data/external/developer-portal/openshift-pipelines/1.21.2.yaml)

Reference MR: https://gitlab.cee.redhat.com/releng/konflux-release-data/-/merge_requests/14753

Repository: https://gitlab.cee.redhat.com/releng/konflux-release-data

**IMPORTANT:** The product version MR must be merged BEFORE creating the CDN release in step 6. ReleasePlanAdmission cannot reference a version that is not yet defined.

## Remaining Steps

| # | Step | Status |
|---|------|--------|
| 5 | ReleasePlan and ReleasePlanAdmission | Not checked (RP/RPA exist on cluster — likely DONE) |
| 6 | CDN production release | Not checked |
