# Table: stripe_quotes

https://stripe.com/docs/api/quotes

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|amount_subtotal|Int|
|amount_total|Int|
|application|JSON|
|application_fee_amount|Int|
|application_fee_percent|Float|
|automatic_tax|JSON|
|collection_method|String|
|computed|JSON|
|created|Timestamp|
|currency|String|
|customer|JSON|
|default_tax_rates|JSON|
|description|String|
|discounts|JSON|
|expires_at|Int|
|footer|String|
|from_quote|JSON|
|header|String|
|invoice|JSON|
|invoice_settings|JSON|
|line_items|JSON|
|livemode|Bool|
|metadata|JSON|
|number|String|
|object|String|
|on_behalf_of|JSON|
|status|String|
|status_transitions|JSON|
|subscription|JSON|
|subscription_data|JSON|
|subscription_schedule|JSON|
|test_clock|JSON|
|total_details|JSON|
|transfer_data|JSON|