# Table: aws_shield_attacks



The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|id (PK)|String|
|attack_counters|JSON|
|attack_properties|JSON|
|end_time|Timestamp|
|mitigations|JSON|
|resource_arn|String|
|start_time|Timestamp|
|sub_resources|JSON|