# Table: azure_network_interface_ip_configurations

This table shows data for Azure Network Interface IP Configurations.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-interface-ip-configurations/list?tabs=HTTP#ipconfiguration

The primary key for this table is **id**.

## Relations

This table depends on [azure_network_interfaces](azure_network_interfaces).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|properties|`json`|
|type|`utf8`|
|etag|`utf8`|