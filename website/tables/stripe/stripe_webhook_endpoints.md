# Table: stripe_webhook_endpoints

This table shows data for Stripe Webhook Endpoints.

https://stripe.com/docs/api/webhook_endpoints

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|api_version|`utf8`|
|application|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|deleted|`bool`|
|description|`utf8`|
|enabled_events|`list<item: utf8, nullable>`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|secret|`utf8`|
|status|`utf8`|
|url|`utf8`|