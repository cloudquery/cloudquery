# Table: aws_docdb_cluster_parameter_groups

This table shows data for Amazon DocumentDB Cluster Parameter Groups.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterParameterGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|parameters|`json`|
|db_cluster_parameter_group_name|`utf8`|
|db_parameter_group_family|`utf8`|
|db_cluster_parameter_group_arn|`utf8`|
|description|`utf8`|