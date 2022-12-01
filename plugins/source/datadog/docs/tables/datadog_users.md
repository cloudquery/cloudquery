# Table: datadog_users



The primary key for this table is **id**.

## Relations

The following tables depend on datadog_users:
  - [datadog_user_permissions](datadog_user_permissions.md)
  - [datadog_user_organizations](datadog_user_organizations.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_name|String|
|id (PK)|String|
|attributes|JSON|
|relationships|JSON|
|type|String|