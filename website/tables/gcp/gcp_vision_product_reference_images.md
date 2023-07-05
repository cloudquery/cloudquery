# Table: gcp_vision_product_reference_images

This table shows data for GCP Vision Product Reference Images.

https://cloud.google.com/vision/docs/reference/rest/v1/projects.locations.products.referenceImages

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vision_products](gcp_vision_products).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|uri|`utf8`|
|bounding_polys|`json`|