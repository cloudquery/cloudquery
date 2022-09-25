# Table: aws_ec2_security_groups


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|