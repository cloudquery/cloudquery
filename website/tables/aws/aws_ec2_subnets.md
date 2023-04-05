# Table: aws_ec2_subnets

This table shows data for Amazon Elastic Compute Cloud (EC2) Subnets.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Subnet.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|request_account_id (PK)|String|
|request_region (PK)|String|
|arn (PK)|String|
|tags|JSON|
|assign_ipv6_address_on_creation|Bool|
|availability_zone|String|
|availability_zone_id|String|
|available_ip_address_count|Int|
|cidr_block|String|
|customer_owned_ipv4_pool|String|
|default_for_az|Bool|
|enable_dns64|Bool|
|enable_lni_at_device_index|Int|
|ipv6_cidr_block_association_set|JSON|
|ipv6_native|Bool|
|map_customer_owned_ip_on_launch|Bool|
|map_public_ip_on_launch|Bool|
|outpost_arn|String|
|owner_id|String|
|private_dns_name_options_on_launch|JSON|
|state|String|
|subnet_arn|String|
|subnet_id|String|
|vpc_id|String|