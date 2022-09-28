# Table: heroku_invoices
https://devcenter.heroku.com/articles/platform-api-reference#invoice-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|charges_total|Float|
|created_at|Timestamp|
|credits_total|Float|
|id (PK)|String|
|number|Int|
|period_end|String|
|period_start|String|
|state|Int|
|total|Float|
|updated_at|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|