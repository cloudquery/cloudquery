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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|authorization_source|String|
|subscription_policies|JSON|
|display_name|String|
|id (PK)|String|
|state|String|
|subscription_id|String|