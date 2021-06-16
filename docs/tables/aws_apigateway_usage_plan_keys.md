
# Table: aws_apigateway_usage_plan_keys
Represents a usage plan key to identify a plan customer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|usage_plan_id|uuid|Unique ID of aws_apigateway_usage_plans table (FK)|
|resource_id|text|The Id of a usage plan key.|
|name|text|The name of a usage plan key.|
|type|text|The type of a usage plan key. Currently, the valid key type is API_KEY.|
|value|text|The value of a usage plan key.|
