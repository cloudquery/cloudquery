# Table: stripe_reporting_report_types

This table shows data for Stripe Reporting Report Types.

https://stripe.com/docs/api/reporting/report_type

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|data_available_end|`int64`|
|data_available_start|`int64`|
|default_columns|`list<item: utf8, nullable>`|
|livemode|`bool`|
|name|`utf8`|
|object|`utf8`|
|updated|`int64`|
|version|`int64`|