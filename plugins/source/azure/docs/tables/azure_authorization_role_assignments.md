# Table: azure_authorization_role_assignments

This table shows data for Azure Authorization Role Assignments.

https://learn.microsoft.com/en-us/rest/api/authorization/role-assignments/get?tabs=HTTP#roleassignment

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### There should be more than one owner assigned to your subscription

```sql
WITH
  owners_in_sub
    AS (
      SELECT
        a.subscription_id, count(*) AS owners, d.id AS id
      FROM
        azure_authorization_role_assignments AS a
        JOIN azure_authorization_role_definitions AS d ON
            a.properties->>'roleDefinitionId' = d.id
      WHERE
        a.properties->>'roleName' = 'Owner'
        AND a.properties->>'roleType' = 'BuiltInRole'
      GROUP BY
        d.id, a.subscription_id
    )
SELECT
  'There should be more than one owner assigned to your subscription' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE WHEN owners > 1 THEN 'fail' ELSE 'pass' END AS status
FROM
  owners_in_sub;
```

### A maximum of 3 owners should be designated for your subscription

```sql
WITH
  owners_in_sub
    AS (
      SELECT
        a.subscription_id, count(*) AS owners, d.id AS id
      FROM
        azure_authorization_role_assignments AS a
        JOIN azure_authorization_role_definitions AS d ON
            a.properties->>'roleDefinitionId' = d.id
      WHERE
        a.properties->>'roleName' = 'Owner'
        AND a.properties->>'roleType' = 'BuiltInRole'
      GROUP BY
        d.id, a.subscription_id
    )
SELECT
  'A maximum of 3 owners should be designated for your subscription' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE WHEN owners > 3 THEN 'fail' ELSE 'pass' END AS status
FROM
  owners_in_sub;
```


