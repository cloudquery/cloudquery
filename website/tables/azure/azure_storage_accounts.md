# Table: azure_storage_accounts

This table shows data for Azure Storage Accounts.

https://learn.microsoft.com/en-us/rest/api/storagerp/storage-accounts/list?tabs=HTTP#storageaccount

The primary key for this table is **id**.

## Relations

The following tables depend on azure_storage_accounts:
  - [azure_storage_blob_services](azure_storage_blob_services)
  - [azure_storage_containers](azure_storage_containers)
  - [azure_storage_file_shares](azure_storage_file_shares)
  - [azure_storage_queue_services](azure_storage_queue_services)
  - [azure_storage_queues](azure_storage_queues)
  - [azure_storage_tables](azure_storage_tables)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|identity|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|sku|`json`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure default network access rule for Storage Accounts is set to deny

```sql
SELECT
  'Ensure default network access rule for Storage Accounts is set to deny'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN properties->'networkAcls'->>'defaultAction' = 'Allow' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_storage_accounts;
```

### Ensure storage for critical data are encrypted with Customer Managed Key

```sql
SELECT
  'Ensure storage for critical data are encrypted with Customer Managed Key'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN properties->'encryption'->>'keySource' = 'Microsoft.Keyvault'
  AND properties->'encryption'->'keyvaultproperties' IS DISTINCT FROM NULL
  THEN 'pass'
  ELSE 'fail'
  END
    AS status
FROM
  azure_storage_accounts;
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

### Ensure that ''Public access level'' is set to Private for blob containers

```sql
SELECT
  e'Ensure that \'Public access level\' is set to Private for blob containers'
    AS title,
  azsc.subscription_id AS subscription_id,
  azsc.id AS resrouce_id,
  CASE
  WHEN azsc.properties->>'publicAccess' = 'None'
  AND NOT (asa.properties->>'allowBlobPublicAccess')::BOOL
  THEN 'pass'
  ELSE 'fail'
  END
    AS status
FROM
  azure_storage_containers AS azsc
  JOIN azure_storage_accounts AS asa ON azsc._cq_parent_id = asa._cq_id;
```


