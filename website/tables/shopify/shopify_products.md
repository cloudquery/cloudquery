# Table: shopify_products

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.
## Relations

The following tables depend on shopify_products:
  - [shopify_product_images](shopify_product_images)
  - [shopify_product_variants](shopify_product_variants)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|updated_at (Incremental Key)|Timestamp|
|title|String|
|body_html|String|
|vendor|String|
|product_type|String|
|created_at|Timestamp|
|handle|String|
|published_at|Timestamp|
|template_suffix|String|
|status|String|
|published_scope|String|
|tags|StringArray|
|image|JSON|
|admin_graphql_api_id|String|