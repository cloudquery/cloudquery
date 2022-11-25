# Table: aws_appstream_image_builders

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ImageBuilder.html

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
|name|String|
|access_endpoints|JSON|
|appstream_agent_version|String|
|created_time|Timestamp|
|description|String|
|display_name|String|
|domain_join_info|JSON|
|enable_default_internet_access|Bool|
|iam_role_arn|String|
|image_arn|String|
|image_builder_errors|JSON|
|instance_type|String|
|network_access_configuration|JSON|
|platform|String|
|state|String|
|state_change_reason|JSON|
|vpc_config|JSON|