# Table: aws_quicksight_templates

This table shows data for QuickSight Templates.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_TemplateSummary.html

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
|last_updated_time|`timestamp[us, tz=UTC]`|
|latest_version_number|`int64`|
|name|`utf8`|
|template_id|`utf8`|