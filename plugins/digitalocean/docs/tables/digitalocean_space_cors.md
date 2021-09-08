
# Table: digitalocean_space_cors
 list of elements describing allowed methods for a specific origin.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|space_cq_id|uuid|Unique CloudQuery ID of digitalocean_spaces table (FK)|
|space_name|text|name of the space.|
|allowed_methods|text[]|HTTP methods (e.g. GET) which are allowed from the specified origin.|
|allowed_origins|text[]|orgins from which requests using the specified methods are allowed.|
|allowed_headers|text[]|headers that will be included in the CORS preflight request’s Access-Control-Request-Headers|
|max_age_seconds|integer|access control max age in seconds|
