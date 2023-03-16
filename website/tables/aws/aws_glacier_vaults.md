# Table: aws_glacier_vaults

This table shows data for Glacier Vaults.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vaults-get.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_glacier_vaults:
  - [aws_glacier_vault_access_policies](aws_glacier_vault_access_policies)
  - [aws_glacier_vault_lock_policies](aws_glacier_vault_lock_policies)
  - [aws_glacier_vault_notifications](aws_glacier_vault_notifications)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|creation_date|String|
|last_inventory_date|String|
|number_of_archives|Int|
|size_in_bytes|Int|
|vault_arn|String|
|vault_name|String|