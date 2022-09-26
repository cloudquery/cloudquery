# Table: azure_monitor_log_profiles


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|storage_account_id|String|
|service_bus_rule_id|String|
|locations|StringArray|
|categories|StringArray|
|retention_policy|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|kind|String|
|etag|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|