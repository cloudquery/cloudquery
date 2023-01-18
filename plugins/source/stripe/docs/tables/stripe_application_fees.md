# Table: stripe_application_fees

https://stripe.com/docs/api/application_fees

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.
## Relations

The following tables depend on stripe_application_fees:
  - [stripe_fee_refunds](stripe_fee_refunds.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|account|JSON|
|amount|Int|
|amount_refunded|Int|
|application|JSON|
|balance_transaction|JSON|
|charge|JSON|
|currency|String|
|livemode|Bool|
|object|String|
|originating_transaction|JSON|
|refunded|Bool|
|refunds|JSON|