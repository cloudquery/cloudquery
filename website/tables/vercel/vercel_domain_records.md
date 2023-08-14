# Table: vercel_domain_records

This table shows data for Vercel Domain Records.

The composite primary key for this table is (**domain_name**, **id**).
It supports incremental syncs.
## Relations

This table depends on [vercel_domains](vercel_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|domain_name (PK)|`utf8`|
|id (PK)|`utf8`|
|slug|`utf8`|
|name|`utf8`|
|type|`utf8`|
|value|`utf8`|
|mx_priority|`int64`|
|priority|`int64`|
|creator|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|ttl|`int64`|