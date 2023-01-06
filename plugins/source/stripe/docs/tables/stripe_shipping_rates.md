# Table: stripe_shipping_rates

https://stripe.com/docs/api/shipping_rates

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|active|Bool|
|created|Int|
|delivery_estimate|JSON|
|display_name|String|
|fixed_amount|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|tax_behavior|String|
|tax_code|JSON|
|type|String|