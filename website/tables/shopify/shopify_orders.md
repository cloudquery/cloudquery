# Table: shopify_orders

This table shows data for Shopify Orders.

The primary key for this table is **id**.
It supports incremental syncs based on the (**created_at**, **updated_at**) columns.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|int64|
|created_at (Incremental Key)|timestamp[us, tz=UTC]|
|updated_at (Incremental Key)|timestamp[us, tz=UTC]|
|admin_graphql_api_id|extension_type<storage=binary>|
|app_id|extension_type<storage=binary>|
|browser_ip|extension_type<storage=binary>|
|buyer_accepts_marketing|extension_type<storage=binary>|
|cancelled_at|extension_type<storage=binary>|
|cart_token|extension_type<storage=binary>|
|checkout_id|extension_type<storage=binary>|
|checkout_token|extension_type<storage=binary>|
|closed_at|extension_type<storage=binary>|
|confirmed|extension_type<storage=binary>|
|contact_email|extension_type<storage=binary>|
|currency|extension_type<storage=binary>|
|current_subtotal_price|extension_type<storage=binary>|
|current_total_discounts|extension_type<storage=binary>|
|current_total_price|extension_type<storage=binary>|
|current_total_tax|extension_type<storage=binary>|
|customer_locale|extension_type<storage=binary>|
|discount_codes|extension_type<storage=binary>|
|email|extension_type<storage=binary>|
|estimated_taxes|extension_type<storage=binary>|
|financial_status|extension_type<storage=binary>|
|gateway|extension_type<storage=binary>|
|landing_site|extension_type<storage=binary>|
|name|extension_type<storage=binary>|
|note_attributes|extension_type<storage=binary>|
|number|extension_type<storage=binary>|
|order_number|extension_type<storage=binary>|
|order_status_url|extension_type<storage=binary>|
|payment_gateway_names|extension_type<storage=binary>|
|phone|extension_type<storage=binary>|
|presentment_currency|extension_type<storage=binary>|
|processed_at|extension_type<storage=binary>|
|processing_method|extension_type<storage=binary>|
|referring_site|extension_type<storage=binary>|
|source_name|extension_type<storage=binary>|
|subtotal_price|extension_type<storage=binary>|
|tags|extension_type<storage=binary>|
|tax_lines|extension_type<storage=binary>|
|taxes_included|extension_type<storage=binary>|
|test|extension_type<storage=binary>|
|token|extension_type<storage=binary>|
|total_discounts|extension_type<storage=binary>|
|total_line_items_price|extension_type<storage=binary>|
|total_outstanding|extension_type<storage=binary>|
|total_price|extension_type<storage=binary>|
|total_price_usd|extension_type<storage=binary>|
|total_tax|extension_type<storage=binary>|
|total_tip_received|extension_type<storage=binary>|
|total_weight|extension_type<storage=binary>|
|user_id|extension_type<storage=binary>|
|customer|extension_type<storage=binary>|
|discount_applications|extension_type<storage=binary>|
|fulfillments|extension_type<storage=binary>|
|line_items|extension_type<storage=binary>|
|refunds|extension_type<storage=binary>|
|shipping_lines|extension_type<storage=binary>|