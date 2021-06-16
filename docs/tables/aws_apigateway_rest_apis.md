
# Table: aws_apigateway_rest_apis
Represents a REST API.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|api_key_source|text|The source of the API key for metering requests according to a usage plan. Valid values are:|
|binary_media_types|text[]|The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.|
|created_date|timestamp without time zone|The timestamp when the API was created.|
|description|text|The API's description.|
|disable_execute_api_endpoint|boolean|Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.|
|endpoint_configuration_types|text[]|A list of endpoint types of an API (RestApi) or its custom domain name (DomainName). For an edge-optimized API and its custom domain name, the endpoint type is "EDGE". For a regional API and its custom domain name, the endpoint type is REGIONAL. For a private API, the endpoint type is PRIVATE.|
|endpoint_configuration_vpc_endpoint_ids|text[]|A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes. It is only supported for PRIVATE endpoint type.|
|resource_id|text|The API's identifier. This identifier is unique across all of your APIs in API Gateway.|
|minimum_compression_size|integer|A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.|
|name|text|The API's name.|
|policy|text|A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
|version|text|A version identifier for the API.|
|warnings|text[]|The warning messages reported when failonwarnings is turned on during API import.|
