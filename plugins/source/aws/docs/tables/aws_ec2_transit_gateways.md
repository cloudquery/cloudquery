# Table: aws_ec2_transit_gateways

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGateway.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

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