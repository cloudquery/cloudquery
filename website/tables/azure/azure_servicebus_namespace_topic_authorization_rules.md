# Table: azure_servicebus_namespace_topic_authorization_rules

https://learn.microsoft.com/en-us/rest/api/servicebus/stable/topics/list-by-namespace?tabs=HTTP

The primary key for this table is **id**.

## Relations

This table depends on [azure_servicebus_namespace_topics](azure_servicebus_namespace_topics).

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