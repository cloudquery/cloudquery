
# Table: aws_rds_instance_db_parameter_groups
The status of the DB parameter group
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_rds_instances table (FK)|
|instance_id|text|The AWS Region-unique, immutable identifier for the DB instance|
|db_parameter_group_name|text|The name of the DB parameter group.|
|parameter_apply_status|text|The status of parameter updates.|
