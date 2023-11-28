# Table: azure_appservice_web_apps

This table shows data for Azure App Service Web Apps.

https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list#site

The primary key for this table is **id**.

## Relations

The following tables depend on azure_appservice_web_apps:
  - [azure_appservice_web_app_auth_settings](azure_appservice_web_app_auth_settings.md)
  - [azure_appservice_web_app_configurations](azure_appservice_web_app_configurations.md)
  - [azure_appservice_web_app_vnet_connections](azure_appservice_web_app_vnet_connections.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|identity|`json`|
|kind|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|