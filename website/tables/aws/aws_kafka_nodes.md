# Table: aws_kafka_nodes

This table shows data for Kafka Nodes.

https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-nodes.html#ListNodes

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
|added_to_cluster_time|`utf8`|
|broker_node_info|`json`|
|instance_type|`utf8`|
|node_arn|`utf8`|
|node_type|`utf8`|
|zookeeper_node_info|`json`|