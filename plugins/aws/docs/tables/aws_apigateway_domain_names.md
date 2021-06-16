
# Table: aws_apigateway_domain_names
Represents a custom domain name as a user-friendly host name of an API (RestApi).
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|certificate_arn|text|The reference to an AWS-managed certificate that will be used by edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.|
|certificate_name|text|The name of the certificate that will be used by edge-optimized endpoint for this domain name.|
|certificate_upload_date|timestamp without time zone|The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded.|
|distribution_domain_name|text|The domain name of the Amazon CloudFront distribution associated with this custom domain name for an edge-optimized endpoint. You set up this association when adding a DNS record pointing the custom domain name to this distribution name. For more information about CloudFront distributions, see the Amazon CloudFront documentation (https://aws.amazon.com/documentation/cloudfront/).|
|distribution_hosted_zone_id|text|The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint. The valid value is Z2FDTNDATAQYW2 for all the regions. For more information, see Set up a Regional Custom Domain Name (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html) and AWS Regions and Endpoints for API Gateway (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).|
|domain_name|text|The custom domain name as an API host name, for example, my-api.example.com.|
|domain_name_status|text|The status of the DomainName migration. The valid values are AVAILABLE and UPDATING. If the status is UPDATING, the domain cannot be modified further until the existing operation is complete. If it is AVAILABLE, the domain can be updated.|
|domain_name_status_message|text|An optional text message containing detailed information about status of the DomainName migration.|
|endpoint_configuration_types|text[]|A list of endpoint types of an API (RestApi) or its custom domain name (DomainName). For an edge-optimized API and its custom domain name, the endpoint type is "EDGE". For a regional API and its custom domain name, the endpoint type is REGIONAL. For a private API, the endpoint type is PRIVATE.|
|endpoint_configuration_vpc_endpoint_ids|text[]|A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes. It is only supported for PRIVATE endpoint type.|
|mutual_tls_authentication_truststore_uri|text|An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.|
|mutual_tls_authentication_truststore_version|text|The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.|
|mutual_tls_authentication_truststore_warnings|text[]|A list of warnings that API Gateway returns while processing your truststore. Invalid certificates produce warnings. Mutual TLS is still enabled, but some clients might not be able to access your API. To resolve warnings, upload a new truststore to S3, and then update you domain name to use the new version.|
|regional_certificate_arn|text|The reference to an AWS-managed certificate that will be used for validating the regional domain name. AWS Certificate Manager is the only supported source.|
|regional_certificate_name|text|The name of the certificate that will be used for validating the regional domain name.|
|regional_domain_name|text|The domain name associated with the regional endpoint for this custom domain name. You set up this association by adding a DNS record that points the custom domain name to this regional domain name. The regional domain name is returned by API Gateway when you create a regional endpoint.|
|regional_hosted_zone_id|text|The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint. For more information, see Set up a Regional Custom Domain Name (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html) and AWS Regions and Endpoints for API Gateway (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).|
|security_policy|text|The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are TLS_1_0 and TLS_1_2.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
