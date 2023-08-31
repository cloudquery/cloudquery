# Table: azure_security_pricings

This table shows data for Azure Security Pricings.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/pricings/list?tabs=HTTP#pricing

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

### Ensure that Azure Defender is set to On for App Service (Automatic)

```sql
SELECT
  'Ensure that Azure Defender is set to On for App Service (Automatic)'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'pricingTier' = 'Standard' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_pricings AS asp
WHERE
  name = 'AppServices';
```

### Ensure that Azure Defender is set to On for Container Registries (Automatic)

```sql
SELECT
  'Ensure that Azure Defender is set to On for Container Registries (Automatic)'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'pricingTier' = 'Standard' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_pricings AS asp
WHERE
  name = 'ContainerRegistry';
```

### Ensure that Azure Defender is set to On for Kubernetes (Automatic)

```sql
SELECT
  'Ensure that Azure Defender is set to On for Kubernetes (Automatic)' AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'pricingTier' = 'Standard' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_pricings AS asp
WHERE
  name = 'KubernetesService';
```

### Ensure that Azure Defender is set to On for Key Vault (Manual)

```sql
SELECT
  'Ensure that Azure Defender is set to On for Key Vault (Manual)' AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'pricingTier' = 'Standard' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_pricings AS asp
WHERE
  name = 'KeyVaults';
```

### Ensure that Azure Defender is set to On for Servers (Automatic)

```sql
SELECT
  'Ensure that Azure Defender is set to On for Servers (Automatic)' AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'pricingTier' = 'Standard' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_pricings AS asp
WHERE
  name = 'VirtualMachines';
```

### Ensure that Azure Defender is set to On for Azure SQL database servers (Automatic)

```sql
SELECT
  'Ensure that Azure Defender is set to On for Azure SQL database servers (Automatic)'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'pricingTier' = 'Standard' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_pricings AS asp
WHERE
  name = 'SqlServers';
```

### Ensure that Azure Defender is set to On for SQL servers on machines (Automatic)

```sql
SELECT
  'Ensure that Azure Defender is set to On for SQL servers on machines (Automatic)'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'pricingTier' = 'Standard' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_pricings AS asp
WHERE
  name = 'SqlserverVirtualMachines';
```

### Ensure that Azure Defender is set to On for Storage (Automatic)

```sql
SELECT
  'Ensure that Azure Defender is set to On for Storage (Automatic)' AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'pricingTier' = 'Standard' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_pricings AS asp
WHERE
  name = 'StorageAccounts';
```


