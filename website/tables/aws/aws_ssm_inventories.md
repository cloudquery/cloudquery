# Table: aws_ssm_inventories

This table shows data for AWS Systems Manager (SSM) Inventories.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryResultEntity.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|id (PK)|`utf8`|
|data|`json`|