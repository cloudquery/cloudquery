# Table: stripe_topups

https://stripe.com/docs/api/topups

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
|amount|Int|
|balance_transaction|JSON|
|currency|String|
|description|String|
|expected_availability_date|Int|
|failure_code|String|
|failure_message|String|
|livemode|Bool|
|metadata|JSON|
|object|String|
|source|JSON|
|statement_descriptor|String|
|status|String|
|transfer_group|String|
|arrival_date|Int|