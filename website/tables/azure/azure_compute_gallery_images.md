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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|