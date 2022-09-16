
# Table: aws_backup_region_settings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|resource_type_management_preference|jsonb|Returns whether Backup fully manages the backups for a resource type|
|resource_type_opt_in_preference|jsonb|Returns a list of all services along with the opt-in preferences in the Region.|
