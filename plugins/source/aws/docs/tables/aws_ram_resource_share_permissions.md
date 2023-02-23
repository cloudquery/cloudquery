# Table: aws_ram_resource_share_permissions

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceSharePermissionSummary.html

The composite primary key for this table is (**account_id**, **region**, **arn**, **version**).

## Relations

This table depends on [aws_ram_resource_shares](aws_ram_resource_shares.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|permission|JSON|
|arn (PK)|String|
|creation_time|Timestamp|
|default_version|Bool|
|is_resource_type_default|Bool|
|last_updated_time|Timestamp|
|name|String|
|resource_type|String|
|status|String|
|version (PK)|String|