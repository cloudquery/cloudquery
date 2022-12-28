# Table: shopify_product_images

The composite primary key for this table is (**product_id**, **id**).

## Relations

This table depends on [shopify_products](shopify_products.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|product_id (PK)|Int|
|id (PK)|Int|
|position|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|alt|String|
|width|Int|
|height|Int|
|src|String|
|variant_ids|JSON|
|admin_graphql_api_id|String|