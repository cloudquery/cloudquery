# Table: azure_subscription_subscriptions

This table shows data for Azure Subscription Subscriptions.

https://learn.microsoft.com/en-us/rest/api/resources/subscriptions/list?tabs=HTTP#subscription

The primary key for this table is **id**.

## Relations

The following tables depend on azure_subscription_subscriptions:
  - [azure_subscription_subscription_locations](azure_subscription_subscription_locations.md)

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