# Table: stripe_plans

This table shows data for Stripe Plans.

https://stripe.com/docs/api/plans

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
|aggregate_usage|`utf8`|
|amount|`int64`|
|amount_decimal|`float64`|
|billing_scheme|`utf8`|
|currency|`utf8`|
|deleted|`bool`|
|interval|`utf8`|
|interval_count|`int64`|
|livemode|`bool`|
|metadata|`json`|
|nickname|`utf8`|
|object|`utf8`|
|product|`json`|
|tiers|`json`|
|tiers_mode|`utf8`|
|transform_usage|`json`|
|trial_period_days|`int64`|
|usage_type|`utf8`|