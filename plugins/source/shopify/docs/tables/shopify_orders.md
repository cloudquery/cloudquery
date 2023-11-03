# Table: shopify_orders

This table shows data for Shopify Orders.

The primary key for this table is **id**.
It supports incremental syncs based on the (**created_at**, **updated_at**) columns.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|created_at (Incremental Key)|`timestamp[us, tz=UTC]`|
|updated_at (Incremental Key)|`timestamp[us, tz=UTC]`|
|admin_graphql_api_id|`utf8`|
|app_id|`int64`|
|browser_ip|`utf8`|
|buyer_accepts_marketing|`bool`|
|cancelled_at|`timestamp[us, tz=UTC]`|
|cart_token|`utf8`|
|checkout_id|`int64`|
|checkout_token|`utf8`|
|closed_at|`timestamp[us, tz=UTC]`|
|confirmed|`bool`|
|contact_email|`utf8`|
|currency|`utf8`|
|current_subtotal_price|`utf8`|
|current_total_discounts|`utf8`|
|current_total_price|`utf8`|
|current_total_tax|`utf8`|
|customer_locale|`utf8`|
|discount_codes|`json`|
|email|`utf8`|
|estimated_taxes|`bool`|
|financial_status|`utf8`|
|gateway|`utf8`|
|landing_site|`utf8`|
|name|`utf8`|
|note_attributes|`json`|
|number|`int64`|
|order_number|`int64`|
|order_status_url|`utf8`|
|payment_gateway_names|`list<item: utf8, nullable>`|
|phone|`utf8`|
|presentment_currency|`utf8`|
|processed_at|`timestamp[us, tz=UTC]`|
|processing_method|`utf8`|
|referring_site|`utf8`|
|source_name|`utf8`|
|subtotal_price|`utf8`|
|tags|`utf8`|
|tax_lines|`json`|
|taxes_included|`bool`|
|test|`bool`|
|token|`utf8`|
|total_discounts|`utf8`|
|total_line_items_price|`utf8`|
|total_outstanding|`utf8`|
|total_price|`utf8`|
|total_price_usd|`utf8`|
|total_tax|`utf8`|
|total_tip_received|`utf8`|
|total_weight|`int64`|
|user_id|`int64`|
|customer|`json`|
|discount_applications|`json`|
|fulfillments|`json`|
|line_items|`json`|
|refunds|`json`|
|shipping_lines|`json`|