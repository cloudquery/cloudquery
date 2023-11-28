# Table: azure_network_application_gateways

This table shows data for Azure Network Application Gateways.

https://learn.microsoft.com/en-us/rest/api/application-gateway/application-gateways/list?tabs=HTTP#applicationgateway

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|identity|`json`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|