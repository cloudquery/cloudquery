# Table: azure_resources_providers

This table shows data for Azure Resources Providers.

https://docs.microsoft.com/en-us/rest/api/resources/providers/list

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|namespace|`utf8`|
|provider_authorization_consent_state|`utf8`|
|id (PK)|`utf8`|
|registration_policy|`utf8`|
|registration_state|`utf8`|
|resource_types|`json`|