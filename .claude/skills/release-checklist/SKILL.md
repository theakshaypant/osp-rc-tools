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

Read and follow the instructions in `SETUP.md`.

The version is passed as the skill argument. Example: `/release-checklist 1.21.3`
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
**Steps 6–8: Component PRs, OLM bundle version, upstream sync, and nudges.**

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
**Steps 11–12: OLM render/index status and code freeze.**

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

Read and follow the instructions in `REPORT.md` to create or update `reports/checklist-${VERSION}.md`.

Since this is the full orchestrator run, update ALL step rows and ALL detail sections from the results collected above.

Use the **Generated** timestamp in absolute local time:
```bash
GENERATED=$(date +"${TZ_FMT}")
```

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
