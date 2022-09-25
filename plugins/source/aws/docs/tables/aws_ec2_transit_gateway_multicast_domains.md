# Table: aws_ec2_transit_gateway_multicast_domains


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
|options|JSON|
|owner_id|String|
|state|String|
|transit_gateway_id|String|
|transit_gateway_multicast_domain_arn|String|
|transit_gateway_multicast_domain_id|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|