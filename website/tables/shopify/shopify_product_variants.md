# Table: shopify_product_variants

This table shows data for Shopify Product Variants.

The composite primary key for this table is (**product_id**, **id**).

## Relations

This table depends on [shopify_products](shopify_products).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|product_id (PK)|`int64`|
|id (PK)|`int64`|
|title|`utf8`|
|price|`utf8`|
|sku|`utf8`|
|position|`int64`|
|inventory_policy|`utf8`|
|compare_at_price|`utf8`|
|fulfillment_service|`utf8`|
|inventory_management|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|taxable|`bool`|
|barcode|`utf8`|
|grams|`int64`|
|weight|`float64`|
|weight_unit|`utf8`|
|inventory_item_id|`int64`|
|inventory_quantity|`int64`|
|old_inventory_quantity|`int64`|
|requires_shipping|`bool`|
|admin_graphql_api_id|`utf8`|