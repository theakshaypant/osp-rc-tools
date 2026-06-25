---
description: Search for Jira tickets on unlinked commits and generate missing release notes
---

# Audit Release

<objective>
Given a release audit result JSON (produced by `go run cmd/audit/main.go`), do two things:

1. **Find Jira tickets** for commits that the Go tool could not link (empty `jira` field) using PR body extraction and keyword JQL search via curl.
2. **Generate release notes** for commits/tickets that have no release note text, following the skipjira pattern: extract from PR description first, then generate using AI from commit + PR context. Always suggest a release note type from the allowed enum.

Skip all `dependabot[bot]` commits for both Jira search and release note generation.
</objective>

<execution_context>
**What the Go tool already provides per commit:**
```json
{
  "sha": "abc123...",
  "message": "fix: something",
  "author": "Name",
  "date": "2026-06-21",
  "pr": "https://github.com/org/repo/pull/123",
  "original_pr": "https://github.com/org/repo/pull/100",
  "jira": "https://redhat.atlassian.net/browse/SRVKP-1234",
  "release_note": {
    "text": "...",
    "type": "Bug Fix",
    "status": "Approved"
  }
}
```

**Jira REST API (curl with basic auth):**
```bash
# Search via JQL
curl -s -u "$JIRA_EMAIL:$JIRA_TOKEN" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -X POST "$JIRA_URL/rest/api/3/search/jql" \
  -d '{"jql":"...","fields":["key","summary","status"],"maxResults":5}'

# Fetch single issue (field IDs come from env: JIRA_RN_TEXT_FIELD, JIRA_RN_TYPE_FIELD, JIRA_RN_STATUS_FIELD)
curl -s -u "$JIRA_EMAIL:$JIRA_TOKEN" \
  -H "Accept: application/json" \
  "$JIRA_URL/rest/api/3/issue/SRVKP-1234?fields=summary,status,$JIRA_RN_TEXT_FIELD,$JIRA_RN_TYPE_FIELD,$JIRA_RN_STATUS_FIELD"
```

**GitHub PR info:**
```bash
gh pr view {number} --repo {org/repo} --json title,body,labels
gh pr diff {number} --repo {org/repo}
```

**Jira key pattern:** `[A-Z]{2,}-[0-9]+`

**Release note custom field IDs** (read from env vars, defaults in `jira.go`):
- `JIRA_RN_TEXT_FIELD` (default: `customfield_10783`) — Release Note Text
- `JIRA_RN_TYPE_FIELD` (default: `customfield_10785`) — Release Note Type (select)
- `JIRA_RN_STATUS_FIELD` (default: `customfield_10807`) — Release Note Status (select)

**Release Note Type enum** (the only valid values for the Type column):
Bug Fix · CVE · Deprecated Functionality · Developer Preview · Enhancement · Feature · Known Issue · Removed Functionality · Tech Preview · Release Notes Not Required · Unspecified
</execution_context>

<process>
<step name="setup">
**Source environment and validate.**

```bash
source .env
```

Verify `JIRA_URL`, `JIRA_EMAIL`, `JIRA_TOKEN` are set.

Read the audit result JSON. The user provides the path as an argument, or default to the most recent `release_*.json` in the current directory.

Parse the JSON. Report the count of commits needing work (non-dependabot, missing jira or release_note).
</step>

<step name="tier2_pr_body_search">
**Tier 2 — Extract Jira keys from PR body/title (confidence: high)**

For each non-dependabot commit where `jira` is empty AND `pr` or `original_pr` is set:

1. Prefer `original_pr` when available (upstream PRs more likely to contain Jira keys). Only fall back to `pr` if `original_pr` yields no keys.
2. Fetch the PR body and title via `gh pr view` (see curl patterns above). **Cache all results** in a dictionary keyed by PR URL — reuse this cache in later steps. Never call `gh pr view` for the same PR twice.
3. Extract Jira keys: `echo "$pr_body $pr_title" | grep -oE '[A-Z]{2,}-[0-9]+' | sort -u`
4. Record matches with **confidence: high** regardless of project key — a key explicitly referenced in a PR body is a strong signal whether it is SRVKP, PLNSRVCE, or any other project.
</step>

<step name="tier3_keyword_search">
**Tier 3 — Keyword-based JQL search (confidence: medium/low)**

For commits still without a Jira link after Tier 2 (and not dependabot):

1. Extract search keywords from the commit message:
   - Strip conventional commit prefixes (`fix:`, `feat:`, `build(deps):`, etc.)
   - Strip `Merge pull request #NNN from ...` — use PR title instead if cached
   - For CVE/GHSA IDs, search by that ID directly
   - Take the first meaningful phrase (up to ~5 words)

2. Search Jira via JQL (see curl pattern above). Batch where possible using `OR` clauses to reduce API calls.

3. Score results:
   - **medium**: keyword match + ticket is open/in-progress + summary is clearly related
   - **low**: keyword match but weak signal (generic terms, old/closed ticket, different component)

4. For CVE commits with no keyword match, also try: `"project = SRVKP AND text ~ \"CVE-2026-XXXX\""`
</step>

<step name="match_unmatched_jiras">
**Match unmatched_jiras to commits.**

The audit result may contain `unmatched_jiras` — tickets in the fixVersion that were not linked to any commit by the Go tool.

For each unmatched Jira ticket, use the `summary` field already present in the JSON (do NOT fetch ticket descriptions from Jira). If the summary keywords match a commit message, record the association with **confidence: low**.
</step>

<step name="generate_release_notes">
**Generate release notes for commits missing them.**

Process commits that either:
- Have a `jira` field but `release_note` is null/missing OR `release_note.text` is empty
- Were newly linked to Jira in Tier 2/3 above

**Step A — Extract from PR description:**
Look for a "Release Notes" section in the PR body (normalize `\r\n` to `\n` first) — code blocks (` ```release-note ``` `), markdown headings (`## Release Notes`), or labeled text (`Release Notes:`). Extract content, strip fences. If the text is "NONE", treat as absent. Mark as `extracted`.

**Step B — Generate using AI:**
If no release notes found in the PR description, gather context:
- Commit message
- PR title and description (from cached `gh pr view` results)
- PR diff only if commit message + PR description are too terse to determine what changed: `gh pr diff {number} --repo {org/repo} | head -200`

Generate a 1-3 sentence release note: present tense, user-facing perspective, no PR numbers or internal references. Mark as `generated`.

**Step C — Classify the release note type:**
Skip commits where `release_note.type` is already set and non-empty. For the rest, classify using these signals:
- PR labels: `kind/bug` → Bug Fix, `kind/feature` → Feature, `kind/enhancement` → Enhancement
- CVE/GHSA IDs → CVE
- `release-note-none` label or "NONE" text → Release Notes Not Required
- Dependency bumps for CVE fixes → CVE
- Internal chores, CI, image sha updates → Release Notes Not Required
- Default when uncertain → Unspecified
</step>

<step name="generate_report">
**Generate the markdown report.**

Organized **per component**. Each component section contains all commits in a single table:

```markdown
# Release Audit: {version}

**Source:** `{json_file_path}`<br>
**Total commits:** N across M components<br>
**Already linked (by Go tool):** N · **Newly linked:** N · **Unlinked:** N · **Dependabot (skipped):** N<br>

---

## {component-name}

**Upstream:** {upstream} · **Commits:** N · **Unsynced:** N

| SHA | Message | Author | PR | Jira | Release Note |
|-----|---------|--------|----|------|--------------|
| abc123 | fix: something | Name | [#100](orig-url) → [#123](pr-url) | [SRVKP-1234](url) | **Type:** Bug Fix<br>**Text:** Fixes panic on nil context. *(Jira)* |
| def456 | feat: new thing | Name | [#123](pr-url) | [SRVKP-5678](url) *(PR body)* | **Type:** Feature<br>**Text:** Adds custom CA bundles. *(extracted)* |
| ghi789 | chore: update CI | Name | [#45](pr-url) | [SRVKP-9012](url) *(keyword)* | **Type:** Enhancement<br>**Text:** Updates CI pipeline. *(generated)* |
| jkl012 | fix: edge case | Name | [#42](pr-url) | — | — |
| mno345 | bump deps | dependabot[bot] | [#50](pr-url) | — | skipped |
```

- **PR**: `[#orig](url) → [#cherry](url)` when both exist and differ; just `[#N](url)` otherwise; `—` when missing.
- **Jira**: Append confidence for audit-linked commits: *(PR body)*, *(keyword)*. No annotation for go-tool-linked. `—` when unlinked.
- **Release Note**: `**Type:**` and `**Text:**` with source: *(Jira)*, *(extracted)*, *(generated)*. `skipped` for dependabot. `—` when absent.

Include unsynced commits in a subsection if present. End with a cross-component summary table:

| Component | Commits | Linked (tool) | Linked (audit) | Unlinked | Has RN | Generated RN | Dependabot |
|-----------|---------|---------------|----------------|----------|--------|--------------|------------|

Write the report to `reports/audit-release-{version}.md`.
</step>
</process>

<success_criteria>
- [ ] All Jira API calls use curl with basic auth (not jira CLI) and read field IDs from env vars
- [ ] No duplicate API calls for same PR or Jira ticket (cache `gh pr view` results)
- [ ] `original_pr` checked before `pr` for Jira key extraction
- [ ] Report written to `reports/audit-release-{version}.md`
</success_criteria>
