# Table: gcp_sql_users

This table shows data for GCP SQL Users.

https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/users#User

The composite primary key for this table is (**project_id**, **instance**, **name**).

## Relations

This table depends on [gcp_sql_instances](gcp_sql_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|dual_password_type|String|
|etag|String|
|host|String|
|instance (PK)|String|
|kind|String|
|name (PK)|String|
|password|String|
|password_policy|JSON|
|project|String|
|sqlserver_user_details|JSON|
|type|String|