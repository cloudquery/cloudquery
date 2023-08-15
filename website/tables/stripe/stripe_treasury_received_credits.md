# Table: stripe_treasury_received_credits

This table shows data for Stripe Treasury Received Credits.

https://stripe.com/docs/api/treasury/received_credits

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
|description|`utf8`|
|failure_code|`utf8`|
|financial_account|`utf8`|
|hosted_regulatory_receipt_url|`utf8`|
|initiating_payment_method_details|`json`|
|linked_flows|`json`|
|livemode|`bool`|
|network|`utf8`|
|object|`utf8`|
|reversal_details|`json`|
|status|`utf8`|
|transaction|`json`|