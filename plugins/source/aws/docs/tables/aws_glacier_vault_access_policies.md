# Table: aws_glacier_vault_access_policies

This table shows data for Glacier Vault Access Policies.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultAccessPolicy.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **vault_arn**.
## Relations

This table depends on [aws_glacier_vaults](aws_glacier_vaults.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|vault_arn|`utf8`|
|policy|`json`|