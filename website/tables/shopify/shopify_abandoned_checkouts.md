# Table: shopify_abandoned_checkouts

This table shows data for Shopify Abandoned Checkouts.

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|int64|
|updated_at (Incremental Key)|timestamp[us, tz=UTC]|
|token|extension_type<storage=binary>|
|cart_token|extension_type<storage=binary>|
|email|extension_type<storage=binary>|
|gateway|extension_type<storage=binary>|
|buyer_accepts_marketing|extension_type<storage=binary>|
|created_at|extension_type<storage=binary>|
|landing_site|extension_type<storage=binary>|
|note_attributes|extension_type<storage=binary>|
|referring_site|extension_type<storage=binary>|
|shipping_lines|extension_type<storage=binary>|
|taxes_included|extension_type<storage=binary>|
|total_weight|extension_type<storage=binary>|
|currency|extension_type<storage=binary>|
|completed_at|extension_type<storage=binary>|
|closed_at|extension_type<storage=binary>|
|user_id|extension_type<storage=binary>|
|customer_locale|extension_type<storage=binary>|
|line_items|extension_type<storage=binary>|
|name|extension_type<storage=binary>|
|abandoned_checkout_url|extension_type<storage=binary>|
|discount_codes|extension_type<storage=binary>|
|tax_lines|extension_type<storage=binary>|
|source_name|extension_type<storage=binary>|
|presentment_currency|extension_type<storage=binary>|
|buyer_accepts_sms_marketing|extension_type<storage=binary>|
|total_discounts|extension_type<storage=binary>|
|total_line_items_price|extension_type<storage=binary>|
|total_price|extension_type<storage=binary>|
|total_tax|extension_type<storage=binary>|
|subtotal_price|extension_type<storage=binary>|
|customer|extension_type<storage=binary>|