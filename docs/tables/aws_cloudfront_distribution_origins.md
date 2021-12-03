
# Table: aws_cloudfront_distribution_origins
An origin
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_cq_id|uuid|Unique CloudQuery ID of aws_cloudfront_distributions table (FK)|
|domain_name|text|The domain name for the origin|
|id|text|A unique identifier for the origin|
|connection_attempts|integer|The number of times that CloudFront attempts to connect to the origin|
|connection_timeout|integer|The number of seconds that CloudFront waits when trying to establish a connection to the origin|
|custom_headers|jsonb|A list of HTTP header names and values that CloudFront adds to the requests that it sends to the origin|
|custom_origin_config_http_port|integer|The HTTP port that CloudFront uses to connect to the origin|
|custom_origin_config_https_port|integer|The HTTPS port that CloudFront uses to connect to the origin|
|custom_origin_config_protocol_policy|text|Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin|
|custom_origin_config_keepalive_timeout|integer|Specifies how long, in seconds, CloudFront persists its connection to the origin|
|custom_origin_config_read_timeout|integer|Specifies how long, in seconds, CloudFront waits for a response from the origin. This is also known as the origin response timeout|
|custom_origin_config_ssl_protocols|text[]|A list that contains allowed SSL/TLS protocols for this distribution.|
|origin_path|text|An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin|
|origin_shield_enabled|boolean|A flag that specifies whether Origin Shield is enabled|
|origin_shield_region|text|The AWS Region for Origin Shield|
|s3_origin_config_origin_access_identity|text|The CloudFront origin access identity to associate with the origin|
