# Table: aws_ec2_transit_gateway_route_tables

This table shows data for Amazon Elastic Compute Cloud (EC2) Transit Gateway Route Tables.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayRouteTable.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ec2_transit_gateways](aws_ec2_transit_gateways).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|transit_gateway_arn|`utf8`|
|tags|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|default_association_route_table|`bool`|
|default_propagation_route_table|`bool`|
|state|`utf8`|
|transit_gateway_id|`utf8`|
|transit_gateway_route_table_id|`utf8`|