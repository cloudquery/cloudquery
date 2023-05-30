# Table: stripe_balance

This table shows data for Stripe Balance.

https://stripe.com/docs/api/balance

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|available|`json`|
|connect_reserved|`json`|
|instant_available|`json`|
|issuing|`json`|
|livemode|`bool`|
|object|`utf8`|
|pending|`json`|