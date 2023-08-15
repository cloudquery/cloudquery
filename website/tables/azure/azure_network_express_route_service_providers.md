# Table: azure_network_express_route_service_providers

This table shows data for Azure Network Express Route Service Providers.

https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-service-providers/list?tabs=HTTP#expressrouteserviceprovider

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|name|`utf8`|
|type|`utf8`|