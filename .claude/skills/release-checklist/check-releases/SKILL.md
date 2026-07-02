---
description: Verify stage/prod releases, QA handover, and advisory (steps 13-16)
---

# Release Checklist: Releases

<objective>
Run the release verification subset of the release checklist (steps 13-16). Checks stage release status, QA handover readiness, production release, and advisory completion.

Invoked as: `/release-checklist:check-releases <version>`

Example: `/release-checklist:check-releases 1.22.4`

This subskill can run independently or as part of the full `/release-checklist` run. When run independently, it updates only its owned section of the checklist report.

**External systems** (credentials in `.env`):
- GitHub: `gh` CLI (read/write via existing auth)
- Konflux cluster: `KONFLUX_SERVER` + `KONFLUX_TOKEN` — **READ-ONLY** (`oc get`/`kubectl get` only)

**Konflux namespace:** All Konflux operations target `tekton-ecosystem-tenant`.

**Formatting rules:**
- **PR links:** Always render PR numbers as markdown links: `[#NUMBER](https://github.com/OWNER/REPO/pull/NUMBER)`
- **SHA links:** Always render commit SHAs as markdown links (12 chars)
- **Timestamps:** All timestamps must be absolute local time with timezone. Never use relative times.
</objective>

<process>
<step name="setup">
Read and follow the instructions in `../SETUP.md`.

The version argument is passed as the skill argument (e.g. `1.22.4`).
</step>

<step name="check">
Read and follow the instructions in `../CHECK_RELEASES.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT` to the subskill context.
</step>

<step name="report">
Read and follow the instructions in `../REPORT.md` to create or update `reports/checklist-${VERSION}.md`.

Only update the rows and sections owned by CHECK_RELEASES (steps 13, 14, 15, 16 and `## Release Status Details`).

Use the **Generated** timestamp in absolute local time:
```bash
GENERATED=$(date +"${TZ_FMT}")
```

Print the results to the conversation.
</step>
</process>

<success_criteria>
- [ ] Setup completed: version parsed, credentials loaded, release config fetched
- [ ] Steps 13-16 verified (with early stop at first blocker within this group)
- [ ] All Konflux commands target `tekton-ecosystem-tenant` namespace
- [ ] Konflux cluster accessed with READ-ONLY operations only
- [ ] All PR numbers rendered as markdown links
- [ ] All timestamps use absolute local time with timezone
- [ ] Report at `reports/checklist-${VERSION}.md` created or updated (only steps 13-16 and Release Status Details)
</success_criteria>
