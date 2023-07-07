# Table: stripe_tax_rates

This table shows data for Stripe Tax Rates.

https://stripe.com/docs/api/tax_rates

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|active|`bool`|
|country|`utf8`|
|description|`utf8`|
|display_name|`utf8`|
|inclusive|`bool`|
|jurisdiction|`utf8`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|percentage|`float64`|
|state|`utf8`|
|tax_type|`utf8`|