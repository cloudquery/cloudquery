# Table: aws_ec2_instance_statuses


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|availability_zone|String|
|events|JSON|
|instance_id|String|
|instance_state|JSON|
|instance_status|JSON|
|outpost_arn|String|
|system_status|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|