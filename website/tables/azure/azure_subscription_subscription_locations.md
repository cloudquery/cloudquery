# Table: azure_subscription_subscription_locations

This table shows data for Azure Subscription Subscription Locations.

https://learn.microsoft.com/en-us/rest/api/resources/subscriptions/list-locations?tabs=HTTP#location

The primary key for this table is **id**.

## Relations

This table depends on [azure_subscription_subscriptions](azure_subscription_subscriptions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|display_name|`utf8`|
|id (PK)|`utf8`|
|latitude|`utf8`|
|longitude|`utf8`|
|name|`utf8`|