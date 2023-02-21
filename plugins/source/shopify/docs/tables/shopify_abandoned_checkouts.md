# Table: shopify_abandoned_checkouts

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|updated_at (Incremental Key)|Timestamp|
|token|String|
|cart_token|String|
|email|String|
|gateway|String|
|buyer_accepts_marketing|Bool|
|created_at|Timestamp|
|landing_site|String|
|note_attributes|JSON|
|referring_site|String|
|shipping_lines|JSON|
|taxes_included|Bool|
|total_weight|Int|
|currency|String|
|completed_at|Timestamp|
|closed_at|Timestamp|
|user_id|Int|
|customer_locale|String|
|line_items|JSON|
|name|String|
|abandoned_checkout_url|String|
|discount_codes|JSON|
|tax_lines|JSON|
|source_name|String|
|presentment_currency|String|
|buyer_accepts_sms_marketing|Bool|
|total_discounts|String|
|total_line_items_price|String|
|total_price|String|
|total_tax|String|
|subtotal_price|String|
|customer|JSON|