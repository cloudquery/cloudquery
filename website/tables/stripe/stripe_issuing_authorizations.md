# Table: stripe_issuing_authorizations

This table shows data for Stripe Issuing Authorizations.

https://stripe.com/docs/api/issuing/authorizations

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|amount|`int64`|
|amount_details|`json`|
|approved|`bool`|
|authorization_method|`utf8`|
|balance_transactions|`json`|
|card|`json`|
|cardholder|`json`|
|currency|`utf8`|
|livemode|`bool`|
|merchant_amount|`int64`|
|merchant_currency|`utf8`|
|merchant_data|`json`|
|metadata|`json`|
|network_data|`json`|
|object|`utf8`|
|pending_request|`json`|
|request_history|`json`|
|status|`utf8`|
|transactions|`json`|
|treasury|`json`|
|verification_data|`json`|
|wallet|`utf8`|