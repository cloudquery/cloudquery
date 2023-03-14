# Table: heroku_team_invoices

This table shows data for Heroku Team Invoices.

https://devcenter.heroku.com/articles/platform-api-reference#team-invoice

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|addons_total|Int|
|charges_total|Int|
|created_at|Timestamp|
|credits_total|Int|
|database_total|Int|
|dyno_units|Float|
|number|Int|
|payment_status|String|
|period_end|String|
|period_start|String|
|platform_total|Int|
|state|Int|
|total|Int|
|updated_at|Timestamp|
|weighted_dyno_hours|Float|