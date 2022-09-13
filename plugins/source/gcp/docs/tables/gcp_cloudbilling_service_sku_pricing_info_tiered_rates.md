
# Table: gcp_cloudbilling_service_sku_pricing_info_tiered_rates
The price rate indicating starting usage and its corresponding price
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_sku_pricing_info_cq_id|uuid|Unique CloudQuery ID of gcp_cloudbilling_service_sku_pricing_info table (FK)|
|start_usage_amount|float|Usage is priced at this rate only after this amount|
|unit_price_currency_code|text|The three-letter currency code defined in ISO 4217|
|unit_price_nanos|bigint|Number of nano (10^-9) units of the amount|
|unit_price_units|bigint|The whole units of the amount|
