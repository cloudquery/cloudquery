# Table: stripe_issuing_cards

This table shows data for Stripe Issuing Cards.

https://stripe.com/docs/api/issuing/cards

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|brand|`utf8`|
|cancellation_reason|`utf8`|
|cardholder|`json`|
|currency|`utf8`|
|cvc|`utf8`|
|exp_month|`int64`|
|exp_year|`int64`|
|financial_account|`utf8`|
|last4|`utf8`|
|livemode|`bool`|
|metadata|`json`|
|number|`utf8`|
|object|`utf8`|
|replaced_by|`json`|
|replacement_for|`json`|
|replacement_reason|`utf8`|
|shipping|`json`|
|spending_controls|`json`|
|status|`utf8`|
|type|`utf8`|
|wallets|`json`|