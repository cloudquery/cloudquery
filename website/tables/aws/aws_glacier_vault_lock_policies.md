# Table: aws_glacier_vault_lock_policies

This table shows data for Glacier Vault Lock Policies.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultLock.html

The primary key for this table is **vault_arn**.

## Relations

This table depends on [aws_glacier_vaults](aws_glacier_vaults).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|vault_arn (PK)|`utf8`|
|policy|`json`|
|creation_date|`utf8`|
|expiration_date|`utf8`|
|state|`utf8`|