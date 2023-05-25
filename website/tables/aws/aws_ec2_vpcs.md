# Table: aws_ec2_vpcs

This table shows data for Amazon Elastic Compute Cloud (EC2) VPCs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Vpc.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|tags|json|
|cidr_block|utf8|
|cidr_block_association_set|json|
|dhcp_options_id|utf8|
|instance_tenancy|utf8|
|ipv6_cidr_block_association_set|json|
|is_default|bool|
|owner_id|utf8|
|state|utf8|
|vpc_id|utf8|