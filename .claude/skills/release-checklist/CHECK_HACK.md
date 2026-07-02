# Check Hack Repo Configuration

Verify the hack repo has the correct release version configured and the release-manager PR is merged.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `RELEASE_TAG` (from release config), `TZ_FMT` (date format string for local time)

**Report ownership:** Steps 1, 2. No detail section.

**Formatting:** All PR numbers must be rendered as markdown links `[#NUM](https://github.com/openshift-pipelines/hack/pull/NUM)`. All commit SHAs as `[SHORT](https://github.com/OWNER/REPO/commit/FULL)`. All timestamps as absolute local time.

## Step 1: Verify release version

Check if `RELEASE_TAG` matches `VERSION`:
- If `RELEASE_TAG == VERSION` → **DONE**
- If `RELEASE_TAG` is a different version (e.g. `1.21.2` when we want `1.21.3`) → **ACTION NEEDED**

The `create-new-patch` function in the hack repo increments the last number in `release-tag` (e.g. `1.21.2` → `1.21.3`) and also sets `code-freeze: false`.

If not done, print:
```
NEXT ACTION: Generate new patch version in hack repo.

Go to: https://github.com/openshift-pipelines/hack/actions/workflows/release-new-patch.yaml
Click "Run workflow" with:
  - Branch: main
  - Version: ${MAJOR_MINOR}

This creates a PR on branch "actions/main/new-patch-${MAJOR_MINOR}" with label "automated".
Review and merge the PR to bump release-tag to the next patch version.
```
Stop here — remaining steps depend on this.

## Step 2: Check hack release-manager PR is merged

The release-manager workflow creates a PR with this exact pattern:
- Branch: `actions/main/new-patch-${MAJOR_MINOR}`
- Label: `automated`
- Title contains: `[bot: ${MAJOR_MINOR}] Release Action: new-patch`

Check:
```bash
gh pr list --repo openshift-pipelines/hack \
  --head "actions/main/new-patch-${MAJOR_MINOR}" \
  --state all --limit 5 \
  --json number,title,state,mergedAt
```

- If the most recent PR is merged → **DONE**. Show merge time as absolute local time: `date -d "${mergedAt}" +"${TZ_FMT}"`
- If the most recent PR is open → **ACTION NEEDED**: review and merge it
- If the most recent PR is closed (not merged) → may need to re-trigger the workflow

**Return:** Status for steps 1 and 2 with details. After producing results, read and follow the report update instructions in `REPORT.md` to write/update `reports/checklist-${VERSION}.md`.
