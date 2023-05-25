# Table: shopify_product_images

This table shows data for Shopify Product Images.

The composite primary key for this table is (**product_id**, **id**).

## Relations

This table depends on [shopify_products](shopify_products).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|product_id (PK)|int64|
|id (PK)|int64|
|position|extension_type<storage=binary>|
|created_at|extension_type<storage=binary>|
|updated_at|extension_type<storage=binary>|
|alt|extension_type<storage=binary>|
|width|extension_type<storage=binary>|
|height|extension_type<storage=binary>|
|src|extension_type<storage=binary>|
|variant_ids|extension_type<storage=binary>|
|admin_graphql_api_id|extension_type<storage=binary>|