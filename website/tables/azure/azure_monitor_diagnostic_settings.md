# Table: azure_monitor_diagnostic_settings

This table shows data for Azure Monitor Diagnostic Settings.

https://learn.microsoft.com/en-us/rest/api/monitor/diagnostic-settings/list?tabs=HTTP#diagnosticsettingsresource

The primary key for this table is **id**.

## Relations

This table depends on [azure_monitor_resources](azure_monitor_resources).

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
|resource_id|`utf8`|

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

### Ensure the storage account containing the container with activity logs is encrypted with BYOK (Use Your Own Key)

```sql
SELECT
  'Ensure the storage account containing the container with activity logs is encrypted with BYOK (Use Your Own Key)'
    AS title,
  asa.subscription_id AS subscription_id,
  asa.id AS resource_id,
  CASE
  WHEN asa.properties->'encryption'->>'keySource' = 'Microsoft.Keyvault'
  AND asa.properties->'encryption'->'keyvaultproperties' IS DISTINCT FROM NULL
  THEN 'pass'
  ELSE 'fail'
  END
    AS status
FROM
  azure_storage_accounts AS asa
  JOIN azure_monitor_diagnostic_settings AS amds ON
      asa.id = amds.properties->>'storageAccountId'
WHERE
  (amds.properties->>'storageAccountId') IS NOT NULL;
```


