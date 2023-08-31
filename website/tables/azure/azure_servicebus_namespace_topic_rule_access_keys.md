# Table: azure_servicebus_namespace_topic_rule_access_keys

This table shows data for Azure Service Bus Namespace Topic Rule Access Keys.

https://learn.microsoft.com/en-us/rest/api/servicebus/stable/topics%20%E2%80%93%20authorization%20rules/list-keys?tabs=HTTP#accesskeys

The composite primary key for this table is (**rule_id**, **key_name**).

## Relations

This table depends on [azure_servicebus_namespace_topic_authorization_rules](azure_servicebus_namespace_topic_authorization_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|rule_id (PK)|`utf8`|
|alias_primary_connection_string|`utf8`|
|alias_secondary_connection_string|`utf8`|
|key_name (PK)|`utf8`|
|primary_connection_string|`utf8`|
|primary_key|`utf8`|
|secondary_connection_string|`utf8`|
|secondary_key|`utf8`|