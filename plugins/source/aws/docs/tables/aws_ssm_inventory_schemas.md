# Table: aws_ssm_inventory_schemas

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryItemSchema.html

The composite primary key for this table is (**account_id**, **region**, **type_name**, **version**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|type_name (PK)|String|
|version (PK)|String|
|attributes|JSON|
|display_name|String|