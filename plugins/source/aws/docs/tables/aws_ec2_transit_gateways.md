# Table: aws_ec2_transit_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) Transit Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGateway.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).
## Relations

The following tables depend on aws_ec2_transit_gateways:
  - [aws_ec2_transit_gateway_attachments](aws_ec2_transit_gateway_attachments.md)
  - [aws_ec2_transit_gateway_multicast_domains](aws_ec2_transit_gateway_multicast_domains.md)
  - [aws_ec2_transit_gateway_peering_attachments](aws_ec2_transit_gateway_peering_attachments.md)
  - [aws_ec2_transit_gateway_route_tables](aws_ec2_transit_gateway_route_tables.md)
  - [aws_ec2_transit_gateway_vpc_attachments](aws_ec2_transit_gateway_vpc_attachments.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|options|`json`|
|owner_id|`utf8`|
|state|`utf8`|
|tags|`json`|
|transit_gateway_arn|`utf8`|
|transit_gateway_id|`utf8`|