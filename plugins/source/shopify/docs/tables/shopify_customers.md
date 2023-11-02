# Table: shopify_customers

This table shows data for Shopify Customers.

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|updated_at (Incremental Key)|`timestamp[us, tz=UTC]`|
|email|`utf8`|
|accepts_marketing|`bool`|
|created_at|`timestamp[us, tz=UTC]`|
|first_name|`utf8`|
|last_name|`utf8`|
|orders_count|`int64`|
|state|`utf8`|
|total_spent|`utf8`|
|last_order_id|`int64`|
|verified_email|`bool`|
|tax_exempt|`bool`|
|tags|`utf8`|
|last_order_name|`utf8`|
|currency|`utf8`|
|addresses|`json`|
|accepts_marketing_updated_at|`timestamp[us, tz=UTC]`|
|marketing_opt_in_level|`utf8`|
|tax_exemptions|`json`|
|email_marketing_consent|`json`|
|admin_graphql_api_id|`utf8`|
|default_address|`json`|