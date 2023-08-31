# Table: aws_neptune_cluster_parameter_groups

This table shows data for Neptune Cluster Parameter Groups.

https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBParameters

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_neptune_cluster_parameter_groups:
  - [aws_neptune_cluster_parameter_group_parameters](aws_neptune_cluster_parameter_group_parameters)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|db_cluster_parameter_group_arn|`utf8`|
|db_cluster_parameter_group_name|`utf8`|
|db_parameter_group_family|`utf8`|
|description|`utf8`|