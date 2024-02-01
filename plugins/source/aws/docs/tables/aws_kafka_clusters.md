# Table: aws_kafka_clusters

This table shows data for Kafka Clusters.

https://docs.aws.amazon.com/MSK/2.0/APIReference/v2-clusters-clusterarn.html#v2-clusters-clusterarn-properties

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_kafka_clusters:
  - [aws_kafka_cluster_operations](aws_kafka_cluster_operations.md)
  - [aws_kafka_nodes](aws_kafka_nodes.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|active_operation_arn|`utf8`|
|cluster_arn|`utf8`|
|cluster_name|`utf8`|
|cluster_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|current_version|`utf8`|
|provisioned|`json`|
|serverless|`json`|
|state|`utf8`|
|state_info|`json`|
|tags|`json`|