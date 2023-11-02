# Table: fastly_account_events

This table shows data for Fastly Account Events.

https://developer.fastly.com/reference/api/account/events/

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|ip|`utf8`|
|admin|`bool`|
|created_at|`timestamp[us, tz=UTC]`|
|customer_id|`utf8`|
|description|`utf8`|
|event_type|`utf8`|
|metadata|`json`|
|service_id|`utf8`|
|user_id|`utf8`|