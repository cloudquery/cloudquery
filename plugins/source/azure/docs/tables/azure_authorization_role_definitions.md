# Table: azure_authorization_role_definitions



The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|name|String|
|type|String|
|role_name|String|
|description|String|
|permissions|JSON|
|assignable_scopes|StringArray|
|role_type|String|