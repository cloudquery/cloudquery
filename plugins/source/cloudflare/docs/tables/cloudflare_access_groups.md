# Table: cloudflare_access_groups


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|zone_id|String|
|id (PK)|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|name|String|
|include|JSON|
|exclude|JSON|
|require|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|