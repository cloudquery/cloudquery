# Table: aws_rds_db_security_groups

This table shows data for Amazon Relational Database Service (RDS) DB Security Groups.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSecurityGroup.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|db_security_group_arn|`utf8`|
|db_security_group_description|`utf8`|
|db_security_group_name|`utf8`|
|ec2_security_groups|`json`|
|ip_ranges|`json`|
|owner_id|`utf8`|
|vpc_id|`utf8`|