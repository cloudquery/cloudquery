
# Table: aws_cloudfront_cache_policies
Contains a cache policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|min_ttl|bigint|The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated|
|name|text|A unique name to identify the cache policy|
|comment|text|A comment to describe the cache policy|
|default_ttl|bigint|The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated|
|max_ttl|bigint|The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated|
|cookies_behavior|text|Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin|
|cookies_quantity|integer|The number of cookie names in the Items list|
|cookies|text[]|A list of cookie names|
|enable_accept_encoding_gzip|boolean|A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin|
|headers_behavior|text|Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin|
|headers_quantity|integer|The number of header names in the Items list|
|headers|text[]|A list of HTTP header names|
|query_strings_behavior|text|Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin|
|query_strings_quantity|integer|The number of query string names in the Items list|
|query_strings|text[]|A list of query string names|
|enable_accept_encoding_brotli|boolean|A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin|
|resource_id|text|The unique identifier for the cache policy|
|last_modified_time|timestamp without time zone|The date and time when the cache policy was last modified|
|type|text|The type of cache policy, either managed (created by AWS) or custom (created in this AWS account)|
