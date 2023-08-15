# Table: shopify_abandoned_checkouts

This table shows data for Shopify Abandoned Checkouts.

The primary key for this table is **id**.
It supports incremental syncs based on the **updated_at** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|updated_at (Incremental Key)|`timestamp[us, tz=UTC]`|
|token|`utf8`|
|cart_token|`utf8`|
|email|`utf8`|
|gateway|`utf8`|
|buyer_accepts_marketing|`bool`|
|created_at|`timestamp[us, tz=UTC]`|
|landing_site|`utf8`|
|note_attributes|`json`|
|referring_site|`utf8`|
|shipping_lines|`json`|
|taxes_included|`bool`|
|total_weight|`int64`|
|currency|`utf8`|
|completed_at|`timestamp[us, tz=UTC]`|
|closed_at|`timestamp[us, tz=UTC]`|
|user_id|`int64`|
|customer_locale|`utf8`|
|line_items|`json`|
|name|`utf8`|
|abandoned_checkout_url|`utf8`|
|discount_codes|`json`|
|tax_lines|`json`|
|source_name|`utf8`|
|presentment_currency|`utf8`|
|buyer_accepts_sms_marketing|`bool`|
|total_discounts|`utf8`|
|total_line_items_price|`utf8`|
|total_price|`utf8`|
|total_tax|`utf8`|
|subtotal_price|`utf8`|
|customer|`json`|