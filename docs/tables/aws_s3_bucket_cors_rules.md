
# Table: aws_s3_bucket_cors_rules
Specifies a cross-origin access rule for an Amazon S3 bucket.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_id|uuid|Unique ID of aws_s3_buckets table (FK)|
|allowed_methods|text[]|An HTTP method that you allow the origin to execute|
|allowed_origins|text[]|One or more origins you want customers to be able to access the bucket from.  |
|allowed_headers|text[]|Headers that are specified in the Access-Control-Request-Headers header|
|expose_headers|text[]|One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).|
|resource_id|text|Unique identifier for the rule|
|max_age_seconds|integer|The time in seconds that your browser is to cache the preflight response for the specified resource.|
