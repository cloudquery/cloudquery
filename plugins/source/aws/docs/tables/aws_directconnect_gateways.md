# Table: aws_directconnect_gateways

This table shows data for AWS Direct Connect Gateways.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGateway.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **arn**).
## Relations

The following tables depend on aws_directconnect_gateways:
  - [aws_directconnect_gateway_associations](aws_directconnect_gateway_associations.md)
  - [aws_directconnect_gateway_attachments](aws_directconnect_gateway_attachments.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|arn|`utf8`|
|id|`utf8`|
|amazon_side_asn|`int64`|
|direct_connect_gateway_id|`utf8`|
|direct_connect_gateway_name|`utf8`|
|direct_connect_gateway_state|`utf8`|
|owner_account|`utf8`|
|state_change_error|`utf8`|