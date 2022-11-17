# Table: azure_keyvault_keys

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault#Key

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
|curve_name|String|
|key_ops|StringArray|
|key_size|Int|
|kty|String|
|key_uri|String|
|key_uri_with_version|String|
|id (PK)|String|
|location|String|
|name|String|
|tags|JSON|
|type|String|
|vault_id|String|