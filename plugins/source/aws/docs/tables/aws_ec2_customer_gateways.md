# Table: aws_ec2_customer_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) Customer Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CustomerGateway.html

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
|bgp_asn|`utf8`|
|certificate_arn|`utf8`|
|customer_gateway_id|`utf8`|
|device_name|`utf8`|
|ip_address|`utf8`|
|state|`utf8`|
|type|`utf8`|