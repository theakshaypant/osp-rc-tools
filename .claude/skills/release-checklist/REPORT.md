# Report Management

Instructions for creating and updating the checklist report at `reports/checklist-${VERSION}.md`.

## Section Ownership

Each CHECK file owns specific step numbers and detail sections. Only update rows and sections you own.

| CHECK File | Step Numbers | Detail Section |
|-----------|-------------|----------------|
| CHECK_HACK | 1, 2 | (none) |
| CHECK_KONFLUX_CONFIG | 3, 3b, 4, 5 | `## Konflux Config Details` |
| CHECK_COMPONENTS | 6, 6b, 7, 8 | `## Component Details` |
| CHECK_BUILDS | 9, 9b, 10 | `## Build Validation Details` |
| CHECK_OLM | 11, 12 | `## OLM and Code Freeze Details` |
| CHECK_RELEASES | 13, 14, 15, 16 | `## Release Status Details` |

## Step Number Ordering

When inserting rows, maintain this canonical order: 1, 2, 3, 3b, 4, 5, 6, 6b, 7, 8, 9, 9b, 10, 11, 12, 13, 14, 15, 16.

## Creating a New Report

If `reports/checklist-${VERSION}.md` does not exist, create it with this structure:

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
${ROWS_FOR_STEPS_CHECKED}

${DETAIL_SECTIONS_FOR_STEPS_CHECKED}

## Next Action

${NEXT_ACTION_IF_BLOCKER_FOUND_OR_ALL_DONE}

## Remaining Steps

${STEPS_NOT_YET_CHECKED}
```

Only include summary table rows and detail sections for the steps actually checked by this invocation. The "Remaining Steps" section lists all 16 steps minus those present in the summary table.

## Updating an Existing Report

When `reports/checklist-${VERSION}.md` already exists:

1. **Read the entire file.**

2. **Update the header:**
   - Replace the `**Generated:**` line with the current timestamp
   - Update `**Release tag:**` and `**Code freeze:**` from the current release config

3. **Update the summary table:**
   - For each step number owned by the current CHECK file, find the matching row by the step number in column 1 (e.g. `| 9 |` or `| 3b |`)
   - If the row exists, replace it with the new status
   - If the row does not exist, insert it in the correct position (see Step Number Ordering above)
   - **Never modify rows belonging to other CHECK files**

4. **Update detail sections:**
   - For each detail section owned by the current CHECK file, find the section by its `## ` heading
   - If the section exists, replace everything from the heading to the next `## ` heading (exclusive) with the new content
   - If the section does not exist, insert it before `## Next Action`
   - **Never modify detail sections belonging to other CHECK files**

5. **Update Next Action:**
   - Scan the summary table for the first non-DONE step (in canonical order). That step's action guidance becomes the Next Action content.
   - If all steps in the table are DONE, write: "All checked steps are complete. Run remaining steps or the full checklist to continue."

6. **Update Remaining Steps:**
   - List all 16 steps that do NOT have a row in the summary table.

7. **Write the updated file** back to `reports/checklist-${VERSION}.md`.

## Formatting Reminders

These apply to all report content:
- All PR numbers must be markdown links: `[#NUM](https://github.com/OWNER/REPO/pull/NUM)`
- All commit SHAs must be markdown links: `[SHORT](https://github.com/OWNER/REPO/commit/FULL)` (12 chars for SHORT)
- All timestamps must be absolute local time with timezone (e.g. `2026-07-02 10:45 IST`)
