# Table: stripe_issuing_disputes

https://stripe.com/docs/api/issuing_disputes

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
|balance_transactions|JSON|
|currency|String|
|evidence|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|status|String|
|transaction|JSON|
|treasury|JSON|