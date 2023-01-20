# Table: mixpanel_engage_revenues

The composite primary key for this table is (**project_id**, **date**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|Int|
|date (PK)|Timestamp|
|amount|Float|
|count|Int|
|paid_count|Int|