# Table: stripe_topups

This table shows data for Stripe Topups.

https://stripe.com/docs/api/topups

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|amount|`int64`|
|balance_transaction|`json`|
|currency|`utf8`|
|description|`utf8`|
|expected_availability_date|`int64`|
|failure_code|`utf8`|
|failure_message|`utf8`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|source|`json`|
|statement_descriptor|`utf8`|
|status|`utf8`|
|transfer_group|`utf8`|
|arrival_date|`int64`|