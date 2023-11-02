# Table: azure_security_auto_provisioning_settings

This table shows data for Azure Security Auto Provisioning Settings.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/auto-provisioning-settings/list?tabs=HTTP#autoprovisioningsetting

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

### Ensure that "Automatic provisioning of monitoring agent" is set to "On" (Automated)

```sql
SELECT
  'Ensure that "Automatic provisioning of monitoring agent" is set to "On" (Automated)'
    AS title,
  subscription_id,
  id,
  CASE
  WHEN properties->>'autoProvision' = 'On' THEN 'pass'
  ELSE 'fail'
  END
FROM
  azure_security_auto_provisioning_settings AS asaps
WHERE
  name = 'default';
```


