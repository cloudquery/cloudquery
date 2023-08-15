# Table: azure_sql_server_admins

This table shows data for Azure SQL Server Admins.

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/server-azure-ad-administrators/list-by-server?tabs=HTTP#serverazureadadministrator

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
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Azure Active Directory Admin is configured (Automated)

```sql
WITH
  ad_admins_count
    AS (
      SELECT
        ass._cq_id, count(*) AS admins_count
      FROM
        azure_sql_servers AS ass
        LEFT JOIN azure_sql_server_admins AS assa ON
            ass._cq_id = assa._cq_parent_id
      WHERE
        assa.properties->>'administratorType' = 'ActiveDirectory'
      GROUP BY
        ass._cq_id, assa.properties->>'administratorType'
    )
SELECT
  'Ensure that Azure Active Directory Admin is configured (Automated)' AS title,
  s.subscription_id,
  s.id,
  CASE
  WHEN a.admins_count IS NULL OR a.admins_count = 0 THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_servers AS s LEFT JOIN ad_admins_count AS a ON s._cq_id = a._cq_id;
```


