# Table: azure_keyvault_keyvault_managed_hsms

This table shows data for Azure Key Vault Key Vault Managed Hsms.

https://learn.microsoft.com/en-us/rest/api/keyvault/managedhsm/managed-hsms/list-by-subscription?tabs=HTTP#managedhsm

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
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


