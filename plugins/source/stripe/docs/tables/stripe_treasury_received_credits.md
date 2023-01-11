# Table: stripe_treasury_received_credits

https://stripe.com/docs/api/treasury_received_credits

The primary key for this table is **id**.

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
|amount|Int|
|created|Timestamp|
|currency|String|
|description|String|
|failure_code|String|
|financial_account|String|
|hosted_regulatory_receipt_url|String|
|initiating_payment_method_details|JSON|
|linked_flows|JSON|
|livemode|Bool|
|network|String|
|object|String|
|reversal_details|JSON|
|status|String|
|transaction|JSON|