# Table: azure_compute_gallery_images

This table shows data for Azure Compute Gallery Images.

https://learn.microsoft.com/en-us/rest/api/compute/gallery-images/list-by-gallery?tabs=HTTP#galleryimage

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_galleries](azure_compute_galleries).

The following tables depend on azure_compute_gallery_images:
  - [azure_compute_gallery_image_versions](azure_compute_gallery_image_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|