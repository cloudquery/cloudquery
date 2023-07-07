# Table: stripe_refunds

This table shows data for Stripe Refunds.

https://stripe.com/docs/api/refunds

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
|charge|`json`|
|currency|`utf8`|
|description|`utf8`|
|failure_balance_transaction|`json`|
|failure_reason|`utf8`|
|instructions_email|`utf8`|
|metadata|`json`|
|next_action|`json`|
|object|`utf8`|
|payment_intent|`json`|
|reason|`utf8`|
|receipt_number|`utf8`|
|source_transfer_reversal|`json`|
|status|`utf8`|
|transfer_reversal|`json`|