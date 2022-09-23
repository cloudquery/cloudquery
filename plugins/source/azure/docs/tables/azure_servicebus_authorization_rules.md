# Table: azure_servicebus_authorization_rules


The primary key for this table is **id**.

## Relations
The following tables depend on `azure_servicebus_authorization_rules`:
  - [`azure_servicebus_access_keys`](azure_servicebus_access_keys.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|servicebus_topic_id|UUID|
|rights|StringArray|
|system_data|JSON|
|id (PK)|String|
|name|String|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|