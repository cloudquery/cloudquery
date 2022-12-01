# Table: cloudflare_waf_overrides



The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|zone_id|String|
|id (PK)|String|
|description|String|
|urls|StringArray|
|priority|Int|
|groups|JSON|
|rewrite_action|JSON|
|rules|JSON|
|paused|Bool|