# Table: azure_cdn_rule_sets



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
|cdn_profile_id|String|
|provisioning_state|String|
|deployment_status|String|
|id (PK)|String|
|name|String|
|type|String|
|system_data|JSON|