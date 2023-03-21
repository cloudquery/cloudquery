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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
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