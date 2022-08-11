
# Table: aws_cloudfront_distribution_cache_behaviors
A complex type that describes how CloudFront processes requests
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_cq_id|uuid|Unique CloudQuery ID of aws_cloudfront_distributions table (FK)|
|path_pattern|text|The pattern (for example, images/*.jpg) that specifies which requests to apply the behavior to|
|target_origin_id|text|The value of ID for the origin that you want CloudFront to route requests to when they match this cache behavior.|
|viewer_protocol_policy|text|The protocol that viewers can use to access the files in the origin specified by TargetOriginId when a request matches the path pattern in PathPattern|
|allowed_methods|text[]|A complex type that contains the HTTP methods that you want CloudFront to process and forward to your origin.|
|cached_methods|text[]|A complex type that contains the HTTP methods that you want CloudFront to cache responses to.|
|cache_policy_id|text|The unique identifier of the cache policy that is attached to this cache behavior|
|compress|boolean|Whether you want CloudFront to automatically compress certain files for this cache behavior|
|default_ttl|bigint|This field is deprecated|
|field_level_encryption_id|text|The value of ID for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for this cache behavior.|
|forwarded_values_cookies_forward|text|This field is deprecated|
|forwarded_values_cookies_whitelisted_names|text[]|A list of cookie names.|
|forwarded_values_query_string|boolean|This field is deprecated|
|forwarded_values_headers|text[]|A list of HTTP header names.|
|forwarded_values_query_string_cache_keys|text[]|A list that contains the query string parameters that you want CloudFront to use as a basis for caching for a cache behavior|
|max_ttl|bigint|This field is deprecated|
|min_ttl|bigint|This field is deprecated|
|origin_request_policy_id|text|The unique identifier of the origin request policy that is attached to this cache behavior|
|realtime_log_config_arn|text|The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior|
|smooth_streaming|boolean|Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify true; if not, specify false|
|trusted_key_groups_enabled|boolean|This field is true if any of the key groups in the list have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies. If not, this field is false.|
|trusted_key_groups|text[]|A list of key groups identifiers.|
|trusted_signers_enabled|boolean|This field is true if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies|
|trusted_signers|text[]|A list of AWS account identifiers.|
