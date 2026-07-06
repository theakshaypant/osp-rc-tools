# Check Component PRs, Upstream Sync, and Nudges

Verify component PRs are merged on release branches, upstream sources are synced, and nudge PRs have landed.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `COMPONENTS` (list from release config), `TZ_FMT`

**Report ownership:** Steps 6, 6b, 7, 8. Detail section: `## Component Details`.

**Formatting:** All PR numbers must be rendered as markdown links `[#NUM](https://github.com/openshift-pipelines/REPO/pull/NUM)` using the correct repo. All commit SHAs as `[SHORT](https://github.com/OWNER/REPO/commit/FULL)`. All timestamps as absolute local time.

**Early stop:** Check steps 6, 6b, 7, 8 in order. If any step requires action, return immediately with the blocking step's status and action details. Do not check subsequent steps.

## Step 6: Check component PRs on release branches

The generate-konflux workflow creates PRs in each downstream component repo. These PRs:
- Branch: `hack/openshift-pipelines-core/${RELEASE_BRANCH}`
- Labels: `hack`, `automated`
- Title: `[bot:hack/openshift-pipelines-core/${RELEASE_BRANCH}] update konflux configuration`

They are auto-merged by the `merge-hack-pull-requests` workflow (runs hourly, merges PRs with `hack` label).

For each component in the release config, resolve its downstream repo name:
```bash
gh api repos/openshift-pipelines/hack/contents/config/downstream/repos/${COMPONENT}.yaml \
  --jq '.content' | base64 -d | grep '^repo:' | awk '{print $2}'
```

Then check for the component PR:
```bash
gh pr list --repo openshift-pipelines/${REPO} \
  --head "hack/openshift-pipelines-core/${RELEASE_BRANCH}" \
  --state all --limit 3 \
  --json number,title,state,mergedAt
```

Track: total components, merged PRs, open PRs, missing PRs. Show all timestamps as absolute local time.

**Important — before the operator PR merges:** Verify the bundle version in `project.yaml` is updated:
```bash
gh api repos/openshift-pipelines/operator/contents/project.yaml?ref=${RELEASE_BRANCH} \
  --jq '.content' | base64 -d | head -5
```

Report: `N/M components with merged PRs`. List any open or missing.

If all merged → **DONE**. If some open, note they should auto-merge within the hour (merge-hack-pull-requests runs hourly). If missing, the generate-konflux workflow may need re-triggering.

## Step 6b: Check olm/config.yaml bundle version

The `olm/config.yaml` on the release branch lists the bundle versions that `render-olm-catalog` processes. The current release version **must** be present before the upstream source sync (step 7) and before code freeze (step 12). The `update-sources` workflow is disabled during code freeze (`if: false`), so if the bundle version is missing at that point, fixing it requires manually re-enabling the workflow ([#25134](https://github.com/openshift-pipelines/operator/pull/25134)).

```bash
gh api repos/openshift-pipelines/operator/contents/olm/config.yaml?ref=${RELEASE_BRANCH} \
  --jq '.content' | base64 -d
```

Parse the `bundles:` list and extract all `version:` values. The last bundle entry should be `VERSION`. For patch releases within the same minor (e.g. `1.21.2` → `1.21.3`), the previous patch version is **replaced** — not kept alongside the new one.

- If `VERSION` is the last bundle entry → **DONE**
- If `VERSION` is **not present** → **ACTION NEEDED** — this is a **hard blocker**. Without it, `render-olm-catalog` will not render a catalog for the current release. The `update-olm-bundle.sh` script only processes bundle versions listed in this config. If the current version is missing, the render will either use a stale bundle image (which may not exist in the target registry) or skip the version entirely, producing **empty catalog.json files** that break all index builds.

If the config needs updating, create a PR on `${RELEASE_BRANCH}` that:
1. Replaces the previous patch version with `VERSION` in the `version:` field
2. Keeps the existing `image:` line unchanged (the `operator-update-images` workflow overwrites it later)
3. Removes the previous patch version entry entirely — only `VERSION` should remain (alongside any older cross-minor entries like `1.18.0`)

The `render-catalog.sh` CI check will automatically run on the PR and push additional commits that update all `catalog-template.json` files across every OCP version (v4.14+). It may also re-add the previous patch version entry — if so, remove it again. CI will show a `render-catalog.sh` failure for the initial commit (stale image SHA) — this is expected.

```
NEXT ACTION: Update olm/config.yaml bundle version to ${VERSION}.

The olm/config.yaml on ${RELEASE_BRANCH} has version ${PREVIOUS_VERSION}
as the last bundle entry. It must be ${VERSION} for render-olm-catalog
to produce a valid catalog.

Create a PR on ${RELEASE_BRANCH} to:
1. Replace ${PREVIOUS_VERSION} with ${VERSION}, keeping the existing
   image reference (a workflow updates it later)
2. Remove the ${PREVIOUS_VERSION} entry entirely

Result should look like:
  bundles:
    - version: 1.18.0
      tag: 1.18.0
    - version: ${VERSION}
      image: <kept from ${PREVIOUS_VERSION} entry>

The render-catalog.sh CI will auto-update catalog-template.json files.
If it re-adds ${PREVIOUS_VERSION}, remove it again.

IMPORTANT: This must be done BEFORE code freeze (step 12). The
update-sources workflow is disabled during freeze, so if this is
missed, the sync cannot propagate the change without manual
intervention (see #25134).

Reference: https://github.com/openshift-pipelines/operator/commit/3bc737e023a9
```

## Step 7: Check upstream source sync

The `update-sources` workflow is a dispatcher that triggers `update-sources.yaml` on each release branch. It runs daily at midnight UTC or can be triggered manually.

It creates PRs in the operator repo with:
- Branch: `actions/update/sources-${RELEASE_BRANCH}`
- Labels: `automated`, `upstream`
- Title: `[bot] Update ${RELEASE_BRANCH} from tektoncd/operator:...`

Check for update-sources PRs:
```bash
gh pr list --repo openshift-pipelines/operator \
  --head "actions/update/sources-${RELEASE_BRANCH}" \
  --state all --limit 5 \
  --json number,title,state,mergedAt,createdAt
```

- If a recent merged PR exists → **DONE**. Show merge time as absolute local time.
- If a PR is open → it may be waiting for CI or review. Show creation time. Check its status:
  ```bash
  gh pr checks ${PR_NUMBER} --repo openshift-pipelines/operator
  ```
- If no PRs exist → trigger manually:
  ```
  NEXT ACTION: Trigger update-sources workflow.

  Go to: https://github.com/openshift-pipelines/operator/actions/workflows/update-sources.yaml
  Click "Run workflow" — it auto-detects all release branches.

  Or wait for the scheduled daily run (midnight UTC).
  ```

This is fully automated — the PR is auto-merged when all pipelines pass.

## Step 8: Check component builds and nudge PRs

Once components are built on Konflux, nudge PRs are created in the operator repo to update image SHAs in `project.yaml`. These PRs have a specific pattern:
- Branch: `konflux/component-updates/operator-${MM_DASHED}-bundle-component-update-{component-image}`
- Labels: `konflux-nudge`, `konflux`, `automated`, `lgtm`, `approved`
- Title: `chore(deps): update {component-image} to {sha}`

Check for nudge PRs using the `konflux-nudge` label:
```bash
gh pr list --repo openshift-pipelines/operator \
  --base ${RELEASE_BRANCH} \
  --label konflux-nudge \
  --state all --limit 100 \
  --json number,title,state,mergedAt,createdAt
```

Categorize:
- Merged nudge PRs (SHA updates landed) — show merge time as absolute local time
- Open nudge PRs (waiting for CI/merge) — show creation time as absolute local time

If open nudge PRs exist, check their CI status:
```bash
gh pr checks ${PR_NUMBER} --repo openshift-pipelines/operator
```

Report which component images have been nudged. If all done → **DONE**. Otherwise list open PRs with their CI status.

This is fully automated. If builds are stuck, you can:
1. Check the component monitor dashboard: https://openshift-pipelines.github.io/hack/
2. Trigger image rebuilds: https://github.com/openshift-pipelines/hack/actions/workflows/trigger-image-rebuilds.yaml
   (Parameters: version=${MAJOR_MINOR}, optionally a specific repo name)

**Return:** Status for steps 6, 6b, 7, 8 with details. Include counts (merged/open/missing PRs, merged/open nudges). After producing results, read and follow the report update instructions in `REPORT.md` to write/update `reports/checklist-${VERSION}.md`.
