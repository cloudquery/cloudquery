# Table: cloudflare_images



The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|id (PK)|String|
|filename|String|
|metadata|JSON|
|require_signed_urls|Bool|
|variants|StringArray|
|uploaded|Timestamp|