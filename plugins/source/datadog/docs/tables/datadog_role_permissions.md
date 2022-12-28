# Table: datadog_role_permissions

The primary key for this table is **_cq_id**.

## Relations

This table depends on [datadog_roles](datadog_roles.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_name|String|
|attributes|JSON|
|id|String|
|type|String|