# Table: azure_compute_gallery_image_versions

This table shows data for Azure Compute Gallery Image Versions.

https://learn.microsoft.com/en-us/rest/api/compute/gallery-image-versions/list-by-gallery-image?tabs=HTTP#galleryimageversion

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_gallery_images](azure_compute_gallery_images).

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