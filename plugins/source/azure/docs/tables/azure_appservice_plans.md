# Table: azure_appservice_plans

This table shows data for Azure App Service Plans.

https://learn.microsoft.com/en-us/rest/api/appservice/app-service-environments/list-app-service-plans?tabs=HTTP#appserviceplan

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|extended_location|`json`|
|kind|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|