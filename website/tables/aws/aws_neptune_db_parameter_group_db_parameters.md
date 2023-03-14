# Table: aws_neptune_db_parameter_group_db_parameters

This table shows data for Neptune DB Parameter Group DB Parameters.

https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBClusterParameters

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_neptune_db_parameter_groups](aws_neptune_db_parameter_groups).

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