
# Table: aws_apigatewayv2_api_stages
Represents an API stage
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_cq_id|uuid|Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)|
|api_id|text|The API id|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|stage_name|text|The name of the stage|
|access_log_settings_destination_arn|text|The ARN of the CloudWatch Logs log group to receive access logs|
|access_log_settings_format|text|A single line format of the access logs of data, as specified by selected $context variables|
|api_gateway_managed|boolean|Specifies whether a stage is managed by API Gateway|
|auto_deploy|boolean|Specifies whether updates to an API automatically trigger a new deployment|
|client_certificate_id|text|The identifier of a client certificate for a Stage|
|created_date|timestamp without time zone|The timestamp when the stage was created|
|route_settings_data_trace_enabled|boolean|Specifies whether (true) or not (false) data trace logging is enabled for this route|
|route_settings_detailed_metrics_enabled|boolean|Specifies whether detailed metrics are enabled|
|route_settings_logging_level|text|Specifies the logging level for this route: INFO, ERROR, or OFF|
|route_settings_throttling_burst_limit|bigint|Specifies the throttling burst limit|
|route_settings_throttling_rate_limit|float|Specifies the throttling rate limit|
|deployment_id|text|The identifier of the Deployment that the Stage is associated with|
|description|text|The description of the stage|
|last_deployment_status_message|text|Describes the status of the last deployment of a stage|
|last_updated_date|timestamp without time zone|The timestamp when the stage was last updated|
|route_settings|jsonb|Route settings for the stage, by routeKey|
|stage_variables|jsonb|A map that defines the stage variables for a stage resource|
|tags|jsonb|The collection of tags|
