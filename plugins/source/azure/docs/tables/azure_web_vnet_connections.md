# Table: azure_web_vnet_connections

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web#VnetInfo

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
|vnet_resource_id|String|
|cert_thumbprint|String|
|cert_blob|String|
|routes|JSON|
|resync_required|Bool|
|dns_servers|String|
|is_swift|Bool|
|id (PK)|String|
|name|String|
|kind|String|
|type|String|