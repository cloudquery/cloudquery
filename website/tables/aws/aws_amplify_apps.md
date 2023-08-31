# Table: aws_amplify_apps

This table shows data for Amplify Apps.

https://docs.aws.amazon.com/amplify/latest/APIReference/API_ListApps.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|app_arn|`utf8`|
|app_id|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|default_domain|`utf8`|
|description|`utf8`|
|enable_basic_auth|`bool`|
|enable_branch_auto_build|`bool`|
|environment_variables|`json`|
|name|`utf8`|
|platform|`utf8`|
|repository|`utf8`|
|update_time|`timestamp[us, tz=UTC]`|
|auto_branch_creation_config|`json`|
|auto_branch_creation_patterns|`list<item: utf8, nullable>`|
|basic_auth_credentials|`utf8`|
|build_spec|`utf8`|
|custom_headers|`utf8`|
|custom_rules|`json`|
|enable_auto_branch_creation|`bool`|
|enable_branch_auto_deletion|`bool`|
|iam_service_role_arn|`utf8`|
|production_branch|`json`|
|repository_clone_method|`utf8`|
|tags|`json`|