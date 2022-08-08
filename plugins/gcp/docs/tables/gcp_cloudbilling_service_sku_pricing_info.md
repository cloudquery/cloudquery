
# Table: gcp_cloudbilling_service_sku_pricing_info
Represents the pricing information for a SKU at a single point of time
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_sku_cq_id|uuid|Unique CloudQuery ID of gcp_cloudbilling_service_skus table (FK)|
|aggregation_count|bigint|The number of intervals to aggregate over|
|aggregation_interval|text|"AGGREGATION_INTERVAL_UNSPECIFIED"   "DAILY"   "MONTHLY"|
|aggregation_level|text|"AGGREGATION_LEVEL_UNSPECIFIED"   "ACCOUNT"   "PROJECT"|
|currency_conversion_rate|float|Conversion rate used for currency conversion, from USD to the currency specified in the request|
|effective_time|text|The timestamp from which this pricing was effective within the requested time range|
|base_unit|text|The base unit for the SKU which is the unit used in usage exports|
|base_unit_conversion_factor|float|Conversion factor for converting from price per usage_unit to price per base_unit, and start_usage_amount to start_usage_amount in base_unit|
|base_unit_description|text|The base unit in human readable form|
|display_quantity|float|The recommended quantity of units for displaying pricing info|
|usage_unit|text|The short hand for unit of usage this pricing is specified in|
|usage_unit_description|text|"gibi byte"|
|summary|text|An optional human readable summary of the pricing information, has a maximum length of 256 characters|
