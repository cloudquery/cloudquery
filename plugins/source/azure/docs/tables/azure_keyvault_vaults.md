# Table: azure_keyvault_vaults

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault#Vault

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
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|properties_tenant_id|UUID|
|properties_sku|JSON|
|properties_access_policies|JSON|
|properties_vault_uri|String|
|properties_enabled_for_deployment|Bool|
|properties_enabled_for_disk_encryption|Bool|
|properties_enabled_for_template_deployment|Bool|
|properties_enable_soft_delete|Bool|
|properties_soft_delete_retention_in_days|Int|
|properties_enable_rbac_authorization|Bool|
|properties_create_mode|String|
|properties_enable_purge_protection|Bool|
|properties_network_acls|JSON|
|properties_private_endpoint_connections|JSON|