# Table: aws_ec2_transit_gateway_peering_attachments

This table shows data for Amazon Elastic Compute Cloud (EC2) Transit Gateway Peering Attachments.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayPeeringAttachment.html

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
|accepter_tgw_info|`json`|
|accepter_transit_gateway_attachment_id|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|options|`json`|
|requester_tgw_info|`json`|
|state|`utf8`|
|status|`json`|
|transit_gateway_attachment_id|`utf8`|