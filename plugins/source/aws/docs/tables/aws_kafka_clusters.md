# Table: aws_kafka_clusters

https://docs.aws.amazon.com/MSK/2.0/APIReference/v2-clusters-clusterarn.html#v2-clusters-clusterarn-properties

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_kafka_clusters:
  - [aws_kafka_nodes](aws_kafka_nodes.md)
  - [aws_kafka_cluster_operations](aws_kafka_cluster_operations.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|active_operation_arn|String|
|cluster_name|String|
|cluster_type|String|
|creation_time|Timestamp|
|current_version|String|
|provisioned|JSON|
|serverless|JSON|
|state|String|
|state_info|JSON|
|tags|JSON|