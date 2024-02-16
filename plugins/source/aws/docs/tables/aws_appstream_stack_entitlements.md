# Table: aws_appstream_stack_entitlements

This table shows data for Amazon AppStream Stack Entitlements.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Entitlement.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **stack_name**, **name**).
## Relations

This table depends on [aws_appstream_stacks](aws_appstream_stacks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|stack_name|`utf8`|
|name|`utf8`|
|app_visibility|`utf8`|
|attributes|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|