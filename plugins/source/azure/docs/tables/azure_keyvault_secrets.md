# Table: azure_keyvault_secrets

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault#Secret

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
|attributes|JSON|
|content_type|String|
|value|String|
|secret_uri|String|
|secret_uri_with_version|String|
|id (PK)|String|
|location|String|
|name|String|
|tags|JSON|
|type|String|
|vault_id|String|