# Table: shopify_price_rules

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.
## Relations

The following tables depend on shopify_price_rules:
  - [shopify_price_rule_discount_codes](shopify_price_rule_discount_codes.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|updated_at (Incremental Key)|Timestamp|
|value_type|String|
|value|String|
|customer_selection|String|
|target_type|String|
|target_selection|String|
|allocation_method|String|
|once_per_customer|Bool|
|usage_limit|Int|
|starts_at|Timestamp|
|ends_at|Timestamp|
|created_at|Timestamp|
|entitled_product_ids|JSON|
|entitled_variant_ids|JSON|
|entitled_collection_ids|JSON|
|entitled_country_ids|JSON|
|prerequisite_product_ids|JSON|
|prerequisite_variant_ids|JSON|
|prerequisite_collection_ids|JSON|
|customer_segment_prerequisite_ids|JSON|
|prerequisite_customer_ids|JSON|
|prerequisite_to_entitlement_quantity_ratio|JSON|
|prerequisite_to_entitlement_purchase|JSON|
|title|String|
|admin_graphql_api_id|String|