# Table: stripe_refunds

https://stripe.com/docs/api/refunds

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|amount|Int|
|balance_transaction|JSON|
|charge|JSON|
|currency|String|
|description|String|
|failure_balance_transaction|JSON|
|failure_reason|String|
|instructions_email|String|
|metadata|JSON|
|next_action|JSON|
|object|String|
|payment_intent|JSON|
|reason|String|
|receipt_number|String|
|source_transfer_reversal|JSON|
|status|String|
|transfer_reversal|JSON|