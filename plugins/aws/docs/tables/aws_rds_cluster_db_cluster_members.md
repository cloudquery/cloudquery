
# Table: aws_rds_cluster_db_cluster_members
Contains information about an instance that is part of a DB cluster. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid|Unique ID of aws_rds_clusters table (FK)|
|db_cluster_parameter_group_status|text|Specifies the status of the DB cluster parameter group for this member of the DB cluster.|
|db_instance_identifier|text|Specifies the instance identifier for this member of the DB cluster.|
|is_cluster_writer|boolean|Value that is true if the cluster member is the primary instance for the DB cluster and false otherwise.|
|promotion_tier|integer|A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance|
