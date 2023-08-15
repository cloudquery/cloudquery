# Table: stripe_payment_link_line_items

This table shows data for Stripe Payment Link Line Items.

https://stripe.com/docs/api/payment_links/line_items

The composite primary key for this table is (**id**, **payment_link_id**).

## Relations

This table depends on [stripe_payment_links](stripe_payment_links).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|payment_link_id (PK)|`utf8`|
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