# Table: stripe_issuing_cardholders

This table shows data for Stripe Issuing Cardholders.

https://stripe.com/docs/api/issuing/cardholders

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|billing|`json`|
|company|`json`|
|email|`utf8`|
|individual|`json`|
|livemode|`bool`|
|metadata|`json`|
|name|`utf8`|
|object|`utf8`|
|phone_number|`utf8`|
|requirements|`json`|
|spending_controls|`json`|
|status|`utf8`|
|type|`utf8`|