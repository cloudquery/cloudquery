
# Table: aws_apigateway_usage_plan_api_stages
API stage name of the associated API stage in a usage plan.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|usage_plan_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_usage_plans table (FK)|
|usage_plan_id|text|The identifier of a UsagePlan resource.|
|api_id|text|API Id of the associated API stage in a usage plan.|
|stage|text|API stage name of the associated API stage in a usage plan.|
|throttle|jsonb|Map containing method level throttling information for API stage in a usage plan.|
