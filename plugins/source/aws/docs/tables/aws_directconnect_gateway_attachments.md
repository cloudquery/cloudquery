# Table: aws_directconnect_gateway_attachments


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_directconnect_gateways`](aws_directconnect_gateways.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|