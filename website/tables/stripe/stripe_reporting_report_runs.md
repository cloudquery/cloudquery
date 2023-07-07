# Table: stripe_reporting_report_runs

This table shows data for Stripe Reporting Report Runs.

https://stripe.com/docs/api/reporting/report_run

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|error|`utf8`|
|livemode|`bool`|
|object|`utf8`|
|parameters|`json`|
|report_type|`utf8`|
|result|`json`|
|status|`utf8`|
|succeeded_at|`int64`|