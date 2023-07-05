# Table: stripe_payouts

This table shows data for Stripe Payouts.

https://stripe.com/docs/api/payouts

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
|arrival_date|`int64`|
|automatic|`bool`|
|balance_transaction|`json`|
|currency|`utf8`|
|description|`utf8`|
|destination|`json`|
|failure_balance_transaction|`json`|
|failure_code|`utf8`|
|failure_message|`utf8`|
|livemode|`bool`|
|metadata|`json`|
|method|`utf8`|
|object|`utf8`|
|original_payout|`json`|
|reconciliation_status|`utf8`|
|reversed_by|`json`|
|source_type|`utf8`|
|statement_descriptor|`utf8`|
|status|`utf8`|
|type|`utf8`|