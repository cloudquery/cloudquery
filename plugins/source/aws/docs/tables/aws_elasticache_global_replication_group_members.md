
# Table: aws_elasticache_global_replication_group_members
A member of a Global datastore
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|global_replication_group_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_global_replication_groups table (FK)|
|automatic_failover|text|Indicates whether automatic failover is enabled for the replication group.|
|replication_group_id|text|The replication group id of the Global datastore member.|
|replication_group_region|text|The Amazon region of the Global datastore member.|
|role|text|Indicates the role of the replication group, primary or secondary.|
|status|text|The status of the membership of the replication group.|
