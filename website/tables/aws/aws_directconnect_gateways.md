# Table: aws_directconnect_gateways

This table shows data for AWS Direct Connect Gateways.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGateway.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_directconnect_gateways:
  - [aws_directconnect_gateway_associations](aws_directconnect_gateway_associations)
  - [aws_directconnect_gateway_attachments](aws_directconnect_gateway_attachments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|arn (PK)|String|
|id|String|
|amazon_side_asn|Int|
|direct_connect_gateway_id|String|
|direct_connect_gateway_name|String|
|direct_connect_gateway_state|String|
|owner_account|String|
|state_change_error|String|