# Table: aws_elasticache_reserved_cache_nodes

This table shows data for Elasticache Reserved Cache Nodes.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReservedCacheNode.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|cache_node_count|`int64`|
|cache_node_type|`utf8`|
|duration|`int64`|
|fixed_price|`float64`|
|offering_type|`utf8`|
|product_description|`utf8`|
|recurring_charges|`json`|
|reservation_arn|`utf8`|
|reserved_cache_node_id|`utf8`|
|reserved_cache_nodes_offering_id|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|usage_price|`float64`|