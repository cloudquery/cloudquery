# Table: azure_compute_images

This table shows data for Azure Compute Images.

https://learn.microsoft.com/en-us/rest/api/compute/images/list?tabs=HTTP#image

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|