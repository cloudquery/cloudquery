# Table: stripe_payment_link_line_items

This table shows data for Stripe Payment Link Line Items.

https://stripe.com/docs/api/payment_links/line_items

The composite primary key for this table is (**id**, **payment_link_id**).

## Relations

This table depends on [stripe_payment_links](stripe_payment_links).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|payment_link_id (PK)|String|
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