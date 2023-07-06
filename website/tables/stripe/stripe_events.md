# Table: stripe_events

This table shows data for Stripe Events.

https://stripe.com/docs/api/events

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|account|`utf8`|
|api_version|`utf8`|
|data|`json`|
|livemode|`bool`|
|object|`utf8`|
|pending_webhooks|`int64`|
|request|`json`|
|type|`utf8`|