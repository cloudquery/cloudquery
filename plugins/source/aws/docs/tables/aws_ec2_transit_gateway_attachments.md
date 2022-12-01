# Table: aws_ec2_transit_gateway_attachments

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayAttachment.html

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
|association|JSON|
|creation_time|Timestamp|
|resource_id|String|
|resource_owner_id|String|
|resource_type|String|
|state|String|
|tags|JSON|
|transit_gateway_attachment_id|String|
|transit_gateway_id|String|
|transit_gateway_owner_id|String|