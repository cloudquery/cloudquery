# Table: launchdarkly_project_environments

https://apidocs.launchdarkly.com/tag/Environments#operation/getEnvironment

The composite primary key for this table is (**project_id**, **id**).

## Relations

This table depends on [launchdarkly_projects](launchdarkly_projects.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|id (PK)|String|
|key|String|
|name|String|
|api_key|String|
|mobile_key|String|
|color|String|
|default_ttl|Int|
|secure_mode|Bool|
|default_track_events|Bool|
|require_comments|Bool|
|confirm_changes|Bool|
|tags|StringArray|
|approval_settings|JSON|