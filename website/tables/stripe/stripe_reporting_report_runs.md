# Table: stripe_reporting_report_runs

https://stripe.com/docs/api/reporting_report_runs

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
|error|String|
|livemode|Bool|
|object|String|
|parameters|JSON|
|report_type|String|
|result|JSON|
|status|String|
|succeeded_at|Int|