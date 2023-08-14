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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|