# Table: stripe_treasury_transactions

This table shows data for Stripe Treasury Transactions.

https://stripe.com/docs/api/treasury/transactions

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.
## Relations

This table depends on [stripe_treasury_financial_accounts](stripe_treasury_financial_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|amount|`int64`|
|balance_impact|`json`|
|currency|`utf8`|
|description|`utf8`|
|entries|`json`|
|financial_account|`utf8`|
|flow|`utf8`|
|flow_details|`json`|
|flow_type|`utf8`|
|livemode|`bool`|
|object|`utf8`|
|status|`utf8`|
|status_transitions|`json`|