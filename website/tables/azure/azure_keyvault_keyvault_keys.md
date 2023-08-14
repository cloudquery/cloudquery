# Table: azure_keyvault_keyvault_keys

This table shows data for Azure Key Vault Key Vault Keys.

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault@v1.0.0#Key

The primary key for this table is **id**.

## Relations

This table depends on [azure_keyvault_keyvault](azure_keyvault_keyvault).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|tags|`json`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


