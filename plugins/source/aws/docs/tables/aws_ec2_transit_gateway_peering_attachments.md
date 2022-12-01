# Table: aws_ec2_transit_gateway_peering_attachments

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayPeeringAttachment.html

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
|accepter_tgw_info|JSON|
|accepter_transit_gateway_attachment_id|String|
|creation_time|Timestamp|
|options|JSON|
|requester_tgw_info|JSON|
|state|String|
|status|JSON|
|tags|JSON|
|transit_gateway_attachment_id|String|