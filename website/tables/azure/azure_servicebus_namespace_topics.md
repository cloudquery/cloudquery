# Table: azure_servicebus_namespace_topics

This table shows data for Azure Service Bus Namespace Topics.

https://learn.microsoft.com/en-us/rest/api/servicebus/stable/topics/list-by-namespace?tabs=HTTP#sbtopic

The primary key for this table is **id**.

## Relations

This table depends on [azure_servicebus_namespaces](azure_servicebus_namespaces).

The following tables depend on azure_servicebus_namespace_topics:
  - [azure_servicebus_namespace_topic_authorization_rules](azure_servicebus_namespace_topic_authorization_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|