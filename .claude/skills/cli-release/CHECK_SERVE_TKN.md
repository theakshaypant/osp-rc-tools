# Check serve-tkn-cli Submodules

Verify that the serve-tkn-cli repository has the correct submodule references on the release branch.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `RELEASE_BRANCH`, `TZ_FMT`

**Formatting:** All PR numbers must be rendered as markdown links. All commit SHAs as links (first 12 chars). All timestamps as absolute local time.

## Step 3: Validate submodule SHAs

The `serve-tkn-cli` repo uses git submodules under `sources/` to track upstream repos.

### Fetch .gitmodules

```bash
gh api repos/openshift-pipelines/serve-tkn-cli/contents/.gitmodules?ref=${RELEASE_BRANCH} \
  --jq '.content' | base64 -d
```

This shows the submodule definitions with their `branch` fields:
- `sources/cli` → `tektoncd/cli` (branch like `release-v0.43.2`)
- `sources/opc` → `openshift-pipelines/opc` (branch like `release-v1.21.x`)
- `sources/pac` → `openshift-pipelines/pipelines-as-code` (downstream; branch like `release-v0.39.x`)

**Important:** Use the repo URL from `.gitmodules` to determine which GitHub org/repo to query. The PAC submodule points to the **downstream** `openshift-pipelines/pipelines-as-code`, not the upstream `tektoncd/pipelines-as-code`.

### Fetch actual submodule SHAs

```bash
gh api repos/openshift-pipelines/serve-tkn-cli/contents/sources?ref=${RELEASE_BRANCH} \
  --jq '.[] | "\(.name) \(.sha)"'
```

This returns the commit SHA each submodule points to.

### Verify each submodule

For each submodule, check if the SHA it points to matches the HEAD of the expected branch:

**sources/cli** (tektoncd/cli):
```bash
CLI_HEAD=$(gh api repos/tektoncd/cli/commits/$(BRANCH_FROM_GITMODULES) --jq '.sha' 2>/dev/null)
```
Compare with the submodule SHA. If they match → CURRENT. If not → OUTDATED.

**sources/opc** (openshift-pipelines/opc):
```bash
OPC_HEAD=$(gh api repos/openshift-pipelines/opc/commits/${RELEASE_BRANCH} --jq '.sha' 2>/dev/null)
```
Compare with the submodule SHA.

**sources/pac** (openshift-pipelines/pipelines-as-code — downstream):
```bash
PAC_HEAD=$(gh api repos/openshift-pipelines/pipelines-as-code/commits/$(BRANCH_FROM_GITMODULES) --jq '.sha' 2>/dev/null)
```
Compare with the submodule SHA.

### Check for recent update PRs

```bash
gh pr list --repo openshift-pipelines/serve-tkn-cli \
  --base "${RELEASE_BRANCH}" \
  --state all --limit 5 \
  --json number,title,state,mergedAt
```

### Present results

Show a comparison table:

```markdown
| Submodule | Repo | Branch | Submodule SHA | Branch HEAD | Status |
|-----------|------|--------|---------------|-------------|--------|
| sources/cli | tektoncd/cli | release-v0.44.2 | [abc123...] | [abc123...] | CURRENT |
| sources/opc | openshift-pipelines/opc | release-v1.22.x | [def456...] | [ghi789...] | **OUTDATED** |
| sources/pac | tektoncd/pipelines-as-code | release-v0.42.x | [jkl012...] | [jkl012...] | CURRENT |
```

If ALL submodules are CURRENT → **DONE**.

If any submodule is OUTDATED:
```
NEXT ACTION: Update serve-tkn-cli submodules on ${RELEASE_BRANCH}.

The following submodules need updating:
${LIST_OF_OUTDATED_SUBMODULES}

Repository: https://github.com/openshift-pipelines/serve-tkn-cli/tree/${RELEASE_BRANCH}

To advance submodules to their branch HEADs:
  git checkout ${RELEASE_BRANCH}
  git submodule update --init --remote --force --checkout
  git add sources/
  git commit -m "Update submodules"

Note: plain `git submodule update` checks out the currently recorded SHA
(yields "already up to date"). The `--remote --force --checkout` flags
fetch from the tracking branch and force-checkout the new HEAD.

Create a PR targeting ${RELEASE_BRANCH}.
```

**Return:** Status for step 3 with the submodule comparison table and details.
