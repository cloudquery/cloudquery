# Table: aws_appstream_stacks

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Stack.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_appstream_stacks:
  - [aws_appstream_stack_entitlements](aws_appstream_stack_entitlements.md)
  - [aws_appstream_stack_user_associations](aws_appstream_stack_user_associations.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|name|String|
|access_endpoints|JSON|
|application_settings|JSON|
|created_time|Timestamp|
|description|String|
|display_name|String|
|embed_host_domains|StringArray|
|feedback_url|String|
|redirect_url|String|
|stack_errors|JSON|
|storage_connectors|JSON|
|streaming_experience_settings|JSON|
|user_settings|JSON|