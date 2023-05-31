# Table: aws_ec2_security_groups

This table shows data for Amazon Elastic Compute Cloud (EC2) Security Groups.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SecurityGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|description|`utf8`|
|group_id|`utf8`|
|group_name|`utf8`|
|ip_permissions|`json`|
|ip_permissions_egress|`json`|
|owner_id|`utf8`|
|vpc_id|`utf8`|