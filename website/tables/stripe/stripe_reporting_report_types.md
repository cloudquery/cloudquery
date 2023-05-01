# Table: stripe_reporting_report_types

This table shows data for Stripe Reporting Report Types.

https://stripe.com/docs/api/reporting/report_type

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|data_available_end|Int|
|data_available_start|Int|
|default_columns|StringArray|
|livemode|Bool|
|name|String|
|object|String|
|updated|Int|
|version|Int|