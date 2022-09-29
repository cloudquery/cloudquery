# Table: azure_keyvault_keys


The primary key for this table is **kid**.

## Relations
This table depends on [`azure_keyvault_vaults`](azure_keyvault_vaults.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|keyvault_vault_id|String|
|kid (PK)|String|
|attributes|JSON|
|tags|JSON|
|managed|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|