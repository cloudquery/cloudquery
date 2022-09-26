# Table: azure_authorization_role_definitions


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|id (PK)|String|
|name|String|
|type|String|
|role_name|String|
|description|String|
|permissions|JSON|
|assignable_scopes|StringArray|
|role_type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|