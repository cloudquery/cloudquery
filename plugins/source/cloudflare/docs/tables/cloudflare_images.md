# Table: cloudflare_images


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|id (PK)|String|
|filename|String|
|metadata|JSON|
|require_signed_ur_ls|Bool|
|variants|StringArray|
|uploaded|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|