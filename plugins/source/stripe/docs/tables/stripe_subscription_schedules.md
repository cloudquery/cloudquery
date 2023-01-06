# Table: stripe_subscription_schedules

https://stripe.com/docs/api/subscription_schedules

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|application|JSON|
|canceled_at|Int|
|completed_at|Int|
|created|Int|
|current_phase|JSON|
|customer|JSON|
|default_settings|JSON|
|end_behavior|String|
|livemode|Bool|
|metadata|JSON|
|object|String|
|phases|JSON|
|released_at|Int|
|released_subscription|JSON|
|status|String|
|subscription|JSON|
|test_clock|JSON|