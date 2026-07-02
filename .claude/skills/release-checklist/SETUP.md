# Release Checklist Setup

Shared setup for the release checklist skill and its subskills.

**Input:** A version string argument matching `X.Y.Z` (e.g. `1.21.3`).

## Parse Version

Parse the version argument. It must match `X.Y.Z` (patch) or `X.Y.0` (minor).

Derive:
- `MAJOR_MINOR` — e.g. `1.21`
- `MM_DASHED` — e.g. `1-21` (dots replaced with dashes)
- `VERSION` — e.g. `1.21.3`
- `RELEASE_BRANCH` — `release-v${MAJOR_MINOR}.x`
- `IS_PATCH` — true if Z > 0
- `KONFLUX_NS` — always `tekton-ecosystem-tenant`

## Source Credentials

```bash
source .env
```

Verify `GITLAB_URL`, `GITLAB_TOKEN`, `KONFLUX_SERVER`, `KONFLUX_TOKEN` are set. If any are missing, note which external checks will be skipped.

## Time Formatting

Capture the local timezone and define a format string for absolute timestamps:
```bash
LOCAL_TZ=$(date +%Z)
TZ_FMT='%Y-%m-%d %H:%M %Z'
```

**CRITICAL:** All timestamps in the report and conversation output MUST use absolute local time. Convert all RFC3339/ISO8601 timestamps from `gh`, `oc`, Jira, or any API to local time:
```bash
date -d "${TIMESTAMP}" +"${TZ_FMT}"
```

**Never** use relative time expressions like "3 min ago", "~30 min ago", "6d5h", "N days ago", or similar. Always show the actual date and time with timezone.

## Fetch Release Config

```bash
RELEASE_CFG=$(gh api repos/openshift-pipelines/hack/contents/config/downstream/releases/${MAJOR_MINOR}.yaml --jq '.content' | base64 -d)
```

Parse key fields from the YAML:
- `RELEASE_TAG` — the `release-tag` field (e.g. `1.21.3`)
- `CODE_FREEZE` — the `code-freeze` field (true/false)
- `COMPONENTS` — list of component names from the `branches` map

If the release config doesn't exist, stop: "Release config for ${MAJOR_MINOR} not found in hack repo."

## Return

All variables above (`VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `IS_PATCH`, `KONFLUX_NS`, `RELEASE_TAG`, `CODE_FREEZE`, `COMPONENTS`, `TZ_FMT`) are available for use by the calling skill.
