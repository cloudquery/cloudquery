# Table: azure_cdn_endpoints

This table shows data for Azure Content Delivery Network (CDN) Endpoints.

https://learn.microsoft.com/en-us/rest/api/cdn/endpoints/list-by-profile?tabs=HTTP#endpoint

The primary key for this table is **id**.

## Relations

This table depends on [azure_cdn_profiles](azure_cdn_profiles).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|