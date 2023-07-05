# Table: stripe_customers

This table shows data for Stripe Customers.

https://stripe.com/docs/api/customers

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|address|`json`|
|balance|`int64`|
|cash_balance|`json`|
|currency|`utf8`|
|default_source|`json`|
|deleted|`bool`|
|delinquent|`bool`|
|description|`utf8`|
|discount|`json`|
|email|`utf8`|
|invoice_credit_balance|`json`|
|invoice_prefix|`utf8`|
|invoice_settings|`json`|
|livemode|`bool`|
|metadata|`json`|
|name|`utf8`|
|next_invoice_sequence|`int64`|
|object|`utf8`|
|phone|`utf8`|
|preferred_locales|`list<item: utf8, nullable>`|
|shipping|`json`|
|sources|`json`|
|subscriptions|`json`|
|tax|`json`|
|tax_exempt|`utf8`|
|tax_ids|`json`|
|test_clock|`json`|