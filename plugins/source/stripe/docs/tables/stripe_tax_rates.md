# Table: stripe_tax_rates

https://stripe.com/docs/api/tax_rates

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|active|Bool|
|country|String|
|created|Int|
|description|String|
|display_name|String|
|inclusive|Bool|
|jurisdiction|String|
|livemode|Bool|
|metadata|JSON|
|object|String|
|percentage|Float|
|state|String|
|tax_type|String|