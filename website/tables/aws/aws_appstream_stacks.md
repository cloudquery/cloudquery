# Table: aws_appstream_stacks

This table shows data for Amazon AppStream Stacks.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Stack.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_appstream_stacks:
  - [aws_appstream_stack_entitlements](aws_appstream_stack_entitlements)
  - [aws_appstream_stack_user_associations](aws_appstream_stack_user_associations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|name|`utf8`|
|access_endpoints|`json`|
|application_settings|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|display_name|`utf8`|
|embed_host_domains|`list<item: utf8, nullable>`|
|feedback_url|`utf8`|
|redirect_url|`utf8`|
|stack_errors|`json`|
|storage_connectors|`json`|
|streaming_experience_settings|`json`|
|user_settings|`json`|