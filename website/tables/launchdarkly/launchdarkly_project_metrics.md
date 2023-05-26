# Table: launchdarkly_project_metrics

This table shows data for LaunchDarkly Project Metrics.

https://apidocs.launchdarkly.com/tag/Metrics#operation/getMetrics

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
|experiment_count|int64|
|id (PK)|utf8|
|key|utf8|
|name|utf8|
|kind|utf8|
|attached_flag_count|int64|
|access|json|
|tags|list<item: utf8, nullable>|
|creation_date|int64|
|last_modified|json|
|maintainer_id|utf8|
|maintainer|json|
|description|utf8|
|is_numeric|bool|
|success_criteria|utf8|
|unit|utf8|
|event_key|utf8|