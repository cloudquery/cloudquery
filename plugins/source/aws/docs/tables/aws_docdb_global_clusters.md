# Table: aws_docdb_global_clusters

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_GlobalCluster.html

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
|arn (PK)|String|
|database_name|String|
|deletion_protection|Bool|
|engine|String|
|engine_version|String|
|global_cluster_arn|String|
|global_cluster_identifier|String|
|global_cluster_members|JSON|
|global_cluster_resource_id|String|
|status|String|
|storage_encrypted|Bool|