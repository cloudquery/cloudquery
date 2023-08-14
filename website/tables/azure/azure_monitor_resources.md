# Table: azure_monitor_resources

This table shows data for Azure Monitor Resources.

https://learn.microsoft.com/en-us/rest/api/resources/resources/list#genericresourceexpanded

The primary key for this table is **id**.

## Relations

The following tables depend on azure_monitor_resources:
  - [azure_monitor_diagnostic_settings](azure_monitor_diagnostic_settings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|extended_location|`json`|
|identity|`json`|
|kind|`utf8`|
|location|`utf8`|
|managed_by|`utf8`|
|plan|`json`|
|sku|`json`|
|tags|`json`|
|changed_time|`timestamp[us, tz=UTC]`|
|created_time|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|name|`utf8`|
|provisioning_state|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that Diagnostic Logs are enabled for all services which support it.

```sql
SELECT
  'Ensure that Diagnostic Logs are enabled for all services which support it.'
    AS title,
  amr.subscription_id AS subscription_id,
  amr.id AS resource_id,
  CASE WHEN amds.id IS DISTINCT FROM NULL THEN 'pass' ELSE 'fail' END AS status
FROM
  azure_monitor_resources AS amr
  LEFT JOIN azure_monitor_diagnostic_settings AS amds ON
      amr._cq_id = amds._cq_parent_id;
```

### Ensure that logging for Azure Key Vault is ''Enabled''

```sql
WITH
  diagnosis_logs
    AS (
      SELECT
        amds.subscription_id,
        amds.id
        || '/'
        || (COALESCE(logs->>'category', logs->>'categoryGroup'))::STRING
          AS id,
        logs->>'category' IS DISTINCT FROM NULL AS hascategory,
        (logs->'retentionPolicy'->>'days')::INT8 >= 180 AS satisfyretentiondays
      FROM
        azure_monitor_resources AS amr
        LEFT JOIN azure_monitor_diagnostic_settings AS amds ON
            amr._cq_id = amds._cq_parent_id,
        jsonb_array_elements(amds.properties->'logs') AS logs
      WHERE
        amr.type = 'Microsoft.KeyVault/vaults'
    )
SELECT
  e'Ensure that logging for Azure Key Vault is \'Enabled\'' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN hascategory AND satisfyretentiondays THEN 'pass'
  ELSE 'fail'
  END
    AS status
FROM
  diagnosis_logs;
```

### Ensure that a ''Diagnostics Setting'' exists

```sql
SELECT
  e'Ensure that a \'Diagnostics Setting\' exists' AS title,
  amr.subscription_id AS subscription_id,
  amr.id AS resource_id,
  CASE WHEN amds.properties IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  azure_monitor_resources AS amr
  LEFT JOIN azure_monitor_diagnostic_settings AS amds ON
      amr._cq_id = amds._cq_parent_id;
```


