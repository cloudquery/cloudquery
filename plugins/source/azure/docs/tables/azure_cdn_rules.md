# Table: azure_cdn_rules



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
|cdn_rule_set_id|String|
|order|Int|
|conditions|JSON|
|actions|JSON|
|match_processing_behavior|String|
|provisioning_state|String|
|deployment_status|String|
|id (PK)|String|
|name|String|
|type|String|
|system_data|JSON|