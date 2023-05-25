# Table: shopify_product_variants

This table shows data for Shopify Product Variants.

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
|title|extension_type<storage=binary>|
|price|extension_type<storage=binary>|
|sku|extension_type<storage=binary>|
|position|extension_type<storage=binary>|
|inventory_policy|extension_type<storage=binary>|
|compare_at_price|extension_type<storage=binary>|
|fulfillment_service|extension_type<storage=binary>|
|inventory_management|extension_type<storage=binary>|
|created_at|extension_type<storage=binary>|
|updated_at|extension_type<storage=binary>|
|taxable|extension_type<storage=binary>|
|barcode|extension_type<storage=binary>|
|grams|extension_type<storage=binary>|
|weight|extension_type<storage=binary>|
|weight_unit|extension_type<storage=binary>|
|inventory_item_id|extension_type<storage=binary>|
|inventory_quantity|extension_type<storage=binary>|
|old_inventory_quantity|extension_type<storage=binary>|
|requires_shipping|extension_type<storage=binary>|
|admin_graphql_api_id|extension_type<storage=binary>|