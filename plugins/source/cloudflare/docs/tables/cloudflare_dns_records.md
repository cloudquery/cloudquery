# Table: cloudflare_dns_records


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|created_on|Timestamp|
|modified_on|Timestamp|
|type|String|
|name|String|
|content|String|
|id (PK)|String|
|zone_id|String|
|zone_name|String|
|priority|Int|
|ttl|Int|
|proxied|Bool|
|proxiable|Bool|
|locked|Bool|
|meta|JSON|
|data|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|