# Table: azure_network_azure_firewalls

This table shows data for Azure Network Azure Firewalls.

https://learn.microsoft.com/en-us/rest/api/firewall/azure-firewalls/list?tabs=HTTP#azurefirewall

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|