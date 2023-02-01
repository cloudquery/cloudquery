# Table: github_hook_deliveries

The composite primary key for this table is (**org**, **hook_id**, **id**).

## Relations

This table depends on [github_hooks](github_hooks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|hook_id (PK)|Int|
|id (PK)|Int|
|guid|String|
|delivered_at|Timestamp|
|redelivery|Bool|
|duration|Float|
|status|String|
|status_code|Int|
|event|String|
|action|String|
|installation_id|Int|
|repository_id|Int|
|request|JSON|
|response|JSON|