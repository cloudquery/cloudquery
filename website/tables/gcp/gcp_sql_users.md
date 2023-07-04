# Table: gcp_sql_users

This table shows data for GCP SQL Users.

https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/users#User

The composite primary key for this table is (**project_id**, **instance**, **name**).

## Relations

This table depends on [gcp_sql_instances](gcp_sql_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|dual_password_type|`utf8`|
|etag|`utf8`|
|host|`utf8`|
|instance (PK)|`utf8`|
|kind|`utf8`|
|name (PK)|`utf8`|
|password|`utf8`|
|password_policy|`json`|
|project|`utf8`|
|sqlserver_user_details|`json`|
|type|`utf8`|