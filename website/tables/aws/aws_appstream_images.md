# Table: aws_appstream_images

This table shows data for Amazon AppStream Images.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Image.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|name|`utf8`|
|applications|`json`|
|appstream_agent_version|`utf8`|
|arn (PK)|`utf8`|
|base_image_arn|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|display_name|`utf8`|
|image_builder_name|`utf8`|
|image_builder_supported|`bool`|
|image_errors|`json`|
|image_permissions|`json`|
|platform|`utf8`|
|public_base_image_released_date|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|state_change_reason|`json`|
|visibility|`utf8`|