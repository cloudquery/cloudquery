# Table: azure_container_registries

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry#Registry

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
|location|String|
|sku|JSON|
|identity|JSON|
|admin_user_enabled|Bool|
|anonymous_pull_enabled|Bool|
|data_endpoint_enabled|Bool|
|encryption|JSON|
|network_rule_bypass_options|String|
|network_rule_set|JSON|
|policies|JSON|
|public_network_access|String|
|zone_redundancy|String|
|creation_date|Timestamp|
|data_endpoint_host_names|StringArray|
|login_server|String|
|private_endpoint_connections|JSON|
|provisioning_state|String|
|status|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|