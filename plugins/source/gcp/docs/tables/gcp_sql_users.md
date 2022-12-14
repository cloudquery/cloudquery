# Table: gcp_sql_users



The composite primary key for this table is (**project_id**, **instance**, **name**).

## Relations
This table depends on [gcp_sql_instances](gcp_sql_instances.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|instance (PK)|String|
|name (PK)|String|
|dual_password_type|String|
|etag|String|
|host|String|
|kind|String|
|password|String|
|password_policy|JSON|
|project|String|
|sqlserver_user_details|JSON|
|type|String|