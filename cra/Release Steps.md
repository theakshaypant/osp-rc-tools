# Release Steps

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
Even though this is an idempotent step, we can checkpoint this and skip this step after it has beenb run/verified once.
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

    check_is_minor -->|true| minor[TODO: Add steps for minor release]
    check_is_minor -->|false| pr_exists{"PR Exists: bot: major.minor Release Action: new-patch?"}

    pr_exists -->|true| check_ci_status{Status of CI on PR?}
    pr_exists -->|false| wait_pr[Sleep and recheck]
    wait_pr --> pr_exists

    check_ci_status -->|Running| wait_ci[Sleep and recheck]
    wait_ci --> check_ci_status
    check_ci_status -->|Green| merge[Approve and Merge PR]
    check_ci_status -->|Red| ERROR([ERROR: Share on Slack @rc])

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

    check_is_minor -->|true| minor[TODO: Add steps for minor release]
    check_is_minor -->|false| pr_exists{"PR Exists: bot: major.minor Update Generated Konflux Config?"}

    pr_exists -->|true| check_ci_status{Status of CI on PR?}
    pr_exists -->|false| wait_pr[Sleep and recheck]
    wait_pr --> pr_exists

    check_ci_status -->|Running| wait_ci[Sleep and recheck]
    wait_ci --> check_ci_status
    check_ci_status -->|Green| merge[Approve and Merge PR]
    check_ci_status -->|Red| ERROR([ERROR: Share on Slack @rc])

    merge --> SUCCESS([SUCCESS])
```

#### Outputs
- update_konflux_config_status: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
  - Once the stage has been run/verified successfully once, we can skip verifying again.
  - Error at this stage implies an issue in `hack` config which needs to be fixed manually

### Apply Konflux Config (apply_konf_config)
#### Inputs
- update_konflux_config_status: should have SUCCESS value

```mermaid
flowchart TD
    check_is_minor{is_minor?}

    check_is_minor -->|true| minor[TODO: Add steps for minor release]
    check_is_minor -->|false| apply_config[Apply `.konflux/openshift-pipelines/1-24` dir to the konflux cluster]

    apply_config --> |success| SUCCESS([SUCCESS])
    apply_config --> |Error| ERROR([ERROR: Share on Slack @rc])
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
    ds_repos([Start: For all downstream repos]) --> check_is_minor{is_minor?}

    check_is_minor -->|true| minor[TODO: Add steps for minor release]
    check_is_minor -->|false| pr_exists{"PR Exists: bot update konflux configuration?"}

    pr_exists -->|true| check_ci_status{Status of CI on PR?}
    pr_exists -->|false| wait_pr[Sleep and recheck]
    wait_pr --> pr_exists

    check_ci_status -->|Running| wait_ci[Sleep and recheck]
    wait_ci --> check_ci_status
    check_ci_status -->|Green| merge[Wait for auto-merge]
    check_ci_status -->|Red| ERROR([ERROR: Share on Slack @rc])

    merge --> SUCCESS([SUCCESS])
```

#### Outputs
- ds_update_konf_config: SUCCESS/ERROR

#### Notes
- Checkpoint: yes
  - Once the stage has been run/verified successfully once, we can skip verifying again.
  - Error at this stage implies an issue in `hack` config which needs to be fixed manually
- These PRs are merged automatically sowe only need to wait for the merge to happen.
  - The flowchart describes the proper e2e flow which also helps in debugging.

---

## Dev Release
### Update Sources Downstream (ds_update_source)
#### Inputs
- ds_update_konf_config: should have SUCCESS value

```mermaid
flowchart TD
    ds_repos([Start: Downstream Repos]) --> check_is_minor{is_minor?}

    check_is_minor -->|true| minor[TODO: Add steps for minor release]
    check_is_minor -->|false| pr_exists{"PR Exists: bot Update release-major.minor.x?"}

    pr_exists -->|true| check_ci_status{Status of CI on PR?}
    pr_exists -->|false| wait_pr[Sleep and recheck]
    wait_pr --> pr_exists
    pr_exists -->|Timeout after t time| check_workflow{Was update_sources workflow run?}

    check_ci_status -->|Running| wait_ci[Sleep and recheck]
    wait_ci --> check_ci_status
    check_ci_status -->|Green| merge[Wait for auto-merge]
    check_ci_status -->|Red| ERROR([ERROR: Share on Slack @rc])

    merge --> SUCCESS([SUCCESS])

    check_workflow -->|false| trigger_workflow[Trigger update_sources workflow]
    trigger_workflow --> workflow_status{Was workflow green for release branch?}
    check_workflow -->|true| workflow_status

    workflow_status -->|true| SUCCESS
    workflow_status -->|false| ERROR
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
    ds_repos([Start: Downstream Repos]) --> on_push_completed{"on-push PipelineRuns completed for HEAD commit on release branch?"}

    on_push_completed -->|true| verify_pr{"Component nudge PR created on operator?"}
    on_push_completed -->|false| fetch_prun[Get PipelineRun name from GitHub]

    fetch_prun --> check_status{Status of PipelineRun on Konflux}

    check_status -->|In-Progress| wait_prun[Sleep and wait]
    wait_prun --> check_status
    check_status -->|Failed| ERROR([ERROR: Share on Slack @rc @component-leads])
    check_status -->|Successful| verify_pr

    verify_pr -->|true| check_pr_status{"PR diff contains the same image SHA as PipelineRun?"}
    verify_pr -->|false| create_new_pr[Close current PR, if any, and create a new one with right image SHA]

    check_pr_status -->|true| ci_passing{"Is CI on the PR green?"}
    check_pr_status -->|false| create_new_pr

    ci_passing -->|true| get_snapshot[Get name of core snapshot created by PipelineRun]
    ci_passing -->|false| ERROR

    create_new_pr --> ci_passing
    get_snapshot --> SUCCESS([SUCCESS])
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
    check_status -->|Failed| ERROR([ERROR: Share on Slack @rc])
    check_status -->|Successful| verify_pr

    verify_pr -->|true| check_pr_status{"PR diff contains the same image SHA as PipelineRun?"}
    verify_pr -->|false| create_new_pr[Close current PR, if any, and create a new one with right image SHA]

    check_pr_status -->|true| ci_passing{"Is CI on the PR green?"}
    check_pr_status -->|false| create_new_pr

    ci_passing -->|true| get_snapshot[Get name of bundle snapshot created by PipelineRun]
    ci_passing -->|false| ERROR

    create_new_pr --> ci_passing
    get_snapshot --> SUCCESS([SUCCESS])
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
  check_status -->|Failed| ERROR([ERROR: Share on Slack @rc])
  check_status -->|Successful| get_snapshot

  get_snapshot --> SUCCESS([SUCCESS])
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
    checkpoint --> ds_update_konf_config[Update Downstream Konflux Config]
    
    hack_ver_gen --> merge_release_action
    merge_release_action --> apply_konf_config
    apply_konf_config --> update_konflux_config
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
  testing --> |Dev/Stage Rejected| fix_upstream[Fix Upstream]
  fix_upstream --> ds_update_source

  subgraph stage[Stage Builds]
    stage_core[Core Build]
    stage_core --> stage_bundle[Bundle Build]
    stage_bundle --> stage_index[Index Build]
  end

  testing --> |Dev Accepted| stage_core
  stage_index --> testing
  
  subgraph prod[Prod Builds]
    prod_core[Core Build]
    prod_core --> prod_bundle[Bundle Build]
    prod_bundle --> prod_index[Index Build]
  end

  testing --> |Stage Accepted| prod_core
```

## References
- `hack` repo: https://github.com/openshift-pipelines/hack
-
