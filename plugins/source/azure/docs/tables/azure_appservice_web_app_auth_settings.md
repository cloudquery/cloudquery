# Table: azure_appservice_web_app_auth_settings

https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/get-auth-settings#siteauthsettings

The primary key for this table is **id**.

## Relations

This table depends on [azure_appservice_web_apps](azure_appservice_web_apps.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|kind|String|
|properties|JSON|
|name|String|
|type|String|