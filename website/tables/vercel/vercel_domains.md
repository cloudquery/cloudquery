# Table: vercel_domains

This table shows data for Vercel Domains.

The primary key for this table is **id**.
It supports incremental syncs.
## Relations

The following tables depend on vercel_domains:
  - [vercel_domain_records](vercel_domain_records)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|bought_at|`timestamp[us, tz=UTC]`|
|cdn_enabled|`bool`|
|config_verified_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|expires_at|`timestamp[us, tz=UTC]`|
|intended_nameservers|`list<item: utf8, nullable>`|
|custom_nameservers|`list<item: utf8, nullable>`|
|name|`utf8`|
|nameservers|`list<item: utf8, nullable>`|
|ns_verified_at|`timestamp[us, tz=UTC]`|
|ordered_at|`timestamp[us, tz=UTC]`|
|renew|`bool`|
|service_type|`utf8`|
|transfer_started_at|`timestamp[us, tz=UTC]`|
|transferred_at|`timestamp[us, tz=UTC]`|
|txt_verified_at|`timestamp[us, tz=UTC]`|
|verification_record|`utf8`|
|verified|`bool`|
|zone|`bool`|
|creator|`json`|