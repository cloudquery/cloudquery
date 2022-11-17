# Table: azure_cdn_rule_sets

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn#RuleSet

The primary key for this table is **id**.

## Relations
This table depends on [azure_cdn_profiles](azure_cdn_profiles.md).

The following tables depend on azure_cdn_rule_sets:
  - [azure_cdn_rules](azure_cdn_rules.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|deployment_status|String|
|profile_name|String|
|provisioning_state|String|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|
|profile_id|String|