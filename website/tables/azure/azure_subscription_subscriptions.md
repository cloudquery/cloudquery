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
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|authorization_source|`utf8`|
|subscription_policies|`json`|
|display_name|`utf8`|
|id (PK)|`utf8`|
|state|`utf8`|
|subscription_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


