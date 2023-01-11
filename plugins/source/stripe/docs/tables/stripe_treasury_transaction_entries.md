# Table: stripe_treasury_transaction_entries

https://stripe.com/docs/api/treasury_transaction_entries

The primary key for this table is **id**.
It supports incremental syncs.
## Relations

This table depends on [stripe_treasury_financial_accounts](stripe_treasury_financial_accounts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|balance_impact|JSON|
|created|Int|
|currency|String|
|effective_at|Int|
|financial_account|String|
|flow|String|
|flow_details|JSON|
|flow_type|String|
|livemode|Bool|
|object|String|
|transaction|JSON|
|type|String|