# Table: azure_keyvault_keyvault_keys

This table shows data for Azure Key Vault Key Vault Keys.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault@v1.0.0#Key

The primary key for this table is **id**.

## Relations

This table depends on [azure_keyvault_keyvault](azure_keyvault_keyvault).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|properties|JSON|
|id (PK)|String|
|location|String|
|name|String|
|tags|JSON|
|type|String|