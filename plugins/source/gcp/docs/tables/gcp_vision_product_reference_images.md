# Table: gcp_vision_product_reference_images

https://cloud.google.com/vision/docs/reference/rest/v1/projects.locations.products.referenceImages

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vision_products](gcp_vision_products.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|uri|String|
|bounding_polys|JSON|