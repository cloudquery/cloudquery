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

### Ensure web app redirects all HTTP traffic to HTTPS in Azure App Service (Automated)

```sql
SELECT
  'Ensure web app redirects all HTTP traffic to HTTPS in Azure App Service (Automated)'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (properties->>'httpsOnly')::BOOL IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_appservice_web_apps;
```

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

### Ensure the web app has ''Client Certificates (Incoming client certificates)'' set to ''On'' (Automated)

```sql
SELECT
  e'Ensure the web app has \'Client Certificates (Incoming client certificates)\' set to \'On\' (Automated)'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN kind LIKE 'app%' AND (properties->>'clientCertEnabled')::BOOL IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_appservice_web_apps;
```

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

### Ensure that Register with Azure Active Directory is enabled on App Service (Automated)

```sql
SELECT
  'Ensure that Register with Azure Active Directory is enabled on App Service (Automated)'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (identity->>'principalId') IS NULL OR identity->>'principalId' = ''
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_appservice_web_apps;
```

### Ensure web app is using the latest version of TLS encryption (Automated)

```sql
SELECT
  'Ensure web app is using the latest version of TLS encryption (Automated)'
    AS title,
  subscription_id AS subscription_id,
  id AS resource_id,
  CASE
  WHEN (properties->'siteConfig'->'minTlsVersion') IS NULL
  OR properties->'siteConfig'->>'minTlsVersion' IS DISTINCT FROM '1.2'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  azure_appservice_web_apps;
```


