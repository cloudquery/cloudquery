
# Table: aws_rds_instance_status_infos
Provides a list of status information for a DB instance. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid|Unique ID of aws_rds_instances table (FK)|
|message|text|Details of the error if there is an error for the instance|
|normal|boolean|Boolean value that is true if the instance is operating normally, or false if the instance is in an error state.|
|status|text|Status of the DB instance|
|status_type|text|This value is currently "read replication."|
