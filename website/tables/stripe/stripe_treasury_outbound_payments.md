# Table: stripe_treasury_outbound_payments

This table shows data for Stripe Treasury Outbound Payments.

https://stripe.com/docs/api/treasury/outbound_payments

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
|customer|`utf8`|
|description|`utf8`|
|destination_payment_method|`utf8`|
|destination_payment_method_details|`json`|
|end_user_details|`json`|
|expected_arrival_date|`int64`|
|financial_account|`utf8`|
|hosted_regulatory_receipt_url|`utf8`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|returned_details|`json`|
|statement_descriptor|`utf8`|
|status|`utf8`|
|status_transitions|`json`|
|transaction|`json`|