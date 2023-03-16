# Table: github_hooks

This table shows data for Github Hooks.

The composite primary key for this table is (**org**, **id**).

## Relations

The following tables depend on github_hooks:
  - [github_hook_deliveries](github_hook_deliveries)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|url|String|
|id (PK)|Int|
|type|String|
|name|String|
|test_url|String|
|ping_url|String|
|last_response|JSON|
|config|JSON|
|events|StringArray|
|active|Bool|