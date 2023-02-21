# Table: stripe_treasury_inbound_transfers

https://stripe.com/docs/api/treasury_inbound_transfers

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
|cancelable|Bool|
|created|Timestamp|
|currency|String|
|description|String|
|failure_details|JSON|
|financial_account|String|
|hosted_regulatory_receipt_url|String|
|linked_flows|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|origin_payment_method|String|
|origin_payment_method_details|JSON|
|returned|Bool|
|statement_descriptor|String|
|status|String|
|status_transitions|JSON|
|transaction|JSON|