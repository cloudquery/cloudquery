
# Table: aws_apigatewayv2_domain_names
Represents a domain name.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|domain_name|text|The name of the DomainName resource.|
|api_mapping_selection_expression|text|The API mapping selection expression.|
|mutual_tls_authentication_truststore_uri|text|An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.|
|mutual_tls_authentication_truststore_version|text|The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.|
|mutual_tls_authentication_truststore_warnings|text[]|A list of warnings that API Gateway returns while processing your truststore. Invalid certificates produce warnings. Mutual TLS is still enabled, but some clients might not be able to access your API. To resolve warnings, upload a new truststore to S3, and then update you domain name to use the new version.|
|tags|jsonb|The collection of tags associated with a domain name.|
