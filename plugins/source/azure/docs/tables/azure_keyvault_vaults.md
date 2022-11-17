# Table: azure_keyvault_vaults

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault#Vault

The primary key for this table is **id**.

## Relations

The following tables depend on azure_keyvault_vaults:
  - [azure_keyvault_keys](azure_keyvault_keys.md)
  - [azure_keyvault_secrets](azure_keyvault_secrets.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sku|JSON|
|tenant_id|String|
|access_policies|JSON|
|create_mode|String|
|enable_purge_protection|Bool|
|enable_rbac_authorization|Bool|
|enable_soft_delete|Bool|
|enabled_for_deployment|Bool|
|enabled_for_disk_encryption|Bool|
|enabled_for_template_deployment|Bool|
|network_acls|JSON|
|provisioning_state|String|
|public_network_access|String|
|soft_delete_retention_in_days|Int|
|vault_uri|String|
|hsm_pool_resource_id|String|
|private_endpoint_connections|JSON|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|