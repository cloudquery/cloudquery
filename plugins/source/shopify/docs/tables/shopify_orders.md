# Table: shopify_orders

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|admin_graphql_api_id|String|
|app_id|Int|
|browser_ip|String|
|buyer_accepts_marketing|Bool|
|cancelled_at|Timestamp|
|cart_token|String|
|checkout_id|Int|
|checkout_token|String|
|closed_at|Timestamp|
|confirmed|Bool|
|contact_email|String|
|created_at|Timestamp|
|currency|String|
|current_subtotal_price|String|
|current_total_discounts|String|
|current_total_price|String|
|current_total_tax|String|
|customer_locale|String|
|discount_codes|JSON|
|email|String|
|estimated_taxes|Bool|
|financial_status|String|
|gateway|String|
|landing_site|String|
|name|String|
|note_attributes|JSON|
|number|Int|
|order_number|Int|
|order_status_url|String|
|payment_gateway_names|StringArray|
|phone|String|
|presentment_currency|String|
|processed_at|Timestamp|
|processing_method|String|
|referring_site|String|
|source_name|String|
|subtotal_price|String|
|tags|String|
|tax_lines|JSON|
|taxes_included|Bool|
|test|Bool|
|token|String|
|total_discounts|String|
|total_line_items_price|String|
|total_outstanding|String|
|total_price|String|
|total_price_usd|String|
|total_tax|String|
|total_tip_received|String|
|total_weight|Int|
|updated_at|Timestamp|
|user_id|Int|
|customer|JSON|
|discount_applications|JSON|
|fulfillments|JSON|
|line_items|JSON|
|refunds|JSON|
|shipping_lines|JSON|