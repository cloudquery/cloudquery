# Table: stripe_checkout_session_line_items

This table shows data for Stripe Checkout Session Line Items.

https://stripe.com/docs/api/checkout/sessions/line_items

The composite primary key for this table is (**id**, **session_id**).

## Relations

This table depends on [stripe_checkout_sessions](stripe_checkout_sessions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|session_id (PK)|String|
|amount_discount|Int|
|amount_subtotal|Int|
|amount_tax|Int|
|amount_total|Int|
|currency|String|
|description|String|
|discounts|JSON|
|object|String|
|price|JSON|
|quantity|Int|
|taxes|JSON|