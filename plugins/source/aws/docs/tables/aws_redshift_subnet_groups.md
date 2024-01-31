# Table: aws_redshift_subnet_groups

This table shows data for Redshift Subnet Groups.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterSubnetGroup.html

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
|cluster_subnet_group_name|`utf8`|
|description|`utf8`|
|subnet_group_status|`utf8`|
|subnets|`json`|
|supported_cluster_ip_address_types|`list<item: utf8, nullable>`|
|vpc_id|`utf8`|