#!/usr/bin/env bash
# audit-snapshot.sh - Audit release commits using snapshot-based bounds
#
# Usage: ./audit-snapshot.sh <snapshot-name> <version>
#
# Example:
#   ./audit-snapshot.sh openshift-pipelines-core-1-24-20261019-170342-000 1.24.2
#
# Requirements:
#   - oc (authenticated to Konflux cluster)
#   - gh (authenticated to GitHub)
#   - jq
#   - KONFLUX_SERVER and KONFLUX_TOKEN env vars
#   - GITHUB_TOKEN env var or gh auth login

set -euo pipefail

# ============================================================================
# CONFIGURATION
# ============================================================================

readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly REPO_ROOT="$(dirname "$SCRIPT_DIR")"
readonly CONFIG_FILE="$REPO_ROOT/config/components.json"
readonly REPORTS_DIR="$REPO_ROOT/config"
readonly KONFLUX_NS="${KONFLUX_NS:-tekton-ecosystem-tenant}"

# Global variables
SNAPSHOT_NAME=""
VERSION=""
MAJOR_MINOR=""
PREVIOUS_VERSION=""
RELEASE_BRANCH=""
REPORT_FILE=""
COMPONENT_COUNT=0
PROCESSED_COUNT=0
ERROR_COUNT=0

# Associative arrays
declare -A COMPONENT_CONFIG
declare -A SNAPSHOT_COMPONENTS
declare -A PR_CACHE
declare -A HEAD_FILE_CACHE

# Jira configuration
JIRA_PROJECT="${JIRA_PROJECT:-SRVKP}"
JIRA_PR_FIELD="${JIRA_PR_FIELD:-customfield_10875}"
JIRA_RN_TEXT_FIELD="${JIRA_RN_TEXT_FIELD:-customfield_10783}"
JIRA_RN_TYPE_FIELD="${JIRA_RN_TYPE_FIELD:-customfield_10785}"
JIRA_RN_STATUS_FIELD="${JIRA_RN_STATUS_FIELD:-customfield_10807}"

# ============================================================================
# UTILITY FUNCTIONS
# ============================================================================

log() {
    local level="$1"
    shift
    case "$level" in
        debug)
            if [[ "${DEBUG:-}" == "true" ]]; then
                echo "DEBUG: $*" >&2
            fi
            ;;
        progress)
            echo "$*" >&2
            ;;
        *)
            echo "${level^^}: $*" >&2
            ;;
    esac
}

log_info() { log info "$@"; }
log_warn() { log warn "$@"; }
log_error() { log error "$@"; }
log_debug() { log debug "$@"; }
log_progress() { log progress "$@"; }

die() {
    log_error "$*"
    exit 1
}

# Parse version into major.minor and calculate previous version
# Args: version (e.g., "1.24.1")
# Sets: MAJOR_MINOR, PREVIOUS_VERSION
parse_version() {
    local version="$1"

    if [[ ! "$version" =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
        die "Invalid version format: $version (expected X.Y.Z)"
    fi

    local major="${BASH_REMATCH[1]}"
    local minor="${BASH_REMATCH[2]}"
    local patch="${BASH_REMATCH[3]}"

    MAJOR_MINOR="${major}.${minor}"

    if [[ $patch -gt 0 ]]; then
        # Patch release: 1.24.1 -> 1.24.0
        PREVIOUS_VERSION="${major}.${minor}.$((patch - 1))"
    else
        # Minor release: 1.24.0 -> 1.23.0
        if [[ $minor -lt 1 ]]; then
            die "Minor version must be >= 1, got $version"
        fi
        PREVIOUS_VERSION="${major}.$((minor - 1)).0"
    fi

    RELEASE_BRANCH="release-v${MAJOR_MINOR}.x"

    log_debug "Version: $version"
    log_debug "Major.Minor: $MAJOR_MINOR"
    log_debug "Previous: $PREVIOUS_VERSION"
    log_debug "Branch: $RELEASE_BRANCH"
}

# Validate required environment and tools
validate_env() {
    # Check tools
    local missing_tools=()
    for tool in oc gh jq curl; do
        if ! command -v "$tool" &>/dev/null; then
            missing_tools+=("$tool")
        fi
    done

    if [[ ${#missing_tools[@]} -gt 0 ]]; then
        die "Missing required tools: ${missing_tools[*]}"
    fi

    # Check Konflux credentials
    if [[ -z "${KONFLUX_SERVER:-}" ]]; then
        die "KONFLUX_SERVER environment variable not set"
    fi

    if [[ -z "${KONFLUX_TOKEN:-}" ]]; then
        die "KONFLUX_TOKEN environment variable not set"
    fi

    # Check GitHub auth
    if ! gh auth status &>/dev/null; then
        die "GitHub authentication required: run 'gh auth login'"
    fi

    # Check Jira config if enabled
    if [[ -n "${JIRA_EMAIL:-}" ]] || [[ -n "${JIRA_TOKEN:-}" ]] || [[ -n "${JIRA_URL:-}" ]]; then
        # If any Jira var is set, all must be set
        [[ -z "${JIRA_EMAIL:-}" ]] && die "JIRA_EMAIL required when Jira is enabled"
        [[ -z "${JIRA_TOKEN:-}" ]] && die "JIRA_TOKEN required when Jira is enabled"
        [[ -z "${JIRA_URL:-}" ]] && die "JIRA_URL required when Jira is enabled"
    fi

    log_debug "Environment validation passed"
}

# ============================================================================
# CONFIGURATION LOADING
# ============================================================================

# Load component configuration from config/components.json
# Populates COMPONENT_CONFIG associative array
load_config() {
    if [[ ! -f "$CONFIG_FILE" ]]; then
        die "Config file not found: $CONFIG_FILE"
    fi

    log_progress "Loading component config for $MAJOR_MINOR..."

    # Check if version exists in config
    if ! jq -e --arg v "$MAJOR_MINOR" '.[$v]' "$CONFIG_FILE" &>/dev/null; then
        die "Version $MAJOR_MINOR not found in $CONFIG_FILE"
    fi

    # Load components into associative array
    # Format: COMPONENT_CONFIG["component:field"] = "value"
    local components_data
    components_data=$(jq -r --arg v "$MAJOR_MINOR" \
        '.[$v] | to_entries[] | "\(.key)|\(.value.upstream)|\(.value.branch)"' \
        "$CONFIG_FILE")

    while IFS='|' read -r component upstream branch; do
        if [[ "$upstream" == "TODO" ]]; then
            die "Component $component has upstream='TODO' - configure it in $CONFIG_FILE before running audit"
        fi

        COMPONENT_CONFIG["${component}:upstream"]="$upstream"
        COMPONENT_CONFIG["${component}:branch"]="$branch"
        COMPONENT_COUNT=$((COMPONENT_COUNT + 1))
    done <<< "$components_data"

    log_progress "  Loaded $COMPONENT_COUNT components"

    if [[ $COMPONENT_COUNT -eq 0 ]]; then
        die "No valid components found in config for version $MAJOR_MINOR"
    fi
}

# ============================================================================
# SNAPSHOT OPERATIONS
# ============================================================================

# Fetch snapshot from Konflux cluster
# Args: snapshot_name
# Returns: JSON on stdout
fetch_snapshot() {
    local snapshot_name="$1"

    log_progress "Fetching snapshot: $snapshot_name"

    oc get snapshot "$snapshot_name" \
        -n "$KONFLUX_NS" \
        --server="$KONFLUX_SERVER" \
        --token="$KONFLUX_TOKEN" \
        --insecure-skip-tls-verify \
        -o json 2>/dev/null || die "Failed to fetch snapshot: $snapshot_name"
}

# Extract components from snapshot JSON
# Args: snapshot_json
# Populates: SNAPSHOT_COMPONENTS associative array
extract_snapshot_components() {
    local snapshot_json="$1"

    log_progress "Extracting snapshot components..."

    local count=0
    while IFS='|' read -r name repo revision; do
        # Extract repo name from full path (e.g., openshift-pipelines/pipeline -> pipeline)
        local repo_name="${repo##*/}"

        # Try to match to component config
        # Snapshot names have suffixes like "tektoncd-chains-1-24-controller"
        # Config names are simple like "tektoncd-chains"
        local component=""
        for comp in "${!COMPONENT_CONFIG[@]}"; do
            if [[ "$comp" =~ ^([^:]+):upstream$ ]]; then
                local comp_name="${BASH_REMATCH[1]}"
                # Match if: exact name, exact repo, or name starts with component name
                if [[ "$name" == "$comp_name" ]] || \
                   [[ "$repo_name" == "$comp_name" ]] || \
                   [[ "$name" == "$comp_name"-* ]]; then
                    component="$comp_name"
                    break
                fi
            fi
        done

        if [[ -n "$component" ]]; then
            SNAPSHOT_COMPONENTS["${component}:revision"]="$revision"
            SNAPSHOT_COMPONENTS["${component}:repo"]="$repo"
            count=$((count + 1))
        else
            log_debug "Snapshot component '$name' not found in config, skipping"
        fi
    done < <(echo "$snapshot_json" | jq -r '
        .spec.components[] |
        "\(.name)|\((.source.git.url // "") | gsub("\\.git$";"") | gsub("^https://github\\.com/";""))|\(.source.git.revision // "")"
    ')

    log_progress "  Found $count components in snapshot"
}

# ============================================================================
# TAG RESOLUTION
# ============================================================================

# Read upstream SHA from downstream repo's HEAD file at a given tag/ref
# Args: downstream_repo, ref
# Returns: SHA on stdout, or exits with error
read_head_file() {
    local repo="$1"
    local ref="$2"
    local cache_key="${repo}:${ref}"

    # Check cache first
    if [[ -n "${HEAD_FILE_CACHE[$cache_key]:-}" ]]; then
        echo "${HEAD_FILE_CACHE[$cache_key]}"
        return 0
    fi

    log_debug "  Reading HEAD file from $repo at $ref"

    local sha
    sha=$(gh api "repos/$repo/contents/head?ref=$ref" \
        --jq '.content' 2>/dev/null | base64 -d | tr -d '\n') || {
        log_debug "  HEAD file not found at $ref"
        return 1
    }

    # Cache the result
    HEAD_FILE_CACHE[$cache_key]="$sha"
    echo "$sha"
}

# Resolve previous tag to upstream SHA
# Args: component, downstream_repo
# Returns: SHA on stdout
resolve_previous_tag() {
    local component="$1"
    local downstream_repo="$2"
    local prev_tag="v${PREVIOUS_VERSION}"

    log_debug "  Resolving $prev_tag in $downstream_repo"

    # Try reading HEAD file at tag
    local sha
    if sha=$(read_head_file "$downstream_repo" "$prev_tag" 2>/dev/null); then
        log_debug "  Resolved via tag: ${sha:0:12}"
        echo "$sha"
        return 0
    fi

    log_warn "$component: tag $prev_tag not found, skipping"
    return 1
}

# ============================================================================
# JIRA OPERATIONS
# ============================================================================

# Check if Jira integration is enabled
jira_enabled() {
    [[ -n "${JIRA_EMAIL:-}" ]] && [[ -n "${JIRA_TOKEN:-}" ]] && [[ -n "${JIRA_URL:-}" ]]
}

# Extract Jira keys from text (commit message, PR title/body)
# Args: text
# Returns: unique Jira keys on stdout, one per line
extract_jira_keys() {
    local text="$1"

    # Match all PROJECT-NNNN patterns (e.g., SRVKP-1234)
    printf '%s\n' "$text" | grep -Eo "${JIRA_PROJECT}-[0-9]+" | sort -u || true
}

# Merge two newline-separated lists of Jira keys
# Args: first_keys, second_keys
# Returns: unique keys on stdout, one per line
merge_jira_keys() {
    local first_keys="$1"
    local second_keys="$2"

    printf '%s\n%s\n' "$first_keys" "$second_keys" | sed '/^$/d' | sort -u
}

# Search Jira for tickets with a specific PR URL
# Args: pr_url
# Returns: Jira keys on stdout, one per line
search_jira_by_pr() {
    local pr_url="$1"

    if ! jira_enabled; then
        echo ""
        return 0
    fi

    log_debug "    Searching Jira for PR: $pr_url"

    # Use Jira REST API v3 with POST (custom fields require direct API, not acli)
    # The field stores URLs in format [url|url], so we search for the URL substring
    local jql="project = ${JIRA_PROJECT} AND ${JIRA_PR_FIELD} ~ \"${pr_url}\""

    local payload
    payload=$(jq -n --arg jql "$jql" '{jql: $jql, fields: ["key"], maxResults: 100}') || {
        log_debug "    JSON generation failed"
        echo ""
        return 0
    }

    local result
    result=$(curl -s -u "$JIRA_EMAIL:$JIRA_TOKEN" \
        -X POST \
        -H "Content-Type: application/json" \
        "${JIRA_URL}/rest/api/3/search/jql" \
        -d "$payload" 2>/dev/null) || {
        log_debug "    Jira API request failed"
        echo ""
        return 0
    }

    echo "$result" | jq -r '.issues[]?.key' 2>/dev/null
}

# Get Jira issue details including release notes
# Args: jira_key
# Returns: JSON with key, summary, status, release notes
get_jira_issue() {
    local key="$1"

    if ! jira_enabled; then
        echo "{}"
        return 0
    fi

    log_debug "    Fetching Jira: $key"

    # Use REST API v3 to get issue details with custom fields
    local fields="summary,status,${JIRA_RN_TEXT_FIELD},${JIRA_RN_TYPE_FIELD},${JIRA_RN_STATUS_FIELD}"

    curl -s -u "$JIRA_EMAIL:$JIRA_TOKEN" \
        -H "Content-Type: application/json" \
        "${JIRA_URL}/rest/api/3/issue/${key}?fields=${fields}" 2>/dev/null | \
        jq --arg tf "$JIRA_RN_TEXT_FIELD" --arg tyf "$JIRA_RN_TYPE_FIELD" --arg sf "$JIRA_RN_STATUS_FIELD" '{
            key: .key,
            url: ("https://redhat.atlassian.net/browse/" + .key),
            summary: .fields.summary,
            status: .fields.status.name,
            rn_text: (
                if .fields[$tf] then
                    # Extract text from Atlassian Document Format (ADF)
                    # Recursively find all text nodes and join them
                    [.fields[$tf] | .. | objects | select(.type == "text") | .text] | join(" ")
                else
                    null
                end
            ),
            rn_type: (if .fields[$tyf] then .fields[$tyf].value else null end),
            rn_status: (if .fields[$sf] then .fields[$sf].value else null end)
        }' 2>/dev/null || echo "{}"
}

# Get details for multiple Jira issues
# Args: jira_keys (one key per line)
# Returns: JSON array of Jira issue details
get_jira_issues() {
    local keys="$1"
    local jira_json="[]"

    if ! jira_enabled; then
        echo "[]"
        return 0
    fi

    local key issue
    while IFS= read -r key; do
        [[ -z "$key" ]] && continue

        issue=$(get_jira_issue "$key")
        if [[ $(echo "$issue" | jq -r '.key // empty' 2>/dev/null) == "$key" ]]; then
            jira_json=$(echo "$jira_json" | jq --argjson issue "$issue" '. + [$issue]')
        else
            log_debug "    Skipping invalid Jira response for $key"
        fi
    done <<< "$keys"

    echo "$jira_json"
}

# Search for unmatched Jiras (tickets with fixVersion but no matching commit)
# Args: version, matched_keys_json (array of matched Jira keys)
# Returns: JSON array of unmatched tickets
search_unmatched_jiras() {
    local version="$1"
    local matched_keys="$2"

    if ! jira_enabled; then
        echo "[]"
        return 0
    fi

    log_progress "Searching for unmatched Jiras with fixVersion=$version..."

    # Search using acli
    local search_result
    search_result=$(acli jira workitem search \
        --jql "project = ${JIRA_PROJECT} AND fixVersion = '$version'" \
        --fields "key,summary,status" \
        --limit 100 \
        --json 2>/dev/null) || {
        log_warn "Jira search failed"
        echo "[]"
        return 0
    }

    # Filter out matched keys
    echo "$search_result" | jq --argjson matched "$matched_keys" '
        [.[] | select(.key as $k | $matched | index($k) | not)]
    ' 2>/dev/null || echo "[]"
}

# ============================================================================
# COMMIT OPERATIONS
# ============================================================================

# Compare commits between two SHAs
# Args: upstream_repo, base_sha, head_sha
# Returns: JSON array of commits on stdout
compare_commits() {
    local upstream_repo="$1"
    local base_sha="$2"
    local head_sha="$3"

    log_debug "  Comparing $upstream_repo: ${base_sha:0:12}...${head_sha:0:12}"

    gh api "repos/$upstream_repo/compare/${base_sha:0:12}...${head_sha:0:12}" \
        --jq '.commits[] | {
            sha: .sha,
            message: (.commit.message | split("\n")[0]),
            author: .commit.author.name,
            date: .commit.committer.date[:10],
            full_message: .commit.message
        }' 2>/dev/null | jq -s . || {
        log_error "Failed to compare commits in $upstream_repo"
        return 1
    }
}

# Extract cherry-pick SHA from commit message
# Args: full_message
# Returns: SHA on stdout if found, empty otherwise
extract_cherry_pick_sha() {
    local message="$1"

    # Look for "(cherry picked from commit <sha>)"
    if [[ "$message" =~ \(cherry\ picked\ from\ commit\ ([0-9a-f]{40})\) ]]; then
        echo "${BASH_REMATCH[1]}"
        return 0
    fi

    return 1
}

# Resolve PR for a commit SHA
# Args: upstream_repo, sha
# Returns: PR URL on stdout, empty if not found
resolve_pr() {
    local upstream_repo="$1"
    local sha="$2"
    local cache_key="${upstream_repo}:${sha}"

    # Check cache first
    if [[ -n "${PR_CACHE[$cache_key]:-}" ]]; then
        echo "${PR_CACHE[$cache_key]}"
        return 0
    fi

    log_debug "    Resolving PR for ${sha:0:12}"

    local pr_url
    pr_url=$(gh api "repos/$upstream_repo/commits/$sha/pulls" \
        --jq '.[0].html_url // empty' 2>/dev/null || echo "")

    # Cache the result (even if empty)
    PR_CACHE[$cache_key]="$pr_url"
    echo "$pr_url"
}

# Enrich commits with PR, cherry-pick, and Jira information
# Args: upstream_repo, commits_json
# Returns: enriched JSON array on stdout
enrich_commits() {
    local upstream_repo="$1"
    local commits_json="$2"

    local commit_count
    commit_count=$(echo "$commits_json" | jq 'length')

    if [[ $commit_count -eq 0 ]]; then
        echo "[]"
        return 0
    fi

    log_debug "  Enriching $commit_count commits"

    # Use jq to process and enrich each commit
    echo "$commits_json" | jq -c '.[]' | while read -r commit; do
        local sha
        sha=$(echo "$commit" | jq -r '.sha')
        local message
        message=$(echo "$commit" | jq -r '.message')
        local full_message
        full_message=$(echo "$commit" | jq -r '.full_message')

        # Resolve PR
        local pr
        pr=$(resolve_pr "$upstream_repo" "$sha")

        # Extract cherry-pick SHA and resolve original PR
        local cherry_sha
        cherry_sha=$(extract_cherry_pick_sha "$full_message")
        local original_pr=""
        if [[ -n "$cherry_sha" ]]; then
            log_debug "    Resolving original PR for cherry-pick ${cherry_sha:0:12}"
            original_pr=$(resolve_pr "$upstream_repo" "$cherry_sha")
        fi

        # Resolve all Jiras linked to the commit, PR, or original PR
        local jira_keys=""
        local jira_json="[]"

        if jira_enabled; then
            jira_keys=$(extract_jira_keys "$message")

            if [[ -n "$pr" ]]; then
                jira_keys=$(merge_jira_keys "$jira_keys" "$(search_jira_by_pr "$pr")")
            fi

            if [[ -n "$original_pr" ]]; then
                jira_keys=$(merge_jira_keys "$jira_keys" "$(search_jira_by_pr "$original_pr")")
            fi

            # Fetch details for every linked Jira
            jira_json=$(get_jira_issues "$jira_keys")
        fi

        # Add all enrichment fields to commit JSON
        echo "$commit" | jq \
            --arg pr "$pr" \
            --arg original_pr "$original_pr" \
            --argjson jiras "$jira_json" \
            '. + {pr: $pr, original_pr: $original_pr, jiras: $jiras}'
    done | jq -s .
}

# ============================================================================
# REPORT GENERATION
# ============================================================================

# Initialize markdown report
init_report() {
    local report_file="$1"

    cat > "$report_file" <<EOF
# Release Audit: $VERSION

**Snapshot:** $SNAPSHOT_NAME
**Previous:** $PREVIOUS_VERSION
**Generated:** $(date -u '+%Y-%m-%d %H:%M UTC')

---

## Summary

_Processing..._

---

## Components

EOF

    log_debug "Initialized report: $report_file"
}

# Append component section to report
# Args: component, upstream_repo, from_sha, to_sha, commits_json, error_msg
append_component() {
    local component="$1"
    local upstream_repo="$2"
    local from_sha="$3"
    local to_sha="$4"
    local commits_json="$5"
    local error_msg="${6:-}"

    local report_file="$REPORT_FILE"

    # If error, write error section
    if [[ -n "$error_msg" ]]; then
        cat >> "$report_file" <<EOF
### $component

**Error:** $error_msg

---

EOF
        return 0
    fi

    local commit_count
    commit_count=$(echo "$commits_json" | jq 'length')

    # Count dependabot commits
    local dependabot_count
    dependabot_count=$(echo "$commits_json" | jq '[.[] | select(.author == "dependabot[bot]")] | length')

    # Filter out dependabot commits for display
    local display_commits
    display_commits=$(echo "$commits_json" | jq '[.[] | select(.author != "dependabot[bot]")]')
    local display_count
    display_count=$(echo "$display_commits" | jq 'length')

    # Write component header
    cat >> "$report_file" <<EOF
### $component

**Upstream:** $upstream_repo
**Commits:** $commit_count ($display_count shown, $dependabot_count dependabot hidden)
**Range:** [\`${from_sha:0:12}\`](https://github.com/$upstream_repo/commit/$from_sha) → [\`${to_sha:0:12}\`](https://github.com/$upstream_repo/commit/$to_sha)

EOF

    # Write commit table if we have non-dependabot commits
    if [[ $display_count -gt 0 ]]; then
        # Check if any commit has Jira data
        local has_jira
        has_jira=$(echo "$display_commits" | jq '[.[] | select((.jiras // []) | length > 0)] | length > 0')

        if [[ "$has_jira" == "true" ]]; then
            # Table with Jira columns including all 3 release note fields
            cat >> "$report_file" <<EOF
| SHA | Message | Author | Date | PR | Original PR | Jira | RN Type | RN Status | RN Text |
|-----|---------|--------|------|----|----|------|---------|-----------|---------|
EOF

            echo "$display_commits" | jq -r '
                def jira_links:
                    (.jiras // [])
                    | if length > 0 then map("[\(.key)](\(.url))") | join("<br>") else "—" end;
                def jira_field($field):
                    (.jiras // [])
                    | if length > 0 then map(.[$field] // "—") | join("<br>") else "—" end;
                .[] |
                "| [\(.sha[:12])](https://github.com/'"$upstream_repo"'/commit/\(.sha)) | \(.message) | \(.author) | \(.date) | \(if .pr != "" then "[\(.pr | split("/")[-1])](\(.pr))" else "—" end) | \(if .original_pr != "" then "[\(.original_pr | split("/")[-1])](\(.original_pr))" else "—" end) | \(jira_links) | \(jira_field("rn_type")) | \(jira_field("rn_status")) | \(jira_field("rn_text")) |"
            ' >> "$report_file"
        else
            # Table without Jira columns
            cat >> "$report_file" <<EOF
| SHA | Message | Author | Date | PR | Original PR |
|-----|---------|--------|------|----|----|
EOF

            echo "$display_commits" | jq -r '.[] |
                "| [\(.sha[:12])](https://github.com/'"$upstream_repo"'/commit/\(.sha)) | \(.message) | \(.author) | \(.date) | \(if .pr != "" then "[\(.pr | split("/")[-1])](\(.pr))" else "—" end) | \(if .original_pr != "" then "[\(.original_pr | split("/")[-1])](\(.original_pr))" else "—" end) |"
            ' >> "$report_file"
        fi
    elif [[ $commit_count -gt 0 ]]; then
        # All commits are dependabot
        cat >> "$report_file" <<EOF
_All commits are from dependabot._

EOF
    fi

    echo "" >> "$report_file"
    echo "---" >> "$report_file"
    echo "" >> "$report_file"
}

# Generate report of PRs with no linked Jira, grouped by author
generate_missing_jira_report() {
    local report_file="$REPORT_FILE"

    log_progress "  Analyzing PRs without Jira..."

    # Extract commits with PR but no Jira from the markdown report
    # Parse the markdown tables to find commits without Jira
    local missing_jira_file="${REPORT_FILE%.md}-missing-jira.txt"

    cat > "$missing_jira_file" << 'EOF'
# PRs Without Linked Jira Tickets

This report lists all PRs that have no linked Jira ticket, grouped by author.

## By Author

EOF

    # Use a temp file to collect data
    local temp_data=$(mktemp)

    # Extract data from markdown report - look for table rows
    grep -E '^\| \[' "$report_file" | while IFS='|' read -r _ sha message author date pr original_pr rest; do
        # Clean up fields
        sha=$(echo "$sha" | sed 's/\[//;s/\].*//;s/^ //;s/ $//')
        message=$(echo "$message" | sed 's/^ //;s/ $//')
        author=$(echo "$author" | sed 's/^ //;s/ $//')
        pr=$(echo "$pr" | sed 's/^ //;s/ $//')

        # Check if we have PR but the rest doesn't contain Jira project (meaning no Jira)
        if [[ "$pr" != "—" ]] && [[ ! "$rest" =~ ${JIRA_PROJECT} ]]; then
            echo "$author|$pr|$message|$sha"
        fi
    done | sort > "$temp_data"

    # Group by author
    local current_author=""
    local count=0
    while IFS='|' read -r author pr message sha; do
        if [[ "$author" != "$current_author" ]]; then
            if [[ -n "$current_author" ]]; then
                echo "" >> "$missing_jira_file"
            fi
            current_author="$author"
            count=$((count + 1))
            echo "### $author" >> "$missing_jira_file"
            echo "" >> "$missing_jira_file"
        fi
        echo "- $pr: $message (\`$sha\`)" >> "$missing_jira_file"
    done < "$temp_data"

    rm -f "$temp_data"

    if [[ $count -eq 0 ]]; then
        echo "_No PRs without Jira found._" >> "$missing_jira_file"
    fi

    log_progress "    Written to: ${missing_jira_file}"
}

# Generate report of Jiras with missing release note fields, grouped by author
generate_missing_rn_report() {
    local report_file="$REPORT_FILE"

    log_progress "  Analyzing Jiras with incomplete release notes..."

    local missing_rn_file="${REPORT_FILE%.md}-missing-release-notes.txt"

    cat > "$missing_rn_file" << 'EOF'
# Jira Tickets With Incomplete Release Notes

This report lists all Jira tickets that are missing release note fields (type, status, or text), grouped by author.

Note: Tickets marked as "Release Note Not Required" are excluded from this report.

## By Author

EOF

    local temp_data=$(mktemp)

    # Extract data from markdown report - look for rows with Jira but incomplete RN fields
    grep -E '^\| \[' "$report_file" | while IFS='|' read -r _ sha message author date pr original_pr jira rn_type rn_status rn_text rest; do
        # Check if this line has Jira columns (more than 6 pipes)
        if [[ $(echo "$_|$sha|$message|$author|$date|$pr|$original_pr|$jira|$rn_type|$rn_status|$rn_text|$rest" | tr -cd '|' | wc -c) -lt 9 ]]; then
            continue
        fi

        # Clean up fields
        author=$(echo "$author" | sed 's/^ //;s/ $//')
        jira=$(echo "$jira" | sed 's/^ //;s/ $//')
        rn_type=$(echo "$rn_type" | sed 's/^ //;s/ $//')
        rn_status=$(echo "$rn_status" | sed 's/^ //;s/ $//')
        rn_text=$(echo "$rn_text" | sed 's/^ //;s/ $//')

        # Skip if RN type is "Release Note Not Required"
        if [[ "$rn_type" =~ "Release Note Not Required" ]]; then
            continue
        fi

        # Audit each ticket in a multi-ticket cell independently
        local jira_entries rn_types rn_statuses rn_texts
        readarray -t jira_entries < <(printf '%s\n' "$jira" | sed 's/<br>/\n/g')
        readarray -t rn_types < <(printf '%s\n' "$rn_type" | sed 's/<br>/\n/g')
        readarray -t rn_statuses < <(printf '%s\n' "$rn_status" | sed 's/<br>/\n/g')
        readarray -t rn_texts < <(printf '%s\n' "$rn_text" | sed 's/<br>/\n/g')

        local index jira_entry jira_key rn_type_value rn_status_value rn_text_value
        for ((index = 0; index < ${#jira_entries[@]}; index++)); do
            jira_entry="${jira_entries[$index]}"
            rn_type_value="${rn_types[$index]:-—}"
            rn_status_value="${rn_statuses[$index]:-—}"
            rn_text_value="${rn_texts[$index]:-—}"

            # Skip tickets marked as not requiring a release note
            if [[ "$rn_type_value" =~ "Release Note Not Required" ]]; then
                continue
            fi

            if [[ "$jira_entry" =~ ${JIRA_PROJECT} ]] && [[ "$rn_type_value" == "—" || "$rn_status_value" == "—" || "$rn_text_value" == "—" ]]; then
                # Extract just the Jira key
                jira_key=$(echo "$jira_entry" | grep -oE "${JIRA_PROJECT}-[0-9]+" || echo "")
                if [[ -n "$jira_key" ]]; then
                    # Determine which fields are missing
                    missing=""
                    [[ "$rn_type_value" == "—" ]] && missing="${missing}type,"
                    [[ "$rn_status_value" == "—" ]] && missing="${missing}status,"
                    [[ "$rn_text_value" == "—" ]] && missing="${missing}text,"
                    missing=${missing%,}

                    echo "$author|$jira_key|$missing"
                fi
            fi
        done
    done | sort -u > "$temp_data"

    # Group by author
    local current_author=""
    local count=0
    while IFS='|' read -r author jira_key missing; do
        if [[ "$author" != "$current_author" ]]; then
            if [[ -n "$current_author" ]]; then
                echo "" >> "$missing_rn_file"
            fi
            current_author="$author"
            count=$((count + 1))
            echo "### $author" >> "$missing_rn_file"
            echo "" >> "$missing_rn_file"
        fi
        echo "- [$jira_key](https://redhat.atlassian.net/browse/$jira_key): Missing $missing" >> "$missing_rn_file"
    done < "$temp_data"

    rm -f "$temp_data"

    if [[ $count -eq 0 ]]; then
        echo "_No Jiras with incomplete release notes found._" >> "$missing_rn_file"
    fi

    log_progress "    Written to: ${missing_rn_file}"
}

# Update summary section with final stats
finalize_report() {
    local report_file="$REPORT_FILE"

    local total_commits=0

    # Create summary - write it properly with real newlines
    local summary_text="**Total Components:** $COMPONENT_COUNT
**Processed:** $PROCESSED_COUNT
**Successful:** $((PROCESSED_COUNT - ERROR_COUNT))
**Errors:** $ERROR_COUNT"

    # Replace "Processing..." with actual summary using a temp file
    if command -v sed &>/dev/null; then
        local tmp=$(mktemp)
        awk -v summary="$summary_text" '{
            if ($0 == "_Processing..._") {
                print summary
            } else {
                print $0
            }
        }' "$report_file" > "$tmp" && mv "$tmp" "$report_file"
    fi

    log_progress ""
    log_progress "Report written to: $report_file"
    log_progress "  Components processed: $PROCESSED_COUNT/$COMPONENT_COUNT"
    log_progress "  Errors: $ERROR_COUNT"
}

# ============================================================================
# COMPONENT PROCESSING
# ============================================================================

# Process a single component
# Args: component
process_component() {
    local component="$1"

    PROCESSED_COUNT=$((PROCESSED_COUNT + 1))
    log_progress ""
    log_progress "[$PROCESSED_COUNT/$COMPONENT_COUNT] Processing $component..."

    # Get component config
    local upstream="${COMPONENT_CONFIG["${component}:upstream"]}"
    local branch="${COMPONENT_CONFIG["${component}:branch"]}"
    local downstream="openshift-pipelines/${component}"

    # Check if component is in snapshot
    local snapshot_revision="${SNAPSHOT_COMPONENTS["${component}:revision"]:-}"
    if [[ -z "$snapshot_revision" ]]; then
        log_warn "$component: not found in snapshot, skipping"
        append_component "$component" "$upstream" "" "" "[]" "Not found in snapshot"
        ERROR_COUNT=$((ERROR_COUNT + 1))
        return 1
    fi

    # Resolve previous tag SHA
    local from_sha
    if ! from_sha=$(resolve_previous_tag "$component" "$downstream"); then
        append_component "$component" "$upstream" "" "" "[]" "Failed to resolve previous tag v$PREVIOUS_VERSION"
        ERROR_COUNT=$((ERROR_COUNT + 1))
        return 1
    fi

    # Read current snapshot SHA (from HEAD file at snapshot revision)
    local to_sha
    if ! to_sha=$(read_head_file "$downstream" "$snapshot_revision"); then
        log_error "$component: failed to read HEAD file at $snapshot_revision"
        append_component "$component" "$upstream" "$from_sha" "" "[]" "Failed to read HEAD file at snapshot revision"
        ERROR_COUNT=$((ERROR_COUNT + 1))
        return 1
    fi

    # Check if SHAs are the same (no commits)
    if [[ "$from_sha" == "$to_sha" ]]; then
        log_progress "  No changes (${from_sha:0:12})"
        append_component "$component" "$upstream" "$from_sha" "$to_sha" "[]" ""
        return 0
    fi

    # Compare commits
    local commits_json
    if ! commits_json=$(compare_commits "$upstream" "$from_sha" "$to_sha"); then
        append_component "$component" "$upstream" "$from_sha" "$to_sha" "[]" "Failed to compare commits"
        ERROR_COUNT=$((ERROR_COUNT + 1))
        return 1
    fi

    local commit_count
    commit_count=$(echo "$commits_json" | jq 'length')
    log_progress "  Found $commit_count commits"

    # Enrich commits (PR resolution)
    if [[ $commit_count -gt 0 ]]; then
        log_progress "  Enriching commits..."
        if ! commits_json=$(enrich_commits "$upstream" "$commits_json"); then
            log_warn "$component: enrichment failed, using basic commit data"
        fi
    fi

    # Append to report
    append_component "$component" "$upstream" "$from_sha" "$to_sha" "$commits_json" ""

    log_progress "  ✓ Complete"
}

# ============================================================================
# MAIN
# ============================================================================

usage() {
    cat <<EOF
Usage: $0 <snapshot-name> <version>

Arguments:
  snapshot-name    Konflux snapshot name (e.g., openshift-pipelines-core-1-24-20260702-034604-000)
  version          Release version (e.g., 1.24.1)

Environment:
  KONFLUX_SERVER   Konflux API server URL (required)
  KONFLUX_TOKEN    Konflux authentication token (required)
  JIRA_URL         Jira base URL (optional, for Jira integration)
  JIRA_EMAIL       Jira email (optional, for Jira integration)
  JIRA_TOKEN       Jira API token (optional, for Jira integration)
  DEBUG            Enable debug output (optional)

Examples:
  $0 openshift-pipelines-core-1-24-20260702-034604-000 1.24.1
  DEBUG=true $0 openshift-pipelines-core-1-24-20260702-034604-000 1.24.1

EOF
}

main() {
    # Parse arguments
    if [[ $# -ne 2 ]]; then
        usage
        exit 1
    fi

    SNAPSHOT_NAME="$1"
    VERSION="$2"

    # Validate and setup
    log_progress "Release Audit Tool (Snapshot-Based)"
    log_progress "===================================="
    log_progress ""

    validate_env
    parse_version "$VERSION"

    # Create reports directory
    mkdir -p "$REPORTS_DIR"
    REPORT_FILE="$REPORTS_DIR/audit-release-${VERSION}.md"

    # Load configuration
    load_config

    # Fetch and parse snapshot
    log_progress ""
    local snapshot_json
    snapshot_json=$(fetch_snapshot "$SNAPSHOT_NAME")
    extract_snapshot_components "$snapshot_json"

    # Initialize report
    init_report "$REPORT_FILE"

    # Process each component from config
    log_progress ""
    log_progress "Processing components..."

    # Temporary file to collect all commits across components
    local all_commits_file=$(mktemp)
    echo "[]" > "$all_commits_file"

    for component in $(echo "${!COMPONENT_CONFIG[@]}" | tr ' ' '\n' | grep ':upstream$' | sed 's/:upstream$//' | sort); do
        process_component "$component" || true
    done

    # Generate additional reports if Jira is enabled
    if jira_enabled; then
        log_progress ""
        log_progress "Generating additional reports..."

        # Extract all commits from the report
        generate_missing_jira_report
        generate_missing_rn_report
    fi

    # Finalize report
    finalize_report

    log_progress ""
    log_progress "✓ Audit complete!"
}

# Run main if executed directly
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi
