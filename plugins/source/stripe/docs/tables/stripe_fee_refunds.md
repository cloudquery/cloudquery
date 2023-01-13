# Table: stripe_fee_refunds

https://stripe.com/docs/api/fee_refunds

The primary key for this table is **id**.

## Relations

This table depends on [stripe_application_fees](stripe_application_fees.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|amount|Int|
|balance_transaction|JSON|
|created|Timestamp|
|currency|String|
|fee|JSON|
|metadata|JSON|
|object|String|