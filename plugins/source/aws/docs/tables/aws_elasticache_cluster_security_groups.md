
# Table: aws_elasticache_cluster_security_groups
Represents a single cache security group and its status.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_clusters table (FK)|
|security_group_id|text|The identifier of the cache security group.|
|status|text|The status of the cache security group membership|
