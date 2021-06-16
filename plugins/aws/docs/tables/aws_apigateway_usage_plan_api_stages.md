
# Table: aws_apigateway_usage_plan_api_stages
API stage name of the associated API stage in a usage plan.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|usage_plan_id|uuid|Unique ID of aws_apigateway_usage_plans table (FK)|
|api_id|text|API Id of the associated API stage in a usage plan.|
|stage|text|API stage name of the associated API stage in a usage plan.|
|throttle|jsonb|Map containing method level throttling information for API stage in a usage plan.|
