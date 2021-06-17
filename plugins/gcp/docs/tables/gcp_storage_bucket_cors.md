
# Table: gcp_storage_bucket_cors
The bucket's Cross-Origin Resource Sharing (CORS) configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_id|uuid|Unique ID of gcp_storage_buckets table (FK)|
|max_age_seconds|bigint|The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses|
|method|text[]|The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method"|
|origin|text[]|The list of Origins eligible to receive CORS response headers Note: "*" is permitted in the list of origins, and means "any Origin"|
|response_header|text[]|The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains|
