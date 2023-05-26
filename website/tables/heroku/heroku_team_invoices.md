# Table: heroku_team_invoices

This table shows data for Heroku Team Invoices.

https://devcenter.heroku.com/articles/platform-api-reference#team-invoice

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|addons_total|int64|
|charges_total|int64|
|created_at|timestamp[us, tz=UTC]|
|credits_total|int64|
|database_total|int64|
|dyno_units|float64|
|number|int64|
|payment_status|utf8|
|period_end|utf8|
|period_start|utf8|
|platform_total|int64|
|state|int64|
|total|int64|
|updated_at|timestamp[us, tz=UTC]|
|weighted_dyno_hours|float64|