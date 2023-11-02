# Table: azure_compute_cloud_services

This table shows data for Azure Compute Cloud Services.

https://learn.microsoft.com/en-us/rest/api/compute/cloud-services/list?tabs=HTTP#cloudservice

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|system_data|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|