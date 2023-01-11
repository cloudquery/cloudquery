# Table: shopify_customers

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|updated_at (Incremental Key)|Timestamp|
|email|String|
|accepts_marketing|Bool|
|created_at|Timestamp|
|first_name|String|
|last_name|String|
|orders_count|Int|
|state|String|
|total_spent|String|
|last_order_id|Int|
|verified_email|Bool|
|tax_exempt|Bool|
|tags|String|
|last_order_name|String|
|currency|String|
|addresses|JSON|
|accepts_marketing_updated_at|Timestamp|
|marketing_opt_in_level|String|
|tax_exemptions|JSON|
|email_marketing_consent|JSON|
|admin_graphql_api_id|String|
|default_address|JSON|