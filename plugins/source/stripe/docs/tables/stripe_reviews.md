# Table: stripe_reviews

https://stripe.com/docs/api/reviews

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|billing_zip|String|
|charge|JSON|
|closed_reason|String|
|ip_address|String|
|ip_address_location|JSON|
|livemode|Bool|
|object|String|
|open|Bool|
|opened_reason|String|
|payment_intent|JSON|
|reason|String|
|session|JSON|