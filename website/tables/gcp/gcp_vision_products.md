# Table: gcp_vision_products

This table shows data for GCP Vision Products.

https://cloud.google.com/vision/docs/reference/rest/v1/projects.locations.products

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_vision_products:
  - [gcp_vision_product_reference_images](gcp_vision_product_reference_images)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|description|`utf8`|
|product_category|`utf8`|
|product_labels|`json`|