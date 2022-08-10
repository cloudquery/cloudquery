
# Table: aws_lightsail_database_pending_maintenance_actions
Describes a pending database maintenance action
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_databases table (FK)|
|action|text|The type of pending database maintenance action|
|current_apply_date|timestamp without time zone|The effective date of the pending database maintenance action|
|description|text|Additional detail about the pending database maintenance action|
