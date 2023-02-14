# Table: aws_ram_resources

https://docs.aws.amazon.com/ram/latest/APIReference/API_Resource.html

The composite primary key for this table is (**account_id**, **region**, **arn**, **resource_share_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|arn (PK)|String|
|creation_time|Timestamp|
|last_updated_time|Timestamp|
|resource_group_arn|String|
|resource_region_scope|String|
|resource_share_arn (PK)|String|
|status|String|
|status_message|String|
|type|String|