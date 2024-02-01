# Table: aws_docdb_subnet_groups

This table shows data for Amazon DocumentDB Subnet Groups.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBSubnetGroup.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|db_subnet_group_arn|`utf8`|
|db_subnet_group_description|`utf8`|
|db_subnet_group_name|`utf8`|
|subnet_group_status|`utf8`|
|subnets|`json`|
|vpc_id|`utf8`|