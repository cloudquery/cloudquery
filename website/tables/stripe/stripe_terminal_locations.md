# Table: stripe_terminal_locations

https://stripe.com/docs/api/terminal_locations

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|address|JSON|
|configuration_overrides|String|
|deleted|Bool|
|display_name|String|
|livemode|Bool|
|metadata|JSON|
|object|String|