# Table: datadog_role_permissions

This table shows data for Datadog Role Permissions.

The composite primary key for this table is (**account_name**, **role_id**, **id**).

## Relations

This table depends on [datadog_roles](datadog_roles).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|role_id (PK)|`utf8`|
|attributes|`json`|
|id (PK)|`utf8`|
|type|`utf8`|
|additional_properties|`json`|