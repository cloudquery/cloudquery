# Table: aws_ec2_transit_gateway_vpc_attachments

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayVpcAttachment.html

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
|options|JSON|
|state|String|
|subnet_ids|StringArray|
|tags|JSON|
|transit_gateway_attachment_id|String|
|transit_gateway_id|String|
|vpc_id|String|
|vpc_owner_id|String|