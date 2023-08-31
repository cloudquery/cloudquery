# Table: stripe_accounts

This table shows data for Stripe Accounts.

https://stripe.com/docs/api/accounts

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.
## Relations

The following tables depend on stripe_accounts:
  - [stripe_capabilities](stripe_capabilities)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|business_profile|`json`|
|business_type|`utf8`|
|capabilities|`json`|
|charges_enabled|`bool`|
|company|`json`|
|controller|`json`|
|country|`utf8`|
|default_currency|`utf8`|
|deleted|`bool`|
|details_submitted|`bool`|
|email|`utf8`|
|external_accounts|`json`|
|future_requirements|`json`|
|individual|`json`|
|metadata|`json`|
|object|`utf8`|
|payouts_enabled|`bool`|
|requirements|`json`|
|settings|`json`|
|tos_acceptance|`json`|
|type|`utf8`|