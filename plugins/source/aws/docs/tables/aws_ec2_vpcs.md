# Table: aws_ec2_vpcs

This table shows data for Amazon Elastic Compute Cloud (EC2) VPCs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Vpc.html

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
|cidr_block|`utf8`|
|cidr_block_association_set|`json`|
|dhcp_options_id|`utf8`|
|instance_tenancy|`utf8`|
|ipv6_cidr_block_association_set|`json`|
|is_default|`bool`|
|owner_id|`utf8`|
|state|`utf8`|
|vpc_id|`utf8`|