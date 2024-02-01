# Table: aws_ec2_vpn_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) VPN Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpnGateway.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|amazon_side_asn|`int64`|
|availability_zone|`utf8`|
|state|`utf8`|
|type|`utf8`|
|vpc_attachments|`json`|
|vpn_gateway_id|`utf8`|