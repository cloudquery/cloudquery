# Table: azure_resources_policy_assignments

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy#Assignment

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|display_name|String|
|policy_definition_id|String|
|scope|String|
|not_scopes|StringArray|
|parameters|JSON|
|description|String|
|enforcement_mode|String|
|id (PK)|String|
|type|String|
|name|String|
|sku|JSON|
|location|String|
|identity|JSON|