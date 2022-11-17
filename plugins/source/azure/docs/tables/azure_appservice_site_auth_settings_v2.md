# Table: azure_appservice_site_auth_settings_v2

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2#SiteAuthSettingsV2

The primary key for this table is **id**.

## Relations
This table depends on [azure_appservice_sites](azure_appservice_sites.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|kind|String|
|global_validation|JSON|
|http_settings|JSON|
|identity_providers|JSON|
|login|JSON|
|platform|JSON|
|id (PK)|String|
|name|String|
|type|String|
|site_id|String|