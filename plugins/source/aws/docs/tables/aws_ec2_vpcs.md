# Table: aws_ec2_vpcs



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
|cidr_block|String|
|cidr_block_association_set|JSON|
|dhcp_options_id|String|
|instance_tenancy|String|
|ipv6_cidr_block_association_set|JSON|
|is_default|Bool|
|owner_id|String|
|state|String|
|tags|JSON|
|vpc_id|String|