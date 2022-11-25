# Table: aws_ec2_security_groups

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SecurityGroup.html

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
|description|String|
|group_id|String|
|group_name|String|
|ip_permissions|JSON|
|ip_permissions_egress|JSON|
|owner_id|String|
|tags|JSON|
|vpc_id|String|