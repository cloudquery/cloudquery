# Table: azure_servicebus_namespace_topic_rule_access_keys

https://learn.microsoft.com/en-us/rest/api/servicebus/stable/topics%20%E2%80%93%20authorization%20rules/list-keys?tabs=HTTP#accesskeys

The primary key for this table is **_cq_id**.

## Relations

This table depends on [azure_servicebus_namespace_topic_authorization_rules](azure_servicebus_namespace_topic_authorization_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|alias_primary_connection_string|String|
|alias_secondary_connection_string|String|
|key_name|String|
|primary_connection_string|String|
|primary_key|String|
|secondary_connection_string|String|
|secondary_key|String|