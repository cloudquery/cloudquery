
# Table: azure_storage_blob_service_cors_rules
CorsRule specifies a CORS rule for the Blob service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|blob_service_cq_id|uuid|Unique CloudQuery ID of azure_storage_blob_services table (FK)|
|blob_service_id|text|Unique ID of azure_storage_blob_services table (FK)|
|allowed_origins|text[]|A list of origin domains that will be allowed via CORS, or "*" to allow all domains|
|allowed_methods|text[]|A list of HTTP methods that are allowed to be executed by the origin|
|max_age_in_seconds|integer|The number of seconds that the client/browser should cache a preflight response|
|exposed_headers|text[]|A list of response headers to expose to CORS clients|
|allowed_headers|text[]|A list of headers allowed to be part of the cross-origin request|
