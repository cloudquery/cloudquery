
# Table: aws_elasticache_cluster_cache_security_groups
Represents a cluster's status within a particular cache security group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_clusters table (FK)|
|name|text|The name of the cache security group.|
|status|text|The membership status in the cache security group|
