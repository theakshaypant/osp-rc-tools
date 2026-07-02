# Check p12n-opc Sync

Verify that the p12n-opc repository is synced with the OPC repository on the release branch.

**Inputs:** `VERSION`, `MAJOR_MINOR`, `RELEASE_BRANCH`, `TZ_FMT`

**Formatting:** All PR numbers must be rendered as markdown links. All commit SHAs as links. All timestamps as absolute local time.

## Step 2: Compare p12n-opc with OPC

The `p12n-opc` repo mirrors the OPC repo under its `upstream/` directory. The `upstream/pkg/version.json` in p12n-opc should match `pkg/version.json` in OPC.

### Fetch OPC version.json

```bash
OPC_VERSION=$(gh api repos/openshift-pipelines/opc/contents/pkg/version.json?ref=${RELEASE_BRANCH} \
  --jq '.content' | base64 -d)
```

### Fetch p12n-opc version.json

```bash
P12N_VERSION=$(gh api repos/openshift-pipelines/p12n-opc/contents/upstream/pkg/version.json?ref=${RELEASE_BRANCH} \
  --jq '.content' | base64 -d)
```

### Compare

Parse both JSON objects and compare each field. Also compare the latest commit on the release branch for both repos:

```bash
OPC_HEAD=$(gh api repos/openshift-pipelines/opc/commits/${RELEASE_BRANCH} --jq '.sha')
P12N_OPC_HEAD=$(gh api repos/openshift-pipelines/p12n-opc/commits/${RELEASE_BRANCH} --jq '.sha')
```

Check for recent sync PRs on p12n-opc:
```bash
gh pr list --repo openshift-pipelines/p12n-opc \
  --base "${RELEASE_BRANCH}" \
  --state all --limit 5 \
  --json number,title,state,mergedAt
```

### Present results

If version.json files match → **DONE**. Show both HEADs and last sync time.

If they differ:
```
NEXT ACTION: Sync p12n-opc with OPC.

The p12n-opc upstream/pkg/version.json is out of sync with OPC pkg/version.json on ${RELEASE_BRANCH}.

Differences:
${DIFF_TABLE}

Repository: https://github.com/openshift-pipelines/p12n-opc/tree/${RELEASE_BRANCH}

Sync the upstream/ directory in p12n-opc to match the current state of OPC on ${RELEASE_BRANCH}.
Create a PR targeting ${RELEASE_BRANCH}.
```

**Return:** Status for step 2 with sync details.
