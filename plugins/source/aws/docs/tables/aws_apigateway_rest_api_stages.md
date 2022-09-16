
# Table: aws_apigateway_rest_api_stages
Represents a unique identifier for a version of a deployed RestApi that is callable by users
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|access_log_settings_destination_arn|text|The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs|
|access_log_settings_format|text|A single line format of the access logs of data, as specified by selected $context variables|
|cache_cluster_enabled|boolean|Specifies whether a cache cluster is enabled for the stage|
|cache_cluster_size|text|The size of the cache cluster for the stage, if enabled|
|cache_cluster_status|text|The status of the cache cluster for the stage, if enabled|
|canary_settings_deployment_id|text|The ID of the canary deployment|
|canary_settings_percent_traffic|float|The percent (0-100) of traffic diverted to a canary deployment|
|canary_settings_stage_variable_overrides|jsonb|Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary|
|canary_settings_use_stage_cache|boolean|A Boolean flag to indicate whether the canary deployment uses the stage cache or not|
|client_certificate_id|text|The identifier of a client certificate for an API stage|
|created_date|timestamp without time zone|The timestamp when the stage was created|
|deployment_id|text|The identifier of the Deployment that the stage points to|
|description|text|The stage's description|
|documentation_version|text|The version of the associated API documentation|
|last_updated_date|timestamp without time zone|The timestamp when the stage last updated|
|method_settings|jsonb|A map that defines the method settings for a Stage resource|
|stage_name|text|The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway|
|tags|jsonb|The collection of tags|
|tracing_enabled|boolean|Specifies whether active tracing with X-ray is enabled for the Stage|
|variables|jsonb|A map that defines the stage variables for a Stage resource|
|web_acl_arn|text|The ARN of the WebAcl associated with the Stage|
