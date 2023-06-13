# Table: stripe_tax_codes

This table shows data for Stripe Tax Codes.

https://stripe.com/docs/api/tax_codes

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|description|`utf8`|
|name|`utf8`|
|object|`utf8`|