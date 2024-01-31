# Table: aws_ec2_transit_gateway_multicast_domains

This table shows data for Amazon Elastic Compute Cloud (EC2) Transit Gateway Multicast Domains.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayMulticastDomain.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **transit_gateway_arn**, **arn**).
## Relations

This table depends on [aws_ec2_transit_gateways](aws_ec2_transit_gateways.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|transit_gateway_arn|`utf8`|
|arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|options|`json`|
|owner_id|`utf8`|
|state|`utf8`|
|tags|`json`|
|transit_gateway_id|`utf8`|
|transit_gateway_multicast_domain_arn|`utf8`|
|transit_gateway_multicast_domain_id|`utf8`|