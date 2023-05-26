# Table: launchdarkly_project_environments

This table shows data for LaunchDarkly Project Environments.

https://apidocs.launchdarkly.com/tag/Environments#operation/getEnvironment

The composite primary key for this table is (**project_id**, **id**).

## Relations

This table depends on [launchdarkly_projects](launchdarkly_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id (PK)|utf8|
|id (PK)|utf8|
|key|utf8|
|name|utf8|
|api_key|utf8|
|mobile_key|utf8|
|color|utf8|
|default_ttl|int64|
|secure_mode|bool|
|default_track_events|bool|
|require_comments|bool|
|confirm_changes|bool|
|tags|list<item: utf8, nullable>|
|approval_settings|json|