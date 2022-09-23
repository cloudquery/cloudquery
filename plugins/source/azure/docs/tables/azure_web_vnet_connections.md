# Table: azure_web_vnet_connections


The primary key for this table is **id**.

## Relations
This table depends on [`azure_web_apps`](azure_web_apps.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|web_app_id|UUID|
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|