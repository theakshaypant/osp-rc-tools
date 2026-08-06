# CLI Release: 1.23.1

**Generated:** 2026-07-28 17:44 IST
**Release branch:** release-v1.23.x
**Hack repo release-tag:** 1.23.1
**Code freeze:** false

| # | Step | Status | Details |
|---|------|--------|---------|
| 1 | OPC version.json | DONE | `opc: "1.23.1"` — all components current |
| 2 | p12n-opc sync | DONE | version.json matches OPC; [#305](https://github.com/openshift-pipelines/p12n-opc/pull/305) merged 2026-07-28 17:05 IST |
| 3 | serve-tkn-cli submodules | **ACTION NEEDED** | `sources/opc` outdated — points to pre-bump commit |

## Step 1 Detail — OPC version.json

| Component | OPC version.json | Latest upstream | Status |
|-----------|-----------------|-----------------|--------|
| pac | 0.48.1 | v0.48.1 (latest in 0.48.x series) | CURRENT |
| tkn | 0.45.1 | v0.45.1 | CURRENT |
| results | 0.19.0 | v0.19.0 | CURRENT |
| manualapprovalgate | 0.9.0 | v0.9.0 | CURRENT |
| assist | 0.1.1 | v0.1.1 | CURRENT |
| opc | 1.23.1 | 1.23.1 (target) | CURRENT |

**OPC HEAD:** [a4b68c825829](https://github.com/openshift-pipelines/opc/commit/a4b68c82582902605e18b5023ed97d38ae38b6dc)

## Step 2 Detail — p12n-opc Sync

| Field | Value |
|-------|-------|
| OPC HEAD (release-v1.23.x) | [a4b68c825829](https://github.com/openshift-pipelines/opc/commit/a4b68c82582902605e18b5023ed97d38ae38b6dc) |
| p12n-opc HEAD (release-v1.23.x) | [c54d484423c9](https://github.com/openshift-pipelines/p12n-opc/commit/c54d484423c9df2845a832642b213b5c86f4918d) |
| p12n-opc version.json `opc` | 1.23.1 ✓ |
| Sync PR | [#305](https://github.com/openshift-pipelines/p12n-opc/pull/305) merged 2026-07-28 17:05 IST |

## Step 3 Detail — serve-tkn-cli Submodules

| Submodule | Repo | Branch | Submodule SHA | Branch HEAD | Status |
|-----------|------|--------|---------------|-------------|--------|
| sources/cli | tektoncd/cli | release-v0.45.x | [492f193a1361](https://github.com/tektoncd/cli/commit/492f193a136170a29c6d2a78dbb7466102bdc035) | [492f193a1361](https://github.com/tektoncd/cli/commit/492f193a136170a29c6d2a78dbb7466102bdc035) | CURRENT |
| sources/opc | openshift-pipelines/opc | release-v1.23.x | [2524083719f8](https://github.com/openshift-pipelines/opc/commit/2524083719f88a4efe20c578591cb510c53f647c) | [a4b68c825829](https://github.com/openshift-pipelines/opc/commit/a4b68c82582902605e18b5023ed97d38ae38b6dc) | **OUTDATED** |
| sources/pac | openshift-pipelines/pipelines-as-code | release-v0.48.x | [87637d31d933](https://github.com/openshift-pipelines/pipelines-as-code/commit/87637d31d933f65b7677b228360eb80e032455f6) | [87637d31d933](https://github.com/openshift-pipelines/pipelines-as-code/commit/87637d31d933f65b7677b228360eb80e032455f6) | CURRENT |

Last submodule PR [#342](https://github.com/openshift-pipelines/serve-tkn-cli/pull/342) merged 2026-07-27 08:36 IST — predates the OPC version bump (commit `a4b68c825829`). No new PR raised yet.

## Next Action

**Update `sources/opc` submodule in serve-tkn-cli to the OPC version bump commit.**

The OPC version bump (`a4b68c825829`) is not yet reflected in the serve-tkn-cli submodule. A bot PR should be raised automatically — check for open PRs:

```
https://github.com/openshift-pipelines/serve-tkn-cli/pulls?q=is:pr+base:release-v1.23.x+is:open
```

If the bot hasn't raised one, create it manually:

```bash
git clone -b release-v1.23.x https://github.com/openshift-pipelines/serve-tkn-cli.git /tmp/serve-tkn-cli-update
cd /tmp/serve-tkn-cli-update
git submodule update --init --remote --force --checkout sources/opc
git checkout -b "release/1.23.1/update-submodules"
git add sources/opc
git commit -m "[bot:1.23] Update submodules to latest upstream"
git push origin "release/1.23.1/update-submodules"
gh pr create --repo openshift-pipelines/serve-tkn-cli \
  --base release-v1.23.x \
  --head "release/1.23.1/update-submodules" \
  --title "[bot:1.23] Update submodules to latest upstream" \
  --label automated
rm -rf /tmp/serve-tkn-cli-update
```

## Remaining Steps (not checked)

| # | Step |
|---|------|
| 4 | Product version configuration (konflux-release-data) |
| 5 | ReleasePlan and ReleasePlanAdmission |
| 6 | CDN production release |
