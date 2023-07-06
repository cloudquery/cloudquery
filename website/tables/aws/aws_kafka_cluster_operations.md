# Table: aws_kafka_cluster_operations

This table shows data for Kafka Cluster Operations.

https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-operations.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_kafka_clusters](aws_kafka_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|cluster_arn|`utf8`|
|tags|`json`|
|client_request_id|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|error_info|`json`|
|operation_arn|`utf8`|
|operation_state|`utf8`|
|operation_steps|`json`|
|operation_type|`utf8`|
|source_cluster_info|`json`|
|target_cluster_info|`json`|
|vpc_connection_info|`json`|