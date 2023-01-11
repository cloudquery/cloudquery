# Table: stripe_prices

https://stripe.com/docs/api/prices

The primary key for this table is **id**.
It supports incremental syncs.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|active|Bool|
|billing_scheme|String|
|created|Int|
|currency|String|
|currency_options|JSON|
|custom_unit_amount|JSON|
|deleted|Bool|
|livemode|Bool|
|lookup_key|String|
|metadata|JSON|
|nickname|String|
|object|String|
|product|JSON|
|recurring|JSON|
|tax_behavior|String|
|tiers|JSON|
|tiers_mode|String|
|transform_quantity|JSON|
|type|String|
|unit_amount|Int|
|unit_amount_decimal|Float|