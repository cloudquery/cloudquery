# Table: stripe_issuing_transactions

This table shows data for Stripe Issuing Transactions.

https://stripe.com/docs/api/issuing/transactions

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
|authorization|`json`|
|balance_transaction|`json`|
|card|`json`|
|cardholder|`json`|
|currency|`utf8`|
|dispute|`json`|
|livemode|`bool`|
|merchant_amount|`int64`|
|merchant_currency|`utf8`|
|merchant_data|`json`|
|metadata|`json`|
|object|`utf8`|
|purchase_details|`json`|
|treasury|`json`|
|type|`utf8`|
|wallet|`utf8`|