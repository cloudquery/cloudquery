# Table: azure_containerinstance_container_groups

This table shows data for Azure Container Instance Container Groups.

https://learn.microsoft.com/en-us/rest/api/container-instances/container-groups/list?tabs=HTTP#containergroup

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|identity|`json`|
|location|`utf8`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|