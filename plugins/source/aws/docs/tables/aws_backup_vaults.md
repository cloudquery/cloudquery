
# Table: aws_backup_vaults
Contains metadata about a backup vault.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.|
|name|text|The name of a logical container where backups are stored.|
|creation_date|timestamp without time zone|The date and time a resource backup is created, in Unix format and Coordinated Universal Time (UTC).|
|creator_request_id|text|A unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.|
|encryption_key_arn|text|A server-side encryption key you can specify to encrypt your backups from services that support full Backup management.|
|lock_date|timestamp without time zone|The date and time when Backup Vault Lock configuration becomes immutable, meaning it cannot be changed or deleted.|
|locked|boolean|A Boolean value that indicates whether Backup Vault Lock applies to the selected backup vault.|
|max_retention_days|bigint|The Backup Vault Lock setting that specifies the maximum retention period that the vault retains its recovery points.|
|min_retention_days|bigint|The Backup Vault Lock setting that specifies the minimum retention period that the vault retains its recovery points.|
|number_of_recovery_points|bigint|The number of recovery points that are stored in a backup vault.|
|access_policy|jsonb|The backup vault access policy document in JSON format.|
|notification_events|text[]|An array of events that indicate the status of jobs to back up resources to the backup vault.|
|notification_sns_topic_arn|text|An ARN that uniquely identifies an Amazon Simple Notification Service topic.|
|tags|jsonb|Resource tags|
