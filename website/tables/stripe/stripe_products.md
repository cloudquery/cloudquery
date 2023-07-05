# Table: stripe_products

This table shows data for Stripe Products.

https://stripe.com/docs/api/products

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|active|`bool`|
|attributes|`list<item: utf8, nullable>`|
|caption|`utf8`|
|deactivate_on|`list<item: utf8, nullable>`|
|default_price|`json`|
|deleted|`bool`|
|description|`utf8`|
|images|`list<item: utf8, nullable>`|
|livemode|`bool`|
|metadata|`json`|
|name|`utf8`|
|object|`utf8`|
|package_dimensions|`json`|
|shippable|`bool`|
|statement_descriptor|`utf8`|
|tax_code|`json`|
|type|`utf8`|
|unit_label|`utf8`|
|updated|`int64`|
|url|`utf8`|