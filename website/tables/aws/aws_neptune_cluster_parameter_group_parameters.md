# Table: aws_neptune_cluster_parameter_group_parameters

This table shows data for Neptune Cluster Parameter Group Parameters.

https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBParameterGroups

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_neptune_cluster_parameter_groups](aws_neptune_cluster_parameter_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_parameter_group_arn|`utf8`|
|allowed_values|`utf8`|
|apply_method|`utf8`|
|apply_type|`utf8`|
|data_type|`utf8`|
|description|`utf8`|
|is_modifiable|`bool`|
|minimum_engine_version|`utf8`|
|parameter_name|`utf8`|
|parameter_value|`utf8`|
|source|`utf8`|