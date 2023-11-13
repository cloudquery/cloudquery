# Table: aws_ec2_eips

This table shows data for Amazon Elastic Compute Cloud (EC2) Eips.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html

The composite primary key for this table is (**account_id**, **region**, **allocation_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|allocation_id (PK)|`utf8`|
|association_id|`utf8`|
|carrier_ip|`utf8`|
|customer_owned_ip|`utf8`|
|customer_owned_ipv4_pool|`utf8`|
|domain|`utf8`|
|instance_id|`utf8`|
|network_border_group|`utf8`|
|network_interface_id|`utf8`|
|network_interface_owner_id|`utf8`|
|private_ip_address|`utf8`|
|public_ip|`utf8`|
|public_ipv4_pool|`utf8`|