# Table: stripe_issuing_transactions

https://stripe.com/docs/api/issuing_transactions

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|amount|Int|
|amount_details|JSON|
|authorization|JSON|
|balance_transaction|JSON|
|card|JSON|
|cardholder|JSON|
|currency|String|
|dispute|JSON|
|livemode|Bool|
|merchant_amount|Int|
|merchant_currency|String|
|merchant_data|JSON|
|metadata|JSON|
|object|String|
|purchase_details|JSON|
|treasury|JSON|
|type|String|
|wallet|String|