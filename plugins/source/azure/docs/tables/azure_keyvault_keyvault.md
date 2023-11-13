# Table: azure_keyvault_keyvault

This table shows data for Azure Key Vault Key Vault.

https://learn.microsoft.com/en-us/rest/api/keyvault/keyvault/vaults/get?tabs=HTTP#vault

The primary key for this table is **id**.

## Relations

The following tables depend on azure_keyvault_keyvault:
  - [azure_keyvault_keyvault_keys](azure_keyvault_keyvault_keys.md)
  - [azure_keyvault_keyvault_secrets](azure_keyvault_keyvault_secrets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|location|`utf8`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|