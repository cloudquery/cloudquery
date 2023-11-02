# Table: stripe_prices

This table shows data for Stripe Prices.

https://stripe.com/docs/api/prices

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
|billing_scheme|`utf8`|
|currency|`utf8`|
|currency_options|`json`|
|custom_unit_amount|`json`|
|deleted|`bool`|
|livemode|`bool`|
|lookup_key|`utf8`|
|metadata|`json`|
|nickname|`utf8`|
|object|`utf8`|
|product|`json`|
|recurring|`json`|
|tax_behavior|`utf8`|
|tiers|`json`|
|tiers_mode|`utf8`|
|transform_quantity|`json`|
|type|`utf8`|
|unit_amount|`int64`|
|unit_amount_decimal|`float64`|