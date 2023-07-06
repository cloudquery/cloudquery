# Table: stripe_balance_transactions

This table shows data for Stripe Balance Transactions.

https://stripe.com/docs/api/balance_transactions

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
|available_on|`int64`|
|currency|`utf8`|
|description|`utf8`|
|exchange_rate|`float64`|
|fee|`int64`|
|fee_details|`json`|
|net|`int64`|
|object|`utf8`|
|reporting_category|`utf8`|
|source|`json`|
|status|`utf8`|
|type|`utf8`|