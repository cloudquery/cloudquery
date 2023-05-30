# Table: aws_ec2_subnets

This table shows data for Amazon Elastic Compute Cloud (EC2) Subnets.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Subnet.html
The 'request_account_id' and 'request_region' columns are added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|assign_ipv6_address_on_creation|`bool`|
|availability_zone|`utf8`|
|availability_zone_id|`utf8`|
|available_ip_address_count|`int64`|
|cidr_block|`utf8`|
|customer_owned_ipv4_pool|`utf8`|
|default_for_az|`bool`|
|enable_dns64|`bool`|
|enable_lni_at_device_index|`int64`|
|ipv6_cidr_block_association_set|`json`|
|ipv6_native|`bool`|
|map_customer_owned_ip_on_launch|`bool`|
|map_public_ip_on_launch|`bool`|
|outpost_arn|`utf8`|
|owner_id|`utf8`|
|private_dns_name_options_on_launch|`json`|
|state|`utf8`|
|subnet_arn|`utf8`|
|subnet_id|`utf8`|
|vpc_id|`utf8`|