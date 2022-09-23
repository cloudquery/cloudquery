# Table: azure_keyvault_secrets


The primary key for this table is **id**.

## Relations
This table depends on [`azure_keyvault_vaults`](azure_keyvault_vaults.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|keyvault_vault_id|UUID|
|id (PK)|String|
|attributes|JSON|
|tags|JSON|
|content_type|String|
|managed|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|