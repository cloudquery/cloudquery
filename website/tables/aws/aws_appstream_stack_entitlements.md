# Table: aws_appstream_stack_entitlements

This table shows data for Amazon AppStream Stack Entitlements.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Entitlement.html

The composite primary key for this table is (**account_id**, **region**, **stack_name**, **name**).

## Relations

This table depends on [aws_appstream_stacks](aws_appstream_stacks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|stack_name (PK)|`utf8`|
|name (PK)|`utf8`|
|app_visibility|`utf8`|
|attributes|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|