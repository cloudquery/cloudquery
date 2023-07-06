# Table: stripe_promotion_codes

This table shows data for Stripe Promotion Codes.

https://stripe.com/docs/api/promotion_codes

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
|code|`utf8`|
|coupon|`json`|
|customer|`json`|
|expires_at|`int64`|
|livemode|`bool`|
|max_redemptions|`int64`|
|metadata|`json`|
|object|`utf8`|
|restrictions|`json`|
|times_redeemed|`int64`|