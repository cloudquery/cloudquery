# Table: aws_redshift_subnet_groups

This table shows data for Redshift Subnet Groups.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterSubnetGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|cluster_subnet_group_name|`utf8`|
|description|`utf8`|
|subnet_group_status|`utf8`|
|subnets|`json`|
|vpc_id|`utf8`|