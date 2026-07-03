# CLI Release: 1.21.3

**Generated:** 2026-07-02 16:44 IST
**Release branch:** release-v1.21.x

| # | Step | Status | Details |
|---|------|--------|---------|
| 1 | OPC version.json | DONE | All upstream versions current for 1.21.x series |
| 2 | p12n-opc sync | DONE | version.json matches OPC; synced via [#247](https://github.com/openshift-pipelines/p12n-opc/pull/247) 2026-06-25 10:14 IST |
| 3 | serve-tkn-cli submodules | DONE | All 3 submodules at branch HEAD; bumped via [#308](https://github.com/openshift-pipelines/serve-tkn-cli/pull/308) 2026-07-02 13:44 IST |
| 4 | Product version config | **ACTION NEEDED** | `1.21.3.yaml` not found in konflux-release-data |

## OPC Version Check

| Component | OPC version.json | Latest upstream | Status |
|-----------|-----------------|-----------------|--------|
| pac | 0.39.7 | 0.48.0 | CURRENT (pinned to 0.39.x) |
| tkn | 0.43.2 | 0.45.0 | CURRENT (pinned to 0.43.x) |
| results | 0.17.2 | 0.19.0 | CURRENT (pinned to 0.17.x) |
| manualapprovalgate | 0.7.0 | 0.9.0 | CURRENT (pinned to 0.7.x) |
| assist | 0.1.1 | 0.1.1 | CURRENT |
| opc | 1.21.3 | 1.21.3 (target) | CURRENT |

## serve-tkn-cli Submodule Details

| Submodule | Repo | Branch | Submodule SHA | Branch HEAD | Status |
|-----------|------|--------|---------------|-------------|--------|
| sources/cli | tektoncd/cli | release-v0.43.2 | [`db2b9a10b000`](https://github.com/tektoncd/cli/commit/db2b9a10b000a53d7b9d0c69587fb1e8dfceb9a4) | [`db2b9a10b000`](https://github.com/tektoncd/cli/commit/db2b9a10b000a53d7b9d0c69587fb1e8dfceb9a4) | CURRENT |
| sources/opc | openshift-pipelines/opc | release-v1.21.x | [`b705f0fcdf32`](https://github.com/openshift-pipelines/opc/commit/b705f0fcdf326ffbea6be481b9c84d656b0663d9) | [`b705f0fcdf32`](https://github.com/openshift-pipelines/opc/commit/b705f0fcdf326ffbea6be481b9c84d656b0663d9) | CURRENT |
| sources/pac | openshift-pipelines/pipelines-as-code | release-v0.39.x | [`6acde9dbaa6d`](https://github.com/openshift-pipelines/pipelines-as-code/commit/6acde9dbaa6da67cfbd67b14b9e0c0a2043c32bb) | [`6acde9dbaa6d`](https://github.com/openshift-pipelines/pipelines-as-code/commit/6acde9dbaa6da67cfbd67b14b9e0c0a2043c32bb) | CURRENT |

## Next Action

**Step 4: Add product version configuration to konflux-release-data.**

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
releaseDate: "2026-07-02"
```

Notes:
- **ga:** Set to `false` for patch versions. Set to `true` for minor releases and patches on the latest minor release.
- **invisible:** Set to `true` initially. Change to `false` after QA sign-off.
- **releaseDate:** Set to the actual release date.

Reference MR: https://gitlab.cee.redhat.com/releng/konflux-release-data/-/merge_requests/14753

Repository: https://gitlab.cee.redhat.com/releng/konflux-release-data

**IMPORTANT:** The product version MR must be merged BEFORE creating the RP/RPA in step 5. ReleasePlanAdmission cannot reference a version that is not yet defined.

## Remaining Steps

| # | Step |
|---|------|
| 5 | ReleasePlan and ReleasePlanAdmission |
| 6 | CDN production release |
