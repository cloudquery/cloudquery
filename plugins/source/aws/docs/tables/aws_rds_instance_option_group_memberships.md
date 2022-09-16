
# Table: aws_rds_instance_option_group_memberships
Provides information on the option groups the DB instance is a member of. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_rds_instances table (FK)|
|option_group_name|text|The name of the option group that the instance belongs to.|
|status|text|The status of the DB instance's option group membership|
