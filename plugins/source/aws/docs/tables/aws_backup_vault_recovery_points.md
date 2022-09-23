# Table: aws_backup_vault_recovery_points


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_backup_vaults`](aws_backup_vaults.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|vault_arn|String|
|tags|JSON|
|backup_size_in_bytes|Int|
|backup_vault_arn|String|
|backup_vault_name|String|
|calculated_lifecycle|JSON|
|completion_date|Timestamp|
|created_by|JSON|
|creation_date|Timestamp|
|encryption_key_arn|String|
|iam_role_arn|String|
|is_encrypted|Bool|
|last_restore_time|Timestamp|
|lifecycle|JSON|
|recovery_point_arn|String|
|resource_arn|String|
|resource_type|String|
|source_backup_vault_arn|String|
|status|String|
|status_message|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|