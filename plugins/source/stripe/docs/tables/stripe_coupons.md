# Table: stripe_coupons

https://stripe.com/docs/api/coupons

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|amount_off|Int|
|applies_to|JSON|
|currency|String|
|currency_options|JSON|
|deleted|Bool|
|duration|String|
|duration_in_months|Int|
|livemode|Bool|
|max_redemptions|Int|
|metadata|JSON|
|name|String|
|object|String|
|percent_off|Float|
|redeem_by|Int|
|times_redeemed|Int|
|valid|Bool|