# Table: azure_keyvault_keyvault_keys

This table shows data for Azure Key Vault Key Vault Keys.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault@v1.0.0#Key

The primary key for this table is **id**.

## Relations

This table depends on [azure_keyvault_keyvault](azure_keyvault_keyvault.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|tags|`json`|
|type|`utf8`|