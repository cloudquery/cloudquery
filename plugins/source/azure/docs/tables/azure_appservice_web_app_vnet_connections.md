# Table: azure_appservice_web_app_vnet_connections

The primary key for this table is **_cq_id**.

## Relations

This table depends on [azure_appservice_web_apps](azure_appservice_web_apps.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|kind|String|
|properties|JSON|
|id|String|
|name|String|
|type|String|