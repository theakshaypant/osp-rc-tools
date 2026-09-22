## Setup

### Verify all credentials (verify_creds)
Check for the presence of `.env` file and verify that all credentials are valid by making a request to a service endpoint.

**No checkpointing**, run every time when CRA is invoked.

### Verify all services are reachable (verify_services)
Check if all services are reachable.

### Version Parsing (version_parsing)

#### Inputs
- release_version: x.y.{z}
  - patch version z can be treated as optional, for example, `1.25` is a valid release_version

#### Outputs
- major: x
- minor: y
- patch: z
- is_minor: z == 0 || z == nil
- release_version

---
## Config
### Generate version in hack (hack_ver_gen)
Even though this is an idempotent step, we can checkpoint this and skip this step after it has been run/verified once.
Only needs to be rerun if `hack_ver_gen_completed` has ERROR value and some manual fix/changes were done by the p12n team.

#### Inputs
- major
- minor
- patch
- is_minor

```mermaid
flowchart TD
    file_check{"major.minor.yaml present in hack repo?"}

    file_check -->|false| new_minor[Trigger Release Action - New Release workflow]
    new_minor --> SUCCESS([SUCCESS])

    file_check -->|true| check_is_minor{is_minor?}

    check_is_minor -->|true| minor_release{"version == major.minor AND release-tag == major.minor.0?"}
    minor_release -->|true| SUCCESS
    minor_release -->|false| ERROR([ERROR: Share on Slack @rc])

    check_is_minor -->|false| patch_release{"version == major.minor AND release-tag == release_version?"}
    patch_release -->|true| SUCCESS
    patch_release -->|false| new_patch[Trigger Release Action - New Patch workflow]
    new_patch --> SUCCESS
```

#### Outputs
- hack_ver_gen_completed: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
  - Once the stage has been run/verified successfully once, we can skip verifying again.
  - Error at this stage implies an issue in `hack` config which needs to be fixed manually

### Merge Release Action PR (merge_release_action)
#### Inputs
- hack_ver_gen_completed: should have SUCCESS value
- release_version
- major
- minor
- is_minor

```mermaid
flowchart TD
    check_is_minor{is_minor?}

    check_is_minor -->|true| minor[TODO - Add steps for minor release]
    check_is_minor -->|false| pr_exists{"PR Exists - bot major.minor Release Action new-patch?"}

    pr_exists -->|true| check_ci_status{Status of CI on PR?}
    pr_exists -->|false| wait_pr[Sleep and recheck]
    wait_pr --> pr_exists

    check_ci_status -->|Running| wait_ci[Sleep and recheck]
    wait_ci --> check_ci_status
    check_ci_status -->|Green| merge[Approve and Merge PR]
    check_ci_status -->|Red| ERROR([ERROR - Share on Slack @rc])

    merge --> SUCCESS([SUCCESS])
```

#### Outputs
- merge_release_action_status: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
  - Once the stage has been run/verified successfully once, we can skip verifying again.
  - Error at this stage implies an issue in `hack` config which needs to be fixed manually

### Update Konflux Config (update_konflux_config)
#### Inputs
- merge_release_action_status: should have SUCCESS value

```mermaid
flowchart TD
    check_is_minor{is_minor?}

    check_is_minor -->|true| minor[TODO - Add steps for minor release]
    check_is_minor -->|false| pr_exists{"PR Exists - bot major.minor Update Generated Konflux Config?"}

    pr_exists -->|true| check_ci_status{Status of CI on PR?}
    pr_exists -->|false| wait_pr[Sleep and recheck]
    wait_pr --> pr_exists

    check_ci_status -->|Running| wait_ci[Sleep and recheck]
    wait_ci --> check_ci_status
    check_ci_status -->|Green| merge[Approve and Merge PR]
    check_ci_status -->|Red| ERROR([ERROR - Share on Slack @rc])

    merge --> SUCCESS([SUCCESS])
```

#### Outputs
- update_konflux_config_status: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
  - Once the stage has been run/verified successfully once, we can skip verifying again.
  - Error at this stage implies an issue in `hack` config which needs to be fixed manually

### Gitlab konflux-release-data Configuration (glab_rpa_config)
#### Inputs
- update_konflux_config_status: should have SUCCESS value
- release_date

```mermaid
flowchart TD
    copy_dir[Copy .konflux/konflux-release-data dir in hack repo to<br>konflux-release-data Gitlab repo] -->
    set_date[Set release_date in CDN manifest] -->
    is_latest{Is release_version the latest} --> |true| set_ga[Set ga true in CDN release]
    set_ga --> create_mr[Create MR with the changes]
    is_latest -->|false| create_mr
    create_mr --> SUCCESS([SUCCESS])
```

#### Outputs
- mr_link

#### Notes
- Checkpoint: yes

### Apply Konflux Config (apply_konf_config)
#### Inputs
- update_konflux_config_status: should have SUCCESS value

```mermaid
flowchart TD
    check_is_minor{is_minor?}

    check_is_minor -->|true| minor[TODO - Add steps for minor release]
    check_is_minor -->|false| apply_config[Apply .konflux/openshift-pipelines/1-24 dir to the konflux cluster]

    apply_config --> |success| SUCCESS([SUCCESS])
    apply_config --> |Error| ERROR([ERROR - Share on Slack @rc])
```

#### Outputs
- apply_konf_config_status: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
  - Once the stage has been run/verified successfully once, we can skip verifying again.
  - Error at this stage implies an issue in `hack` config which needs to be fixed manually


### Update Downstream Konflux Config (ds_update_konf_config)
#### Inputs
- apply_konf_config_status: should have SUCCESS value

```mermaid
flowchart TD
    ds_repos([Start - For all downstream repos]) --> check_is_minor{is_minor?}

    check_is_minor -->|true| minor[TODO - Add steps for minor release]
    check_is_minor -->|false| pr_exists{"PR Exists - bot update konflux configuration?"}

    pr_exists -->|true| check_ci_status{Status of CI on PR?}
    pr_exists -->|false| wait_pr[Sleep and recheck]
    wait_pr --> pr_exists

    check_ci_status -->|Running| wait_ci[Sleep and recheck]
    wait_ci --> check_ci_status
    check_ci_status -->|Green| merge[Wait for auto-merge]
    check_ci_status -->|Red| ERROR([ERROR - Share on Slack @rc])

    merge --> SUCCESS([SUCCESS])
```

#### Outputs
- ds_update_konf_config: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
  - Once the stage has been run/verified successfully once, we can skip verifying again.
  - Error at this stage implies an issue in `hack` config which needs to be fixed manually
- These PRs are merged automatically so we only need to wait for the merge to happen.
  - The flowchart describes the proper e2e flow which also helps in debugging.



### Operator Version Update (operator_version)
#### Inputs
- ds_update_konf_config: should have SUCCESS 
- major, minor, patch
- is_minor

```mermaid
flowchart TD
    check_file{operator/project.yaml has correct version fields for current and previous?}
    
    check_file -->|true| SUCCESS([SUCCESS])
    check_file -->|false| check_is_minor{is_minor?}
    
    check_is_minor -->|true| update_major_minor[Update version-major-minor field]
    check_is_minor -->|false| update_patch[Update version-patch field]
    
    update_major_minor -->|Committed| SUCCESS
    update_patch -->|Committed| SUCCESS
    
    update_major_minor -->|Failed| ERROR([ERROR])
    update_patch -->|Failed| ERROR([ERROR])
```

---

## Dev Release
### Update Sources Downstream (ds_update_source)
#### Inputs
- ds_update_konf_config: should have SUCCESS value

```mermaid
flowchart TD
    ds_repos([Start - Downstream Repos]) --> check_is_minor{is_minor?}

    check_is_minor -->|true| minor[TODO - Add steps for minor release]
    check_is_minor -->|false| pr_exists{"PR Exists - bot Update release-major.minor.x?"}

    pr_exists -->|true| check_ci_status{Status of CI on PR?}
    pr_exists -->|false| wait_pr[Sleep and recheck]
    wait_pr --> pr_exists
    pr_exists -->|Timeout after t time| check_workflow{Was update_sources workflow run?}

    check_ci_status -->|Running| wait_ci[Sleep and recheck]
    wait_ci --> check_ci_status
    check_ci_status -->|Green| merge[Wait for auto-merge]
    check_ci_status -->|Red| ERROR_NODE([ERROR - Share on Slack @rc])

    merge --> SUCCESS_NODE([SUCCESS])

    check_workflow -->|false| trigger_workflow[Trigger update_sources workflow]
    trigger_workflow --> workflow_status{Was workflow green for release branch?}
    check_workflow -->|true| workflow_status

    workflow_status -->|true| SUCCESS_NODE
    workflow_status -->|false| ERROR_NODE
```
#### Outputs
- update_ds_konf_config_status: SUCCESS/ERROR

#### Notes
- These PRs are merged automatically so we only need to wait for the merge to happen.
  - The flowchart describes the proper e2e flow which also helps in debugging.
- This step may need to be run again when new builds are required.

--- 

## Builds
### Core Build (core)
#### Inputs
- ds_update_konf_config: should have SUCCESS value
- env: dev/staging/prod

```mermaid
flowchart TD
    ds_repos([Start - Downstream Repos]) --> on_push_completed{"on-push PipelineRuns completed for HEAD commit on release branch?"}

    on_push_completed -->|true| verify_pr{"Component nudge PR created on operator?"}
    on_push_completed -->|false| fetch_prun[Get PipelineRun name from GitHub]

    fetch_prun --> check_status{Status of PipelineRun on Konflux}

    check_status -->|In-Progress| wait_prun[Sleep and wait]
    wait_prun --> check_status
    check_status -->|Failed| ERROR_NODE([ERROR - Share on Slack @rc @component-leads])
    check_status -->|Successful| verify_pr

    verify_pr -->|true| check_pr_status{"PR diff contains the same image SHA as PipelineRun?"}
    verify_pr -->|false| create_new_pr[Close current PR if any and create a new one with right image SHA]

    check_pr_status -->|true| ci_passing{"Is CI on the PR green?"}
    check_pr_status -->|false| create_new_pr

    ci_passing -->|true| get_snapshot[Get name of core snapshot created by PipelineRun]
    ci_passing -->|false| ERROR_NODE

    create_new_pr --> ci_passing
    get_snapshot --> SUCCESS_NODE([SUCCESS])
```

#### Outputs
- core_{env}_status: SUCCESS/ERROR
- nudge_pr_list
- core_snapshot

#### Notes
- TODO

### Bundle (bundle)
#### Inputs
- core_{env}_status: should have SUCCESS value
- env: should match `core_{env}_status` key
- nudge_pr_list

```mermaid
flowchart TD
    merge_nudge[Merge All Nudge PRs] --> bundle_completed{"Bundle on-push PipelineRun completed?"}

    bundle_completed -->|true| verify_pr{"Bundle nudge PR created on operator?"}
    bundle_completed -->|false| fetch_prun[Get bundle-on-push PipelineRun name from GitHub]

    fetch_prun --> check_status{Status of PipelineRun on Konflux}

    check_status -->|In-Progress| wait_prun[Sleep and wait]
    wait_prun --> check_status
    check_status -->|Failed| ERROR_NODE([ERROR - Share on Slack @rc])
    check_status -->|Successful| verify_pr

    verify_pr -->|true| check_pr_status{"PR diff contains the same image SHA as PipelineRun?"}
    verify_pr -->|false| create_new_pr[Close current PR if any and create a new one with right image SHA]

    check_pr_status -->|true| ci_passing{"Is CI on the PR green?"}
    check_pr_status -->|false| create_new_pr

    ci_passing -->|true| get_snapshot[Get name of bundle snapshot created by PipelineRun]
    ci_passing -->|false| ERROR_NODE

    create_new_pr --> ci_passing
    get_snapshot --> SUCCESS_NODE([SUCCESS])
```

#### Outputs
- bundle_{env}_status: SUCCESS/ERROR
- bundle_nudge_pr
- bundle_snapshot

#### Notes
- TODO

### Index (index)
#### Inputs
- bundle_{env}_status: should have SUCCESS value
- env: should match `bundle_{env}_status` key
- bundle_nudge_pr

```mermaid
flowchart TD
  merge_nudge[Merge Bundle Nudge PR] --> index_completed{"Index on-push PipelineRuns completed?"}

  index_completed --> |true| get_snapshot[Get Index snapshots for all compatible OCP versions]
  index_completed --> |false| fetch_prun[Get index-on-push PipelineRun names from GitHub]

  fetch_prun --> check_status{Status of PipelineRuns on Konflux}

  check_status -->|In-Progress| wait_prun[Sleep and wait]
  wait_prun --> check_status
  check_status -->|Failed| ERROR_NODE([ERROR - Share on Slack @rc])
  check_status -->|Successful| get_snapshot

  get_snapshot --> SUCCESS_NODE([SUCCESS])
```

#### Outputs
- index_{env}_status: SUCCESS/ERROR
- index_snapshots

#### Notes
- TODO

---

## End to End Flow
```mermaid
flowchart TD
  start(START)

  subgraph setup[Setup]
    verify_creds[Verify all credentials] --> verify_services[Verify all services are reachable]
    verify_services --> version_parsing[Version Parsing]
  end

  start --> verify_creds


  subgraph config[Configuration]
    version_parsing --> checkpoint{Last saved checkpoint}
    checkpoint --> hack_ver_gen[Generate version in hack]
    checkpoint --> merge_release_action[Merge Release Action PR]
    checkpoint --> update_konflux_config[Update Konflux Config]
    checkpoint --> apply_konf_config[Apply Konflux Config]
    checkpoint --> glab_rpa_config[Gitlab konflux-release-data Configuration]
    checkpoint --> ds_update_konf_config[Update Downstream Konflux Config]
    
    hack_ver_gen --> merge_release_action
    merge_release_action --> apply_konf_config
    apply_konf_config --> update_konflux_config
    update_konflux_config --> glab_rpa_config
    update_konflux_config --> ds_update_konf_config
  end

  subgraph dev[Dev Builds]
    ds_update_source[Update Sources Downstream] --> dev_core[Core Build]
    dev_core --> dev_bundle[Bundle Build]
    dev_bundle --> dev_index[Index Build]
  end

  ds_update_konf_config --> ds_update_source

  subgraph QE
    testing
  end

  dev_index --> testing
  testing -->|Dev or Stage Rejected| fix_upstream[Fix Upstream]
  fix_upstream --> ds_update_source

  subgraph stage[Stage Builds]
    stage_core[Core Build]
    stage_core --> stage_bundle[Bundle Build]
    stage_bundle --> stage_index[Index Build]
  end

  testing -->|Dev Accepted| stage_core
  stage_index --> testing
  
  subgraph prod[Prod Builds]
    prod_core[Core Build]
    prod_core --> prod_bundle[Bundle Build]
    prod_bundle --> prod_index[Index Build]
  end

  testing -->|Stage Accepted| prod_core
```

---

## Missing Steps (To Be Integrated)
### Pyxis Configuration (pyxis_config)
#### Inputs
- rpa_config_status: should have SUCCESS or SKIP value
- is_minor

```mermaid
flowchart TD
    check_is_minor{is_minor?}
    
    check_is_minor -->|false| skip[SKIP - Patch releases do not add components]
    check_is_minor -->|true| check_new_components{New components in this minor release?}
    
    check_new_components -->|false| skip
    check_new_components -->|true| check_pyxis{Pyxis repo configs exist for new components?}
    
    check_pyxis -->|true| SUCCESS_NODE([SUCCESS])
    check_pyxis -->|false| create_mr[Create GitLab MR with Pyxis configs]
    
    create_mr -->|Merged| SUCCESS_NODE
    create_mr -->|Failed| ERROR_NODE([ERROR - Share on Slack @rc])
    
    skip --> SUCCESS_NODE
```

#### Outputs
- pyxis_config_status: SUCCESS/SKIP/ERROR

#### Notes
- Checkpoint: yes
- Usually SKIP for patch releases (no new components)
- Only needed when new images/components added
- GitLab repo: releng/pyxis-repo-configs


#### Outputs
- operator_version_status: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
- Fields used by Konflux for image tagging
- Minor: updates `version-major-minor`, Patch: updates `version-patch`

### OPC Version Update (opc_version)
#### Inputs
- operator_version_status: should have SUCCESS value
- is_minor

```mermaid
flowchart TD
    check_is_minor{is_minor?}
    
    check_is_minor -->|false| skip[SKIP - Patch reuses OPC version]
    check_is_minor -->|true| check_version{pkg/version.json has correct versions?}
    
    check_version -->|true| SUCCESS_NODE([SUCCESS])
    check_version -->|false| update_versions[Update all component versions in version.json]
    
    update_versions -->|PR Created| check_pr{PR CI green?}
    
    check_pr -->|Green| merge_pr[Merge PR]
    check_pr -->|Red| ERROR_NODE([ERROR - Fix CI])
    
    merge_pr --> SUCCESS_NODE
    
    skip --> SUCCESS_NODE
```

#### Outputs
- opc_version_status: SUCCESS/SKIP/ERROR

#### Notes
- Checkpoint: yes
- Minor releases only (major version bumps for components)
- Updates pkg/version.json for pac, tkn, results, etc.
- Requires GITHUB_USER and GITHUB_EMAIL for automated commits

### Configure CLI Release Infrastructure (configure_cli_release)
#### Inputs
- opc_version_status: should have SUCCESS or SKIP value
- release_version
- release_branch

```mermaid
flowchart TD
    check_p12n_sync{p12n-opc upstream/pkg/version.json<br>matches OPC pkg/version.json?}
    
    check_p12n_sync -->|true| check_submodules{serve-tkn-cli submodules<br>point to branch HEADs?}
    check_p12n_sync -->|false| sync_p12n[Sync p12n-opc upstream from OPC]
    
    sync_p12n -->|PR Merged| check_submodules
    sync_p12n -->|PR Failed| error1([ERROR - p12n-opc sync failed])
    
    check_submodules -->|true| check_product_yaml{Product version YAML exists<br>in konflux-release-data?}
    check_submodules -->|false| update_submodules[Update serve-tkn-cli submodules<br>sources/cli sources/opc sources/pac]
    
    update_submodules -->|PR Merged| check_product_yaml
    update_submodules -->|Failed| error2([ERROR - Submodule update failed])
    
    check_product_yaml -->|true| success([SUCCESS])
    check_product_yaml -->|false| create_product_yaml[Create product version YAML<br>in GitLab konflux-release-data]
    
    create_product_yaml -->|MR Merged| success
    create_product_yaml -->|Failed| error3([ERROR - Product version MR failed])
```

#### Outputs
- configure_cli_release_status: SUCCESS/ERROR

#### Notes
- Checkpoint: yes (checkpoints after each substep)
- Required for ALL releases (both minor and patch)
- Sequential flow (based on cli-release skill):
  1. **p12n-opc sync**: Compare `upstream/pkg/version.json` in p12n-opc with `pkg/version.json` in OPC on release branch. Sync if different.
  2. **serve-tkn-cli submodules**: Validate git submodules (sources/cli → tektoncd/cli, sources/opc → openshift-pipelines/opc, sources/pac → openshift-pipelines/pipelines-as-code) point to their branch HEADs. Update with `git submodule update --init --remote --force --checkout` if outdated.
  3. **Product version**: Ensure `data/external/developer-portal/openshift-pipelines/{VERSION}.yaml` exists in GitLab konflux-release-data with correct versionName, ga flag, releaseDate.
- GitLab repo: https://gitlab.cee.redhat.com/releng/konflux-release-data
- **IMPORTANT**: Product version YAML must exist before RP/RPA can be created (next step dependency)


### OLM Catalog Render (olm_catalog_render)
#### Inputs
- bundle_{env}_status: should have SUCCESS value
- env: dev/staging/production

```mermaid
flowchart TD
    trigger_workflow[Trigger render-olm-catalog workflow] --> check_auto{Auto-triggered run exists?}
    
    check_auto -->|true| wait_auto[Wait for auto-run to complete]
    check_auto -->|false| wait_manual[Wait for manual trigger to complete]
    
    wait_auto -->|Completed| dispatch_env[Dispatch with environment param]
    wait_manual -->|Completed| check_env{Environment matches?}
    
    check_env -->|true| verify_catalog{Catalog JSONs created?}
    check_env -->|false| dispatch_env
    
    dispatch_env --> workflow_status{Workflow status?}
    
    workflow_status -->|Success| verify_catalog
    workflow_status -->|Failed| ERROR_NODE([ERROR - Share on Slack @rc])
    
    verify_catalog -->|true| SUCCESS_NODE([SUCCESS])
    verify_catalog -->|false| ERROR_ALT([ERROR - Catalog generation failed])
```

#### Outputs
- olm_catalog_render_{env}_status: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
- Auto-triggers on bundle.yaml changes but with wrong environment (devel)
- Must dispatch manually with correct environment parameter (dev/staging/production)
- Generates File-Based Catalog JSONs required for index builds

### Code Freeze (code_freeze)
#### Inputs
- index_staging_status: should have SUCCESS value

```mermaid
flowchart TD
    check_freeze{code-freeze true in hack config?}
    
    check_freeze -->|true| SUCCESS_NODE([SUCCESS])
    check_freeze -->|false| update_config[Set code-freeze true in hack config]
    
    update_config -->|PR Created| check_pr{PR merged?}
    
    check_pr -->|true| SUCCESS_NODE
    check_pr -->|false| merge_pr[Merge PR]
    
    merge_pr --> SUCCESS_NODE
```

#### Outputs
- code_freeze_status: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
- Done after staging builds complete
- Prevents accidental upstream syncs during release
- Disables update-sources workflow

### Stage Release (stage_release)
#### Inputs
- index_dev_status: should have SUCCESS value
- core_snapshot, bundle_snapshot, index_snapshots

```mermaid
flowchart TD
    create_core_cr[Generate Core Release CR for staging] --> save_core[Save to reports/manifest/stage/]
    save_core --> create_core[oc create -f core-stage.yaml]
    
    create_core -->|Released| create_bundle_cr[Generate Bundle Release CR]
    create_core -->|Failed| ERROR_NODE([ERROR])
    
    create_bundle_cr --> save_bundle[Save to reports/manifest/stage/]
    save_bundle --> create_bundle[oc create -f bundle-stage.yaml]
    
    create_bundle -->|Released| create_index_crs[Generate Index Release CRs for all OCP versions]
    create_bundle -->|Failed| ERROR_NODE
    
    create_index_crs --> save_indexes[Save all to reports/manifest/stage/]
    save_indexes --> create_indexes[oc create -f index-*.yaml]
    
    create_indexes -->|All Released| SUCCESS_NODE([SUCCESS])
    create_indexes -->|Any Failed| ERROR_NODE
```

#### Outputs
- stage_release_status: SUCCESS/ERROR
- stage_core_release, stage_bundle_release, stage_index_releases

#### Notes
- Checkpoint: yes
- **Hard ordering:** core → bundle → index (sequential, never parallel)
- Save all manifests before applying (for audit trail)
- Wait up to 5 minutes per release, continue if pending
- Release CRs use `generateName` (apply with `oc create`, not `oc apply`)

---

### Production Release (prod_release)
#### Inputs
- stage_release_status: should have SUCCESS value
- STAGE_CORE_SNAPSHOT: from staging (MUST reuse, not pick new one)

#### Flow
```mermaid
flowchart TD
    user_gate{User confirms - Ready for production?}
    
    user_gate -->|No| STOP([STOP])
    user_gate -->|Yes| verify_stage{All stage releases succeeded?}
    
    verify_stage -->|No| ERROR_STAGE([ERROR - Complete stage first])
    verify_stage -->|Yes| create_core[Create Core Production Release]
    
    create_core -->|Released| trigger_csv[Trigger operator-update-images workflow]
    create_core -->|Failed| ERROR_NODE([ERROR])
    
    trigger_csv -->|Success| wait_csv_pr{CSV PR merged?}
    trigger_csv -->|Failed| ERROR_NODE
    
    wait_csv_pr -->|Merged| create_bundle[Create Bundle Production Release]
    wait_csv_pr -->|Timeout| ERROR_NODE
    
    create_bundle -->|Released| render_catalog[Render OLM catalog for production]
    create_bundle -->|Failed| ERROR_NODE
    
    render_catalog -->|Success| create_indexes[Create Index Production Releases]
    render_catalog -->|Failed| ERROR_NODE
    
    create_indexes -->|All Released| create_cdn[Create CDN Production Release]
    create_indexes -->|Failed| ERROR_NODE
    
    create_cdn -->|Released| SUCCESS_NODE([SUCCESS])
    create_cdn -->|Failed| ERROR_NODE
```

#### Key Steps

1. **Core Production Release**
   - MUST reuse STAGE_CORE_SNAPSHOT (verified in Step 4.1)
   - Do NOT pick latest snapshot

2. **CSV Update**
   - Trigger `operator-update-images` workflow with `environment=production`
   - Wait for auto-generated PR
   - Verify production registry references
   - Merge PR

3. **Bundle Production Release**
   - Bundle rebuilds with production CSV
   - Wait for new bundle snapshot
   - Create Bundle Release CR

4. **OLM Catalog Render**
   - Dispatch `render-olm-catalog` with `environment=production`
   - Generates catalog JSONs

5. **Index Production Releases**
   - One per OCP version
   - Wait for index snapshots
   - Create Index Release CRs

6. **CDN Production Release**
   - Releases CLI binaries to CDN
   - Creates Release CR for serve-tkn-cli

#### Outputs
- prod_release_status: SUCCESS/ERROR
- prod_core_release, prod_bundle_release, prod_index_releases, prod_cdn_release

#### Notes
- **User gate required** before starting
- **Snapshot reuse critical** for core (same bits tested in stage)
- **Hard ordering:** core → CSV → bundle → catalog → indexes → CDN

---

### Image Copy to Quay (image_copy_quay)
#### Inputs
- stage_index_releases: list of stage index Release CRs

#### Flow
```mermaid
flowchart TD
    extract_digests[Extract IIB image digests from Release CR artifacts] --> login[Login to quay.io]
    
    login -->|Success| copy_images[For each index - skopeo copy from registry to quay.io]
    login -->|Failed| ERROR_AUTH([ERROR - Quay auth failed])
    
    copy_images -->|All copied| verify[Verify images in quay.io]
    copy_images -->|Any failed| retry{Retry count less than 3?}
    
    retry -->|Yes| copy_images
    retry -->|No| ERROR_COPY([ERROR - Image copy failed])
    
    verify -->|Success| SUCCESS_NODE([SUCCESS])
    verify -->|Failed| ERROR_VERIFY([ERROR - Verification failed])
```

#### Execute
```bash
# Login to Quay
echo "$QUAY_PASSWORD" | skopeo login quay.io -u "$QUAY_USERNAME" --password-stdin

# For each index release, extract IIB digest
IIB_DIGEST=$(oc get release ${RELEASE_NAME} -n ${KONFLUX_NS} \
  -o jsonpath='{.status.artifacts.iib-images[0].digest}')

# Copy to Quay
skopeo copy --all --preserve-digests \
  docker://registry.redhat.io/...@${IIB_DIGEST} \
  docker://quay.io/openshift-pipeline/pipelines-index-${OCP}:v${VERSION}-stage
```

#### Outputs
- image_copy_quay_status: SUCCESS/ERROR
- copied_images: list of quay.io URLs

#### Notes
- Temporary step for QE testing (stage images)
- Uses skopeo for multi-arch copy
- May need VPN for registry access
- Fallback: generate copy script if skopeo unavailable

---

## References
- `hack` repo: https://github.com/openshift-pipelines/hack
- `operator` repo: https://github.com/openshift-pipelines/operator
- `opc` repo: https://github.com/openshift-pipelines/opc
- `p12n-opc` repo: https://github.com/openshift-pipelines/p12n-opc
- `serve-tkn-cli` repo: https://github.com/openshift-pipelines/serve-tkn-cli
- GitLab konflux-release-data: https://gitlab.cee.redhat.com/releng/konflux-release-data
- GitLab pyxis-repo-configs: https://gitlab.cee.redhat.com/releng/pyxis-repo-configs
