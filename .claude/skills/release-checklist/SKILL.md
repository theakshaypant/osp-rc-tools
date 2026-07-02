---
description: Verify release checklist progress and identify next steps for a patch or minor release
---

# Release Checklist

<objective>
Given a release version (e.g. `1.21.3`), walk through the OpenShift Pipelines release workflow steps in order, verify which are complete by querying GitHub, GitLab, Konflux cluster, and local tools, and **stop at the first step that requires action**.

Steps are **sequential** — each step depends on the previous one being complete. The skill validates not just that PRs merged or builds exist, but that builds originate from the **correct commit SHAs** on the downstream repos.

**Early stop rule:** Check steps in order. As soon as a step's status is anything other than DONE (i.e. ACTION NEEDED, IN PROGRESS, STALE, FAILED, etc.), **stop checking further steps**. Report all completed steps plus the first blocking step with its actionable guidance. Do not query systems for steps beyond the blocker.

The version is passed as the skill argument. Example: `/release-checklist 1.21.3`

**External systems** (credentials in `.env`):
- GitHub: `gh` CLI (read/write via existing auth)
- GitLab (gitlab.cee.redhat.com): `GITLAB_URL` + `GITLAB_TOKEN` — **READ-ONLY** (`GET` requests only)
- Konflux cluster: `KONFLUX_SERVER` + `KONFLUX_TOKEN` — **READ-ONLY** (`oc get`/`kubectl get` only)
- Jira: `JIRA_URL` + `JIRA_EMAIL` + `JIRA_TOKEN`

**Konflux namespace:** All Konflux operations target `tekton-ecosystem-tenant`.

**Formatting rules** (apply everywhere — report, conversation output, and all subskill output):
- **PR links:** Always render PR numbers as markdown links: `[#NUMBER](https://github.com/OWNER/REPO/pull/NUMBER)`. For GitLab MRs: `[!NUMBER](GITLAB_URL/path/-/merge_requests/NUMBER)`.
- **SHA links:** Always render commit SHAs as markdown links: `[SHORT_SHA](https://github.com/OWNER/REPO/commit/FULL_SHA)`. Use the first 12 characters for SHORT_SHA. When comparing SHAs (e.g. in build validation tables), both the HEAD SHA and Snapshot SHA columns should be links.
- **Timestamps:** All timestamps must be absolute local time with timezone (e.g. `2026-07-01 13:11 IST`). Never use relative times.
</objective>

<process>
<step name="setup">
**Parse arguments, initialize, and set up time formatting.**

Parse the version argument. It must match `X.Y.Z` (patch) or `X.Y.0` (minor).

Derive:
- `MAJOR_MINOR` — e.g. `1.21`
- `MM_DASHED` — e.g. `1-21` (dots replaced with dashes)
- `VERSION` — e.g. `1.21.3`
- `RELEASE_BRANCH` — `release-v${MAJOR_MINOR}.x`
- `IS_PATCH` — true if Z > 0
- `KONFLUX_NS` — always `tekton-ecosystem-tenant`

Source environment for external credentials:
```bash
source .env
```

Verify `GITLAB_URL`, `GITLAB_TOKEN`, `KONFLUX_SERVER`, `KONFLUX_TOKEN` are set. If any are missing, note which external checks will be skipped.

**Time formatting:**

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

Fetch the release config from the hack repo:
```bash
RELEASE_CFG=$(gh api repos/openshift-pipelines/hack/contents/config/downstream/releases/${MAJOR_MINOR}.yaml --jq '.content' | base64 -d)
```

Parse key fields from the YAML:
- `RELEASE_TAG` — the `release-tag` field (e.g. `1.21.3`)
- `CODE_FREEZE` — the `code-freeze` field (true/false)
- `COMPONENTS` — list of component names from the `branches` map

If the release config doesn't exist, stop: "Release config for ${MAJOR_MINOR} not found in hack repo."
</step>

<step name="check_hack">
**Steps 1–2: Hack repo configuration.**

Read and follow the instructions in `CHECK_HACK.md`.

Pass `VERSION`, `MAJOR_MINOR`, `RELEASE_TAG`, `TZ_FMT` to the subskill context.

If step 1 is not done, stop — remaining steps depend on this.
</step>

<step name="check_konflux_config">
**Steps 3–5: Konflux configuration, RPA, and Pyxis.**

Read and follow the instructions in `CHECK_KONFLUX_CONFIG.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `GITLAB_URL`, `GITLAB_TOKEN`, `TZ_FMT` to the subskill context.

**If any step in this group requires action, stop and skip to the report step.**
</step>

<step name="check_components">
**Steps 6–8: Component PRs, upstream sync, and nudges.**

Read and follow the instructions in `CHECK_COMPONENTS.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `COMPONENTS`, `TZ_FMT` to the subskill context.

**If any step in this group requires action, stop and skip to the report step.**
</step>

<step name="check_builds">
**Steps 9–10: Operator project.yaml version and build validation (SHA comparison).**

Read and follow the instructions in `CHECK_BUILDS.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT` to the subskill context.

**If project.yaml version is wrong (step 9), stop — images must not be built with the wrong version.**
**If builds are stale or split (step 10), stop and skip to the report step.**
</step>

<step name="check_olm">
**Steps 11–12: OLM config and code freeze.**

Read and follow the instructions in `CHECK_OLM.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `CODE_FREEZE`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT` to the subskill context.

**If any step in this group requires action, stop and skip to the report step.**
</step>

<step name="check_releases">
**Steps 13–16: Stage release, QA handover, production release, and advisory.**

Read and follow the instructions in `CHECK_RELEASES.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT` to the subskill context.

**If any step in this group requires action, stop and skip to the report step.**
</step>

<step name="report">
**Step 17: Generate summary report.**

Compile all step results into a summary table and write to `reports/checklist-${VERSION}.md`.

Use the **Generated** timestamp in absolute local time:
```bash
GENERATED=$(date +"${TZ_FMT}")
```

Include:
- Summary table with step status (only steps that were checked)
- Detail sections for checked steps (build validation, open PRs, Konflux releases — only if those steps were reached)
- The first blocking step with actionable next-action guidance
- List of remaining unchecked steps

**Formatting reminders:**
- All PR numbers must be markdown links: `[#NUM](https://github.com/OWNER/REPO/pull/NUM)`
- All commit SHAs must be markdown links: `[SHORT](https://github.com/OWNER/REPO/commit/FULL)`
- All timestamps must be absolute local time with timezone

```markdown
# Release Checklist: ${VERSION}

**Generated:** ${GENERATED}
**Release branch:** ${RELEASE_BRANCH}
**Release tag:** ${RELEASE_TAG}
**Code freeze:** ${CODE_FREEZE}
**Component monitor:** https://openshift-pipelines.github.io/hack/
**Konflux UI:** https://konflux-ui.apps.kflux-prd-rh02.0fk9.p1.openshiftapps.com

| # | Step | Status | Details |
|---|------|--------|---------|
| ... | (only steps checked so far) | ... | ... |

## Next Action

${NEXT_ACTION_DETAILS_FOR_FIRST_BLOCKING_STEP}

## Remaining Steps

${LIST_OF_UNCHECKED_STEPS}
```

Write the report to `reports/checklist-${VERSION}.md`.

Print the summary to the conversation.
</step>
</process>

<success_criteria>
- [ ] Release config fetched and parsed correctly from hack repo
- [ ] Steps verified sequentially — stopped at first step requiring action
- [ ] All Konflux commands target `tekton-ecosystem-tenant` namespace
- [ ] GitLab and Konflux cluster accessed with READ-ONLY operations only
- [ ] First incomplete step identified with actionable guidance
- [ ] Manual steps include exact URLs, commands, and parameters
- [ ] All PR numbers rendered as markdown links: `[#NUM](URL)`
- [ ] All commit SHAs rendered as markdown links: `[SHORT_SHA](URL)`
- [ ] All timestamps use absolute local time with timezone (e.g. `2026-07-01 13:11 IST`) — no relative times
- [ ] Summary table written to `reports/checklist-${VERSION}.md`
</success_criteria>
