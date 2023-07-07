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

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### SQL managed instances should use customer-managed keys to encrypt data at rest

```sql
WITH
  protected_instances
    AS (
      SELECT
        s.id AS instance_id
      FROM
        azure_sql_managed_instances AS s
        LEFT JOIN azure_sql_managed_instance_encryption_protectors AS ep ON
            s._cq_id = ep._cq_parent_id
      WHERE
        ep.properties->>'serverKeyType' = 'AzureKeyVault'
        AND (ep.properties->>'uri') IS NOT NULL
    )
SELECT
  'SQL managed instances should use customer-managed keys to encrypt data at rest'
    AS title,
  i.subscription_id,
  i.id AS instance_id,
  CASE
  WHEN p.instance_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_managed_instances AS i
  LEFT JOIN protected_instances AS p ON p.instance_id = i.id;
```

### Vulnerability assessment should be enabled on SQL Managed Instance

```sql
WITH
  protected_instances
    AS (
      SELECT
        s.id AS instance_id
      FROM
        azure_sql_managed_instances AS s
        LEFT JOIN azure_sql_managed_instance_vulnerability_assessments AS va ON
            s._cq_id = va._cq_parent_id
      WHERE
        (va.properties->'recurringScans'->>'isEnabled')::BOOL IS true
    )
SELECT
  'Vulnerability assessment should be enabled on SQL Managed Instance' AS title,
  i.subscription_id,
  i.id AS instance_id,
  CASE
  WHEN p.instance_id IS NULL THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_managed_instances AS i
  LEFT JOIN protected_instances AS p ON p.instance_id = i.id;
```


