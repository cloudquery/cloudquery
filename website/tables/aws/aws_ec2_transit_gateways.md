# Table: aws_ec2_transit_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) Transit Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGateway.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Relations

The following tables depend on aws_ec2_transit_gateways:
  - [aws_ec2_transit_gateway_attachments](aws_ec2_transit_gateway_attachments)
  - [aws_ec2_transit_gateway_multicast_domains](aws_ec2_transit_gateway_multicast_domains)
  - [aws_ec2_transit_gateway_peering_attachments](aws_ec2_transit_gateway_peering_attachments)
  - [aws_ec2_transit_gateway_route_tables](aws_ec2_transit_gateway_route_tables)
  - [aws_ec2_transit_gateway_vpc_attachments](aws_ec2_transit_gateway_vpc_attachments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|region (PK)|utf8|
|id|utf8|
|arn (PK)|utf8|
|tags|json|
|creation_time|timestamp[us, tz=UTC]|
|description|utf8|
|options|json|
|owner_id|utf8|
|state|utf8|
|transit_gateway_arn|utf8|
|transit_gateway_id|utf8|