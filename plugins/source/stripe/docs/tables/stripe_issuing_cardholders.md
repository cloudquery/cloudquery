# Table: stripe_issuing_cardholders

https://stripe.com/docs/api/issuing_cardholders

The primary key for this table is **id**.
It supports incremental syncs.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|billing|JSON|
|company|JSON|
|created|Int|
|email|String|
|individual|JSON|
|livemode|Bool|
|metadata|JSON|
|name|String|
|object|String|
|phone_number|String|
|requirements|JSON|
|spending_controls|JSON|
|status|String|
|type|String|