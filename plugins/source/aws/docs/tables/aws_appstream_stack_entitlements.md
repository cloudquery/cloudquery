# Table: aws_appstream_stack_entitlements

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Entitlement.html

The composite primary key for this table is (**account_id**, **region**, **stack_name**, **name**).

## Relations
This table depends on [aws_appstream_stacks](aws_appstream_stacks.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|stack_name (PK)|String|
|name (PK)|String|
|app_visibility|String|
|attributes|JSON|
|created_time|Timestamp|
|description|String|
|last_modified_time|Timestamp|