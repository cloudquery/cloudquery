# Table: aws_ram_resource_shares

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShare.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_ram_resource_shares:
  - [aws_ram_resource_share_permissions](aws_ram_resource_share_permissions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|allow_external_principals|Bool|
|creation_time|Timestamp|
|feature_set|String|
|last_updated_time|Timestamp|
|name|String|
|owning_account_id|String|
|arn (PK)|String|
|status|String|
|status_message|String|
|tags|JSON|