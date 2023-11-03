# Table: azure_monitor_activity_log_alerts

This table shows data for Azure Monitor Activity Log Alerts.

https://learn.microsoft.com/en-us/rest/api/monitor/activity-log-alerts/list-by-subscription-id?tabs=HTTP#activitylogalertresource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Activity Log Alert exists for Create or Update Network Security Group

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals = 'Microsoft.Network/networkSecurityGroups/write'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Create or Update Network Security Group'
    AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```

### Ensure that Activity Log Alert exists for Create or Update Network Security Group Rule

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals
          = 'Microsoft.Network/networkSecurityGroups/securityRules/write'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Create or Update Network Security Group Rule'
    AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```

### Ensure that Activity Log Alert exists for Create or Update or Delete SQL Server Firewall Rule

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals = 'Microsoft.Sql/servers/firewallRules/write'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Create or Update or Delete SQL Server Firewall Rule'
    AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```

### Ensure that Activity Log Alert exists for Create or Update Security Solution

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals = 'Microsoft.Security/securitySolutions/write'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Create or Update Security Solution'
    AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```

### Ensure that Activity Log Alert exists for Create Policy Assignment

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals = 'Microsoft.Security/policies/write'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Create Policy Assignment' AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```

### Ensure that Activity Log Alert exists for Delete Network Security Group

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals = 'Microsoft.Network/networkSecurityGroups/delete'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Delete Network Security Group'
    AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```

### Ensure that Activity Log Alert exists for Delete Network Security Group Rule

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals
          = 'Microsoft.Network/networkSecurityGroups/securityRules/delete'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Delete Network Security Group Rule'
    AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```

### Ensure that Activity Log Alert exists for Delete Policy Assignment

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals = 'Microsoft.Security/policies/delete'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Delete Policy Assignment' AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```

### Ensure that Activity Log Alert exists for Delete Security Solution

```sql
WITH
  fields
    AS (
      SELECT
        subscription_id,
        id,
        location,
        (properties->'enabled')::BOOL AS enabled,
        conditions->>'field' AS field,
        conditions->>'equals' AS equals
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements(properties->'condition'->'allOf') AS conditions
    ),
  scopes
    AS (
      SELECT
        subscription_id, id, scope
      FROM
        azure_monitor_activity_log_alerts,
        jsonb_array_elements_text(properties->'scopes') AS scope
    ),
  conditions
    AS (
      SELECT
        fields.subscription_id AS subscription_id,
        fields.id AS id,
        scopes.scope AS scope,
        location = 'global'
        AND enabled
        AND equals = 'Microsoft.Security/securitySolutions/delete'
        AND scopes.scope
          ~ e'^\\/subscriptions\\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$'
          AS condition
      FROM
        fields JOIN scopes ON fields.id = scopes.id
      WHERE
        field = 'operationName'
    )
SELECT
  'Ensure that Activity Log Alert exists for Delete Security Solution' AS title,
  subscription_id AS subscription_id,
  scope AS resrouce_id,
  bool_or(condition) AS status
FROM
  conditions
GROUP BY
  subscription_id, scope;
```


