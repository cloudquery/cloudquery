# Table: shopify_price_rules

This table shows data for Shopify Price Rules.

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.
## Relations

The following tables depend on shopify_price_rules:
  - [shopify_price_rule_discount_codes](shopify_price_rule_discount_codes)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|updated_at (Incremental Key)|`timestamp[us, tz=UTC]`|
|value_type|`utf8`|
|value|`utf8`|
|customer_selection|`utf8`|
|target_type|`utf8`|
|target_selection|`utf8`|
|allocation_method|`utf8`|
|once_per_customer|`bool`|
|usage_limit|`int64`|
|starts_at|`timestamp[us, tz=UTC]`|
|ends_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|entitled_product_ids|`json`|
|entitled_variant_ids|`json`|
|entitled_collection_ids|`json`|
|entitled_country_ids|`json`|
|prerequisite_product_ids|`json`|
|prerequisite_variant_ids|`json`|
|prerequisite_collection_ids|`json`|
|customer_segment_prerequisite_ids|`json`|
|prerequisite_customer_ids|`json`|
|prerequisite_to_entitlement_quantity_ratio|`json`|
|prerequisite_to_entitlement_purchase|`json`|
|title|`utf8`|
|admin_graphql_api_id|`utf8`|