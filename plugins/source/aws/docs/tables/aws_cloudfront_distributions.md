
# Table: aws_cloudfront_distributions
A summary of the information about a CloudFront distribution.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|tags|jsonb||
|arn|text|The ARN (Amazon Resource Name) for the distribution|
|caller_reference|text|A unique value (for example, a date-time stamp) that ensures that the request can't be replayed|
|comment|text|Any comments you want to include about the distribution|
|cache_behavior_target_origin_id|text|The value of ID for the origin that you want CloudFront to route requests to when they use the default cache behavior.|
|cache_behavior_viewer_protocol_policy|text|The protocol that viewers can use to access the files in the origin specified by TargetOriginId when a request matches the path pattern in PathPattern|
|cache_behavior_allowed_methods|text[]|A complex type that contains the HTTP methods that you want CloudFront to process and forward to your origin.|
|cache_behavior_allowed_methods_cached_methods|text[]|A complex type that contains the HTTP methods that you want CloudFront to cache responses to.|
|cache_behavior_cache_policy_id|text|The unique identifier of the cache policy that is attached to the default cache behavior|
|cache_behavior_compress|boolean|Whether you want CloudFront to automatically compress certain files for this cache behavior|
|cache_behavior_default_ttl|bigint|This field is deprecated|
|cache_behavior_field_level_encryption_id|text|The value of ID for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for the default cache behavior.|
|cache_behavior_forwarded_values_cookies_forward|text|This field is deprecated|
|cache_behavior_forwarded_values_cookies_whitelisted_names|text[]|A list of cookie names.|
|cache_behavior_forwarded_values_query_string|boolean|This field is deprecated|
|cache_behavior_forwarded_values_headers|text[]|A list of HTTP header names.|
|cache_behavior_forwarded_values_query_string_cache_keys|text[]|A list that contains the query string parameters that you want CloudFront to use as a basis for caching for a cache behavior|
|cache_behavior_max_ttl|bigint|This field is deprecated|
|cache_behavior_min_ttl|bigint|This field is deprecated|
|cache_behavior_origin_request_policy_id|text|The unique identifier of the origin request policy that is attached to the default cache behavior|
|cache_behavior_realtime_log_config_arn|text|The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior|
|cache_behavior_smooth_streaming|boolean|Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify true; if not, specify false|
|cache_behavior_trusted_key_groups_enabled|boolean|This field is true if any of the key groups in the list have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies. If not, this field is false.|
|cache_behavior_trusted_key_groups|text[]|A list of key groups identifiers.|
|cache_behavior_trusted_signers_enabled|boolean|This field is true if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies|
|cache_behavior_trusted_signers|text[]|A list of AWS account identifiers.|
|enabled|boolean|From this field, you can enable or disable the selected distribution.|
|aliases|text[]|A complex type that contains the CNAME aliases, if any, that you want to associate with this distribution.|
|default_root_object|text|The object that you want CloudFront to request from your origin (for example, index.html) when a viewer requests the root URL for your distribution (http://www.example.com) instead of an object in your distribution (http://www.example.com/product-description.html)|
|http_version|text|(Optional) Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront|
|ipv6_enabled|boolean|If you want CloudFront to respond to IPv6 DNS requests with an IPv6 address for your distribution, specify true|
|logging_bucket|text|The Amazon S3 bucket to store the access logs in, for example, myawslogbucket.s3.amazonaws.com.|
|logging_enabled|boolean|Specifies whether you want CloudFront to save access logs to an Amazon S3 bucket|
|logging_include_cookies|boolean|Specifies whether you want CloudFront to include cookies in access logs, specify true for IncludeCookies|
|logging_prefix|text|An optional string that you want CloudFront to prefix to the access log filenames for this distribution, for example, myprefix/|
|price_class|text|The price class that corresponds with the maximum price that you want to pay for CloudFront service|
|geo_restriction_type|text|The method that you want to use to restrict distribution of your content by country:  * none: No geo restriction is enabled, meaning access to content is not restricted by client geo location.  * blacklist: The Location elements specify the countries in which you don't want CloudFront to distribute your content.  * whitelist: The Location elements specify the countries in which you want CloudFront to distribute your content.|
|geo_restrictions|text[]|A complex type that contains a Location element for each country in which you want CloudFront either to distribute your content (whitelist) or not distribute your content (blacklist)|
|viewer_certificate_acm_certificate_arn|text|If the distribution uses Aliases (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in AWS Certificate Manager (ACM) (https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html), provide the Amazon Resource Name (ARN) of the ACM certificate|
|viewer_certificate|text|This field is deprecated|
|viewer_certificate_source|text|This field is deprecated|
|viewer_certificate_cloudfront_default_certificate|boolean|If the distribution uses the CloudFront domain name such as d111111abcdef8.cloudfront.net, set this field to true|
|viewer_certificate_iam_certificate_id|text|If the distribution uses Aliases (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in AWS Identity and Access Management (AWS IAM) (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html), provide the ID of the IAM certificate|
|viewer_certificate_minimum_protocol_version|text|If the distribution uses Aliases (alternate domain names or CNAMEs), specify the security policy that you want CloudFront to use for HTTPS connections with viewers|
|viewer_certificate_ssl_support_method|text|If the distribution uses Aliases (alternate domain names or CNAMEs), specify which viewers the distribution accepts HTTPS connections from.  * sni-only – The distribution accepts HTTPS connections from only viewers that support server name indication (SNI) (https://en.wikipedia.org/wiki/Server_Name_Indication). This is recommended|
|web_acl_id|text|A unique identifier that specifies the AWS WAF web ACL, if any, to associate with this distribution|
|domain_name|text|The domain name corresponding to the distribution, for example, d111111abcdef8.cloudfront.net.|
|id|text|The identifier for the distribution|
|in_progress_invalidation_batches|integer|The number of invalidation batches currently in progress.|
|last_modified_time|timestamp without time zone|The date and time the distribution was last modified.|
|status|text|This response element indicates the current status of the distribution|
|active_trusted_key_groups_enabled|boolean|This field is true if any of the key groups have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies|
|active_trusted_key_groups|jsonb|A list of key groups, including the identifiers of the public keys in each key group that CloudFront can use to verify the signatures of signed URLs and signed cookies.|
|active_trusted_signers_enabled|boolean|This field is true if any of the AWS accounts in the list have active CloudFront key pairs that CloudFront can use to verify the signatures of signed URLs and signed cookies|
|active_trusted_signers|jsonb|A list of AWS accounts and the identifiers of active CloudFront key pairs in each account that CloudFront can use to verify the signatures of signed URLs and signed cookies.|
|alias_icp_recordals|jsonb|AWS services in China customers must file for an Internet Content Provider (ICP) recordal if they want to serve content publicly on an alternate domain name, also known as a CNAME, that they've added to CloudFront|
