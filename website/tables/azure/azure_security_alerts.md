# Table: azure_security_alerts

This table shows data for Azure Security Alerts.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/alerts/list?tabs=HTTP#alert

The primary key for this table is **id**.

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
|name|String|
|type|String|