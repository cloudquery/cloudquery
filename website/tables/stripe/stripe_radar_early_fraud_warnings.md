# Table: stripe_radar_early_fraud_warnings

This table shows data for Stripe Radar Early Fraud Warnings.

https://stripe.com/docs/api/radar/early_fraud_warnings

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|actionable|`bool`|
|charge|`json`|
|created|`timestamp[us, tz=UTC]`|
|fraud_type|`utf8`|
|livemode|`bool`|
|object|`utf8`|
|payment_intent|`json`|