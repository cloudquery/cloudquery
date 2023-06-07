# Table: azure_authorization_role_definitions

This table shows data for Azure Authorization Role Definitions.

https://learn.microsoft.com/en-us/rest/api/authorization/role-definitions/list?tabs=HTTP#roledefinition

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


