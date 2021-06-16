
# Table: aws_lambda_function_aliases
Provides configuration information about a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html). 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid|Unique ID of aws_lambda_functions table (FK)|
|alias_arn|text|The Amazon Resource Name (ARN) of the alias.|
|description|text|A description of the alias.|
|function_version|text|The function version that the alias invokes.|
|name|text|The name of the alias.|
|revision_id|text|A unique identifier that changes when you update the alias.|
|routing_config_additional_version_weights|jsonb|The second version, and the percentage of traffic that's routed to it.|
