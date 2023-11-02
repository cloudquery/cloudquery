# Table: azure_network_bastion_hosts

This table shows data for Azure Network Bastion Hosts.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/bastion-hosts/list?tabs=HTTP#bastionhost

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
|sku|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|