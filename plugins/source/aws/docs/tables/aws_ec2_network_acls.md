# Table: aws_ec2_network_acls



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
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