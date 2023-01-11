# Table: stripe_transfers

https://stripe.com/docs/api/transfers

The primary key for this table is **id**.
It supports incremental syncs.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|amount|Int|
|amount_reversed|Int|
|balance_transaction|JSON|
|created|Int|
|currency|String|
|description|String|
|destination|JSON|
|destination_payment|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|reversals|JSON|
|reversed|Bool|
|source_transaction|JSON|
|source_type|String|
|transfer_group|String|