# Table: aws_ec2_dhcp_options

This table shows data for Amazon Elastic Compute Cloud (EC2) DHCP Options.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DhcpOptions.html

The composite primary key for this table is (**account_id**, **region**, **dhcp_options_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|dhcp_configurations|`json`|
|dhcp_options_id (PK)|`utf8`|
|owner_id|`utf8`|