# Table: stripe_sigma_scheduled_query_runs

https://stripe.com/docs/api/sigma_scheduled_query_runs

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created|Timestamp|
|data_load_time|Int|
|error|JSON|
|file|JSON|
|livemode|Bool|
|object|String|
|result_available_until|Int|
|sql|String|
|status|String|
|title|String|