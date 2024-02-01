# Table: aws_quicksight_templates

This table shows data for QuickSight Templates.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_TemplateSummary.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|latest_version_number|`int64`|
|name|`utf8`|
|template_id|`utf8`|