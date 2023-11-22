# Table: aws_directconnect_gateway_associations

This table shows data for AWS Direct Connect Gateway Associations.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAssociation.html

The composite primary key for this table is (**gateway_arn**, **association_id**).

## Relations

This table depends on [aws_directconnect_gateways](aws_directconnect_gateways.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|gateway_arn (PK)|`utf8`|
|gateway_id|`utf8`|
|allowed_prefixes_to_direct_connect_gateway|`json`|
|associated_gateway|`json`|
|association_id (PK)|`utf8`|
|association_state|`utf8`|
|direct_connect_gateway_id|`utf8`|
|direct_connect_gateway_owner_account|`utf8`|
|state_change_error|`utf8`|
|virtual_gateway_id|`utf8`|
|virtual_gateway_owner_account|`utf8`|
|virtual_gateway_region|`utf8`|