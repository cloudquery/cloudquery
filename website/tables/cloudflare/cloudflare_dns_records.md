# Table: cloudflare_dns_records

This table shows data for Cloudflare DNS Records.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|meta|`json`|
|data|`json`|
|id (PK)|`utf8`|
|created_on|`timestamp[us, tz=UTC]`|
|modified_on|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|name|`utf8`|
|content|`utf8`|
|zone_id|`utf8`|
|zone_name|`utf8`|
|priority|`int64`|
|ttl|`int64`|
|proxied|`bool`|
|proxiable|`bool`|
|locked|`bool`|