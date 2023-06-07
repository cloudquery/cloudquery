# Table: azure_appservice_web_apps

This table shows data for Azure App Service Web Apps.

https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/list#site

The primary key for this table is **id**.

## Relations

The following tables depend on azure_appservice_web_apps:
  - [azure_appservice_web_app_auth_settings](azure_appservice_web_app_auth_settings)
  - [azure_appservice_web_app_configurations](azure_appservice_web_app_configurations)
  - [azure_appservice_web_app_vnet_connections](azure_appservice_web_app_vnet_connections)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|identity|`json`|
|kind|`utf8`|
|properties|`json`|
|tags|`json`|
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


