# Table: azure_network_load_balancers

This table shows data for Azure Network Load Balancers.

https://learn.microsoft.com/en-us/rest/api/load-balancer/load-balancers/list?tabs=HTTP#loadbalancer

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|extended_location|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|