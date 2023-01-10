# Table: aws_backup_vault_recovery_points

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_RecoveryPointByBackupVault.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_backup_vaults](aws_backup_vaults.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|vault_arn|String|
|arn (PK)|String|
|tags|JSON|
|backup_size_in_bytes|Int|
|backup_vault_arn|String|
|backup_vault_name|String|
|calculated_lifecycle|JSON|
|completion_date|Timestamp|
|composite_member_identifier|String|
|created_by|JSON|
|creation_date|Timestamp|
|encryption_key_arn|String|
|iam_role_arn|String|
|is_encrypted|Bool|
|is_parent|Bool|
|last_restore_time|Timestamp|
|lifecycle|JSON|
|parent_recovery_point_arn|String|
|recovery_point_arn|String|
|resource_arn|String|
|resource_type|String|
|source_backup_vault_arn|String|
|status|String|
|status_message|String|