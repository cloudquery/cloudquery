# Table: aws_ec2_transit_gateways

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGateway.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ec2_transit_gateways:
  - [aws_ec2_transit_gateway_attachments](aws_ec2_transit_gateway_attachments.md)
  - [aws_ec2_transit_gateway_route_tables](aws_ec2_transit_gateway_route_tables.md)
  - [aws_ec2_transit_gateway_vpc_attachments](aws_ec2_transit_gateway_vpc_attachments.md)
  - [aws_ec2_transit_gateway_peering_attachments](aws_ec2_transit_gateway_peering_attachments.md)
  - [aws_ec2_transit_gateway_multicast_domains](aws_ec2_transit_gateway_multicast_domains.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|id|String|
|arn (PK)|String|
|creation_time|Timestamp|
|description|String|
|options|JSON|
|owner_id|String|
|state|String|
|tags|JSON|