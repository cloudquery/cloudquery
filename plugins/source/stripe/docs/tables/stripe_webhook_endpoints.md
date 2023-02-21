# Table: stripe_webhook_endpoints

https://stripe.com/docs/api/webhook_endpoints

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|api_version|String|
|application|String|
|created|Timestamp|
|deleted|Bool|
|description|String|
|enabled_events|StringArray|
|livemode|Bool|
|metadata|JSON|
|object|String|
|secret|String|
|status|String|
|url|String|