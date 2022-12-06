# Table: aws_ec2_transit_gateway_route_tables

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayRouteTable.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_ec2_transit_gateways](aws_ec2_transit_gateways.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|transit_gateway_arn|String|
|creation_time|Timestamp|
|default_association_route_table|Bool|
|default_propagation_route_table|Bool|
|state|String|
|tags|JSON|
|transit_gateway_id|String|
|transit_gateway_route_table_id|String|