# Table: datadog_dashboard_lists

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name (PK)|String|
|id (PK)|Int|
|author|JSON|
|created|Timestamp|
|dashboard_count|Int|
|is_favorite|Bool|
|modified|Timestamp|
|name|String|
|type|String|