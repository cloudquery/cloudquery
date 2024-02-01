# Table: aws_neptune_db_parameter_groups

This table shows data for Neptune DB Parameter Groups.

https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBClusterParameterGroups

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_neptune_db_parameter_groups:
  - [aws_neptune_db_parameter_group_db_parameters](aws_neptune_db_parameter_group_db_parameters.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|db_parameter_group_arn|`utf8`|
|db_parameter_group_family|`utf8`|
|db_parameter_group_name|`utf8`|
|description|`utf8`|