
# Table: aws_apigatewayv2_api_stages
Represents an API stage.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|stage_name|text|The name of the stage.|
|access_log_settings_destination_arn|text|The ARN of the CloudWatch Logs log group to receive access logs.|
|access_log_settings_format|text|A single line format of the access logs of data, as specified by selected $context variables. The format must include at least $context.requestId.|
|api_gateway_managed|boolean|Specifies whether a stage is managed by API Gateway. If you created an API using quick create, the $default stage is managed by API Gateway. You can't modify the $default stage.|
|auto_deploy|boolean|Specifies whether updates to an API automatically trigger a new deployment. The default value is false.|
|client_certificate_id|text|The identifier of a client certificate for a Stage. Supported only for WebSocket APIs.|
|created_date|timestamp without time zone|The timestamp when the stage was created.|
|route_settings_data_trace_enabled|boolean|Specifies whether (true) or not (false) data trace logging is enabled for this route. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.|
|route_settings_detailed_metrics_enabled|boolean|Specifies whether detailed metrics are enabled.|
|route_settings_logging_level|text|Specifies the logging level for this route: INFO, ERROR, or OFF. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.|
|route_settings_throttling_burst_limit|integer|Specifies the throttling burst limit.|
|route_settings_throttling_rate_limit|float|Specifies the throttling rate limit.|
|deployment_id|text|The identifier of the Deployment that the Stage is associated with. Can't be updated if autoDeploy is enabled.|
|description|text|The description of the stage.|
|last_deployment_status_message|text|Describes the status of the last deployment of a stage. Supported only for stages with autoDeploy enabled.|
|last_updated_date|timestamp without time zone|The timestamp when the stage was last updated.|
|route_settings|jsonb|Route settings for the stage, by routeKey.|
|stage_variables|jsonb|A map that defines the stage variables for a stage resource. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
