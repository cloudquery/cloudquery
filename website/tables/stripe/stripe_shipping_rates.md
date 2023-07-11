# Table: stripe_shipping_rates

This table shows data for Stripe Shipping Rates.

https://stripe.com/docs/api/shipping_rates

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|active|`bool`|
|delivery_estimate|`json`|
|display_name|`utf8`|
|fixed_amount|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|tax_behavior|`utf8`|
|tax_code|`json`|
|type|`utf8`|