
# Table: aws_apigatewayv2_domain_name_configurations
The domain name configuration
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_cq_id|uuid|Unique CloudQuery ID of aws_apigatewayv2_domain_names table (FK)|
|api_gateway_domain_name|text|A domain name for the API|
|certificate_arn|text|An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name|
|certificate_name|text|The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name|
|certificate_upload_date|timestamp without time zone|The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded|
|domain_name_status|text|The status of the domain name migration|
|domain_name_status_message|text|An optional text message containing detailed information about status of the domain name migration|
|endpoint_type|text|The endpoint type|
|hosted_zone_id|text|The Amazon Route 53 Hosted Zone ID of the endpoint|
|ownership_verification_certificate_arn|text|The ARN of the public certificate issued by ACM to validate ownership of your custom domain|
|security_policy|text|The Transport Layer Security (TLS) version of the security policy for this domain name|
