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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|id|`utf8`|
|amazon_side_asn|`int64`|
|direct_connect_gateway_id|`utf8`|
|direct_connect_gateway_name|`utf8`|
|direct_connect_gateway_state|`utf8`|
|owner_account|`utf8`|
|state_change_error|`utf8`|