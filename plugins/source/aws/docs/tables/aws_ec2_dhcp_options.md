# Table: aws_ec2_dhcp_options

This table shows data for Amazon Elastic Compute Cloud (EC2) DHCP Options.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DhcpOptions.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **dhcp_options_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|dhcp_configurations|`json`|
|dhcp_options_id|`utf8`|
|owner_id|`utf8`|