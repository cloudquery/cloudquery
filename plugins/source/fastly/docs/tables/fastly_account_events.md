# Table: fastly_account_events

https://developer.fastly.com/reference/api/account/events/

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|ip|String|
|admin|Bool|
|created_at|Timestamp|
|customer_id|String|
|description|String|
|event_type|String|
|metadata|JSON|
|service_id|String|
|user_id|String|