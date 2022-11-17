# Table: azure_container_replications

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry#Replication

The primary key for this table is **id**.

## Relations
This table depends on [azure_container_registries](azure_container_registries.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|region_endpoint_enabled|Bool|
|zone_redundancy|String|
|provisioning_state|String|
|status|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|
|registry_id|String|