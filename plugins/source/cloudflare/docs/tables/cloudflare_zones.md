# Table: cloudflare_zones


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|id (PK)|String|
|name|String|
|dev_mode|Int|
|original_ns|StringArray|
|original_registrar|String|
|original_dns_host|String|
|created_on|Timestamp|
|modified_on|Timestamp|
|name_servers|StringArray|
|owner|JSON|
|permissions|StringArray|
|plan|JSON|
|plan_pending|JSON|
|status|String|
|paused|Bool|
|type|String|
|host|JSON|
|vanity_ns|StringArray|
|betas|StringArray|
|deact_reason|String|
|meta|JSON|
|account|JSON|
|verification_key|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|