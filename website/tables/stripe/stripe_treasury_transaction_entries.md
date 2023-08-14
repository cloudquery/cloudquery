# Table: stripe_treasury_transaction_entries

This table shows data for Stripe Treasury Transaction Entries.

https://stripe.com/docs/api/treasury/transaction_entries

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
|balance_impact|`json`|
|currency|`utf8`|
|effective_at|`int64`|
|financial_account|`utf8`|
|flow|`utf8`|
|flow_details|`json`|
|flow_type|`utf8`|
|livemode|`bool`|
|object|`utf8`|
|transaction|`json`|
|type|`utf8`|