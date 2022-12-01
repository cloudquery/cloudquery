# Table: aws_neptune_cluster_parameter_groups

https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBParameters

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_neptune_cluster_parameter_groups:
  - [aws_neptune_cluster_parameter_group_parameters](aws_neptune_cluster_parameter_group_parameters.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|db_cluster_parameter_group_name|String|
|db_parameter_group_family|String|
|description|String|