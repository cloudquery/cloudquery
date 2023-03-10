# Table: aws_dax_clusters

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_dax_Cluster.html

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
|active_nodes|Int|
|cluster_arn|String|
|cluster_discovery_endpoint|JSON|
|cluster_endpoint_encryption_type|String|
|cluster_name|String|
|description|String|
|iam_role_arn|String|
|node_ids_to_remove|StringArray|
|node_type|String|
|nodes|JSON|
|notification_configuration|JSON|
|parameter_group|JSON|
|preferred_maintenance_window|String|
|sse_description|JSON|
|security_groups|JSON|
|status|String|
|subnet_group|String|
|total_nodes|Int|