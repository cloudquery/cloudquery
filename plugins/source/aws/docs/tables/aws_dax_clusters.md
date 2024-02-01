# Table: aws_dax_clusters

This table shows data for Dax Clusters.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_dax_Cluster.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|active_nodes|`int64`|
|cluster_arn|`utf8`|
|cluster_discovery_endpoint|`json`|
|cluster_endpoint_encryption_type|`utf8`|
|cluster_name|`utf8`|
|description|`utf8`|
|iam_role_arn|`utf8`|
|node_ids_to_remove|`list<item: utf8, nullable>`|
|node_type|`utf8`|
|nodes|`json`|
|notification_configuration|`json`|
|parameter_group|`json`|
|preferred_maintenance_window|`utf8`|
|sse_description|`json`|
|security_groups|`json`|
|status|`utf8`|
|subnet_group|`utf8`|
|total_nodes|`int64`|