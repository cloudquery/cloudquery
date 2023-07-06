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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|updated_at (Incremental Key)|`timestamp[us, tz=UTC]`|
|title|`utf8`|
|body_html|`utf8`|
|vendor|`utf8`|
|product_type|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|handle|`utf8`|
|published_at|`timestamp[us, tz=UTC]`|
|template_suffix|`utf8`|
|status|`utf8`|
|published_scope|`utf8`|
|tags|`list<item: utf8, nullable>`|
|image|`json`|
|admin_graphql_api_id|`utf8`|