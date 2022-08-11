
# Table: aws_rds_instance_db_instance_automated_backups_replications
Automated backups of a DB instance replicated to another AWS Region
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_rds_instances table (FK)|
|instance_id|text|The AWS Region-unique, immutable identifier for the DB instance|
|db_instance_automated_backups_arn|text|The Amazon Resource Name (ARN) of the replicated automated backups.|
