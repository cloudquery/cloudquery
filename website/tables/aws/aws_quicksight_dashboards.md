# Table: aws_quicksight_dashboards

This table shows data for QuickSight Dashboards.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DashboardSummary.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|dashboard_id|`utf8`|
|last_published_time|`timestamp[us, tz=UTC]`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|published_version_number|`int64`|