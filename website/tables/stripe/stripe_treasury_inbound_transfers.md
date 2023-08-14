# Table: stripe_treasury_inbound_transfers

This table shows data for Stripe Treasury Inbound Transfers.

https://stripe.com/docs/api/treasury/inbound_transfers

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
|cancelable|`bool`|
|created|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|description|`utf8`|
|failure_details|`json`|
|financial_account|`utf8`|
|hosted_regulatory_receipt_url|`utf8`|
|linked_flows|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|origin_payment_method|`utf8`|
|origin_payment_method_details|`json`|
|returned|`bool`|
|statement_descriptor|`utf8`|
|status|`utf8`|
|status_transitions|`json`|
|transaction|`json`|