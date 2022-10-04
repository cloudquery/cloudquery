# Table: aws_ec2_subnets



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
|tags|JSON|
|vpc_id|String|