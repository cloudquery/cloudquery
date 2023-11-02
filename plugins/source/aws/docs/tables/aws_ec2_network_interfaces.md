# Table: aws_ec2_network_interfaces

This table shows data for Amazon Elastic Compute Cloud (EC2) Network Interfaces.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInterface.html

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
|association|`json`|
|attachment|`json`|
|availability_zone|`utf8`|
|deny_all_igw_traffic|`bool`|
|description|`utf8`|
|groups|`json`|
|interface_type|`utf8`|
|ipv4_prefixes|`json`|
|ipv6_address|`utf8`|
|ipv6_addresses|`json`|
|ipv6_native|`bool`|
|ipv6_prefixes|`json`|
|mac_address|`utf8`|
|network_interface_id|`utf8`|
|outpost_arn|`utf8`|
|owner_id|`utf8`|
|private_dns_name|`utf8`|
|private_ip_address|`utf8`|
|private_ip_addresses|`json`|
|requester_id|`utf8`|
|requester_managed|`bool`|
|source_dest_check|`bool`|
|status|`utf8`|
|subnet_id|`utf8`|
|vpc_id|`utf8`|