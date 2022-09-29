# Table: cloudflare_certificate_packs


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|zone_id|String|
|id (PK)|String|
|type|String|
|hosts|StringArray|
|certificates|JSON|
|primary_certificate|String|
|validation_records|JSON|
|validation_errors|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|