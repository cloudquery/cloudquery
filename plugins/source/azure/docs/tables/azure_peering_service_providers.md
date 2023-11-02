# Table: azure_peering_service_providers

This table shows data for Azure Peering Service Providers.

https://learn.microsoft.com/en-us/rest/api/peering/peering-service-providers/list?tabs=HTTP#peeringserviceprovider

The composite primary key for this table is (**subscription_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|properties|`json`|
|id|`utf8`|
|name (PK)|`utf8`|
|type|`utf8`|