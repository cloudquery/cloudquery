# Table: stripe_payouts

https://stripe.com/docs/api/payouts

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
|arrival_date|Int|
|automatic|Bool|
|balance_transaction|JSON|
|currency|String|
|description|String|
|destination|JSON|
|failure_balance_transaction|JSON|
|failure_code|String|
|failure_message|String|
|livemode|Bool|
|metadata|JSON|
|method|String|
|object|String|
|original_payout|JSON|
|reversed_by|JSON|
|source_type|String|
|statement_descriptor|String|
|status|String|
|type|String|