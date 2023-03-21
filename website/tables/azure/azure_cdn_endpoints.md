# Table: azure_cdn_endpoints

This table shows data for Azure Content Delivery Network (CDN) Endpoints.

https://learn.microsoft.com/en-us/rest/api/cdn/endpoints/list-by-profile?tabs=HTTP#endpoint

The primary key for this table is **id**.

## Relations

This table depends on [azure_cdn_profiles](azure_cdn_profiles).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|location|String|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|