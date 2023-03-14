# Table: cloudflare_dns_records

This table shows data for Cloudflare DNS Records.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|meta|JSON|
|data|JSON|
|id (PK)|String|
|created_on|Timestamp|
|modified_on|Timestamp|
|type|String|
|name|String|
|content|String|
|zone_id|String|
|zone_name|String|
|priority|Int|
|ttl|Int|
|proxied|Bool|
|proxiable|Bool|
|locked|Bool|