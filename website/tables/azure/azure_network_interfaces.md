# Table: azure_network_interfaces

This table shows data for Azure Network Interfaces.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-interfaces/list?tabs=HTTP#networkinterface

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_interfaces:
  - [azure_network_interface_ip_configurations](azure_network_interface_ip_configurations)

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
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|