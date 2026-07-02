---
description: Verify OLM config, index images, and code freeze (steps 11-12)
---

# Release Checklist: OLM and Code Freeze

<objective>
Run the OLM verification subset of the release checklist (steps 11-12). Checks OLM catalog configuration, index image builds, and code freeze status.

Invoked as: `/release-checklist:check-olm <version>`

Example: `/release-checklist:check-olm 1.22.4`

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
Read and follow the instructions in `../CHECK_OLM.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `CODE_FREEZE`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `TZ_FMT` to the subskill context.
</step>

<step name="report">
Read and follow the instructions in `../REPORT.md` to create or update `reports/checklist-${VERSION}.md`.

Only update the rows and sections owned by CHECK_OLM (steps 11, 12 and `## OLM and Code Freeze Details`).

Use the **Generated** timestamp in absolute local time:
```bash
GENERATED=$(date +"${TZ_FMT}")
```

Print the results to the conversation.
</step>
</process>

<success_criteria>
- [ ] Setup completed: version parsed, credentials loaded, release config fetched
- [ ] Steps 11 and 12 verified
- [ ] All Konflux commands target `tekton-ecosystem-tenant` namespace
- [ ] Konflux cluster accessed with READ-ONLY operations only
- [ ] All PR numbers rendered as markdown links
- [ ] All timestamps use absolute local time with timezone
- [ ] Report at `reports/checklist-${VERSION}.md` created or updated (only steps 11-12 and OLM and Code Freeze Details)
</success_criteria>
