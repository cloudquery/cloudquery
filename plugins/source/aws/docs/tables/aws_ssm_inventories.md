# Table: aws_ssm_inventories

This table shows data for AWS Systems Manager (SSM) Inventories.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryResultEntity.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|data|`json`|