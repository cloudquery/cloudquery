# Table: azure_security_topology

This table shows data for Azure Security Topology.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/topology/list?tabs=HTTP#topologyresource

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|properties|`json`|
|type|`utf8`|