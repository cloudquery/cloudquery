# Table: azure_keyvault_keyvault



The primary key for this table is **id**.

## Relations

The following tables depend on azure_keyvault_keyvault:
  - [azure_keyvault_keyvault_keys](azure_keyvault_keyvault_keys.md)
  - [azure_keyvault_keyvault_secrets](azure_keyvault_keyvault_secrets.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|location|String|
|name|String|
|tags|JSON|
|type|String|