# Table: aws_neptune_global_clusters

This table shows data for Neptune Global Clusters.

https://docs.aws.amazon.com/neptune/latest/userguide/api-global-dbs.html#GlobalCluster

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|deletion_protection|`bool`|
|engine|`utf8`|
|engine_version|`utf8`|
|global_cluster_arn|`utf8`|
|global_cluster_identifier|`utf8`|
|global_cluster_members|`json`|
|global_cluster_resource_id|`utf8`|
|status|`utf8`|
|storage_encrypted|`bool`|