
# Table: aws_rds_instance_db_security_groups
This data type is used as a response element in the following actions:  * ModifyDBInstance  * RebootDBInstance  * RestoreDBInstanceFromDBSnapshot  * RestoreDBInstanceToPointInTime 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_rds_instances table (FK)|
|instance_id|text|The AWS Region-unique, immutable identifier for the DB instance|
|db_security_group_name|text|The name of the DB security group.|
|status|text|The status of the DB security group.|
