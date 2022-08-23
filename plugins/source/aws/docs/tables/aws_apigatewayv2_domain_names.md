
# Table: aws_apigatewayv2_domain_names
Represents a domain name
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|domain_name|text|The name of the DomainName resource|
|api_mapping_selection_expression|text|The API mapping selection expression|
|mutual_tls_authentication_truststore_uri|text|An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name|
|mutual_tls_authentication_truststore_version|text|The version of the S3 object that contains your truststore|
|mutual_tls_authentication_truststore_warnings|text[]|A list of warnings that API Gateway returns while processing your truststore|
|tags|jsonb|The collection of tags associated with a domain name|
