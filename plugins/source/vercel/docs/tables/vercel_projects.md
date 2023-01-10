# Table: vercel_projects

The primary key for this table is **id**.

## Relations

The following tables depend on vercel_projects:
  - [vercel_project_envs](vercel_project_envs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|account_id|String|
|auto_expose_system_envs|Bool|
|build_command|String|
|command_for_ignoring_build_step|String|
|created_at|Timestamp|
|dev_command|String|
|directory_listing|Bool|
|env|JSON|
|framework|String|
|git_fork_protection|Bool|
|install_command|String|
|name|String|
|node_version|String|
|output_directory|String|
|public_source|Bool|
|root_directory|String|
|serverless_function_region|String|
|source_files_outside_root_directory|Bool|
|updated_at|Timestamp|
|live|Bool|
|latest_deployments|JSON|
|transfer_started_at|Timestamp|
|transfer_completed_at|Timestamp|
|transferred_from_account_id|String|