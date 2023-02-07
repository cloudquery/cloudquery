# Table: aws_amplify_apps

https://docs.aws.amazon.com/amplify/latest/APIReference/API_ListApps.html

The primary key for this table is **arn**.

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
|app_arn|String|
|app_id|String|
|create_time|Timestamp|
|default_domain|String|
|description|String|
|enable_basic_auth|Bool|
|enable_branch_auto_build|Bool|
|environment_variables|JSON|
|name|String|
|platform|String|
|repository|String|
|update_time|Timestamp|
|auto_branch_creation_config|JSON|
|auto_branch_creation_patterns|StringArray|
|basic_auth_credentials|String|
|build_spec|String|
|custom_headers|String|
|custom_rules|JSON|
|enable_auto_branch_creation|Bool|
|enable_branch_auto_deletion|Bool|
|iam_service_role_arn|String|
|production_branch|JSON|
|repository_clone_method|String|
|tags|JSON|