# Table: aws_appstream_images

This table shows data for Amazon AppStream Images.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Image.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|name|String|
|applications|JSON|
|appstream_agent_version|String|
|arn (PK)|String|
|base_image_arn|String|
|created_time|Timestamp|
|description|String|
|display_name|String|
|image_builder_name|String|
|image_builder_supported|Bool|
|image_errors|JSON|
|image_permissions|JSON|
|platform|String|
|public_base_image_released_date|Timestamp|
|state|String|
|state_change_reason|JSON|
|visibility|String|