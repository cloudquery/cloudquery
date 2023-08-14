# Table: aws_ec2_instance_statuses

This table shows data for Amazon Elastic Compute Cloud (EC2) Instance Statuses.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceStatus.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|availability_zone|`utf8`|
|events|`json`|
|instance_id|`utf8`|
|instance_state|`json`|
|instance_status|`json`|
|outpost_arn|`utf8`|
|system_status|`json`|