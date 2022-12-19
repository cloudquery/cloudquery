# Table: datadog_dashboards

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name (PK)|String|
|id (PK)|String|
|author_handle|String|
|created_at|Timestamp|
|description|JSON|
|is_read_only|Bool|
|layout_type|String|
|modified_at|Timestamp|
|title|String|
|url|String|