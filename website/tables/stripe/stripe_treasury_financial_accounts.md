# Table: stripe_treasury_financial_accounts

This table shows data for Stripe Treasury Financial Accounts.

https://stripe.com/docs/api/treasury/financial_accounts

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.
## Relations

The following tables depend on stripe_treasury_financial_accounts:
  - [stripe_treasury_credit_reversals](stripe_treasury_credit_reversals)
  - [stripe_treasury_debit_reversals](stripe_treasury_debit_reversals)
  - [stripe_treasury_inbound_transfers](stripe_treasury_inbound_transfers)
  - [stripe_treasury_outbound_payments](stripe_treasury_outbound_payments)
  - [stripe_treasury_outbound_transfers](stripe_treasury_outbound_transfers)
  - [stripe_treasury_received_credits](stripe_treasury_received_credits)
  - [stripe_treasury_received_debits](stripe_treasury_received_debits)
  - [stripe_treasury_transaction_entries](stripe_treasury_transaction_entries)
  - [stripe_treasury_transactions](stripe_treasury_transactions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|active_features|`list<item: utf8, nullable>`|
|balance|`json`|
|country|`utf8`|
|features|`json`|
|financial_addresses|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|pending_features|`list<item: utf8, nullable>`|
|platform_restrictions|`json`|
|restricted_features|`list<item: utf8, nullable>`|
|status|`utf8`|
|status_details|`json`|
|supported_currencies|`list<item: utf8, nullable>`|