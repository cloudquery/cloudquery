
# Table: gcp_cloudbilling_service_skus
Encapsulates a single SKU in Google Cloud Platform
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_cloudbilling_services table (FK)|
|resource_family|text|The type of product the SKU refers to|
|resource_group|text|A group classification for related SKUs|
|service_display_name|text|The display name of the service this SKU belongs to|
|usage_type|text|Represents how the SKU is consumed|
|description|text|A human readable description of the SKU, has a maximum length of 256 characters|
|geo_taxonomy_regions|text[]|The list of regions associated with a sku|
|geo_taxonomy_type|text|"TYPE_UNSPECIFIED" - The type is not specified   "GLOBAL" - The sku is global in nature, eg|
|name|text|The resource name for the SKU|
|service_provider_name|text|Identifies the service provider|
|service_regions|text[]|"asia-east1" Service regions can be found at https://cloudgooglecom/about/locations/|
|sku_id|text|The identifier for the SKU|
