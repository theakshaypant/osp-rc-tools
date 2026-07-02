---
description: Track CLI release progress for serve-tkn-cli binaries to CDN
---

# CLI Release

<objective>
Given a release version (e.g. `1.22.4`), walk through the CLI release workflow steps in order, verify which are complete by querying GitHub, GitLab, and Konflux cluster, and **stop at the first step that requires action**.

Steps are **sequential** — each step depends on the previous one being complete. The skill validates OPC upstream versions, p12n-opc sync, serve-tkn-cli submodules, product version configuration, RP/RPA, and CDN production release.

**Early stop rule:** Check steps in order. As soon as a step's status is anything other than DONE, **stop checking further steps**. Report all completed steps plus the first blocking step with its actionable guidance. Do not query systems for steps beyond the blocker.

The version is passed as the skill argument. Example: `/cli-release 1.22.4`

**External systems** (credentials in `.env`):
- GitHub: `gh` CLI (read via existing auth)
- GitLab (gitlab.cee.redhat.com): `GITLAB_URL` + `GITLAB_TOKEN` — **READ-ONLY** (`GET` requests only)
- Konflux cluster: `KONFLUX_SERVER` + `KONFLUX_TOKEN` — **READ-ONLY** (`oc get`/`kubectl get` only)

**Konflux namespace:** All Konflux operations target `tekton-ecosystem-tenant`.

**Formatting rules** (apply everywhere — report, conversation output, and all subskill output):
- **PR links:** Always render PR numbers as markdown links: `[#NUMBER](https://github.com/OWNER/REPO/pull/NUMBER)`. For GitLab MRs: `[!NUMBER](GITLAB_URL/path/-/merge_requests/NUMBER)`.
- **SHA links:** Always render commit SHAs as markdown links: `[SHORT_SHA](https://github.com/OWNER/REPO/commit/FULL_SHA)`. Use the first 12 characters for SHORT_SHA.
- **Timestamps:** All timestamps must be absolute local time with timezone (e.g. `2026-07-01 13:11 IST`). Never use relative times.
</objective>

<process>
<step name="setup">
**Parse arguments, initialize, and set up time formatting.**

Read and follow the instructions in the shared `../release-checklist/SETUP.md`.

The version argument must match `X.Y.Z`. This skill uses the same variables: `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `IS_PATCH`, `KONFLUX_NS`, `RELEASE_TAG`, `CODE_FREEZE`, `TZ_FMT`.
</step>

<step name="check_opc">
**Step 1: OPC version.json — upstream version check.**

Read and follow the instructions in `CHECK_OPC.md`.

Pass `VERSION`, `MAJOR_MINOR`, `RELEASE_BRANCH`, `TZ_FMT` to the subskill context.

If any upstream version is outdated → **ACTION NEEDED**. Stop here.
</step>

<step name="check_p12n_opc">
**Step 2: p12n-opc sync.**

Read and follow the instructions in `CHECK_P12N_OPC.md`.

Pass `VERSION`, `MAJOR_MINOR`, `RELEASE_BRANCH`, `TZ_FMT` to the subskill context.

If p12n-opc is out of sync → **ACTION NEEDED**. Stop here.
</step>

<step name="check_serve_tkn">
**Step 3: serve-tkn-cli submodule validation.**

Read and follow the instructions in `CHECK_SERVE_TKN.md`.

Pass `VERSION`, `MAJOR_MINOR`, `RELEASE_BRANCH`, `TZ_FMT` to the subskill context.

If submodules are outdated → **ACTION NEEDED**. Stop here.
</step>

<step name="check_product_version">
**Step 4: Product version configuration.**

Read and follow the instructions in `CHECK_PRODUCT_VERSION.md`.

Pass `VERSION`, `MAJOR_MINOR`, `IS_PATCH`, `GITLAB_URL`, `GITLAB_TOKEN`, `TZ_FMT` to the subskill context.

If product version config is missing → **ACTION NEEDED**. Stop here.
</step>

<step name="check_rp_rpa">
**Step 5: ReleasePlan and ReleasePlanAdmission.**

Read and follow the instructions in `CHECK_RP_RPA.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `GITLAB_URL`, `GITLAB_TOKEN`, `TZ_FMT` to the subskill context.

If RP/RPA is missing → **ACTION NEEDED**. Stop here.
</step>

<step name="check_cdn_release">
**Step 6: CDN production release.**

Read and follow the instructions in `CHECK_CDN_RELEASE.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT` to the subskill context.

If release is not done → **ACTION NEEDED** with release YAML template.
</step>

<step name="report">
**Step 7: Generate summary report.**

Compile all step results into a summary table and write to `reports/cli-release-${VERSION}.md`.

Use the **Generated** timestamp in absolute local time:
```bash
GENERATED=$(date +"${TZ_FMT}")
```

Include:
- Summary table with step status (only steps that were checked)
- Detail sections for checked steps (version comparison tables, submodule status, Konflux releases — only if those steps were reached)
- The first blocking step with actionable next-action guidance
- List of remaining unchecked steps

**Formatting reminders:**
- All PR numbers must be markdown links: `[#NUM](https://github.com/OWNER/REPO/pull/NUM)`
- All commit SHAs must be markdown links: `[SHORT](https://github.com/OWNER/REPO/commit/FULL)`
- All timestamps must be absolute local time with timezone

```markdown
# CLI Release: ${VERSION}

**Generated:** ${GENERATED}
**Release branch:** ${RELEASE_BRANCH}

| # | Step | Status | Details |
|---|------|--------|---------|
| ... | (only steps checked so far) | ... | ... |

## Next Action

${NEXT_ACTION_DETAILS_FOR_FIRST_BLOCKING_STEP}

## Remaining Steps

${LIST_OF_UNCHECKED_STEPS}
```

Write the report to `reports/cli-release-${VERSION}.md`.

Print the summary to the conversation.
</step>
</process>

<success_criteria>
- [ ] Version parsed and release config fetched from hack repo
- [ ] Steps verified sequentially — stopped at first step requiring action
- [ ] All Konflux commands target `tekton-ecosystem-tenant` namespace
- [ ] GitLab and Konflux cluster accessed with READ-ONLY operations only
- [ ] First incomplete step identified with actionable guidance
- [ ] Manual steps include exact URLs, commands, and parameters
- [ ] All PR numbers rendered as markdown links: `[#NUM](URL)`
- [ ] All commit SHAs rendered as markdown links: `[SHORT_SHA](URL)`
- [ ] All timestamps use absolute local time with timezone — no relative times
- [ ] Summary table written to `reports/cli-release-${VERSION}.md`
</success_criteria>
