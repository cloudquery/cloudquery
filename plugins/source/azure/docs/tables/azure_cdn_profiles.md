# Table: azure_cdn_profiles

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn#Profile

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
|sku|JSON|
|resource_state|String|
|provisioning_state|String|
|frontdoor_id|String|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|
|system_data|JSON|