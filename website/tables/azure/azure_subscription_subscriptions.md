# Table: azure_subscription_subscriptions

This table shows data for Azure Subscription Subscriptions.

https://learn.microsoft.com/en-us/rest/api/resources/subscriptions/list?tabs=HTTP#subscription

The primary key for this table is **id**.

## Relations

The following tables depend on azure_subscription_subscriptions:
  - [azure_subscription_subscription_locations](azure_subscription_subscription_locations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|authorization_source|`utf8`|
|managed_by_tenants|`json`|
|subscription_policies|`json`|
|tags|`json`|
|display_name|`utf8`|
|id (PK)|`utf8`|
|state|`utf8`|
|subscription_id|`utf8`|
|tenant_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### External accounts with owner permissions should be removed from your subscription

```sql
SELECT
  'External accounts with owner permissions should be removed from your subscription'
    AS title,
  mc.subscription_id AS subscription_id,
  mc.id AS resource_id,
  CASE
  WHEN (properties->>'enableRBAC')::BOOL IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_containerservice_managed_clusters AS mc
  INNER JOIN azure_subscription_subscriptions AS sub ON
      sub.id = mc.subscription_id;
```

### Auto provisioning of the Log Analytics agent should be enabled on your subscription

```sql
SELECT
  'Auto provisioning of the Log Analytics agent should be enabled on your subscription'
    AS title,
  azure_subscription_subscriptions.id AS subscription_id,
  azure_security_auto_provisioning_settings._cq_id,
  CASE
  WHEN properties->>'autoProvision' IS DISTINCT FROM 'On' THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_security_auto_provisioning_settings
  RIGHT JOIN azure_subscription_subscriptions ON
      azure_security_auto_provisioning_settings.subscription_id
      = azure_subscription_subscriptions.id;
```

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

### Auditing on SQL server should be enabled

```sql
SELECT
  'Auditing on SQL server should be enabled' AS title,
  sub.id,
  sub.display_name AS subscription_name,
  CASE
  WHEN azure_sql_server_blob_auditing_policies._cq_parent_id
  = azure_sql_servers._cq_id
  AND sub.id = azure_sql_servers.subscription_id
  AND azure_sql_server_blob_auditing_policies.properties->>'state' = 'Disabled'
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  azure_sql_server_blob_auditing_policies,
  azure_sql_servers,
  azure_subscription_subscriptions AS sub;
```


