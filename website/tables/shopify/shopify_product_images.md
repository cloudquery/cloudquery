# Table: shopify_product_images

This table shows data for Shopify Product Images.

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
|position|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|alt|`utf8`|
|width|`int64`|
|height|`int64`|
|src|`utf8`|
|variant_ids|`json`|
|admin_graphql_api_id|`utf8`|