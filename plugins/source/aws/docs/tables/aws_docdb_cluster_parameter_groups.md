# Table: aws_docdb_cluster_parameter_groups

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterParameterGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|parameters|JSON|
|db_cluster_parameter_group_name|String|
|db_parameter_group_family|String|
|db_cluster_parameter_group_arn|String|
|description|String|