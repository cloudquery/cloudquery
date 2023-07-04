# Table: stripe_apple_pay_domains

This table shows data for Stripe Apple Pay Domains.

https://stripe.com/docs/api

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|deleted|`bool`|
|domain_name|`utf8`|
|livemode|`bool`|
|object|`utf8`|