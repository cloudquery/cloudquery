# Table: aws_lightsail_database_parameters

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseParameter.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lightsail_databases](aws_lightsail_databases.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|database_arn|String|
|allowed_values|String|
|apply_method|String|
|apply_type|String|
|data_type|String|
|description|String|
|is_modifiable|Bool|
|parameter_name|String|
|parameter_value|String|