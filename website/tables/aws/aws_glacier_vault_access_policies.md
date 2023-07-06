# Table: aws_glacier_vault_access_policies

This table shows data for Glacier Vault Access Policies.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultAccessPolicy.html

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