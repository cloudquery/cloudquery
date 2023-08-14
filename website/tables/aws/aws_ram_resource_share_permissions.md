# Table: aws_ram_resource_share_permissions

This table shows data for RAM Resource Share Permissions.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceSharePermissionSummary.html

The composite primary key for this table is (**account_id**, **region**, **resource_share_arn**, **arn**, **version**).

## Relations

This table depends on [aws_ram_resource_shares](aws_ram_resource_shares).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|resource_share_arn (PK)|`utf8`|
|permission|`json`|
|tags|`json`|
|arn (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|default_version|`bool`|
|feature_set|`utf8`|
|is_resource_type_default|`bool`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|permission_type|`utf8`|
|resource_type|`utf8`|
|status|`utf8`|
|version (PK)|`utf8`|