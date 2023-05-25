# Table: shopify_customers

This table shows data for Shopify Customers.

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|int64|
|updated_at (Incremental Key)|timestamp[us, tz=UTC]|
|email|extension_type<storage=binary>|
|accepts_marketing|extension_type<storage=binary>|
|created_at|extension_type<storage=binary>|
|first_name|extension_type<storage=binary>|
|last_name|extension_type<storage=binary>|
|orders_count|extension_type<storage=binary>|
|state|extension_type<storage=binary>|
|total_spent|extension_type<storage=binary>|
|last_order_id|extension_type<storage=binary>|
|verified_email|extension_type<storage=binary>|
|tax_exempt|extension_type<storage=binary>|
|tags|extension_type<storage=binary>|
|last_order_name|extension_type<storage=binary>|
|currency|extension_type<storage=binary>|
|addresses|extension_type<storage=binary>|
|accepts_marketing_updated_at|extension_type<storage=binary>|
|marketing_opt_in_level|extension_type<storage=binary>|
|tax_exemptions|extension_type<storage=binary>|
|email_marketing_consent|extension_type<storage=binary>|
|admin_graphql_api_id|extension_type<storage=binary>|
|default_address|extension_type<storage=binary>|