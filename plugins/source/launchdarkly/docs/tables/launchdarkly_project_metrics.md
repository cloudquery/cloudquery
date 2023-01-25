# Table: launchdarkly_project_metrics

https://apidocs.launchdarkly.com/tag/Metrics#operation/getMetrics

The composite primary key for this table is (**project_id**, **_id**).

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
|experiment_count|Int|
|_id (PK)|String|
|key|String|
|name|String|
|kind|String|
|_attached_flag_count|Int|
|_access|JSON|
|tags|StringArray|
|_creation_date|Int|
|last_modified|JSON|
|maintainer_id|String|
|_maintainer|JSON|
|description|String|
|is_numeric|Bool|
|success_criteria|String|
|unit|String|
|event_key|String|