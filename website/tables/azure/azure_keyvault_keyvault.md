# Table: azure_keyvault_keyvault

This table shows data for Azure Key Vault Key Vault.

https://learn.microsoft.com/en-us/rest/api/keyvault/keyvault/vaults/get?tabs=HTTP#vault

The primary key for this table is **id**.

## Relations

The following tables depend on azure_keyvault_keyvault:
  - [azure_keyvault_keyvault_keys](azure_keyvault_keyvault_keys)
  - [azure_keyvault_keyvault_secrets](azure_keyvault_keyvault_secrets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|location|`utf8`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Azure Key Vault Managed HSM should have purge protection enabled

```sql
SELECT
  'Azure Key Vault Managed HSM should have purge protection enabled' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (properties->>'enablePurgeProtection')::BOOL IS NOT true
  OR (properties->>'enableSoftDelete')::BOOL IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_keyvault_keyvault_managed_hsms;
```

### Ensure that the expiration date is set on all keys (Automated)

```sql
SELECT
  'Ensure that the expiration date is set on all keys (Automated)' AS title,
  akv.subscription_id AS subscription_id,
  akvk.id AS resource_id,
  CASE
  WHEN (akvk.properties->'attributes'->>'enabled')::BOOL = true
  AND (akvk.properties->'attributes'->>'exp') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_keyvault_keyvault AS akv
  JOIN azure_keyvault_keyvault_keys AS akvk ON akv._cq_id = akvk._cq_parent_id;
```

### Ensure the key vault is recoverable (Automated)

```sql
SELECT
  'Ensure the key vault is recoverable (Automated)' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (properties->>'enableSoftDelete')::BOOL IS NOT true
  OR (properties->>'enablePurgeProtection')::BOOL IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_keyvault_keyvault;
```

### Ensure that the expiration date is set on all Secrets (Automated)

```sql
SELECT
  'Ensure that the expiration date is set on all Secrets (Automated)' AS title,
  akv.subscription_id AS subscription_id,
  akvs.id AS resource_id,
  CASE
  WHEN (akvs.properties->'attributes'->>'enabled')::BOOL = true
  AND (akvs.properties->'attributes'->>'exp') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_keyvault_keyvault AS akv
  JOIN azure_keyvault_keyvault_secrets AS akvs ON
      akv._cq_id = akvs._cq_parent_id;
```


