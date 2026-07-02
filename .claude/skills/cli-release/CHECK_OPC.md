# Check OPC version.json

Verify that the OPC `pkg/version.json` on the release branch has up-to-date upstream versions.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `RELEASE_BRANCH`, `TZ_FMT`

**Formatting:** All PR numbers must be rendered as markdown links. All commit SHAs as links. All timestamps as absolute local time.

## Step 1: Fetch and compare upstream versions

Fetch `pkg/version.json` from the OPC repo on the release branch:
```bash
gh api repos/openshift-pipelines/opc/contents/pkg/version.json?ref=${RELEASE_BRANCH} \
  --jq '.content' | base64 -d
```

This returns a JSON object like:
```json
{"pac": "0.42.2", "tkn": "0.44.2", "results": "0.18.0", "manualapprovalgate": "0.8.0", "assist": "0.1.1", "opc": "1.22.4"}
```

For each upstream component, fetch the latest release tag and compare:

### pac (Pipelines as Code)

```bash
gh api repos/openshift-pipelines/pipelines-as-code/releases/latest --jq '.tag_name' 2>/dev/null
```

The tag may have a `v` prefix (e.g. `v0.42.2`). Strip it for comparison.

### tkn (Tekton CLI)

```bash
gh api repos/tektoncd/cli/releases/latest --jq '.tag_name' 2>/dev/null
```

### results (Tekton Results)

```bash
gh api repos/tektoncd/results/releases/latest --jq '.tag_name' 2>/dev/null
```

### manualapprovalgate

```bash
gh api repos/openshift-pipelines/manual-approval-gate/releases/latest --jq '.tag_name' 2>/dev/null
```

### assist

```bash
gh api repos/openshift-pipelines/tekton-assist/releases/latest --jq '.tag_name' 2>/dev/null
```

### opc version

The `opc` field should match `VERSION`. If it doesn't, the version bump hasn't been done yet.

## Present results

Show a comparison table:

```markdown
| Component | OPC version.json | Latest upstream | Status |
|-----------|-----------------|-----------------|--------|
| pac       | 0.42.2          | 0.42.2          | CURRENT |
| tkn       | 0.44.2          | 0.45.0          | **OUTDATED** |
| results   | 0.18.0          | 0.18.0          | CURRENT |
| manualapprovalgate | 0.8.0  | 0.8.0           | CURRENT |
| assist    | 0.1.1           | 0.1.1           | CURRENT |
| opc       | 1.22.3          | 1.22.4 (target) | **OUTDATED** |
```

If ALL components are CURRENT → **DONE**.

If any component is OUTDATED:
```
NEXT ACTION: Update OPC version.json on the release branch.

The following upstream versions need updating in pkg/version.json on branch ${RELEASE_BRANCH}:
${LIST_OF_OUTDATED_COMPONENTS}

Repository: https://github.com/openshift-pipelines/opc/tree/${RELEASE_BRANCH}/pkg

Update pkg/version.json with the new versions and also update the corresponding
dependencies in go.mod. Then run `go mod tidy && go mod vendor` to sync.

Create a PR targeting ${RELEASE_BRANCH}.
```

**Note:** Not all upstream versions need to track the absolute latest release. Some components may intentionally pin to an older version for the release branch. If the upstream latest is a different major/minor series than what's in version.json (e.g. version.json has `0.42.x` but upstream latest is `0.43.x`), report it but note it may be intentional for this release branch. Flag it as **REVIEW** instead of **OUTDATED**.

**Return:** Status for step 1 with the comparison table and details.
