# Table: datadog_users

This table shows data for Datadog Users.

The composite primary key for this table is (**account_name**, **id**).

## Relations

The following tables depend on datadog_users:
  - [datadog_user_permissions](datadog_user_permissions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|attributes|`json`|
|id (PK)|`utf8`|
|relationships|`json`|
|type|`utf8`|
|additional_properties|`json`|