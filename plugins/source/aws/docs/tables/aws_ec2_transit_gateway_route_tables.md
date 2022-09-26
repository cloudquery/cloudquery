# Table: aws_ec2_transit_gateway_route_tables


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_ec2_transit_gateways`](aws_ec2_transit_gateways.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|transit_gateway_arn|String|
|tags|JSON|
|creation_time|Timestamp|
|default_association_route_table|Bool|
|default_propagation_route_table|Bool|
|state|String|
|transit_gateway_id|String|
|transit_gateway_route_table_id|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|