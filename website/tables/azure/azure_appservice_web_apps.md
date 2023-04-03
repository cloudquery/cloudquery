# Table: azure_appservice_web_apps

This table shows data for Azure App Service Web Apps.

https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list#site

The primary key for this table is **id**.

## Relations

The following tables depend on azure_appservice_web_apps:
  - [azure_appservice_web_app_auth_settings](azure_appservice_web_app_auth_settings)
  - [azure_appservice_web_app_configurations](azure_appservice_web_app_configurations)
  - [azure_appservice_web_app_vnet_connections](azure_appservice_web_app_vnet_connections)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|kind|String|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|