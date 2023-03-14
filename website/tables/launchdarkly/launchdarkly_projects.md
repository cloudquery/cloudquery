# Table: launchdarkly_projects

This table shows data for Launchdarkly Projects.

https://apidocs.launchdarkly.com/tag/Projects#operation/getProjects

The primary key for this table is **id**.

## Relations

The following tables depend on launchdarkly_projects:
  - [launchdarkly_project_environments](launchdarkly_project_environments)
  - [launchdarkly_project_flag_defaults](launchdarkly_project_flag_defaults)
  - [launchdarkly_project_flags](launchdarkly_project_flags)
  - [launchdarkly_project_metrics](launchdarkly_project_metrics)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|key|String|
|include_in_snippet_by_default|Bool|
|default_client_side_availability|JSON|
|name|String|
|tags|StringArray|