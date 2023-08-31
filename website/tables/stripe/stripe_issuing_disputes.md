# Table: stripe_issuing_disputes

This table shows data for Stripe Issuing Disputes.

https://stripe.com/docs/api/issuing/disputes

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
|balance_transactions|`json`|
|currency|`utf8`|
|evidence|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|status|`utf8`|
|transaction|`json`|
|treasury|`json`|