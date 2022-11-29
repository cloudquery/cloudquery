# Table: aws_ec2_network_acls

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkAcl.html

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
|associations|JSON|
|entries|JSON|
|is_default|Bool|
|network_acl_id|String|
|owner_id|String|
|tags|JSON|
|vpc_id|String|