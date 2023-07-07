# Table: stripe_files

This table shows data for Stripe Files.

https://stripe.com/docs/api/files

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|expires_at|`int64`|
|filename|`utf8`|
|links|`json`|
|object|`utf8`|
|purpose|`utf8`|
|size|`int64`|
|title|`utf8`|
|type|`utf8`|
|url|`utf8`|