# Table: stripe_credit_notes

This table shows data for Stripe Credit Notes.

https://stripe.com/docs/api/credit_notes

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|amount|`int64`|
|amount_shipping|`int64`|
|created|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|customer|`json`|
|customer_balance_transaction|`json`|
|discount_amount|`int64`|
|discount_amounts|`json`|
|invoice|`json`|
|lines|`json`|
|livemode|`bool`|
|memo|`utf8`|
|metadata|`json`|
|number|`utf8`|
|object|`utf8`|
|out_of_band_amount|`int64`|
|pdf|`utf8`|
|reason|`utf8`|
|refund|`json`|
|shipping_cost|`json`|
|status|`utf8`|
|subtotal|`int64`|
|subtotal_excluding_tax|`int64`|
|tax_amounts|`json`|
|total|`int64`|
|total_excluding_tax|`int64`|
|type|`utf8`|
|voided_at|`int64`|