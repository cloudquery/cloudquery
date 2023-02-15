# Table: azure_network_interface_ip_configurations

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-interface-ip-configurations/list?tabs=HTTP#ipconfiguration

The composite primary key for this table is (**network_interface_name**, **id**).

## Relations

This table depends on [azure_network_interfaces](azure_network_interfaces.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|network_interface_name (PK)|String|
|id (PK)|String|
|name|String|
|properties|JSON|
|type|String|
|etag|String|