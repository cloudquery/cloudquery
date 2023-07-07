# Table: azure_storage_containers

This table shows data for Azure Storage Containers.

https://learn.microsoft.com/en-us/rest/api/storagerp/blob-containers/list?tabs=HTTP#listcontaineritem

The primary key for this table is **id**.

## Relations

This table depends on [azure_storage_accounts](azure_storage_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|etag|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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

### Ensure the storage container storing the activity logs is not publicly accessible

```sql
SELECT
  'Ensure the storage container storing the activity logs is not publicly accessible'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN properties->>'publicAccess' = 'None' THEN 'pass'
  ELSE 'fail'
  END
    AS status
FROM
  azure_storage_containers
WHERE
  name = 'insights-activity-logs';
```


