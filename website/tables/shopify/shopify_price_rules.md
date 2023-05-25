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
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|int64|
|updated_at (Incremental Key)|timestamp[us, tz=UTC]|
|value_type|extension_type<storage=binary>|
|value|extension_type<storage=binary>|
|customer_selection|extension_type<storage=binary>|
|target_type|extension_type<storage=binary>|
|target_selection|extension_type<storage=binary>|
|allocation_method|extension_type<storage=binary>|
|once_per_customer|extension_type<storage=binary>|
|usage_limit|extension_type<storage=binary>|
|starts_at|extension_type<storage=binary>|
|ends_at|extension_type<storage=binary>|
|created_at|extension_type<storage=binary>|
|entitled_product_ids|extension_type<storage=binary>|
|entitled_variant_ids|extension_type<storage=binary>|
|entitled_collection_ids|extension_type<storage=binary>|
|entitled_country_ids|extension_type<storage=binary>|
|prerequisite_product_ids|extension_type<storage=binary>|
|prerequisite_variant_ids|extension_type<storage=binary>|
|prerequisite_collection_ids|extension_type<storage=binary>|
|customer_segment_prerequisite_ids|extension_type<storage=binary>|
|prerequisite_customer_ids|extension_type<storage=binary>|
|prerequisite_to_entitlement_quantity_ratio|extension_type<storage=binary>|
|prerequisite_to_entitlement_purchase|extension_type<storage=binary>|
|title|extension_type<storage=binary>|
|admin_graphql_api_id|extension_type<storage=binary>|