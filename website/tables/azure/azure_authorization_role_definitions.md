# Table: azure_authorization_role_definitions

This table shows data for Azure Authorization Role Definitions.

https://learn.microsoft.com/en-us/rest/api/authorization/role-definitions/list?tabs=HTTP#roledefinition

The primary key for this table is **id**.

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

### Ensure That No Custom Subscription Administrator Roles Exist

```sql
--check if definition matches scopes
WITH
  custom_roles
    AS (
      SELECT
        *
      FROM
        azure_authorization_role_definitions
      WHERE
        properties->>'type' = 'CustomRole'
    ),
  assignable_scopes
    AS (
      SELECT
        _cq_id, scope AS assignable_scope
      FROM
        custom_roles,
        jsonb_array_elements_text(properties->'assignableScopes') AS scope
    ),
  meets_scopes
    AS (
      SELECT
        _cq_id,
        bool_or(
          assignable_scope = '/'
          OR assignable_scope
            ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
        )
          AS has_wide_scope
      FROM
        assignable_scopes
      GROUP BY
        _cq_id
    ),
  definition_actions
    AS (
      SELECT
        _cq_id, actions AS action
      FROM
        custom_roles,
        jsonb_array_elements(properties->'permissions') AS p,
        jsonb_array_elements_text(p->'actions') AS actions
    ),
  meets_actions
    AS (
      SELECT
        _cq_id, bool_or(action = '*') AS has_all_action
      FROM
        definition_actions
      GROUP BY
        _cq_id
    )
SELECT
  'Ensure That No Custom Subscription Administrator Roles Exist' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN has_wide_scope AND has_all_action THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  custom_roles
  JOIN meets_scopes USING (_cq_id)
  JOIN meets_actions USING (_cq_id);
```


