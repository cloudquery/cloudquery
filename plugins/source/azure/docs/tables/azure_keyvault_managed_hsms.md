# Table: azure_keyvault_managed_hsms

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault#ManagedHsm

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|create_mode|String|
|enable_purge_protection|Bool|
|enable_soft_delete|Bool|
|initial_admin_object_ids|StringArray|
|network_acls|JSON|
|public_network_access|String|
|soft_delete_retention_in_days|Int|
|tenant_id|String|
|hsm_uri|String|
|private_endpoint_connections|JSON|
|provisioning_state|String|
|scheduled_purge_date|Timestamp|
|status_message|String|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|