
# Table: aws_elbv2_load_balancer_attributes
Load balancer attributes
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_cq_id|uuid|Unique CloudQuery ID of aws_elbv2_load_balancers table (FK)|
|access_logs_s3_enabled|boolean|Indicates whether access logs stored in Amazon S3 are enabled.|
|access_logs_s3_bucket|text|The name of the Amazon S3 bucket for the access logs.|
|access_logs_s3_prefix|text|The prefix for the location in the Amazon S3 bucket.|
|deletion_protection|boolean|Indicates whether deletion protection is enabled.|
|idle_timeout|integer|The idle timeout value, in seconds.|
|routing_http_desync_mitigation_mode|text|Determines how the load balancer handles requests that might pose a security risk to your application.|
|routing_http_drop_invalid_header_fields|boolean|Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer.|
|routing_http_xamzntls_enabled|boolean|Indicates whether the two headers (x-amzn-tls-{version,cipher-suite}) are added to the client request before sending it to the target.|
|routing_http_xff_client_port|boolean|Indicates whether the X-Forwarded-For header should preserve the source port that the client used to connect to the load balancer.|
|routing_http2|boolean|Indicates whether HTTP/2 is enabled.|
|waf_fail_open|boolean|Indicates whether to allow a AWS WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF.|
|load_balancing_cross_zone|boolean|Indicates whether cross-zone load balancing is enabled.|
