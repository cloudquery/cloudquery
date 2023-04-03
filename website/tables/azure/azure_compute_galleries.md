# Table: azure_compute_galleries

This table shows data for Azure Compute Galleries.

https://learn.microsoft.com/en-us/rest/api/compute/galleries/list?tabs=HTTP#gallery

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_galleries:
  - [azure_compute_gallery_images](azure_compute_gallery_images)

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