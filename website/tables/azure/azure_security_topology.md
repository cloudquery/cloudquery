# Table: azure_security_topology

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/topology/list?tabs=HTTP#topologyresource

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|id (PK)|String|
|location|String|
|name|String|
|properties|JSON|
|type|String|