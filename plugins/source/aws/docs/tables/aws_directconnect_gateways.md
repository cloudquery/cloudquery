# Table: aws_directconnect_gateways

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGateway.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_directconnect_gateways:
  - [aws_directconnect_gateway_associations](aws_directconnect_gateway_associations.md)
  - [aws_directconnect_gateway_attachments](aws_directconnect_gateway_attachments.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|id|String|
|amazon_side_asn|Int|
|direct_connect_gateway_id|String|
|direct_connect_gateway_name|String|
|direct_connect_gateway_state|String|
|owner_account|String|
|state_change_error|String|