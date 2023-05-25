# Table: aws_backup_vaults

This table shows data for Backup Vaults.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BackupVaultListMember.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_backup_vaults:
  - [aws_backup_vault_recovery_points](aws_backup_vault_recovery_points)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|access_policy|json|
|notifications|json|
|tags|json|
|backup_vault_arn|utf8|
|backup_vault_name|utf8|
|creation_date|timestamp[us, tz=UTC]|
|creator_request_id|utf8|
|encryption_key_arn|utf8|
|lock_date|timestamp[us, tz=UTC]|
|locked|bool|
|max_retention_days|int64|
|min_retention_days|int64|
|number_of_recovery_points|int64|