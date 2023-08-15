# Table: azure_security_assessments

This table shows data for Azure Security Assessments.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/assessments/list?tabs=HTTP#securityassessment

The primary key for this table is **id**.

## Relations

The following tables depend on azure_security_assessments:
  - [azure_security_sub_assessments](azure_security_sub_assessments)

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

### Deprecated accounts with owner permissions should be removed from your subscription

```sql
SELECT
  'Deprecated accounts with owner permissions should be removed from your subscription'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE WHEN a.code IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  azure_subscription_subscriptions AS s
  LEFT JOIN azure_security_assessments AS a ON
      s.id = '/subscriptions/' || a.subscription_id
      AND a.name = 'e52064aa-6853-e252-a11e-dffc675689c2'
      AND (
          a.code IS NOT DISTINCT FROM 'NotApplicable'
          OR a.code IS NOT DISTINCT FROM 'Healthy'
        );
```

### External accounts with owner permissions should be removed from your subscription

```sql
SELECT
  'External accounts with owner permissions should be removed from your subscription'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (a.properties->>'code') IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_subscription_subscriptions AS s
  LEFT JOIN azure_security_assessments AS a ON
      s.id = '/subscriptions/' || a.subscription_id
      AND a.name = 'c3b6ae71-f1f0-31b4-e6c1-d5951285d03d'
      AND (
          a.properties->>'code' IS NOT DISTINCT FROM 'NotApplicable'
          OR a.properties->>'code' IS NOT DISTINCT FROM 'Healthy'
        );
```

### MFA should be enabled on accounts with write permissions on your subscription

```sql
SELECT
  'MFA should be enabled on accounts with write permissions on your subscription'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN properties->'status'->>'code' IS DISTINCT FROM 'NotApplicable'
  AND properties->'status'->>'code' IS DISTINCT FROM 'Healthy'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_security_assessments
WHERE
  name = '57e98606-6b1e-6193-0e3d-fe621387c16b';
```

### MFA should be enabled on accounts with owner permissions on your subscription

```sql
SELECT
  'MFA should be enabled on accounts with owner permissions on your subscription'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN properties->'status'->>'code' IS DISTINCT FROM 'NotApplicable'
  AND properties->'status'->>'code' IS DISTINCT FROM 'Healthy'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_security_assessments
WHERE
  name = '94290b00-4d0c-d7b4-7cea-064a9554e681';
```

### MFA should be enabled on accounts with owner permissions on your subscription

```sql
SELECT
  'MFA should be enabled on accounts with owner permissions on your subscription'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN properties->'status'->>'code' IS DISTINCT FROM 'NotApplicable'
  AND properties->'status'->>'code' IS DISTINCT FROM 'Healthy'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_security_assessments
WHERE
  name = '151e82c5-5341-a74b-1eb0-bc38d2c84bb5';
```


