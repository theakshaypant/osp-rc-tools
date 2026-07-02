---
description: Verify component PRs, upstream sync, and nudge PRs (steps 6-8)
---

# Release Checklist: Components

<objective>
Run the component verification subset of the release checklist (steps 6-8). Checks that component PRs are merged on release branches, upstream sources are synced, and nudge PRs have landed.

Invoked as: `/release-checklist:check-components <version>`

Example: `/release-checklist:check-components 1.22.4`

This subskill can run independently or as part of the full `/release-checklist` run. When run independently, it updates only its owned section of the checklist report.

**External systems** (credentials in `.env`):
- GitHub: `gh` CLI (read/write via existing auth)

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
Read and follow the instructions in `../CHECK_COMPONENTS.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `COMPONENTS`, `TZ_FMT` to the subskill context.
</step>

<step name="report">
Read and follow the instructions in `../REPORT.md` to create or update `reports/checklist-${VERSION}.md`.

Only update the rows and sections owned by CHECK_COMPONENTS (steps 6, 7, 8 and `## Component Details`).

Use the **Generated** timestamp in absolute local time:
```bash
GENERATED=$(date +"${TZ_FMT}")
```

Print the results to the conversation.
</step>
</process>

<success_criteria>
- [ ] Setup completed: version parsed, credentials loaded, release config fetched
- [ ] Steps 6, 7, and 8 verified
- [ ] All PR numbers rendered as markdown links
- [ ] All timestamps use absolute local time with timezone
- [ ] Report at `reports/checklist-${VERSION}.md` created or updated (only steps 6-8 and Component Details)
</success_criteria>
