# Table: azure_cdn_rules


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|cdn_rule_set_id|UUID|
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|