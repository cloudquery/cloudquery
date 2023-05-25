# Table: shopify_products

This table shows data for Shopify Products.

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.
## Relations

The following tables depend on shopify_products:
  - [shopify_product_images](shopify_product_images)
  - [shopify_product_variants](shopify_product_variants)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|int64|
|updated_at (Incremental Key)|timestamp[us, tz=UTC]|
|title|extension_type<storage=binary>|
|body_html|extension_type<storage=binary>|
|vendor|extension_type<storage=binary>|
|product_type|extension_type<storage=binary>|
|created_at|extension_type<storage=binary>|
|handle|extension_type<storage=binary>|
|published_at|extension_type<storage=binary>|
|template_suffix|extension_type<storage=binary>|
|status|extension_type<storage=binary>|
|published_scope|extension_type<storage=binary>|
|tags|extension_type<storage=binary>|
|image|extension_type<storage=binary>|
|admin_graphql_api_id|extension_type<storage=binary>|