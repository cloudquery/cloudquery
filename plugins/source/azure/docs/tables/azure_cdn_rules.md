# Table: azure_cdn_rules

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn#Rule

The primary key for this table is **id**.

## Relations
This table depends on [azure_cdn_rule_sets](azure_cdn_rule_sets.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|actions|JSON|
|conditions|JSON|
|match_processing_behavior|String|
|order|Int|
|deployment_status|String|
|provisioning_state|String|
|rule_set_name|String|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|
|rule_set_id|String|