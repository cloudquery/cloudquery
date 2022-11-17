# Table: azure_resource_policy_assignments

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy#Assignment

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|identity|JSON|
|location|String|
|description|String|
|display_name|String|
|enforcement_mode|String|
|non_compliance_messages|JSON|
|not_scopes|StringArray|
|parameters|JSON|
|policy_definition_id|String|
|scope|String|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|