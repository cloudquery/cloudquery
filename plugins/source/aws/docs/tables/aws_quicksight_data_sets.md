# Table: aws_quicksight_data_sets



The primary key for this table is **arn**.

## Relations

The following tables depend on aws_quicksight_data_sets:
  - [aws_quicksight_ingestions](aws_quicksight_ingestions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|column_level_permission_rules_applied|Bool|
|created_time|Timestamp|
|data_set_id|String|
|import_mode|String|
|last_updated_time|Timestamp|
|name|String|
|row_level_permission_data_set|JSON|
|row_level_permission_tag_configuration_applied|Bool|