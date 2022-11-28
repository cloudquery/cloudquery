# Table: azure_frontdoor_doors

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor#FrontDoor

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|resource_state|String|
|provisioning_state|String|
|cname|String|
|frontdoor_id|String|
|rules_engines|JSON|
|friendly_name|String|
|routing_rules|JSON|
|load_balancing_settings|JSON|
|health_probe_settings|JSON|
|backend_pools|JSON|
|frontend_endpoints|JSON|
|backend_pools_settings|JSON|
|enabled_state|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|