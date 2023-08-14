# Table: stripe_country_specs

This table shows data for Stripe Country Specs.

https://stripe.com/docs/api/country_specs

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|default_currency|`utf8`|
|object|`utf8`|
|supported_bank_account_currencies|`json`|
|supported_payment_currencies|`list<item: utf8, nullable>`|
|supported_payment_methods|`list<item: utf8, nullable>`|
|supported_transfer_countries|`list<item: utf8, nullable>`|
|verification_fields|`json`|