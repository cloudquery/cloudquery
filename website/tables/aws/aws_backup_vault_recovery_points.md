# Table: aws_backup_vault_recovery_points

This table shows data for Backup Vault Recovery Points.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_RecoveryPointByBackupVault.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_backup_vaults](aws_backup_vaults).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|vault_arn|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|backup_size_in_bytes|`int64`|
|backup_vault_arn|`utf8`|
|backup_vault_name|`utf8`|
|calculated_lifecycle|`json`|
|completion_date|`timestamp[us, tz=UTC]`|
|composite_member_identifier|`utf8`|
|created_by|`json`|
|creation_date|`timestamp[us, tz=UTC]`|
|encryption_key_arn|`utf8`|
|iam_role_arn|`utf8`|
|is_encrypted|`bool`|
|is_parent|`bool`|
|last_restore_time|`timestamp[us, tz=UTC]`|
|lifecycle|`json`|
|parent_recovery_point_arn|`utf8`|
|recovery_point_arn|`utf8`|
|resource_arn|`utf8`|
|resource_name|`utf8`|
|resource_type|`utf8`|
|source_backup_vault_arn|`utf8`|
|status|`utf8`|
|status_message|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Vaults with no recovery points

```sql
WITH
  "point" AS (SELECT DISTINCT vault_arn FROM aws_backup_vault_recovery_points)
SELECT
  'Vaults with no recovery points' AS title,
  vault.account_id,
  vault.arn AS resource_id,
  'fail' AS status
FROM
  aws_backup_vaults AS vault LEFT JOIN "point" ON "point".vault_arn = vault.arn
WHERE
  "point".vault_arn IS NULL;
```


