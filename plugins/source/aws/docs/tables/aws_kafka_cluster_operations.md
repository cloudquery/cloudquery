# Table: aws_kafka_cluster_operations

https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-operations.html

The primary key for this table is **arn**.

## Relations
This table depends on [aws_kafka_clusters](aws_kafka_clusters.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|cluster_arn|String|
|tags|JSON|
|client_request_id|String|
|creation_time|Timestamp|
|end_time|Timestamp|
|error_info|JSON|
|operation_state|String|
|operation_steps|JSON|
|operation_type|String|
|source_cluster_info|JSON|
|target_cluster_info|JSON|