# Table: azure_web_site_auth_settings_v2

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web#SiteAuthSettingsV2

The primary key for this table is **id**.

## Relations
This table depends on [azure_web_apps](azure_web_apps.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|web_app_id|String|
|platform|JSON|
|global_validation|JSON|
|identity_providers|JSON|
|login|JSON|
|http_settings|JSON|
|id (PK)|String|
|name|String|
|kind|String|
|type|String|