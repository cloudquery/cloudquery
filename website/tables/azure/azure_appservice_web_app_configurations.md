# Table: azure_appservice_web_app_configurations

This table shows data for Azure App Service Web App Configurations.

https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list-configurations#siteconfigresource

The primary key for this table is **id**.

## Relations

This table depends on [azure_appservice_web_apps](azure_appservice_web_apps).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|kind|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure FTP deployments are disabled (Automated)

```sql
SELECT
  'Ensure FTP deployments are disabled (Automated)' AS title,
  aawac.subscription_id AS subscription_id,
  aawac.id AS resource_id,
  CASE
  WHEN aawac.properties->>'ftpsState' = 'AllAllowed' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_appservice_web_apps AS aawa
  JOIN azure_appservice_web_app_configurations AS aawac ON
      aawa._cq_id = aawac._cq_parent_id;
```


