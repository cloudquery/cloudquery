# Table: stripe_coupons

This table shows data for Stripe Coupons.

https://stripe.com/docs/api/coupons

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|amount_off|`int64`|
|applies_to|`json`|
|currency|`utf8`|
|currency_options|`json`|
|deleted|`bool`|
|duration|`utf8`|
|duration_in_months|`int64`|
|livemode|`bool`|
|max_redemptions|`int64`|
|metadata|`json`|
|name|`utf8`|
|object|`utf8`|
|percent_off|`float64`|
|redeem_by|`int64`|
|times_redeemed|`int64`|
|valid|`bool`|