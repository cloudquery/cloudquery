# Table: stripe_radar_early_fraud_warnings

https://stripe.com/docs/api/radar_early_fraud_warnings

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|actionable|Bool|
|charge|JSON|
|created|Timestamp|
|fraud_type|String|
|livemode|Bool|
|object|String|
|payment_intent|JSON|