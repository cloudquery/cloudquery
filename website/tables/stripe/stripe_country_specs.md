# Table: stripe_country_specs

https://stripe.com/docs/api/country_specs

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|default_currency|String|
|object|String|
|supported_bank_account_currencies|JSON|
|supported_payment_currencies|StringArray|
|supported_payment_methods|StringArray|
|supported_transfer_countries|StringArray|
|verification_fields|JSON|