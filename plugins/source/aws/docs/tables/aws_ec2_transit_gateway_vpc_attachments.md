# Table: aws_ec2_transit_gateway_vpc_attachments

This table shows data for Amazon Elastic Compute Cloud (EC2) Transit Gateway VPC Attachments.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayVpcAttachment.html

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
|creation_time|`timestamp[us, tz=UTC]`|
|options|`json`|
|state|`utf8`|
|subnet_ids|`list<item: utf8, nullable>`|
|tags|`json`|
|transit_gateway_attachment_id|`utf8`|
|transit_gateway_id|`utf8`|
|vpc_id|`utf8`|
|vpc_owner_id|`utf8`|