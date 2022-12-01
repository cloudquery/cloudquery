# Table: aws_ec2_network_interfaces

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInterface.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|association|JSON|
|attachment|JSON|
|availability_zone|String|
|deny_all_igw_traffic|Bool|
|description|String|
|groups|JSON|
|interface_type|String|
|ipv4_prefixes|JSON|
|ipv6_address|String|
|ipv6_addresses|JSON|
|ipv6_native|Bool|
|ipv6_prefixes|JSON|
|mac_address|String|
|network_interface_id|String|
|outpost_arn|String|
|owner_id|String|
|private_dns_name|String|
|private_ip_address|String|
|private_ip_addresses|JSON|
|requester_id|String|
|requester_managed|Bool|
|source_dest_check|Bool|
|status|String|
|subnet_id|String|
|vpc_id|String|