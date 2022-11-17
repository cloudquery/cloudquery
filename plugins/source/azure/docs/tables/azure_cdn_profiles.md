# Table: azure_cdn_profiles

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn#Profile

The primary key for this table is **id**.

## Relations

The following tables depend on azure_cdn_profiles:
  - [azure_cdn_endpoints](azure_cdn_endpoints.md)
  - [azure_cdn_rule_sets](azure_cdn_rule_sets.md)
  - [azure_cdn_security_policies](azure_cdn_security_policies.md)

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
|origin_response_timeout_seconds|Int|
|front_door_id|String|
|provisioning_state|String|
|resource_state|String|
|tags|JSON|
|id (PK)|String|
|kind|String|
|name|String|
|system_data|JSON|
|type|String|