# Table: aws_directconnect_virtual_gateways

This table shows data for AWS Direct Connect Virtual Gateways.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualGateway.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **request_region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|id|`utf8`|
|virtual_gateway_id|`utf8`|
|virtual_gateway_state|`utf8`|