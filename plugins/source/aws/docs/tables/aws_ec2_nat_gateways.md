# Table: aws_ec2_nat_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) NAT Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NatGateway.html

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
|connectivity_type|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|delete_time|`timestamp[us, tz=UTC]`|
|failure_code|`utf8`|
|failure_message|`utf8`|
|nat_gateway_addresses|`json`|
|nat_gateway_id|`utf8`|
|provisioned_bandwidth|`json`|
|state|`utf8`|
|subnet_id|`utf8`|
|vpc_id|`utf8`|