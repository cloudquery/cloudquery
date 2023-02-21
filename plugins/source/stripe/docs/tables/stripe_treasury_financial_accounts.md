# Table: stripe_treasury_financial_accounts

https://stripe.com/docs/api/treasury_financial_accounts

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.
## Relations

The following tables depend on stripe_treasury_financial_accounts:
  - [stripe_treasury_credit_reversals](stripe_treasury_credit_reversals.md)
  - [stripe_treasury_debit_reversals](stripe_treasury_debit_reversals.md)
  - [stripe_treasury_inbound_transfers](stripe_treasury_inbound_transfers.md)
  - [stripe_treasury_outbound_payments](stripe_treasury_outbound_payments.md)
  - [stripe_treasury_outbound_transfers](stripe_treasury_outbound_transfers.md)
  - [stripe_treasury_received_credits](stripe_treasury_received_credits.md)
  - [stripe_treasury_received_debits](stripe_treasury_received_debits.md)
  - [stripe_treasury_transaction_entries](stripe_treasury_transaction_entries.md)
  - [stripe_treasury_transactions](stripe_treasury_transactions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|active_features|StringArray|
|balance|JSON|
|country|String|
|features|JSON|
|financial_addresses|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|pending_features|StringArray|
|platform_restrictions|JSON|
|restricted_features|StringArray|
|status|String|
|status_details|JSON|
|supported_currencies|StringArray|