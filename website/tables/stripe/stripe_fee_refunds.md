# Table: stripe_fee_refunds

This table shows data for Stripe Fee Refunds.

https://stripe.com/docs/api/fee_refunds

The primary key for this table is **id**.

## Relations

This table depends on [stripe_application_fees](stripe_application_fees).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|amount|`int64`|
|balance_transaction|`json`|
|created|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|fee|`json`|
|metadata|`json`|
|object|`utf8`|