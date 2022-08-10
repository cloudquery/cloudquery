
# Table: aws_backup_plan_rules
Specifies a scheduled task used to back up a selection of resources.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|plan_cq_id|uuid|Unique CloudQuery ID of aws_backup_plan table (FK)|
|name|text|A display name for a backup rule.|
|target_backup_vault_name|text|The name of a logical container where backups are stored.|
|completion_window_minutes|bigint|A value in minutes after a backup job is successfully started before it must be completed or it will be canceled by Backup.|
|copy_actions|jsonb|The details of the copy operation.|
|enable_continuous_backup|boolean|Specifies whether Backup creates continuous backups.|
|delete_after_days|bigint|Specifies the number of days after creation that a recovery point is deleted.|
|move_to_cold_storage_after_days|bigint|Specifies the number of days after creation that a recovery point is moved to cold storage.|
|recovery_point_tags|jsonb|An array of key-value pair strings that are assigned to resources that are associated with this rule when restored from backup.|
|id|text|Uniquely identifies a rule that is used to schedule the backup of a selection of resources.|
|schedule_expression|text|A cron expression in UTC specifying when Backup initiates a backup job.|
|start_window_minutes|bigint|A value in minutes after a backup is scheduled before a job will be canceled if it doesn't start successfully.|
