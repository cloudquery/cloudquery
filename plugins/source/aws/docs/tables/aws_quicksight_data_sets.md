# Table: aws_quicksight_data_sets

This table shows data for QuickSight Data Sets.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSetSummary.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).
## Relations

The following tables depend on aws_quicksight_data_sets:
  - [aws_quicksight_ingestions](aws_quicksight_ingestions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|column_level_permission_rules_applied|`bool`|
|created_time|`timestamp[us, tz=UTC]`|
|data_set_id|`utf8`|
|import_mode|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|row_level_permission_data_set|`json`|
|row_level_permission_tag_configuration_applied|`bool`|