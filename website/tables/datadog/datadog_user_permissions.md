# Table: datadog_user_permissions

This table shows data for Datadog User Permissions.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [datadog_users](datadog_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id (PK)|uuid|
|_cq_parent_id|uuid|
|account_name|utf8|
|attributes|json|
|id|utf8|
|type|utf8|
|additional_properties|json|