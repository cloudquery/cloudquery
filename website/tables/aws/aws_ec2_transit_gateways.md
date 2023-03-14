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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|id|String|
|arn (PK)|String|
|tags|JSON|
|creation_time|Timestamp|
|description|String|
|options|JSON|
|owner_id|String|
|state|String|
|transit_gateway_arn|String|
|transit_gateway_id|String|