# Table: stripe_treasury_debit_reversals

This table shows data for Stripe Treasury Debit Reversals.

https://stripe.com/docs/api/treasury/debit_reversals

The primary key for this table is **id**.

## Relations

This table depends on [stripe_treasury_financial_accounts](stripe_treasury_financial_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|amount|`int64`|
|created|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|financial_account|`utf8`|
|hosted_regulatory_receipt_url|`utf8`|
|linked_flows|`json`|
|livemode|`bool`|
|metadata|`json`|
|network|`utf8`|
|object|`utf8`|
|received_debit|`utf8`|
|status|`utf8`|
|status_transitions|`json`|
|transaction|`json`|