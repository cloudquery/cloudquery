# Table: azure_servicebus_access_keys


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|servicebus_authorization_rule_id|String|
|primary_connection_string|String|
|secondary_connection_string|String|
|alias_primary_connection_string|String|
|alias_secondary_connection_string|String|
|primary_key|String|
|secondary_key|String|
|key_name|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|