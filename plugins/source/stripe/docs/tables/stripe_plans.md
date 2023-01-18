# Table: stripe_plans

https://stripe.com/docs/api/plans

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
|active|Bool|
|aggregate_usage|String|
|amount|Int|
|amount_decimal|Float|
|billing_scheme|String|
|currency|String|
|deleted|Bool|
|interval|String|
|interval_count|Int|
|livemode|Bool|
|metadata|JSON|
|nickname|String|
|object|String|
|product|JSON|
|tiers|JSON|
|tiers_mode|String|
|transform_usage|JSON|
|trial_period_days|Int|
|usage_type|String|