# Table: azure_consumption_subscription_reservation_recommendations

This table shows data for Azure Consumption Subscription Reservation Recommendations.

https://learn.microsoft.com/en-us/rest/api/consumption/reservation-recommendations/list?tabs=HTTP#legacyreservationrecommendation

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|kind|`utf8`|
|etag|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|sku|`utf8`|
|tags|`json`|
|type|`utf8`|