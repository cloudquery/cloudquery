# Table: azure_sql_managed_instances

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/managed-instances/list?tabs=HTTP#managedinstance

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_managed_instances:
  - [azure_sql_managed_instance_encryption_protectors](azure_sql_managed_instance_encryption_protectors.md)
  - [azure_sql_managed_instance_vulnerability_assessments](azure_sql_managed_instance_vulnerability_assessments.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|identity|JSON|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|