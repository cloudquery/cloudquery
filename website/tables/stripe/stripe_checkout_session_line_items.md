# Table: stripe_checkout_session_line_items

This table shows data for Stripe Checkout Session Line Items.

https://stripe.com/docs/api/checkout/sessions/line_items

The composite primary key for this table is (**id**, **session_id**).

## Relations

This table depends on [stripe_checkout_sessions](stripe_checkout_sessions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|session_id (PK)|`utf8`|
|amount_discount|`int64`|
|amount_subtotal|`int64`|
|amount_tax|`int64`|
|amount_total|`int64`|
|currency|`utf8`|
|description|`utf8`|
|discounts|`json`|
|object|`utf8`|
|price|`json`|
|quantity|`int64`|
|taxes|`json`|