# Table: aws_rds_db_parameter_groups

This table shows data for Amazon Relational Database Service (RDS) DB Parameter Groups.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBParameterGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_rds_db_parameter_groups:
  - [aws_rds_db_parameter_group_db_parameters](aws_rds_db_parameter_group_db_parameters)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|db_parameter_group_arn|`utf8`|
|db_parameter_group_family|`utf8`|
|db_parameter_group_name|`utf8`|
|description|`utf8`|