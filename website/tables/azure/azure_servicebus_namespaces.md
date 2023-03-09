# Table: azure_servicebus_namespaces

https://learn.microsoft.com/en-us/rest/api/servicebus/stable/namespaces/list?tabs=HTTP#sbnamespace

The primary key for this table is **id**.

## Relations

The following tables depend on azure_servicebus_namespaces:
  - [azure_servicebus_namespace_topics](azure_servicebus_namespace_topics)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|identity|JSON|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|