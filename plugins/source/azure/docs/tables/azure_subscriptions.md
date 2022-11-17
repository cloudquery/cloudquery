# Table: azure_subscriptions

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription#Subscription

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|authorization_source|String|
|subscription_policies|JSON|
|display_name|String|
|id (PK)|String|
|state|String|
|subscription_id|String|