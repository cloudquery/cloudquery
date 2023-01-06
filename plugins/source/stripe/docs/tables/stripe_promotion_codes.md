# Table: stripe_promotion_codes

https://stripe.com/docs/api/promotion_codes

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
|code|String|
|coupon|JSON|
|created|Int|
|customer|JSON|
|expires_at|Int|
|livemode|Bool|
|max_redemptions|Int|
|metadata|JSON|
|object|String|
|restrictions|JSON|
|times_redeemed|Int|