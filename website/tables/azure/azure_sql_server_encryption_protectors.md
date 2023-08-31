# Table: azure_sql_server_encryption_protectors

This table shows data for Azure SQL Server Encryption Protectors.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/encryption-protectors/list-by-server?tabs=HTTP#encryptionprotector

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
|location|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure SQL server"s TDE protector is encrypted with Customer-managed key (Automated)

```sql
SELECT
  'Ensure SQL server"s TDE protector is encrypted with Customer-managed key (Automated)'
    AS title,
  s.subscription_id,
  s.id AS server_id,
  CASE
  WHEN p.kind != 'azurekeyvault'
  OR p.properties->>'serverKeyType' IS DISTINCT FROM 'AzureKeyVault'
  OR (p.properties->>'uri') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s
  LEFT JOIN azure_sql_server_encryption_protectors AS p ON
      s._cq_id = p._cq_parent_id;
```


