# Table: aws_lightsail_database_parameters

This table shows data for Lightsail Database Parameters.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseParameter.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lightsail_databases](aws_lightsail_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|database_arn|`utf8`|
|allowed_values|`utf8`|
|apply_method|`utf8`|
|apply_type|`utf8`|
|data_type|`utf8`|
|description|`utf8`|
|is_modifiable|`bool`|
|parameter_name|`utf8`|
|parameter_value|`utf8`|