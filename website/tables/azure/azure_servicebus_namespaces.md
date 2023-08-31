# Table: azure_servicebus_namespaces

This table shows data for Azure Service Bus Namespaces.

https://learn.microsoft.com/en-us/rest/api/servicebus/stable/namespaces/list?tabs=HTTP#sbnamespace

The primary key for this table is **id**.

## Relations

The following tables depend on azure_servicebus_namespaces:
  - [azure_servicebus_namespace_topics](azure_servicebus_namespace_topics)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|