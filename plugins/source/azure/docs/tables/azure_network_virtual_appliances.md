# Table: azure_network_virtual_appliances

This table shows data for Azure Network Virtual Appliances.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-virtual-appliances/list?tabs=HTTP#networkvirtualappliance

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
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|