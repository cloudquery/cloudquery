# Table: stripe_credit_notes

https://stripe.com/docs/api/credit_notes

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
|created|Timestamp|
|currency|String|
|customer|JSON|
|customer_balance_transaction|JSON|
|discount_amount|Int|
|discount_amounts|JSON|
|invoice|JSON|
|lines|JSON|
|livemode|Bool|
|memo|String|
|metadata|JSON|
|number|String|
|object|String|
|out_of_band_amount|Int|
|pdf|String|
|reason|String|
|refund|JSON|
|status|String|
|subtotal|Int|
|subtotal_excluding_tax|Int|
|tax_amounts|JSON|
|total|Int|
|total_excluding_tax|Int|
|type|String|
|voided_at|Int|