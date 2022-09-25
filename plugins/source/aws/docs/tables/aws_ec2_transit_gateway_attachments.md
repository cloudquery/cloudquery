# Table: aws_ec2_transit_gateway_attachments


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
|association|JSON|
|creation_time|Timestamp|
|resource_id|String|
|resource_owner_id|String|
|resource_type|String|
|state|String|
|transit_gateway_attachment_id|String|
|transit_gateway_id|String|
|transit_gateway_owner_id|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|