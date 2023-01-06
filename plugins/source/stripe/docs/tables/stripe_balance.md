# Table: stripe_balance

https://stripe.com/docs/api/balance

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|available|JSON|
|connect_reserved|JSON|
|instant_available|JSON|
|issuing|JSON|
|livemode|Bool|
|object|String|
|pending|JSON|