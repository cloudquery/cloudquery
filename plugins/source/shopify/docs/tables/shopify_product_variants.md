# Table: shopify_product_variants

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
|title|String|
|price|String|
|sku|String|
|position|Int|
|inventory_policy|String|
|compare_at_price|String|
|fulfillment_service|String|
|inventory_management|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|taxable|Bool|
|barcode|String|
|grams|Int|
|weight|Float|
|weight_unit|String|
|inventory_item_id|Int|
|inventory_quantity|Int|
|old_inventory_quantity|Int|
|requires_shipping|Bool|
|admin_graphql_api_id|String|