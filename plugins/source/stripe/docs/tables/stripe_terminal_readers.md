# Table: stripe_terminal_readers

https://stripe.com/docs/api/terminal_readers

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|action|JSON|
|deleted|Bool|
|device_sw_version|String|
|device_type|String|
|ip_address|String|
|label|String|
|livemode|Bool|
|location|JSON|
|metadata|JSON|
|object|String|
|serial_number|String|
|status|String|