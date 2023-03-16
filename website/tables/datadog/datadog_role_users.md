# Table: datadog_role_users

This table shows data for Datadog Role Users.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [datadog_roles](datadog_roles).

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
|relationships|JSON|
|type|String|
|additional_properties|JSON|