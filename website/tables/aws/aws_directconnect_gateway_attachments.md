# Table: aws_directconnect_gateway_attachments

This table shows data for AWS Direct Connect Gateway Attachments.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAttachment.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_directconnect_gateways](aws_directconnect_gateways).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|gateway_arn|`utf8`|
|gateway_id|`utf8`|
|attachment_state|`utf8`|
|attachment_type|`utf8`|
|direct_connect_gateway_id|`utf8`|
|state_change_error|`utf8`|
|virtual_interface_id|`utf8`|
|virtual_interface_owner_account|`utf8`|
|virtual_interface_region|`utf8`|