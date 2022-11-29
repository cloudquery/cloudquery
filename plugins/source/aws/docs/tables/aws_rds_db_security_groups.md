# Table: aws_rds_db_security_groups

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSecurityGroup.html

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
|tags|JSON|
|db_security_group_description|String|
|db_security_group_name|String|
|ec2_security_groups|JSON|
|ip_ranges|JSON|
|owner_id|String|
|vpc_id|String|