# Table: azure_keyvault_managed_hsms

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault#ManagedHsm

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties_tenant_id|UUID|
|properties_initial_admin_object_ids|StringArray|
|properties_hsm_uri|String|
|properties_enable_soft_delete|Bool|
|properties_soft_delete_retention_in_days|Int|
|properties_enable_purge_protection|Bool|
|properties_create_mode|String|
|properties_status_message|String|
|properties_provisioning_state|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|sku|JSON|
|tags|JSON|