# Table: aws_ssm_inventories

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryResultEntity.html

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
|data|JSON|