# Table: azure_resources_policy_assignments


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|