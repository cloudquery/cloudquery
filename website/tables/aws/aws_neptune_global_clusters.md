# Table: aws_neptune_global_clusters

This table shows data for Neptune Global Clusters.

https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#DescribeDBInstances

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
|tags|JSON|
|deletion_protection|Bool|
|engine|String|
|engine_version|String|
|global_cluster_arn|String|
|global_cluster_identifier|String|
|global_cluster_members|JSON|
|global_cluster_resource_id|String|
|status|String|
|storage_encrypted|Bool|