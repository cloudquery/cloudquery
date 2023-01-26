# Table: launchdarkly_project_metrics

https://apidocs.launchdarkly.com/tag/Metrics#operation/getMetrics

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
|experiment_count|Int|
|id (PK)|String|
|key|String|
|name|String|
|kind|String|
|attached_flag_count|Int|
|access|JSON|
|tags|StringArray|
|creation_date|Int|
|last_modified|JSON|
|maintainer_id|String|
|maintainer|JSON|
|description|String|
|is_numeric|Bool|
|success_criteria|String|
|unit|String|
|event_key|String|