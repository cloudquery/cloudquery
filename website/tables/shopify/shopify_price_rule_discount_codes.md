# Table: shopify_price_rule_discount_codes

This table shows data for Shopify Price Rule Discount Codes.

The composite primary key for this table is (**id**, **price_rule_id**).

## Relations

This table depends on [shopify_price_rules](shopify_price_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|price_rule_id (PK)|`int64`|
|code|`utf8`|
|usage_count|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|