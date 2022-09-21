
# Table: azure_cdn_profile_endpoint_geo_filters
GeoFilter rules defining user's geo access within a CDN endpoint
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_endpoint_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)|
|relative_path|text|Relative path applicable to geo filter|
|action|text|Action of the geo filter, ie|
|country_codes|text[]|Two letter country codes defining user country access in a geo filter, eg|
