# Table: azure_monitor_subscription_diagnostic_settings

This table shows data for Azure Monitor Subscription Diagnostic Settings.

https://learn.microsoft.com/en-us/rest/api/monitor/subscription-diagnostic-settings/list?tabs=HTTP#subscriptiondiagnosticsettingsresource

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
|system_data|`json`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure Diagnostic Setting captures appropriate categories

```sql
WITH
  diagnostic_settings
    AS (
      SELECT
        subscription_id,
        id,
        (logs->>'enabled')::BOOL AS enabled,
        logs->>'category' AS category
      FROM
        azure_monitor_subscription_diagnostic_settings AS a,
        jsonb_array_elements(properties->'logs') AS logs
    ),
  required_settings
    AS (
      SELECT
        *
      FROM
        diagnostic_settings
      WHERE
        category IN ('Administrative', 'Alert', 'Policy', 'Security')
    )
SELECT
  'Ensure Diagnostic Setting captures appropriate categories' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE WHEN count(id) = 4 THEN 'pass' ELSE 'fail' END AS status
FROM
  required_settings
WHERE
  enabled
GROUP BY
  subscription_id, id;
```


