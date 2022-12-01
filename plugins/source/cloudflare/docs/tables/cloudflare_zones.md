# Table: cloudflare_zones



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
|name|String|
|development_mode|Int|
|original_name_servers|StringArray|
|original_registrar|String|
|original_dnshost|String|
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
|vanity_name_servers|StringArray|
|betas|StringArray|
|deactivation_reason|String|
|meta|JSON|
|account|JSON|
|verification_key|String|