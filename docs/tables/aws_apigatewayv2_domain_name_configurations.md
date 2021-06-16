
# Table: aws_apigatewayv2_domain_name_configurations
The domain name configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|domain_name_id|uuid|Unique ID of aws_apigatewayv2_domain_names table (FK)|
|api_gateway_domain_name|text|A domain name for the API.|
|certificate_arn|text|An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.|
|certificate_name|text|The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.|
|certificate_upload_date|timestamp without time zone|The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded.|
|domain_name_status|text|The status of the domain name migration. The valid values are AVAILABLE and UPDATING. If the status is UPDATING, the domain cannot be modified further until the existing operation is complete. If it is AVAILABLE, the domain can be updated.|
|domain_name_status_message|text|An optional text message containing detailed information about status of the domain name migration.|
|endpoint_type|text|The endpoint type.|
|hosted_zone_id|text|The Amazon Route 53 Hosted Zone ID of the endpoint.|
|security_policy|text|The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are TLS_1_0 and TLS_1_2.|
