# Table: azure_sql_managed_instance_encryption_protectors

This table shows data for Azure SQL Managed Instance Encryption Protectors.

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/managed-instance-encryption-protectors/list-by-instance?tabs=HTTP#managedinstanceencryptionprotector

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_managed_instances](azure_sql_managed_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
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


