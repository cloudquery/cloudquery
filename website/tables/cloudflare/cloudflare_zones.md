# Table: cloudflare_zones

This table shows data for Cloudflare Zones.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|development_mode|`int64`|
|original_name_servers|`list<item: utf8, nullable>`|
|original_registrar|`utf8`|
|original_dnshost|`utf8`|
|created_on|`timestamp[us, tz=UTC]`|
|modified_on|`timestamp[us, tz=UTC]`|
|name_servers|`list<item: utf8, nullable>`|
|owner|`json`|
|permissions|`list<item: utf8, nullable>`|
|plan|`json`|
|plan_pending|`json`|
|status|`utf8`|
|paused|`bool`|
|type|`utf8`|
|host|`json`|
|vanity_name_servers|`list<item: utf8, nullable>`|
|betas|`list<item: utf8, nullable>`|
|deactivation_reason|`utf8`|
|meta|`json`|
|account|`json`|
|verification_key|`utf8`|