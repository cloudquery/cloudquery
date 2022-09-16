
# Table: aws_apigateway_domain_names
Represents a custom domain name as a user-friendly host name of an API (RestApi)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|certificate_arn|text|The reference to an AWS-managed certificate that will be used by edge-optimized endpoint for this domain name|
|certificate_name|text|The name of the certificate that will be used by edge-optimized endpoint for this domain name|
|certificate_upload_date|timestamp without time zone|The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded|
|distribution_domain_name|text|The domain name of the Amazon CloudFront distribution associated with this custom domain name for an edge-optimized endpoint|
|distribution_hosted_zone_id|text|The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint|
|domain_name|text|The custom domain name as an API host name, for example, my-api.example.com|
|domain_name_status|text|The status of the DomainName migration|
|domain_name_status_message|text|An optional text message containing detailed information about status of the DomainName migration|
|endpoint_configuration_types|text[]|A list of endpoint types of an API (RestApi) or its custom domain name (DomainName)|
|endpoint_configuration_vpc_endpoint_ids|text[]|A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes|
|mutual_tls_authentication_truststore_uri|text|An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example s3://bucket-name/key-name|
|mutual_tls_authentication_truststore_version|text|The version of the S3 object that contains your truststore|
|mutual_tls_authentication_truststore_warnings|text[]|A list of warnings that API Gateway returns while processing your truststore|
|ownership_verification_certificate_arn|text|The ARN of the public certificate issued by ACM to validate ownership of your custom domain|
|regional_certificate_arn|text|The reference to an AWS-managed certificate that will be used for validating the regional domain name|
|regional_certificate_name|text|The name of the certificate that will be used for validating the regional domain name|
|regional_domain_name|text|The domain name associated with the regional endpoint for this custom domain name|
|regional_hosted_zone_id|text|The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint|
|security_policy|text|The Transport Layer Security (TLS) version + cipher suite for this DomainName|
|tags|jsonb|The collection of tags|
