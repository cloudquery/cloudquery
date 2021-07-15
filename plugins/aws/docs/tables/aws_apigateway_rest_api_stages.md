
# Table: aws_apigateway_rest_api_stages
Represents a unique identifier for a version of a deployed RestApi that is callable by users.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier. This identifier is unique across all of your APIs in API Gateway.|
|access_log_settings_destination_arn|text|The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-.|
|access_log_settings_format|text|A single line format of the access logs of data, as specified by selected $context variables (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least $context.requestId.|
|cache_cluster_enabled|boolean|Specifies whether a cache cluster is enabled for the stage.|
|cache_cluster_size|text|The size of the cache cluster for the stage, if enabled.|
|cache_cluster_status|text|The status of the cache cluster for the stage, if enabled.|
|canary_settings_deployment_id|text|The ID of the canary deployment.|
|canary_settings_percent_traffic|float|The percent (0-100) of traffic diverted to a canary deployment.|
|canary_settings_stage_variable_overrides|jsonb|Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.|
|canary_settings_use_stage_cache|boolean|A Boolean flag to indicate whether the canary deployment uses the stage cache or not.|
|client_certificate_id|text|The identifier of a client certificate for an API stage.|
|created_date|timestamp without time zone|The timestamp when the stage was created.|
|deployment_id|text|The identifier of the Deployment that the stage points to.|
|description|text|The stage's description.|
|documentation_version|text|The version of the associated API documentation.|
|last_updated_date|timestamp without time zone|The timestamp when the stage last updated.|
|method_settings|jsonb|A map that defines the method settings for a Stage resource. Keys (designated as /{method_setting_key below) are method paths defined as {resource_path}/{http_method} for an individual method override, or /\*/\* for overriding all methods in the stage.|
|stage_name|text|The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
|tracing_enabled|boolean|Specifies whether active tracing with X-ray is enabled for the Stage.|
|variables|jsonb|A map that defines the stage variables for a Stage resource. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.|
|web_acl_arn|text|The ARN of the WebAcl associated with the Stage.|
