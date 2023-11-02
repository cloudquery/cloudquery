# Table: aws_appstream_image_builders

This table shows data for Amazon AppStream Image Builders.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ImageBuilder.html

The primary key for this table is **arn**.

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
|appstream_agent_version|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|display_name|`utf8`|
|domain_join_info|`json`|
|enable_default_internet_access|`bool`|
|iam_role_arn|`utf8`|
|image_arn|`utf8`|
|image_builder_errors|`json`|
|instance_type|`utf8`|
|network_access_configuration|`json`|
|platform|`utf8`|
|state|`utf8`|
|state_change_reason|`json`|
|vpc_config|`json`|