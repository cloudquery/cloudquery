# Table: stripe_capabilities

This table shows data for Stripe Capabilities.

https://stripe.com/docs/api/capabilities

The primary key for this table is **id**.

## Relations

This table depends on [stripe_accounts](stripe_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|account|`json`|
|future_requirements|`json`|
|object|`utf8`|
|requested|`bool`|
|requested_at|`int64`|
|requirements|`json`|
|status|`utf8`|