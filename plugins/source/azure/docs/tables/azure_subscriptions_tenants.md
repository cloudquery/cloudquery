# Table: azure_subscriptions_tenants

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions#TenantIDDescription

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|country|String|
|country_code|String|
|default_domain|String|
|display_name|String|
|domains|JSON|
|id (PK)|String|
|tenant_branding_logo_url|String|
|tenant_category|String|
|tenant_id|String|
|tenant_type|String|