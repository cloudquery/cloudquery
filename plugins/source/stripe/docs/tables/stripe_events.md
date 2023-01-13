# Table: stripe_events

https://stripe.com/docs/api/events

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
|account|String|
|api_version|String|
|data|JSON|
|livemode|Bool|
|object|String|
|pending_webhooks|Int|
|request|JSON|
|type|String|