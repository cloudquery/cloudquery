# Table: stripe_reviews

This table shows data for Stripe Reviews.

https://stripe.com/docs/api/radar/reviews

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|billing_zip|`utf8`|
|charge|`json`|
|closed_reason|`utf8`|
|ip_address|`utf8`|
|ip_address_location|`json`|
|livemode|`bool`|
|object|`utf8`|
|open|`bool`|
|opened_reason|`utf8`|
|payment_intent|`json`|
|reason|`utf8`|
|session|`json`|