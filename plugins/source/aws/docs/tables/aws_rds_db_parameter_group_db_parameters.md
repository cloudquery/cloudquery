# Table: aws_rds_db_parameter_group_db_parameters

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Parameter.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_rds_db_parameter_groups](aws_rds_db_parameter_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|db_parameter_group_arn|String|
|allowed_values|String|
|apply_method|String|
|apply_type|String|
|data_type|String|
|description|String|
|is_modifiable|Bool|
|minimum_engine_version|String|
|parameter_name|String|
|parameter_value|String|
|source|String|
|supported_engine_modes|StringArray|