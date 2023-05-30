# Table: heroku_invoices

This table shows data for Heroku Invoices.

https://devcenter.heroku.com/articles/platform-api-reference#invoice

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|charges_total|`float64`|
|created_at|`timestamp[us, tz=UTC]`|
|credits_total|`float64`|
|number|`int64`|
|period_end|`utf8`|
|period_start|`utf8`|
|state|`int64`|
|total|`float64`|
|updated_at|`timestamp[us, tz=UTC]`|