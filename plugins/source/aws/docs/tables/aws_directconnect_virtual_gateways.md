# Table: aws_directconnect_virtual_gateways

This table shows data for AWS Direct Connect Virtual Gateways.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualGateway.html

The composite primary key for this table is (**request_account_id**, **request_region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|id (PK)|`utf8`|
|virtual_gateway_id|`utf8`|
|virtual_gateway_state|`utf8`|