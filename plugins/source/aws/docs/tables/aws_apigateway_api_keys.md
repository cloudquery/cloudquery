
# Table: aws_apigateway_api_keys
A resource that can be distributed to callers for executing Method resources that require an API key.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|created_date|timestamp without time zone|The timestamp when the API Key was created.|
|customer_id|text|An AWS Marketplace customer identifier , when integrating with the AWS SaaS Marketplace.|
|description|text|The description of the API Key.|
|enabled|boolean|Specifies whether the API Key can be used by callers.|
|id|text|The identifier of the API Key.|
|last_updated_date|timestamp without time zone|The timestamp when the API Key was last updated.|
|name|text|The name of the API Key.|
|stage_keys|text[]|A list of Stage resources that are associated with the ApiKey resource.|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
|value|text|The value of the API Key.|
