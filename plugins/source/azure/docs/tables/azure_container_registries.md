# Table: azure_container_registries

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry#Registry

The primary key for this table is **id**.

## Relations

The following tables depend on azure_container_registries:
  - [azure_container_replications](azure_container_replications.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sku|JSON|
|login_server|String|
|creation_date|Timestamp|
|provisioning_state|String|
|status|JSON|
|admin_user_enabled|Bool|
|storage_account|JSON|
|network_rule_set|JSON|
|policies|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|