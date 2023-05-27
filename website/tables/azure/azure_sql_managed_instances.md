# Table: azure_sql_managed_instances

This table shows data for Azure SQL Managed Instances.

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/managed-instances/list?tabs=HTTP#managedinstance

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_managed_instances:
  - [azure_sql_managed_instance_encryption_protectors](azure_sql_managed_instance_encryption_protectors)
  - [azure_sql_managed_instance_vulnerability_assessments](azure_sql_managed_instance_vulnerability_assessments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|