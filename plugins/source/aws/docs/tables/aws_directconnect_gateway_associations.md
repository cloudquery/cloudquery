# Table: aws_directconnect_gateway_associations

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAssociation.html

The primary key for this table is **account_id**.

## Relations

This table depends on [aws_directconnect_gateways](aws_directconnect_gateways.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|gateway_arn|String|
|gateway_id|String|
|allowed_prefixes_to_direct_connect_gateway|JSON|
|associated_gateway|JSON|
|association_id|String|
|association_state|String|
|direct_connect_gateway_id|String|
|direct_connect_gateway_owner_account|String|
|state_change_error|String|
|virtual_gateway_id|String|
|virtual_gateway_owner_account|String|
|virtual_gateway_region|String|