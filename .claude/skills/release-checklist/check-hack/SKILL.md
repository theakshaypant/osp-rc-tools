---
description: Verify hack repo release version and release-manager PR (steps 1-2)
---

# Release Checklist: Hack Configuration

<objective>
Run the hack repo verification subset of the release checklist (steps 1-2). Checks that the release version is configured correctly in the hack repo and that the release-manager PR is merged.

Invoked as: `/release-checklist:check-hack <version>`

Example: `/release-checklist:check-hack 1.22.4`

This subskill can run independently or as part of the full `/release-checklist` run. When run independently, it updates only its owned section of the checklist report.

**External systems** (credentials in `.env`):
- GitHub: `gh` CLI (read/write via existing auth)

**Formatting rules:**
- **PR links:** Always render PR numbers as markdown links: `[#NUMBER](https://github.com/OWNER/REPO/pull/NUMBER)`
- **SHA links:** Always render commit SHAs as markdown links: `[SHORT_SHA](https://github.com/OWNER/REPO/commit/FULL_SHA)` (12 chars)
- **Timestamps:** All timestamps must be absolute local time with timezone. Never use relative times.
</objective>

<process>
<step name="setup">
Read and follow the instructions in `../SETUP.md`.

The version argument is passed as the skill argument (e.g. `1.22.4`).
</step>

<step name="check">
Read and follow the instructions in `../CHECK_HACK.md`.

Pass `VERSION`, `MAJOR_MINOR`, `RELEASE_TAG`, `TZ_FMT` to the subskill context.
</step>

<step name="report">
Read and follow the instructions in `../REPORT.md` to create or update `reports/checklist-${VERSION}.md`.

Only update the rows and sections owned by CHECK_HACK (steps 1, 2; no detail section).

Use the **Generated** timestamp in absolute local time:
```bash
GENERATED=$(date +"${TZ_FMT}")
```

Print the results to the conversation.
</step>
</process>

<success_criteria>
- [ ] Setup completed: version parsed, credentials loaded, release config fetched
- [ ] Steps 1 and 2 verified
- [ ] All PR numbers rendered as markdown links
- [ ] All timestamps use absolute local time with timezone
- [ ] Report at `reports/checklist-${VERSION}.md` created or updated (only steps 1-2)
</success_criteria>
