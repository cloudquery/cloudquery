# Table: aws_ec2_transit_gateway_peering_attachments

This table shows data for Amazon Elastic Compute Cloud (EC2) Transit Gateway Peering Attachments.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayPeeringAttachment.html

The composite primary key for this table is (**account_id**, **region**, **transit_gateway_arn**, **id**).

## Relations

This table depends on [aws_ec2_transit_gateways](aws_ec2_transit_gateways.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|transit_gateway_arn (PK)|`utf8`|
|id (PK)|`utf8`|
|accepter_tgw_info|`json`|
|accepter_transit_gateway_attachment_id|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|options|`json`|
|requester_tgw_info|`json`|
|state|`utf8`|
|status|`json`|
|tags|`json`|
|transit_gateway_attachment_id|`utf8`|