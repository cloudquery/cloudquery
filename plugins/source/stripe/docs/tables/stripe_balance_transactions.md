# Table: stripe_balance_transactions

https://stripe.com/docs/api/balance_transactions

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|amount|Int|
|available_on|Int|
|created|Int|
|currency|String|
|description|String|
|exchange_rate|Float|
|fee|Int|
|fee_details|JSON|
|net|Int|
|object|String|
|reporting_category|String|
|source|JSON|
|status|String|
|type|String|