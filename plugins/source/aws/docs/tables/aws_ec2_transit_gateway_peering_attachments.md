# Table: aws_ec2_transit_gateway_peering_attachments


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
|accepter_tgw_info|JSON|
|creation_time|Timestamp|
|requester_tgw_info|JSON|
|state|String|
|status|JSON|
|transit_gateway_attachment_id|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|