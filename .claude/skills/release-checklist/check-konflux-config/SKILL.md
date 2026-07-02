---
description: Verify Konflux configuration, RPA, and Pyxis setup (steps 3-5)
---

# Release Checklist: Konflux Configuration

<objective>
Run the Konflux configuration verification subset of the release checklist (steps 3-5). Checks that the Konflux config PR is merged, cluster config matches the hack repo, ReleasePlanAdmission exists, and Pyxis config is in place.

Invoked as: `/release-checklist:check-konflux-config <version>`

Example: `/release-checklist:check-konflux-config 1.22.4`

This subskill can run independently or as part of the full `/release-checklist` run. When run independently, it updates only its owned section of the checklist report.

**External systems** (credentials in `.env`):
- GitHub: `gh` CLI (read/write via existing auth)
- GitLab (gitlab.cee.redhat.com): `GITLAB_URL` + `GITLAB_TOKEN` — **READ-ONLY** (`GET` requests only)
- Konflux cluster: `KONFLUX_SERVER` + `KONFLUX_TOKEN` — **READ-ONLY** (`oc get`/`kubectl get` only)

**Konflux namespace:** All Konflux operations target `tekton-ecosystem-tenant`.

**Formatting rules:**
- **PR links:** Always render PR numbers as markdown links: `[#NUMBER](https://github.com/OWNER/REPO/pull/NUMBER)`. For GitLab MRs: `[!NUMBER](GITLAB_URL/path/-/merge_requests/NUMBER)`.
- **SHA links:** Always render commit SHAs as markdown links (12 chars)
- **Timestamps:** All timestamps must be absolute local time with timezone. Never use relative times.
</objective>

<process>
<step name="setup">
Read and follow the instructions in `../SETUP.md`.

The version argument is passed as the skill argument (e.g. `1.22.4`).
</step>

<step name="check">
Read and follow the instructions in `../CHECK_KONFLUX_CONFIG.md`.

Pass `VERSION`, `MAJOR_MINOR`, `MM_DASHED`, `RELEASE_BRANCH`, `KONFLUX_NS`, `KONFLUX_SERVER`, `KONFLUX_TOKEN`, `GITLAB_URL`, `GITLAB_TOKEN`, `TZ_FMT` to the subskill context.
</step>

<step name="report">
Read and follow the instructions in `../REPORT.md` to create or update `reports/checklist-${VERSION}.md`.

Only update the rows and sections owned by CHECK_KONFLUX_CONFIG (steps 3, 3b, 4, 5 and `## Konflux Config Details`).

Use the **Generated** timestamp in absolute local time:
```bash
GENERATED=$(date +"${TZ_FMT}")
```

Print the results to the conversation.
</step>
</process>

<success_criteria>
- [ ] Setup completed: version parsed, credentials loaded, release config fetched
- [ ] Steps 3, 3b, 4, and 5 verified
- [ ] All Konflux commands target `tekton-ecosystem-tenant` namespace
- [ ] GitLab and Konflux cluster accessed with READ-ONLY operations only
- [ ] All PR/MR numbers rendered as markdown links
- [ ] All timestamps use absolute local time with timezone
- [ ] Report at `reports/checklist-${VERSION}.md` created or updated (only steps 3-5 and Konflux Config Details)
</success_criteria>
