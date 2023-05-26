# Table: azure_appservice_web_app_configurations

This table shows data for Azure App Service Web App Configurations.

https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list-configurations#siteconfigresource

The primary key for this table is **id**.

## Relations

This table depends on [azure_appservice_web_apps](azure_appservice_web_apps).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|kind|utf8|
|properties|json|
|id (PK)|utf8|
|name|utf8|
|type|utf8|