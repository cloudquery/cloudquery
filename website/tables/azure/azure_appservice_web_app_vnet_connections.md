# Table: azure_appservice_web_app_vnet_connections

This table shows data for Azure App Service Web App Vnet Connections.

https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list-vnet-connections#vnetinforesource

The primary key for this table is **id**.

## Relations

This table depends on [azure_appservice_web_apps](azure_appservice_web_apps).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|kind|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|