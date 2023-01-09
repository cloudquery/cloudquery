# Table: aws_directconnect_virtual_gateways

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualGateway.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|id (PK)|String|
|virtual_gateway_id|String|
|virtual_gateway_state|String|