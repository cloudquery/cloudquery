# Table: shopify_price_rule_discount_codes

The composite primary key for this table is (**id**, **price_rule_id**).

## Relations

This table depends on [shopify_price_rules](shopify_price_rules.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|price_rule_id (PK)|Int|
|code|String|
|usage_count|Int|
|created_at|Timestamp|
|updated_at|Timestamp|