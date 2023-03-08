# Table: azure_servicebus_namespace_private_endpoint_connections

https://learn.microsoft.com/en-us/rest/api/servicebus/stable/private-endpoint-connections/list?tabs=HTTP

The primary key for this table is **id**.

## Relations

This table depends on [azure_servicebus_namespaces](azure_servicebus_namespaces).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties|JSON|
|id (PK)|String|
|location|String|
|name|String|
|system_data|JSON|
|type|String|