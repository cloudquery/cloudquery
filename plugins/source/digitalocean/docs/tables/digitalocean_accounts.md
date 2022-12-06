# Table: digitalocean_accounts



The primary key for this table is **uuid**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|uuid (PK)|String|
|droplet_limit|Int|
|floating_ip_limit|Int|
|reserved_ip_limit|Int|
|volume_limit|Int|
|email|String|
|email_verified|Bool|
|status|String|
|status_message|String|
|team|JSON|