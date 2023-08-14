# Table: stripe_invoice_items

This table shows data for Stripe Invoice Items.

https://stripe.com/docs/api/invoiceitems

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|amount|`int64`|
|currency|`utf8`|
|customer|`json`|
|date|`int64`|
|deleted|`bool`|
|description|`utf8`|
|discountable|`bool`|
|discounts|`json`|
|invoice|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|period|`json`|
|plan|`json`|
|price|`json`|
|proration|`bool`|
|quantity|`int64`|
|subscription|`json`|
|subscription_item|`utf8`|
|tax_rates|`json`|
|test_clock|`json`|
|unit_amount|`int64`|
|unit_amount_decimal|`float64`|