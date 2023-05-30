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