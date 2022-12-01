# Table: aws_ec2_instance_statuses

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceStatus.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
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