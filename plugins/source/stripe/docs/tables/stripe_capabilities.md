# Table: stripe_capabilities

https://stripe.com/docs/api/capabilities

The primary key for this table is **id**.

## Relations

This table depends on [stripe_accounts](stripe_accounts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|account|JSON|
|future_requirements|JSON|
|object|String|
|requested|Bool|
|requested_at|Int|
|requirements|JSON|
|status|String|