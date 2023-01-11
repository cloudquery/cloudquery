# Table: stripe_treasury_transactions

https://stripe.com/docs/api/treasury_transactions

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.
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
|created (Incremental Key)|Timestamp|
|amount|Int|
|balance_impact|JSON|
|currency|String|
|description|String|
|entries|JSON|
|financial_account|String|
|flow|String|
|flow_details|JSON|
|flow_type|String|
|livemode|Bool|
|object|String|
|status|String|
|status_transitions|JSON|