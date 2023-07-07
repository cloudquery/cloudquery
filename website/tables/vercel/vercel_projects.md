# Table: vercel_projects

This table shows data for Vercel Projects.

The primary key for this table is **id**.
It supports incremental syncs.
## Relations

The following tables depend on vercel_projects:
  - [vercel_project_envs](vercel_project_envs)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|account_id|`utf8`|
|auto_expose_system_envs|`bool`|
|build_command|`utf8`|
|command_for_ignoring_build_step|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|dev_command|`utf8`|
|directory_listing|`bool`|
|env|`json`|
|framework|`utf8`|
|git_fork_protection|`bool`|
|install_command|`utf8`|
|name|`utf8`|
|node_version|`utf8`|
|output_directory|`utf8`|
|public_source|`bool`|
|root_directory|`utf8`|
|serverless_function_region|`utf8`|
|source_files_outside_root_directory|`bool`|
|updated_at|`timestamp[us, tz=UTC]`|
|live|`bool`|
|latest_deployments|`json`|
|transfer_started_at|`timestamp[us, tz=UTC]`|
|transfer_completed_at|`timestamp[us, tz=UTC]`|
|transferred_from_account_id|`utf8`|