# Table: azure_keyvault_secrets



The primary key for this table is **id**.

## Relations
This table depends on [azure_keyvault_vaults](azure_keyvault_vaults.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|keyvault_vault_id|String|
|id (PK)|String|
|attributes|JSON|
|tags|JSON|
|content_type|String|
|managed|Bool|