# Table: azure_subscription_tenants

This table shows data for Azure Subscription Tenants.

https://learn.microsoft.com/en-us/rest/api/resources/tenants/list?tabs=HTTP#tenantiddescription

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|country|`utf8`|
|country_code|`utf8`|
|default_domain|`utf8`|
|display_name|`utf8`|
|domains|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|tenant_branding_logo_url|`utf8`|
|tenant_category|`utf8`|
|tenant_id|`utf8`|
|tenant_type|`utf8`|