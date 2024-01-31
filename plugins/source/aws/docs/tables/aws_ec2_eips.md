# Table: aws_ec2_eips

This table shows data for Amazon Elastic Compute Cloud (EC2) Eips.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **allocation_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|allocation_id|`utf8`|
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