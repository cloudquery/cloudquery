# Table: aws_elasticache_reserved_cache_nodes_offerings


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|cache_node_type|String|
|duration|Int|
|fixed_price|Float|
|offering_type|String|
|product_description|String|
|recurring_charges|JSON|
|reserved_cache_nodes_offering_id|String|
|usage_price|Float|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|