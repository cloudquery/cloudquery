# Table: azure_appservice_web_apps

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on azure_appservice_web_apps:
  - [azure_appservice_web_app_auth_settings](azure_appservice_web_app_auth_settings.md)
  - [azure_appservice_web_app_vnet_connections](azure_appservice_web_app_vnet_connections.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|kind|String|
|properties|JSON|
|tags|JSON|
|id|String|
|name|String|
|type|String|