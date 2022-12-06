# Table: azure_subscriptions

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions#Subscription

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|authorization_source|String|
|managed_by_tenants|JSON|
|subscription_policies|JSON|
|tags|JSON|
|display_name|String|
|id (PK)|String|
|state|String|
|tenant_id|String|