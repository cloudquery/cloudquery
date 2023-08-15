# Table: azure_servicebus_namespace_topic_authorization_rules

This table shows data for Azure Service Bus Namespace Topic Authorization Rules.

https://learn.microsoft.com/en-us/rest/api/servicebus/stable/topics%20%E2%80%93%20authorization%20rules/list-authorization-rules?tabs=HTTP#sbauthorizationrule

The primary key for this table is **id**.

## Relations

This table depends on [azure_servicebus_namespace_topics](azure_servicebus_namespace_topics).

The following tables depend on azure_servicebus_namespace_topic_authorization_rules:
  - [azure_servicebus_namespace_topic_rule_access_keys](azure_servicebus_namespace_topic_rule_access_keys)

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