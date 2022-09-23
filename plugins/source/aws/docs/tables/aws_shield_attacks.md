# Table: aws_shield_attacks


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|id (PK)|String|
|attack_counters|JSON|
|attack_properties|JSON|
|end_time|Timestamp|
|mitigations|JSON|
|resource_arn|String|
|start_time|Timestamp|
|sub_resources|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|