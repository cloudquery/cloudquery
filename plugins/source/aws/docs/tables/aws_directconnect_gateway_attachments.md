# Table: aws_directconnect_gateway_attachments

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAttachment.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_directconnect_gateways](aws_directconnect_gateways.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|gateway_arn|String|
|gateway_id|String|
|attachment_state|String|
|attachment_type|String|
|direct_connect_gateway_id|String|
|state_change_error|String|
|virtual_interface_id|String|
|virtual_interface_owner_account|String|
|virtual_interface_region|String|