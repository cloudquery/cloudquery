
# Table: aws_lambda_function_aliases
Provides configuration information about a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_cq_id|uuid|Unique CloudQuery ID of aws_lambda_functions table (FK)|
|function_arn|text|The Amazon Resource Name (ARN) of the lambda function|
|arn|text|The Amazon Resource Name (ARN) of the alias.|
|description|text|A description of the alias.|
|function_version|text|The function version that the alias invokes.|
|name|text|The name of the alias.|
|revision_id|text|A unique identifier that changes when you update the alias.|
|routing_config_additional_version_weights|jsonb|The second version, and the percentage of traffic that's routed to it.|
|url_config_auth_type|text|The type of authentication that your function URL uses|
|url_config_creation_time|timestamp without time zone|When the function URL was created, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).|
|url_config_function_arn|text|The Amazon Resource Name (ARN) of your function.|
|url_config_function_url|text|The HTTP URL endpoint for your function.|
|url_config_last_modified_time|timestamp without time zone|When the function URL configuration was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).|
|url_config_cors|jsonb|The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.|
