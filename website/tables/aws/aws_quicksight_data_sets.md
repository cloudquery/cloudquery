# Table: aws_quicksight_data_sets

This table shows data for QuickSight Data Sets.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSetSummary.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Relations

The following tables depend on aws_quicksight_data_sets:
  - [aws_quicksight_ingestions](aws_quicksight_ingestions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|column_level_permission_rules_applied|`bool`|
|created_time|`timestamp[us, tz=UTC]`|
|data_set_id|`utf8`|
|import_mode|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|row_level_permission_data_set|`json`|
|row_level_permission_tag_configuration_applied|`bool`|