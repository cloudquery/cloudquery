# Table: azure_appservice_web_app_auth_settings

This table shows data for Azure App Service Web App Auth Settings.

https://learn.microsoft.com/en-us/rest/api/appservice/web-apps/get-auth-settings#siteauthsettings

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

### Ensure App Service Authentication is set on Azure App Service (Automated)

```sql
SELECT
  'Ensure App Service Authentication is set on Azure App Service (Automated)'
    AS title,
  awa.subscription_id AS subscription_id,
  awa.id AS resource_id,
  CASE
  WHEN (awaas.properties->>'enabled')::BOOL IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_appservice_web_apps AS awa
  LEFT JOIN azure_appservice_web_app_auth_settings AS awaas ON
      awa._cq_id = awaas._cq_parent_id;
```


