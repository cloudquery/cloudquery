# Table: stripe_products

https://stripe.com/docs/api/products

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
|active|Bool|
|attributes|StringArray|
|caption|String|
|deactivate_on|StringArray|
|default_price|JSON|
|deleted|Bool|
|description|String|
|images|StringArray|
|livemode|Bool|
|metadata|JSON|
|name|String|
|object|String|
|package_dimensions|JSON|
|shippable|Bool|
|statement_descriptor|String|
|tax_code|JSON|
|type|String|
|unit_label|String|
|updated|Int|
|url|String|