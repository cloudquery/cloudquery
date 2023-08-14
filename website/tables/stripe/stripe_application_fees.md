# Table: stripe_application_fees

This table shows data for Stripe Application Fees.

https://stripe.com/docs/api/application_fees

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.
## Relations

The following tables depend on stripe_application_fees:
  - [stripe_fee_refunds](stripe_fee_refunds)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|account|`json`|
|amount|`int64`|
|amount_refunded|`int64`|
|application|`json`|
|balance_transaction|`json`|
|charge|`json`|
|currency|`utf8`|
|livemode|`bool`|
|object|`utf8`|
|originating_transaction|`json`|
|refunded|`bool`|
|refunds|`json`|