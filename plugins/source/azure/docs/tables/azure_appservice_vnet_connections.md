# Table: azure_appservice_vnet_connections

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2#VnetInfoResource

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
|cert_blob|String|
|dns_servers|String|
|is_swift|Bool|
|vnet_resource_id|String|
|cert_thumbprint|String|
|resync_required|Bool|
|routes|JSON|
|id (PK)|String|
|name|String|
|type|String|
|site_id|String|