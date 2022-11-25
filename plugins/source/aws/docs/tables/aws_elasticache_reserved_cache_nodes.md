# Table: aws_elasticache_reserved_cache_nodes

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReservedCacheNode.html

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
|cache_node_count|Int|
|cache_node_type|String|
|duration|Int|
|fixed_price|Float|
|offering_type|String|
|product_description|String|
|recurring_charges|JSON|
|reserved_cache_node_id|String|
|reserved_cache_nodes_offering_id|String|
|start_time|Timestamp|
|state|String|
|usage_price|Float|