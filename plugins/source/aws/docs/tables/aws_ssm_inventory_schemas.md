# Table: aws_ssm_inventory_schemas

This table shows data for AWS Systems Manager (SSM) Inventory Schemas.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryItemSchema.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **type_name**, **version**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|type_name|`utf8`|
|version|`utf8`|
|attributes|`json`|
|display_name|`utf8`|