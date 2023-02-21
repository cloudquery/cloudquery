# Table: gcp_vision_products

https://cloud.google.com/vision/docs/reference/rest/v1/projects.locations.products

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_vision_products:
  - [gcp_vision_product_reference_images](gcp_vision_product_reference_images.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|display_name|String|
|description|String|
|product_category|String|
|product_labels|JSON|