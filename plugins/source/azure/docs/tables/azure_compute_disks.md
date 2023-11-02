# Table: azure_compute_disks

This table shows data for Azure Compute Disks.

https://learn.microsoft.com/en-us/rest/api/compute/disks/list?tabs=HTTP#disk

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|managed_by|`utf8`|
|managed_by_extended|`list<item: utf8, nullable>`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that ''OS and Data'' disks are encrypted with CMK (Automated)

```sql
INSERT
INTO
  azure_policy_results
    (
      execution_time,
      framework,
      check_id,
      title,
      subscription_id,
      resource_id,
      status
    )
SELECT
  e'Ensure that \'OS and Data\' disks are encrypted with CMK (Automated)'
    AS title,
  v.subscription_id AS subscription_id,
  v.id AS resource_id,
  CASE
  WHEN d.properties->'encryption'->>'type' NOT LIKE '%CustomerKey%' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_compute_virtual_machines AS v
  JOIN azure_compute_disks AS d ON
      lower(v.id) = lower(d.properties->>'managedBy');
```

### Ensure that ''Unattached disks'' are encrypted with CMK (Automated)

```sql
INSERT
INTO
  azure_policy_results
    (
      execution_time,
      framework,
      check_id,
      title,
      subscription_id,
      resource_id,
      status
    )
SELECT
  e'Ensure that \'Unattached disks\' are encrypted with CMK (Automated)'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN properties->'encryption'->>'type' NOT LIKE '%CustomerKey%' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_compute_disks
WHERE
  properties->>'diskState' = 'Unattached';
```

### Ensure that VHD''s are encrypted (Manual)

```sql
SELECT
  e'Ensure that VHD\'s are encrypted (Manual)' AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (properties->'encryptionSettingsCollection'->>'enabled')::BOOL
  IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_compute_disks;
```


