# Table: stripe_transfers

This table shows data for Stripe Transfers.

https://stripe.com/docs/api/transfers

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
|amount_reversed|`int64`|
|balance_transaction|`json`|
|currency|`utf8`|
|description|`utf8`|
|destination|`json`|
|destination_payment|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|reversals|`json`|
|reversed|`bool`|
|source_transaction|`json`|
|source_type|`utf8`|
|transfer_group|`utf8`|