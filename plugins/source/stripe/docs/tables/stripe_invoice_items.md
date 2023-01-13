# Table: stripe_invoice_items

https://stripe.com/docs/api/invoiceitems

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|amount|Int|
|currency|String|
|customer|JSON|
|date|Int|
|deleted|Bool|
|description|String|
|discountable|Bool|
|discounts|JSON|
|invoice|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|period|JSON|
|plan|JSON|
|price|JSON|
|proration|Bool|
|quantity|Int|
|subscription|JSON|
|subscription_item|String|
|tax_rates|JSON|
|test_clock|JSON|
|unit_amount|Int|
|unit_amount_decimal|Float|