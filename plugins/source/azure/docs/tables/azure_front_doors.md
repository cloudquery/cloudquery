# Table: azure_front_doors

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor#FrontDoor

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
|backend_pools|JSON|
|backend_pools_settings|JSON|
|enabled_state|String|
|friendly_name|String|
|frontend_endpoints|JSON|
|health_probe_settings|JSON|
|load_balancing_settings|JSON|
|routing_rules|JSON|
|cname|String|
|frontdoor_id|String|
|provisioning_state|String|
|resource_state|String|
|rules_engines|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|