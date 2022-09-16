
# Table: aws_apigateway_rest_apis
Represents a REST API
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|api_key_source|text|The source of the API key for metering requests according to a usage plan|
|binary_media_types|text[]|The list of binary media types supported by the RestApi|
|created_date|timestamp without time zone|The timestamp when the API was created|
|description|text|The API's description|
|disable_execute_api_endpoint|boolean|Specifies whether clients can invoke your API by using the default execute-api endpoint|
|endpoint_configuration_types|text[]|A list of endpoint types of an API (RestApi) or its custom domain name (DomainName)|
|endpoint_configuration_vpc_endpoint_ids|text[]|A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes|
|id|text|The API's identifier|
|minimum_compression_size|bigint|A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API|
|name|text|The API's name|
|policy|text|A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration|
|tags|jsonb|The collection of tags|
|version|text|A version identifier for the API|
|warnings|text[]|The warning messages reported when failonwarnings is turned on during API import|
