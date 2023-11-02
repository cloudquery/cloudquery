# Table: stripe_file_links

This table shows data for Stripe File Links.

https://stripe.com/docs/api/file_links

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|expired|`bool`|
|expires_at|`int64`|
|file|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|url|`utf8`|