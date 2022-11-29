# Table: aws_kafka_nodes

https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-nodes.html#ListNodes

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
|added_to_cluster_time|String|
|broker_node_info|JSON|
|instance_type|String|
|node_type|String|
|zookeeper_node_info|JSON|