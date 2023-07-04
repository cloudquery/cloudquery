# Table: launchdarkly_projects

This table shows data for LaunchDarkly Projects.

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
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|key|`utf8`|
|include_in_snippet_by_default|`bool`|
|default_client_side_availability|`json`|
|name|`utf8`|
|tags|`list<item: utf8, nullable>`|