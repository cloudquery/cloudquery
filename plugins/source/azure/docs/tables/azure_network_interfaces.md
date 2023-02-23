# Table: azure_network_interfaces

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/network-interfaces/list?tabs=HTTP#networkinterface

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_interfaces:
  - [azure_network_interface_ip_configurations](azure_network_interface_ip_configurations.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|extended_location|JSON|
|id (PK)|String|
|location|String|
|properties|JSON|
|tags|JSON|
|etag|String|
|name|String|
|type|String|