# Table: azure_keyvault_keyvault_secrets

This table shows data for Azure Key Vault Key Vault Secrets.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault@v1.0.0#Secret

The primary key for this table is **id**.

## Relations

This table depends on [azure_keyvault_keyvault](azure_keyvault_keyvault).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|properties|json|
|id (PK)|utf8|
|location|utf8|
|name|utf8|
|tags|json|
|type|utf8|